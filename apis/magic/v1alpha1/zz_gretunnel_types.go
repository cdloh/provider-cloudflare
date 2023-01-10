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

type GRETunnelObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GRETunnelParameters struct {

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

	// +kubebuilder:validation:Required
	CloudflareGreEndpoint *string `json:"cloudflareGreEndpoint" tf:"cloudflare_gre_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	CustomerGreEndpoint *string `json:"customerGreEndpoint" tf:"customer_gre_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckTarget *string `json:"healthCheckTarget,omitempty" tf:"health_check_target,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// +kubebuilder:validation:Required
	InterfaceAddress *string `json:"interfaceAddress" tf:"interface_address,omitempty"`

	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// GRETunnelSpec defines the desired state of GRETunnel
type GRETunnelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GRETunnelParameters `json:"forProvider"`
}

// GRETunnelStatus defines the observed state of GRETunnel.
type GRETunnelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GRETunnelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GRETunnel is the Schema for the GRETunnels API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type GRETunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GRETunnelSpec   `json:"spec"`
	Status            GRETunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GRETunnelList contains a list of GRETunnels
type GRETunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GRETunnel `json:"items"`
}

// Repository type metadata.
var (
	GRETunnel_Kind             = "GRETunnel"
	GRETunnel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GRETunnel_Kind}.String()
	GRETunnel_KindAPIVersion   = GRETunnel_Kind + "." + CRDGroupVersion.String()
	GRETunnel_GroupVersionKind = CRDGroupVersion.WithKind(GRETunnel_Kind)
)

func init() {
	SchemeBuilder.Register(&GRETunnel{}, &GRETunnelList{})
}
