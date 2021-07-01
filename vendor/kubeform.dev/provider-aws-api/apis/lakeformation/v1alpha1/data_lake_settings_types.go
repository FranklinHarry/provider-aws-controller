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

type DataLakeSettings struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeSettingsSpec   `json:"spec,omitempty"`
	Status            DataLakeSettingsStatus `json:"status,omitempty"`
}

type DataLakeSettingsSpecCreateDatabaseDefaultPermissions struct {
	// +optional
	Permissions []string `json:"permissions,omitempty" tf:"permissions"`
	// +optional
	Principal *string `json:"principal,omitempty" tf:"principal"`
}

type DataLakeSettingsSpecCreateTableDefaultPermissions struct {
	// +optional
	Permissions []string `json:"permissions,omitempty" tf:"permissions"`
	// +optional
	Principal *string `json:"principal,omitempty" tf:"principal"`
}

type DataLakeSettingsSpec struct {
	KubeformOutput *DataLakeSettingsSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource DataLakeSettingsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DataLakeSettingsSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Admins []string `json:"admins,omitempty" tf:"admins"`
	// +optional
	CatalogID *string `json:"catalogID,omitempty" tf:"catalog_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	CreateDatabaseDefaultPermissions []DataLakeSettingsSpecCreateDatabaseDefaultPermissions `json:"createDatabaseDefaultPermissions,omitempty" tf:"create_database_default_permissions"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	CreateTableDefaultPermissions []DataLakeSettingsSpecCreateTableDefaultPermissions `json:"createTableDefaultPermissions,omitempty" tf:"create_table_default_permissions"`
	// +optional
	TrustedResourceOwners []string `json:"trustedResourceOwners,omitempty" tf:"trusted_resource_owners"`
}

type DataLakeSettingsStatus struct {
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

// DataLakeSettingsList is a list of DataLakeSettingss
type DataLakeSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataLakeSettings CRD objects
	Items []DataLakeSettings `json:"items,omitempty"`
}
