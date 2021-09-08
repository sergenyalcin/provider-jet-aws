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

type ApiGatewayGatewayResponseObservation struct {
}

type ApiGatewayGatewayResponseParameters struct {
	ResponseParameters map[string]string `json:"responseParameters,omitempty" tf:"response_parameters"`

	ResponseTemplates map[string]string `json:"responseTemplates,omitempty" tf:"response_templates"`

	ResponseType string `json:"responseType" tf:"response_type"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	StatusCode *string `json:"statusCode,omitempty" tf:"status_code"`
}

// ApiGatewayGatewayResponseSpec defines the desired state of ApiGatewayGatewayResponse
type ApiGatewayGatewayResponseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayGatewayResponseParameters `json:"forProvider"`
}

// ApiGatewayGatewayResponseStatus defines the observed state of ApiGatewayGatewayResponse.
type ApiGatewayGatewayResponseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayGatewayResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayGatewayResponse is the Schema for the ApiGatewayGatewayResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayGatewayResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayGatewayResponseSpec   `json:"spec"`
	Status            ApiGatewayGatewayResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayGatewayResponseList contains a list of ApiGatewayGatewayResponses
type ApiGatewayGatewayResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayGatewayResponse `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayGatewayResponseKind             = "ApiGatewayGatewayResponse"
	ApiGatewayGatewayResponseGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayGatewayResponseKind}.String()
	ApiGatewayGatewayResponseKindAPIVersion   = ApiGatewayGatewayResponseKind + "." + GroupVersion.String()
	ApiGatewayGatewayResponseGroupVersionKind = GroupVersion.WithKind(ApiGatewayGatewayResponseKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayGatewayResponse{}, &ApiGatewayGatewayResponseList{})
}