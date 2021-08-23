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

// +kubebuilder:object:generate=true
// +groupName=waf.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/waf/v1alpha1"
)

type WafRegexPatternSetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type WafRegexPatternSetParameters struct {
	Name string `json:"name" tf:"name"`

	RegexPatternStrings []string `json:"regexPatternStrings,omitempty" tf:"regex_pattern_strings"`
}

// WafRegexPatternSetSpec defines the desired state of WafRegexPatternSet
type WafRegexPatternSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafRegexPatternSetParameters `json:"forProvider"`
}

// WafRegexPatternSetStatus defines the observed state of WafRegexPatternSet.
type WafRegexPatternSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafRegexPatternSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafRegexPatternSet is the Schema for the WafRegexPatternSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WafRegexPatternSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRegexPatternSetSpec   `json:"spec"`
	Status            WafRegexPatternSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafRegexPatternSetList contains a list of WafRegexPatternSets
type WafRegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafRegexPatternSet `json:"items"`
}

// Repository type metadata.
var (
	WafRegexPatternSetKind             = "WafRegexPatternSet"
	WafRegexPatternSetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: WafRegexPatternSetKind}.String()
	WafRegexPatternSetKindAPIVersion   = WafRegexPatternSetKind + "." + v1alpha1.GroupVersion.String()
	WafRegexPatternSetGroupVersionKind = v1alpha1.GroupVersion.WithKind(WafRegexPatternSetKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&WafRegexPatternSet{}, &WafRegexPatternSetList{})
}