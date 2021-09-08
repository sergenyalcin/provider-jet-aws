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

type IamServiceLinkedRoleObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreateDate string `json:"createDate" tf:"create_date"`

	Name string `json:"name" tf:"name"`

	Path string `json:"path" tf:"path"`

	UniqueId string `json:"uniqueId" tf:"unique_id"`
}

type IamServiceLinkedRoleParameters struct {
	AwsServiceName string `json:"awsServiceName" tf:"aws_service_name"`

	CustomSuffix *string `json:"customSuffix,omitempty" tf:"custom_suffix"`

	Description *string `json:"description,omitempty" tf:"description"`
}

// IamServiceLinkedRoleSpec defines the desired state of IamServiceLinkedRole
type IamServiceLinkedRoleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamServiceLinkedRoleParameters `json:"forProvider"`
}

// IamServiceLinkedRoleStatus defines the observed state of IamServiceLinkedRole.
type IamServiceLinkedRoleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamServiceLinkedRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamServiceLinkedRole is the Schema for the IamServiceLinkedRoles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamServiceLinkedRoleSpec   `json:"spec"`
	Status            IamServiceLinkedRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamServiceLinkedRoleList contains a list of IamServiceLinkedRoles
type IamServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamServiceLinkedRole `json:"items"`
}

// Repository type metadata.
var (
	IamServiceLinkedRoleKind             = "IamServiceLinkedRole"
	IamServiceLinkedRoleGroupKind        = schema.GroupKind{Group: Group, Kind: IamServiceLinkedRoleKind}.String()
	IamServiceLinkedRoleKindAPIVersion   = IamServiceLinkedRoleKind + "." + GroupVersion.String()
	IamServiceLinkedRoleGroupVersionKind = GroupVersion.WithKind(IamServiceLinkedRoleKind)
)

func init() {
	SchemeBuilder.Register(&IamServiceLinkedRole{}, &IamServiceLinkedRoleList{})
}
