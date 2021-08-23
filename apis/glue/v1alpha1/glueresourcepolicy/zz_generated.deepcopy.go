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
func (in *GlueResourcePolicy) DeepCopyInto(out *GlueResourcePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueResourcePolicy.
func (in *GlueResourcePolicy) DeepCopy() *GlueResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(GlueResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueResourcePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueResourcePolicyList) DeepCopyInto(out *GlueResourcePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlueResourcePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueResourcePolicyList.
func (in *GlueResourcePolicyList) DeepCopy() *GlueResourcePolicyList {
	if in == nil {
		return nil
	}
	out := new(GlueResourcePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueResourcePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueResourcePolicyObservation) DeepCopyInto(out *GlueResourcePolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueResourcePolicyObservation.
func (in *GlueResourcePolicyObservation) DeepCopy() *GlueResourcePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(GlueResourcePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueResourcePolicyParameters) DeepCopyInto(out *GlueResourcePolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueResourcePolicyParameters.
func (in *GlueResourcePolicyParameters) DeepCopy() *GlueResourcePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(GlueResourcePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueResourcePolicySpec) DeepCopyInto(out *GlueResourcePolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueResourcePolicySpec.
func (in *GlueResourcePolicySpec) DeepCopy() *GlueResourcePolicySpec {
	if in == nil {
		return nil
	}
	out := new(GlueResourcePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueResourcePolicyStatus) DeepCopyInto(out *GlueResourcePolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueResourcePolicyStatus.
func (in *GlueResourcePolicyStatus) DeepCopy() *GlueResourcePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(GlueResourcePolicyStatus)
	in.DeepCopyInto(out)
	return out
}