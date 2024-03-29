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

type DeviceSettingsPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DeviceSettingsPolicyParameters struct {

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

	// Whether to allow mode switch for this policy.
	// +kubebuilder:validation:Optional
	AllowModeSwitch *bool `json:"allowModeSwitch,omitempty" tf:"allow_mode_switch,omitempty"`

	// Whether to allow updates under this policy.
	// +kubebuilder:validation:Optional
	AllowUpdates *bool `json:"allowUpdates,omitempty" tf:"allow_updates,omitempty"`

	// Whether to allow devices to leave the organization. Defaults to `true`.
	// +kubebuilder:validation:Optional
	AllowedToLeave *bool `json:"allowedToLeave,omitempty" tf:"allowed_to_leave,omitempty"`

	// The amount of time in minutes to reconnect after having been disabled.
	// +kubebuilder:validation:Optional
	AutoConnect *float64 `json:"autoConnect,omitempty" tf:"auto_connect,omitempty"`

	// The captive portal value for this policy. Defaults to `180`.
	// +kubebuilder:validation:Optional
	CaptivePortal *float64 `json:"captivePortal,omitempty" tf:"captive_portal,omitempty"`

	// Whether the policy refers to the default account policy.
	// +kubebuilder:validation:Optional
	Default *bool `json:"default,omitempty" tf:"default,omitempty"`

	// Whether to disable auto fallback for this policy.
	// +kubebuilder:validation:Optional
	DisableAutoFallback *bool `json:"disableAutoFallback,omitempty" tf:"disable_auto_fallback,omitempty"`

	// Whether the policy is enabled (cannot be set for default policies). Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Wirefilter expression to match a device against when evaluating whether this policy should take effect for that device.
	// +kubebuilder:validation:Optional
	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	// Name of the policy.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The precedence of the policy. Lower values indicate higher precedence.
	// +kubebuilder:validation:Optional
	Precedence *float64 `json:"precedence,omitempty" tf:"precedence,omitempty"`

	// The service mode. Defaults to `warp`.
	// +kubebuilder:validation:Optional
	ServiceModeV2Mode *string `json:"serviceModeV2Mode,omitempty" tf:"service_mode_v2_mode,omitempty"`

	// The port to use for the proxy service mode. Required when using `service_mode_v2_mode`.
	// +kubebuilder:validation:Optional
	ServiceModeV2Port *float64 `json:"serviceModeV2Port,omitempty" tf:"service_mode_v2_port,omitempty"`

	// The support URL that will be opened when sending feedback.
	// +kubebuilder:validation:Optional
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`

	// Enablement of the ZT client switch lock.
	// +kubebuilder:validation:Optional
	SwitchLocked *bool `json:"switchLocked,omitempty" tf:"switch_locked,omitempty"`
}

// DeviceSettingsPolicySpec defines the desired state of DeviceSettingsPolicy
type DeviceSettingsPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceSettingsPolicyParameters `json:"forProvider"`
}

// DeviceSettingsPolicyStatus defines the observed state of DeviceSettingsPolicy.
type DeviceSettingsPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceSettingsPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceSettingsPolicy is the Schema for the DeviceSettingsPolicys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type DeviceSettingsPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeviceSettingsPolicySpec   `json:"spec"`
	Status            DeviceSettingsPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceSettingsPolicyList contains a list of DeviceSettingsPolicys
type DeviceSettingsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeviceSettingsPolicy `json:"items"`
}

// Repository type metadata.
var (
	DeviceSettingsPolicy_Kind             = "DeviceSettingsPolicy"
	DeviceSettingsPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeviceSettingsPolicy_Kind}.String()
	DeviceSettingsPolicy_KindAPIVersion   = DeviceSettingsPolicy_Kind + "." + CRDGroupVersion.String()
	DeviceSettingsPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DeviceSettingsPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DeviceSettingsPolicy{}, &DeviceSettingsPolicyList{})
}
