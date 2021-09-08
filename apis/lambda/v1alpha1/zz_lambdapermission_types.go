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

type LambdaPermissionObservation struct {
}

type LambdaPermissionParameters struct {
	Action string `json:"action" tf:"action"`

	EventSourceToken *string `json:"eventSourceToken,omitempty" tf:"event_source_token"`

	FunctionName string `json:"functionName" tf:"function_name"`

	Principal string `json:"principal" tf:"principal"`

	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier"`

	SourceAccount *string `json:"sourceAccount,omitempty" tf:"source_account"`

	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn"`

	StatementId *string `json:"statementId,omitempty" tf:"statement_id"`

	StatementIdPrefix *string `json:"statementIdPrefix,omitempty" tf:"statement_id_prefix"`
}

// LambdaPermissionSpec defines the desired state of LambdaPermission
type LambdaPermissionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LambdaPermissionParameters `json:"forProvider"`
}

// LambdaPermissionStatus defines the observed state of LambdaPermission.
type LambdaPermissionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LambdaPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaPermission is the Schema for the LambdaPermissions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LambdaPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaPermissionSpec   `json:"spec"`
	Status            LambdaPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaPermissionList contains a list of LambdaPermissions
type LambdaPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaPermission `json:"items"`
}

// Repository type metadata.
var (
	LambdaPermissionKind             = "LambdaPermission"
	LambdaPermissionGroupKind        = schema.GroupKind{Group: Group, Kind: LambdaPermissionKind}.String()
	LambdaPermissionKindAPIVersion   = LambdaPermissionKind + "." + GroupVersion.String()
	LambdaPermissionGroupVersionKind = GroupVersion.WithKind(LambdaPermissionKind)
)

func init() {
	SchemeBuilder.Register(&LambdaPermission{}, &LambdaPermissionList{})
}
