// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheNodesObservation) DeepCopyInto(out *CacheNodesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheNodesObservation.
func (in *CacheNodesObservation) DeepCopy() *CacheNodesObservation {
	if in == nil {
		return nil
	}
	out := new(CacheNodesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheNodesParameters) DeepCopyInto(out *CacheNodesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheNodesParameters.
func (in *CacheNodesParameters) DeepCopy() *CacheNodesParameters {
	if in == nil {
		return nil
	}
	out := new(CacheNodesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheCluster) DeepCopyInto(out *ElasticacheCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheCluster.
func (in *ElasticacheCluster) DeepCopy() *ElasticacheCluster {
	if in == nil {
		return nil
	}
	out := new(ElasticacheCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticacheCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterList) DeepCopyInto(out *ElasticacheClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticacheCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterList.
func (in *ElasticacheClusterList) DeepCopy() *ElasticacheClusterList {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticacheClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterObservation) DeepCopyInto(out *ElasticacheClusterObservation) {
	*out = *in
	if in.CacheNodes != nil {
		in, out := &in.CacheNodes, &out.CacheNodes
		*out = make([]CacheNodesObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterObservation.
func (in *ElasticacheClusterObservation) DeepCopy() *ElasticacheClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterParameters) DeepCopyInto(out *ElasticacheClusterParameters) {
	*out = *in
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.AzMode != nil {
		in, out := &in.AzMode, &out.AzMode
		*out = new(string)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.FinalSnapshotIdentifier != nil {
		in, out := &in.FinalSnapshotIdentifier, &out.FinalSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.NotificationTopicArn != nil {
		in, out := &in.NotificationTopicArn, &out.NotificationTopicArn
		*out = new(string)
		**out = **in
	}
	if in.NumCacheNodes != nil {
		in, out := &in.NumCacheNodes, &out.NumCacheNodes
		*out = new(int64)
		**out = **in
	}
	if in.ParameterGroupName != nil {
		in, out := &in.ParameterGroupName, &out.ParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.PreferredAvailabilityZones != nil {
		in, out := &in.PreferredAvailabilityZones, &out.PreferredAvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplicationGroupId != nil {
		in, out := &in.ReplicationGroupId, &out.ReplicationGroupId
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupNames != nil {
		in, out := &in.SecurityGroupNames, &out.SecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotArns != nil {
		in, out := &in.SnapshotArns, &out.SnapshotArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotName != nil {
		in, out := &in.SnapshotName, &out.SnapshotName
		*out = new(string)
		**out = **in
	}
	if in.SnapshotRetentionLimit != nil {
		in, out := &in.SnapshotRetentionLimit, &out.SnapshotRetentionLimit
		*out = new(int64)
		**out = **in
	}
	if in.SnapshotWindow != nil {
		in, out := &in.SnapshotWindow, &out.SnapshotWindow
		*out = new(string)
		**out = **in
	}
	if in.SubnetGroupName != nil {
		in, out := &in.SubnetGroupName, &out.SubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterParameters.
func (in *ElasticacheClusterParameters) DeepCopy() *ElasticacheClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterSpec) DeepCopyInto(out *ElasticacheClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterSpec.
func (in *ElasticacheClusterSpec) DeepCopy() *ElasticacheClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterStatus) DeepCopyInto(out *ElasticacheClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterStatus.
func (in *ElasticacheClusterStatus) DeepCopy() *ElasticacheClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterStatus)
	in.DeepCopyInto(out)
	return out
}