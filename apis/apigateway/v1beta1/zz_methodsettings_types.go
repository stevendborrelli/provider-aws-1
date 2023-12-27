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

type MethodSettingsInitParameters struct {

	// Method path defined as {resource_path}/{http_method} for an individual method override, or */* for overriding all methods in the stage. Ensure to trim any leading forward slashes in the path (e.g., trimprefix(aws_api_gateway_resource.example.path, "/")).
	MethodPath *string `json:"methodPath,omitempty" tf:"method_path,omitempty"`

	// ID of the REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// Settings block, see below.
	Settings []SettingsInitParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// Name of the stage
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Stage
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("stage_name",false)
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`

	// Reference to a Stage in apigateway to populate stageName.
	// +kubebuilder:validation:Optional
	StageNameRef *v1.Reference `json:"stageNameRef,omitempty" tf:"-"`

	// Selector for a Stage in apigateway to populate stageName.
	// +kubebuilder:validation:Optional
	StageNameSelector *v1.Selector `json:"stageNameSelector,omitempty" tf:"-"`
}

type MethodSettingsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Method path defined as {resource_path}/{http_method} for an individual method override, or */* for overriding all methods in the stage. Ensure to trim any leading forward slashes in the path (e.g., trimprefix(aws_api_gateway_resource.example.path, "/")).
	MethodPath *string `json:"methodPath,omitempty" tf:"method_path,omitempty"`

	// ID of the REST API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Settings block, see below.
	Settings []SettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`

	// Name of the stage
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`
}

type MethodSettingsParameters struct {

	// Method path defined as {resource_path}/{http_method} for an individual method override, or */* for overriding all methods in the stage. Ensure to trim any leading forward slashes in the path (e.g., trimprefix(aws_api_gateway_resource.example.path, "/")).
	// +kubebuilder:validation:Optional
	MethodPath *string `json:"methodPath,omitempty" tf:"method_path,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the REST API
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

	// Settings block, see below.
	// +kubebuilder:validation:Optional
	Settings []SettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// Name of the stage
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Stage
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("stage_name",false)
	// +kubebuilder:validation:Optional
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`

	// Reference to a Stage in apigateway to populate stageName.
	// +kubebuilder:validation:Optional
	StageNameRef *v1.Reference `json:"stageNameRef,omitempty" tf:"-"`

	// Selector for a Stage in apigateway to populate stageName.
	// +kubebuilder:validation:Optional
	StageNameSelector *v1.Selector `json:"stageNameSelector,omitempty" tf:"-"`
}

type SettingsInitParameters struct {

	// Whether the cached responses are encrypted.
	CacheDataEncrypted *bool `json:"cacheDataEncrypted,omitempty" tf:"cache_data_encrypted,omitempty"`

	// Time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
	CacheTTLInSeconds *float64 `json:"cacheTtlInSeconds,omitempty" tf:"cache_ttl_in_seconds,omitempty"`

	// Whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.
	CachingEnabled *bool `json:"cachingEnabled,omitempty" tf:"caching_enabled,omitempty"`

	// Whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are OFF, ERROR, and INFO.
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Whether Amazon CloudWatch metrics are enabled for this method.
	MetricsEnabled *bool `json:"metricsEnabled,omitempty" tf:"metrics_enabled,omitempty"`

	// Whether authorization is required for a cache invalidation request.
	RequireAuthorizationForCacheControl *bool `json:"requireAuthorizationForCacheControl,omitempty" tf:"require_authorization_for_cache_control,omitempty"`

	// Throttling burst limit. Default: -1 (throttling disabled).
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit. Default: -1 (throttling disabled).
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`

	// How to handle unauthorized requests for cache invalidation. The available values are FAIL_WITH_403, SUCCEED_WITH_RESPONSE_HEADER, SUCCEED_WITHOUT_RESPONSE_HEADER.
	UnauthorizedCacheControlHeaderStrategy *string `json:"unauthorizedCacheControlHeaderStrategy,omitempty" tf:"unauthorized_cache_control_header_strategy,omitempty"`
}

