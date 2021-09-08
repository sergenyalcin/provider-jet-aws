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
func (in *NetworkInterfacesObservation) DeepCopyInto(out *NetworkInterfacesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfacesObservation.
func (in *NetworkInterfacesObservation) DeepCopy() *NetworkInterfacesObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfacesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfacesParameters) DeepCopyInto(out *NetworkInterfacesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfacesParameters.
func (in *NetworkInterfacesParameters) DeepCopy() *NetworkInterfacesParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfacesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3OutpostsEndpoint) DeepCopyInto(out *S3OutpostsEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3OutpostsEndpoint.
func (in *S3OutpostsEndpoint) DeepCopy() *S3OutpostsEndpoint {
	if in == nil {
		return nil
	}
	out := new(S3OutpostsEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3OutpostsEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3OutpostsEndpointList) DeepCopyInto(out *S3OutpostsEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3OutpostsEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3OutpostsEndpointList.
func (in *S3OutpostsEndpointList) DeepCopy() *S3OutpostsEndpointList {
	if in == nil {
		return nil
	}
	out := new(S3OutpostsEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3OutpostsEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3OutpostsEndpointObservation) DeepCopyInto(out *S3OutpostsEndpointObservation) {
	*out = *in
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make([]NetworkInterfacesObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3OutpostsEndpointObservation.
func (in *S3OutpostsEndpointObservation) DeepCopy() *S3OutpostsEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(S3OutpostsEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3OutpostsEndpointParameters) DeepCopyInto(out *S3OutpostsEndpointParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3OutpostsEndpointParameters.
func (in *S3OutpostsEndpointParameters) DeepCopy() *S3OutpostsEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(S3OutpostsEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3OutpostsEndpointSpec) DeepCopyInto(out *S3OutpostsEndpointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3OutpostsEndpointSpec.
func (in *S3OutpostsEndpointSpec) DeepCopy() *S3OutpostsEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(S3OutpostsEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3OutpostsEndpointStatus) DeepCopyInto(out *S3OutpostsEndpointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3OutpostsEndpointStatus.
func (in *S3OutpostsEndpointStatus) DeepCopy() *S3OutpostsEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(S3OutpostsEndpointStatus)
	in.DeepCopyInto(out)
	return out
}