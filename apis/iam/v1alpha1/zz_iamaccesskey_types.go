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

type IamAccessKeyObservation struct {
	CreateDate string `json:"createDate" tf:"create_date"`

	EncryptedSecret string `json:"encryptedSecret" tf:"encrypted_secret"`

	EncryptedSesSmtpPasswordV4 string `json:"encryptedSesSmtpPasswordV4" tf:"encrypted_ses_smtp_password_v4"`

	KeyFingerprint string `json:"keyFingerprint" tf:"key_fingerprint"`

	Secret string `json:"secret" tf:"secret"`

	SesSmtpPasswordV4 string `json:"sesSmtpPasswordV4" tf:"ses_smtp_password_v4"`
}

type IamAccessKeyParameters struct {
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key"`

	Status *string `json:"status,omitempty" tf:"status"`

	User string `json:"user" tf:"user"`
}

// IamAccessKeySpec defines the desired state of IamAccessKey
type IamAccessKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamAccessKeyParameters `json:"forProvider"`
}

// IamAccessKeyStatus defines the observed state of IamAccessKey.
type IamAccessKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamAccessKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamAccessKey is the Schema for the IamAccessKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamAccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamAccessKeySpec   `json:"spec"`
	Status            IamAccessKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamAccessKeyList contains a list of IamAccessKeys
type IamAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamAccessKey `json:"items"`
}

// Repository type metadata.
var (
	IamAccessKeyKind             = "IamAccessKey"
	IamAccessKeyGroupKind        = schema.GroupKind{Group: Group, Kind: IamAccessKeyKind}.String()
	IamAccessKeyKindAPIVersion   = IamAccessKeyKind + "." + GroupVersion.String()
	IamAccessKeyGroupVersionKind = GroupVersion.WithKind(IamAccessKeyKind)
)

func init() {
	SchemeBuilder.Register(&IamAccessKey{}, &IamAccessKeyList{})
}
