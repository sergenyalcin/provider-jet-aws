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
func (in *Vpc) DeepCopyInto(out *Vpc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vpc.
func (in *Vpc) DeepCopy() *Vpc {
	if in == nil {
		return nil
	}
	out := new(Vpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vpc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcList) DeepCopyInto(out *VpcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vpc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcList.
func (in *VpcList) DeepCopy() *VpcList {
	if in == nil {
		return nil
	}
	out := new(VpcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcObservation) DeepCopyInto(out *VpcObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DefaultNetworkACLID != nil {
		in, out := &in.DefaultNetworkACLID, &out.DefaultNetworkACLID
		*out = new(string)
		**out = **in
	}
	if in.DefaultRouteTableID != nil {
		in, out := &in.DefaultRouteTableID, &out.DefaultRouteTableID
		*out = new(string)
		**out = **in
	}
	if in.DefaultSecurityGroupID != nil {
		in, out := &in.DefaultSecurityGroupID, &out.DefaultSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.DhcpOptionsID != nil {
		in, out := &in.DhcpOptionsID, &out.DhcpOptionsID
		*out = new(string)
		**out = **in
	}
	if in.IPv6AssociationID != nil {
		in, out := &in.IPv6AssociationID, &out.IPv6AssociationID
		*out = new(string)
		**out = **in
	}
	if in.IPv6CidrBlock != nil {
		in, out := &in.IPv6CidrBlock, &out.IPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.MainRouteTableID != nil {
		in, out := &in.MainRouteTableID, &out.MainRouteTableID
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcObservation.
func (in *VpcObservation) DeepCopy() *VpcObservation {
	if in == nil {
		return nil
	}
	out := new(VpcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcParameters) DeepCopyInto(out *VpcParameters) {
	*out = *in
	if in.AssignGeneratedIPv6CidrBlock != nil {
		in, out := &in.AssignGeneratedIPv6CidrBlock, &out.AssignGeneratedIPv6CidrBlock
		*out = new(bool)
		**out = **in
	}
	if in.CidrBlock != nil {
		in, out := &in.CidrBlock, &out.CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.EnableClassiclink != nil {
		in, out := &in.EnableClassiclink, &out.EnableClassiclink
		*out = new(bool)
		**out = **in
	}
	if in.EnableClassiclinkDNSSupport != nil {
		in, out := &in.EnableClassiclinkDNSSupport, &out.EnableClassiclinkDNSSupport
		*out = new(bool)
		**out = **in
	}
	if in.EnableDNSHostnames != nil {
		in, out := &in.EnableDNSHostnames, &out.EnableDNSHostnames
		*out = new(bool)
		**out = **in
	}
	if in.EnableDNSSupport != nil {
		in, out := &in.EnableDNSSupport, &out.EnableDNSSupport
		*out = new(bool)
		**out = **in
	}
	if in.InstanceTenancy != nil {
		in, out := &in.InstanceTenancy, &out.InstanceTenancy
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcParameters.
func (in *VpcParameters) DeepCopy() *VpcParameters {
	if in == nil {
		return nil
	}
	out := new(VpcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcSpec) DeepCopyInto(out *VpcSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcSpec.
func (in *VpcSpec) DeepCopy() *VpcSpec {
	if in == nil {
		return nil
	}
	out := new(VpcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcStatus) DeepCopyInto(out *VpcStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcStatus.
func (in *VpcStatus) DeepCopy() *VpcStatus {
	if in == nil {
		return nil
	}
	out := new(VpcStatus)
	in.DeepCopyInto(out)
	return out
}
