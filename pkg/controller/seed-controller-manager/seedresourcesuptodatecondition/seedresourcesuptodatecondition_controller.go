/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package seedresourcesuptodatecondition

import (
	"context"
	"fmt"
	"reflect"

	"go.uber.org/zap"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	kubermaticv1helper "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1/helper"
	controllerutil "k8c.io/kubermatic/v2/pkg/controller/util"
	"k8c.io/kubermatic/v2/pkg/version/kubermatic"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const ControllerName = "seed_resources_up_to_date_condition_controller"

type reconciler struct {
	log        *zap.SugaredLogger
	client     ctrlruntimeclient.Client
	recorder   record.EventRecorder
	workerName string
	versions   kubermatic.Versions
}

func Add(
	ctx context.Context,
	log *zap.SugaredLogger,
	mgr manager.Manager,
	numWorkers int,
	workerName string,
	versions kubermatic.Versions,
) error {
	r := &reconciler{
		log:        log.Named(ControllerName),
		client:     mgr.GetClient(),
		recorder:   mgr.GetEventRecorderFor(ControllerName),
		workerName: workerName,
		versions:   versions,
	}

	ctrlOptions := controller.Options{
		Reconciler:              r,
		MaxConcurrentReconciles: numWorkers,
	}
	c, err := controller.New(ControllerName, mgr, ctrlOptions)
	if err != nil {
		return err
	}

	typesToWatch := []ctrlruntimeclient.Object{
		&appsv1.Deployment{},
		&appsv1.StatefulSet{},
	}
	for _, t := range typesToWatch {
		if err := c.Watch(&source.Kind{Type: t}, controllerutil.EnqueueClusterForNamespacedObject(mgr.GetClient())); err != nil {
			return fmt.Errorf("failed to create watch for %T: %w", t, err)
		}
	}

	return c.Watch(&source.Kind{Type: &kubermaticv1.Cluster{}}, &handler.EnqueueRequestForObject{})
}

func (r *reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	cluster := &kubermaticv1.Cluster{}
	if err := r.client.Get(ctx, request.NamespacedName, cluster); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, fmt.Errorf("failed to get cluster %q: %w", request.Name, err)
	}

	// Add a wrapping here so we can emit an event on error
	err := r.reconcile(ctx, cluster)
	if err != nil {
		r.log.With("cluster", request.Name).Errorw("Failed to reconcile cluster", zap.Error(err))
		r.recorder.Event(cluster, corev1.EventTypeWarning, "ReconcilingError", err.Error())
	}
	return reconcile.Result{}, err
}

func (r *reconciler) reconcile(ctx context.Context, cluster *kubermaticv1.Cluster) error {
	if r.workerName != cluster.Labels[kubermaticv1.WorkerNameLabelKey] {
		return nil
	}

	if cluster.Spec.Pause {
		return nil
	}

	upToDate, err := r.seedResourcesUpToDate(ctx, cluster)
	if err != nil {
		return err
	}

	oldCluster := cluster.DeepCopy()
	if !upToDate {
		kubermaticv1helper.SetClusterCondition(
			cluster,
			r.versions,
			kubermaticv1.ClusterConditionSeedResourcesUpToDate,
			corev1.ConditionFalse,
			kubermaticv1.ReasonClusterUpdateSuccessful,
			"Some control plane components did not finish updating",
		)
	} else {
		kubermaticv1helper.SetClusterCondition(
			cluster,
			r.versions,
			kubermaticv1.ClusterConditionSeedResourcesUpToDate,
			corev1.ConditionTrue,
			kubermaticv1.ReasonClusterUpdateSuccessful,
			"All control plane components are up to date",
		)
	}
	if reflect.DeepEqual(oldCluster, cluster) {
		return nil
	}

	return r.client.Status().Patch(ctx, cluster, ctrlruntimeclient.MergeFrom(oldCluster))
}

func (r *reconciler) seedResourcesUpToDate(ctx context.Context, cluster *kubermaticv1.Cluster) (bool, error) {
	listOpts := &ctrlruntimeclient.ListOptions{Namespace: cluster.Status.NamespaceName}

	statefulSets := &appsv1.StatefulSetList{}
	if err := r.client.List(ctx, statefulSets, listOpts); err != nil {
		return false, fmt.Errorf("failed to list statefulSets: %w", err)
	}
	for _, statefulSet := range statefulSets.Items {
		if statefulSet.Spec.Replicas == nil {
			return false, nil
		}
		if *statefulSet.Spec.Replicas != statefulSet.Status.UpdatedReplicas ||
			*statefulSet.Spec.Replicas != statefulSet.Status.CurrentReplicas ||
			*statefulSet.Spec.Replicas != statefulSet.Status.ReadyReplicas {
			return false, nil
		}
	}

	deployments := &appsv1.DeploymentList{}
	if err := r.client.List(ctx, deployments, listOpts); err != nil {
		return false, fmt.Errorf("failed to list deployments: %w", err)
	}

	for _, deployment := range deployments.Items {
		if deployment.Spec.Replicas == nil {
			return false, nil
		}
		if *deployment.Spec.Replicas != deployment.Status.UpdatedReplicas ||
			*deployment.Spec.Replicas != deployment.Status.AvailableReplicas ||
			*deployment.Spec.Replicas != deployment.Status.ReadyReplicas {
			return false, nil
		}
	}

	// This is to avoid setting the resource up-to-date condition in the
	// initial stage when deploymens and statefulsets are not yet deployed.
	// TODO This is not perfect as we may endup in a situation where
	// the available control plane components are ready, but not all components have
	// been deployed yet. This scenario is quite unlikely to happen though and
	// the impact of having the condition set is not big.
	return len(deployments.Items) > 0 || len(statefulSets.Items) > 0, nil
}
