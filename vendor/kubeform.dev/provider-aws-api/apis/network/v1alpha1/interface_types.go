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

type Interface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InterfaceSpec   `json:"spec,omitempty"`
	Status            InterfaceStatus `json:"status,omitempty"`
}

type InterfaceSpecAttachment struct {
	// +optional
	AttachmentID *string `json:"attachmentID,omitempty" tf:"attachment_id"`
	DeviceIndex  *int64  `json:"deviceIndex" tf:"device_index"`
	Instance     *string `json:"instance" tf:"instance"`
}

type InterfaceSpec struct {
	KubeformOutput *InterfaceSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource InterfaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InterfaceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Attachment []InterfaceSpecAttachment `json:"attachment,omitempty" tf:"attachment"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type"`
	// +optional
	Ipv6AddressCount *int64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count"`
	// +optional
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses"`
	// +optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address"`
	// +optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn"`
	// +optional
	PrivateDNSName *string `json:"privateDNSName,omitempty" tf:"private_dns_name"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	PrivateIPS []string `json:"privateIPS,omitempty" tf:"private_ips"`
	// +optional
	PrivateIPSCount *int64 `json:"privateIPSCount,omitempty" tf:"private_ips_count"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	SourceDestCheck *bool   `json:"sourceDestCheck,omitempty" tf:"source_dest_check"`
	SubnetID        *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type InterfaceStatus struct {
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

// InterfaceList is a list of Interfaces
type InterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Interface CRD objects
	Items []Interface `json:"items,omitempty"`
}