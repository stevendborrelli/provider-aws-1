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

type EventSubscriptionInitParameters struct {

	// Whether the event subscription should be enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of event categories to listen for, see DescribeEventCategories for a canonical list.
	// +listType=set
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// SNS topic arn to send events on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Reference to a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`

	// Ids of sources to listen to.
	// +listType=set
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// Type of source for events. Valid values: replication-instance or replication-task
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventSubscriptionObservation struct {

	// Amazon Resource Name (ARN) of the DMS Event Subscription.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether the event subscription should be enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of event categories to listen for, see DescribeEventCategories for a canonical list.
	// +listType=set
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// SNS topic arn to send events on.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Ids of sources to listen to.
	// +listType=set
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// Type of source for events. Valid values: replication-instance or replication-task
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventSubscriptionParameters struct {

	// Whether the event subscription should be enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of event categories to listen for, see DescribeEventCategories for a canonical list.
	// +kubebuilder:validation:Optional
	// +listType=set
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// SNS topic arn to send events on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Reference to a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`

	// Ids of sources to listen to.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// Type of source for events. Valid values: replication-instance or replication-task
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EventSubscriptionSpec defines the desired state of EventSubscription
type EventSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSubscriptionParameters `json:"forProvider"`
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
	InitProvider EventSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// EventSubscriptionStatus defines the observed state of EventSubscription.
type EventSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventSubscription is the Schema for the EventSubscriptions API. Provides a DMS (Data Migration Service) event subscription resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventCategories) || (has(self.initProvider) && has(self.initProvider.eventCategories))",message="spec.forProvider.eventCategories is a required parameter"
	Spec   EventSubscriptionSpec   `json:"spec"`
	Status EventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSubscriptionList contains a list of EventSubscriptions
type EventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSubscription `json:"items"`
}

// Repository type metadata.
var (
	EventSubscription_Kind             = "EventSubscription"
	EventSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSubscription_Kind}.String()
	EventSubscription_KindAPIVersion   = EventSubscription_Kind + "." + CRDGroupVersion.String()
	EventSubscription_GroupVersionKind = CRDGroupVersion.WithKind(EventSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSubscription{}, &EventSubscriptionList{})
}
