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

type Fleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FleetSpec   `json:"spec,omitempty"`
	Status            FleetStatus `json:"status,omitempty"`
}

type FleetSpecEc2InboundPermission struct {
	FromPort *int64  `json:"fromPort" tf:"from_port"`
	IpRange  *string `json:"ipRange" tf:"ip_range"`
	Protocol *string `json:"protocol" tf:"protocol"`
	ToPort   *int64  `json:"toPort" tf:"to_port"`
}

type FleetSpecResourceCreationLimitPolicy struct {
	// +optional
	NewGameSessionsPerCreator *int64 `json:"newGameSessionsPerCreator,omitempty" tf:"new_game_sessions_per_creator"`
	// +optional
	PolicyPeriodInMinutes *int64 `json:"policyPeriodInMinutes,omitempty" tf:"policy_period_in_minutes"`
}

type FleetSpecRuntimeConfigurationServerProcess struct {
	ConcurrentExecutions *int64  `json:"concurrentExecutions" tf:"concurrent_executions"`
	LaunchPath           *string `json:"launchPath" tf:"launch_path"`
	// +optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters"`
}

type FleetSpecRuntimeConfiguration struct {
	// +optional
	GameSessionActivationTimeoutSeconds *int64 `json:"gameSessionActivationTimeoutSeconds,omitempty" tf:"game_session_activation_timeout_seconds"`
	// +optional
	MaxConcurrentGameSessionActivations *int64 `json:"maxConcurrentGameSessionActivations,omitempty" tf:"max_concurrent_game_session_activations"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	ServerProcess []FleetSpecRuntimeConfigurationServerProcess `json:"serverProcess,omitempty" tf:"server_process"`
}

type FleetSpec struct {
	KubeformOutput *FleetSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource FleetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FleetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn     *string `json:"arn,omitempty" tf:"arn"`
	BuildID *string `json:"buildID" tf:"build_id"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	Ec2InboundPermission []FleetSpecEc2InboundPermission `json:"ec2InboundPermission,omitempty" tf:"ec2_inbound_permission"`
	Ec2InstanceType      *string                         `json:"ec2InstanceType" tf:"ec2_instance_type"`
	// +optional
	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type"`
	// +optional
	InstanceRoleArn *string `json:"instanceRoleArn,omitempty" tf:"instance_role_arn"`
	// +optional
	LogPaths []string `json:"logPaths,omitempty" tf:"log_paths"`
	// +optional
	MetricGroups []string `json:"metricGroups,omitempty" tf:"metric_groups"`
	Name         *string  `json:"name" tf:"name"`
	// +optional
	NewGameSessionProtectionPolicy *string `json:"newGameSessionProtectionPolicy,omitempty" tf:"new_game_session_protection_policy"`
	// +optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system"`
	// +optional
	ResourceCreationLimitPolicy *FleetSpecResourceCreationLimitPolicy `json:"resourceCreationLimitPolicy,omitempty" tf:"resource_creation_limit_policy"`
	// +optional
	RuntimeConfiguration *FleetSpecRuntimeConfiguration `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type FleetStatus struct {
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

// FleetList is a list of Fleets
type FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Fleet CRD objects
	Items []Fleet `json:"items,omitempty"`
}