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
func (in *EbsBlockDeviceObservation) DeepCopyInto(out *EbsBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsBlockDeviceObservation.
func (in *EbsBlockDeviceObservation) DeepCopy() *EbsBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(EbsBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsBlockDeviceParameters) DeepCopyInto(out *EbsBlockDeviceParameters) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.SnapshotId != nil {
		in, out := &in.SnapshotId, &out.SnapshotId
		*out = new(string)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsBlockDeviceParameters.
func (in *EbsBlockDeviceParameters) DeepCopy() *EbsBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(EbsBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralBlockDeviceObservation) DeepCopyInto(out *EphemeralBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralBlockDeviceObservation.
func (in *EphemeralBlockDeviceObservation) DeepCopy() *EphemeralBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(EphemeralBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralBlockDeviceParameters) DeepCopyInto(out *EphemeralBlockDeviceParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralBlockDeviceParameters.
func (in *EphemeralBlockDeviceParameters) DeepCopy() *EphemeralBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(EphemeralBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksInstance) DeepCopyInto(out *OpsworksInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksInstance.
func (in *OpsworksInstance) DeepCopy() *OpsworksInstance {
	if in == nil {
		return nil
	}
	out := new(OpsworksInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksInstanceList) DeepCopyInto(out *OpsworksInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpsworksInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksInstanceList.
func (in *OpsworksInstanceList) DeepCopy() *OpsworksInstanceList {
	if in == nil {
		return nil
	}
	out := new(OpsworksInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksInstanceObservation) DeepCopyInto(out *OpsworksInstanceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksInstanceObservation.
func (in *OpsworksInstanceObservation) DeepCopy() *OpsworksInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(OpsworksInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksInstanceParameters) DeepCopyInto(out *OpsworksInstanceParameters) {
	*out = *in
	if in.AgentVersion != nil {
		in, out := &in.AgentVersion, &out.AgentVersion
		*out = new(string)
		**out = **in
	}
	if in.AmiId != nil {
		in, out := &in.AmiId, &out.AmiId
		*out = new(string)
		**out = **in
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.AutoScalingType != nil {
		in, out := &in.AutoScalingType, &out.AutoScalingType
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.DeleteEbs != nil {
		in, out := &in.DeleteEbs, &out.DeleteEbs
		*out = new(bool)
		**out = **in
	}
	if in.DeleteEip != nil {
		in, out := &in.DeleteEip, &out.DeleteEip
		*out = new(bool)
		**out = **in
	}
	if in.EbsBlockDevice != nil {
		in, out := &in.EbsBlockDevice, &out.EbsBlockDevice
		*out = make([]EbsBlockDeviceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EbsOptimized != nil {
		in, out := &in.EbsOptimized, &out.EbsOptimized
		*out = new(bool)
		**out = **in
	}
	if in.EcsClusterArn != nil {
		in, out := &in.EcsClusterArn, &out.EcsClusterArn
		*out = new(string)
		**out = **in
	}
	if in.ElasticIp != nil {
		in, out := &in.ElasticIp, &out.ElasticIp
		*out = new(string)
		**out = **in
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]EphemeralBlockDeviceParameters, len(*in))
		copy(*out, *in)
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.InfrastructureClass != nil {
		in, out := &in.InfrastructureClass, &out.InfrastructureClass
		*out = new(string)
		**out = **in
	}
	if in.InstallUpdatesOnBoot != nil {
		in, out := &in.InstallUpdatesOnBoot, &out.InstallUpdatesOnBoot
		*out = new(bool)
		**out = **in
	}
	if in.InstanceProfileArn != nil {
		in, out := &in.InstanceProfileArn, &out.InstanceProfileArn
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.LastServiceErrorId != nil {
		in, out := &in.LastServiceErrorId, &out.LastServiceErrorId
		*out = new(string)
		**out = **in
	}
	if in.LayerIds != nil {
		in, out := &in.LayerIds, &out.LayerIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Os != nil {
		in, out := &in.Os, &out.Os
		*out = new(string)
		**out = **in
	}
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(string)
		**out = **in
	}
	if in.PrivateDns != nil {
		in, out := &in.PrivateDns, &out.PrivateDns
		*out = new(string)
		**out = **in
	}
	if in.PrivateIp != nil {
		in, out := &in.PrivateIp, &out.PrivateIp
		*out = new(string)
		**out = **in
	}
	if in.PublicDns != nil {
		in, out := &in.PublicDns, &out.PublicDns
		*out = new(string)
		**out = **in
	}
	if in.PublicIp != nil {
		in, out := &in.PublicIp, &out.PublicIp
		*out = new(string)
		**out = **in
	}
	if in.RegisteredBy != nil {
		in, out := &in.RegisteredBy, &out.RegisteredBy
		*out = new(string)
		**out = **in
	}
	if in.ReportedAgentVersion != nil {
		in, out := &in.ReportedAgentVersion, &out.ReportedAgentVersion
		*out = new(string)
		**out = **in
	}
	if in.ReportedOsFamily != nil {
		in, out := &in.ReportedOsFamily, &out.ReportedOsFamily
		*out = new(string)
		**out = **in
	}
	if in.ReportedOsName != nil {
		in, out := &in.ReportedOsName, &out.ReportedOsName
		*out = new(string)
		**out = **in
	}
	if in.ReportedOsVersion != nil {
		in, out := &in.ReportedOsVersion, &out.ReportedOsVersion
		*out = new(string)
		**out = **in
	}
	if in.RootBlockDevice != nil {
		in, out := &in.RootBlockDevice, &out.RootBlockDevice
		*out = make([]RootBlockDeviceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RootDeviceType != nil {
		in, out := &in.RootDeviceType, &out.RootDeviceType
		*out = new(string)
		**out = **in
	}
	if in.RootDeviceVolumeId != nil {
		in, out := &in.RootDeviceVolumeId, &out.RootDeviceVolumeId
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SshHostDsaKeyFingerprint != nil {
		in, out := &in.SshHostDsaKeyFingerprint, &out.SshHostDsaKeyFingerprint
		*out = new(string)
		**out = **in
	}
	if in.SshHostRsaKeyFingerprint != nil {
		in, out := &in.SshHostRsaKeyFingerprint, &out.SshHostRsaKeyFingerprint
		*out = new(string)
		**out = **in
	}
	if in.SshKeyName != nil {
		in, out := &in.SshKeyName, &out.SshKeyName
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
		*out = new(string)
		**out = **in
	}
	if in.Tenancy != nil {
		in, out := &in.Tenancy, &out.Tenancy
		*out = new(string)
		**out = **in
	}
	if in.VirtualizationType != nil {
		in, out := &in.VirtualizationType, &out.VirtualizationType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksInstanceParameters.
func (in *OpsworksInstanceParameters) DeepCopy() *OpsworksInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(OpsworksInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksInstanceSpec) DeepCopyInto(out *OpsworksInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksInstanceSpec.
func (in *OpsworksInstanceSpec) DeepCopy() *OpsworksInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(OpsworksInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksInstanceStatus) DeepCopyInto(out *OpsworksInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksInstanceStatus.
func (in *OpsworksInstanceStatus) DeepCopy() *OpsworksInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(OpsworksInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootBlockDeviceObservation) DeepCopyInto(out *RootBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootBlockDeviceObservation.
func (in *RootBlockDeviceObservation) DeepCopy() *RootBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(RootBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootBlockDeviceParameters) DeepCopyInto(out *RootBlockDeviceParameters) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootBlockDeviceParameters.
func (in *RootBlockDeviceParameters) DeepCopy() *RootBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(RootBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}