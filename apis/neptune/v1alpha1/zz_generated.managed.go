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

// GetCondition of this NeptuneCluster.
func (mg *NeptuneCluster) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneCluster.
func (mg *NeptuneCluster) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneCluster.
func (mg *NeptuneCluster) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneCluster.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneCluster) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneCluster.
func (mg *NeptuneCluster) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneCluster.
func (mg *NeptuneCluster) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneCluster.
func (mg *NeptuneCluster) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneCluster.
func (mg *NeptuneCluster) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneCluster.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneCluster) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneCluster.
func (mg *NeptuneCluster) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneClusterEndpoint.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneClusterEndpoint) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneClusterEndpoint.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneClusterEndpoint) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneClusterEndpoint.
func (mg *NeptuneClusterEndpoint) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneClusterInstance.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneClusterInstance) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneClusterInstance.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneClusterInstance) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneClusterInstance.
func (mg *NeptuneClusterInstance) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneClusterParameterGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneClusterParameterGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneClusterParameterGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneClusterParameterGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneClusterParameterGroup.
func (mg *NeptuneClusterParameterGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneClusterSnapshot.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneClusterSnapshot) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneClusterSnapshot.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneClusterSnapshot) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneClusterSnapshot.
func (mg *NeptuneClusterSnapshot) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneEventSubscription.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneEventSubscription) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneEventSubscription.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneEventSubscription) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneEventSubscription.
func (mg *NeptuneEventSubscription) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneParameterGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneParameterGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneParameterGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneParameterGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneParameterGroup.
func (mg *NeptuneParameterGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NeptuneSubnetGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NeptuneSubnetGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NeptuneSubnetGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NeptuneSubnetGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NeptuneSubnetGroup.
func (mg *NeptuneSubnetGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
