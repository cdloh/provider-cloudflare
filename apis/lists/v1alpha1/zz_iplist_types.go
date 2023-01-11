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

type IPListObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IPListParameters struct {

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

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Item []ItemParameters `json:"item,omitempty" tf:"item,omitempty"`

	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ItemObservation struct {
}

type ItemParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// IPListSpec defines the desired state of IPList
type IPListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPListParameters `json:"forProvider"`
}

// IPListStatus defines the observed state of IPList.
type IPListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPList is the Schema for the IPLists API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type IPList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPListSpec   `json:"spec"`
	Status            IPListStatus `json:"status,omitempty"`
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