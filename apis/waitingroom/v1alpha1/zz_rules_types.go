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

type RulesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of rules to apply to the ruleset.
	// +kubebuilder:validation:Optional
	Rules []RulesRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`
}

type RulesParameters struct {

	// List of rules to apply to the ruleset.
	// +kubebuilder:validation:Optional
	Rules []RulesRulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// The Waiting Room ID the rules should apply to. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=Room
	// +kubebuilder:validation:Optional
	WaitingRoomID *string `json:"waitingRoomId,omitempty" tf:"waiting_room_id,omitempty"`

	// Reference to a Room to populate waitingRoomId.
	// +kubebuilder:validation:Optional
	WaitingRoomIDRef *v1.Reference `json:"waitingRoomIdRef,omitempty" tf:"-"`

	// Selector for a Room to populate waitingRoomId.
	// +kubebuilder:validation:Optional
	WaitingRoomIDSelector *v1.Selector `json:"waitingRoomIdSelector,omitempty" tf:"-"`

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

type RulesRulesObservation struct {

	// Unique rule identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Version of the waiting room rule.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RulesRulesParameters struct {

	// Action to perform in the ruleset rule. Available values: `bypass_waiting_room`.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Brief summary of the waiting room rule and its intended use.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Criteria for an HTTP request to trigger the waiting room rule action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Waiting Room Rules Docs](https://developers.cloudflare.com/waiting-room/additional-options/waiting-room-rules/bypass-rules/).
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// Whether the rule is enabled or disabled. Available values: `enabled`, `disabled`.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// RulesSpec defines the desired state of Rules
type RulesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RulesParameters `json:"forProvider"`
}

// RulesStatus defines the observed state of Rules.
type RulesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RulesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rules is the Schema for the Ruless API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Rules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RulesSpec   `json:"spec"`
	Status            RulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RulesList contains a list of Ruless
type RulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rules `json:"items"`
}

// Repository type metadata.
var (
	Rules_Kind             = "Rules"
	Rules_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rules_Kind}.String()
	Rules_KindAPIVersion   = Rules_Kind + "." + CRDGroupVersion.String()
	Rules_GroupVersionKind = CRDGroupVersion.WithKind(Rules_Kind)
)

func init() {
	SchemeBuilder.Register(&Rules{}, &RulesList{})
}
