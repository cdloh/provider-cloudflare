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

type RoomObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoomParameters struct {

	// This is a templated html file that will be rendered at the edge.
	// +kubebuilder:validation:Optional
	CustomPageHTML *string `json:"customPageHtml,omitempty" tf:"custom_page_html,omitempty"`

	// The language to use for the default waiting room page. Available values: `de-DE`, `es-ES`, `en-US`, `fr-FR`, `id-ID`, `it-IT`, `ja-JP`, `ko-KR`, `nl-NL`, `pl-PL`, `pt-BR`, `tr-TR`, `zh-CN`, `zh-TW`. Defaults to `en-US`.
	// +kubebuilder:validation:Optional
	DefaultTemplateLanguage *string `json:"defaultTemplateLanguage,omitempty" tf:"default_template_language,omitempty"`

	// A description to add more details about the waiting room.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disables automatic renewal of session cookies.
	// +kubebuilder:validation:Optional
	DisableSessionRenewal *bool `json:"disableSessionRenewal,omitempty" tf:"disable_session_renewal,omitempty"`

	// Host name for which the waiting room will be applied (no wildcards).
	// +kubebuilder:validation:Required
	Host *string `json:"host" tf:"host,omitempty"`

	// If true, requests to the waiting room with the header `Accept: application/json` will receive a JSON response object.
	// +kubebuilder:validation:Optional
	JSONResponseEnabled *bool `json:"jsonResponseEnabled,omitempty" tf:"json_response_enabled,omitempty"`

	// A unique name to identify the waiting room. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The number of new users that will be let into the route every minute.
	// +kubebuilder:validation:Required
	NewUsersPerMinute *float64 `json:"newUsersPerMinute" tf:"new_users_per_minute,omitempty"`

	// The path within the host to enable the waiting room on. Defaults to `/`.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// If queue_all is true, then all traffic will be sent to the waiting room.
	// +kubebuilder:validation:Optional
	QueueAll *bool `json:"queueAll,omitempty" tf:"queue_all,omitempty"`

	// The queueing method used by the waiting room. Available values: `fifo`, `random`, `passthrough`, `reject`. Defaults to `fifo`.
	// +kubebuilder:validation:Optional
	QueueingMethod *string `json:"queueingMethod,omitempty" tf:"queueing_method,omitempty"`

	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the origin. Defaults to `5`.
	// +kubebuilder:validation:Optional
	SessionDuration *float64 `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// Suspends the waiting room.
	// +kubebuilder:validation:Optional
	Suspended *bool `json:"suspended,omitempty" tf:"suspended,omitempty"`

	// The total number of active user sessions on the route at a point in time.
	// +kubebuilder:validation:Required
	TotalActiveUsers *float64 `json:"totalActiveUsers" tf:"total_active_users,omitempty"`

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

// RoomSpec defines the desired state of Room
type RoomSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoomParameters `json:"forProvider"`
}

// RoomStatus defines the observed state of Room.
type RoomStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoomObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Room is the Schema for the Rooms API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Room struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoomSpec   `json:"spec"`
	Status            RoomStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoomList contains a list of Rooms
type RoomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Room `json:"items"`
}

// Repository type metadata.
var (
	Room_Kind             = "Room"
	Room_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Room_Kind}.String()
	Room_KindAPIVersion   = Room_Kind + "." + CRDGroupVersion.String()
	Room_GroupVersionKind = CRDGroupVersion.WithKind(Room_Kind)
)

func init() {
	SchemeBuilder.Register(&Room{}, &RoomList{})
}