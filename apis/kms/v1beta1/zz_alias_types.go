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

type AliasInitParameters struct {

	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	// +crossplane:generate:reference:type=Key
	TargetKeyID *string `json:"targetKeyId,omitempty" tf:"target_key_id,omitempty"`

	// Reference to a Key to populate targetKeyId.
	// +kubebuilder:validation:Optional
	TargetKeyIDRef *v1.Reference `json:"targetKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key to populate targetKeyId.
	// +kubebuilder:validation:Optional
	TargetKeyIDSelector *v1.Selector `json:"targetKeyIdSelector,omitempty" tf:"-"`
}

type AliasObservation struct {

	// The Amazon Resource Name (ARN) of the key alias.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the target key identifier.
	TargetKeyArn *string `json:"targetKeyArn,omitempty" tf:"target_key_arn,omitempty"`

	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyID *string `json:"targetKeyId,omitempty" tf:"target_key_id,omitempty"`
}

type AliasParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	// +crossplane:generate:reference:type=Key
	// +kubebuilder:validation:Optional
	TargetKeyID *string `json:"targetKeyId,omitempty" tf:"target_key_id,omitempty"`

	// Reference to a Key to populate targetKeyId.
	// +kubebuilder:validation:Optional
	TargetKeyIDRef *v1.Reference `json:"targetKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key to populate targetKeyId.
	// +kubebuilder:validation:Optional
	TargetKeyIDSelector *v1.Selector `json:"targetKeyIdSelector,omitempty" tf:"-"`
}

// AliasSpec defines the desired state of Alias
type AliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AliasParameters `json:"forProvider"`
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
	InitProvider AliasInitParameters `json:"initProvider,omitempty"`
}

// AliasStatus defines the observed state of Alias.
type AliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Alias is the Schema for the Aliass API. Provides a display name for a customer master key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Alias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AliasSpec   `json:"spec"`
	Status            AliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AliasList contains a list of Aliass
type AliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alias `json:"items"`
}

// Repository type metadata.
var (
	Alias_Kind             = "Alias"
	Alias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alias_Kind}.String()
	Alias_KindAPIVersion   = Alias_Kind + "." + CRDGroupVersion.String()
	Alias_GroupVersionKind = CRDGroupVersion.WithKind(Alias_Kind)
)

func init() {
	SchemeBuilder.Register(&Alias{}, &AliasList{})
}
