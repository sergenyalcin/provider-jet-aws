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
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotId != nil {
		in, out := &in.SnapshotId, &out.SnapshotId
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
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
func (in *LaunchConfiguration) DeepCopyInto(out *LaunchConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfiguration.
func (in *LaunchConfiguration) DeepCopy() *LaunchConfiguration {
	if in == nil {
		return nil
	}
	out := new(LaunchConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationList) DeepCopyInto(out *LaunchConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LaunchConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationList.
func (in *LaunchConfigurationList) DeepCopy() *LaunchConfigurationList {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationObservation) DeepCopyInto(out *LaunchConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationObservation.
func (in *LaunchConfigurationObservation) DeepCopy() *LaunchConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationParameters) DeepCopyInto(out *LaunchConfigurationParameters) {
	*out = *in
	if in.AssociatePublicIpAddress != nil {
		in, out := &in.AssociatePublicIpAddress, &out.AssociatePublicIpAddress
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
	if in.EnableMonitoring != nil {
		in, out := &in.EnableMonitoring, &out.EnableMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]EphemeralBlockDeviceParameters, len(*in))
		copy(*out, *in)
	}
	if in.IamInstanceProfile != nil {
		in, out := &in.IamInstanceProfile, &out.IamInstanceProfile
		*out = new(string)
		**out = **in
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.MetadataOptions != nil {
		in, out := &in.MetadataOptions, &out.MetadataOptions
		*out = make([]MetadataOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.PlacementTenancy != nil {
		in, out := &in.PlacementTenancy, &out.PlacementTenancy
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
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SpotPrice != nil {
		in, out := &in.SpotPrice, &out.SpotPrice
		*out = new(string)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.UserDataBase64 != nil {
		in, out := &in.UserDataBase64, &out.UserDataBase64
		*out = new(string)
		**out = **in
	}
	if in.VpcClassicLinkId != nil {
		in, out := &in.VpcClassicLinkId, &out.VpcClassicLinkId
		*out = new(string)
		**out = **in
	}
	if in.VpcClassicLinkSecurityGroups != nil {
		in, out := &in.VpcClassicLinkSecurityGroups, &out.VpcClassicLinkSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationParameters.
func (in *LaunchConfigurationParameters) DeepCopy() *LaunchConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationSpec) DeepCopyInto(out *LaunchConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationSpec.
func (in *LaunchConfigurationSpec) DeepCopy() *LaunchConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationStatus) DeepCopyInto(out *LaunchConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationStatus.
func (in *LaunchConfigurationStatus) DeepCopy() *LaunchConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptionsObservation) DeepCopyInto(out *MetadataOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptionsObservation.
func (in *MetadataOptionsObservation) DeepCopy() *MetadataOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(MetadataOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptionsParameters) DeepCopyInto(out *MetadataOptionsParameters) {
	*out = *in
	if in.HttpEndpoint != nil {
		in, out := &in.HttpEndpoint, &out.HttpEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HttpPutResponseHopLimit != nil {
		in, out := &in.HttpPutResponseHopLimit, &out.HttpPutResponseHopLimit
		*out = new(int64)
		**out = **in
	}
	if in.HttpTokens != nil {
		in, out := &in.HttpTokens, &out.HttpTokens
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptionsParameters.
func (in *MetadataOptionsParameters) DeepCopy() *MetadataOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(MetadataOptionsParameters)
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
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
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