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
func (in *Ec2TransitGateway) DeepCopyInto(out *Ec2TransitGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGateway.
func (in *Ec2TransitGateway) DeepCopy() *Ec2TransitGateway {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayList) DeepCopyInto(out *Ec2TransitGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2TransitGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayList.
func (in *Ec2TransitGatewayList) DeepCopy() *Ec2TransitGatewayList {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayObservation) DeepCopyInto(out *Ec2TransitGatewayObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayObservation.
func (in *Ec2TransitGatewayObservation) DeepCopy() *Ec2TransitGatewayObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayParameters) DeepCopyInto(out *Ec2TransitGatewayParameters) {
	*out = *in
	if in.AmazonSideAsn != nil {
		in, out := &in.AmazonSideAsn, &out.AmazonSideAsn
		*out = new(int64)
		**out = **in
	}
	if in.AutoAcceptSharedAttachments != nil {
		in, out := &in.AutoAcceptSharedAttachments, &out.AutoAcceptSharedAttachments
		*out = new(string)
		**out = **in
	}
	if in.DefaultRouteTableAssociation != nil {
		in, out := &in.DefaultRouteTableAssociation, &out.DefaultRouteTableAssociation
		*out = new(string)
		**out = **in
	}
	if in.DefaultRouteTablePropagation != nil {
		in, out := &in.DefaultRouteTablePropagation, &out.DefaultRouteTablePropagation
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DnsSupport != nil {
		in, out := &in.DnsSupport, &out.DnsSupport
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
	if in.VpnEcmpSupport != nil {
		in, out := &in.VpnEcmpSupport, &out.VpnEcmpSupport
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayParameters.
func (in *Ec2TransitGatewayParameters) DeepCopy() *Ec2TransitGatewayParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewaySpec) DeepCopyInto(out *Ec2TransitGatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewaySpec.
func (in *Ec2TransitGatewaySpec) DeepCopy() *Ec2TransitGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayStatus) DeepCopyInto(out *Ec2TransitGatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayStatus.
func (in *Ec2TransitGatewayStatus) DeepCopy() *Ec2TransitGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayStatus)
	in.DeepCopyInto(out)
	return out
}