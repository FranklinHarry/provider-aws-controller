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

type Alb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbSpec   `json:"spec,omitempty"`
	Status            AlbStatus `json:"status,omitempty"`
}

type AlbSpecAccessLogs struct {
	Bucket *string `json:"bucket" tf:"bucket"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type AlbSpecSubnetMapping struct {
	// +optional
	AllocationID *string `json:"allocationID,omitempty" tf:"allocation_id"`
	// +optional
	Ipv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address"`
	// +optional
	OutpostID *string `json:"outpostID,omitempty" tf:"outpost_id"`
	// +optional
	PrivateIpv4Address *string `json:"privateIpv4Address,omitempty" tf:"private_ipv4_address"`
	SubnetID           *string `json:"subnetID" tf:"subnet_id"`
}

type AlbSpec struct {
	KubeformOutput *AlbSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource AlbSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AlbSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessLogs *AlbSpecAccessLogs `json:"accessLogs,omitempty" tf:"access_logs"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix"`
	// +optional
	CustomerOwnedIpv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool"`
	// +optional
	DnsName *string `json:"dnsName,omitempty" tf:"dns_name"`
	// +optional
	DropInvalidHeaderFields *bool `json:"dropInvalidHeaderFields,omitempty" tf:"drop_invalid_header_fields"`
	// +optional
	EnableCrossZoneLoadBalancing *bool `json:"enableCrossZoneLoadBalancing,omitempty" tf:"enable_cross_zone_load_balancing"`
	// +optional
	EnableDeletionProtection *bool `json:"enableDeletionProtection,omitempty" tf:"enable_deletion_protection"`
	// +optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2"`
	// +optional
	IdleTimeout *int64 `json:"idleTimeout,omitempty" tf:"idle_timeout"`
	// +optional
	Internal *bool `json:"internal,omitempty" tf:"internal"`
	// +optional
	IpAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type"`
	// +optional
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	SubnetMapping []AlbSpecSubnetMapping `json:"subnetMapping,omitempty" tf:"subnet_mapping"`
	// +optional
	Subnets []string `json:"subnets,omitempty" tf:"subnets"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
}

type AlbStatus struct {
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

// AlbList is a list of Albs
type AlbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Alb CRD objects
	Items []Alb `json:"items,omitempty"`
}