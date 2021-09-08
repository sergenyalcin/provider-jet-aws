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

type GlueSchemaObservation struct {
	Arn string `json:"arn" tf:"arn"`

	LatestSchemaVersion int64 `json:"latestSchemaVersion" tf:"latest_schema_version"`

	NextSchemaVersion int64 `json:"nextSchemaVersion" tf:"next_schema_version"`

	RegistryName string `json:"registryName" tf:"registry_name"`

	SchemaCheckpoint int64 `json:"schemaCheckpoint" tf:"schema_checkpoint"`
}

type GlueSchemaParameters struct {
	Compatibility string `json:"compatibility" tf:"compatibility"`

	DataFormat string `json:"dataFormat" tf:"data_format"`

	Description *string `json:"description,omitempty" tf:"description"`

	RegistryArn *string `json:"registryArn,omitempty" tf:"registry_arn"`

	SchemaDefinition string `json:"schemaDefinition" tf:"schema_definition"`

	SchemaName string `json:"schemaName" tf:"schema_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// GlueSchemaSpec defines the desired state of GlueSchema
type GlueSchemaSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueSchemaParameters `json:"forProvider"`
}

// GlueSchemaStatus defines the observed state of GlueSchema.
type GlueSchemaStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueSchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueSchema is the Schema for the GlueSchemas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueSchemaSpec   `json:"spec"`
	Status            GlueSchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueSchemaList contains a list of GlueSchemas
type GlueSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueSchema `json:"items"`
}

// Repository type metadata.
var (
	GlueSchemaKind             = "GlueSchema"
	GlueSchemaGroupKind        = schema.GroupKind{Group: Group, Kind: GlueSchemaKind}.String()
	GlueSchemaKindAPIVersion   = GlueSchemaKind + "." + GroupVersion.String()
	GlueSchemaGroupVersionKind = GroupVersion.WithKind(GlueSchemaKind)
)

func init() {
	SchemeBuilder.Register(&GlueSchema{}, &GlueSchemaList{})
}
