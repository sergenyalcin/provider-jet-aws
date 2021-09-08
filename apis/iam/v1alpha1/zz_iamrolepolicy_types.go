/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type IamRolePolicyObservation struct {
}

type IamRolePolicyParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Policy string `json:"policy" tf:"policy"`

	Role string `json:"role" tf:"role"`
}

// IamRolePolicySpec defines the desired state of IamRolePolicy
type IamRolePolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamRolePolicyParameters `json:"forProvider"`
}

// IamRolePolicyStatus defines the observed state of IamRolePolicy.
type IamRolePolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamRolePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamRolePolicy is the Schema for the IamRolePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamRolePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamRolePolicySpec   `json:"spec"`
	Status            IamRolePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamRolePolicyList contains a list of IamRolePolicys
type IamRolePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamRolePolicy `json:"items"`
}

// Repository type metadata.
var (
	IamRolePolicyKind             = "IamRolePolicy"
	IamRolePolicyGroupKind        = schema.GroupKind{Group: Group, Kind: IamRolePolicyKind}.String()
	IamRolePolicyKindAPIVersion   = IamRolePolicyKind + "." + GroupVersion.String()
	IamRolePolicyGroupVersionKind = GroupVersion.WithKind(IamRolePolicyKind)
)

func init() {
	SchemeBuilder.Register(&IamRolePolicy{}, &IamRolePolicyList{})
}
