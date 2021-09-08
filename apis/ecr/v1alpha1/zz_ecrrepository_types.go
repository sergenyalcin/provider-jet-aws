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

type EcrRepositoryObservation struct {
	Arn string `json:"arn" tf:"arn"`

	RegistryId string `json:"registryId" tf:"registry_id"`

	RepositoryUrl string `json:"repositoryUrl" tf:"repository_url"`
}

type EcrRepositoryParameters struct {
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration"`

	ImageScanningConfiguration []ImageScanningConfigurationParameters `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration"`

	ImageTagMutability *string `json:"imageTagMutability,omitempty" tf:"image_tag_mutability"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type"`

	KmsKey *string `json:"kmsKey,omitempty" tf:"kms_key"`
}

type ImageScanningConfigurationObservation struct {
}

type ImageScanningConfigurationParameters struct {
	ScanOnPush bool `json:"scanOnPush" tf:"scan_on_push"`
}

// EcrRepositorySpec defines the desired state of EcrRepository
type EcrRepositorySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EcrRepositoryParameters `json:"forProvider"`
}

// EcrRepositoryStatus defines the observed state of EcrRepository.
type EcrRepositoryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EcrRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EcrRepository is the Schema for the EcrRepositorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EcrRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcrRepositorySpec   `json:"spec"`
	Status            EcrRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcrRepositoryList contains a list of EcrRepositorys
type EcrRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcrRepository `json:"items"`
}

// Repository type metadata.
var (
	EcrRepositoryKind             = "EcrRepository"
	EcrRepositoryGroupKind        = schema.GroupKind{Group: Group, Kind: EcrRepositoryKind}.String()
	EcrRepositoryKindAPIVersion   = EcrRepositoryKind + "." + GroupVersion.String()
	EcrRepositoryGroupVersionKind = GroupVersion.WithKind(EcrRepositoryKind)
)

func init() {
	SchemeBuilder.Register(&EcrRepository{}, &EcrRepositoryList{})
}
