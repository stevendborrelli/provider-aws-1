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

type InputDataConfigInitParameters struct {

	// IAM role with access to S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	DataAccessRoleArn *string `json:"dataAccessRoleArn,omitempty" tf:"data_access_role_arn,omitempty"`

	// Reference to a Role in iam to populate dataAccessRoleArn.
	// +kubebuilder:validation:Optional
	DataAccessRoleArnRef *v1.Reference `json:"dataAccessRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate dataAccessRoleArn.
	// +kubebuilder:validation:Optional
	DataAccessRoleArnSelector *v1.Selector `json:"dataAccessRoleArnSelector,omitempty" tf:"-"`

	// S3 URI where training data is located.
	S3URI *string `json:"s3Uri,omitempty" tf:"s3_uri,omitempty"`

	// S3 URI where tuning data is located.
	TuningDataS3URI *string `json:"tuningDataS3Uri,omitempty" tf:"tuning_data_s3_uri,omitempty"`
}

type InputDataConfigObservation struct {

	// IAM role with access to S3 bucket.
	DataAccessRoleArn *string `json:"dataAccessRoleArn,omitempty" tf:"data_access_role_arn,omitempty"`

	// S3 URI where training data is located.
	S3URI *string `json:"s3Uri,omitempty" tf:"s3_uri,omitempty"`

	// S3 URI where tuning data is located.
	TuningDataS3URI *string `json:"tuningDataS3Uri,omitempty" tf:"tuning_data_s3_uri,omitempty"`
}

type InputDataConfigParameters struct {

	// IAM role with access to S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	DataAccessRoleArn *string `json:"dataAccessRoleArn,omitempty" tf:"data_access_role_arn,omitempty"`

	// Reference to a Role in iam to populate dataAccessRoleArn.
	// +kubebuilder:validation:Optional
	DataAccessRoleArnRef *v1.Reference `json:"dataAccessRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate dataAccessRoleArn.
	// +kubebuilder:validation:Optional
	DataAccessRoleArnSelector *v1.Selector `json:"dataAccessRoleArnSelector,omitempty" tf:"-"`

	// S3 URI where training data is located.
	// +kubebuilder:validation:Optional
	S3URI *string `json:"s3Uri" tf:"s3_uri,omitempty"`

	// S3 URI where tuning data is located.
	// +kubebuilder:validation:Optional
	TuningDataS3URI *string `json:"tuningDataS3Uri,omitempty" tf:"tuning_data_s3_uri,omitempty"`
}

type LanguageModelInitParameters struct {

	// Name of reference base model.
	BaseModelName *string `json:"baseModelName,omitempty" tf:"base_model_name,omitempty"`

	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig []InputDataConfigInitParameters `json:"inputDataConfig,omitempty" tf:"input_data_config,omitempty"`

	// The language code you selected for your language model. Refer to the supported languages page for accepted codes.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LanguageModelObservation struct {

	// ARN of the LanguageModel.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name of reference base model.
	BaseModelName *string `json:"baseModelName,omitempty" tf:"base_model_name,omitempty"`

	// LanguageModel name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig []InputDataConfigObservation `json:"inputDataConfig,omitempty" tf:"input_data_config,omitempty"`

	// The language code you selected for your language model. Refer to the supported languages page for accepted codes.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LanguageModelParameters struct {

	// Name of reference base model.
	// +kubebuilder:validation:Optional
	BaseModelName *string `json:"baseModelName,omitempty" tf:"base_model_name,omitempty"`

	// The input data config for the LanguageModel. See Input Data Config for more details.
	// +kubebuilder:validation:Optional
	InputDataConfig []InputDataConfigParameters `json:"inputDataConfig,omitempty" tf:"input_data_config,omitempty"`

	// The language code you selected for your language model. Refer to the supported languages page for accepted codes.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LanguageModelSpec defines the desired state of LanguageModel
type LanguageModelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LanguageModelParameters `json:"forProvider"`
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
	InitProvider LanguageModelInitParameters `json:"initProvider,omitempty"`
}

// LanguageModelStatus defines the observed state of LanguageModel.
type LanguageModelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LanguageModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LanguageModel is the Schema for the LanguageModels API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LanguageModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.baseModelName) || (has(self.initProvider) && has(self.initProvider.baseModelName))",message="spec.forProvider.baseModelName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.inputDataConfig) || (has(self.initProvider) && has(self.initProvider.inputDataConfig))",message="spec.forProvider.inputDataConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.languageCode) || (has(self.initProvider) && has(self.initProvider.languageCode))",message="spec.forProvider.languageCode is a required parameter"
	Spec   LanguageModelSpec   `json:"spec"`
	Status LanguageModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LanguageModelList contains a list of LanguageModels
type LanguageModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LanguageModel `json:"items"`
}

// Repository type metadata.
var (
	LanguageModel_Kind             = "LanguageModel"
	LanguageModel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LanguageModel_Kind}.String()
	LanguageModel_KindAPIVersion   = LanguageModel_Kind + "." + CRDGroupVersion.String()
	LanguageModel_GroupVersionKind = CRDGroupVersion.WithKind(LanguageModel_Kind)
)

func init() {
	SchemeBuilder.Register(&LanguageModel{}, &LanguageModelList{})
}
