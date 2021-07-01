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

type LocationFsxWindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationFsxWindowsFileSystemSpec   `json:"spec,omitempty"`
	Status            LocationFsxWindowsFileSystemStatus `json:"status,omitempty"`
}

type LocationFsxWindowsFileSystemSpec struct {
	KubeformOutput *LocationFsxWindowsFileSystemSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource LocationFsxWindowsFileSystemSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type LocationFsxWindowsFileSystemSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time"`
	// +optional
	Domain           *string `json:"domain,omitempty" tf:"domain"`
	FsxFilesystemArn *string `json:"fsxFilesystemArn" tf:"fsx_filesystem_arn"`
	Password         *string `json:"-" sensitive:"true" tf:"password"`
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:MinItems=1
	SecurityGroupArns []string `json:"securityGroupArns" tf:"security_group_arns"`
	// +optional
	Subdirectory *string `json:"subdirectory,omitempty" tf:"subdirectory"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Uri  *string `json:"uri,omitempty" tf:"uri"`
	User *string `json:"user" tf:"user"`
}

type LocationFsxWindowsFileSystemStatus struct {
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

// LocationFsxWindowsFileSystemList is a list of LocationFsxWindowsFileSystems
type LocationFsxWindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LocationFsxWindowsFileSystem CRD objects
	Items []LocationFsxWindowsFileSystem `json:"items,omitempty"`
}