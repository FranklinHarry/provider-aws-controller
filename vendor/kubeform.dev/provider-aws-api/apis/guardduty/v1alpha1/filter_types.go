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

type Filter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FilterSpec   `json:"spec,omitempty"`
	Status            FilterStatus `json:"status,omitempty"`
}

type FilterSpecFindingCriteriaCriterion struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	Equals []string `json:"equals,omitempty" tf:"equals"`
	Field  *string  `json:"field" tf:"field"`
	// +optional
	GreaterThan *string `json:"greaterThan,omitempty" tf:"greater_than"`
	// +optional
	GreaterThanOrEqual *string `json:"greaterThanOrEqual,omitempty" tf:"greater_than_or_equal"`
	// +optional
	LessThan *string `json:"lessThan,omitempty" tf:"less_than"`
	// +optional
	LessThanOrEqual *string `json:"lessThanOrEqual,omitempty" tf:"less_than_or_equal"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	NotEquals []string `json:"notEquals,omitempty" tf:"not_equals"`
}

type FilterSpecFindingCriteria struct {
	// +kubebuilder:validation:MinItems=1
	Criterion []FilterSpecFindingCriteriaCriterion `json:"criterion" tf:"criterion"`
}

type FilterSpec struct {
	KubeformOutput *FilterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource FilterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FilterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action *string `json:"action" tf:"action"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description     *string                    `json:"description,omitempty" tf:"description"`
	DetectorID      *string                    `json:"detectorID" tf:"detector_id"`
	FindingCriteria *FilterSpecFindingCriteria `json:"findingCriteria" tf:"finding_criteria"`
	Name            *string                    `json:"name" tf:"name"`
	Rank            *int64                     `json:"rank" tf:"rank"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type FilterStatus struct {
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

// FilterList is a list of Filters
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Filter CRD objects
	Items []Filter `json:"items,omitempty"`
}