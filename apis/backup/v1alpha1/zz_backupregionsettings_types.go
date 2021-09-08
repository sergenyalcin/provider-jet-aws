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

type BackupRegionSettingsObservation struct {
}

type BackupRegionSettingsParameters struct {
	ResourceTypeOptInPreference map[string]bool `json:"resourceTypeOptInPreference" tf:"resource_type_opt_in_preference"`
}

// BackupRegionSettingsSpec defines the desired state of BackupRegionSettings
type BackupRegionSettingsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BackupRegionSettingsParameters `json:"forProvider"`
}

// BackupRegionSettingsStatus defines the observed state of BackupRegionSettings.
type BackupRegionSettingsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BackupRegionSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupRegionSettings is the Schema for the BackupRegionSettingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type BackupRegionSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupRegionSettingsSpec   `json:"spec"`
	Status            BackupRegionSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupRegionSettingsList contains a list of BackupRegionSettingss
type BackupRegionSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupRegionSettings `json:"items"`
}

// Repository type metadata.
var (
	BackupRegionSettingsKind             = "BackupRegionSettings"
	BackupRegionSettingsGroupKind        = schema.GroupKind{Group: Group, Kind: BackupRegionSettingsKind}.String()
	BackupRegionSettingsKindAPIVersion   = BackupRegionSettingsKind + "." + GroupVersion.String()
	BackupRegionSettingsGroupVersionKind = GroupVersion.WithKind(BackupRegionSettingsKind)
)

func init() {
	SchemeBuilder.Register(&BackupRegionSettings{}, &BackupRegionSettingsList{})
}
