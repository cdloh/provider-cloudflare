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

type ServiceTokenObservation struct {

	// UUID client ID associated with the Service Token. **Modifying this attribute will force creation of a new resource.**
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Date when the token expires.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceTokenParameters struct {

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

	// Defaults to `0`.
	// +kubebuilder:validation:Optional
	MinDaysForRenewal *float64 `json:"minDaysForRenewal,omitempty" tf:"min_days_for_renewal,omitempty"`

	// Friendly name of the token's intent.
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

// ServiceTokenSpec defines the desired state of ServiceToken
type ServiceTokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceTokenParameters `json:"forProvider"`
}

// ServiceTokenStatus defines the observed state of ServiceToken.
type ServiceTokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceToken is the Schema for the ServiceTokens API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type ServiceToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceTokenSpec   `json:"spec"`
	Status            ServiceTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceTokenList contains a list of ServiceTokens
type ServiceTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceToken `json:"items"`
}

// Repository type metadata.
var (
	ServiceToken_Kind             = "ServiceToken"
	ServiceToken_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceToken_Kind}.String()
	ServiceToken_KindAPIVersion   = ServiceToken_Kind + "." + CRDGroupVersion.String()
	ServiceToken_GroupVersionKind = CRDGroupVersion.WithKind(ServiceToken_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceToken{}, &ServiceTokenList{})
}
