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

type Branch struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BranchSpec   `json:"spec,omitempty"`
	Status            BranchStatus `json:"status,omitempty"`
}

type BranchSpec struct {
	KubeformOutput *BranchSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource BranchSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type BranchSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AppID *string `json:"appID" tf:"app_id"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AssociatedResources []string `json:"associatedResources,omitempty" tf:"associated_resources"`
	// +optional
	BackendEnvironmentArn *string `json:"backendEnvironmentArn,omitempty" tf:"backend_environment_arn"`
	// +optional
	BasicAuthCredentials *string `json:"-" sensitive:"true" tf:"basic_auth_credentials"`
	BranchName           *string `json:"branchName" tf:"branch_name"`
	// +optional
	CustomDomains []string `json:"customDomains,omitempty" tf:"custom_domains"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DestinationBranch *string `json:"destinationBranch,omitempty" tf:"destination_branch"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	EnableAutoBuild *bool `json:"enableAutoBuild,omitempty" tf:"enable_auto_build"`
	// +optional
	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth"`
	// +optional
	EnableNotification *bool `json:"enableNotification,omitempty" tf:"enable_notification"`
	// +optional
	EnablePerformanceMode *bool `json:"enablePerformanceMode,omitempty" tf:"enable_performance_mode"`
	// +optional
	EnablePullRequestPreview *bool `json:"enablePullRequestPreview,omitempty" tf:"enable_pull_request_preview"`
	// +optional
	EnvironmentVariables *map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables"`
	// +optional
	Framework *string `json:"framework,omitempty" tf:"framework"`
	// +optional
	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName,omitempty" tf:"pull_request_environment_name"`
	// +optional
	SourceBranch *string `json:"sourceBranch,omitempty" tf:"source_branch"`
	// +optional
	Stage *string `json:"stage,omitempty" tf:"stage"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Ttl *string `json:"ttl,omitempty" tf:"ttl"`
}

type BranchStatus struct {
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

// BranchList is a list of Branchs
type BranchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Branch CRD objects
	Items []Branch `json:"items,omitempty"`
}