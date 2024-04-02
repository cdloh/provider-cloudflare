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

type TieredCacheInitParameters struct {

	// (String) The typed of tiered cache to utilize on the zone. Available values: generic, smart, off.
	// The typed of tiered cache to utilize on the zone. Available values: `generic`, `smart`, `off`.
	CacheType *string `json:"cacheType,omitempty" tf:"cache_type,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=Zone
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type TieredCacheObservation struct {

	// (String) The typed of tiered cache to utilize on the zone. Available values: generic, smart, off.
	// The typed of tiered cache to utilize on the zone. Available values: `generic`, `smart`, `off`.
	CacheType *string `json:"cacheType,omitempty" tf:"cache_type,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TieredCacheParameters struct {

	// (String) The typed of tiered cache to utilize on the zone. Available values: generic, smart, off.
	// The typed of tiered cache to utilize on the zone. Available values: `generic`, `smart`, `off`.
	// +kubebuilder:validation:Optional
	CacheType *string `json:"cacheType,omitempty" tf:"cache_type,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// TieredCacheSpec defines the desired state of TieredCache
type TieredCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TieredCacheParameters `json:"forProvider"`
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
	InitProvider TieredCacheInitParameters `json:"initProvider,omitempty"`
}

// TieredCacheStatus defines the observed state of TieredCache.
type TieredCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TieredCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TieredCache is the Schema for the TieredCaches API. Provides a resource, that manages Cloudflare Tiered Cache settings. This allows you to adjust topologies for your zone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type TieredCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cacheType) || (has(self.initProvider) && has(self.initProvider.cacheType))",message="spec.forProvider.cacheType is a required parameter"
	Spec   TieredCacheSpec   `json:"spec"`
	Status TieredCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TieredCacheList contains a list of TieredCaches
type TieredCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TieredCache `json:"items"`
}

// Repository type metadata.
var (
	TieredCache_Kind             = "TieredCache"
	TieredCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TieredCache_Kind}.String()
	TieredCache_KindAPIVersion   = TieredCache_Kind + "." + CRDGroupVersion.String()
	TieredCache_GroupVersionKind = CRDGroupVersion.WithKind(TieredCache_Kind)
)

func init() {
	SchemeBuilder.Register(&TieredCache{}, &TieredCacheList{})
}
