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

type PrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            PrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

type PrivateVirtualInterfaceSpec struct {
	KubeformOutput *PrivateVirtualInterfaceSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource PrivateVirtualInterfaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PrivateVirtualInterfaceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressFamily *string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`
	// +optional
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device"`
	BgpAsn    *int64  `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey   *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`
	ConnectionID *string `json:"connectionID" tf:"connection_id"`
	// +optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`
	// +optional
	DxGatewayID *string `json:"dxGatewayID,omitempty" tf:"dx_gateway_id"`
	// +optional
	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable"`
	// +optional
	Mtu  *int64  `json:"mtu,omitempty" tf:"mtu"`
	Name *string `json:"name" tf:"name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	Vlan    *int64             `json:"vlan" tf:"vlan"`
	// +optional
	VpnGatewayID *string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id"`
}

type PrivateVirtualInterfaceStatus struct {
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

// PrivateVirtualInterfaceList is a list of PrivateVirtualInterfaces
type PrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PrivateVirtualInterface CRD objects
	Items []PrivateVirtualInterface `json:"items,omitempty"`
}
