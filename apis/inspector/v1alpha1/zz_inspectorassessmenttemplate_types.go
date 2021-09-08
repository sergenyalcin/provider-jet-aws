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

type InspectorAssessmentTemplateObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type InspectorAssessmentTemplateParameters struct {
	Duration int64 `json:"duration" tf:"duration"`

	Name string `json:"name" tf:"name"`

	RulesPackageArns []string `json:"rulesPackageArns" tf:"rules_package_arns"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TargetArn string `json:"targetArn" tf:"target_arn"`
}

// InspectorAssessmentTemplateSpec defines the desired state of InspectorAssessmentTemplate
type InspectorAssessmentTemplateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InspectorAssessmentTemplateParameters `json:"forProvider"`
}

// InspectorAssessmentTemplateStatus defines the observed state of InspectorAssessmentTemplate.
type InspectorAssessmentTemplateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InspectorAssessmentTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorAssessmentTemplate is the Schema for the InspectorAssessmentTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type InspectorAssessmentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InspectorAssessmentTemplateSpec   `json:"spec"`
	Status            InspectorAssessmentTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorAssessmentTemplateList contains a list of InspectorAssessmentTemplates
type InspectorAssessmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InspectorAssessmentTemplate `json:"items"`
}

// Repository type metadata.
var (
	InspectorAssessmentTemplateKind             = "InspectorAssessmentTemplate"
	InspectorAssessmentTemplateGroupKind        = schema.GroupKind{Group: Group, Kind: InspectorAssessmentTemplateKind}.String()
	InspectorAssessmentTemplateKindAPIVersion   = InspectorAssessmentTemplateKind + "." + GroupVersion.String()
	InspectorAssessmentTemplateGroupVersionKind = GroupVersion.WithKind(InspectorAssessmentTemplateKind)
)

func init() {
	SchemeBuilder.Register(&InspectorAssessmentTemplate{}, &InspectorAssessmentTemplateList{})
}
