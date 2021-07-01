// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Accelerator) DeepCopyInto(out *Accelerator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Accelerator.
func (in *Accelerator) DeepCopy() *Accelerator {
	if in == nil {
		return nil
	}
	out := new(Accelerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Accelerator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorList) DeepCopyInto(out *AcceleratorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Accelerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorList.
func (in *AcceleratorList) DeepCopy() *AcceleratorList {
	if in == nil {
		return nil
	}
	out := new(AcceleratorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcceleratorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorSpec) DeepCopyInto(out *AcceleratorSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(AcceleratorSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorSpec.
func (in *AcceleratorSpec) DeepCopy() *AcceleratorSpec {
	if in == nil {
		return nil
	}
	out := new(AcceleratorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorSpecAttributes) DeepCopyInto(out *AcceleratorSpecAttributes) {
	*out = *in
	if in.FlowLogsEnabled != nil {
		in, out := &in.FlowLogsEnabled, &out.FlowLogsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FlowLogsS3Bucket != nil {
		in, out := &in.FlowLogsS3Bucket, &out.FlowLogsS3Bucket
		*out = new(string)
		**out = **in
	}
	if in.FlowLogsS3Prefix != nil {
		in, out := &in.FlowLogsS3Prefix, &out.FlowLogsS3Prefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorSpecAttributes.
func (in *AcceleratorSpecAttributes) DeepCopy() *AcceleratorSpecAttributes {
	if in == nil {
		return nil
	}
	out := new(AcceleratorSpecAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorSpecIpSets) DeepCopyInto(out *AcceleratorSpecIpSets) {
	*out = *in
	if in.IpAddresses != nil {
		in, out := &in.IpAddresses, &out.IpAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IpFamily != nil {
		in, out := &in.IpFamily, &out.IpFamily
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorSpecIpSets.
func (in *AcceleratorSpecIpSets) DeepCopy() *AcceleratorSpecIpSets {
	if in == nil {
		return nil
	}
	out := new(AcceleratorSpecIpSets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorSpecResource) DeepCopyInto(out *AcceleratorSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = new(AcceleratorSpecAttributes)
		(*in).DeepCopyInto(*out)
	}
	if in.DnsName != nil {
		in, out := &in.DnsName, &out.DnsName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.HostedZoneID != nil {
		in, out := &in.HostedZoneID, &out.HostedZoneID
		*out = new(string)
		**out = **in
	}
	if in.IpAddressType != nil {
		in, out := &in.IpAddressType, &out.IpAddressType
		*out = new(string)
		**out = **in
	}
	if in.IpSets != nil {
		in, out := &in.IpSets, &out.IpSets
		*out = make([]AcceleratorSpecIpSets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorSpecResource.
func (in *AcceleratorSpecResource) DeepCopy() *AcceleratorSpecResource {
	if in == nil {
		return nil
	}
	out := new(AcceleratorSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorStatus) DeepCopyInto(out *AcceleratorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorStatus.
func (in *AcceleratorStatus) DeepCopy() *AcceleratorStatus {
	if in == nil {
		return nil
	}
	out := new(AcceleratorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroup) DeepCopyInto(out *EndpointGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroup.
func (in *EndpointGroup) DeepCopy() *EndpointGroup {
	if in == nil {
		return nil
	}
	out := new(EndpointGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupList) DeepCopyInto(out *EndpointGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EndpointGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupList.
func (in *EndpointGroupList) DeepCopy() *EndpointGroupList {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupSpec) DeepCopyInto(out *EndpointGroupSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(EndpointGroupSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupSpec.
func (in *EndpointGroupSpec) DeepCopy() *EndpointGroupSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupSpecEndpointConfiguration) DeepCopyInto(out *EndpointGroupSpecEndpointConfiguration) {
	*out = *in
	if in.ClientIPPreservationEnabled != nil {
		in, out := &in.ClientIPPreservationEnabled, &out.ClientIPPreservationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupSpecEndpointConfiguration.
func (in *EndpointGroupSpecEndpointConfiguration) DeepCopy() *EndpointGroupSpecEndpointConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupSpecEndpointConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupSpecPortOverride) DeepCopyInto(out *EndpointGroupSpecPortOverride) {
	*out = *in
	if in.EndpointPort != nil {
		in, out := &in.EndpointPort, &out.EndpointPort
		*out = new(int64)
		**out = **in
	}
	if in.ListenerPort != nil {
		in, out := &in.ListenerPort, &out.ListenerPort
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupSpecPortOverride.
func (in *EndpointGroupSpecPortOverride) DeepCopy() *EndpointGroupSpecPortOverride {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupSpecPortOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupSpecResource) DeepCopyInto(out *EndpointGroupSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.EndpointConfiguration != nil {
		in, out := &in.EndpointConfiguration, &out.EndpointConfiguration
		*out = make([]EndpointGroupSpecEndpointConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndpointGroupRegion != nil {
		in, out := &in.EndpointGroupRegion, &out.EndpointGroupRegion
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckIntervalSeconds != nil {
		in, out := &in.HealthCheckIntervalSeconds, &out.HealthCheckIntervalSeconds
		*out = new(int64)
		**out = **in
	}
	if in.HealthCheckPath != nil {
		in, out := &in.HealthCheckPath, &out.HealthCheckPath
		*out = new(string)
		**out = **in
	}
	if in.HealthCheckPort != nil {
		in, out := &in.HealthCheckPort, &out.HealthCheckPort
		*out = new(int64)
		**out = **in
	}
	if in.HealthCheckProtocol != nil {
		in, out := &in.HealthCheckProtocol, &out.HealthCheckProtocol
		*out = new(string)
		**out = **in
	}
	if in.ListenerArn != nil {
		in, out := &in.ListenerArn, &out.ListenerArn
		*out = new(string)
		**out = **in
	}
	if in.PortOverride != nil {
		in, out := &in.PortOverride, &out.PortOverride
		*out = make([]EndpointGroupSpecPortOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ThresholdCount != nil {
		in, out := &in.ThresholdCount, &out.ThresholdCount
		*out = new(int64)
		**out = **in
	}
	if in.TrafficDialPercentage != nil {
		in, out := &in.TrafficDialPercentage, &out.TrafficDialPercentage
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupSpecResource.
func (in *EndpointGroupSpecResource) DeepCopy() *EndpointGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointGroupStatus) DeepCopyInto(out *EndpointGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointGroupStatus.
func (in *EndpointGroupStatus) DeepCopy() *EndpointGroupStatus {
	if in == nil {
		return nil
	}
	out := new(EndpointGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listener) DeepCopyInto(out *Listener) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Listener) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerList) DeepCopyInto(out *ListenerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Listener, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerList.
func (in *ListenerList) DeepCopy() *ListenerList {
	if in == nil {
		return nil
	}
	out := new(ListenerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListenerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerSpec) DeepCopyInto(out *ListenerSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ListenerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerSpec.
func (in *ListenerSpec) DeepCopy() *ListenerSpec {
	if in == nil {
		return nil
	}
	out := new(ListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerSpecPortRange) DeepCopyInto(out *ListenerSpecPortRange) {
	*out = *in
	if in.FromPort != nil {
		in, out := &in.FromPort, &out.FromPort
		*out = new(int64)
		**out = **in
	}
	if in.ToPort != nil {
		in, out := &in.ToPort, &out.ToPort
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerSpecPortRange.
func (in *ListenerSpecPortRange) DeepCopy() *ListenerSpecPortRange {
	if in == nil {
		return nil
	}
	out := new(ListenerSpecPortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerSpecResource) DeepCopyInto(out *ListenerSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AcceleratorArn != nil {
		in, out := &in.AcceleratorArn, &out.AcceleratorArn
		*out = new(string)
		**out = **in
	}
	if in.ClientAffinity != nil {
		in, out := &in.ClientAffinity, &out.ClientAffinity
		*out = new(string)
		**out = **in
	}
	if in.PortRange != nil {
		in, out := &in.PortRange, &out.PortRange
		*out = make([]ListenerSpecPortRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerSpecResource.
func (in *ListenerSpecResource) DeepCopy() *ListenerSpecResource {
	if in == nil {
		return nil
	}
	out := new(ListenerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerStatus) DeepCopyInto(out *ListenerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerStatus.
func (in *ListenerStatus) DeepCopy() *ListenerStatus {
	if in == nil {
		return nil
	}
	out := new(ListenerStatus)
	in.DeepCopyInto(out)
	return out
}