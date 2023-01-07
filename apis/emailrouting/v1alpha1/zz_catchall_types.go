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

type ActionObservation struct {
}

type ActionParameters struct {

	// Type of supported action. Available values: `drop`, `forward`, `worker`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// A list with items in the following form.
	// +kubebuilder:validation:Required
	Value []*string `json:"value" tf:"value,omitempty"`
}

type CatchAllObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Routing rule identifier.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type CatchAllParameters struct {

	// List actions patterns.
	// +kubebuilder:validation:Required
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// Routing rule status.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Matching patterns to forward to your actions.
	// +kubebuilder:validation:Required
	Matcher []MatcherParameters `json:"matcher" tf:"matcher,omitempty"`

	// Routing rule name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The zone identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone
	// +crossplane:generate:reference:refFieldName=ZoneIDRefs
	// +crossplane:generate:reference:selectorFieldName=ZoneIDSelector
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRefs *v1.Reference `json:"zoneIdRefs,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type MatcherObservation struct {
}

type MatcherParameters struct {

	// Type of matcher. Available values: `all`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// CatchAllSpec defines the desired state of CatchAll
type CatchAllSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatchAllParameters `json:"forProvider"`
}

// CatchAllStatus defines the observed state of CatchAll.
type CatchAllStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatchAllObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CatchAll is the Schema for the CatchAlls API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type CatchAll struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatchAllSpec   `json:"spec"`
	Status            CatchAllStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatchAllList contains a list of CatchAlls
type CatchAllList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatchAll `json:"items"`
}

// Repository type metadata.
var (
	CatchAll_Kind             = "CatchAll"
	CatchAll_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatchAll_Kind}.String()
	CatchAll_KindAPIVersion   = CatchAll_Kind + "." + CRDGroupVersion.String()
	CatchAll_GroupVersionKind = CRDGroupVersion.WithKind(CatchAll_Kind)
)

func init() {
	SchemeBuilder.Register(&CatchAll{}, &CatchAllList{})
}
