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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecRestoreToPointInTime struct {
	// +optional
	RestoreToTime *string `json:"restoreToTime,omitempty" tf:"restore_to_time"`
	// +optional
	RestoreType             *string `json:"restoreType,omitempty" tf:"restore_type"`
	SourceClusterIdentifier *string `json:"sourceClusterIdentifier" tf:"source_cluster_identifier"`
	// +optional
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time"`
}

type ClusterSpecS3Import struct {
	BucketName *string `json:"bucketName" tf:"bucket_name"`
	// +optional
	BucketPrefix        *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix"`
	IngestionRole       *string `json:"ingestionRole" tf:"ingestion_role"`
	SourceEngine        *string `json:"sourceEngine" tf:"source_engine"`
	SourceEngineVersion *string `json:"sourceEngineVersion" tf:"source_engine_version"`
}

type ClusterSpecScalingConfiguration struct {
	// +optional
	AutoPause *bool `json:"autoPause,omitempty" tf:"auto_pause"`
	// +optional
	MaxCapacity *int64 `json:"maxCapacity,omitempty" tf:"max_capacity"`
	// +optional
	MinCapacity *int64 `json:"minCapacity,omitempty" tf:"min_capacity"`
	// +optional
	SecondsUntilAutoPause *int64 `json:"secondsUntilAutoPause,omitempty" tf:"seconds_until_auto_pause"`
	// +optional
	TimeoutAction *string `json:"timeoutAction,omitempty" tf:"timeout_action"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade"`
	// +optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`
	// +optional
	BacktrackWindow *int64 `json:"backtrackWindow,omitempty" tf:"backtrack_window"`
	// +optional
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period"`
	// +optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier"`
	// +optional
	ClusterIdentifierPrefix *string `json:"clusterIdentifierPrefix,omitempty" tf:"cluster_identifier_prefix"`
	// +optional
	ClusterMembers []string `json:"clusterMembers,omitempty" tf:"cluster_members"`
	// +optional
	ClusterResourceID *string `json:"clusterResourceID,omitempty" tf:"cluster_resource_id"`
	// +optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot"`
	// +optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name"`
	// +optional
	DbClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name"`
	// +optional
	DbSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name"`
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// +optional
	EnableHTTPEndpoint *bool `json:"enableHTTPEndpoint,omitempty" tf:"enable_http_endpoint"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	EngineMode *string `json:"engineMode,omitempty" tf:"engine_mode"`
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// +optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier"`
	// +optional
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier"`
	// +optional
	HostedZoneID *string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id"`
	// +optional
	IamDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled"`
	// +optional
	IamRoles []string `json:"iamRoles,omitempty" tf:"iam_roles"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	MasterPassword *string `json:"-" sensitive:"true" tf:"master_password"`
	// +optional
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window"`
	// +optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window"`
	// +optional
	ReaderEndpoint *string `json:"readerEndpoint,omitempty" tf:"reader_endpoint"`
	// +optional
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier"`
	// +optional
	RestoreToPointInTime *ClusterSpecRestoreToPointInTime `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time"`
	// +optional
	S3Import *ClusterSpecS3Import `json:"s3Import,omitempty" tf:"s3_import"`
	// +optional
	ScalingConfiguration *ClusterSpecScalingConfiguration `json:"scalingConfiguration,omitempty" tf:"scaling_configuration"`
	// +optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot"`
	// +optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier"`
	// +optional
	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region"`
	// +optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
