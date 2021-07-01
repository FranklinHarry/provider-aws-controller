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
func (in *AdminAccount) DeepCopyInto(out *AdminAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminAccount.
func (in *AdminAccount) DeepCopy() *AdminAccount {
	if in == nil {
		return nil
	}
	out := new(AdminAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminAccountList) DeepCopyInto(out *AdminAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdminAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminAccountList.
func (in *AdminAccountList) DeepCopy() *AdminAccountList {
	if in == nil {
		return nil
	}
	out := new(AdminAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminAccountSpec) DeepCopyInto(out *AdminAccountSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(AdminAccountSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminAccountSpec.
func (in *AdminAccountSpec) DeepCopy() *AdminAccountSpec {
	if in == nil {
		return nil
	}
	out := new(AdminAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminAccountSpecResource) DeepCopyInto(out *AdminAccountSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminAccountSpecResource.
func (in *AdminAccountSpecResource) DeepCopy() *AdminAccountSpecResource {
	if in == nil {
		return nil
	}
	out := new(AdminAccountSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminAccountStatus) DeepCopyInto(out *AdminAccountStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminAccountStatus.
func (in *AdminAccountStatus) DeepCopy() *AdminAccountStatus {
	if in == nil {
		return nil
	}
	out := new(AdminAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(PolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecExcludeMap) DeepCopyInto(out *PolicySpecExcludeMap) {
	*out = *in
	if in.Account != nil {
		in, out := &in.Account, &out.Account
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Orgunit != nil {
		in, out := &in.Orgunit, &out.Orgunit
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecExcludeMap.
func (in *PolicySpecExcludeMap) DeepCopy() *PolicySpecExcludeMap {
	if in == nil {
		return nil
	}
	out := new(PolicySpecExcludeMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecIncludeMap) DeepCopyInto(out *PolicySpecIncludeMap) {
	*out = *in
	if in.Account != nil {
		in, out := &in.Account, &out.Account
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Orgunit != nil {
		in, out := &in.Orgunit, &out.Orgunit
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecIncludeMap.
func (in *PolicySpecIncludeMap) DeepCopy() *PolicySpecIncludeMap {
	if in == nil {
		return nil
	}
	out := new(PolicySpecIncludeMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecResource) DeepCopyInto(out *PolicySpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DeleteAllPolicyResources != nil {
		in, out := &in.DeleteAllPolicyResources, &out.DeleteAllPolicyResources
		*out = new(bool)
		**out = **in
	}
	if in.ExcludeMap != nil {
		in, out := &in.ExcludeMap, &out.ExcludeMap
		*out = new(PolicySpecExcludeMap)
		(*in).DeepCopyInto(*out)
	}
	if in.ExcludeResourceTags != nil {
		in, out := &in.ExcludeResourceTags, &out.ExcludeResourceTags
		*out = new(bool)
		**out = **in
	}
	if in.IncludeMap != nil {
		in, out := &in.IncludeMap, &out.IncludeMap
		*out = new(PolicySpecIncludeMap)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyUpdateToken != nil {
		in, out := &in.PolicyUpdateToken, &out.PolicyUpdateToken
		*out = new(string)
		**out = **in
	}
	if in.RemediationEnabled != nil {
		in, out := &in.RemediationEnabled, &out.RemediationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceTags != nil {
		in, out := &in.ResourceTags, &out.ResourceTags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.ResourceTypeList != nil {
		in, out := &in.ResourceTypeList, &out.ResourceTypeList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityServicePolicyData != nil {
		in, out := &in.SecurityServicePolicyData, &out.SecurityServicePolicyData
		*out = new(PolicySpecSecurityServicePolicyData)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecResource.
func (in *PolicySpecResource) DeepCopy() *PolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(PolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecSecurityServicePolicyData) DeepCopyInto(out *PolicySpecSecurityServicePolicyData) {
	*out = *in
	if in.ManagedServiceData != nil {
		in, out := &in.ManagedServiceData, &out.ManagedServiceData
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecSecurityServicePolicyData.
func (in *PolicySpecSecurityServicePolicyData) DeepCopy() *PolicySpecSecurityServicePolicyData {
	if in == nil {
		return nil
	}
	out := new(PolicySpecSecurityServicePolicyData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}
