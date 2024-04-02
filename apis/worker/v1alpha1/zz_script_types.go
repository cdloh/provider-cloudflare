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

type AnalyticsEngineBindingInitParameters struct {

	// (String) The name of the Analytics Engine dataset to write to.
	// The name of the Analytics Engine dataset to write to.
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AnalyticsEngineBindingObservation struct {

	// (String) The name of the Analytics Engine dataset to write to.
	// The name of the Analytics Engine dataset to write to.
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AnalyticsEngineBindingParameters struct {

	// (String) The name of the Analytics Engine dataset to write to.
	// The name of the Analytics Engine dataset to write to.
	// +kubebuilder:validation:Optional
	Dataset *string `json:"dataset" tf:"dataset,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type KvNamespaceBindingInitParameters struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the KV namespace you want to use.
	// ID of the KV namespace you want to use.
	// +crossplane:generate:reference:type=KVNamespace
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a KVNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a KVNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`
}

type KvNamespaceBindingObservation struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the KV namespace you want to use.
	// ID of the KV namespace you want to use.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`
}

type KvNamespaceBindingParameters struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) ID of the KV namespace you want to use.
	// ID of the KV namespace you want to use.
	// +crossplane:generate:reference:type=KVNamespace
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a KVNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a KVNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`
}

type PlainTextBindingInitParameters struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The plain text you want to store.
	// The plain text you want to store.
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

type PlainTextBindingObservation struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The plain text you want to store.
	// The plain text you want to store.
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

type PlainTextBindingParameters struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The plain text you want to store.
	// The plain text you want to store.
	// +kubebuilder:validation:Optional
	Text *string `json:"text" tf:"text,omitempty"`
}

type R2BucketBindingInitParameters struct {

	// (String) The name of the Bucket to bind to.
	// The name of the Bucket to bind to.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type R2BucketBindingObservation struct {

	// (String) The name of the Bucket to bind to.
	// The name of the Bucket to bind to.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type R2BucketBindingParameters struct {

	// (String) The name of the Bucket to bind to.
	// The name of the Bucket to bind to.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ScriptInitParameters struct {

	// (String) The account identifier to target for the resource.
	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Block Set) (see below for nested schema)
	AnalyticsEngineBinding []AnalyticsEngineBindingInitParameters `json:"analyticsEngineBinding,omitempty" tf:"analytics_engine_binding,omitempty"`

	// (String) The script content.
	// The script content.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (Block Set) (see below for nested schema)
	KvNamespaceBinding []KvNamespaceBindingInitParameters `json:"kvNamespaceBinding,omitempty" tf:"kv_namespace_binding,omitempty"`

	// (Boolean) Whether to upload Worker as a module.
	// Whether to upload Worker as a module.
	Module *bool `json:"module,omitempty" tf:"module,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The name for the script. **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set) (see below for nested schema)
	PlainTextBinding []PlainTextBindingInitParameters `json:"plainTextBinding,omitempty" tf:"plain_text_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	R2BucketBinding []R2BucketBindingInitParameters `json:"r2BucketBinding,omitempty" tf:"r2_bucket_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	SecretTextBinding []SecretTextBindingInitParameters `json:"secretTextBinding,omitempty" tf:"secret_text_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	ServiceBinding []ServiceBindingInitParameters `json:"serviceBinding,omitempty" tf:"service_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	WebassemblyBinding []WebassemblyBindingInitParameters `json:"webassemblyBinding,omitempty" tf:"webassembly_binding,omitempty"`
}

type ScriptObservation struct {

	// (String) The account identifier to target for the resource.
	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Block Set) (see below for nested schema)
	AnalyticsEngineBinding []AnalyticsEngineBindingObservation `json:"analyticsEngineBinding,omitempty" tf:"analytics_engine_binding,omitempty"`

	// (String) The script content.
	// The script content.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block Set) (see below for nested schema)
	KvNamespaceBinding []KvNamespaceBindingObservation `json:"kvNamespaceBinding,omitempty" tf:"kv_namespace_binding,omitempty"`

	// (Boolean) Whether to upload Worker as a module.
	// Whether to upload Worker as a module.
	Module *bool `json:"module,omitempty" tf:"module,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The name for the script. **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set) (see below for nested schema)
	PlainTextBinding []PlainTextBindingObservation `json:"plainTextBinding,omitempty" tf:"plain_text_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	R2BucketBinding []R2BucketBindingObservation `json:"r2BucketBinding,omitempty" tf:"r2_bucket_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	SecretTextBinding []SecretTextBindingObservation `json:"secretTextBinding,omitempty" tf:"secret_text_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	ServiceBinding []ServiceBindingObservation `json:"serviceBinding,omitempty" tf:"service_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	WebassemblyBinding []WebassemblyBindingObservation `json:"webassemblyBinding,omitempty" tf:"webassembly_binding,omitempty"`
}

type ScriptParameters struct {

	// (String) The account identifier to target for the resource.
	// The account identifier to target for the resource.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	AnalyticsEngineBinding []AnalyticsEngineBindingParameters `json:"analyticsEngineBinding,omitempty" tf:"analytics_engine_binding,omitempty"`

	// (String) The script content.
	// The script content.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	KvNamespaceBinding []KvNamespaceBindingParameters `json:"kvNamespaceBinding,omitempty" tf:"kv_namespace_binding,omitempty"`

	// (Boolean) Whether to upload Worker as a module.
	// Whether to upload Worker as a module.
	// +kubebuilder:validation:Optional
	Module *bool `json:"module,omitempty" tf:"module,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The name for the script. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	PlainTextBinding []PlainTextBindingParameters `json:"plainTextBinding,omitempty" tf:"plain_text_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	R2BucketBinding []R2BucketBindingParameters `json:"r2BucketBinding,omitempty" tf:"r2_bucket_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	SecretTextBinding []SecretTextBindingParameters `json:"secretTextBinding,omitempty" tf:"secret_text_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ServiceBinding []ServiceBindingParameters `json:"serviceBinding,omitempty" tf:"service_binding,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	WebassemblyBinding []WebassemblyBindingParameters `json:"webassemblyBinding,omitempty" tf:"webassembly_binding,omitempty"`
}

