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

package stack

import (
	"context"

	"github.com/sirupsen/logrus"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/install/helm"
	"k8c.io/kubermatic/v2/pkg/provider"
	"k8c.io/kubermatic/v2/pkg/util/yamled"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type DeployOptions struct {
	HelmClient                 helm.Client
	HelmValues                 *yamled.Document
	KubeClient                 ctrlruntimeclient.Client
	StorageClassProvider       string
	KubermaticConfiguration    *kubermaticv1.KubermaticConfiguration
	RawKubermaticConfiguration *unstructured.Unstructured
	ForceHelmReleaseUpgrade    bool
	ChartsDirectory            string

	SeedsGetter      provider.SeedsGetter
	SeedClientGetter provider.SeedClientGetter

	Logger                             *logrus.Entry
	EnableCertManagerV2Migration       bool
	EnableCertManagerUpstreamMigration bool
	EnableNginxIngressMigration        bool
	EnableOpenstackCSIDriverMigration  bool
	EnableLogrotateMigration           bool
	DisableTelemetry                   bool
}

type Stack interface {
	Name() string
	ValidateConfiguration(config *kubermaticv1.KubermaticConfiguration, helmValues *yamled.Document, opt DeployOptions, logger logrus.FieldLogger) (*kubermaticv1.KubermaticConfiguration, *yamled.Document, []error)
	ValidateState(ctx context.Context, opt DeployOptions) []error
	Deploy(ctx context.Context, opt DeployOptions) error
}
