package packet

import (
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	kuberneteshelper "github.com/kubermatic/kubermatic/api/pkg/kubernetes"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
)

const (
	defaultBillingCycle = "hourly"
)

type packet struct {
}

// NewCloudProvider creates a new packet provider.
func NewCloudProvider() provider.CloudProvider {
	return &packet{}
}

// DefaultCloudSpec adds defaults to the CloudSpec.
func (p *packet) DefaultCloudSpec(spec *kubermaticv1.CloudSpec) error {
	return nil
}

// ValidateCloudSpec validates the given CloudSpec.
func (p *packet) ValidateCloudSpec(spec kubermaticv1.CloudSpec) error {
	if spec.Packet.APIKey == "" {
		if err := kuberneteshelper.ValidateSecretKeySelector(spec.Packet.APIKeyReference, "apiKey"); err != nil {
			return err
		}
	}

	if spec.Packet.ProjectID == "" {
		if err := kuberneteshelper.ValidateSecretKeySelector(spec.Packet.ProjectIDReference, "projectID"); err != nil {
			return err
		}
	}

	return nil
}

// InitializeCloudProvider initializes a cluster, in particular
// updates BillingCycle to the defaultBillingCycle, if it is not set.
func (p *packet) InitializeCloudProvider(cluster *kubermaticv1.Cluster, update provider.ClusterUpdater) (*kubermaticv1.Cluster, error) {
	var err error
	if cluster.Spec.Cloud.Packet.BillingCycle == "" {
		cluster, err = update(cluster.Name, func(cluster *kubermaticv1.Cluster) {
			cluster.Spec.Cloud.Packet.BillingCycle = defaultBillingCycle
		})
		if err != nil {
			return nil, err
		}
	}

	return cluster, nil
}

// CleanUpCloudProvider
func (p *packet) CleanUpCloudProvider(cluster *kubermaticv1.Cluster, update provider.ClusterUpdater) (*kubermaticv1.Cluster, error) {
	return cluster, nil
}

// ValidateCloudSpecUpdate verifies whether an update of cloud spec is valid and permitted
func (p *packet) ValidateCloudSpecUpdate(oldSpec kubermaticv1.CloudSpec, newSpec kubermaticv1.CloudSpec) error {
	return nil
}
