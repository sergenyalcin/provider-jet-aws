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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ElastictranscoderPipeline.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ElastictranscoderPipeline) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ElastictranscoderPipeline.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ElastictranscoderPipeline) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ElastictranscoderPipeline.
func (mg *ElastictranscoderPipeline) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ElastictranscoderPreset.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ElastictranscoderPreset) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ElastictranscoderPreset.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ElastictranscoderPreset) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ElastictranscoderPreset.
func (mg *ElastictranscoderPreset) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