type SecretTextBindingInitParameters struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecretTextBindingObservation struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecretTextBindingParameters struct {

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The plain text you want to store.
	// The secret text you want to store.
	// +kubebuilder:validation:Required
	TextSecretRef v1.SecretKeySelector `json:"textSecretRef" tf:"-"`
}

type ServiceBindingInitParameters struct {

	// (String) The name of the Worker environment to bind to.
	// The name of the Worker environment to bind to.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The name of the Worker to bind to.
	// The name of the Worker to bind to.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceBindingObservation struct {

	// (String) The name of the Worker environment to bind to.
	// The name of the Worker environment to bind to.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The name of the Worker to bind to.
	// The name of the Worker to bind to.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceBindingParameters struct {

	// (String) The name of the Worker environment to bind to.
	// The name of the Worker environment to bind to.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The name of the Worker to bind to.
	// The name of the Worker to bind to.
	// +kubebuilder:validation:Optional
	Service *string `json:"service" tf:"service,omitempty"`
}

type WebassemblyBindingInitParameters struct {

	// (Boolean) Whether to upload Worker as a module.
	// The base64 encoded wasm module you want to store.
	Module *string `json:"module,omitempty" tf:"module,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type WebassemblyBindingObservation struct {

	// (Boolean) Whether to upload Worker as a module.
	// The base64 encoded wasm module you want to store.
	Module *string `json:"module,omitempty" tf:"module,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type WebassemblyBindingParameters struct {

	// (Boolean) Whether to upload Worker as a module.
	// The base64 encoded wasm module you want to store.
	// +kubebuilder:validation:Optional
	Module *string `json:"module" tf:"module,omitempty"`

	// (String) The name for the script. Modifying this attribute will force creation of a new resource.
	// The global variable for the binding in your Worker code.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// ScriptSpec defines the desired state of Script
type ScriptSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScriptParameters `json:"forProvider"`
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
	InitProvider ScriptInitParameters `json:"initProvider,omitempty"`
}

// ScriptStatus defines the observed state of Script.
type ScriptStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScriptObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Script is the Schema for the Scripts API. Provides a Cloudflare worker script resource. In order for a script to be active, you'll also need to setup a cloudflare_worker_route.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Script struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ScriptSpec   `json:"spec"`
	Status ScriptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScriptList contains a list of Scripts
type ScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Script `json:"items"`
}

// Repository type metadata.
var (
	Script_Kind             = "Script"
	Script_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Script_Kind}.String()
	Script_KindAPIVersion   = Script_Kind + "." + CRDGroupVersion.String()
	Script_GroupVersionKind = CRDGroupVersion.WithKind(Script_Kind)
)

func init() {
	SchemeBuilder.Register(&Script{}, &ScriptList{})
}
