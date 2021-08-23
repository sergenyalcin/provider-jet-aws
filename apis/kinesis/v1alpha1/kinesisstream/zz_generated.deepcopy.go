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
func (in *KinesisStream) DeepCopyInto(out *KinesisStream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStream.
func (in *KinesisStream) DeepCopy() *KinesisStream {
	if in == nil {
		return nil
	}
	out := new(KinesisStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KinesisStream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamList) DeepCopyInto(out *KinesisStreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KinesisStream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamList.
func (in *KinesisStreamList) DeepCopy() *KinesisStreamList {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KinesisStreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamObservation) DeepCopyInto(out *KinesisStreamObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamObservation.
func (in *KinesisStreamObservation) DeepCopy() *KinesisStreamObservation {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamParameters) DeepCopyInto(out *KinesisStreamParameters) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.EnforceConsumerDeletion != nil {
		in, out := &in.EnforceConsumerDeletion, &out.EnforceConsumerDeletion
		*out = new(bool)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.RetentionPeriod != nil {
		in, out := &in.RetentionPeriod, &out.RetentionPeriod
		*out = new(int64)
		**out = **in
	}
	if in.ShardLevelMetrics != nil {
		in, out := &in.ShardLevelMetrics, &out.ShardLevelMetrics
		*out = make([]string, len(*in))
		copy(*out, *in)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamParameters.
func (in *KinesisStreamParameters) DeepCopy() *KinesisStreamParameters {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamSpec) DeepCopyInto(out *KinesisStreamSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamSpec.
func (in *KinesisStreamSpec) DeepCopy() *KinesisStreamSpec {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KinesisStreamStatus) DeepCopyInto(out *KinesisStreamStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KinesisStreamStatus.
func (in *KinesisStreamStatus) DeepCopy() *KinesisStreamStatus {
	if in == nil {
		return nil
	}
	out := new(KinesisStreamStatus)
	in.DeepCopyInto(out)
	return out
}