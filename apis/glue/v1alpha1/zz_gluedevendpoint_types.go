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

type GlueDevEndpointObservation struct {
	Arn string `json:"arn" tf:"arn"`

	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	FailureReason string `json:"failureReason" tf:"failure_reason"`

	PrivateAddress string `json:"privateAddress" tf:"private_address"`

	PublicAddress string `json:"publicAddress" tf:"public_address"`

	Status string `json:"status" tf:"status"`

	VpcId string `json:"vpcId" tf:"vpc_id"`

	YarnEndpointAddress string `json:"yarnEndpointAddress" tf:"yarn_endpoint_address"`

	ZeppelinRemoteSparkInterpreterPort int64 `json:"zeppelinRemoteSparkInterpreterPort" tf:"zeppelin_remote_spark_interpreter_port"`
}

type GlueDevEndpointParameters struct {
	Arguments map[string]string `json:"arguments,omitempty" tf:"arguments"`

	ExtraJarsS3Path *string `json:"extraJarsS3Path,omitempty" tf:"extra_jars_s3_path"`

	ExtraPythonLibsS3Path *string `json:"extraPythonLibsS3Path,omitempty" tf:"extra_python_libs_s3_path"`

	GlueVersion *string `json:"glueVersion,omitempty" tf:"glue_version"`

	Name string `json:"name" tf:"name"`

	NumberOfNodes *int64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes"`

	NumberOfWorkers *int64 `json:"numberOfWorkers,omitempty" tf:"number_of_workers"`

	PublicKey *string `json:"publicKey,omitempty" tf:"public_key"`

	PublicKeys []string `json:"publicKeys,omitempty" tf:"public_keys"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration"`

	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	WorkerType *string `json:"workerType,omitempty" tf:"worker_type"`
}

// GlueDevEndpointSpec defines the desired state of GlueDevEndpoint
type GlueDevEndpointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueDevEndpointParameters `json:"forProvider"`
}

// GlueDevEndpointStatus defines the observed state of GlueDevEndpoint.
type GlueDevEndpointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueDevEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueDevEndpoint is the Schema for the GlueDevEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueDevEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueDevEndpointSpec   `json:"spec"`
	Status            GlueDevEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueDevEndpointList contains a list of GlueDevEndpoints
type GlueDevEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueDevEndpoint `json:"items"`
}

// Repository type metadata.
var (
	GlueDevEndpointKind             = "GlueDevEndpoint"
	GlueDevEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: GlueDevEndpointKind}.String()
	GlueDevEndpointKindAPIVersion   = GlueDevEndpointKind + "." + GroupVersion.String()
	GlueDevEndpointGroupVersionKind = GroupVersion.WithKind(GlueDevEndpointKind)
)

func init() {
	SchemeBuilder.Register(&GlueDevEndpoint{}, &GlueDevEndpointList{})
}
