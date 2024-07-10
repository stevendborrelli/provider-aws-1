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

type InstancePublicPortsInitParameters struct {

	// Name of the Lightsail Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta2.Instance
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// Configuration block with port information. AWS closes all currently open ports that are not included in the port_info. Detailed below.
	PortInfo []PortInfoInitParameters `json:"portInfo,omitempty" tf:"port_info,omitempty"`
}

type InstancePublicPortsObservation struct {

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Lightsail Instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Configuration block with port information. AWS closes all currently open ports that are not included in the port_info. Detailed below.
	PortInfo []PortInfoObservation `json:"portInfo,omitempty" tf:"port_info,omitempty"`
}

type InstancePublicPortsParameters struct {

	// Name of the Lightsail Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta2.Instance
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in lightsail to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// Configuration block with port information. AWS closes all currently open ports that are not included in the port_info. Detailed below.
	// +kubebuilder:validation:Optional
	PortInfo []PortInfoParameters `json:"portInfo,omitempty" tf:"port_info,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type PortInfoInitParameters struct {

	// Set of CIDR aliases that define access for a preconfigured range of IP addresses.
	// +listType=set
	CidrListAliases []*string `json:"cidrListAliases,omitempty" tf:"cidr_list_aliases,omitempty"`

	// Set of CIDR blocks.
	// +listType=set
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// First port in a range of open ports on an instance.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// +listType=set
	IPv6Cidrs []*string `json:"ipv6Cidrs,omitempty" tf:"ipv6_cidrs,omitempty"`

	// IP protocol name. Valid values are tcp, all, udp, and icmp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Last port in a range of open ports on an instance.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type PortInfoObservation struct {

	// Set of CIDR aliases that define access for a preconfigured range of IP addresses.
	// +listType=set
	CidrListAliases []*string `json:"cidrListAliases,omitempty" tf:"cidr_list_aliases,omitempty"`

	// Set of CIDR blocks.
	// +listType=set
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// First port in a range of open ports on an instance.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// +listType=set
	IPv6Cidrs []*string `json:"ipv6Cidrs,omitempty" tf:"ipv6_cidrs,omitempty"`

	// IP protocol name. Valid values are tcp, all, udp, and icmp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Last port in a range of open ports on an instance.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type PortInfoParameters struct {

	// Set of CIDR aliases that define access for a preconfigured range of IP addresses.
	// +kubebuilder:validation:Optional
	// +listType=set
	CidrListAliases []*string `json:"cidrListAliases,omitempty" tf:"cidr_list_aliases,omitempty"`

	// Set of CIDR blocks.
	// +kubebuilder:validation:Optional
	// +listType=set
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// First port in a range of open ports on an instance.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	// +listType=set
	IPv6Cidrs []*string `json:"ipv6Cidrs,omitempty" tf:"ipv6_cidrs,omitempty"`

	// IP protocol name. Valid values are tcp, all, udp, and icmp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Last port in a range of open ports on an instance.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort" tf:"to_port,omitempty"`
}

// InstancePublicPortsSpec defines the desired state of InstancePublicPorts
type InstancePublicPortsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstancePublicPortsParameters `json:"forProvider"`
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
	InitProvider InstancePublicPortsInitParameters `json:"initProvider,omitempty"`
}

// InstancePublicPortsStatus defines the observed state of InstancePublicPorts.
type InstancePublicPortsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstancePublicPortsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstancePublicPorts is the Schema for the InstancePublicPortss API. Provides an Lightsail Instance
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstancePublicPorts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.portInfo) || (has(self.initProvider) && has(self.initProvider.portInfo))",message="spec.forProvider.portInfo is a required parameter"
	Spec   InstancePublicPortsSpec   `json:"spec"`
	Status InstancePublicPortsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstancePublicPortsList contains a list of InstancePublicPortss
type InstancePublicPortsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstancePublicPorts `json:"items"`
}

// Repository type metadata.
var (
	InstancePublicPorts_Kind             = "InstancePublicPorts"
	InstancePublicPorts_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstancePublicPorts_Kind}.String()
	InstancePublicPorts_KindAPIVersion   = InstancePublicPorts_Kind + "." + CRDGroupVersion.String()
	InstancePublicPorts_GroupVersionKind = CRDGroupVersion.WithKind(InstancePublicPorts_Kind)
)

func init() {
	SchemeBuilder.Register(&InstancePublicPorts{}, &InstancePublicPortsList{})
}
