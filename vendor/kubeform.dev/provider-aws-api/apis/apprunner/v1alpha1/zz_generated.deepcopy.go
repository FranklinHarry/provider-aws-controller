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
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationVersion) DeepCopyInto(out *AutoScalingConfigurationVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationVersion.
func (in *AutoScalingConfigurationVersion) DeepCopy() *AutoScalingConfigurationVersion {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoScalingConfigurationVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationVersionList) DeepCopyInto(out *AutoScalingConfigurationVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutoScalingConfigurationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationVersionList.
func (in *AutoScalingConfigurationVersionList) DeepCopy() *AutoScalingConfigurationVersionList {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoScalingConfigurationVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationVersionSpec) DeepCopyInto(out *AutoScalingConfigurationVersionSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(AutoScalingConfigurationVersionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationVersionSpec.
func (in *AutoScalingConfigurationVersionSpec) DeepCopy() *AutoScalingConfigurationVersionSpec {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationVersionSpecResource) DeepCopyInto(out *AutoScalingConfigurationVersionSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoScalingConfigurationName != nil {
		in, out := &in.AutoScalingConfigurationName, &out.AutoScalingConfigurationName
		*out = new(string)
		**out = **in
	}
	if in.AutoScalingConfigurationRevision != nil {
		in, out := &in.AutoScalingConfigurationRevision, &out.AutoScalingConfigurationRevision
		*out = new(int64)
		**out = **in
	}
	if in.Latest != nil {
		in, out := &in.Latest, &out.Latest
		*out = new(bool)
		**out = **in
	}
	if in.MaxConcurrency != nil {
		in, out := &in.MaxConcurrency, &out.MaxConcurrency
		*out = new(int64)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int64)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationVersionSpecResource.
func (in *AutoScalingConfigurationVersionSpecResource) DeepCopy() *AutoScalingConfigurationVersionSpecResource {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationVersionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationVersionStatus) DeepCopyInto(out *AutoScalingConfigurationVersionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationVersionStatus.
func (in *AutoScalingConfigurationVersionStatus) DeepCopy() *AutoScalingConfigurationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Connection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionList) DeepCopyInto(out *ConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Connection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionList.
func (in *ConnectionList) DeepCopy() *ConnectionList {
	if in == nil {
		return nil
	}
	out := new(ConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpec) DeepCopyInto(out *ConnectionSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ConnectionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpec.
func (in *ConnectionSpec) DeepCopy() *ConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpecResource) DeepCopyInto(out *ConnectionSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ConnectionName != nil {
		in, out := &in.ConnectionName, &out.ConnectionName
		*out = new(string)
		**out = **in
	}
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpecResource.
func (in *ConnectionSpecResource) DeepCopy() *ConnectionSpecResource {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStatus) DeepCopyInto(out *ConnectionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStatus.
func (in *ConnectionStatus) DeepCopy() *ConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainAssociation) DeepCopyInto(out *CustomDomainAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainAssociation.
func (in *CustomDomainAssociation) DeepCopy() *CustomDomainAssociation {
	if in == nil {
		return nil
	}
	out := new(CustomDomainAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomDomainAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainAssociationList) DeepCopyInto(out *CustomDomainAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomDomainAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainAssociationList.
func (in *CustomDomainAssociationList) DeepCopy() *CustomDomainAssociationList {
	if in == nil {
		return nil
	}
	out := new(CustomDomainAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomDomainAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainAssociationSpec) DeepCopyInto(out *CustomDomainAssociationSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(CustomDomainAssociationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainAssociationSpec.
func (in *CustomDomainAssociationSpec) DeepCopy() *CustomDomainAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(CustomDomainAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainAssociationSpecCertificateValidationRecords) DeepCopyInto(out *CustomDomainAssociationSpecCertificateValidationRecords) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainAssociationSpecCertificateValidationRecords.
func (in *CustomDomainAssociationSpecCertificateValidationRecords) DeepCopy() *CustomDomainAssociationSpecCertificateValidationRecords {
	if in == nil {
		return nil
	}
	out := new(CustomDomainAssociationSpecCertificateValidationRecords)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainAssociationSpecResource) DeepCopyInto(out *CustomDomainAssociationSpecResource) {
	*out = *in
	if in.CertificateValidationRecords != nil {
		in, out := &in.CertificateValidationRecords, &out.CertificateValidationRecords
		*out = make([]CustomDomainAssociationSpecCertificateValidationRecords, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DnsTarget != nil {
		in, out := &in.DnsTarget, &out.DnsTarget
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.EnableWWWSubdomain != nil {
		in, out := &in.EnableWWWSubdomain, &out.EnableWWWSubdomain
		*out = new(bool)
		**out = **in
	}
	if in.ServiceArn != nil {
		in, out := &in.ServiceArn, &out.ServiceArn
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainAssociationSpecResource.
func (in *CustomDomainAssociationSpecResource) DeepCopy() *CustomDomainAssociationSpecResource {
	if in == nil {
		return nil
	}
	out := new(CustomDomainAssociationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainAssociationStatus) DeepCopyInto(out *CustomDomainAssociationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainAssociationStatus.
func (in *CustomDomainAssociationStatus) DeepCopy() *CustomDomainAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(CustomDomainAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Service) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceList) DeepCopyInto(out *ServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceList.
func (in *ServiceList) DeepCopy() *ServiceList {
	if in == nil {
		return nil
	}
	out := new(ServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ServiceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecEncryptionConfiguration) DeepCopyInto(out *ServiceSpecEncryptionConfiguration) {
	*out = *in
	if in.KmsKey != nil {
		in, out := &in.KmsKey, &out.KmsKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecEncryptionConfiguration.
func (in *ServiceSpecEncryptionConfiguration) DeepCopy() *ServiceSpecEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecEncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecHealthCheckConfiguration) DeepCopyInto(out *ServiceSpecHealthCheckConfiguration) {
	*out = *in
	if in.HealthyThreshold != nil {
		in, out := &in.HealthyThreshold, &out.HealthyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.UnhealthyThreshold != nil {
		in, out := &in.UnhealthyThreshold, &out.UnhealthyThreshold
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecHealthCheckConfiguration.
func (in *ServiceSpecHealthCheckConfiguration) DeepCopy() *ServiceSpecHealthCheckConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecHealthCheckConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecInstanceConfiguration) DeepCopyInto(out *ServiceSpecInstanceConfiguration) {
	*out = *in
	if in.Cpu != nil {
		in, out := &in.Cpu, &out.Cpu
		*out = new(string)
		**out = **in
	}
	if in.InstanceRoleArn != nil {
		in, out := &in.InstanceRoleArn, &out.InstanceRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecInstanceConfiguration.
func (in *ServiceSpecInstanceConfiguration) DeepCopy() *ServiceSpecInstanceConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecInstanceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecResource) DeepCopyInto(out *ServiceSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoScalingConfigurationArn != nil {
		in, out := &in.AutoScalingConfigurationArn, &out.AutoScalingConfigurationArn
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = new(ServiceSpecEncryptionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheckConfiguration != nil {
		in, out := &in.HealthCheckConfiguration, &out.HealthCheckConfiguration
		*out = new(ServiceSpecHealthCheckConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceConfiguration != nil {
		in, out := &in.InstanceConfiguration, &out.InstanceConfiguration
		*out = new(ServiceSpecInstanceConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.ServiceURL != nil {
		in, out := &in.ServiceURL, &out.ServiceURL
		*out = new(string)
		**out = **in
	}
	if in.SourceConfiguration != nil {
		in, out := &in.SourceConfiguration, &out.SourceConfiguration
		*out = new(ServiceSpecSourceConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecResource.
func (in *ServiceSpecResource) DeepCopy() *ServiceSpecResource {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfiguration) DeepCopyInto(out *ServiceSpecSourceConfiguration) {
	*out = *in
	if in.AuthenticationConfiguration != nil {
		in, out := &in.AuthenticationConfiguration, &out.AuthenticationConfiguration
		*out = new(ServiceSpecSourceConfigurationAuthenticationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoDeploymentsEnabled != nil {
		in, out := &in.AutoDeploymentsEnabled, &out.AutoDeploymentsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CodeRepository != nil {
		in, out := &in.CodeRepository, &out.CodeRepository
		*out = new(ServiceSpecSourceConfigurationCodeRepository)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageRepository != nil {
		in, out := &in.ImageRepository, &out.ImageRepository
		*out = new(ServiceSpecSourceConfigurationImageRepository)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfiguration.
func (in *ServiceSpecSourceConfiguration) DeepCopy() *ServiceSpecSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfigurationAuthenticationConfiguration) DeepCopyInto(out *ServiceSpecSourceConfigurationAuthenticationConfiguration) {
	*out = *in
	if in.AccessRoleArn != nil {
		in, out := &in.AccessRoleArn, &out.AccessRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ConnectionArn != nil {
		in, out := &in.ConnectionArn, &out.ConnectionArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfigurationAuthenticationConfiguration.
func (in *ServiceSpecSourceConfigurationAuthenticationConfiguration) DeepCopy() *ServiceSpecSourceConfigurationAuthenticationConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfigurationAuthenticationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfigurationCodeRepository) DeepCopyInto(out *ServiceSpecSourceConfigurationCodeRepository) {
	*out = *in
	if in.CodeConfiguration != nil {
		in, out := &in.CodeConfiguration, &out.CodeConfiguration
		*out = new(ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.RepositoryURL != nil {
		in, out := &in.RepositoryURL, &out.RepositoryURL
		*out = new(string)
		**out = **in
	}
	if in.SourceCodeVersion != nil {
		in, out := &in.SourceCodeVersion, &out.SourceCodeVersion
		*out = new(ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfigurationCodeRepository.
func (in *ServiceSpecSourceConfigurationCodeRepository) DeepCopy() *ServiceSpecSourceConfigurationCodeRepository {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfigurationCodeRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration) DeepCopyInto(out *ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration) {
	*out = *in
	if in.CodeConfigurationValues != nil {
		in, out := &in.CodeConfigurationValues, &out.CodeConfigurationValues
		*out = new(ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigurationSource != nil {
		in, out := &in.ConfigurationSource, &out.ConfigurationSource
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration.
func (in *ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration) DeepCopy() *ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues) DeepCopyInto(out *ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues) {
	*out = *in
	if in.BuildCommand != nil {
		in, out := &in.BuildCommand, &out.BuildCommand
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.RuntimeEnvironmentVariables != nil {
		in, out := &in.RuntimeEnvironmentVariables, &out.RuntimeEnvironmentVariables
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.StartCommand != nil {
		in, out := &in.StartCommand, &out.StartCommand
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues.
func (in *ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues) DeepCopy() *ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion) DeepCopyInto(out *ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion.
func (in *ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion) DeepCopy() *ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfigurationImageRepository) DeepCopyInto(out *ServiceSpecSourceConfigurationImageRepository) {
	*out = *in
	if in.ImageConfiguration != nil {
		in, out := &in.ImageConfiguration, &out.ImageConfiguration
		*out = new(ServiceSpecSourceConfigurationImageRepositoryImageConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageIdentifier != nil {
		in, out := &in.ImageIdentifier, &out.ImageIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ImageRepositoryType != nil {
		in, out := &in.ImageRepositoryType, &out.ImageRepositoryType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfigurationImageRepository.
func (in *ServiceSpecSourceConfigurationImageRepository) DeepCopy() *ServiceSpecSourceConfigurationImageRepository {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfigurationImageRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecSourceConfigurationImageRepositoryImageConfiguration) DeepCopyInto(out *ServiceSpecSourceConfigurationImageRepositoryImageConfiguration) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.RuntimeEnvironmentVariables != nil {
		in, out := &in.RuntimeEnvironmentVariables, &out.RuntimeEnvironmentVariables
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.StartCommand != nil {
		in, out := &in.StartCommand, &out.StartCommand
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecSourceConfigurationImageRepositoryImageConfiguration.
func (in *ServiceSpecSourceConfigurationImageRepositoryImageConfiguration) DeepCopy() *ServiceSpecSourceConfigurationImageRepositoryImageConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecSourceConfigurationImageRepositoryImageConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}