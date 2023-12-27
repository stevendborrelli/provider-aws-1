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

type DocumentationVersionInitParameters struct {

	// Description of the API documentation version.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the associated Rest API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// Version identifier of the API documentation snapshot.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type DocumentationVersionObservation struct {

	// Description of the API documentation version.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the associated Rest API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Version identifier of the API documentation snapshot.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type DocumentationVersionParameters struct {

	// Description of the API documentation version.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the associated Rest API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// Version identifier of the API documentation snapshot.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// DocumentationVersionSpec defines the desired state of DocumentationVersion
type DocumentationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentationVersionParameters `json:"forProvider"`
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
	InitProvider DocumentationVersionInitParameters `json:"initProvider,omitempty"`
}

// DocumentationVersionStatus defines the observed state of DocumentationVersion.
type DocumentationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationVersion is the Schema for the DocumentationVersions API. Provides a resource to manage an API Gateway Documentation Version.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocumentationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   DocumentationVersionSpec   `json:"spec"`
	Status DocumentationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationVersionList contains a list of DocumentationVersions
type DocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentationVersion `json:"items"`
}

// Repository type metadata.
var (
	DocumentationVersion_Kind             = "DocumentationVersion"
	DocumentationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DocumentationVersion_Kind}.String()
	DocumentationVersion_KindAPIVersion   = DocumentationVersion_Kind + "." + CRDGroupVersion.String()
	DocumentationVersion_GroupVersionKind = CRDGroupVersion.WithKind(DocumentationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&DocumentationVersion{}, &DocumentationVersionList{})
}
