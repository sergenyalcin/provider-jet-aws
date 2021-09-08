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

type BackupPolicyObservation struct {
}

type BackupPolicyParameters struct {
	Status string `json:"status" tf:"status"`
}

type EfsBackupPolicyObservation struct {
}

type EfsBackupPolicyParameters struct {
	BackupPolicy []BackupPolicyParameters `json:"backupPolicy" tf:"backup_policy"`

	FileSystemId string `json:"fileSystemId" tf:"file_system_id"`
}

// EfsBackupPolicySpec defines the desired state of EfsBackupPolicy
type EfsBackupPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EfsBackupPolicyParameters `json:"forProvider"`
}

// EfsBackupPolicyStatus defines the observed state of EfsBackupPolicy.
type EfsBackupPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EfsBackupPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EfsBackupPolicy is the Schema for the EfsBackupPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EfsBackupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EfsBackupPolicySpec   `json:"spec"`
	Status            EfsBackupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsBackupPolicyList contains a list of EfsBackupPolicys
type EfsBackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsBackupPolicy `json:"items"`
}

// Repository type metadata.
var (
	EfsBackupPolicyKind             = "EfsBackupPolicy"
	EfsBackupPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: EfsBackupPolicyKind}.String()
	EfsBackupPolicyKindAPIVersion   = EfsBackupPolicyKind + "." + GroupVersion.String()
	EfsBackupPolicyGroupVersionKind = GroupVersion.WithKind(EfsBackupPolicyKind)
)

func init() {
	SchemeBuilder.Register(&EfsBackupPolicy{}, &EfsBackupPolicyList{})
}
