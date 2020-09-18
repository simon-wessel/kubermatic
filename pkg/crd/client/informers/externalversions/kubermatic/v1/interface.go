// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "k8c.io/kubermatic/v2/pkg/crd/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Addons returns a AddonInformer.
	Addons() AddonInformer
	// AddonConfigs returns a AddonConfigInformer.
	AddonConfigs() AddonConfigInformer
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ConstraintTemplates returns a ConstraintTemplateInformer.
	ConstraintTemplates() ConstraintTemplateInformer
	// ExternalClusters returns a ExternalClusterInformer.
	ExternalClusters() ExternalClusterInformer
	// KubermaticSettings returns a KubermaticSettingInformer.
	KubermaticSettings() KubermaticSettingInformer
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
	// Users returns a UserInformer.
	Users() UserInformer
	// UserProjectBindings returns a UserProjectBindingInformer.
	UserProjectBindings() UserProjectBindingInformer
	// UserSSHKeys returns a UserSSHKeyInformer.
	UserSSHKeys() UserSSHKeyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Addons returns a AddonInformer.
func (v *version) Addons() AddonInformer {
	return &addonInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AddonConfigs returns a AddonConfigInformer.
func (v *version) AddonConfigs() AddonConfigInformer {
	return &addonConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ConstraintTemplates returns a ConstraintTemplateInformer.
func (v *version) ConstraintTemplates() ConstraintTemplateInformer {
	return &constraintTemplateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ExternalClusters returns a ExternalClusterInformer.
func (v *version) ExternalClusters() ExternalClusterInformer {
	return &externalClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KubermaticSettings returns a KubermaticSettingInformer.
func (v *version) KubermaticSettings() KubermaticSettingInformer {
	return &kubermaticSettingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// UserProjectBindings returns a UserProjectBindingInformer.
func (v *version) UserProjectBindings() UserProjectBindingInformer {
	return &userProjectBindingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// UserSSHKeys returns a UserSSHKeyInformer.
func (v *version) UserSSHKeys() UserSSHKeyInformer {
	return &userSSHKeyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
