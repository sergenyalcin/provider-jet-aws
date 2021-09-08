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

type SecretsmanagerSecretPolicyObservation struct {
}

type SecretsmanagerSecretPolicyParameters struct {
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy"`

	Policy string `json:"policy" tf:"policy"`

	SecretArn string `json:"secretArn" tf:"secret_arn"`
}

// SecretsmanagerSecretPolicySpec defines the desired state of SecretsmanagerSecretPolicy
type SecretsmanagerSecretPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecretsmanagerSecretPolicyParameters `json:"forProvider"`
}

// SecretsmanagerSecretPolicyStatus defines the observed state of SecretsmanagerSecretPolicy.
type SecretsmanagerSecretPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecretsmanagerSecretPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecretPolicy is the Schema for the SecretsmanagerSecretPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SecretsmanagerSecretPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretsmanagerSecretPolicySpec   `json:"spec"`
	Status            SecretsmanagerSecretPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecretPolicyList contains a list of SecretsmanagerSecretPolicys
type SecretsmanagerSecretPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretsmanagerSecretPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecretsmanagerSecretPolicyKind             = "SecretsmanagerSecretPolicy"
	SecretsmanagerSecretPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: SecretsmanagerSecretPolicyKind}.String()
	SecretsmanagerSecretPolicyKindAPIVersion   = SecretsmanagerSecretPolicyKind + "." + GroupVersion.String()
	SecretsmanagerSecretPolicyGroupVersionKind = GroupVersion.WithKind(SecretsmanagerSecretPolicyKind)
)

func init() {
	SchemeBuilder.Register(&SecretsmanagerSecretPolicy{}, &SecretsmanagerSecretPolicyList{})
}
