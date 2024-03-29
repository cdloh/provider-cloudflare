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

type PackObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	ValidationErrors []ValidationErrorsObservation `json:"validationErrors,omitempty" tf:"validation_errors,omitempty"`
}

type PackParameters struct {

	// Which certificate authority to issue the certificate pack. Available values: `digicert`, `lets_encrypt`, `google`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	CertificateAuthority *string `json:"certificateAuthority" tf:"certificate_authority,omitempty"`

	// Whether or not to include Cloudflare branding. This will add `sni.cloudflaressl.com` as the Common Name if set to `true`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	CloudflareBranding *bool `json:"cloudflareBranding,omitempty" tf:"cloudflare_branding,omitempty"`

	// List of hostnames to provision the certificate pack for. The zone name must be included as a host. Note: If using Let's Encrypt, you cannot use individual subdomains and only a wildcard for subdomain is available. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	Hosts []*string `json:"hosts" tf:"hosts,omitempty"`

	// Certificate pack configuration type. Available values: `advanced`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	ValidationErrors []ValidationErrorsParameters `json:"validationErrors,omitempty" tf:"validation_errors,omitempty"`

	// Which validation method to use in order to prove domain ownership. Available values: `txt`, `http`, `email`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	ValidationMethod *string `json:"validationMethod" tf:"validation_method,omitempty"`

	// +kubebuilder:validation:Optional
	ValidationRecords []ValidationRecordsParameters `json:"validationRecords,omitempty" tf:"validation_records,omitempty"`

	// How long the certificate is valid for. Note: If using Let's Encrypt, this value can only be 90 days. Available values: `14`, `30`, `90`, `365`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	ValidityDays *float64 `json:"validityDays" tf:"validity_days,omitempty"`

	// Whether or not to wait for a certificate pack to reach status `active` during creation. Defaults to `false`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	WaitForActiveStatus *bool `json:"waitForActiveStatus,omitempty" tf:"wait_for_active_status,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
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

type ValidationErrorsObservation struct {
	Message *string `json:"message,omitempty" tf:"message,omitempty"`
}

type ValidationErrorsParameters struct {
}

type ValidationRecordsObservation struct {
}

type ValidationRecordsParameters struct {

	// +kubebuilder:validation:Optional
	CnameName *string `json:"cnameName,omitempty" tf:"cname_name,omitempty"`

	// +kubebuilder:validation:Optional
	CnameTarget *string `json:"cnameTarget,omitempty" tf:"cname_target,omitempty"`

	// +kubebuilder:validation:Optional
	Emails []*string `json:"emails,omitempty" tf:"emails,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPBody *string `json:"httpBody,omitempty" tf:"http_body,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPURL *string `json:"httpUrl,omitempty" tf:"http_url,omitempty"`

	// +kubebuilder:validation:Optional
	TxtName *string `json:"txtName,omitempty" tf:"txt_name,omitempty"`

	// +kubebuilder:validation:Optional
	TxtValue *string `json:"txtValue,omitempty" tf:"txt_value,omitempty"`
}

// PackSpec defines the desired state of Pack
type PackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PackParameters `json:"forProvider"`
}

// PackStatus defines the observed state of Pack.
type PackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pack is the Schema for the Packs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Pack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PackSpec   `json:"spec"`
	Status            PackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PackList contains a list of Packs
type PackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pack `json:"items"`
}

// Repository type metadata.
var (
	Pack_Kind             = "Pack"
	Pack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pack_Kind}.String()
	Pack_KindAPIVersion   = Pack_Kind + "." + CRDGroupVersion.String()
	Pack_GroupVersionKind = CRDGroupVersion.WithKind(Pack_Kind)
)

func init() {
	SchemeBuilder.Register(&Pack{}, &PackList{})
}
