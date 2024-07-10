// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AttachmentAccepterInitParameters struct {

	// The ID of the attachment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta2.VPCAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`

	// Reference to a VPCAttachment in networkmanager to populate attachmentId.
	// +kubebuilder:validation:Optional
	AttachmentIDRef *v1.Reference `json:"attachmentIdRef,omitempty" tf:"-"`

	// Selector for a VPCAttachment in networkmanager to populate attachmentId.
	// +kubebuilder:validation:Optional
	AttachmentIDSelector *v1.Selector `json:"attachmentIdSelector,omitempty" tf:"-"`

	// The type of attachment. Valid values can be found in the AWS Documentation
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta2.VPCAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("attachment_type",true)
	AttachmentType *string `json:"attachmentType,omitempty" tf:"attachment_type,omitempty"`

	// Reference to a VPCAttachment in networkmanager to populate attachmentType.
	// +kubebuilder:validation:Optional
	AttachmentTypeRef *v1.Reference `json:"attachmentTypeRef,omitempty" tf:"-"`

	// Selector for a VPCAttachment in networkmanager to populate attachmentType.
	// +kubebuilder:validation:Optional
	AttachmentTypeSelector *v1.Selector `json:"attachmentTypeSelector,omitempty" tf:"-"`
}

type AttachmentAccepterObservation struct {

	// The ID of the attachment.
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`

	// The policy rule number associated with the attachment.
	AttachmentPolicyRuleNumber *float64 `json:"attachmentPolicyRuleNumber,omitempty" tf:"attachment_policy_rule_number,omitempty"`

	// The type of attachment. Valid values can be found in the AWS Documentation
	AttachmentType *string `json:"attachmentType,omitempty" tf:"attachment_type,omitempty"`

	// The ARN of a core network.
	CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

	// The id of a core network.
	CoreNetworkID *string `json:"coreNetworkId,omitempty" tf:"core_network_id,omitempty"`

	// The Region where the edge is located.
	EdgeLocation *string `json:"edgeLocation,omitempty" tf:"edge_location,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the attachment account owner.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// The attachment resource ARN.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// The name of the segment attachment.
	SegmentName *string `json:"segmentName,omitempty" tf:"segment_name,omitempty"`

	// The state of the attachment.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type AttachmentAccepterParameters struct {

	// The ID of the attachment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta2.VPCAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`

	// Reference to a VPCAttachment in networkmanager to populate attachmentId.
	// +kubebuilder:validation:Optional
	AttachmentIDRef *v1.Reference `json:"attachmentIdRef,omitempty" tf:"-"`

	// Selector for a VPCAttachment in networkmanager to populate attachmentId.
	// +kubebuilder:validation:Optional
	AttachmentIDSelector *v1.Selector `json:"attachmentIdSelector,omitempty" tf:"-"`

	// The type of attachment. Valid values can be found in the AWS Documentation
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta2.VPCAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("attachment_type",true)
	// +kubebuilder:validation:Optional
	AttachmentType *string `json:"attachmentType,omitempty" tf:"attachment_type,omitempty"`

	// Reference to a VPCAttachment in networkmanager to populate attachmentType.
	// +kubebuilder:validation:Optional
	AttachmentTypeRef *v1.Reference `json:"attachmentTypeRef,omitempty" tf:"-"`

	// Selector for a VPCAttachment in networkmanager to populate attachmentType.
	// +kubebuilder:validation:Optional
	AttachmentTypeSelector *v1.Selector `json:"attachmentTypeSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AttachmentAccepterSpec defines the desired state of AttachmentAccepter
type AttachmentAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentAccepterParameters `json:"forProvider"`
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
	InitProvider AttachmentAccepterInitParameters `json:"initProvider,omitempty"`
}

// AttachmentAccepterStatus defines the observed state of AttachmentAccepter.
type AttachmentAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AttachmentAccepter is the Schema for the AttachmentAccepters API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachmentAccepterSpec   `json:"spec"`
	Status            AttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentAccepterList contains a list of AttachmentAccepters
type AttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AttachmentAccepter `json:"items"`
}

// Repository type metadata.
var (
	AttachmentAccepter_Kind             = "AttachmentAccepter"
	AttachmentAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AttachmentAccepter_Kind}.String()
	AttachmentAccepter_KindAPIVersion   = AttachmentAccepter_Kind + "." + CRDGroupVersion.String()
	AttachmentAccepter_GroupVersionKind = CRDGroupVersion.WithKind(AttachmentAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&AttachmentAccepter{}, &AttachmentAccepterList{})
}
