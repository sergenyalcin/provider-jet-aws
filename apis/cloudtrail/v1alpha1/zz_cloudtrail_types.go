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

type CloudtrailObservation struct {
	Arn string `json:"arn" tf:"arn"`

	HomeRegion string `json:"homeRegion" tf:"home_region"`
}

type CloudtrailParameters struct {
	CloudWatchLogsGroupArn *string `json:"cloudWatchLogsGroupArn,omitempty" tf:"cloud_watch_logs_group_arn"`

	CloudWatchLogsRoleArn *string `json:"cloudWatchLogsRoleArn,omitempty" tf:"cloud_watch_logs_role_arn"`

	EnableLogFileValidation *bool `json:"enableLogFileValidation,omitempty" tf:"enable_log_file_validation"`

	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging"`

	EventSelector []EventSelectorParameters `json:"eventSelector,omitempty" tf:"event_selector"`

	IncludeGlobalServiceEvents *bool `json:"includeGlobalServiceEvents,omitempty" tf:"include_global_service_events"`

	InsightSelector []InsightSelectorParameters `json:"insightSelector,omitempty" tf:"insight_selector"`

	IsMultiRegionTrail *bool `json:"isMultiRegionTrail,omitempty" tf:"is_multi_region_trail"`

	IsOrganizationTrail *bool `json:"isOrganizationTrail,omitempty" tf:"is_organization_trail"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Name string `json:"name" tf:"name"`

	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`

	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix"`

	SnsTopicName *string `json:"snsTopicName,omitempty" tf:"sns_topic_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type DataResourceObservation struct {
}

type DataResourceParameters struct {
	Type string `json:"type" tf:"type"`

	Values []string `json:"values" tf:"values"`
}

type EventSelectorObservation struct {
}

type EventSelectorParameters struct {
	DataResource []DataResourceParameters `json:"dataResource,omitempty" tf:"data_resource"`

	IncludeManagementEvents *bool `json:"includeManagementEvents,omitempty" tf:"include_management_events"`

	ReadWriteType *string `json:"readWriteType,omitempty" tf:"read_write_type"`
}

type InsightSelectorObservation struct {
}

type InsightSelectorParameters struct {
	InsightType string `json:"insightType" tf:"insight_type"`
}

// CloudtrailSpec defines the desired state of Cloudtrail
type CloudtrailSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudtrailParameters `json:"forProvider"`
}

// CloudtrailStatus defines the observed state of Cloudtrail.
type CloudtrailStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudtrailObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cloudtrail is the Schema for the Cloudtrails API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Cloudtrail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudtrailSpec   `json:"spec"`
	Status            CloudtrailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudtrailList contains a list of Cloudtrails
type CloudtrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cloudtrail `json:"items"`
}

// Repository type metadata.
var (
	CloudtrailKind             = "Cloudtrail"
	CloudtrailGroupKind        = schema.GroupKind{Group: Group, Kind: CloudtrailKind}.String()
	CloudtrailKindAPIVersion   = CloudtrailKind + "." + GroupVersion.String()
	CloudtrailGroupVersionKind = GroupVersion.WithKind(CloudtrailKind)
)

func init() {
	SchemeBuilder.Register(&Cloudtrail{}, &CloudtrailList{})
}
