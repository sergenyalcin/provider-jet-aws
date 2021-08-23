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
func (in *NeptuneParameterGroup) DeepCopyInto(out *NeptuneParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneParameterGroup.
func (in *NeptuneParameterGroup) DeepCopy() *NeptuneParameterGroup {
	if in == nil {
		return nil
	}
	out := new(NeptuneParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeptuneParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneParameterGroupList) DeepCopyInto(out *NeptuneParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NeptuneParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneParameterGroupList.
func (in *NeptuneParameterGroupList) DeepCopy() *NeptuneParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(NeptuneParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeptuneParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneParameterGroupObservation) DeepCopyInto(out *NeptuneParameterGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneParameterGroupObservation.
func (in *NeptuneParameterGroupObservation) DeepCopy() *NeptuneParameterGroupObservation {
	if in == nil {
		return nil
	}
	out := new(NeptuneParameterGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneParameterGroupParameters) DeepCopyInto(out *NeptuneParameterGroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]ParameterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneParameterGroupParameters.
func (in *NeptuneParameterGroupParameters) DeepCopy() *NeptuneParameterGroupParameters {
	if in == nil {
		return nil
	}
	out := new(NeptuneParameterGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneParameterGroupSpec) DeepCopyInto(out *NeptuneParameterGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneParameterGroupSpec.
func (in *NeptuneParameterGroupSpec) DeepCopy() *NeptuneParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(NeptuneParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneParameterGroupStatus) DeepCopyInto(out *NeptuneParameterGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneParameterGroupStatus.
func (in *NeptuneParameterGroupStatus) DeepCopy() *NeptuneParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(NeptuneParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterObservation) DeepCopyInto(out *ParameterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterObservation.
func (in *ParameterObservation) DeepCopy() *ParameterObservation {
	if in == nil {
		return nil
	}
	out := new(ParameterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterParameters) DeepCopyInto(out *ParameterParameters) {
	*out = *in
	if in.ApplyMethod != nil {
		in, out := &in.ApplyMethod, &out.ApplyMethod
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterParameters.
func (in *ParameterParameters) DeepCopy() *ParameterParameters {
	if in == nil {
		return nil
	}
	out := new(ParameterParameters)
	in.DeepCopyInto(out)
	return out
}