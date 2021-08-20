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

// +kubebuilder:object:generate=true
// +groupName=codedeploy.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/codedeploy/v1alpha1"
)

type AlarmConfigurationObservation struct {
}

type AlarmConfigurationParameters struct {
	Alarms []string `json:"alarms,omitempty" tf:"alarms"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	IgnorePollAlarmFailure *bool `json:"ignorePollAlarmFailure,omitempty" tf:"ignore_poll_alarm_failure"`
}

type AutoRollbackConfigurationObservation struct {
}

type AutoRollbackConfigurationParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Events []string `json:"events,omitempty" tf:"events"`
}

type BlueGreenDeploymentConfigObservation struct {
}

type BlueGreenDeploymentConfigParameters struct {
	DeploymentReadyOption []DeploymentReadyOptionParameters `json:"deploymentReadyOption,omitempty" tf:"deployment_ready_option"`

	GreenFleetProvisioningOption []GreenFleetProvisioningOptionParameters `json:"greenFleetProvisioningOption,omitempty" tf:"green_fleet_provisioning_option"`

	TerminateBlueInstancesOnDeploymentSuccess []TerminateBlueInstancesOnDeploymentSuccessParameters `json:"terminateBlueInstancesOnDeploymentSuccess,omitempty" tf:"terminate_blue_instances_on_deployment_success"`
}

type CodedeployDeploymentGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ComputePlatform string `json:"computePlatform" tf:"compute_platform"`

	DeploymentGroupId string `json:"deploymentGroupId" tf:"deployment_group_id"`
}

type CodedeployDeploymentGroupParameters struct {
	AlarmConfiguration []AlarmConfigurationParameters `json:"alarmConfiguration,omitempty" tf:"alarm_configuration"`

	AppName string `json:"appName" tf:"app_name"`

	AutoRollbackConfiguration []AutoRollbackConfigurationParameters `json:"autoRollbackConfiguration,omitempty" tf:"auto_rollback_configuration"`

	AutoscalingGroups []string `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups"`

	BlueGreenDeploymentConfig []BlueGreenDeploymentConfigParameters `json:"blueGreenDeploymentConfig,omitempty" tf:"blue_green_deployment_config"`

	DeploymentConfigName *string `json:"deploymentConfigName,omitempty" tf:"deployment_config_name"`

	DeploymentGroupName string `json:"deploymentGroupName" tf:"deployment_group_name"`

	DeploymentStyle []DeploymentStyleParameters `json:"deploymentStyle,omitempty" tf:"deployment_style"`

	Ec2TagFilter []Ec2TagFilterParameters `json:"ec2TagFilter,omitempty" tf:"ec2_tag_filter"`

	Ec2TagSet []Ec2TagSetParameters `json:"ec2TagSet,omitempty" tf:"ec2_tag_set"`

	EcsService []EcsServiceParameters `json:"ecsService,omitempty" tf:"ecs_service"`

	LoadBalancerInfo []LoadBalancerInfoParameters `json:"loadBalancerInfo,omitempty" tf:"load_balancer_info"`

	OnPremisesInstanceTagFilter []OnPremisesInstanceTagFilterParameters `json:"onPremisesInstanceTagFilter,omitempty" tf:"on_premises_instance_tag_filter"`

	ServiceRoleArn string `json:"serviceRoleArn" tf:"service_role_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TriggerConfiguration []TriggerConfigurationParameters `json:"triggerConfiguration,omitempty" tf:"trigger_configuration"`
}

type DeploymentReadyOptionObservation struct {
}

type DeploymentReadyOptionParameters struct {
	ActionOnTimeout *string `json:"actionOnTimeout,omitempty" tf:"action_on_timeout"`

	WaitTimeInMinutes *int64 `json:"waitTimeInMinutes,omitempty" tf:"wait_time_in_minutes"`
}

type DeploymentStyleObservation struct {
}

type DeploymentStyleParameters struct {
	DeploymentOption *string `json:"deploymentOption,omitempty" tf:"deployment_option"`

	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type"`
}

type Ec2TagFilterObservation struct {
}

