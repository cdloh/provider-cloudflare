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

type MutualTLSCertificateObservation struct {
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MutualTLSCertificateParameters struct {

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

	// The hostnames that will be prompted for this certificate.
	// +kubebuilder:validation:Optional
	AssociatedHostnames []*string `json:"associatedHostnames,omitempty" tf:"associated_hostnames,omitempty"`

	// The Root CA for your certificates.
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// The name of the certificate.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

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

// MutualTLSCertificateSpec defines the desired state of MutualTLSCertificate
type MutualTLSCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MutualTLSCertificateParameters `json:"forProvider"`
}

// MutualTLSCertificateStatus defines the observed state of MutualTLSCertificate.
type MutualTLSCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MutualTLSCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MutualTLSCertificate is the Schema for the MutualTLSCertificates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type MutualTLSCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MutualTLSCertificateSpec   `json:"spec"`
	Status            MutualTLSCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MutualTLSCertificateList contains a list of MutualTLSCertificates
type MutualTLSCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MutualTLSCertificate `json:"items"`
}

// Repository type metadata.
var (
	MutualTLSCertificate_Kind             = "MutualTLSCertificate"
	MutualTLSCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MutualTLSCertificate_Kind}.String()
	MutualTLSCertificate_KindAPIVersion   = MutualTLSCertificate_Kind + "." + CRDGroupVersion.String()
	MutualTLSCertificate_GroupVersionKind = CRDGroupVersion.WithKind(MutualTLSCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&MutualTLSCertificate{}, &MutualTLSCertificateList{})
}
