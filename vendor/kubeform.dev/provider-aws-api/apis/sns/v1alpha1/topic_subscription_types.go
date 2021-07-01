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

type TopicSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSubscriptionSpec   `json:"spec,omitempty"`
	Status            TopicSubscriptionStatus `json:"status,omitempty"`
}

type TopicSubscriptionSpec struct {
	KubeformOutput *TopicSubscriptionSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource TopicSubscriptionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TopicSubscriptionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ConfirmationTimeoutInMinutes *int64 `json:"confirmationTimeoutInMinutes,omitempty" tf:"confirmation_timeout_in_minutes"`
	// +optional
	ConfirmationWasAuthenticated *bool `json:"confirmationWasAuthenticated,omitempty" tf:"confirmation_was_authenticated"`
	// +optional
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy"`
	Endpoint       *string `json:"endpoint" tf:"endpoint"`
	// +optional
	EndpointAutoConfirms *bool `json:"endpointAutoConfirms,omitempty" tf:"endpoint_auto_confirms"`
	// +optional
	FilterPolicy *string `json:"filterPolicy,omitempty" tf:"filter_policy"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	PendingConfirmation *bool   `json:"pendingConfirmation,omitempty" tf:"pending_confirmation"`
	Protocol            *string `json:"protocol" tf:"protocol"`
	// +optional
	RawMessageDelivery *bool `json:"rawMessageDelivery,omitempty" tf:"raw_message_delivery"`
	// +optional
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy"`
	// +optional
	SubscriptionRoleArn *string `json:"subscriptionRoleArn,omitempty" tf:"subscription_role_arn"`
	TopicArn            *string `json:"topicArn" tf:"topic_arn"`
}

type TopicSubscriptionStatus struct {
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

// TopicSubscriptionList is a list of TopicSubscriptions
type TopicSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TopicSubscription CRD objects
	Items []TopicSubscription `json:"items,omitempty"`
}
