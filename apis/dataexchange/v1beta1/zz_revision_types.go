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

type RevisionInitParameters struct {

	// An optional comment about the revision.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The dataset id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dataexchange/v1beta1.DataSet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DataSetID *string `json:"dataSetId,omitempty" tf:"data_set_id,omitempty"`

	// Reference to a DataSet in dataexchange to populate dataSetId.
	// +kubebuilder:validation:Optional
	DataSetIDRef *v1.Reference `json:"dataSetIdRef,omitempty" tf:"-"`

	// Selector for a DataSet in dataexchange to populate dataSetId.
	// +kubebuilder:validation:Optional
	DataSetIDSelector *v1.Selector `json:"dataSetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RevisionObservation struct {

	// The Amazon Resource Name of this data set.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// An optional comment about the revision.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The dataset id.
	DataSetID *string `json:"dataSetId,omitempty" tf:"data_set_id,omitempty"`

	// The Id of the data set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Id of the revision.
	RevisionID *string `json:"revisionId,omitempty" tf:"revision_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RevisionParameters struct {

	// An optional comment about the revision.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The dataset id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dataexchange/v1beta1.DataSet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataSetID *string `json:"dataSetId,omitempty" tf:"data_set_id,omitempty"`

	// Reference to a DataSet in dataexchange to populate dataSetId.
	// +kubebuilder:validation:Optional
	DataSetIDRef *v1.Reference `json:"dataSetIdRef,omitempty" tf:"-"`

	// Selector for a DataSet in dataexchange to populate dataSetId.
	// +kubebuilder:validation:Optional
	DataSetIDSelector *v1.Selector `json:"dataSetIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RevisionSpec defines the desired state of Revision
type RevisionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RevisionParameters `json:"forProvider"`
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
	InitProvider RevisionInitParameters `json:"initProvider,omitempty"`
}

// RevisionStatus defines the observed state of Revision.
type RevisionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RevisionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Revision is the Schema for the Revisions API. Provides a DataExchange Revision
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Revision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RevisionSpec   `json:"spec"`
	Status            RevisionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RevisionList contains a list of Revisions
type RevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Revision `json:"items"`
}

// Repository type metadata.
var (
	Revision_Kind             = "Revision"
	Revision_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Revision_Kind}.String()
	Revision_KindAPIVersion   = Revision_Kind + "." + CRDGroupVersion.String()
	Revision_GroupVersionKind = CRDGroupVersion.WithKind(Revision_Kind)
)

func init() {
	SchemeBuilder.Register(&Revision{}, &RevisionList{})
}
