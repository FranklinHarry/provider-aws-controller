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

type CloudformationType struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationTypeSpec   `json:"spec,omitempty"`
	Status            CloudformationTypeStatus `json:"status,omitempty"`
}

type CloudformationTypeSpecLoggingConfig struct {
	LogGroupName *string `json:"logGroupName" tf:"log_group_name"`
	LogRoleArn   *string `json:"logRoleArn" tf:"log_role_arn"`
}

type CloudformationTypeSpec struct {
	KubeformOutput *CloudformationTypeSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource CloudformationTypeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CloudformationTypeSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	DefaultVersionID *string `json:"defaultVersionID,omitempty" tf:"default_version_id"`
	// +optional
	DeprecatedStatus *string `json:"deprecatedStatus,omitempty" tf:"deprecated_status"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DocumentationURL *string `json:"documentationURL,omitempty" tf:"documentation_url"`
	// +optional
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn"`
	// +optional
	IsDefaultVersion *bool `json:"isDefaultVersion,omitempty" tf:"is_default_version"`
	// +optional
	LoggingConfig *CloudformationTypeSpecLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config"`
	// +optional
	ProvisioningType *string `json:"provisioningType,omitempty" tf:"provisioning_type"`
	// +optional
	Schema               *string `json:"schema,omitempty" tf:"schema"`
	SchemaHandlerPackage *string `json:"schemaHandlerPackage" tf:"schema_handler_package"`
	// +optional
	SourceURL *string `json:"sourceURL,omitempty" tf:"source_url"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	TypeArn  *string `json:"typeArn,omitempty" tf:"type_arn"`
	TypeName *string `json:"typeName" tf:"type_name"`
	// +optional
	VersionID *string `json:"versionID,omitempty" tf:"version_id"`
	// +optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility"`
}

type CloudformationTypeStatus struct {
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

// CloudformationTypeList is a list of CloudformationTypes
type CloudformationTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudformationType CRD objects
	Items []CloudformationType `json:"items,omitempty"`
}