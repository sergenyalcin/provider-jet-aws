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

type ConfigConfigurationRecorderObservation struct {
}

type ConfigConfigurationRecorderParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	RecordingGroup []RecordingGroupParameters `json:"recordingGroup,omitempty" tf:"recording_group"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type RecordingGroupObservation struct {
}

type RecordingGroupParameters struct {
	AllSupported *bool `json:"allSupported,omitempty" tf:"all_supported"`

	IncludeGlobalResourceTypes *bool `json:"includeGlobalResourceTypes,omitempty" tf:"include_global_resource_types"`

	ResourceTypes []string `json:"resourceTypes,omitempty" tf:"resource_types"`
}

// ConfigConfigurationRecorderSpec defines the desired state of ConfigConfigurationRecorder
type ConfigConfigurationRecorderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConfigConfigurationRecorderParameters `json:"forProvider"`
}

// ConfigConfigurationRecorderStatus defines the observed state of ConfigConfigurationRecorder.
type ConfigConfigurationRecorderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConfigConfigurationRecorderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigConfigurationRecorder is the Schema for the ConfigConfigurationRecorders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ConfigConfigurationRecorder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationRecorderSpec   `json:"spec"`
	Status            ConfigConfigurationRecorderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigConfigurationRecorderList contains a list of ConfigConfigurationRecorders
type ConfigConfigurationRecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigConfigurationRecorder `json:"items"`
}

// Repository type metadata.
var (
	ConfigConfigurationRecorderKind             = "ConfigConfigurationRecorder"
	ConfigConfigurationRecorderGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigConfigurationRecorderKind}.String()
	ConfigConfigurationRecorderKindAPIVersion   = ConfigConfigurationRecorderKind + "." + GroupVersion.String()
	ConfigConfigurationRecorderGroupVersionKind = GroupVersion.WithKind(ConfigConfigurationRecorderKind)
)

func init() {
	SchemeBuilder.Register(&ConfigConfigurationRecorder{}, &ConfigConfigurationRecorderList{})
}
