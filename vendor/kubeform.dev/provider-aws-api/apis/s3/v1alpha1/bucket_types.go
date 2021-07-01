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

type Bucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec,omitempty"`
	Status            BucketStatus `json:"status,omitempty"`
}

type BucketSpecCorsRule struct {
	// +optional
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers"`
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	ExposeHeaders []string `json:"exposeHeaders,omitempty" tf:"expose_headers"`
	// +optional
	MaxAgeSeconds *int64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds"`
}

type BucketSpecGrant struct {
	// +optional
	ID          *string  `json:"ID,omitempty" tf:"id"`
	Permissions []string `json:"permissions" tf:"permissions"`
	Type        *string  `json:"type" tf:"type"`
	// +optional
	Uri *string `json:"uri,omitempty" tf:"uri"`
}

type BucketSpecLifecycleRuleExpiration struct {
	// +optional
	Date *string `json:"date,omitempty" tf:"date"`
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
	// +optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker"`
}

type BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
}

type BucketSpecLifecycleRuleNoncurrentVersionTransition struct {
	// +optional
	Days         *int64  `json:"days,omitempty" tf:"days"`
	StorageClass *string `json:"storageClass" tf:"storage_class"`
}

type BucketSpecLifecycleRuleTransition struct {
	// +optional
	Date *string `json:"date,omitempty" tf:"date"`
	// +optional
	Days         *int64  `json:"days,omitempty" tf:"days"`
	StorageClass *string `json:"storageClass" tf:"storage_class"`
}

type BucketSpecLifecycleRule struct {
	// +optional
	AbortIncompleteMultipartUploadDays *int64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days"`
	Enabled                            *bool  `json:"enabled" tf:"enabled"`
	// +optional
	Expiration *BucketSpecLifecycleRuleExpiration `json:"expiration,omitempty" tf:"expiration"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	NoncurrentVersionExpiration *BucketSpecLifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration"`
	// +optional
	NoncurrentVersionTransition []BucketSpecLifecycleRuleNoncurrentVersionTransition `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Transition []BucketSpecLifecycleRuleTransition `json:"transition,omitempty" tf:"transition"`
}

type BucketSpecLogging struct {
	TargetBucket *string `json:"targetBucket" tf:"target_bucket"`
	// +optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix"`
}

type BucketSpecObjectLockConfigurationRuleDefaultRetention struct {
	// +optional
	Days *int64  `json:"days,omitempty" tf:"days"`
	Mode *string `json:"mode" tf:"mode"`
	// +optional
	Years *int64 `json:"years,omitempty" tf:"years"`
}

type BucketSpecObjectLockConfigurationRule struct {
	DefaultRetention *BucketSpecObjectLockConfigurationRuleDefaultRetention `json:"defaultRetention" tf:"default_retention"`
}

type BucketSpecObjectLockConfiguration struct {
	ObjectLockEnabled *string `json:"objectLockEnabled" tf:"object_lock_enabled"`
	// +optional
	Rule *BucketSpecObjectLockConfigurationRule `json:"rule,omitempty" tf:"rule"`
}

type BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation struct {
	Owner *string `json:"owner" tf:"owner"`
}

type BucketSpecReplicationConfigurationRulesDestination struct {
	// +optional
	AccessControlTranslation *BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation `json:"accessControlTranslation,omitempty" tf:"access_control_translation"`
	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	Bucket    *string `json:"bucket" tf:"bucket"`
	// +optional
	ReplicaKmsKeyID *string `json:"replicaKmsKeyID,omitempty" tf:"replica_kms_key_id"`
	// +optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`
}

type BucketSpecReplicationConfigurationRulesFilter struct {
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type BucketSpecReplicationConfigurationRulesSourceSelectionCriteria struct {
	// +optional
	SseKmsEncryptedObjects *BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects"`
}

type BucketSpecReplicationConfigurationRules struct {
	Destination *BucketSpecReplicationConfigurationRulesDestination `json:"destination" tf:"destination"`
	// +optional
	Filter *BucketSpecReplicationConfigurationRulesFilter `json:"filter,omitempty" tf:"filter"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	SourceSelectionCriteria *BucketSpecReplicationConfigurationRulesSourceSelectionCriteria `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria"`
	Status                  *string                                                         `json:"status" tf:"status"`
}

type BucketSpecReplicationConfiguration struct {
	Role  *string                                   `json:"role" tf:"role"`
	Rules []BucketSpecReplicationConfigurationRules `json:"rules" tf:"rules"`
}

type BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	// +optional
	KmsMasterKeyID *string `json:"kmsMasterKeyID,omitempty" tf:"kms_master_key_id"`
	SseAlgorithm   *string `json:"sseAlgorithm" tf:"sse_algorithm"`
}

type BucketSpecServerSideEncryptionConfigurationRule struct {
	ApplyServerSideEncryptionByDefault *BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `json:"applyServerSideEncryptionByDefault" tf:"apply_server_side_encryption_by_default"`
	// +optional
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled"`
}

type BucketSpecServerSideEncryptionConfiguration struct {
	Rule *BucketSpecServerSideEncryptionConfigurationRule `json:"rule" tf:"rule"`
}

type BucketSpecVersioning struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete"`
}

type BucketSpecWebsite struct {
	// +optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document"`
	// +optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document"`
	// +optional
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to"`
	// +optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules"`
}

type BucketSpec struct {
	KubeformOutput *BucketSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource BucketSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BucketSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccelerationStatus *string `json:"accelerationStatus,omitempty" tf:"acceleration_status"`
	// +optional
	Acl *string `json:"acl,omitempty" tf:"acl"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`
	// +optional
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name"`
	// +optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix"`
	// +optional
	BucketRegionalDomainName *string `json:"bucketRegionalDomainName,omitempty" tf:"bucket_regional_domain_name"`
	// +optional
	CorsRule []BucketSpecCorsRule `json:"corsRule,omitempty" tf:"cors_rule"`
	// +optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`
	// +optional
	Grant []BucketSpecGrant `json:"grant,omitempty" tf:"grant"`
	// +optional
	HostedZoneID *string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id"`
	// +optional
	LifecycleRule []BucketSpecLifecycleRule `json:"lifecycleRule,omitempty" tf:"lifecycle_rule"`
	// +optional
	Logging []BucketSpecLogging `json:"logging,omitempty" tf:"logging"`
	// +optional
	ObjectLockConfiguration *BucketSpecObjectLockConfiguration `json:"objectLockConfiguration,omitempty" tf:"object_lock_configuration"`
	// +optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	ReplicationConfiguration *BucketSpecReplicationConfiguration `json:"replicationConfiguration,omitempty" tf:"replication_configuration"`
	// +optional
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer"`
	// +optional
	ServerSideEncryptionConfiguration *BucketSpecServerSideEncryptionConfiguration `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Versioning *BucketSpecVersioning `json:"versioning,omitempty" tf:"versioning"`
	// +optional
	Website *BucketSpecWebsite `json:"website,omitempty" tf:"website"`
	// +optional
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain"`
	// +optional
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint"`
}

type BucketStatus struct {
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

// BucketList is a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Bucket CRD objects
	Items []Bucket `json:"items,omitempty"`
}
