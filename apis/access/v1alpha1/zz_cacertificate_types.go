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

type CACertificateInitParameters struct {

	// (String) The account identifier to target for the resource. Conflicts with zone_id.
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// (String) The Access Application ID to associate with the CA certificate.
	// The Access Application ID to associate with the CA certificate.
	// +crossplane:generate:reference:type=Application
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// (String) The zone identifier to target for the resource. Conflicts with account_id.
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type CACertificateObservation struct {

	// (String) The account identifier to target for the resource. Conflicts with zone_id.
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) The Access Application ID to associate with the CA certificate.
	// The Access Application ID to associate with the CA certificate.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String) Application Audience (AUD) Tag of the CA certificate.
	// Application Audience (AUD) Tag of the CA certificate.
	Aud *string `json:"aud,omitempty" tf:"aud,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Cryptographic public key of the generated CA certificate.
	// Cryptographic public key of the generated CA certificate.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// (String) The zone identifier to target for the resource. Conflicts with account_id.
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type CACertificateParameters struct {

	// (String) The account identifier to target for the resource. Conflicts with zone_id.
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// (String) The Access Application ID to associate with the CA certificate.
	// The Access Application ID to associate with the CA certificate.
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// (String) The zone identifier to target for the resource. Conflicts with account_id.
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// CACertificateSpec defines the desired state of CACertificate
type CACertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CACertificateParameters `json:"forProvider"`
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
	InitProvider CACertificateInitParameters `json:"initProvider,omitempty"`
}

// CACertificateStatus defines the observed state of CACertificate.
type CACertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CACertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CACertificate is the Schema for the CACertificates API. Cloudflare Access can replace traditional SSH key models with short-lived certificates issued to your users based on the token generated by their Access login.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type CACertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CACertificateSpec   `json:"spec"`
	Status            CACertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CACertificateList contains a list of CACertificates
type CACertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CACertificate `json:"items"`
}

// Repository type metadata.
var (
	CACertificate_Kind             = "CACertificate"
	CACertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CACertificate_Kind}.String()
	CACertificate_KindAPIVersion   = CACertificate_Kind + "." + CRDGroupVersion.String()
	CACertificate_GroupVersionKind = CRDGroupVersion.WithKind(CACertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&CACertificate{}, &CACertificateList{})
}
