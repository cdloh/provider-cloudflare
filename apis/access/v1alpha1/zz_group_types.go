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

type AzureObservation struct {
}

type AzureParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type ExcludeObservation struct {
}

type ExcludeParameters struct {

	// +kubebuilder:validation:Optional
	AnyValidServiceToken *bool `json:"anyValidServiceToken,omitempty" tf:"any_valid_service_token,omitempty"`

	// +kubebuilder:validation:Optional
	AuthMethod *string `json:"authMethod,omitempty" tf:"auth_method,omitempty"`

	// +kubebuilder:validation:Optional
	Azure []AzureParameters `json:"azure,omitempty" tf:"azure,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate *bool `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// +kubebuilder:validation:Optional
	DevicePosture []*string `json:"devicePosture,omitempty" tf:"device_posture,omitempty"`

	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	EmailDomain []*string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	// +kubebuilder:validation:Optional
	Everyone *bool `json:"everyone,omitempty" tf:"everyone,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalEvaluation []ExternalEvaluationParameters `json:"externalEvaluation,omitempty" tf:"external_evaluation,omitempty"`

	// +kubebuilder:validation:Optional
	Geo []*string `json:"geo,omitempty" tf:"geo,omitempty"`

	// +kubebuilder:validation:Optional
	Github []GithubParameters `json:"github,omitempty" tf:"github,omitempty"`

	// +kubebuilder:validation:Optional
	Group []*string `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Optional
	Gsuite []GsuiteParameters `json:"gsuite,omitempty" tf:"gsuite,omitempty"`

	// +kubebuilder:validation:Optional
	IP []*string `json:"ip,omitempty" tf:"ip,omitempty"`

	// +kubebuilder:validation:Optional
	LoginMethod []*string `json:"loginMethod,omitempty" tf:"login_method,omitempty"`

	// +kubebuilder:validation:Optional
	Okta []OktaParameters `json:"okta,omitempty" tf:"okta,omitempty"`

	// +kubebuilder:validation:Optional
	SAML []SAMLParameters `json:"saml,omitempty" tf:"saml,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceToken []*string `json:"serviceToken,omitempty" tf:"service_token,omitempty"`
}

type ExternalEvaluationObservation struct {
}

type ExternalEvaluationParameters struct {

	// +kubebuilder:validation:Optional
	EvaluateURL *string `json:"evaluateUrl,omitempty" tf:"evaluate_url,omitempty"`

	// +kubebuilder:validation:Optional
	KeysURL *string `json:"keysUrl,omitempty" tf:"keys_url,omitempty"`
}

type GithubObservation struct {
}

type GithubParameters struct {

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Teams []*string `json:"teams,omitempty" tf:"teams,omitempty"`
}

type GroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GroupParameters struct {

	// The account identifier to target for the resource. Conflicts with `zone_id`. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Exclude []ExcludeParameters `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// +kubebuilder:validation:Required
	Include []IncludeParameters `json:"include" tf:"include,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Require []RequireParameters `json:"require,omitempty" tf:"require,omitempty"`

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

type GsuiteObservation struct {
}

type GsuiteParameters struct {

	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type IncludeAzureObservation struct {
}

type IncludeAzureParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type IncludeExternalEvaluationObservation struct {
}

type IncludeExternalEvaluationParameters struct {

	// +kubebuilder:validation:Optional
	EvaluateURL *string `json:"evaluateUrl,omitempty" tf:"evaluate_url,omitempty"`

	// +kubebuilder:validation:Optional
	KeysURL *string `json:"keysUrl,omitempty" tf:"keys_url,omitempty"`
}

type IncludeGithubObservation struct {
}

type IncludeGithubParameters struct {

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Teams []*string `json:"teams,omitempty" tf:"teams,omitempty"`
}

type IncludeGsuiteObservation struct {
}

type IncludeGsuiteParameters struct {

	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type IncludeObservation struct {
}

type IncludeOktaObservation struct {
}

type IncludeOktaParameters struct {

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name []*string `json:"name,omitempty" tf:"name,omitempty"`
}

type IncludeParameters struct {

	// +kubebuilder:validation:Optional
	AnyValidServiceToken *bool `json:"anyValidServiceToken,omitempty" tf:"any_valid_service_token,omitempty"`

	// +kubebuilder:validation:Optional
	AuthMethod *string `json:"authMethod,omitempty" tf:"auth_method,omitempty"`

	// +kubebuilder:validation:Optional
	Azure []IncludeAzureParameters `json:"azure,omitempty" tf:"azure,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate *bool `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// +kubebuilder:validation:Optional
	DevicePosture []*string `json:"devicePosture,omitempty" tf:"device_posture,omitempty"`

	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	EmailDomain []*string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	// +kubebuilder:validation:Optional
	Everyone *bool `json:"everyone,omitempty" tf:"everyone,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalEvaluation []IncludeExternalEvaluationParameters `json:"externalEvaluation,omitempty" tf:"external_evaluation,omitempty"`

	// +kubebuilder:validation:Optional
	Geo []*string `json:"geo,omitempty" tf:"geo,omitempty"`

	// +kubebuilder:validation:Optional
	Github []IncludeGithubParameters `json:"github,omitempty" tf:"github,omitempty"`

	// +kubebuilder:validation:Optional
	Group []*string `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Optional
	Gsuite []IncludeGsuiteParameters `json:"gsuite,omitempty" tf:"gsuite,omitempty"`

	// +kubebuilder:validation:Optional
	IP []*string `json:"ip,omitempty" tf:"ip,omitempty"`

	// +kubebuilder:validation:Optional
	LoginMethod []*string `json:"loginMethod,omitempty" tf:"login_method,omitempty"`

	// +kubebuilder:validation:Optional
	Okta []IncludeOktaParameters `json:"okta,omitempty" tf:"okta,omitempty"`

	// +kubebuilder:validation:Optional
	SAML []IncludeSAMLParameters `json:"saml,omitempty" tf:"saml,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceToken []*string `json:"serviceToken,omitempty" tf:"service_token,omitempty"`
}

type IncludeSAMLObservation struct {
}

type IncludeSAMLParameters struct {

	// +kubebuilder:validation:Optional
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// +kubebuilder:validation:Optional
	AttributeValue *string `json:"attributeValue,omitempty" tf:"attribute_value,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type OktaObservation struct {
}

