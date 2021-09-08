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

type VpnGatewayRoutePropagationObservation struct {
}

type VpnGatewayRoutePropagationParameters struct {
	RouteTableId string `json:"routeTableId" tf:"route_table_id"`

	VpnGatewayId string `json:"vpnGatewayId" tf:"vpn_gateway_id"`
}

// VpnGatewayRoutePropagationSpec defines the desired state of VpnGatewayRoutePropagation
type VpnGatewayRoutePropagationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpnGatewayRoutePropagationParameters `json:"forProvider"`
}

// VpnGatewayRoutePropagationStatus defines the observed state of VpnGatewayRoutePropagation.
type VpnGatewayRoutePropagationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpnGatewayRoutePropagationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpnGatewayRoutePropagation is the Schema for the VpnGatewayRoutePropagations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type VpnGatewayRoutePropagation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewayRoutePropagationSpec   `json:"spec"`
	Status            VpnGatewayRoutePropagationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnGatewayRoutePropagationList contains a list of VpnGatewayRoutePropagations
type VpnGatewayRoutePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnGatewayRoutePropagation `json:"items"`
}

// Repository type metadata.
var (
	VpnGatewayRoutePropagationKind             = "VpnGatewayRoutePropagation"
	VpnGatewayRoutePropagationGroupKind        = schema.GroupKind{Group: Group, Kind: VpnGatewayRoutePropagationKind}.String()
	VpnGatewayRoutePropagationKindAPIVersion   = VpnGatewayRoutePropagationKind + "." + GroupVersion.String()
	VpnGatewayRoutePropagationGroupVersionKind = GroupVersion.WithKind(VpnGatewayRoutePropagationKind)
)

func init() {
	SchemeBuilder.Register(&VpnGatewayRoutePropagation{}, &VpnGatewayRoutePropagationList{})
}
