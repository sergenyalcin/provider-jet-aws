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
func (in *DestinationObservation) DeepCopyInto(out *DestinationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationObservation.
func (in *DestinationObservation) DeepCopy() *DestinationObservation {
	if in == nil {
		return nil
	}
	out := new(DestinationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationParameters) DeepCopyInto(out *DestinationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationParameters.
func (in *DestinationParameters) DeepCopy() *DestinationParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcrReplicationConfiguration) DeepCopyInto(out *EcrReplicationConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcrReplicationConfiguration.
func (in *EcrReplicationConfiguration) DeepCopy() *EcrReplicationConfiguration {
	if in == nil {
		return nil
	}
	out := new(EcrReplicationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EcrReplicationConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcrReplicationConfigurationList) DeepCopyInto(out *EcrReplicationConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EcrReplicationConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcrReplicationConfigurationList.
func (in *EcrReplicationConfigurationList) DeepCopy() *EcrReplicationConfigurationList {
	if in == nil {
		return nil
	}
	out := new(EcrReplicationConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EcrReplicationConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcrReplicationConfigurationObservation) DeepCopyInto(out *EcrReplicationConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcrReplicationConfigurationObservation.
func (in *EcrReplicationConfigurationObservation) DeepCopy() *EcrReplicationConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(EcrReplicationConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcrReplicationConfigurationParameters) DeepCopyInto(out *EcrReplicationConfigurationParameters) {
	*out = *in
	if in.ReplicationConfiguration != nil {
		in, out := &in.ReplicationConfiguration, &out.ReplicationConfiguration
		*out = make([]ReplicationConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcrReplicationConfigurationParameters.
func (in *EcrReplicationConfigurationParameters) DeepCopy() *EcrReplicationConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(EcrReplicationConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcrReplicationConfigurationSpec) DeepCopyInto(out *EcrReplicationConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcrReplicationConfigurationSpec.
func (in *EcrReplicationConfigurationSpec) DeepCopy() *EcrReplicationConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(EcrReplicationConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcrReplicationConfigurationStatus) DeepCopyInto(out *EcrReplicationConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcrReplicationConfigurationStatus.
func (in *EcrReplicationConfigurationStatus) DeepCopy() *EcrReplicationConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(EcrReplicationConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationObservation) DeepCopyInto(out *ReplicationConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationObservation.
func (in *ReplicationConfigurationObservation) DeepCopy() *ReplicationConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationParameters) DeepCopyInto(out *ReplicationConfigurationParameters) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationParameters.
func (in *ReplicationConfigurationParameters) DeepCopy() *ReplicationConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleObservation) DeepCopyInto(out *RuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleObservation.
func (in *RuleObservation) DeepCopy() *RuleObservation {
	if in == nil {
		return nil
	}
	out := new(RuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleParameters) DeepCopyInto(out *RuleParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = make([]DestinationParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleParameters.
func (in *RuleParameters) DeepCopy() *RuleParameters {
	if in == nil {
		return nil
	}
	out := new(RuleParameters)
	in.DeepCopyInto(out)
	return out
}