type SettingsObservation struct {

	// Whether the cached responses are encrypted.
	CacheDataEncrypted *bool `json:"cacheDataEncrypted,omitempty" tf:"cache_data_encrypted,omitempty"`

	// Time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
	CacheTTLInSeconds *float64 `json:"cacheTtlInSeconds,omitempty" tf:"cache_ttl_in_seconds,omitempty"`

	// Whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.
	CachingEnabled *bool `json:"cachingEnabled,omitempty" tf:"caching_enabled,omitempty"`

	// Whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are OFF, ERROR, and INFO.
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Whether Amazon CloudWatch metrics are enabled for this method.
	MetricsEnabled *bool `json:"metricsEnabled,omitempty" tf:"metrics_enabled,omitempty"`

	// Whether authorization is required for a cache invalidation request.
	RequireAuthorizationForCacheControl *bool `json:"requireAuthorizationForCacheControl,omitempty" tf:"require_authorization_for_cache_control,omitempty"`

	// Throttling burst limit. Default: -1 (throttling disabled).
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit. Default: -1 (throttling disabled).
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`

	// How to handle unauthorized requests for cache invalidation. The available values are FAIL_WITH_403, SUCCEED_WITH_RESPONSE_HEADER, SUCCEED_WITHOUT_RESPONSE_HEADER.
	UnauthorizedCacheControlHeaderStrategy *string `json:"unauthorizedCacheControlHeaderStrategy,omitempty" tf:"unauthorized_cache_control_header_strategy,omitempty"`
}

type SettingsParameters struct {

	// Whether the cached responses are encrypted.
	// +kubebuilder:validation:Optional
	CacheDataEncrypted *bool `json:"cacheDataEncrypted,omitempty" tf:"cache_data_encrypted,omitempty"`

	// Time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
	// +kubebuilder:validation:Optional
	CacheTTLInSeconds *float64 `json:"cacheTtlInSeconds,omitempty" tf:"cache_ttl_in_seconds,omitempty"`

	// Whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.
	// +kubebuilder:validation:Optional
	CachingEnabled *bool `json:"cachingEnabled,omitempty" tf:"caching_enabled,omitempty"`

	// Whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	// +kubebuilder:validation:Optional
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are OFF, ERROR, and INFO.
	// +kubebuilder:validation:Optional
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Whether Amazon CloudWatch metrics are enabled for this method.
	// +kubebuilder:validation:Optional
	MetricsEnabled *bool `json:"metricsEnabled,omitempty" tf:"metrics_enabled,omitempty"`

	// Whether authorization is required for a cache invalidation request.
	// +kubebuilder:validation:Optional
	RequireAuthorizationForCacheControl *bool `json:"requireAuthorizationForCacheControl,omitempty" tf:"require_authorization_for_cache_control,omitempty"`

	// Throttling burst limit. Default: -1 (throttling disabled).
	// +kubebuilder:validation:Optional
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit. Default: -1 (throttling disabled).
	// +kubebuilder:validation:Optional
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`

	// How to handle unauthorized requests for cache invalidation. The available values are FAIL_WITH_403, SUCCEED_WITH_RESPONSE_HEADER, SUCCEED_WITHOUT_RESPONSE_HEADER.
	// +kubebuilder:validation:Optional
	UnauthorizedCacheControlHeaderStrategy *string `json:"unauthorizedCacheControlHeaderStrategy,omitempty" tf:"unauthorized_cache_control_header_strategy,omitempty"`
}

// MethodSettingsSpec defines the desired state of MethodSettings
type MethodSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MethodSettingsParameters `json:"forProvider"`
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
	InitProvider MethodSettingsInitParameters `json:"initProvider,omitempty"`
}

// MethodSettingsStatus defines the observed state of MethodSettings.
type MethodSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MethodSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MethodSettings is the Schema for the MethodSettingss API. Manages API Gateway Stage Method Settings
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MethodSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.methodPath) || (has(self.initProvider) && has(self.initProvider.methodPath))",message="spec.forProvider.methodPath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.settings) || (has(self.initProvider) && has(self.initProvider.settings))",message="spec.forProvider.settings is a required parameter"
	Spec   MethodSettingsSpec   `json:"spec"`
	Status MethodSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MethodSettingsList contains a list of MethodSettingss
type MethodSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MethodSettings `json:"items"`
}

// Repository type metadata.
var (
	MethodSettings_Kind             = "MethodSettings"
	MethodSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MethodSettings_Kind}.String()
	MethodSettings_KindAPIVersion   = MethodSettings_Kind + "." + CRDGroupVersion.String()
	MethodSettings_GroupVersionKind = CRDGroupVersion.WithKind(MethodSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&MethodSettings{}, &MethodSettingsList{})
}
