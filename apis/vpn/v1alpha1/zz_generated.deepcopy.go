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
func (in *RoutesObservation) DeepCopyInto(out *RoutesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutesObservation.
func (in *RoutesObservation) DeepCopy() *RoutesObservation {
	if in == nil {
		return nil
	}
	out := new(RoutesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutesParameters) DeepCopyInto(out *RoutesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutesParameters.
func (in *RoutesParameters) DeepCopy() *RoutesParameters {
	if in == nil {
		return nil
	}
	out := new(RoutesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VgwTelemetryObservation) DeepCopyInto(out *VgwTelemetryObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VgwTelemetryObservation.
func (in *VgwTelemetryObservation) DeepCopy() *VgwTelemetryObservation {
	if in == nil {
		return nil
	}
	out := new(VgwTelemetryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VgwTelemetryParameters) DeepCopyInto(out *VgwTelemetryParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VgwTelemetryParameters.
func (in *VgwTelemetryParameters) DeepCopy() *VgwTelemetryParameters {
	if in == nil {
		return nil
	}
	out := new(VgwTelemetryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnection) DeepCopyInto(out *VpnConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnection.
func (in *VpnConnection) DeepCopy() *VpnConnection {
	if in == nil {
		return nil
	}
	out := new(VpnConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionList) DeepCopyInto(out *VpnConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionList.
func (in *VpnConnectionList) DeepCopy() *VpnConnectionList {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionObservation) DeepCopyInto(out *VpnConnectionObservation) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RoutesObservation, len(*in))
		copy(*out, *in)
	}
	if in.VgwTelemetry != nil {
		in, out := &in.VgwTelemetry, &out.VgwTelemetry
		*out = make([]VgwTelemetryObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionObservation.
func (in *VpnConnectionObservation) DeepCopy() *VpnConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionParameters) DeepCopyInto(out *VpnConnectionParameters) {
	*out = *in
	if in.EnableAcceleration != nil {
		in, out := &in.EnableAcceleration, &out.EnableAcceleration
		*out = new(bool)
		**out = **in
	}
	if in.LocalIpv4NetworkCidr != nil {
		in, out := &in.LocalIpv4NetworkCidr, &out.LocalIpv4NetworkCidr
		*out = new(string)
		**out = **in
	}
	if in.LocalIpv6NetworkCidr != nil {
		in, out := &in.LocalIpv6NetworkCidr, &out.LocalIpv6NetworkCidr
		*out = new(string)
		**out = **in
	}
	if in.RemoteIpv4NetworkCidr != nil {
		in, out := &in.RemoteIpv4NetworkCidr, &out.RemoteIpv4NetworkCidr
		*out = new(string)
		**out = **in
	}
	if in.RemoteIpv6NetworkCidr != nil {
		in, out := &in.RemoteIpv6NetworkCidr, &out.RemoteIpv6NetworkCidr
		*out = new(string)
		**out = **in
	}
	if in.StaticRoutesOnly != nil {
		in, out := &in.StaticRoutesOnly, &out.StaticRoutesOnly
		*out = new(bool)
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
	if in.TransitGatewayId != nil {
		in, out := &in.TransitGatewayId, &out.TransitGatewayId
		*out = new(string)
		**out = **in
	}
	if in.Tunnel1DpdTimeoutAction != nil {
		in, out := &in.Tunnel1DpdTimeoutAction, &out.Tunnel1DpdTimeoutAction
		*out = new(string)
		**out = **in
	}
	if in.Tunnel1DpdTimeoutSeconds != nil {
		in, out := &in.Tunnel1DpdTimeoutSeconds, &out.Tunnel1DpdTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel1IkeVersions != nil {
		in, out := &in.Tunnel1IkeVersions, &out.Tunnel1IkeVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel1InsideCidr != nil {
		in, out := &in.Tunnel1InsideCidr, &out.Tunnel1InsideCidr
		*out = new(string)
		**out = **in
	}
	if in.Tunnel1InsideIpv6Cidr != nil {
		in, out := &in.Tunnel1InsideIpv6Cidr, &out.Tunnel1InsideIpv6Cidr
		*out = new(string)
		**out = **in
	}
	if in.Tunnel1Phase1DhGroupNumbers != nil {
		in, out := &in.Tunnel1Phase1DhGroupNumbers, &out.Tunnel1Phase1DhGroupNumbers
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel1Phase1EncryptionAlgorithms != nil {
		in, out := &in.Tunnel1Phase1EncryptionAlgorithms, &out.Tunnel1Phase1EncryptionAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel1Phase1IntegrityAlgorithms != nil {
		in, out := &in.Tunnel1Phase1IntegrityAlgorithms, &out.Tunnel1Phase1IntegrityAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel1Phase1LifetimeSeconds != nil {
		in, out := &in.Tunnel1Phase1LifetimeSeconds, &out.Tunnel1Phase1LifetimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel1Phase2DhGroupNumbers != nil {
		in, out := &in.Tunnel1Phase2DhGroupNumbers, &out.Tunnel1Phase2DhGroupNumbers
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel1Phase2EncryptionAlgorithms != nil {
		in, out := &in.Tunnel1Phase2EncryptionAlgorithms, &out.Tunnel1Phase2EncryptionAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel1Phase2IntegrityAlgorithms != nil {
		in, out := &in.Tunnel1Phase2IntegrityAlgorithms, &out.Tunnel1Phase2IntegrityAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel1Phase2LifetimeSeconds != nil {
		in, out := &in.Tunnel1Phase2LifetimeSeconds, &out.Tunnel1Phase2LifetimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel1PresharedKey != nil {
		in, out := &in.Tunnel1PresharedKey, &out.Tunnel1PresharedKey
		*out = new(string)
		**out = **in
	}
	if in.Tunnel1RekeyFuzzPercentage != nil {
		in, out := &in.Tunnel1RekeyFuzzPercentage, &out.Tunnel1RekeyFuzzPercentage
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel1RekeyMarginTimeSeconds != nil {
		in, out := &in.Tunnel1RekeyMarginTimeSeconds, &out.Tunnel1RekeyMarginTimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel1ReplayWindowSize != nil {
		in, out := &in.Tunnel1ReplayWindowSize, &out.Tunnel1ReplayWindowSize
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel1StartupAction != nil {
		in, out := &in.Tunnel1StartupAction, &out.Tunnel1StartupAction
		*out = new(string)
		**out = **in
	}
	if in.Tunnel2DpdTimeoutAction != nil {
		in, out := &in.Tunnel2DpdTimeoutAction, &out.Tunnel2DpdTimeoutAction
		*out = new(string)
		**out = **in
	}
	if in.Tunnel2DpdTimeoutSeconds != nil {
		in, out := &in.Tunnel2DpdTimeoutSeconds, &out.Tunnel2DpdTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel2IkeVersions != nil {
		in, out := &in.Tunnel2IkeVersions, &out.Tunnel2IkeVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel2InsideCidr != nil {
		in, out := &in.Tunnel2InsideCidr, &out.Tunnel2InsideCidr
		*out = new(string)
		**out = **in
	}
	if in.Tunnel2InsideIpv6Cidr != nil {
		in, out := &in.Tunnel2InsideIpv6Cidr, &out.Tunnel2InsideIpv6Cidr
		*out = new(string)
		**out = **in
	}
	if in.Tunnel2Phase1DhGroupNumbers != nil {
		in, out := &in.Tunnel2Phase1DhGroupNumbers, &out.Tunnel2Phase1DhGroupNumbers
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel2Phase1EncryptionAlgorithms != nil {
		in, out := &in.Tunnel2Phase1EncryptionAlgorithms, &out.Tunnel2Phase1EncryptionAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel2Phase1IntegrityAlgorithms != nil {
		in, out := &in.Tunnel2Phase1IntegrityAlgorithms, &out.Tunnel2Phase1IntegrityAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel2Phase1LifetimeSeconds != nil {
		in, out := &in.Tunnel2Phase1LifetimeSeconds, &out.Tunnel2Phase1LifetimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel2Phase2DhGroupNumbers != nil {
		in, out := &in.Tunnel2Phase2DhGroupNumbers, &out.Tunnel2Phase2DhGroupNumbers
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel2Phase2EncryptionAlgorithms != nil {
		in, out := &in.Tunnel2Phase2EncryptionAlgorithms, &out.Tunnel2Phase2EncryptionAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel2Phase2IntegrityAlgorithms != nil {
		in, out := &in.Tunnel2Phase2IntegrityAlgorithms, &out.Tunnel2Phase2IntegrityAlgorithms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tunnel2Phase2LifetimeSeconds != nil {
		in, out := &in.Tunnel2Phase2LifetimeSeconds, &out.Tunnel2Phase2LifetimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel2PresharedKey != nil {
		in, out := &in.Tunnel2PresharedKey, &out.Tunnel2PresharedKey
		*out = new(string)
		**out = **in
	}
	if in.Tunnel2RekeyFuzzPercentage != nil {
		in, out := &in.Tunnel2RekeyFuzzPercentage, &out.Tunnel2RekeyFuzzPercentage
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel2RekeyMarginTimeSeconds != nil {
		in, out := &in.Tunnel2RekeyMarginTimeSeconds, &out.Tunnel2RekeyMarginTimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel2ReplayWindowSize != nil {
		in, out := &in.Tunnel2ReplayWindowSize, &out.Tunnel2ReplayWindowSize
		*out = new(int64)
		**out = **in
	}
	if in.Tunnel2StartupAction != nil {
		in, out := &in.Tunnel2StartupAction, &out.Tunnel2StartupAction
		*out = new(string)
		**out = **in
	}
	if in.TunnelInsideIpVersion != nil {
		in, out := &in.TunnelInsideIpVersion, &out.TunnelInsideIpVersion
		*out = new(string)
		**out = **in
	}
	if in.VpnGatewayId != nil {
		in, out := &in.VpnGatewayId, &out.VpnGatewayId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionParameters.
func (in *VpnConnectionParameters) DeepCopy() *VpnConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionRoute) DeepCopyInto(out *VpnConnectionRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionRoute.
func (in *VpnConnectionRoute) DeepCopy() *VpnConnectionRoute {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnConnectionRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionRouteList) DeepCopyInto(out *VpnConnectionRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnConnectionRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionRouteList.
func (in *VpnConnectionRouteList) DeepCopy() *VpnConnectionRouteList {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnConnectionRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionRouteObservation) DeepCopyInto(out *VpnConnectionRouteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionRouteObservation.
func (in *VpnConnectionRouteObservation) DeepCopy() *VpnConnectionRouteObservation {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionRouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionRouteParameters) DeepCopyInto(out *VpnConnectionRouteParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionRouteParameters.
func (in *VpnConnectionRouteParameters) DeepCopy() *VpnConnectionRouteParameters {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionRouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionRouteSpec) DeepCopyInto(out *VpnConnectionRouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionRouteSpec.
func (in *VpnConnectionRouteSpec) DeepCopy() *VpnConnectionRouteSpec {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionRouteStatus) DeepCopyInto(out *VpnConnectionRouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionRouteStatus.
func (in *VpnConnectionRouteStatus) DeepCopy() *VpnConnectionRouteStatus {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionSpec) DeepCopyInto(out *VpnConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionSpec.
func (in *VpnConnectionSpec) DeepCopy() *VpnConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnConnectionStatus) DeepCopyInto(out *VpnConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnConnectionStatus.
func (in *VpnConnectionStatus) DeepCopy() *VpnConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(VpnConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGateway) DeepCopyInto(out *VpnGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGateway.
func (in *VpnGateway) DeepCopy() *VpnGateway {
	if in == nil {
		return nil
	}
	out := new(VpnGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayAttachment) DeepCopyInto(out *VpnGatewayAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayAttachment.
func (in *VpnGatewayAttachment) DeepCopy() *VpnGatewayAttachment {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGatewayAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayAttachmentList) DeepCopyInto(out *VpnGatewayAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnGatewayAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayAttachmentList.
func (in *VpnGatewayAttachmentList) DeepCopy() *VpnGatewayAttachmentList {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGatewayAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayAttachmentObservation) DeepCopyInto(out *VpnGatewayAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayAttachmentObservation.
func (in *VpnGatewayAttachmentObservation) DeepCopy() *VpnGatewayAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayAttachmentParameters) DeepCopyInto(out *VpnGatewayAttachmentParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayAttachmentParameters.
func (in *VpnGatewayAttachmentParameters) DeepCopy() *VpnGatewayAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayAttachmentSpec) DeepCopyInto(out *VpnGatewayAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayAttachmentSpec.
func (in *VpnGatewayAttachmentSpec) DeepCopy() *VpnGatewayAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayAttachmentStatus) DeepCopyInto(out *VpnGatewayAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayAttachmentStatus.
func (in *VpnGatewayAttachmentStatus) DeepCopy() *VpnGatewayAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayList) DeepCopyInto(out *VpnGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayList.
func (in *VpnGatewayList) DeepCopy() *VpnGatewayList {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayObservation) DeepCopyInto(out *VpnGatewayObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayObservation.
func (in *VpnGatewayObservation) DeepCopy() *VpnGatewayObservation {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayParameters) DeepCopyInto(out *VpnGatewayParameters) {
	*out = *in
	if in.AmazonSideAsn != nil {
		in, out := &in.AmazonSideAsn, &out.AmazonSideAsn
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
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
	if in.VpcId != nil {
		in, out := &in.VpcId, &out.VpcId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayParameters.
func (in *VpnGatewayParameters) DeepCopy() *VpnGatewayParameters {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagation) DeepCopyInto(out *VpnGatewayRoutePropagation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagation.
func (in *VpnGatewayRoutePropagation) DeepCopy() *VpnGatewayRoutePropagation {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGatewayRoutePropagation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationList) DeepCopyInto(out *VpnGatewayRoutePropagationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnGatewayRoutePropagation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationList.
func (in *VpnGatewayRoutePropagationList) DeepCopy() *VpnGatewayRoutePropagationList {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGatewayRoutePropagationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationObservation) DeepCopyInto(out *VpnGatewayRoutePropagationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationObservation.
func (in *VpnGatewayRoutePropagationObservation) DeepCopy() *VpnGatewayRoutePropagationObservation {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationParameters) DeepCopyInto(out *VpnGatewayRoutePropagationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationParameters.
func (in *VpnGatewayRoutePropagationParameters) DeepCopy() *VpnGatewayRoutePropagationParameters {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationSpec) DeepCopyInto(out *VpnGatewayRoutePropagationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationSpec.
func (in *VpnGatewayRoutePropagationSpec) DeepCopy() *VpnGatewayRoutePropagationSpec {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationStatus) DeepCopyInto(out *VpnGatewayRoutePropagationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationStatus.
func (in *VpnGatewayRoutePropagationStatus) DeepCopy() *VpnGatewayRoutePropagationStatus {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewaySpec) DeepCopyInto(out *VpnGatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewaySpec.
func (in *VpnGatewaySpec) DeepCopy() *VpnGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(VpnGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayStatus) DeepCopyInto(out *VpnGatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayStatus.
func (in *VpnGatewayStatus) DeepCopy() *VpnGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayStatus)
	in.DeepCopyInto(out)
	return out
}
