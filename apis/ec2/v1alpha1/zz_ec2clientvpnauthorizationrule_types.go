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

type Ec2ClientVpnAuthorizationRuleObservation struct {
}

type Ec2ClientVpnAuthorizationRuleParameters struct {
	AccessGroupId *string `json:"accessGroupId,omitempty" tf:"access_group_id"`

	AuthorizeAllGroups *bool `json:"authorizeAllGroups,omitempty" tf:"authorize_all_groups"`

	ClientVpnEndpointId string `json:"clientVpnEndpointId" tf:"client_vpn_endpoint_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	TargetNetworkCidr string `json:"targetNetworkCidr" tf:"target_network_cidr"`
}

// Ec2ClientVpnAuthorizationRuleSpec defines the desired state of Ec2ClientVpnAuthorizationRule
type Ec2ClientVpnAuthorizationRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2ClientVpnAuthorizationRuleParameters `json:"forProvider"`
}

// Ec2ClientVpnAuthorizationRuleStatus defines the observed state of Ec2ClientVpnAuthorizationRule.
type Ec2ClientVpnAuthorizationRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2ClientVpnAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnAuthorizationRule is the Schema for the Ec2ClientVpnAuthorizationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2ClientVpnAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ClientVpnAuthorizationRuleSpec   `json:"spec"`
	Status            Ec2ClientVpnAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnAuthorizationRuleList contains a list of Ec2ClientVpnAuthorizationRules
type Ec2ClientVpnAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ClientVpnAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	Ec2ClientVpnAuthorizationRuleKind             = "Ec2ClientVpnAuthorizationRule"
	Ec2ClientVpnAuthorizationRuleGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2ClientVpnAuthorizationRuleKind}.String()
	Ec2ClientVpnAuthorizationRuleKindAPIVersion   = Ec2ClientVpnAuthorizationRuleKind + "." + GroupVersion.String()
	Ec2ClientVpnAuthorizationRuleGroupVersionKind = GroupVersion.WithKind(Ec2ClientVpnAuthorizationRuleKind)
)

func init() {
	SchemeBuilder.Register(&Ec2ClientVpnAuthorizationRule{}, &Ec2ClientVpnAuthorizationRuleList{})
}
