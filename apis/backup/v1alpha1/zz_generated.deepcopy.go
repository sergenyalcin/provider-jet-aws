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
func (in *AdvancedBackupSettingObservation) DeepCopyInto(out *AdvancedBackupSettingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedBackupSettingObservation.
func (in *AdvancedBackupSettingObservation) DeepCopy() *AdvancedBackupSettingObservation {
	if in == nil {
		return nil
	}
	out := new(AdvancedBackupSettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedBackupSettingParameters) DeepCopyInto(out *AdvancedBackupSettingParameters) {
	*out = *in
	if in.BackupOptions != nil {
		in, out := &in.BackupOptions, &out.BackupOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedBackupSettingParameters.
func (in *AdvancedBackupSettingParameters) DeepCopy() *AdvancedBackupSettingParameters {
	if in == nil {
		return nil
	}
	out := new(AdvancedBackupSettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGlobalSettings) DeepCopyInto(out *BackupGlobalSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGlobalSettings.
func (in *BackupGlobalSettings) DeepCopy() *BackupGlobalSettings {
	if in == nil {
		return nil
	}
	out := new(BackupGlobalSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupGlobalSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGlobalSettingsList) DeepCopyInto(out *BackupGlobalSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupGlobalSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGlobalSettingsList.
func (in *BackupGlobalSettingsList) DeepCopy() *BackupGlobalSettingsList {
	if in == nil {
		return nil
	}
	out := new(BackupGlobalSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupGlobalSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGlobalSettingsObservation) DeepCopyInto(out *BackupGlobalSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGlobalSettingsObservation.
func (in *BackupGlobalSettingsObservation) DeepCopy() *BackupGlobalSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(BackupGlobalSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGlobalSettingsParameters) DeepCopyInto(out *BackupGlobalSettingsParameters) {
	*out = *in
	if in.GlobalSettings != nil {
		in, out := &in.GlobalSettings, &out.GlobalSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGlobalSettingsParameters.
func (in *BackupGlobalSettingsParameters) DeepCopy() *BackupGlobalSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(BackupGlobalSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGlobalSettingsSpec) DeepCopyInto(out *BackupGlobalSettingsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGlobalSettingsSpec.
func (in *BackupGlobalSettingsSpec) DeepCopy() *BackupGlobalSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(BackupGlobalSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGlobalSettingsStatus) DeepCopyInto(out *BackupGlobalSettingsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGlobalSettingsStatus.
func (in *BackupGlobalSettingsStatus) DeepCopy() *BackupGlobalSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(BackupGlobalSettingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlan) DeepCopyInto(out *BackupPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlan.
func (in *BackupPlan) DeepCopy() *BackupPlan {
	if in == nil {
		return nil
	}
	out := new(BackupPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanList) DeepCopyInto(out *BackupPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanList.
func (in *BackupPlanList) DeepCopy() *BackupPlanList {
	if in == nil {
		return nil
	}
	out := new(BackupPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanObservation) DeepCopyInto(out *BackupPlanObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanObservation.
func (in *BackupPlanObservation) DeepCopy() *BackupPlanObservation {
	if in == nil {
		return nil
	}
	out := new(BackupPlanObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanParameters) DeepCopyInto(out *BackupPlanParameters) {
	*out = *in
	if in.AdvancedBackupSetting != nil {
		in, out := &in.AdvancedBackupSetting, &out.AdvancedBackupSetting
		*out = make([]AdvancedBackupSettingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleParameters, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanParameters.
func (in *BackupPlanParameters) DeepCopy() *BackupPlanParameters {
	if in == nil {
		return nil
	}
	out := new(BackupPlanParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanSpec) DeepCopyInto(out *BackupPlanSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanSpec.
func (in *BackupPlanSpec) DeepCopy() *BackupPlanSpec {
	if in == nil {
		return nil
	}
	out := new(BackupPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPlanStatus) DeepCopyInto(out *BackupPlanStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPlanStatus.
func (in *BackupPlanStatus) DeepCopy() *BackupPlanStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRegionSettings) DeepCopyInto(out *BackupRegionSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRegionSettings.
func (in *BackupRegionSettings) DeepCopy() *BackupRegionSettings {
	if in == nil {
		return nil
	}
	out := new(BackupRegionSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupRegionSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRegionSettingsList) DeepCopyInto(out *BackupRegionSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupRegionSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRegionSettingsList.
func (in *BackupRegionSettingsList) DeepCopy() *BackupRegionSettingsList {
	if in == nil {
		return nil
	}
	out := new(BackupRegionSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupRegionSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRegionSettingsObservation) DeepCopyInto(out *BackupRegionSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRegionSettingsObservation.
func (in *BackupRegionSettingsObservation) DeepCopy() *BackupRegionSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(BackupRegionSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRegionSettingsParameters) DeepCopyInto(out *BackupRegionSettingsParameters) {
	*out = *in
	if in.ResourceTypeOptInPreference != nil {
		in, out := &in.ResourceTypeOptInPreference, &out.ResourceTypeOptInPreference
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRegionSettingsParameters.
func (in *BackupRegionSettingsParameters) DeepCopy() *BackupRegionSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(BackupRegionSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRegionSettingsSpec) DeepCopyInto(out *BackupRegionSettingsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRegionSettingsSpec.
func (in *BackupRegionSettingsSpec) DeepCopy() *BackupRegionSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(BackupRegionSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRegionSettingsStatus) DeepCopyInto(out *BackupRegionSettingsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRegionSettingsStatus.
func (in *BackupRegionSettingsStatus) DeepCopy() *BackupRegionSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(BackupRegionSettingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSelection) DeepCopyInto(out *BackupSelection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSelection.
func (in *BackupSelection) DeepCopy() *BackupSelection {
	if in == nil {
		return nil
	}
	out := new(BackupSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupSelection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSelectionList) DeepCopyInto(out *BackupSelectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSelectionList.
func (in *BackupSelectionList) DeepCopy() *BackupSelectionList {
	if in == nil {
		return nil
	}
	out := new(BackupSelectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupSelectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSelectionObservation) DeepCopyInto(out *BackupSelectionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSelectionObservation.
func (in *BackupSelectionObservation) DeepCopy() *BackupSelectionObservation {
	if in == nil {
		return nil
	}
	out := new(BackupSelectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSelectionParameters) DeepCopyInto(out *BackupSelectionParameters) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SelectionTag != nil {
		in, out := &in.SelectionTag, &out.SelectionTag
		*out = make([]SelectionTagParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSelectionParameters.
func (in *BackupSelectionParameters) DeepCopy() *BackupSelectionParameters {
	if in == nil {
		return nil
	}
	out := new(BackupSelectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSelectionSpec) DeepCopyInto(out *BackupSelectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSelectionSpec.
func (in *BackupSelectionSpec) DeepCopy() *BackupSelectionSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSelectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSelectionStatus) DeepCopyInto(out *BackupSelectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSelectionStatus.
func (in *BackupSelectionStatus) DeepCopy() *BackupSelectionStatus {
	if in == nil {
		return nil
	}
	out := new(BackupSelectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVault) DeepCopyInto(out *BackupVault) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVault.
func (in *BackupVault) DeepCopy() *BackupVault {
	if in == nil {
		return nil
	}
	out := new(BackupVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupVault) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultList) DeepCopyInto(out *BackupVaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupVault, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultList.
func (in *BackupVaultList) DeepCopy() *BackupVaultList {
	if in == nil {
		return nil
	}
	out := new(BackupVaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupVaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultNotifications) DeepCopyInto(out *BackupVaultNotifications) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultNotifications.
func (in *BackupVaultNotifications) DeepCopy() *BackupVaultNotifications {
	if in == nil {
		return nil
	}
	out := new(BackupVaultNotifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupVaultNotifications) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultNotificationsList) DeepCopyInto(out *BackupVaultNotificationsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupVaultNotifications, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultNotificationsList.
func (in *BackupVaultNotificationsList) DeepCopy() *BackupVaultNotificationsList {
	if in == nil {
		return nil
	}
	out := new(BackupVaultNotificationsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupVaultNotificationsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultNotificationsObservation) DeepCopyInto(out *BackupVaultNotificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultNotificationsObservation.
func (in *BackupVaultNotificationsObservation) DeepCopy() *BackupVaultNotificationsObservation {
	if in == nil {
		return nil
	}
	out := new(BackupVaultNotificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultNotificationsParameters) DeepCopyInto(out *BackupVaultNotificationsParameters) {
	*out = *in
	if in.BackupVaultEvents != nil {
		in, out := &in.BackupVaultEvents, &out.BackupVaultEvents
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultNotificationsParameters.
func (in *BackupVaultNotificationsParameters) DeepCopy() *BackupVaultNotificationsParameters {
	if in == nil {
		return nil
	}
	out := new(BackupVaultNotificationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultNotificationsSpec) DeepCopyInto(out *BackupVaultNotificationsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultNotificationsSpec.
func (in *BackupVaultNotificationsSpec) DeepCopy() *BackupVaultNotificationsSpec {
	if in == nil {
		return nil
	}
	out := new(BackupVaultNotificationsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultNotificationsStatus) DeepCopyInto(out *BackupVaultNotificationsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultNotificationsStatus.
func (in *BackupVaultNotificationsStatus) DeepCopy() *BackupVaultNotificationsStatus {
	if in == nil {
		return nil
	}
	out := new(BackupVaultNotificationsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultObservation) DeepCopyInto(out *BackupVaultObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultObservation.
func (in *BackupVaultObservation) DeepCopy() *BackupVaultObservation {
	if in == nil {
		return nil
	}
	out := new(BackupVaultObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultParameters) DeepCopyInto(out *BackupVaultParameters) {
	*out = *in
	if in.KmsKeyArn != nil {
		in, out := &in.KmsKeyArn, &out.KmsKeyArn
		*out = new(string)
		**out = **in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultParameters.
func (in *BackupVaultParameters) DeepCopy() *BackupVaultParameters {
	if in == nil {
		return nil
	}
	out := new(BackupVaultParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultPolicy) DeepCopyInto(out *BackupVaultPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultPolicy.
func (in *BackupVaultPolicy) DeepCopy() *BackupVaultPolicy {
	if in == nil {
		return nil
	}
	out := new(BackupVaultPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupVaultPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultPolicyList) DeepCopyInto(out *BackupVaultPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupVaultPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultPolicyList.
func (in *BackupVaultPolicyList) DeepCopy() *BackupVaultPolicyList {
	if in == nil {
		return nil
	}
	out := new(BackupVaultPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupVaultPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultPolicyObservation) DeepCopyInto(out *BackupVaultPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultPolicyObservation.
func (in *BackupVaultPolicyObservation) DeepCopy() *BackupVaultPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(BackupVaultPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultPolicyParameters) DeepCopyInto(out *BackupVaultPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultPolicyParameters.
func (in *BackupVaultPolicyParameters) DeepCopy() *BackupVaultPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(BackupVaultPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultPolicySpec) DeepCopyInto(out *BackupVaultPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultPolicySpec.
func (in *BackupVaultPolicySpec) DeepCopy() *BackupVaultPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BackupVaultPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultPolicyStatus) DeepCopyInto(out *BackupVaultPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultPolicyStatus.
func (in *BackupVaultPolicyStatus) DeepCopy() *BackupVaultPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BackupVaultPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultSpec) DeepCopyInto(out *BackupVaultSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultSpec.
func (in *BackupVaultSpec) DeepCopy() *BackupVaultSpec {
	if in == nil {
		return nil
	}
	out := new(BackupVaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVaultStatus) DeepCopyInto(out *BackupVaultStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVaultStatus.
func (in *BackupVaultStatus) DeepCopy() *BackupVaultStatus {
	if in == nil {
		return nil
	}
	out := new(BackupVaultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopyActionObservation) DeepCopyInto(out *CopyActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyActionObservation.
func (in *CopyActionObservation) DeepCopy() *CopyActionObservation {
	if in == nil {
		return nil
	}
	out := new(CopyActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopyActionParameters) DeepCopyInto(out *CopyActionParameters) {
	*out = *in
	if in.Lifecycle != nil {
		in, out := &in.Lifecycle, &out.Lifecycle
		*out = make([]LifecycleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyActionParameters.
func (in *CopyActionParameters) DeepCopy() *CopyActionParameters {
	if in == nil {
		return nil
	}
	out := new(CopyActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleObservation) DeepCopyInto(out *LifecycleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleObservation.
func (in *LifecycleObservation) DeepCopy() *LifecycleObservation {
	if in == nil {
		return nil
	}
	out := new(LifecycleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleParameters) DeepCopyInto(out *LifecycleParameters) {
	*out = *in
	if in.ColdStorageAfter != nil {
		in, out := &in.ColdStorageAfter, &out.ColdStorageAfter
		*out = new(int64)
		**out = **in
	}
	if in.DeleteAfter != nil {
		in, out := &in.DeleteAfter, &out.DeleteAfter
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleParameters.
func (in *LifecycleParameters) DeepCopy() *LifecycleParameters {
	if in == nil {
		return nil
	}
	out := new(LifecycleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleLifecycleObservation) DeepCopyInto(out *RuleLifecycleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleLifecycleObservation.
func (in *RuleLifecycleObservation) DeepCopy() *RuleLifecycleObservation {
	if in == nil {
		return nil
	}
	out := new(RuleLifecycleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleLifecycleParameters) DeepCopyInto(out *RuleLifecycleParameters) {
	*out = *in
	if in.ColdStorageAfter != nil {
		in, out := &in.ColdStorageAfter, &out.ColdStorageAfter
		*out = new(int64)
		**out = **in
	}
	if in.DeleteAfter != nil {
		in, out := &in.DeleteAfter, &out.DeleteAfter
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleLifecycleParameters.
func (in *RuleLifecycleParameters) DeepCopy() *RuleLifecycleParameters {
	if in == nil {
		return nil
	}
	out := new(RuleLifecycleParameters)
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
	if in.CompletionWindow != nil {
		in, out := &in.CompletionWindow, &out.CompletionWindow
		*out = new(int64)
		**out = **in
	}
	if in.CopyAction != nil {
		in, out := &in.CopyAction, &out.CopyAction
		*out = make([]CopyActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableContinuousBackup != nil {
		in, out := &in.EnableContinuousBackup, &out.EnableContinuousBackup
		*out = new(bool)
		**out = **in
	}
	if in.Lifecycle != nil {
		in, out := &in.Lifecycle, &out.Lifecycle
		*out = make([]RuleLifecycleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecoveryPointTags != nil {
		in, out := &in.RecoveryPointTags, &out.RecoveryPointTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.StartWindow != nil {
		in, out := &in.StartWindow, &out.StartWindow
		*out = new(int64)
		**out = **in
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectionTagObservation) DeepCopyInto(out *SelectionTagObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectionTagObservation.
func (in *SelectionTagObservation) DeepCopy() *SelectionTagObservation {
	if in == nil {
		return nil
	}
	out := new(SelectionTagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectionTagParameters) DeepCopyInto(out *SelectionTagParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectionTagParameters.
func (in *SelectionTagParameters) DeepCopy() *SelectionTagParameters {
	if in == nil {
		return nil
	}
	out := new(SelectionTagParameters)
	in.DeepCopyInto(out)
	return out
}
