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

type DomainAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainAssociationSpec   `json:"spec,omitempty"`
	Status            DomainAssociationStatus `json:"status,omitempty"`
}

type DomainAssociationSpecSubDomain struct {
	BranchName *string `json:"branchName" tf:"branch_name"`
	// +optional
	DnsRecord *string `json:"dnsRecord,omitempty" tf:"dns_record"`
	Prefix    *string `json:"prefix" tf:"prefix"`
	// +optional
	Verified *bool `json:"verified,omitempty" tf:"verified"`
}

type DomainAssociationSpec struct {
	KubeformOutput *DomainAssociationSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource DomainAssociationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DomainAssociationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AppID *string `json:"appID" tf:"app_id"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CertificateVerificationDNSRecord *string                          `json:"certificateVerificationDNSRecord,omitempty" tf:"certificate_verification_dns_record"`
	DomainName                       *string                          `json:"domainName" tf:"domain_name"`
	SubDomain                        []DomainAssociationSpecSubDomain `json:"subDomain" tf:"sub_domain"`
	// +optional
	WaitForVerification *bool `json:"waitForVerification,omitempty" tf:"wait_for_verification"`
}

type DomainAssociationStatus struct {
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

// DomainAssociationList is a list of DomainAssociations
type DomainAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DomainAssociation CRD objects
	Items []DomainAssociation `json:"items,omitempty"`
}
