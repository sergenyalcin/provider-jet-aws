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

type IamGroupPolicyAttachmentObservation struct {
}

type IamGroupPolicyAttachmentParameters struct {

	// +crossplane:generate:reference:type=IamGroup
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group"`

	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=IamPolicy
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-aws/config/iam.PolicyARNExtractor()
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn"`

	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// IamGroupPolicyAttachmentSpec defines the desired state of IamGroupPolicyAttachment
type IamGroupPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IamGroupPolicyAttachmentParameters `json:"forProvider"`
}

// IamGroupPolicyAttachmentStatus defines the observed state of IamGroupPolicyAttachment.
type IamGroupPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IamGroupPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamGroupPolicyAttachment is the Schema for the IamGroupPolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamGroupPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamGroupPolicyAttachmentSpec   `json:"spec"`
	Status            IamGroupPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamGroupPolicyAttachmentList contains a list of IamGroupPolicyAttachments
type IamGroupPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamGroupPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	IamGroupPolicyAttachmentKind             = "IamGroupPolicyAttachment"
	IamGroupPolicyAttachmentGroupKind        = schema.GroupKind{Group: Group, Kind: IamGroupPolicyAttachmentKind}.String()
	IamGroupPolicyAttachmentKindAPIVersion   = IamGroupPolicyAttachmentKind + "." + GroupVersion.String()
	IamGroupPolicyAttachmentGroupVersionKind = GroupVersion.WithKind(IamGroupPolicyAttachmentKind)
)

func init() {
	SchemeBuilder.Register(&IamGroupPolicyAttachment{}, &IamGroupPolicyAttachmentList{})
}
