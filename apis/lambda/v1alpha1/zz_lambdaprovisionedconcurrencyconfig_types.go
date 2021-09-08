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

type LambdaProvisionedConcurrencyConfigObservation struct {
}

type LambdaProvisionedConcurrencyConfigParameters struct {
	FunctionName string `json:"functionName" tf:"function_name"`

	ProvisionedConcurrentExecutions int64 `json:"provisionedConcurrentExecutions" tf:"provisioned_concurrent_executions"`

	Qualifier string `json:"qualifier" tf:"qualifier"`
}

// LambdaProvisionedConcurrencyConfigSpec defines the desired state of LambdaProvisionedConcurrencyConfig
type LambdaProvisionedConcurrencyConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LambdaProvisionedConcurrencyConfigParameters `json:"forProvider"`
}

// LambdaProvisionedConcurrencyConfigStatus defines the observed state of LambdaProvisionedConcurrencyConfig.
type LambdaProvisionedConcurrencyConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LambdaProvisionedConcurrencyConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaProvisionedConcurrencyConfig is the Schema for the LambdaProvisionedConcurrencyConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LambdaProvisionedConcurrencyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaProvisionedConcurrencyConfigSpec   `json:"spec"`
	Status            LambdaProvisionedConcurrencyConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaProvisionedConcurrencyConfigList contains a list of LambdaProvisionedConcurrencyConfigs
type LambdaProvisionedConcurrencyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaProvisionedConcurrencyConfig `json:"items"`
}

// Repository type metadata.
var (
	LambdaProvisionedConcurrencyConfigKind             = "LambdaProvisionedConcurrencyConfig"
	LambdaProvisionedConcurrencyConfigGroupKind        = schema.GroupKind{Group: Group, Kind: LambdaProvisionedConcurrencyConfigKind}.String()
	LambdaProvisionedConcurrencyConfigKindAPIVersion   = LambdaProvisionedConcurrencyConfigKind + "." + GroupVersion.String()
	LambdaProvisionedConcurrencyConfigGroupVersionKind = GroupVersion.WithKind(LambdaProvisionedConcurrencyConfigKind)
)

func init() {
	SchemeBuilder.Register(&LambdaProvisionedConcurrencyConfig{}, &LambdaProvisionedConcurrencyConfigList{})
}
