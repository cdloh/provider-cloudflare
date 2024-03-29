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

type AddressObservation struct {

	// The date and time the destination address has been created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time the destination address was last modified.
	Modified *string `json:"modified,omitempty" tf:"modified,omitempty"`

	// Destination address identifier.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The date and time the destination address has been verified. Null means not verified yet.
	Verified *string `json:"verified,omitempty" tf:"verified,omitempty"`
}

type AddressParameters struct {

	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// The contact email address of the user. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`
}

// AddressSpec defines the desired state of Address
type AddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressParameters `json:"forProvider"`
}

// AddressStatus defines the observed state of Address.
type AddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Address is the Schema for the Addresss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Address struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddressSpec   `json:"spec"`
	Status            AddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressList contains a list of Addresss
type AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Address `json:"items"`
}

// Repository type metadata.
var (
	Address_Kind             = "Address"
	Address_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Address_Kind}.String()
	Address_KindAPIVersion   = Address_Kind + "." + CRDGroupVersion.String()
	Address_GroupVersionKind = CRDGroupVersion.WithKind(Address_Kind)
)

func init() {
	SchemeBuilder.Register(&Address{}, &AddressList{})
}
