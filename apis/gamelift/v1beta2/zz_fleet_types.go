// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CertificateConfigurationInitParameters struct {

	// Indicates whether a TLS/SSL certificate is generated for a fleet. Valid values are DISABLED and GENERATED. Default value is DISABLED.
	CertificateType *string `json:"certificateType,omitempty" tf:"certificate_type,omitempty"`
}

type CertificateConfigurationObservation struct {

	// Indicates whether a TLS/SSL certificate is generated for a fleet. Valid values are DISABLED and GENERATED. Default value is DISABLED.
	CertificateType *string `json:"certificateType,omitempty" tf:"certificate_type,omitempty"`
}

type CertificateConfigurationParameters struct {

	// Indicates whether a TLS/SSL certificate is generated for a fleet. Valid values are DISABLED and GENERATED. Default value is DISABLED.
	// +kubebuilder:validation:Optional
	CertificateType *string `json:"certificateType,omitempty" tf:"certificate_type,omitempty"`
}

type EC2InboundPermissionInitParameters struct {

	// Starting value for a range of allowed port numbers.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Range of allowed IP addresses expressed in CIDR notationE.g., 000.000.000.000/[subnet mask] or 0.0.0.0/[subnet mask].
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// Network communication protocol used by the fleetE.g., TCP or UDP
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than from_port.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type EC2InboundPermissionObservation struct {

	// Starting value for a range of allowed port numbers.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Range of allowed IP addresses expressed in CIDR notationE.g., 000.000.000.000/[subnet mask] or 0.0.0.0/[subnet mask].
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// Network communication protocol used by the fleetE.g., TCP or UDP
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than from_port.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type EC2InboundPermissionParameters struct {

	// Starting value for a range of allowed port numbers.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// Range of allowed IP addresses expressed in CIDR notationE.g., 000.000.000.000/[subnet mask] or 0.0.0.0/[subnet mask].
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange" tf:"ip_range,omitempty"`

	// Network communication protocol used by the fleetE.g., TCP or UDP
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than from_port.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort" tf:"to_port,omitempty"`
}

type FleetInitParameters struct {

	// ID of the GameLift Build to be deployed on the fleet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/gamelift/v1beta2.Build
	BuildID *string `json:"buildId,omitempty" tf:"build_id,omitempty"`

	// Reference to a Build in gamelift to populate buildId.
	// +kubebuilder:validation:Optional
	BuildIDRef *v1.Reference `json:"buildIdRef,omitempty" tf:"-"`

	// Selector for a Build in gamelift to populate buildId.
	// +kubebuilder:validation:Optional
	BuildIDSelector *v1.Selector `json:"buildIdSelector,omitempty" tf:"-"`

	// Prompts GameLift to generate a TLS/SSL certificate for the fleet. See certificate_configuration.
	CertificateConfiguration *CertificateConfigurationInitParameters `json:"certificateConfiguration,omitempty" tf:"certificate_configuration,omitempty"`

	// Human-readable description of the fleet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	EC2InboundPermission []EC2InboundPermissionInitParameters `json:"ec2InboundPermission,omitempty" tf:"ec2_inbound_permission,omitempty"`

	// Name of an EC2 instance typeE.g., t2.micro
	EC2InstanceType *string `json:"ec2InstanceType,omitempty" tf:"ec2_instance_type,omitempty"`

	// Type of fleet. This value must be ON_DEMAND or SPOT. Defaults to ON_DEMAND.
	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`

	// ARN of an IAM role that instances in the fleet can assume.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	InstanceRoleArn *string `json:"instanceRoleArn,omitempty" tf:"instance_role_arn,omitempty"`

	// Reference to a Role in iam to populate instanceRoleArn.
	// +kubebuilder:validation:Optional
	InstanceRoleArnRef *v1.Reference `json:"instanceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate instanceRoleArn.
	// +kubebuilder:validation:Optional
	InstanceRoleArnSelector *v1.Selector `json:"instanceRoleArnSelector,omitempty" tf:"-"`

	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to default.
	MetricGroups []*string `json:"metricGroups,omitempty" tf:"metric_groups,omitempty"`

	// The name of the fleet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Game session protection policy to apply to all instances in this fleetE.g., FullProtection. Defaults to NoProtection.
	NewGameSessionProtectionPolicy *string `json:"newGameSessionProtectionPolicy,omitempty" tf:"new_game_session_protection_policy,omitempty"`

	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicyInitParameters `json:"resourceCreationLimitPolicy,omitempty" tf:"resource_creation_limit_policy,omitempty"`

	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration *RuntimeConfigurationInitParameters `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration,omitempty"`

	// ID of the GameLift Script to be deployed on the fleet.
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FleetObservation struct {

	// Fleet ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Build ARN.
	BuildArn *string `json:"buildArn,omitempty" tf:"build_arn,omitempty"`

	// ID of the GameLift Build to be deployed on the fleet.
	BuildID *string `json:"buildId,omitempty" tf:"build_id,omitempty"`

	// Prompts GameLift to generate a TLS/SSL certificate for the fleet. See certificate_configuration.
	CertificateConfiguration *CertificateConfigurationObservation `json:"certificateConfiguration,omitempty" tf:"certificate_configuration,omitempty"`

	// Human-readable description of the fleet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	EC2InboundPermission []EC2InboundPermissionObservation `json:"ec2InboundPermission,omitempty" tf:"ec2_inbound_permission,omitempty"`

	// Name of an EC2 instance typeE.g., t2.micro
	EC2InstanceType *string `json:"ec2InstanceType,omitempty" tf:"ec2_instance_type,omitempty"`

	// Type of fleet. This value must be ON_DEMAND or SPOT. Defaults to ON_DEMAND.
	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`

	// Fleet ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn *string `json:"instanceRoleArn,omitempty" tf:"instance_role_arn,omitempty"`

	LogPaths []*string `json:"logPaths,omitempty" tf:"log_paths,omitempty"`

	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to default.
	MetricGroups []*string `json:"metricGroups,omitempty" tf:"metric_groups,omitempty"`

	// The name of the fleet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Game session protection policy to apply to all instances in this fleetE.g., FullProtection. Defaults to NoProtection.
	NewGameSessionProtectionPolicy *string `json:"newGameSessionProtectionPolicy,omitempty" tf:"new_game_session_protection_policy,omitempty"`

	// Operating system of the fleet's computing resources.
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicyObservation `json:"resourceCreationLimitPolicy,omitempty" tf:"resource_creation_limit_policy,omitempty"`

	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration *RuntimeConfigurationObservation `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration,omitempty"`

	// Script ARN.
	ScriptArn *string `json:"scriptArn,omitempty" tf:"script_arn,omitempty"`

	// ID of the GameLift Script to be deployed on the fleet.
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FleetParameters struct {

	// ID of the GameLift Build to be deployed on the fleet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/gamelift/v1beta2.Build
	// +kubebuilder:validation:Optional
	BuildID *string `json:"buildId,omitempty" tf:"build_id,omitempty"`

	// Reference to a Build in gamelift to populate buildId.
	// +kubebuilder:validation:Optional
	BuildIDRef *v1.Reference `json:"buildIdRef,omitempty" tf:"-"`

	// Selector for a Build in gamelift to populate buildId.
	// +kubebuilder:validation:Optional
	BuildIDSelector *v1.Selector `json:"buildIdSelector,omitempty" tf:"-"`

	// Prompts GameLift to generate a TLS/SSL certificate for the fleet. See certificate_configuration.
	// +kubebuilder:validation:Optional
	CertificateConfiguration *CertificateConfigurationParameters `json:"certificateConfiguration,omitempty" tf:"certificate_configuration,omitempty"`

	// Human-readable description of the fleet.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	// +kubebuilder:validation:Optional
	EC2InboundPermission []EC2InboundPermissionParameters `json:"ec2InboundPermission,omitempty" tf:"ec2_inbound_permission,omitempty"`

	// Name of an EC2 instance typeE.g., t2.micro
	// +kubebuilder:validation:Optional
	EC2InstanceType *string `json:"ec2InstanceType,omitempty" tf:"ec2_instance_type,omitempty"`

	// Type of fleet. This value must be ON_DEMAND or SPOT. Defaults to ON_DEMAND.
	// +kubebuilder:validation:Optional
	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`

	// ARN of an IAM role that instances in the fleet can assume.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	InstanceRoleArn *string `json:"instanceRoleArn,omitempty" tf:"instance_role_arn,omitempty"`

	// Reference to a Role in iam to populate instanceRoleArn.
	// +kubebuilder:validation:Optional
	InstanceRoleArnRef *v1.Reference `json:"instanceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate instanceRoleArn.
	// +kubebuilder:validation:Optional
	InstanceRoleArnSelector *v1.Selector `json:"instanceRoleArnSelector,omitempty" tf:"-"`

	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to default.
	// +kubebuilder:validation:Optional
	MetricGroups []*string `json:"metricGroups,omitempty" tf:"metric_groups,omitempty"`

	// The name of the fleet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Game session protection policy to apply to all instances in this fleetE.g., FullProtection. Defaults to NoProtection.
	// +kubebuilder:validation:Optional
	NewGameSessionProtectionPolicy *string `json:"newGameSessionProtectionPolicy,omitempty" tf:"new_game_session_protection_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	// +kubebuilder:validation:Optional
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicyParameters `json:"resourceCreationLimitPolicy,omitempty" tf:"resource_creation_limit_policy,omitempty"`

	// Instructions for launching server processes on each instance in the fleet. See below.
	// +kubebuilder:validation:Optional
	RuntimeConfiguration *RuntimeConfigurationParameters `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration,omitempty"`

	// ID of the GameLift Script to be deployed on the fleet.
	// +kubebuilder:validation:Optional
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourceCreationLimitPolicyInitParameters struct {

	// Maximum number of game sessions that an individual can create during the policy period.
	NewGameSessionsPerCreator *float64 `json:"newGameSessionsPerCreator,omitempty" tf:"new_game_sessions_per_creator,omitempty"`

	// Time span used in evaluating the resource creation limit policy.
	PolicyPeriodInMinutes *float64 `json:"policyPeriodInMinutes,omitempty" tf:"policy_period_in_minutes,omitempty"`
}

type ResourceCreationLimitPolicyObservation struct {

	// Maximum number of game sessions that an individual can create during the policy period.
	NewGameSessionsPerCreator *float64 `json:"newGameSessionsPerCreator,omitempty" tf:"new_game_sessions_per_creator,omitempty"`

	// Time span used in evaluating the resource creation limit policy.
	PolicyPeriodInMinutes *float64 `json:"policyPeriodInMinutes,omitempty" tf:"policy_period_in_minutes,omitempty"`
}

type ResourceCreationLimitPolicyParameters struct {

	// Maximum number of game sessions that an individual can create during the policy period.
	// +kubebuilder:validation:Optional
	NewGameSessionsPerCreator *float64 `json:"newGameSessionsPerCreator,omitempty" tf:"new_game_sessions_per_creator,omitempty"`

	// Time span used in evaluating the resource creation limit policy.
	// +kubebuilder:validation:Optional
	PolicyPeriodInMinutes *float64 `json:"policyPeriodInMinutes,omitempty" tf:"policy_period_in_minutes,omitempty"`
}

type RuntimeConfigurationInitParameters struct {

	// Maximum amount of time (in seconds) that a game session can remain in status ACTIVATING.
	GameSessionActivationTimeoutSeconds *float64 `json:"gameSessionActivationTimeoutSeconds,omitempty" tf:"game_session_activation_timeout_seconds,omitempty"`

	// Maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously.
	MaxConcurrentGameSessionActivations *float64 `json:"maxConcurrentGameSessionActivations,omitempty" tf:"max_concurrent_game_session_activations,omitempty"`

	// Collection of server process configurations that describe which server processes to run on each instance in a fleet. See below.
	ServerProcess []ServerProcessInitParameters `json:"serverProcess,omitempty" tf:"server_process,omitempty"`
}

type RuntimeConfigurationObservation struct {

	// Maximum amount of time (in seconds) that a game session can remain in status ACTIVATING.
	GameSessionActivationTimeoutSeconds *float64 `json:"gameSessionActivationTimeoutSeconds,omitempty" tf:"game_session_activation_timeout_seconds,omitempty"`

	// Maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously.
	MaxConcurrentGameSessionActivations *float64 `json:"maxConcurrentGameSessionActivations,omitempty" tf:"max_concurrent_game_session_activations,omitempty"`

	// Collection of server process configurations that describe which server processes to run on each instance in a fleet. See below.
	ServerProcess []ServerProcessObservation `json:"serverProcess,omitempty" tf:"server_process,omitempty"`
}

type RuntimeConfigurationParameters struct {

	// Maximum amount of time (in seconds) that a game session can remain in status ACTIVATING.
	// +kubebuilder:validation:Optional
	GameSessionActivationTimeoutSeconds *float64 `json:"gameSessionActivationTimeoutSeconds,omitempty" tf:"game_session_activation_timeout_seconds,omitempty"`

	// Maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously.
	// +kubebuilder:validation:Optional
	MaxConcurrentGameSessionActivations *float64 `json:"maxConcurrentGameSessionActivations,omitempty" tf:"max_concurrent_game_session_activations,omitempty"`

	// Collection of server process configurations that describe which server processes to run on each instance in a fleet. See below.
	// +kubebuilder:validation:Optional
	ServerProcess []ServerProcessParameters `json:"serverProcess,omitempty" tf:"server_process,omitempty"`
}

type ServerProcessInitParameters struct {

	// Number of server processes using this configuration to run concurrently on an instance.
	ConcurrentExecutions *float64 `json:"concurrentExecutions,omitempty" tf:"concurrent_executions,omitempty"`

	// Location of the server executable in a game build. All game builds are installed on instances at the root : for Windows instances C:\game, and for Linux instances /local/game.
	LaunchPath *string `json:"launchPath,omitempty" tf:"launch_path,omitempty"`

	// Optional list of parameters to pass to the server executable on launch.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type ServerProcessObservation struct {

	// Number of server processes using this configuration to run concurrently on an instance.
	ConcurrentExecutions *float64 `json:"concurrentExecutions,omitempty" tf:"concurrent_executions,omitempty"`

	// Location of the server executable in a game build. All game builds are installed on instances at the root : for Windows instances C:\game, and for Linux instances /local/game.
	LaunchPath *string `json:"launchPath,omitempty" tf:"launch_path,omitempty"`

	// Optional list of parameters to pass to the server executable on launch.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type ServerProcessParameters struct {

	// Number of server processes using this configuration to run concurrently on an instance.
	// +kubebuilder:validation:Optional
	ConcurrentExecutions *float64 `json:"concurrentExecutions" tf:"concurrent_executions,omitempty"`

	// Location of the server executable in a game build. All game builds are installed on instances at the root : for Windows instances C:\game, and for Linux instances /local/game.
	// +kubebuilder:validation:Optional
	LaunchPath *string `json:"launchPath" tf:"launch_path,omitempty"`

	// Optional list of parameters to pass to the server executable on launch.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// FleetSpec defines the desired state of Fleet
type FleetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FleetParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FleetInitParameters `json:"initProvider,omitempty"`
}

// FleetStatus defines the observed state of Fleet.
type FleetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Fleet is the Schema for the Fleets API. Provides a GameLift Fleet resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws},path=fleet
type Fleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ec2InstanceType) || (has(self.initProvider) && has(self.initProvider.ec2InstanceType))",message="spec.forProvider.ec2InstanceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FleetSpec   `json:"spec"`
	Status FleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FleetList contains a list of Fleets
type FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Fleet `json:"items"`
}

// Repository type metadata.
var (
	Fleet_Kind             = "Fleet"
	Fleet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Fleet_Kind}.String()
	Fleet_KindAPIVersion   = Fleet_Kind + "." + CRDGroupVersion.String()
	Fleet_GroupVersionKind = CRDGroupVersion.WithKind(Fleet_Kind)
)

func init() {
	SchemeBuilder.Register(&Fleet{}, &FleetList{})
}
