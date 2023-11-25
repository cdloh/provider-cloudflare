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

type KVInitParameters struct {

	// Name of the KV pair. **Modifying this attribute will force creation of a new resource.**
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of the KV pair.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type KVObservation struct {

	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the KV pair. **Modifying this attribute will force creation of a new resource.**
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The ID of the Workers KV namespace in which you want to create the KV pair. **Modifying this attribute will force creation of a new resource.**
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Value of the KV pair.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type KVParameters struct {

	// The account identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// Name of the KV pair. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The ID of the Workers KV namespace in which you want to create the KV pair. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=KVNamespace
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a KVNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a KVNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// Value of the KV pair.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// KVSpec defines the desired state of KV
type KVSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KVParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider KVInitParameters `json:"initProvider,omitempty"`
}

// KVStatus defines the observed state of KV.
type KVStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KVObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KV is the Schema for the KVs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type KV struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || has(self.initProvider.key)",message="key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || has(self.initProvider.value)",message="value is a required parameter"
	Spec   KVSpec   `json:"spec"`
	Status KVStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KVList contains a list of KVs
type KVList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KV `json:"items"`
}

// Repository type metadata.
var (
	KV_Kind             = "KV"
	KV_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KV_Kind}.String()
	KV_KindAPIVersion   = KV_Kind + "." + CRDGroupVersion.String()
	KV_GroupVersionKind = CRDGroupVersion.WithKind(KV_Kind)
)

func init() {
	SchemeBuilder.Register(&KV{}, &KVList{})
}
