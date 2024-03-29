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

type CronTriggerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CronTriggerParameters struct {

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

	// Cron expressions to execute the Worker script.
	// +kubebuilder:validation:Required
	Schedules []*string `json:"schedules" tf:"schedules,omitempty"`

	// Worker script to target for the schedules.
	// +crossplane:generate:reference:type=Script
	// +kubebuilder:validation:Optional
	ScriptName *string `json:"scriptName,omitempty" tf:"script_name,omitempty"`

	// Reference to a Script to populate scriptName.
	// +kubebuilder:validation:Optional
	ScriptNameRef *v1.Reference `json:"scriptNameRef,omitempty" tf:"-"`

	// Selector for a Script to populate scriptName.
	// +kubebuilder:validation:Optional
	ScriptNameSelector *v1.Selector `json:"scriptNameSelector,omitempty" tf:"-"`
}

// CronTriggerSpec defines the desired state of CronTrigger
type CronTriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CronTriggerParameters `json:"forProvider"`
}

// CronTriggerStatus defines the observed state of CronTrigger.
type CronTriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CronTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CronTrigger is the Schema for the CronTriggers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type CronTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CronTriggerSpec   `json:"spec"`
	Status            CronTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronTriggerList contains a list of CronTriggers
type CronTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronTrigger `json:"items"`
}

// Repository type metadata.
var (
	CronTrigger_Kind             = "CronTrigger"
	CronTrigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CronTrigger_Kind}.String()
	CronTrigger_KindAPIVersion   = CronTrigger_Kind + "." + CRDGroupVersion.String()
	CronTrigger_GroupVersionKind = CRDGroupVersion.WithKind(CronTrigger_Kind)
)

func init() {
	SchemeBuilder.Register(&CronTrigger{}, &CronTriggerList{})
}
