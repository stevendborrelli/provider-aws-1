// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InputDestinationsInitParameters struct {

	// A unique name for the location the RTMP stream is being pushed to.
	StreamName *string `json:"streamName,omitempty" tf:"stream_name,omitempty"`
}

type InputDestinationsObservation struct {

	// A unique name for the location the RTMP stream is being pushed to.
	StreamName *string `json:"streamName,omitempty" tf:"stream_name,omitempty"`
}

type InputDestinationsParameters struct {

	// A unique name for the location the RTMP stream is being pushed to.
	// +kubebuilder:validation:Optional
	StreamName *string `json:"streamName" tf:"stream_name,omitempty"`
}

type InputDevicesInitParameters struct {

	// The unique ID for the device.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InputDevicesObservation struct {

	// The unique ID for the device.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InputDevicesParameters struct {

	// The unique ID for the device.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`
}

type InputInitParameters struct {

	// Destination settings for PUSH type inputs. See Destinations for more details.
	Destinations []InputDestinationsInitParameters `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// Settings for the devices. See Input Devices for more details.
	InputDevices []InputDevicesInitParameters `json:"inputDevices,omitempty" tf:"input_devices,omitempty"`

	// List of input security groups.
	InputSecurityGroups []*string `json:"inputSecurityGroups,omitempty" tf:"input_security_groups,omitempty"`

	// A list of the MediaConnect Flows. See Media Connect Flows for more details.
	MediaConnectFlows []MediaConnectFlowsInitParameters `json:"mediaConnectFlows,omitempty" tf:"media_connect_flows,omitempty"`

	// Name of the input.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the role this input assumes during and after creation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The source URLs for a PULL-type input. See Sources for more details.
	Sources []SourcesInitParameters `json:"sources,omitempty" tf:"sources,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The different types of inputs that AWS Elemental MediaLive supports.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Settings for a private VPC Input. See VPC for more details.
	VPC []InputVPCInitParameters `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type InputObservation struct {

	// ARN of the Input.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Channels attached to Input.
	AttachedChannels []*string `json:"attachedChannels,omitempty" tf:"attached_channels,omitempty"`

	// Destination settings for PUSH type inputs. See Destinations for more details.
	Destinations []InputDestinationsObservation `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// The unique ID for the device.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The input class.
	InputClass *string `json:"inputClass,omitempty" tf:"input_class,omitempty"`

	// Settings for the devices. See Input Devices for more details.
	InputDevices []InputDevicesObservation `json:"inputDevices,omitempty" tf:"input_devices,omitempty"`

	// A list of IDs for all Inputs which are partners of this one.
	InputPartnerIds []*string `json:"inputPartnerIds,omitempty" tf:"input_partner_ids,omitempty"`

	// List of input security groups.
	InputSecurityGroups []*string `json:"inputSecurityGroups,omitempty" tf:"input_security_groups,omitempty"`

	// Source type of the input.
	InputSourceType *string `json:"inputSourceType,omitempty" tf:"input_source_type,omitempty"`

	// A list of the MediaConnect Flows. See Media Connect Flows for more details.
	MediaConnectFlows []MediaConnectFlowsObservation `json:"mediaConnectFlows,omitempty" tf:"media_connect_flows,omitempty"`

	// Name of the input.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the role this input assumes during and after creation.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The source URLs for a PULL-type input. See Sources for more details.
	Sources []SourcesObservation `json:"sources,omitempty" tf:"sources,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The different types of inputs that AWS Elemental MediaLive supports.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Settings for a private VPC Input. See VPC for more details.
	VPC []InputVPCObservation `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type InputParameters struct {

	// Destination settings for PUSH type inputs. See Destinations for more details.
	// +kubebuilder:validation:Optional
	Destinations []InputDestinationsParameters `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// Settings for the devices. See Input Devices for more details.
	// +kubebuilder:validation:Optional
	InputDevices []InputDevicesParameters `json:"inputDevices,omitempty" tf:"input_devices,omitempty"`

	// List of input security groups.
	// +kubebuilder:validation:Optional
	InputSecurityGroups []*string `json:"inputSecurityGroups,omitempty" tf:"input_security_groups,omitempty"`

	// A list of the MediaConnect Flows. See Media Connect Flows for more details.
	// +kubebuilder:validation:Optional
	MediaConnectFlows []MediaConnectFlowsParameters `json:"mediaConnectFlows,omitempty" tf:"media_connect_flows,omitempty"`

	// Name of the input.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the role this input assumes during and after creation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The source URLs for a PULL-type input. See Sources for more details.
	// +kubebuilder:validation:Optional
	Sources []SourcesParameters `json:"sources,omitempty" tf:"sources,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The different types of inputs that AWS Elemental MediaLive supports.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Settings for a private VPC Input. See VPC for more details.
	// +kubebuilder:validation:Optional
	VPC []InputVPCParameters `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type InputVPCInitParameters struct {

	// A list of up to 5 EC2 VPC security group IDs to attach to the Input.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of 2 VPC subnet IDs from the same VPC.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type InputVPCObservation struct {

	// A list of up to 5 EC2 VPC security group IDs to attach to the Input.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of 2 VPC subnet IDs from the same VPC.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type InputVPCParameters struct {

	// A list of up to 5 EC2 VPC security group IDs to attach to the Input.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of 2 VPC subnet IDs from the same VPC.
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`
}

type MediaConnectFlowsInitParameters struct {

	// The ARN of the MediaConnect Flow
	FlowArn *string `json:"flowArn,omitempty" tf:"flow_arn,omitempty"`
}

type MediaConnectFlowsObservation struct {

	// The ARN of the MediaConnect Flow
	FlowArn *string `json:"flowArn,omitempty" tf:"flow_arn,omitempty"`
}

type MediaConnectFlowsParameters struct {

	// The ARN of the MediaConnect Flow
	// +kubebuilder:validation:Optional
	FlowArn *string `json:"flowArn" tf:"flow_arn,omitempty"`
}

type SourcesInitParameters struct {

	// The key used to extract the password from EC2 Parameter store.
	PasswordParam *string `json:"passwordParam,omitempty" tf:"password_param,omitempty"`

	// The URL where the stream is pulled from.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username for the input source.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SourcesObservation struct {

	// The key used to extract the password from EC2 Parameter store.
	PasswordParam *string `json:"passwordParam,omitempty" tf:"password_param,omitempty"`

	// The URL where the stream is pulled from.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username for the input source.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SourcesParameters struct {

	// The key used to extract the password from EC2 Parameter store.
	// +kubebuilder:validation:Optional
	PasswordParam *string `json:"passwordParam" tf:"password_param,omitempty"`

	// The URL where the stream is pulled from.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`

	// The username for the input source.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

// InputSpec defines the desired state of Input
type InputSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InputParameters `json:"forProvider"`
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
	InitProvider InputInitParameters `json:"initProvider,omitempty"`
}

// InputStatus defines the observed state of Input.
type InputStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InputObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Input is the Schema for the Inputs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Input struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   InputSpec   `json:"spec"`
	Status InputStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InputList contains a list of Inputs
type InputList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Input `json:"items"`
}

// Repository type metadata.
var (
	Input_Kind             = "Input"
	Input_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Input_Kind}.String()
	Input_KindAPIVersion   = Input_Kind + "." + CRDGroupVersion.String()
	Input_GroupVersionKind = CRDGroupVersion.WithKind(Input_Kind)
)

func init() {
	SchemeBuilder.Register(&Input{}, &InputList{})
}