type OktaParameters struct {

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name []*string `json:"name,omitempty" tf:"name,omitempty"`
}

type RequireAzureObservation struct {
}

type RequireAzureParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type RequireExternalEvaluationObservation struct {
}

type RequireExternalEvaluationParameters struct {

	// +kubebuilder:validation:Optional
	EvaluateURL *string `json:"evaluateUrl,omitempty" tf:"evaluate_url,omitempty"`

	// +kubebuilder:validation:Optional
	KeysURL *string `json:"keysUrl,omitempty" tf:"keys_url,omitempty"`
}

type RequireGithubObservation struct {
}

type RequireGithubParameters struct {

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Teams []*string `json:"teams,omitempty" tf:"teams,omitempty"`
}

type RequireGsuiteObservation struct {
}

type RequireGsuiteParameters struct {

	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type RequireObservation struct {
}

type RequireOktaObservation struct {
}

type RequireOktaParameters struct {

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name []*string `json:"name,omitempty" tf:"name,omitempty"`
}

type RequireParameters struct {

	// +kubebuilder:validation:Optional
	AnyValidServiceToken *bool `json:"anyValidServiceToken,omitempty" tf:"any_valid_service_token,omitempty"`

	// +kubebuilder:validation:Optional
	AuthMethod *string `json:"authMethod,omitempty" tf:"auth_method,omitempty"`

	// +kubebuilder:validation:Optional
	Azure []RequireAzureParameters `json:"azure,omitempty" tf:"azure,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate *bool `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// +kubebuilder:validation:Optional
	DevicePosture []*string `json:"devicePosture,omitempty" tf:"device_posture,omitempty"`

	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	EmailDomain []*string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	// +kubebuilder:validation:Optional
	Everyone *bool `json:"everyone,omitempty" tf:"everyone,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalEvaluation []RequireExternalEvaluationParameters `json:"externalEvaluation,omitempty" tf:"external_evaluation,omitempty"`

	// +kubebuilder:validation:Optional
	Geo []*string `json:"geo,omitempty" tf:"geo,omitempty"`

	// +kubebuilder:validation:Optional
	Github []RequireGithubParameters `json:"github,omitempty" tf:"github,omitempty"`

	// +kubebuilder:validation:Optional
	Group []*string `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Optional
	Gsuite []RequireGsuiteParameters `json:"gsuite,omitempty" tf:"gsuite,omitempty"`

	// +kubebuilder:validation:Optional
	IP []*string `json:"ip,omitempty" tf:"ip,omitempty"`

	// +kubebuilder:validation:Optional
	LoginMethod []*string `json:"loginMethod,omitempty" tf:"login_method,omitempty"`

	// +kubebuilder:validation:Optional
	Okta []RequireOktaParameters `json:"okta,omitempty" tf:"okta,omitempty"`

	// +kubebuilder:validation:Optional
	SAML []RequireSAMLParameters `json:"saml,omitempty" tf:"saml,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceToken []*string `json:"serviceToken,omitempty" tf:"service_token,omitempty"`
}

type RequireSAMLObservation struct {
}

type RequireSAMLParameters struct {

	// +kubebuilder:validation:Optional
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// +kubebuilder:validation:Optional
	AttributeValue *string `json:"attributeValue,omitempty" tf:"attribute_value,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

type SAMLObservation struct {
}

type SAMLParameters struct {

	// +kubebuilder:validation:Optional
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// +kubebuilder:validation:Optional
	AttributeValue *string `json:"attributeValue,omitempty" tf:"attribute_value,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
