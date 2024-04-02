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

type ListInitParameters struct {

	// (String) The account identifier to target for the resource.
	// The account identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// (String) An optional description of the list.
	// An optional description of the list.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block Set) (see below for nested schema)
	Item []ListItemInitParameters `json:"item,omitempty" tf:"item,omitempty"`

	// (String) The type of items the list will contain.
	// The type of items the list will contain.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// (String) The name of the list. Modifying this attribute will force creation of a new resource.
	// The name of the list. **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ListItemInitParameters struct {

	// (String) An optional comment for the item.
	// An optional comment for the item.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Value []ValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type ListItemObservation struct {

	// (String) An optional comment for the item.
	// An optional comment for the item.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Value []ValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type ListItemParameters struct {

	// (String) An optional comment for the item.
	// An optional comment for the item.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Value []ValueParameters `json:"value" tf:"value,omitempty"`
}

type ListObservation struct {

	// (String) The account identifier to target for the resource.
	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) An optional description of the list.
	// An optional description of the list.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block Set) (see below for nested schema)
	Item []ListItemObservation `json:"item,omitempty" tf:"item,omitempty"`

	// (String) The type of items the list will contain.
	// The type of items the list will contain.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// (String) The name of the list. Modifying this attribute will force creation of a new resource.
	// The name of the list. **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ListParameters struct {

	// (String) The account identifier to target for the resource.
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

	// (String) An optional description of the list.
	// An optional description of the list.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Item []ListItemParameters `json:"item,omitempty" tf:"item,omitempty"`

	// (String) The type of items the list will contain.
	// The type of items the list will contain.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// (String) The name of the list. Modifying this attribute will force creation of a new resource.
	// The name of the list. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RedirectInitParameters struct {

	// (String) Whether the redirect also matches subdomains of the source url. Available values: disabled, enabled.
	// Whether the redirect also matches subdomains of the source url. Available values: `disabled`, `enabled`.
	IncludeSubdomains *string `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// (String) Whether to preserve the path suffix when doing subpath matching. Available values: disabled, enabled.
	// Whether to preserve the path suffix when doing subpath matching. Available values: `disabled`, `enabled`.
	PreservePathSuffix *string `json:"preservePathSuffix,omitempty" tf:"preserve_path_suffix,omitempty"`

	// (String) Whether the redirect target url should keep the query string of the request's url. Available values: disabled, enabled.
	// Whether the redirect target url should keep the query string of the request's url. Available values: `disabled`, `enabled`.
	PreserveQueryString *string `json:"preserveQueryString,omitempty" tf:"preserve_query_string,omitempty"`

	// (String) The source url of the redirect.
	// The source url of the redirect.
	SourceURL *string `json:"sourceUrl,omitempty" tf:"source_url,omitempty"`

	// (Number) The status code to be used when redirecting a request.
	// The status code to be used when redirecting a request.
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// (String) Whether the redirect also matches subpaths of the source url. Available values: disabled, enabled.
	// Whether the redirect also matches subpaths of the source url. Available values: `disabled`, `enabled`.
	SubpathMatching *string `json:"subpathMatching,omitempty" tf:"subpath_matching,omitempty"`

	// (String) The target url of the redirect.
	// The target url of the redirect.
	TargetURL *string `json:"targetUrl,omitempty" tf:"target_url,omitempty"`
}

type RedirectObservation struct {

	// (String) Whether the redirect also matches subdomains of the source url. Available values: disabled, enabled.
	// Whether the redirect also matches subdomains of the source url. Available values: `disabled`, `enabled`.
	IncludeSubdomains *string `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// (String) Whether to preserve the path suffix when doing subpath matching. Available values: disabled, enabled.
	// Whether to preserve the path suffix when doing subpath matching. Available values: `disabled`, `enabled`.
	PreservePathSuffix *string `json:"preservePathSuffix,omitempty" tf:"preserve_path_suffix,omitempty"`

	// (String) Whether the redirect target url should keep the query string of the request's url. Available values: disabled, enabled.
	// Whether the redirect target url should keep the query string of the request's url. Available values: `disabled`, `enabled`.
	PreserveQueryString *string `json:"preserveQueryString,omitempty" tf:"preserve_query_string,omitempty"`

	// (String) The source url of the redirect.
	// The source url of the redirect.
	SourceURL *string `json:"sourceUrl,omitempty" tf:"source_url,omitempty"`

	// (Number) The status code to be used when redirecting a request.
	// The status code to be used when redirecting a request.
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// (String) Whether the redirect also matches subpaths of the source url. Available values: disabled, enabled.
	// Whether the redirect also matches subpaths of the source url. Available values: `disabled`, `enabled`.
	SubpathMatching *string `json:"subpathMatching,omitempty" tf:"subpath_matching,omitempty"`

	// (String) The target url of the redirect.
	// The target url of the redirect.
	TargetURL *string `json:"targetUrl,omitempty" tf:"target_url,omitempty"`
}

