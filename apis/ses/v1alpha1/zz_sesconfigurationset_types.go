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

type DeliveryOptionsObservation struct {
}

type DeliveryOptionsParameters struct {
	TlsPolicy *string `json:"tlsPolicy,omitempty" tf:"tls_policy"`
}

type SesConfigurationSetObservation struct {
	Arn string `json:"arn" tf:"arn"`

	LastFreshStart string `json:"lastFreshStart" tf:"last_fresh_start"`
}

type SesConfigurationSetParameters struct {
	DeliveryOptions []DeliveryOptionsParameters `json:"deliveryOptions,omitempty" tf:"delivery_options"`

	Name string `json:"name" tf:"name"`

	ReputationMetricsEnabled *bool `json:"reputationMetricsEnabled,omitempty" tf:"reputation_metrics_enabled"`

	SendingEnabled *bool `json:"sendingEnabled,omitempty" tf:"sending_enabled"`
}

// SesConfigurationSetSpec defines the desired state of SesConfigurationSet
type SesConfigurationSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SesConfigurationSetParameters `json:"forProvider"`
}

// SesConfigurationSetStatus defines the observed state of SesConfigurationSet.
type SesConfigurationSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SesConfigurationSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SesConfigurationSet is the Schema for the SesConfigurationSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SesConfigurationSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesConfigurationSetSpec   `json:"spec"`
	Status            SesConfigurationSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesConfigurationSetList contains a list of SesConfigurationSets
type SesConfigurationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesConfigurationSet `json:"items"`
}

// Repository type metadata.
var (
	SesConfigurationSetKind             = "SesConfigurationSet"
	SesConfigurationSetGroupKind        = schema.GroupKind{Group: Group, Kind: SesConfigurationSetKind}.String()
	SesConfigurationSetKindAPIVersion   = SesConfigurationSetKind + "." + GroupVersion.String()
	SesConfigurationSetGroupVersionKind = GroupVersion.WithKind(SesConfigurationSetKind)
)

func init() {
	SchemeBuilder.Register(&SesConfigurationSet{}, &SesConfigurationSetList{})
}
