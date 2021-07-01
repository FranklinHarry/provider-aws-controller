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
func (in *AccountAssignment) DeepCopyInto(out *AccountAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountAssignment.
func (in *AccountAssignment) DeepCopy() *AccountAssignment {
	if in == nil {
		return nil
	}
	out := new(AccountAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountAssignmentList) DeepCopyInto(out *AccountAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountAssignmentList.
func (in *AccountAssignmentList) DeepCopy() *AccountAssignmentList {
	if in == nil {
		return nil
	}
	out := new(AccountAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountAssignmentSpec) DeepCopyInto(out *AccountAssignmentSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(AccountAssignmentSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountAssignmentSpec.
func (in *AccountAssignmentSpec) DeepCopy() *AccountAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(AccountAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountAssignmentSpecResource) DeepCopyInto(out *AccountAssignmentSpecResource) {
	*out = *in
	if in.InstanceArn != nil {
		in, out := &in.InstanceArn, &out.InstanceArn
		*out = new(string)
		**out = **in
	}
	if in.PermissionSetArn != nil {
		in, out := &in.PermissionSetArn, &out.PermissionSetArn
		*out = new(string)
		**out = **in
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(string)
		**out = **in
	}
	if in.TargetID != nil {
		in, out := &in.TargetID, &out.TargetID
		*out = new(string)
		**out = **in
	}
	if in.TargetType != nil {
		in, out := &in.TargetType, &out.TargetType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountAssignmentSpecResource.
func (in *AccountAssignmentSpecResource) DeepCopy() *AccountAssignmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(AccountAssignmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountAssignmentStatus) DeepCopyInto(out *AccountAssignmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountAssignmentStatus.
func (in *AccountAssignmentStatus) DeepCopy() *AccountAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(AccountAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyAttachment) DeepCopyInto(out *ManagedPolicyAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyAttachment.
func (in *ManagedPolicyAttachment) DeepCopy() *ManagedPolicyAttachment {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedPolicyAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyAttachmentList) DeepCopyInto(out *ManagedPolicyAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedPolicyAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyAttachmentList.
func (in *ManagedPolicyAttachmentList) DeepCopy() *ManagedPolicyAttachmentList {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedPolicyAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyAttachmentSpec) DeepCopyInto(out *ManagedPolicyAttachmentSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ManagedPolicyAttachmentSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyAttachmentSpec.
func (in *ManagedPolicyAttachmentSpec) DeepCopy() *ManagedPolicyAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyAttachmentSpecResource) DeepCopyInto(out *ManagedPolicyAttachmentSpecResource) {
	*out = *in
	if in.InstanceArn != nil {
		in, out := &in.InstanceArn, &out.InstanceArn
		*out = new(string)
		**out = **in
	}
	if in.ManagedPolicyArn != nil {
		in, out := &in.ManagedPolicyArn, &out.ManagedPolicyArn
		*out = new(string)
		**out = **in
	}
	if in.ManagedPolicyName != nil {
		in, out := &in.ManagedPolicyName, &out.ManagedPolicyName
		*out = new(string)
		**out = **in
	}
	if in.PermissionSetArn != nil {
		in, out := &in.PermissionSetArn, &out.PermissionSetArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyAttachmentSpecResource.
func (in *ManagedPolicyAttachmentSpecResource) DeepCopy() *ManagedPolicyAttachmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyAttachmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyAttachmentStatus) DeepCopyInto(out *ManagedPolicyAttachmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyAttachmentStatus.
func (in *ManagedPolicyAttachmentStatus) DeepCopy() *ManagedPolicyAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSet) DeepCopyInto(out *PermissionSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSet.
func (in *PermissionSet) DeepCopy() *PermissionSet {
	if in == nil {
		return nil
	}
	out := new(PermissionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PermissionSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetInlinePolicy) DeepCopyInto(out *PermissionSetInlinePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetInlinePolicy.
func (in *PermissionSetInlinePolicy) DeepCopy() *PermissionSetInlinePolicy {
	if in == nil {
		return nil
	}
	out := new(PermissionSetInlinePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PermissionSetInlinePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetInlinePolicyList) DeepCopyInto(out *PermissionSetInlinePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PermissionSetInlinePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetInlinePolicyList.
func (in *PermissionSetInlinePolicyList) DeepCopy() *PermissionSetInlinePolicyList {
	if in == nil {
		return nil
	}
	out := new(PermissionSetInlinePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PermissionSetInlinePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetInlinePolicySpec) DeepCopyInto(out *PermissionSetInlinePolicySpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(PermissionSetInlinePolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetInlinePolicySpec.
func (in *PermissionSetInlinePolicySpec) DeepCopy() *PermissionSetInlinePolicySpec {
	if in == nil {
		return nil
	}
	out := new(PermissionSetInlinePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetInlinePolicySpecResource) DeepCopyInto(out *PermissionSetInlinePolicySpecResource) {
	*out = *in
	if in.InlinePolicy != nil {
		in, out := &in.InlinePolicy, &out.InlinePolicy
		*out = new(string)
		**out = **in
	}
	if in.InstanceArn != nil {
		in, out := &in.InstanceArn, &out.InstanceArn
		*out = new(string)
		**out = **in
	}
	if in.PermissionSetArn != nil {
		in, out := &in.PermissionSetArn, &out.PermissionSetArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetInlinePolicySpecResource.
func (in *PermissionSetInlinePolicySpecResource) DeepCopy() *PermissionSetInlinePolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(PermissionSetInlinePolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetInlinePolicyStatus) DeepCopyInto(out *PermissionSetInlinePolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetInlinePolicyStatus.
func (in *PermissionSetInlinePolicyStatus) DeepCopy() *PermissionSetInlinePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PermissionSetInlinePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetList) DeepCopyInto(out *PermissionSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PermissionSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetList.
func (in *PermissionSetList) DeepCopy() *PermissionSetList {
	if in == nil {
		return nil
	}
	out := new(PermissionSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PermissionSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetSpec) DeepCopyInto(out *PermissionSetSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(PermissionSetSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetSpec.
func (in *PermissionSetSpec) DeepCopy() *PermissionSetSpec {
	if in == nil {
		return nil
	}
	out := new(PermissionSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetSpecResource) DeepCopyInto(out *PermissionSetSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InstanceArn != nil {
		in, out := &in.InstanceArn, &out.InstanceArn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RelayState != nil {
		in, out := &in.RelayState, &out.RelayState
		*out = new(string)
		**out = **in
	}
	if in.SessionDuration != nil {
		in, out := &in.SessionDuration, &out.SessionDuration
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetSpecResource.
func (in *PermissionSetSpecResource) DeepCopy() *PermissionSetSpecResource {
	if in == nil {
		return nil
	}
	out := new(PermissionSetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSetStatus) DeepCopyInto(out *PermissionSetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSetStatus.
func (in *PermissionSetStatus) DeepCopy() *PermissionSetStatus {
	if in == nil {
		return nil
	}
	out := new(PermissionSetStatus)
	in.DeepCopyInto(out)
	return out
}