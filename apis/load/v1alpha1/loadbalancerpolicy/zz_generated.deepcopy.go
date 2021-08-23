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
func (in *LoadBalancerPolicy) DeepCopyInto(out *LoadBalancerPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPolicy.
func (in *LoadBalancerPolicy) DeepCopy() *LoadBalancerPolicy {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerPolicyList) DeepCopyInto(out *LoadBalancerPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPolicyList.
func (in *LoadBalancerPolicyList) DeepCopy() *LoadBalancerPolicyList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerPolicyObservation) DeepCopyInto(out *LoadBalancerPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPolicyObservation.
func (in *LoadBalancerPolicyObservation) DeepCopy() *LoadBalancerPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerPolicyParameters) DeepCopyInto(out *LoadBalancerPolicyParameters) {
	*out = *in
	if in.PolicyAttribute != nil {
		in, out := &in.PolicyAttribute, &out.PolicyAttribute
		*out = make([]PolicyAttributeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPolicyParameters.
func (in *LoadBalancerPolicyParameters) DeepCopy() *LoadBalancerPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerPolicySpec) DeepCopyInto(out *LoadBalancerPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPolicySpec.
func (in *LoadBalancerPolicySpec) DeepCopy() *LoadBalancerPolicySpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerPolicyStatus) DeepCopyInto(out *LoadBalancerPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPolicyStatus.
func (in *LoadBalancerPolicyStatus) DeepCopy() *LoadBalancerPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAttributeObservation) DeepCopyInto(out *PolicyAttributeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAttributeObservation.
func (in *PolicyAttributeObservation) DeepCopy() *PolicyAttributeObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyAttributeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAttributeParameters) DeepCopyInto(out *PolicyAttributeParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAttributeParameters.
func (in *PolicyAttributeParameters) DeepCopy() *PolicyAttributeParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyAttributeParameters)
	in.DeepCopyInto(out)
	return out
}