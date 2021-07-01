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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type LifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LifecyclePolicySpec   `json:"spec,omitempty"`
	Status            LifecyclePolicyStatus `json:"status,omitempty"`
}

type LifecyclePolicySpecPolicyDetailsScheduleCreateRule struct {
	Interval *int64 `json:"interval" tf:"interval"`
	// +optional
	IntervalUnit *string `json:"intervalUnit,omitempty" tf:"interval_unit"`
	// +optional
	Times []string `json:"times,omitempty" tf:"times"`
}

type LifecyclePolicySpecPolicyDetailsScheduleRetainRule struct {
	Count *int64 `json:"count" tf:"count"`
}

type LifecyclePolicySpecPolicyDetailsSchedule struct {
	// +optional
	CopyTags   *bool                                               `json:"copyTags,omitempty" tf:"copy_tags"`
	CreateRule *LifecyclePolicySpecPolicyDetailsScheduleCreateRule `json:"createRule" tf:"create_rule"`
	Name       *string                                             `json:"name" tf:"name"`
	RetainRule *LifecyclePolicySpecPolicyDetailsScheduleRetainRule `json:"retainRule" tf:"retain_rule"`
	// +optional
	TagsToAdd *map[string]string `json:"tagsToAdd,omitempty" tf:"tags_to_add"`
}

type LifecyclePolicySpecPolicyDetails struct {
	ResourceTypes []string                                   `json:"resourceTypes" tf:"resource_types"`
	Schedule      []LifecyclePolicySpecPolicyDetailsSchedule `json:"schedule" tf:"schedule"`
	TargetTags    *map[string]string                         `json:"targetTags" tf:"target_tags"`
}

type LifecyclePolicySpec struct {
	KubeformOutput *LifecyclePolicySpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource LifecyclePolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LifecyclePolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn              *string                           `json:"arn,omitempty" tf:"arn"`
	Description      *string                           `json:"description" tf:"description"`
	ExecutionRoleArn *string                           `json:"executionRoleArn" tf:"execution_role_arn"`
	PolicyDetails    *LifecyclePolicySpecPolicyDetails `json:"policyDetails" tf:"policy_details"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type LifecyclePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LifecyclePolicyList is a list of LifecyclePolicys
type LifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LifecyclePolicy CRD objects
	Items []LifecyclePolicy `json:"items,omitempty"`
}