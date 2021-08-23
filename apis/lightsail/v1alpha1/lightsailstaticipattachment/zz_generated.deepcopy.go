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
func (in *LightsailStaticIpAttachment) DeepCopyInto(out *LightsailStaticIpAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailStaticIpAttachment.
func (in *LightsailStaticIpAttachment) DeepCopy() *LightsailStaticIpAttachment {
	if in == nil {
		return nil
	}
	out := new(LightsailStaticIpAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LightsailStaticIpAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailStaticIpAttachmentList) DeepCopyInto(out *LightsailStaticIpAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LightsailStaticIpAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailStaticIpAttachmentList.
func (in *LightsailStaticIpAttachmentList) DeepCopy() *LightsailStaticIpAttachmentList {
	if in == nil {
		return nil
	}
	out := new(LightsailStaticIpAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LightsailStaticIpAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailStaticIpAttachmentObservation) DeepCopyInto(out *LightsailStaticIpAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailStaticIpAttachmentObservation.
func (in *LightsailStaticIpAttachmentObservation) DeepCopy() *LightsailStaticIpAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(LightsailStaticIpAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailStaticIpAttachmentParameters) DeepCopyInto(out *LightsailStaticIpAttachmentParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailStaticIpAttachmentParameters.
func (in *LightsailStaticIpAttachmentParameters) DeepCopy() *LightsailStaticIpAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(LightsailStaticIpAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailStaticIpAttachmentSpec) DeepCopyInto(out *LightsailStaticIpAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailStaticIpAttachmentSpec.
func (in *LightsailStaticIpAttachmentSpec) DeepCopy() *LightsailStaticIpAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(LightsailStaticIpAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LightsailStaticIpAttachmentStatus) DeepCopyInto(out *LightsailStaticIpAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LightsailStaticIpAttachmentStatus.
func (in *LightsailStaticIpAttachmentStatus) DeepCopy() *LightsailStaticIpAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(LightsailStaticIpAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}