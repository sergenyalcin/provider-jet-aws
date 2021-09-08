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

type GeoMatchConstraintObservation struct {
}

type GeoMatchConstraintParameters struct {
	Type string `json:"type" tf:"type"`

	Value string `json:"value" tf:"value"`
}

type WafregionalGeoMatchSetObservation struct {
}

type WafregionalGeoMatchSetParameters struct {
	GeoMatchConstraint []GeoMatchConstraintParameters `json:"geoMatchConstraint,omitempty" tf:"geo_match_constraint"`

	Name string `json:"name" tf:"name"`
}

// WafregionalGeoMatchSetSpec defines the desired state of WafregionalGeoMatchSet
type WafregionalGeoMatchSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalGeoMatchSetParameters `json:"forProvider"`
}

// WafregionalGeoMatchSetStatus defines the observed state of WafregionalGeoMatchSet.
type WafregionalGeoMatchSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalGeoMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalGeoMatchSet is the Schema for the WafregionalGeoMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafregionalGeoMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalGeoMatchSetSpec   `json:"spec"`
	Status            WafregionalGeoMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalGeoMatchSetList contains a list of WafregionalGeoMatchSets
type WafregionalGeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalGeoMatchSet `json:"items"`
}

// Repository type metadata.
var (
	WafregionalGeoMatchSetKind             = "WafregionalGeoMatchSet"
	WafregionalGeoMatchSetGroupKind        = schema.GroupKind{Group: Group, Kind: WafregionalGeoMatchSetKind}.String()
	WafregionalGeoMatchSetKindAPIVersion   = WafregionalGeoMatchSetKind + "." + GroupVersion.String()
	WafregionalGeoMatchSetGroupVersionKind = GroupVersion.WithKind(WafregionalGeoMatchSetKind)
)

func init() {
	SchemeBuilder.Register(&WafregionalGeoMatchSet{}, &WafregionalGeoMatchSetList{})
}
