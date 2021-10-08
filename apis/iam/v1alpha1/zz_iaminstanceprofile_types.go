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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IamInstanceProfileObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn"`

	CreateDate *string `json:"createDate,omitempty" tf:"create_date"`

	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id"`
}

type IamInstanceProfileParameters struct {

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=IamRole
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role"`

	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// IamInstanceProfileSpec defines the desired state of IamInstanceProfile
type IamInstanceProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IamInstanceProfileParameters `json:"forProvider"`
}

// IamInstanceProfileStatus defines the observed state of IamInstanceProfile.
type IamInstanceProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IamInstanceProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamInstanceProfile is the Schema for the IamInstanceProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamInstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamInstanceProfileSpec   `json:"spec"`
	Status            IamInstanceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamInstanceProfileList contains a list of IamInstanceProfiles
type IamInstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamInstanceProfile `json:"items"`
}

// Repository type metadata.
var (
	IamInstanceProfileKind             = "IamInstanceProfile"
	IamInstanceProfileGroupKind        = schema.GroupKind{Group: Group, Kind: IamInstanceProfileKind}.String()
	IamInstanceProfileKindAPIVersion   = IamInstanceProfileKind + "." + GroupVersion.String()
	IamInstanceProfileGroupVersionKind = GroupVersion.WithKind(IamInstanceProfileKind)
)

func init() {
	SchemeBuilder.Register(&IamInstanceProfile{}, &IamInstanceProfileList{})
}
