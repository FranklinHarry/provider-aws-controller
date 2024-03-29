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

type Service struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec,omitempty"`
	Status            ServiceStatus `json:"status,omitempty"`
}

type ServiceSpecDnsConfigDnsRecords struct {
	Ttl  *int64  `json:"ttl" tf:"ttl"`
	Type *string `json:"type" tf:"type"`
}

type ServiceSpecDnsConfig struct {
	DnsRecords  []ServiceSpecDnsConfigDnsRecords `json:"dnsRecords" tf:"dns_records"`
	NamespaceID *string                          `json:"namespaceID" tf:"namespace_id"`
	// +optional
	RoutingPolicy *string `json:"routingPolicy,omitempty" tf:"routing_policy"`
}

type ServiceSpecHealthCheckConfig struct {
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
	// +optional
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceSpecHealthCheckCustomConfig struct {
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
}

type ServiceSpec struct {
	State *ServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DnsConfig *ServiceSpecDnsConfig `json:"dnsConfig,omitempty" tf:"dns_config"`
	// +optional
	HealthCheckConfig *ServiceSpecHealthCheckConfig `json:"healthCheckConfig,omitempty" tf:"health_check_config"`
	// +optional
	HealthCheckCustomConfig *ServiceSpecHealthCheckCustomConfig `json:"healthCheckCustomConfig,omitempty" tf:"health_check_custom_config"`
	Name                    *string                             `json:"name" tf:"name"`
	// +optional
	NamespaceID *string `json:"namespaceID,omitempty" tf:"namespace_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ServiceStatus struct {
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

// ServiceList is a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Service CRD objects
	Items []Service `json:"items,omitempty"`
}
