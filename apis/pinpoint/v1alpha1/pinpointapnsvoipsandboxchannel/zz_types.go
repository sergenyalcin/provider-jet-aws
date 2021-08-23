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
// +groupName=pinpoint.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/pinpoint/v1alpha1"
)

type PinpointApnsVoipSandboxChannelObservation struct {
}

type PinpointApnsVoipSandboxChannelParameters struct {
	ApplicationId string `json:"applicationId" tf:"application_id"`

	BundleId *string `json:"bundleId,omitempty" tf:"bundle_id"`

	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key"`

	TeamId *string `json:"teamId,omitempty" tf:"team_id"`

	TokenKey *string `json:"tokenKey,omitempty" tf:"token_key"`

	TokenKeyId *string `json:"tokenKeyId,omitempty" tf:"token_key_id"`
}

// PinpointApnsVoipSandboxChannelSpec defines the desired state of PinpointApnsVoipSandboxChannel
type PinpointApnsVoipSandboxChannelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PinpointApnsVoipSandboxChannelParameters `json:"forProvider"`
}

// PinpointApnsVoipSandboxChannelStatus defines the observed state of PinpointApnsVoipSandboxChannel.
type PinpointApnsVoipSandboxChannelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PinpointApnsVoipSandboxChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsVoipSandboxChannel is the Schema for the PinpointApnsVoipSandboxChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PinpointApnsVoipSandboxChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointApnsVoipSandboxChannelSpec   `json:"spec"`
	Status            PinpointApnsVoipSandboxChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsVoipSandboxChannelList contains a list of PinpointApnsVoipSandboxChannels
type PinpointApnsVoipSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApnsVoipSandboxChannel `json:"items"`
}

// Repository type metadata.
var (
	PinpointApnsVoipSandboxChannelKind             = "PinpointApnsVoipSandboxChannel"
	PinpointApnsVoipSandboxChannelGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: PinpointApnsVoipSandboxChannelKind}.String()
	PinpointApnsVoipSandboxChannelKindAPIVersion   = PinpointApnsVoipSandboxChannelKind + "." + v1alpha1.GroupVersion.String()
	PinpointApnsVoipSandboxChannelGroupVersionKind = v1alpha1.GroupVersion.WithKind(PinpointApnsVoipSandboxChannelKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&PinpointApnsVoipSandboxChannel{}, &PinpointApnsVoipSandboxChannelList{})
}