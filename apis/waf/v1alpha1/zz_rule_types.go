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

type RuleInitParameters struct {

	// The mode of the rule, can be one of ["block", "challenge", "default", "disable", "simulate"] or ["on", "off"] depending on the WAF Rule type.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the WAF Rule Package that contains the rule.
	PackageID *string `json:"packageId,omitempty" tf:"package_id,omitempty"`

	// The WAF Rule ID.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// The DNS zone ID to apply to.
	// The zone identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type RuleObservation struct {

	// The ID of the WAF Rule Group that contains the rule.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The WAF Rule ID, the same as rule_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The mode of the rule, can be one of ["block", "challenge", "default", "disable", "simulate"] or ["on", "off"] depending on the WAF Rule type.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the WAF Rule Package that contains the rule.
	PackageID *string `json:"packageId,omitempty" tf:"package_id,omitempty"`

	// The WAF Rule ID.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// The DNS zone ID to apply to.
	// The zone identifier to target for the resource.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RuleParameters struct {

	// The mode of the rule, can be one of ["block", "challenge", "default", "disable", "simulate"] or ["on", "off"] depending on the WAF Rule type.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the WAF Rule Package that contains the rule.
	// +kubebuilder:validation:Optional
	PackageID *string `json:"packageId,omitempty" tf:"package_id,omitempty"`

	// The WAF Rule ID.
	// +kubebuilder:validation:Optional
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// The DNS zone ID to apply to.
	// The zone identifier to target for the resource.
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

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Provides a Cloudflare WAF rule resource for a particular zone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mode) || (has(self.initProvider) && has(self.initProvider.mode))",message="spec.forProvider.mode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleId) || (has(self.initProvider) && has(self.initProvider.ruleId))",message="spec.forProvider.ruleId is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}