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

type AnalyticsConfigurationObservation struct {
}

type AnalyticsConfigurationParameters struct {
	ApplicationArn *string `json:"applicationArn,omitempty" tf:"application_arn"`

	ApplicationId *string `json:"applicationId,omitempty" tf:"application_id"`

	ExternalId *string `json:"externalId,omitempty" tf:"external_id"`

	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`

	UserDataShared *bool `json:"userDataShared,omitempty" tf:"user_data_shared"`
}

type CognitoUserPoolClientObservation struct {
	ClientSecret string `json:"clientSecret" tf:"client_secret"`
}

type CognitoUserPoolClientParameters struct {
	AccessTokenValidity *int64 `json:"accessTokenValidity,omitempty" tf:"access_token_validity"`

	AllowedOauthFlows []string `json:"allowedOauthFlows,omitempty" tf:"allowed_oauth_flows"`

	AllowedOauthFlowsUserPoolClient *bool `json:"allowedOauthFlowsUserPoolClient,omitempty" tf:"allowed_oauth_flows_user_pool_client"`

	AllowedOauthScopes []string `json:"allowedOauthScopes,omitempty" tf:"allowed_oauth_scopes"`

	AnalyticsConfiguration []AnalyticsConfigurationParameters `json:"analyticsConfiguration,omitempty" tf:"analytics_configuration"`

	CallbackUrls []string `json:"callbackUrls,omitempty" tf:"callback_urls"`

	DefaultRedirectUri *string `json:"defaultRedirectUri,omitempty" tf:"default_redirect_uri"`

	EnableTokenRevocation *bool `json:"enableTokenRevocation,omitempty" tf:"enable_token_revocation"`

	ExplicitAuthFlows []string `json:"explicitAuthFlows,omitempty" tf:"explicit_auth_flows"`

	GenerateSecret *bool `json:"generateSecret,omitempty" tf:"generate_secret"`

	IdTokenValidity *int64 `json:"idTokenValidity,omitempty" tf:"id_token_validity"`

	LogoutUrls []string `json:"logoutUrls,omitempty" tf:"logout_urls"`

	Name string `json:"name" tf:"name"`

	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors,omitempty" tf:"prevent_user_existence_errors"`

	ReadAttributes []string `json:"readAttributes,omitempty" tf:"read_attributes"`

	RefreshTokenValidity *int64 `json:"refreshTokenValidity,omitempty" tf:"refresh_token_validity"`

	SupportedIdentityProviders []string `json:"supportedIdentityProviders,omitempty" tf:"supported_identity_providers"`

	TokenValidityUnits []TokenValidityUnitsParameters `json:"tokenValidityUnits,omitempty" tf:"token_validity_units"`

	UserPoolId string `json:"userPoolId" tf:"user_pool_id"`

	WriteAttributes []string `json:"writeAttributes,omitempty" tf:"write_attributes"`
}

type TokenValidityUnitsObservation struct {
}

type TokenValidityUnitsParameters struct {
	AccessToken *string `json:"accessToken,omitempty" tf:"access_token"`

	IdToken *string `json:"idToken,omitempty" tf:"id_token"`

	RefreshToken *string `json:"refreshToken,omitempty" tf:"refresh_token"`
}

// CognitoUserPoolClientSpec defines the desired state of CognitoUserPoolClient
type CognitoUserPoolClientSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoUserPoolClientParameters `json:"forProvider"`
}

// CognitoUserPoolClientStatus defines the observed state of CognitoUserPoolClient.
type CognitoUserPoolClientStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoUserPoolClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolClient is the Schema for the CognitoUserPoolClients API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CognitoUserPoolClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolClientSpec   `json:"spec"`
	Status            CognitoUserPoolClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolClientList contains a list of CognitoUserPoolClients
type CognitoUserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPoolClient `json:"items"`
}

// Repository type metadata.
var (
	CognitoUserPoolClientKind             = "CognitoUserPoolClient"
	CognitoUserPoolClientGroupKind        = schema.GroupKind{Group: Group, Kind: CognitoUserPoolClientKind}.String()
	CognitoUserPoolClientKindAPIVersion   = CognitoUserPoolClientKind + "." + GroupVersion.String()
	CognitoUserPoolClientGroupVersionKind = GroupVersion.WithKind(CognitoUserPoolClientKind)
)

func init() {
	SchemeBuilder.Register(&CognitoUserPoolClient{}, &CognitoUserPoolClientList{})
}
