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

type ReceiptRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReceiptRuleSpec   `json:"spec,omitempty"`
	Status            ReceiptRuleStatus `json:"status,omitempty"`
}

type ReceiptRuleSpecAddHeaderAction struct {
	HeaderName  *string `json:"headerName" tf:"header_name"`
	HeaderValue *string `json:"headerValue" tf:"header_value"`
	Position    *int64  `json:"position" tf:"position"`
}

type ReceiptRuleSpecBounceAction struct {
	Message       *string `json:"message" tf:"message"`
	Position      *int64  `json:"position" tf:"position"`
	Sender        *string `json:"sender" tf:"sender"`
	SmtpReplyCode *string `json:"smtpReplyCode" tf:"smtp_reply_code"`
	// +optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code"`
	// +optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type ReceiptRuleSpecLambdaAction struct {
	FunctionArn *string `json:"functionArn" tf:"function_arn"`
	// +optional
	InvocationType *string `json:"invocationType,omitempty" tf:"invocation_type"`
	Position       *int64  `json:"position" tf:"position"`
	// +optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type ReceiptRuleSpecS3Action struct {
	BucketName *string `json:"bucketName" tf:"bucket_name"`
	// +optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	// +optional
	ObjectKeyPrefix *string `json:"objectKeyPrefix,omitempty" tf:"object_key_prefix"`
	Position        *int64  `json:"position" tf:"position"`
	// +optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type ReceiptRuleSpecSnsAction struct {
	// +optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`
	Position *int64  `json:"position" tf:"position"`
	TopicArn *string `json:"topicArn" tf:"topic_arn"`
}

type ReceiptRuleSpecStopAction struct {
	Position *int64  `json:"position" tf:"position"`
	Scope    *string `json:"scope" tf:"scope"`
	// +optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type ReceiptRuleSpecWorkmailAction struct {
	OrganizationArn *string `json:"organizationArn" tf:"organization_arn"`
	Position        *int64  `json:"position" tf:"position"`
	// +optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn"`
}

type ReceiptRuleSpec struct {
	KubeformOutput *ReceiptRuleSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ReceiptRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ReceiptRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AddHeaderAction []ReceiptRuleSpecAddHeaderAction `json:"addHeaderAction,omitempty" tf:"add_header_action"`
	// +optional
	After *string `json:"after,omitempty" tf:"after"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	BounceAction []ReceiptRuleSpecBounceAction `json:"bounceAction,omitempty" tf:"bounce_action"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	LambdaAction []ReceiptRuleSpecLambdaAction `json:"lambdaAction,omitempty" tf:"lambda_action"`
	Name         *string                       `json:"name" tf:"name"`
	// +optional
	Recipients  []string `json:"recipients,omitempty" tf:"recipients"`
	RuleSetName *string  `json:"ruleSetName" tf:"rule_set_name"`
	// +optional
	S3Action []ReceiptRuleSpecS3Action `json:"s3Action,omitempty" tf:"s3_action"`
	// +optional
	ScanEnabled *bool `json:"scanEnabled,omitempty" tf:"scan_enabled"`
	// +optional
	SnsAction []ReceiptRuleSpecSnsAction `json:"snsAction,omitempty" tf:"sns_action"`
	// +optional
	StopAction []ReceiptRuleSpecStopAction `json:"stopAction,omitempty" tf:"stop_action"`
	// +optional
	TlsPolicy *string `json:"tlsPolicy,omitempty" tf:"tls_policy"`
	// +optional
	WorkmailAction []ReceiptRuleSpecWorkmailAction `json:"workmailAction,omitempty" tf:"workmail_action"`
}

type ReceiptRuleStatus struct {
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

// ReceiptRuleList is a list of ReceiptRules
type ReceiptRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ReceiptRule CRD objects
	Items []ReceiptRule `json:"items,omitempty"`
}