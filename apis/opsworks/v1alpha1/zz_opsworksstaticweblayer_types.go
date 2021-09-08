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

type OpsworksStaticWebLayerEbsVolumeObservation struct {
}

type OpsworksStaticWebLayerEbsVolumeParameters struct {
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	MountPoint string `json:"mountPoint" tf:"mount_point"`

	NumberOfDisks int64 `json:"numberOfDisks" tf:"number_of_disks"`

	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level"`

	Size int64 `json:"size" tf:"size"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type OpsworksStaticWebLayerObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type OpsworksStaticWebLayerParameters struct {
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips"`

	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips"`

	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing"`

	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes"`

	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes"`

	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn"`

	CustomJson *string `json:"customJson,omitempty" tf:"custom_json"`

	CustomSecurityGroupIds []string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids"`

	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes"`

	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes"`

	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes"`

	DrainElbOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown"`

	EbsVolume []OpsworksStaticWebLayerEbsVolumeParameters `json:"ebsVolume,omitempty" tf:"ebs_volume"`

	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer"`

	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot"`

	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout"`

	Name *string `json:"name,omitempty" tf:"name"`

	StackId string `json:"stackId" tf:"stack_id"`

	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	UseEbsOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances"`
}

// OpsworksStaticWebLayerSpec defines the desired state of OpsworksStaticWebLayer
type OpsworksStaticWebLayerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OpsworksStaticWebLayerParameters `json:"forProvider"`
}

// OpsworksStaticWebLayerStatus defines the observed state of OpsworksStaticWebLayer.
type OpsworksStaticWebLayerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OpsworksStaticWebLayerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksStaticWebLayer is the Schema for the OpsworksStaticWebLayers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type OpsworksStaticWebLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksStaticWebLayerSpec   `json:"spec"`
	Status            OpsworksStaticWebLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksStaticWebLayerList contains a list of OpsworksStaticWebLayers
type OpsworksStaticWebLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksStaticWebLayer `json:"items"`
}

// Repository type metadata.
var (
	OpsworksStaticWebLayerKind             = "OpsworksStaticWebLayer"
	OpsworksStaticWebLayerGroupKind        = schema.GroupKind{Group: Group, Kind: OpsworksStaticWebLayerKind}.String()
	OpsworksStaticWebLayerKindAPIVersion   = OpsworksStaticWebLayerKind + "." + GroupVersion.String()
	OpsworksStaticWebLayerGroupVersionKind = GroupVersion.WithKind(OpsworksStaticWebLayerKind)
)

func init() {
	SchemeBuilder.Register(&OpsworksStaticWebLayer{}, &OpsworksStaticWebLayerList{})
}
