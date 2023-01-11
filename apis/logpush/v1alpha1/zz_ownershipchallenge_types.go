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

type OwnershipChallengeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnershipChallengeFilename *string `json:"ownershipChallengeFilename,omitempty" tf:"ownership_challenge_filename,omitempty"`
}

type OwnershipChallengeParameters struct {

	// The account identifier to target for the resource. Must provide only one of `account_id`, `zone_id`.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	DestinationConf *string `json:"destinationConf" tf:"destination_conf,omitempty"`

	// The zone identifier to target for the resource. Must provide only one of `account_id`, `zone_id`.
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

// OwnershipChallengeSpec defines the desired state of OwnershipChallenge
type OwnershipChallengeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OwnershipChallengeParameters `json:"forProvider"`
}

// OwnershipChallengeStatus defines the observed state of OwnershipChallenge.
type OwnershipChallengeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OwnershipChallengeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OwnershipChallenge is the Schema for the OwnershipChallenges API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type OwnershipChallenge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OwnershipChallengeSpec   `json:"spec"`
	Status            OwnershipChallengeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OwnershipChallengeList contains a list of OwnershipChallenges
type OwnershipChallengeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OwnershipChallenge `json:"items"`
}

// Repository type metadata.
var (
	OwnershipChallenge_Kind             = "OwnershipChallenge"
	OwnershipChallenge_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OwnershipChallenge_Kind}.String()
	OwnershipChallenge_KindAPIVersion   = OwnershipChallenge_Kind + "." + CRDGroupVersion.String()
	OwnershipChallenge_GroupVersionKind = CRDGroupVersion.WithKind(OwnershipChallenge_Kind)
)

func init() {
	SchemeBuilder.Register(&OwnershipChallenge{}, &OwnershipChallengeList{})
}
