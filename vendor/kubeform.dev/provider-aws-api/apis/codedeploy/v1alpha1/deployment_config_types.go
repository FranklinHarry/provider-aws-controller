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

type DeploymentConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentConfigSpec   `json:"spec,omitempty"`
	Status            DeploymentConfigStatus `json:"status,omitempty"`
}

type DeploymentConfigSpecMinimumHealthyHosts struct {
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Value *int64 `json:"value,omitempty" tf:"value"`
}

type DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary struct {
	// +optional
	Interval *int64 `json:"interval,omitempty" tf:"interval"`
	// +optional
	Percentage *int64 `json:"percentage,omitempty" tf:"percentage"`
}

type DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear struct {
	// +optional
	Interval *int64 `json:"interval,omitempty" tf:"interval"`
	// +optional
	Percentage *int64 `json:"percentage,omitempty" tf:"percentage"`
}

type DeploymentConfigSpecTrafficRoutingConfig struct {
	// +optional
	TimeBasedCanary *DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary `json:"timeBasedCanary,omitempty" tf:"time_based_canary"`
	// +optional
	TimeBasedLinear *DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear `json:"timeBasedLinear,omitempty" tf:"time_based_linear"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type DeploymentConfigSpec struct {
	KubeformOutput *DeploymentConfigSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource DeploymentConfigSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DeploymentConfigSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform"`
	// +optional
	DeploymentConfigID   *string `json:"deploymentConfigID,omitempty" tf:"deployment_config_id"`
	DeploymentConfigName *string `json:"deploymentConfigName" tf:"deployment_config_name"`
	// +optional
	MinimumHealthyHosts *DeploymentConfigSpecMinimumHealthyHosts `json:"minimumHealthyHosts,omitempty" tf:"minimum_healthy_hosts"`
	// +optional
	TrafficRoutingConfig *DeploymentConfigSpecTrafficRoutingConfig `json:"trafficRoutingConfig,omitempty" tf:"traffic_routing_config"`
}

type DeploymentConfigStatus struct {
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

// DeploymentConfigList is a list of DeploymentConfigs
type DeploymentConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DeploymentConfig CRD objects
	Items []DeploymentConfig `json:"items,omitempty"`
}