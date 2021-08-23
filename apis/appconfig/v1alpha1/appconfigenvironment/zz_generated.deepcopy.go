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
func (in *AppconfigEnvironment) DeepCopyInto(out *AppconfigEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigEnvironment.
func (in *AppconfigEnvironment) DeepCopy() *AppconfigEnvironment {
	if in == nil {
		return nil
	}
	out := new(AppconfigEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppconfigEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigEnvironmentList) DeepCopyInto(out *AppconfigEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppconfigEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigEnvironmentList.
func (in *AppconfigEnvironmentList) DeepCopy() *AppconfigEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(AppconfigEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppconfigEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigEnvironmentObservation) DeepCopyInto(out *AppconfigEnvironmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigEnvironmentObservation.
func (in *AppconfigEnvironmentObservation) DeepCopy() *AppconfigEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(AppconfigEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigEnvironmentParameters) DeepCopyInto(out *AppconfigEnvironmentParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Monitor != nil {
		in, out := &in.Monitor, &out.Monitor
		*out = make([]MonitorParameters, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigEnvironmentParameters.
func (in *AppconfigEnvironmentParameters) DeepCopy() *AppconfigEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(AppconfigEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigEnvironmentSpec) DeepCopyInto(out *AppconfigEnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigEnvironmentSpec.
func (in *AppconfigEnvironmentSpec) DeepCopy() *AppconfigEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(AppconfigEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigEnvironmentStatus) DeepCopyInto(out *AppconfigEnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigEnvironmentStatus.
func (in *AppconfigEnvironmentStatus) DeepCopy() *AppconfigEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(AppconfigEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorObservation) DeepCopyInto(out *MonitorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorObservation.
func (in *MonitorObservation) DeepCopy() *MonitorObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorParameters) DeepCopyInto(out *MonitorParameters) {
	*out = *in
	if in.AlarmRoleArn != nil {
		in, out := &in.AlarmRoleArn, &out.AlarmRoleArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorParameters.
func (in *MonitorParameters) DeepCopy() *MonitorParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorParameters)
	in.DeepCopyInto(out)
	return out
}