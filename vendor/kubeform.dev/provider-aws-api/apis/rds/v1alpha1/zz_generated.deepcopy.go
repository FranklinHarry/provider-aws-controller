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
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpoint) DeepCopyInto(out *ClusterEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpoint.
func (in *ClusterEndpoint) DeepCopy() *ClusterEndpoint {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointList) DeepCopyInto(out *ClusterEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointList.
func (in *ClusterEndpointList) DeepCopy() *ClusterEndpointList {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointSpec) DeepCopyInto(out *ClusterEndpointSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ClusterEndpointSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointSpec.
func (in *ClusterEndpointSpec) DeepCopy() *ClusterEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointSpecResource) DeepCopyInto(out *ClusterEndpointSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ClusterEndpointIdentifier != nil {
		in, out := &in.ClusterEndpointIdentifier, &out.ClusterEndpointIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ClusterIdentifier != nil {
		in, out := &in.ClusterIdentifier, &out.ClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.CustomEndpointType != nil {
		in, out := &in.CustomEndpointType, &out.CustomEndpointType
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ExcludedMembers != nil {
		in, out := &in.ExcludedMembers, &out.ExcludedMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StaticMembers != nil {
		in, out := &in.StaticMembers, &out.StaticMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointSpecResource.
func (in *ClusterEndpointSpecResource) DeepCopy() *ClusterEndpointSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointStatus) DeepCopyInto(out *ClusterEndpointStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointStatus.
func (in *ClusterEndpointStatus) DeepCopy() *ClusterEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstance) DeepCopyInto(out *ClusterInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstance.
func (in *ClusterInstance) DeepCopy() *ClusterInstance {
	if in == nil {
		return nil
	}
	out := new(ClusterInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstanceList) DeepCopyInto(out *ClusterInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstanceList.
func (in *ClusterInstanceList) DeepCopy() *ClusterInstanceList {
	if in == nil {
		return nil
	}
	out := new(ClusterInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstanceSpec) DeepCopyInto(out *ClusterInstanceSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ClusterInstanceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstanceSpec.
func (in *ClusterInstanceSpec) DeepCopy() *ClusterInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstanceSpecResource) DeepCopyInto(out *ClusterInstanceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoMinorVersionUpgrade != nil {
		in, out := &in.AutoMinorVersionUpgrade, &out.AutoMinorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CaCertIdentifier != nil {
		in, out := &in.CaCertIdentifier, &out.CaCertIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ClusterIdentifier != nil {
		in, out := &in.ClusterIdentifier, &out.ClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.CopyTagsToSnapshot != nil {
		in, out := &in.CopyTagsToSnapshot, &out.CopyTagsToSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.DbParameterGroupName != nil {
		in, out := &in.DbParameterGroupName, &out.DbParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.DbSubnetGroupName != nil {
		in, out := &in.DbSubnetGroupName, &out.DbSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.DbiResourceID != nil {
		in, out := &in.DbiResourceID, &out.DbiResourceID
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.Identifier != nil {
		in, out := &in.Identifier, &out.Identifier
		*out = new(string)
		**out = **in
	}
	if in.IdentifierPrefix != nil {
		in, out := &in.IdentifierPrefix, &out.IdentifierPrefix
		*out = new(string)
		**out = **in
	}
	if in.InstanceClass != nil {
		in, out := &in.InstanceClass, &out.InstanceClass
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.MonitoringInterval != nil {
		in, out := &in.MonitoringInterval, &out.MonitoringInterval
		*out = new(int64)
		**out = **in
	}
	if in.MonitoringRoleArn != nil {
		in, out := &in.MonitoringRoleArn, &out.MonitoringRoleArn
		*out = new(string)
		**out = **in
	}
	if in.PerformanceInsightsEnabled != nil {
		in, out := &in.PerformanceInsightsEnabled, &out.PerformanceInsightsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PerformanceInsightsKmsKeyID != nil {
		in, out := &in.PerformanceInsightsKmsKeyID, &out.PerformanceInsightsKmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.PreferredBackupWindow != nil {
		in, out := &in.PreferredBackupWindow, &out.PreferredBackupWindow
		*out = new(string)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.PromotionTier != nil {
		in, out := &in.PromotionTier, &out.PromotionTier
		*out = new(int64)
		**out = **in
	}
	if in.PubliclyAccessible != nil {
		in, out := &in.PubliclyAccessible, &out.PubliclyAccessible
		*out = new(bool)
		**out = **in
	}
	if in.StorageEncrypted != nil {
		in, out := &in.StorageEncrypted, &out.StorageEncrypted
		*out = new(bool)
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
	if in.Writer != nil {
		in, out := &in.Writer, &out.Writer
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstanceSpecResource.
func (in *ClusterInstanceSpecResource) DeepCopy() *ClusterInstanceSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterInstanceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstanceStatus) DeepCopyInto(out *ClusterInstanceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstanceStatus.
func (in *ClusterInstanceStatus) DeepCopy() *ClusterInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterGroup) DeepCopyInto(out *ClusterParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterGroup.
func (in *ClusterParameterGroup) DeepCopy() *ClusterParameterGroup {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterGroupList) DeepCopyInto(out *ClusterParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterGroupList.
func (in *ClusterParameterGroupList) DeepCopy() *ClusterParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterGroupSpec) DeepCopyInto(out *ClusterParameterGroupSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ClusterParameterGroupSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterGroupSpec.
func (in *ClusterParameterGroupSpec) DeepCopy() *ClusterParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterGroupSpecParameter) DeepCopyInto(out *ClusterParameterGroupSpecParameter) {
	*out = *in
	if in.ApplyMethod != nil {
		in, out := &in.ApplyMethod, &out.ApplyMethod
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterGroupSpecParameter.
func (in *ClusterParameterGroupSpecParameter) DeepCopy() *ClusterParameterGroupSpecParameter {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterGroupSpecParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterGroupSpecResource) DeepCopyInto(out *ClusterParameterGroupSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]ClusterParameterGroupSpecParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterGroupSpecResource.
func (in *ClusterParameterGroupSpecResource) DeepCopy() *ClusterParameterGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterGroupStatus) DeepCopyInto(out *ClusterParameterGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterGroupStatus.
func (in *ClusterParameterGroupStatus) DeepCopy() *ClusterParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ClusterSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecResource) DeepCopyInto(out *ClusterSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowMajorVersionUpgrade != nil {
		in, out := &in.AllowMajorVersionUpgrade, &out.AllowMajorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BacktrackWindow != nil {
		in, out := &in.BacktrackWindow, &out.BacktrackWindow
		*out = new(int64)
		**out = **in
	}
	if in.BackupRetentionPeriod != nil {
		in, out := &in.BackupRetentionPeriod, &out.BackupRetentionPeriod
		*out = new(int64)
		**out = **in
	}
	if in.ClusterIdentifier != nil {
		in, out := &in.ClusterIdentifier, &out.ClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ClusterIdentifierPrefix != nil {
		in, out := &in.ClusterIdentifierPrefix, &out.ClusterIdentifierPrefix
		*out = new(string)
		**out = **in
	}
	if in.ClusterMembers != nil {
		in, out := &in.ClusterMembers, &out.ClusterMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterResourceID != nil {
		in, out := &in.ClusterResourceID, &out.ClusterResourceID
		*out = new(string)
		**out = **in
	}
	if in.CopyTagsToSnapshot != nil {
		in, out := &in.CopyTagsToSnapshot, &out.CopyTagsToSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.DbClusterParameterGroupName != nil {
		in, out := &in.DbClusterParameterGroupName, &out.DbClusterParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.DbSubnetGroupName != nil {
		in, out := &in.DbSubnetGroupName, &out.DbSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.EnableHTTPEndpoint != nil {
		in, out := &in.EnableHTTPEndpoint, &out.EnableHTTPEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.EnabledCloudwatchLogsExports != nil {
		in, out := &in.EnabledCloudwatchLogsExports, &out.EnabledCloudwatchLogsExports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineMode != nil {
		in, out := &in.EngineMode, &out.EngineMode
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.FinalSnapshotIdentifier != nil {
		in, out := &in.FinalSnapshotIdentifier, &out.FinalSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.GlobalClusterIdentifier != nil {
		in, out := &in.GlobalClusterIdentifier, &out.GlobalClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.HostedZoneID != nil {
		in, out := &in.HostedZoneID, &out.HostedZoneID
		*out = new(string)
		**out = **in
	}
	if in.IamDatabaseAuthenticationEnabled != nil {
		in, out := &in.IamDatabaseAuthenticationEnabled, &out.IamDatabaseAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IamRoles != nil {
		in, out := &in.IamRoles, &out.IamRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.MasterPassword != nil {
		in, out := &in.MasterPassword, &out.MasterPassword
		*out = new(string)
		**out = **in
	}
	if in.MasterUsername != nil {
		in, out := &in.MasterUsername, &out.MasterUsername
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.PreferredBackupWindow != nil {
		in, out := &in.PreferredBackupWindow, &out.PreferredBackupWindow
		*out = new(string)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.ReaderEndpoint != nil {
		in, out := &in.ReaderEndpoint, &out.ReaderEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSourceIdentifier != nil {
		in, out := &in.ReplicationSourceIdentifier, &out.ReplicationSourceIdentifier
		*out = new(string)
		**out = **in
	}
	if in.RestoreToPointInTime != nil {
		in, out := &in.RestoreToPointInTime, &out.RestoreToPointInTime
		*out = new(ClusterSpecRestoreToPointInTime)
		(*in).DeepCopyInto(*out)
	}
	if in.S3Import != nil {
		in, out := &in.S3Import, &out.S3Import
		*out = new(ClusterSpecS3Import)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalingConfiguration != nil {
		in, out := &in.ScalingConfiguration, &out.ScalingConfiguration
		*out = new(ClusterSpecScalingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.SkipFinalSnapshot != nil {
		in, out := &in.SkipFinalSnapshot, &out.SkipFinalSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotIdentifier != nil {
		in, out := &in.SnapshotIdentifier, &out.SnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.SourceRegion != nil {
		in, out := &in.SourceRegion, &out.SourceRegion
		*out = new(string)
		**out = **in
	}
	if in.StorageEncrypted != nil {
		in, out := &in.StorageEncrypted, &out.StorageEncrypted
		*out = new(bool)
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
	if in.VpcSecurityGroupIDS != nil {
		in, out := &in.VpcSecurityGroupIDS, &out.VpcSecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecResource.
func (in *ClusterSpecResource) DeepCopy() *ClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecRestoreToPointInTime) DeepCopyInto(out *ClusterSpecRestoreToPointInTime) {
	*out = *in
	if in.RestoreToTime != nil {
		in, out := &in.RestoreToTime, &out.RestoreToTime
		*out = new(string)
		**out = **in
	}
	if in.RestoreType != nil {
		in, out := &in.RestoreType, &out.RestoreType
		*out = new(string)
		**out = **in
	}
	if in.SourceClusterIdentifier != nil {
		in, out := &in.SourceClusterIdentifier, &out.SourceClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.UseLatestRestorableTime != nil {
		in, out := &in.UseLatestRestorableTime, &out.UseLatestRestorableTime
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecRestoreToPointInTime.
func (in *ClusterSpecRestoreToPointInTime) DeepCopy() *ClusterSpecRestoreToPointInTime {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecRestoreToPointInTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecS3Import) DeepCopyInto(out *ClusterSpecS3Import) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.BucketPrefix != nil {
		in, out := &in.BucketPrefix, &out.BucketPrefix
		*out = new(string)
		**out = **in
	}
	if in.IngestionRole != nil {
		in, out := &in.IngestionRole, &out.IngestionRole
		*out = new(string)
		**out = **in
	}
	if in.SourceEngine != nil {
		in, out := &in.SourceEngine, &out.SourceEngine
		*out = new(string)
		**out = **in
	}
	if in.SourceEngineVersion != nil {
		in, out := &in.SourceEngineVersion, &out.SourceEngineVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecS3Import.
func (in *ClusterSpecS3Import) DeepCopy() *ClusterSpecS3Import {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecS3Import)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecScalingConfiguration) DeepCopyInto(out *ClusterSpecScalingConfiguration) {
	*out = *in
	if in.AutoPause != nil {
		in, out := &in.AutoPause, &out.AutoPause
		*out = new(bool)
		**out = **in
	}
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(int64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(int64)
		**out = **in
	}
	if in.SecondsUntilAutoPause != nil {
		in, out := &in.SecondsUntilAutoPause, &out.SecondsUntilAutoPause
		*out = new(int64)
		**out = **in
	}
	if in.TimeoutAction != nil {
		in, out := &in.TimeoutAction, &out.TimeoutAction
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecScalingConfiguration.
func (in *ClusterSpecScalingConfiguration) DeepCopy() *ClusterSpecScalingConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecScalingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalCluster) DeepCopyInto(out *GlobalCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalCluster.
func (in *GlobalCluster) DeepCopy() *GlobalCluster {
	if in == nil {
		return nil
	}
	out := new(GlobalCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalClusterList) DeepCopyInto(out *GlobalClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalClusterList.
func (in *GlobalClusterList) DeepCopy() *GlobalClusterList {
	if in == nil {
		return nil
	}
	out := new(GlobalClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalClusterSpec) DeepCopyInto(out *GlobalClusterSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(GlobalClusterSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalClusterSpec.
func (in *GlobalClusterSpec) DeepCopy() *GlobalClusterSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalClusterSpecGlobalClusterMembers) DeepCopyInto(out *GlobalClusterSpecGlobalClusterMembers) {
	*out = *in
	if in.DbClusterArn != nil {
		in, out := &in.DbClusterArn, &out.DbClusterArn
		*out = new(string)
		**out = **in
	}
	if in.IsWriter != nil {
		in, out := &in.IsWriter, &out.IsWriter
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalClusterSpecGlobalClusterMembers.
func (in *GlobalClusterSpecGlobalClusterMembers) DeepCopy() *GlobalClusterSpecGlobalClusterMembers {
	if in == nil {
		return nil
	}
	out := new(GlobalClusterSpecGlobalClusterMembers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalClusterSpecResource) DeepCopyInto(out *GlobalClusterSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.GlobalClusterIdentifier != nil {
		in, out := &in.GlobalClusterIdentifier, &out.GlobalClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.GlobalClusterMembers != nil {
		in, out := &in.GlobalClusterMembers, &out.GlobalClusterMembers
		*out = make([]GlobalClusterSpecGlobalClusterMembers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GlobalClusterResourceID != nil {
		in, out := &in.GlobalClusterResourceID, &out.GlobalClusterResourceID
		*out = new(string)
		**out = **in
	}
	if in.SourceDbClusterIdentifier != nil {
		in, out := &in.SourceDbClusterIdentifier, &out.SourceDbClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.StorageEncrypted != nil {
		in, out := &in.StorageEncrypted, &out.StorageEncrypted
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalClusterSpecResource.
func (in *GlobalClusterSpecResource) DeepCopy() *GlobalClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(GlobalClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalClusterStatus) DeepCopyInto(out *GlobalClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalClusterStatus.
func (in *GlobalClusterStatus) DeepCopy() *GlobalClusterStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalClusterStatus)
	in.DeepCopyInto(out)
	return out
}