type RedirectParameters struct {

	// (String) Whether the redirect also matches subdomains of the source url. Available values: disabled, enabled.
	// Whether the redirect also matches subdomains of the source url. Available values: `disabled`, `enabled`.
	// +kubebuilder:validation:Optional
	IncludeSubdomains *string `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// (String) Whether to preserve the path suffix when doing subpath matching. Available values: disabled, enabled.
	// Whether to preserve the path suffix when doing subpath matching. Available values: `disabled`, `enabled`.
	// +kubebuilder:validation:Optional
	PreservePathSuffix *string `json:"preservePathSuffix,omitempty" tf:"preserve_path_suffix,omitempty"`

	// (String) Whether the redirect target url should keep the query string of the request's url. Available values: disabled, enabled.
	// Whether the redirect target url should keep the query string of the request's url. Available values: `disabled`, `enabled`.
	// +kubebuilder:validation:Optional
	PreserveQueryString *string `json:"preserveQueryString,omitempty" tf:"preserve_query_string,omitempty"`

	// (String) The source url of the redirect.
	// The source url of the redirect.
	// +kubebuilder:validation:Optional
	SourceURL *string `json:"sourceUrl" tf:"source_url,omitempty"`

	// (Number) The status code to be used when redirecting a request.
	// The status code to be used when redirecting a request.
	// +kubebuilder:validation:Optional
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// (String) Whether the redirect also matches subpaths of the source url. Available values: disabled, enabled.
	// Whether the redirect also matches subpaths of the source url. Available values: `disabled`, `enabled`.
	// +kubebuilder:validation:Optional
	SubpathMatching *string `json:"subpathMatching,omitempty" tf:"subpath_matching,omitempty"`

	// (String) The target url of the redirect.
	// The target url of the redirect.
	// +kubebuilder:validation:Optional
	TargetURL *string `json:"targetUrl" tf:"target_url,omitempty"`
}

type ValueInitParameters struct {

	// (String)
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// (Block List) (see below for nested schema)
	Redirect []RedirectInitParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`
}

type ValueObservation struct {

	// (String)
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// (Block List) (see below for nested schema)
	Redirect []RedirectObservation `json:"redirect,omitempty" tf:"redirect,omitempty"`
}

type ValueParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Redirect []RedirectParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`
}

// ListSpec defines the desired state of List
type ListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListParameters `json:"forProvider"`
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
	InitProvider ListInitParameters `json:"initProvider,omitempty"`
}

// ListStatus defines the observed state of List.
type ListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// List is the Schema for the Lists API. Provides Lists (IPs, Redirects) to be used in Edge Rules Engine across all zones within the same account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type List struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kind) || (has(self.initProvider) && has(self.initProvider.kind))",message="spec.forProvider.kind is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ListSpec   `json:"spec"`
	Status ListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListList contains a list of Lists
type ListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []List `json:"items"`
}

// Repository type metadata.
var (
	List_Kind             = "List"
	List_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: List_Kind}.String()
	List_KindAPIVersion   = List_Kind + "." + CRDGroupVersion.String()
	List_GroupVersionKind = CRDGroupVersion.WithKind(List_Kind)
)

func init() {
	SchemeBuilder.Register(&List{}, &ListList{})
}
