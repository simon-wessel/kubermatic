//go:build ee

/*
                  Kubermatic Enterprise Read-Only License
                         Version 1.0 ("KERO-1.0”)
                     Copyright © 2021 Loodse GmbH

   1.	You may only view, read and display for studying purposes the source
      code of the software licensed under this license, and, to the extent
      explicitly provided under this license, the binary code.
   2.	Any use of the software which exceeds the foregoing right, including,
      without limitation, its execution, compilation, copying, modification
      and distribution, is expressly prohibited.
   3.	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
      EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
      MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
      IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
      CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
      TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
      SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

   END OF TERMS AND CONDITIONS
*/

package metering

import (
	"context"
	"testing"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/provider"

	appsv1 "k8s.io/api/apps/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestReconcileMeteringResources(t *testing.T) {
	testCases := []struct {
		name   string
		client ctrlruntimeclient.Client
		seed   *kubermaticv1.Seed
	}{
		{
			name: "scenario 1: default settings",
			client: fakectrlruntimeclient.
				NewClientBuilder().
				WithScheme(scheme.Scheme).
				Build(),
			seed: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name:      provider.DefaultSeedName,
					Namespace: "kubermatic",
				},
				Spec: kubermaticv1.SeedSpec{
					Metering: &kubermaticv1.MeteringConfigurations{
						Enabled:          true,
						StorageClassName: "kubermatic-fast",
						StorageSize:      "100Gi",
					},
				},
			},
		},
		{
			name: "scenario 2: custom storage size and class",
			client: fakectrlruntimeclient.
				NewClientBuilder().
				WithScheme(scheme.Scheme).
				Build(),
			seed: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name:      provider.DefaultSeedName,
					Namespace: "kubermatic",
				},
				Spec: kubermaticv1.SeedSpec{
					Metering: &kubermaticv1.MeteringConfigurations{
						Enabled:          true,
						StorageClassName: "mystorageclass",
						StorageSize:      "75Gi",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			// Reconcile
			if err := ReconcileMeteringResources(ctx, tc.client, tc.seed); err != nil {
				t.Fatalf("reconciling failed: %v", err)
			}

			// Check CronJob
			cronJobList := &batchv1beta1.CronJobList{}
			if err := tc.client.List(ctx, cronJobList); err != nil {
				t.Fatalf("Error listing cronjobs: %v", err)
			}
			if len(cronJobList.Items) != 1 {
				t.Fatalf("Expected exactly one cronjob, got %v", len(cronJobList.Items))
			}
			if len(cronJobList.Items[0].Spec.JobTemplate.Spec.Template.Spec.InitContainers) != 1 {
				t.Fatalf("Expected exactly one init container, got %v", len(cronJobList.Items[0].Spec.JobTemplate.Spec.Template.Spec.InitContainers))
			}
			if len(cronJobList.Items[0].Spec.JobTemplate.Spec.Template.Spec.Containers) != 2 {
				t.Fatalf("Expected exactly two containers, got %v", len(cronJobList.Items[0].Spec.JobTemplate.Spec.Template.Spec.Containers))
			}

			// Check Service Account
			serviceAccountList := &corev1.ServiceAccountList{}
			if err := tc.client.List(ctx, serviceAccountList); err != nil {
				t.Fatalf("Error listing service accounts: %v", err)
			}
			if len(serviceAccountList.Items) != 1 {
				t.Fatalf("Expected exactly one PVC, got %v", len(cronJobList.Items))
			}

			// Check PVC
			pvcList := &corev1.PersistentVolumeClaimList{}
			if err := tc.client.List(ctx, pvcList); err != nil {
				t.Fatalf("Error listing pvcs: %v", err)
			}
			if len(pvcList.Items) != 1 {
				t.Fatalf("Expected exactly one PVC, got %v", len(cronJobList.Items))
			}
			pvc := pvcList.Items[0]
			if *pvc.Spec.StorageClassName != tc.seed.Spec.Metering.StorageClassName {
				t.Fatalf("Expected storageClassName to be %s, but was %s",
					tc.seed.Spec.Metering.StorageClassName,
					*pvc.Spec.StorageClassName)
			}
			storage := *pvc.Spec.Resources.Requests.Storage()
			if storage.String() != tc.seed.Spec.Metering.StorageSize {
				t.Fatalf("Expected pvc storage size to be %s, but was %s",
					tc.seed.Spec.Metering.StorageSize,
					storage.String())
			}

			// Check Cluster Role Binding
			clusterRoleBindingList := &rbacv1.ClusterRoleBindingList{}
			if err := tc.client.List(ctx, clusterRoleBindingList); err != nil {
				t.Fatalf("Error listing cluster role bindings: %v", err)
			}
			if len(clusterRoleBindingList.Items) != 1 {
				t.Fatalf("Expected exactly one cluster role binding, got %v", len(cronJobList.Items))
			}
			clusterRoleBinding := clusterRoleBindingList.Items[0]
			if clusterRoleBinding.RoleRef.Name != "cluster-admin" {
				t.Fatalf("Expected cluster role to be 'cluster admin', got '%v'",
					clusterRoleBinding.RoleRef.Name)
			}

			// Check Deployment
			deploymentList := &appsv1.DeploymentList{}
			if err := tc.client.List(ctx, deploymentList); err != nil {
				t.Fatalf("Error listing deployments: %v", err)
			}
			if len(deploymentList.Items) != 1 {
				t.Fatalf("Expected exactly one deployment, got %v", len(cronJobList.Items))
			}
		})
	}
}
