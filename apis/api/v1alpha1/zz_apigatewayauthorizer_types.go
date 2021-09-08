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

type ApiGatewayAuthorizerObservation struct {
}

type ApiGatewayAuthorizerParameters struct {
	AuthorizerCredentials *string `json:"authorizerCredentials,omitempty" tf:"authorizer_credentials"`

	AuthorizerResultTtlInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds"`

	AuthorizerUri *string `json:"authorizerUri,omitempty" tf:"authorizer_uri"`

	IdentitySource *string `json:"identitySource,omitempty" tf:"identity_source"`

	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression"`

	Name string `json:"name" tf:"name"`

	ProviderArns []string `json:"providerArns,omitempty" tf:"provider_arns"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	Type *string `json:"type,omitempty" tf:"type"`
}

// ApiGatewayAuthorizerSpec defines the desired state of ApiGatewayAuthorizer
type ApiGatewayAuthorizerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayAuthorizerParameters `json:"forProvider"`
}

// ApiGatewayAuthorizerStatus defines the observed state of ApiGatewayAuthorizer.
type ApiGatewayAuthorizerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayAuthorizerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayAuthorizer is the Schema for the ApiGatewayAuthorizers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayAuthorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayAuthorizerSpec   `json:"spec"`
	Status            ApiGatewayAuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayAuthorizerList contains a list of ApiGatewayAuthorizers
type ApiGatewayAuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayAuthorizer `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayAuthorizerKind             = "ApiGatewayAuthorizer"
	ApiGatewayAuthorizerGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayAuthorizerKind}.String()
	ApiGatewayAuthorizerKindAPIVersion   = ApiGatewayAuthorizerKind + "." + GroupVersion.String()
	ApiGatewayAuthorizerGroupVersionKind = GroupVersion.WithKind(ApiGatewayAuthorizerKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayAuthorizer{}, &ApiGatewayAuthorizerList{})
}
