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

type CognitoConfigObservation struct {
}

type CognitoConfigParameters struct {
	ClientId string `json:"clientId" tf:"client_id"`

	UserPool string `json:"userPool" tf:"user_pool"`
}

type OidcConfigObservation struct {
}

type OidcConfigParameters struct {
	AuthorizationEndpoint string `json:"authorizationEndpoint" tf:"authorization_endpoint"`

	ClientId string `json:"clientId" tf:"client_id"`

	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	Issuer string `json:"issuer" tf:"issuer"`

	JwksUri string `json:"jwksUri" tf:"jwks_uri"`

	LogoutEndpoint string `json:"logoutEndpoint" tf:"logout_endpoint"`

	TokenEndpoint string `json:"tokenEndpoint" tf:"token_endpoint"`

	UserInfoEndpoint string `json:"userInfoEndpoint" tf:"user_info_endpoint"`
}

type SagemakerWorkforceObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Subdomain string `json:"subdomain" tf:"subdomain"`
}

type SagemakerWorkforceParameters struct {
	CognitoConfig []CognitoConfigParameters `json:"cognitoConfig,omitempty" tf:"cognito_config"`

	OidcConfig []OidcConfigParameters `json:"oidcConfig,omitempty" tf:"oidc_config"`

	SourceIpConfig []SourceIpConfigParameters `json:"sourceIpConfig,omitempty" tf:"source_ip_config"`

	WorkforceName string `json:"workforceName" tf:"workforce_name"`
}

type SourceIpConfigObservation struct {
}

type SourceIpConfigParameters struct {
	Cidrs []string `json:"cidrs" tf:"cidrs"`
}

// SagemakerWorkforceSpec defines the desired state of SagemakerWorkforce
type SagemakerWorkforceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerWorkforceParameters `json:"forProvider"`
}

// SagemakerWorkforceStatus defines the observed state of SagemakerWorkforce.
type SagemakerWorkforceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerWorkforceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerWorkforce is the Schema for the SagemakerWorkforces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerWorkforce struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerWorkforceSpec   `json:"spec"`
	Status            SagemakerWorkforceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerWorkforceList contains a list of SagemakerWorkforces
type SagemakerWorkforceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerWorkforce `json:"items"`
}

// Repository type metadata.
var (
	SagemakerWorkforceKind             = "SagemakerWorkforce"
	SagemakerWorkforceGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerWorkforceKind}.String()
	SagemakerWorkforceKindAPIVersion   = SagemakerWorkforceKind + "." + GroupVersion.String()
	SagemakerWorkforceGroupVersionKind = GroupVersion.WithKind(SagemakerWorkforceKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerWorkforce{}, &SagemakerWorkforceList{})
}
