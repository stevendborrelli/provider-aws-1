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

type RequestValidatorInitParameters struct {

	// Name of the request validator
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

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

	// Boolean whether to validate request body. Defaults to false.
	ValidateRequestBody *bool `json:"validateRequestBody,omitempty" tf:"validate_request_body,omitempty"`

	// Boolean whether to validate request parameters. Defaults to false.
	ValidateRequestParameters *bool `json:"validateRequestParameters,omitempty" tf:"validate_request_parameters,omitempty"`
}

type RequestValidatorObservation struct {

	// Unique ID of the request validator
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the request validator
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the associated Rest API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Boolean whether to validate request body. Defaults to false.
	ValidateRequestBody *bool `json:"validateRequestBody,omitempty" tf:"validate_request_body,omitempty"`

	// Boolean whether to validate request parameters. Defaults to false.
	ValidateRequestParameters *bool `json:"validateRequestParameters,omitempty" tf:"validate_request_parameters,omitempty"`
}

type RequestValidatorParameters struct {

	// Name of the request validator
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

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

	// Boolean whether to validate request body. Defaults to false.
	// +kubebuilder:validation:Optional
	ValidateRequestBody *bool `json:"validateRequestBody,omitempty" tf:"validate_request_body,omitempty"`

	// Boolean whether to validate request parameters. Defaults to false.
	// +kubebuilder:validation:Optional
	ValidateRequestParameters *bool `json:"validateRequestParameters,omitempty" tf:"validate_request_parameters,omitempty"`
}

// RequestValidatorSpec defines the desired state of RequestValidator
type RequestValidatorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RequestValidatorParameters `json:"forProvider"`
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
	InitProvider RequestValidatorInitParameters `json:"initProvider,omitempty"`
}

// RequestValidatorStatus defines the observed state of RequestValidator.
type RequestValidatorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RequestValidatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RequestValidator is the Schema for the RequestValidators API. Manages an API Gateway Request Validator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RequestValidator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RequestValidatorSpec   `json:"spec"`
	Status RequestValidatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RequestValidatorList contains a list of RequestValidators
type RequestValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RequestValidator `json:"items"`
}

// Repository type metadata.
var (
	RequestValidator_Kind             = "RequestValidator"
	RequestValidator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RequestValidator_Kind}.String()
	RequestValidator_KindAPIVersion   = RequestValidator_Kind + "." + CRDGroupVersion.String()
	RequestValidator_GroupVersionKind = CRDGroupVersion.WithKind(RequestValidator_Kind)
)

func init() {
	SchemeBuilder.Register(&RequestValidator{}, &RequestValidatorList{})
}
