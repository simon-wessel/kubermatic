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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SSHKeyResourceName represents "Resource" defined in Kubernetes.
	SSHKeyResourceName = "usersshkeies"

	// SSHKeyKind represents "Kind" defined in Kubernetes.
	SSHKeyKind = "UserSSHKey"
)

// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=".spec.name",name="HumanReadableName",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.owner",name="Owner",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// UserSSHKey specifies a users UserSSHKey.
type UserSSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SSHKeySpec `json:"spec,omitempty"`
}

type SSHKeySpec struct {
	Owner       string   `json:"owner"`
	Name        string   `json:"name"`
	Fingerprint string   `json:"fingerprint"`
	PublicKey   string   `json:"publicKey"`
	Clusters    []string `json:"clusters"`
}

func (sk *UserSSHKey) IsUsedByCluster(clustername string) bool {
	if sk.Spec.Clusters == nil {
		return false
	}
	for _, name := range sk.Spec.Clusters {
		if name == clustername {
			return true
		}
	}
	return false
}

func (sk *UserSSHKey) RemoveFromCluster(clustername string) {
	for i, cl := range sk.Spec.Clusters {
		if cl != clustername {
			continue
		}
		// Don't break we don't check for duplicates when adding clusters!
		sk.Spec.Clusters = append(sk.Spec.Clusters[:i], sk.Spec.Clusters[i+1:]...)
	}
}

func (sk *UserSSHKey) AddToCluster(clustername string) {
	sk.Spec.Clusters = append(sk.Spec.Clusters, clustername)
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// UserSSHKeyList specifies a users UserSSHKey.
type UserSSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UserSSHKey `json:"items"`
}
