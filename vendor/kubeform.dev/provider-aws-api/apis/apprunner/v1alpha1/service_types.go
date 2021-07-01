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

type ServiceSpecEncryptionConfiguration struct {
	KmsKey *string `json:"kmsKey" tf:"kms_key"`
}

type ServiceSpecHealthCheckConfiguration struct {
	// +optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold"`
	// +optional
	Interval *int64 `json:"interval,omitempty" tf:"interval"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// +optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold"`
}

type ServiceSpecInstanceConfiguration struct {
	// +optional
	Cpu             *string `json:"cpu,omitempty" tf:"cpu"`
	InstanceRoleArn *string `json:"instanceRoleArn" tf:"instance_role_arn"`
	// +optional
	Memory *string `json:"memory,omitempty" tf:"memory"`
}

type ServiceSpecSourceConfigurationAuthenticationConfiguration struct {
	// +optional
	AccessRoleArn *string `json:"accessRoleArn,omitempty" tf:"access_role_arn"`
	// +optional
	ConnectionArn *string `json:"connectionArn,omitempty" tf:"connection_arn"`
}

type ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues struct {
	// +optional
	BuildCommand *string `json:"buildCommand,omitempty" tf:"build_command"`
	// +optional
	Port    *string `json:"port,omitempty" tf:"port"`
	Runtime *string `json:"runtime" tf:"runtime"`
	// +optional
	RuntimeEnvironmentVariables *map[string]string `json:"runtimeEnvironmentVariables,omitempty" tf:"runtime_environment_variables"`
	// +optional
	StartCommand *string `json:"startCommand,omitempty" tf:"start_command"`
}

type ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration struct {
	// +optional
	CodeConfigurationValues *ServiceSpecSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues `json:"codeConfigurationValues,omitempty" tf:"code_configuration_values"`
	ConfigurationSource     *string                                                                               `json:"configurationSource" tf:"configuration_source"`
}

type ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion struct {
	Type  *string `json:"type" tf:"type"`
	Value *string `json:"value" tf:"value"`
}

type ServiceSpecSourceConfigurationCodeRepository struct {
	// +optional
	CodeConfiguration *ServiceSpecSourceConfigurationCodeRepositoryCodeConfiguration `json:"codeConfiguration,omitempty" tf:"code_configuration"`
	RepositoryURL     *string                                                        `json:"repositoryURL" tf:"repository_url"`
	SourceCodeVersion *ServiceSpecSourceConfigurationCodeRepositorySourceCodeVersion `json:"sourceCodeVersion" tf:"source_code_version"`
}

type ServiceSpecSourceConfigurationImageRepositoryImageConfiguration struct {
	// +optional
	Port *string `json:"port,omitempty" tf:"port"`
	// +optional
	RuntimeEnvironmentVariables *map[string]string `json:"runtimeEnvironmentVariables,omitempty" tf:"runtime_environment_variables"`
	// +optional
	StartCommand *string `json:"startCommand,omitempty" tf:"start_command"`
}

type ServiceSpecSourceConfigurationImageRepository struct {
	// +optional
	ImageConfiguration  *ServiceSpecSourceConfigurationImageRepositoryImageConfiguration `json:"imageConfiguration,omitempty" tf:"image_configuration"`
	ImageIdentifier     *string                                                          `json:"imageIdentifier" tf:"image_identifier"`
	ImageRepositoryType *string                                                          `json:"imageRepositoryType" tf:"image_repository_type"`
}

type ServiceSpecSourceConfiguration struct {
	// +optional
	AuthenticationConfiguration *ServiceSpecSourceConfigurationAuthenticationConfiguration `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration"`
	// +optional
	AutoDeploymentsEnabled *bool `json:"autoDeploymentsEnabled,omitempty" tf:"auto_deployments_enabled"`
	// +optional
	CodeRepository *ServiceSpecSourceConfigurationCodeRepository `json:"codeRepository,omitempty" tf:"code_repository"`
	// +optional
	ImageRepository *ServiceSpecSourceConfigurationImageRepository `json:"imageRepository,omitempty" tf:"image_repository"`
}

type ServiceSpec struct {
	KubeformOutput *ServiceSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ServiceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutoScalingConfigurationArn *string `json:"autoScalingConfigurationArn,omitempty" tf:"auto_scaling_configuration_arn"`
	// +optional
	EncryptionConfiguration *ServiceSpecEncryptionConfiguration `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration"`
	// +optional
	HealthCheckConfiguration *ServiceSpecHealthCheckConfiguration `json:"healthCheckConfiguration,omitempty" tf:"health_check_configuration"`
	// +optional
	InstanceConfiguration *ServiceSpecInstanceConfiguration `json:"instanceConfiguration,omitempty" tf:"instance_configuration"`
	// +optional
	ServiceID   *string `json:"serviceID,omitempty" tf:"service_id"`
	ServiceName *string `json:"serviceName" tf:"service_name"`
	// +optional
	ServiceURL          *string                         `json:"serviceURL,omitempty" tf:"service_url"`
	SourceConfiguration *ServiceSpecSourceConfiguration `json:"sourceConfiguration" tf:"source_configuration"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
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