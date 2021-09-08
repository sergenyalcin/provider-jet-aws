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

type AmplifyWebhookObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Url string `json:"url" tf:"url"`
}

type AmplifyWebhookParameters struct {
	AppId string `json:"appId" tf:"app_id"`

	BranchName string `json:"branchName" tf:"branch_name"`

	Description *string `json:"description,omitempty" tf:"description"`
}

// AmplifyWebhookSpec defines the desired state of AmplifyWebhook
type AmplifyWebhookSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AmplifyWebhookParameters `json:"forProvider"`
}

// AmplifyWebhookStatus defines the observed state of AmplifyWebhook.
type AmplifyWebhookStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AmplifyWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AmplifyWebhook is the Schema for the AmplifyWebhooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AmplifyWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmplifyWebhookSpec   `json:"spec"`
	Status            AmplifyWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AmplifyWebhookList contains a list of AmplifyWebhooks
type AmplifyWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AmplifyWebhook `json:"items"`
}

// Repository type metadata.
var (
	AmplifyWebhookKind             = "AmplifyWebhook"
	AmplifyWebhookGroupKind        = schema.GroupKind{Group: Group, Kind: AmplifyWebhookKind}.String()
	AmplifyWebhookKindAPIVersion   = AmplifyWebhookKind + "." + GroupVersion.String()
	AmplifyWebhookGroupVersionKind = GroupVersion.WithKind(AmplifyWebhookKind)
)

func init() {
	SchemeBuilder.Register(&AmplifyWebhook{}, &AmplifyWebhookList{})
}