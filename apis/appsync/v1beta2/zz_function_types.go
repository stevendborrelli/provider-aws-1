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

type FunctionInitParameters struct {

	// ID of the associated AppSync API.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta2.GraphQLAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Function data source name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta2.Datasource
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Reference to a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceRef *v1.Reference `json:"dataSourceRef,omitempty" tf:"-"`

	// Selector for a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceSelector *v1.Selector `json:"dataSourceSelector,omitempty" tf:"-"`

	// Function description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Version of the request mapping template. Currently the supported value is 2018-05-29. Does not apply when specifying code.
	FunctionVersion *string `json:"functionVersion,omitempty" tf:"function_version,omitempty"`

	// Maximum batching size for a resolver. Valid values are between 0 and 2000.
	MaxBatchSize *float64 `json:"maxBatchSize,omitempty" tf:"max_batch_size,omitempty"`

	// Function name. The function name does not have to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `json:"requestMappingTemplate,omitempty" tf:"request_mapping_template,omitempty"`

	// Function response mapping template.
	ResponseMappingTemplate *string `json:"responseMappingTemplate,omitempty" tf:"response_mapping_template,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See runtime Block for details.
	Runtime *RuntimeInitParameters `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Describes a Sync configuration for a resolver. See sync_config Block for details.
	SyncConfig *SyncConfigInitParameters `json:"syncConfig,omitempty" tf:"sync_config,omitempty"`
}

type FunctionObservation struct {

	// ID of the associated AppSync API.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// ARN of the Function object.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Function data source name.
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Function description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Unique ID representing the Function object.
	FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

	// Version of the request mapping template. Currently the supported value is 2018-05-29. Does not apply when specifying code.
	FunctionVersion *string `json:"functionVersion,omitempty" tf:"function_version,omitempty"`

	// API Function ID (Formatted as ApiId-FunctionId)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum batching size for a resolver. Valid values are between 0 and 2000.
	MaxBatchSize *float64 `json:"maxBatchSize,omitempty" tf:"max_batch_size,omitempty"`

	// Function name. The function name does not have to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `json:"requestMappingTemplate,omitempty" tf:"request_mapping_template,omitempty"`

	// Function response mapping template.
	ResponseMappingTemplate *string `json:"responseMappingTemplate,omitempty" tf:"response_mapping_template,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See runtime Block for details.
	Runtime *RuntimeObservation `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Describes a Sync configuration for a resolver. See sync_config Block for details.
	SyncConfig *SyncConfigObservation `json:"syncConfig,omitempty" tf:"sync_config,omitempty"`
}

type FunctionParameters struct {

	// ID of the associated AppSync API.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta2.GraphQLAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	// +kubebuilder:validation:Optional
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Function data source name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta2.Datasource
	// +kubebuilder:validation:Optional
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Reference to a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceRef *v1.Reference `json:"dataSourceRef,omitempty" tf:"-"`

	// Selector for a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceSelector *v1.Selector `json:"dataSourceSelector,omitempty" tf:"-"`

	// Function description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Version of the request mapping template. Currently the supported value is 2018-05-29. Does not apply when specifying code.
	// +kubebuilder:validation:Optional
	FunctionVersion *string `json:"functionVersion,omitempty" tf:"function_version,omitempty"`

	// Maximum batching size for a resolver. Valid values are between 0 and 2000.
	// +kubebuilder:validation:Optional
	MaxBatchSize *float64 `json:"maxBatchSize,omitempty" tf:"max_batch_size,omitempty"`

	// Function name. The function name does not have to be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	// +kubebuilder:validation:Optional
	RequestMappingTemplate *string `json:"requestMappingTemplate,omitempty" tf:"request_mapping_template,omitempty"`

	// Function response mapping template.
	// +kubebuilder:validation:Optional
	ResponseMappingTemplate *string `json:"responseMappingTemplate,omitempty" tf:"response_mapping_template,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See runtime Block for details.
	// +kubebuilder:validation:Optional
	Runtime *RuntimeParameters `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Describes a Sync configuration for a resolver. See sync_config Block for details.
	// +kubebuilder:validation:Optional
	SyncConfig *SyncConfigParameters `json:"syncConfig,omitempty" tf:"sync_config,omitempty"`
}

type LambdaConflictHandlerConfigInitParameters struct {

	// ARN for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn,omitempty"`
}

type LambdaConflictHandlerConfigObservation struct {

	// ARN for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn,omitempty"`
}

type LambdaConflictHandlerConfigParameters struct {

	// ARN for the Lambda function to use as the Conflict Handler.
	// +kubebuilder:validation:Optional
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn,omitempty"`
}

type RuntimeInitParameters struct {

	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

type RuntimeObservation struct {

	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

type RuntimeParameters struct {

	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion" tf:"runtime_version,omitempty"`
}

type SyncConfigInitParameters struct {

	// Conflict Detection strategy to use. Valid values are NONE and VERSION.
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection,omitempty"`

	// Conflict Resolution strategy to perform in the event of a conflict. Valid values are NONE, OPTIMISTIC_CONCURRENCY, AUTOMERGE, and LAMBDA.
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler,omitempty"`

	// Lambda Conflict Handler Config when configuring LAMBDA as the Conflict Handler. See lambda_conflict_handler_config Block for details.
	LambdaConflictHandlerConfig *LambdaConflictHandlerConfigInitParameters `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config,omitempty"`
}

type SyncConfigObservation struct {

	// Conflict Detection strategy to use. Valid values are NONE and VERSION.
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection,omitempty"`

	// Conflict Resolution strategy to perform in the event of a conflict. Valid values are NONE, OPTIMISTIC_CONCURRENCY, AUTOMERGE, and LAMBDA.
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler,omitempty"`

	// Lambda Conflict Handler Config when configuring LAMBDA as the Conflict Handler. See lambda_conflict_handler_config Block for details.
	LambdaConflictHandlerConfig *LambdaConflictHandlerConfigObservation `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config,omitempty"`
}

type SyncConfigParameters struct {

	// Conflict Detection strategy to use. Valid values are NONE and VERSION.
	// +kubebuilder:validation:Optional
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection,omitempty"`

	// Conflict Resolution strategy to perform in the event of a conflict. Valid values are NONE, OPTIMISTIC_CONCURRENCY, AUTOMERGE, and LAMBDA.
	// +kubebuilder:validation:Optional
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler,omitempty"`

	// Lambda Conflict Handler Config when configuring LAMBDA as the Conflict Handler. See lambda_conflict_handler_config Block for details.
	// +kubebuilder:validation:Optional
	LambdaConflictHandlerConfig *LambdaConflictHandlerConfigParameters `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
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
	InitProvider FunctionInitParameters `json:"initProvider,omitempty"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Function is the Schema for the Functions API. Provides an AppSync Function.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FunctionSpec   `json:"spec"`
	Status FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