type Ec2TagFilterParameters struct {
	Key *string `json:"key,omitempty" tf:"key"`

	Type *string `json:"type,omitempty" tf:"type"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type Ec2TagSetEc2TagFilterObservation struct {
}

type Ec2TagSetEc2TagFilterParameters struct {
	Key *string `json:"key,omitempty" tf:"key"`

	Type *string `json:"type,omitempty" tf:"type"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type Ec2TagSetObservation struct {
}

type Ec2TagSetParameters struct {
	Ec2TagFilter []Ec2TagSetEc2TagFilterParameters `json:"ec2TagFilter,omitempty" tf:"ec2_tag_filter"`
}

type EcsServiceObservation struct {
}

type EcsServiceParameters struct {
	ClusterName string `json:"clusterName" tf:"cluster_name"`

	ServiceName string `json:"serviceName" tf:"service_name"`
}

type ElbInfoObservation struct {
}

type ElbInfoParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`
}

type GreenFleetProvisioningOptionObservation struct {
}

type GreenFleetProvisioningOptionParameters struct {
	Action *string `json:"action,omitempty" tf:"action"`
}

type LoadBalancerInfoObservation struct {
}

type LoadBalancerInfoParameters struct {
	ElbInfo []ElbInfoParameters `json:"elbInfo,omitempty" tf:"elb_info"`

	TargetGroupInfo []TargetGroupInfoParameters `json:"targetGroupInfo,omitempty" tf:"target_group_info"`

	TargetGroupPairInfo []TargetGroupPairInfoParameters `json:"targetGroupPairInfo,omitempty" tf:"target_group_pair_info"`
}

type OnPremisesInstanceTagFilterObservation struct {
}

type OnPremisesInstanceTagFilterParameters struct {
	Key *string `json:"key,omitempty" tf:"key"`

	Type *string `json:"type,omitempty" tf:"type"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type ProdTrafficRouteObservation struct {
}

type ProdTrafficRouteParameters struct {
	ListenerArns []string `json:"listenerArns" tf:"listener_arns"`
}

type TargetGroupInfoObservation struct {
}

type TargetGroupInfoParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`
}

type TargetGroupObservation struct {
}

type TargetGroupPairInfoObservation struct {
}

type TargetGroupPairInfoParameters struct {
	ProdTrafficRoute []ProdTrafficRouteParameters `json:"prodTrafficRoute" tf:"prod_traffic_route"`

	TargetGroup []TargetGroupParameters `json:"targetGroup" tf:"target_group"`

	TestTrafficRoute []TestTrafficRouteParameters `json:"testTrafficRoute,omitempty" tf:"test_traffic_route"`
}

type TargetGroupParameters struct {
	Name string `json:"name" tf:"name"`
}

type TerminateBlueInstancesOnDeploymentSuccessObservation struct {
}

type TerminateBlueInstancesOnDeploymentSuccessParameters struct {
	Action *string `json:"action,omitempty" tf:"action"`

	TerminationWaitTimeInMinutes *int64 `json:"terminationWaitTimeInMinutes,omitempty" tf:"termination_wait_time_in_minutes"`
}

type TestTrafficRouteObservation struct {
}

type TestTrafficRouteParameters struct {
	ListenerArns []string `json:"listenerArns" tf:"listener_arns"`
}

type TriggerConfigurationObservation struct {
}

type TriggerConfigurationParameters struct {
	TriggerEvents []string `json:"triggerEvents" tf:"trigger_events"`

	TriggerName string `json:"triggerName" tf:"trigger_name"`

	TriggerTargetArn string `json:"triggerTargetArn" tf:"trigger_target_arn"`
}

// CodedeployDeploymentGroupSpec defines the desired state of CodedeployDeploymentGroup
type CodedeployDeploymentGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodedeployDeploymentGroupParameters `json:"forProvider"`
}

// CodedeployDeploymentGroupStatus defines the observed state of CodedeployDeploymentGroup.
type CodedeployDeploymentGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodedeployDeploymentGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployDeploymentGroup is the Schema for the CodedeployDeploymentGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodedeployDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployDeploymentGroupSpec   `json:"spec"`
	Status            CodedeployDeploymentGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployDeploymentGroupList contains a list of CodedeployDeploymentGroups
type CodedeployDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodedeployDeploymentGroup `json:"items"`
}

// Repository type metadata.
var (
	CodedeployDeploymentGroupKind             = "CodedeployDeploymentGroup"
	CodedeployDeploymentGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CodedeployDeploymentGroupKind}.String()
	CodedeployDeploymentGroupKindAPIVersion   = CodedeployDeploymentGroupKind + "." + v1alpha1.GroupVersion.String()
	CodedeployDeploymentGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(CodedeployDeploymentGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CodedeployDeploymentGroup{}, &CodedeployDeploymentGroupList{})
}
