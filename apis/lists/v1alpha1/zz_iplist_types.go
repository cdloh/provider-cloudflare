// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IPListInitParameters struct {

	// The ID of the account where the IP List is being created.
	// The account identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// A note that can be used to annotate the List. Maximum Length: 500
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Item []ItemInitParameters `json:"item,omitempty" tf:"item,omitempty"`

	// The kind of values in the List. Valid values: ip.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The name of the list (used in filter expressions). Valid pattern: ^[a-zA-Z0-9_]+$. Maximum Length: 50
	// **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IPListObservation struct {

	// The ID of the account where the IP List is being created.
	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A note that can be used to annotate the List. Maximum Length: 500
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Item []ItemObservation `json:"item,omitempty" tf:"item,omitempty"`

	// The kind of values in the List. Valid values: ip.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The name of the list (used in filter expressions). Valid pattern: ^[a-zA-Z0-9_]+$. Maximum Length: 50
	// **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IPListParameters struct {

	// The ID of the account where the IP List is being created.
	// The account identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// A note that can be used to annotate the List. Maximum Length: 500
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Item []ItemParameters `json:"item,omitempty" tf:"item,omitempty"`

	// The kind of values in the List. Valid values: ip.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The name of the list (used in filter expressions). Valid pattern: ^[a-zA-Z0-9_]+$. Maximum Length: 50
	// **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ItemInitParameters struct {

	// A note that can be used to annotate the item.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The IPv4 address, IPv4 CIDR or IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ItemObservation struct {

	// A note that can be used to annotate the item.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The IPv4 address, IPv4 CIDR or IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ItemParameters struct {

	// A note that can be used to annotate the item.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The IPv4 address, IPv4 CIDR or IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// IPListSpec defines the desired state of IPList
type IPListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPListParameters `json:"forProvider"`
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
	InitProvider IPListInitParameters `json:"initProvider,omitempty"`
}

// IPListStatus defines the observed state of IPList.
type IPListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IPList is the Schema for the IPLists API. Provides IP Lists to be used in Firewall Rules across all zones within the same account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type IPList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kind) || (has(self.initProvider) && has(self.initProvider.kind))",message="spec.forProvider.kind is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IPListSpec   `json:"spec"`
	Status IPListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPListList contains a list of IPLists
type IPListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPList `json:"items"`
}

// Repository type metadata.
var (
	IPList_Kind             = "IPList"
	IPList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPList_Kind}.String()
	IPList_KindAPIVersion   = IPList_Kind + "." + CRDGroupVersion.String()
	IPList_GroupVersionKind = CRDGroupVersion.WithKind(IPList_Kind)
)

func init() {
	SchemeBuilder.Register(&IPList{}, &IPListList{})
}