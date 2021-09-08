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

type DxBgpPeerObservation struct {
	AwsDevice string `json:"awsDevice" tf:"aws_device"`

	BgpPeerId string `json:"bgpPeerId" tf:"bgp_peer_id"`

	BgpStatus string `json:"bgpStatus" tf:"bgp_status"`
}

type DxBgpPeerParameters struct {
	AddressFamily string `json:"addressFamily" tf:"address_family"`

	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`

	BgpAsn int64 `json:"bgpAsn" tf:"bgp_asn"`

	BgpAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`

	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`

	VirtualInterfaceId string `json:"virtualInterfaceId" tf:"virtual_interface_id"`
}

// DxBgpPeerSpec defines the desired state of DxBgpPeer
type DxBgpPeerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxBgpPeerParameters `json:"forProvider"`
}

// DxBgpPeerStatus defines the observed state of DxBgpPeer.
type DxBgpPeerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxBgpPeerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxBgpPeer is the Schema for the DxBgpPeers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DxBgpPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxBgpPeerSpec   `json:"spec"`
	Status            DxBgpPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxBgpPeerList contains a list of DxBgpPeers
type DxBgpPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxBgpPeer `json:"items"`
}

// Repository type metadata.
var (
	DxBgpPeerKind             = "DxBgpPeer"
	DxBgpPeerGroupKind        = schema.GroupKind{Group: Group, Kind: DxBgpPeerKind}.String()
	DxBgpPeerKindAPIVersion   = DxBgpPeerKind + "." + GroupVersion.String()
	DxBgpPeerGroupVersionKind = GroupVersion.WithKind(DxBgpPeerKind)
)

func init() {
	SchemeBuilder.Register(&DxBgpPeer{}, &DxBgpPeerList{})
}