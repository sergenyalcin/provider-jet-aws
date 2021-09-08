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

type ApprovalRuleObservation struct {
}

type ApprovalRuleParameters struct {
	ApproveAfterDays *int64 `json:"approveAfterDays,omitempty" tf:"approve_after_days"`

	ApproveUntilDate *string `json:"approveUntilDate,omitempty" tf:"approve_until_date"`

	ComplianceLevel *string `json:"complianceLevel,omitempty" tf:"compliance_level"`

	EnableNonSecurity *bool `json:"enableNonSecurity,omitempty" tf:"enable_non_security"`

	PatchFilter []PatchFilterParameters `json:"patchFilter" tf:"patch_filter"`
}

type GlobalFilterObservation struct {
}

type GlobalFilterParameters struct {
	Key string `json:"key" tf:"key"`

	Values []string `json:"values" tf:"values"`
}

type PatchFilterObservation struct {
}

type PatchFilterParameters struct {
	Key string `json:"key" tf:"key"`

	Values []string `json:"values" tf:"values"`
}

type SourceObservation struct {
}

type SourceParameters struct {
	Configuration string `json:"configuration" tf:"configuration"`

	Name string `json:"name" tf:"name"`

	Products []string `json:"products" tf:"products"`
}

type SsmPatchBaselineObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SsmPatchBaselineParameters struct {
	ApprovalRule []ApprovalRuleParameters `json:"approvalRule,omitempty" tf:"approval_rule"`

	ApprovedPatches []string `json:"approvedPatches,omitempty" tf:"approved_patches"`

	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel,omitempty" tf:"approved_patches_compliance_level"`

	ApprovedPatchesEnableNonSecurity *bool `json:"approvedPatchesEnableNonSecurity,omitempty" tf:"approved_patches_enable_non_security"`

	Description *string `json:"description,omitempty" tf:"description"`

	GlobalFilter []GlobalFilterParameters `json:"globalFilter,omitempty" tf:"global_filter"`

	Name string `json:"name" tf:"name"`

	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system"`

	RejectedPatches []string `json:"rejectedPatches,omitempty" tf:"rejected_patches"`

	RejectedPatchesAction *string `json:"rejectedPatchesAction,omitempty" tf:"rejected_patches_action"`

	Source []SourceParameters `json:"source,omitempty" tf:"source"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SsmPatchBaselineSpec defines the desired state of SsmPatchBaseline
type SsmPatchBaselineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsmPatchBaselineParameters `json:"forProvider"`
}

// SsmPatchBaselineStatus defines the observed state of SsmPatchBaseline.
type SsmPatchBaselineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsmPatchBaselineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsmPatchBaseline is the Schema for the SsmPatchBaselines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SsmPatchBaseline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmPatchBaselineSpec   `json:"spec"`
	Status            SsmPatchBaselineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmPatchBaselineList contains a list of SsmPatchBaselines
type SsmPatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmPatchBaseline `json:"items"`
}

// Repository type metadata.
var (
	SsmPatchBaselineKind             = "SsmPatchBaseline"
	SsmPatchBaselineGroupKind        = schema.GroupKind{Group: Group, Kind: SsmPatchBaselineKind}.String()
	SsmPatchBaselineKindAPIVersion   = SsmPatchBaselineKind + "." + GroupVersion.String()
	SsmPatchBaselineGroupVersionKind = GroupVersion.WithKind(SsmPatchBaselineKind)
)

func init() {
	SchemeBuilder.Register(&SsmPatchBaseline{}, &SsmPatchBaselineList{})
}
