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

type ApplicationInitParameters struct {

	// The identity providers selected for the application.
	AllowedIdps []*string `json:"allowedIdps,omitempty" tf:"allowed_idps,omitempty"`

	// Option to show/hide applications in App Launcher. Defaults to `true`.
	AppLauncherVisible *bool `json:"appLauncherVisible,omitempty" tf:"app_launcher_visible,omitempty"`

	// Option to skip identity provider selection if only one is configured in `allowed_idps`. Defaults to `false`.
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// CORS configuration for the Access Application. See below for reference structure.
	CorsHeaders []CorsHeadersInitParameters `json:"corsHeaders,omitempty" tf:"cors_headers,omitempty"`

	// Option that returns a custom error message when a user is denied access to the application.
	CustomDenyMessage *string `json:"customDenyMessage,omitempty" tf:"custom_deny_message,omitempty"`

	// Option that redirects to a custom URL when a user is denied access to the application via identity based rules.
	CustomDenyURL *string `json:"customDenyUrl,omitempty" tf:"custom_deny_url,omitempty"`

	// Option that redirects to a custom URL when a user is denied access to the application via non identity rules.
	CustomNonIdentityDenyURL *string `json:"customNonIdentityDenyUrl,omitempty" tf:"custom_non_identity_deny_url,omitempty"`

	// The custom pages selected for the application.
	CustomPages []*string `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// The primary hostname and path that Access will secure. If the app is visible in the App Launcher dashboard, this is the domain that will be displayed.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional "binding" cookie on requests. Defaults to `false`.
	EnableBindingCookie *bool `json:"enableBindingCookie,omitempty" tf:"enable_binding_cookie,omitempty"`

	// Option to add the `HttpOnly` cookie flag to access tokens.
	HTTPOnlyCookieAttribute *bool `json:"httpOnlyCookieAttribute,omitempty" tf:"http_only_cookie_attribute,omitempty"`

	// Image URL for the logo shown in the app launcher dashboard.
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`

	// Friendly name of the Access Application.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// SaaS configuration for the Access Application.
	SaasApp []SaasAppInitParameters `json:"saasApp,omitempty" tf:"saas_app,omitempty"`

	// Defines the same-site cookie setting for access tokens. Available values: `none`, `lax`, `strict`.
	SameSiteCookieAttribute *string `json:"sameSiteCookieAttribute,omitempty" tf:"same_site_cookie_attribute,omitempty"`

	// List of domains that access will secure. Only present for self_hosted, vnc, and ssh applications. Always includes the value set as `domain`.
	SelfHostedDomains []*string `json:"selfHostedDomains,omitempty" tf:"self_hosted_domains,omitempty"`

	// Option to return a 401 status code in service authentication rules on failed requests. Defaults to `false`.
	ServiceAuth401Redirect *bool `json:"serviceAuth401Redirect,omitempty" tf:"service_auth_401_redirect,omitempty"`

	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`. Defaults to `24h`.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// Option to skip the authorization interstitial when using the CLI. Defaults to `false`.
	SkipInterstitial *bool `json:"skipInterstitial,omitempty" tf:"skip_interstitial,omitempty"`

	// The itags associated with the application.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The application type. Available values: `app_launcher`, `bookmark`, `biso`, `dash_sso`, `saas`, `self_hosted`, `ssh`, `vnc`, `warp`. Defaults to `self_hosted`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ApplicationObservation struct {

	// The account identifier to target for the resource. Conflicts with `zone_id`.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The identity providers selected for the application.
	AllowedIdps []*string `json:"allowedIdps,omitempty" tf:"allowed_idps,omitempty"`

	// Option to show/hide applications in App Launcher. Defaults to `true`.
	AppLauncherVisible *bool `json:"appLauncherVisible,omitempty" tf:"app_launcher_visible,omitempty"`

	// Application Audience (AUD) Tag of the application.
	Aud *string `json:"aud,omitempty" tf:"aud,omitempty"`

	// Option to skip identity provider selection if only one is configured in `allowed_idps`. Defaults to `false`.
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// CORS configuration for the Access Application. See below for reference structure.
	CorsHeaders []CorsHeadersObservation `json:"corsHeaders,omitempty" tf:"cors_headers,omitempty"`

	// Option that returns a custom error message when a user is denied access to the application.
	CustomDenyMessage *string `json:"customDenyMessage,omitempty" tf:"custom_deny_message,omitempty"`

	// Option that redirects to a custom URL when a user is denied access to the application via identity based rules.
	CustomDenyURL *string `json:"customDenyUrl,omitempty" tf:"custom_deny_url,omitempty"`

	// Option that redirects to a custom URL when a user is denied access to the application via non identity rules.
	CustomNonIdentityDenyURL *string `json:"customNonIdentityDenyUrl,omitempty" tf:"custom_non_identity_deny_url,omitempty"`

	// The custom pages selected for the application.
	CustomPages []*string `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// The primary hostname and path that Access will secure. If the app is visible in the App Launcher dashboard, this is the domain that will be displayed.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional "binding" cookie on requests. Defaults to `false`.
	EnableBindingCookie *bool `json:"enableBindingCookie,omitempty" tf:"enable_binding_cookie,omitempty"`

	// Option to add the `HttpOnly` cookie flag to access tokens.
	HTTPOnlyCookieAttribute *bool `json:"httpOnlyCookieAttribute,omitempty" tf:"http_only_cookie_attribute,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Image URL for the logo shown in the app launcher dashboard.
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`

	// Friendly name of the Access Application.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// SaaS configuration for the Access Application.
	SaasApp []SaasAppObservation `json:"saasApp,omitempty" tf:"saas_app,omitempty"`

	// Defines the same-site cookie setting for access tokens. Available values: `none`, `lax`, `strict`.
	SameSiteCookieAttribute *string `json:"sameSiteCookieAttribute,omitempty" tf:"same_site_cookie_attribute,omitempty"`

	// List of domains that access will secure. Only present for self_hosted, vnc, and ssh applications. Always includes the value set as `domain`.
	SelfHostedDomains []*string `json:"selfHostedDomains,omitempty" tf:"self_hosted_domains,omitempty"`

	// Option to return a 401 status code in service authentication rules on failed requests. Defaults to `false`.
	ServiceAuth401Redirect *bool `json:"serviceAuth401Redirect,omitempty" tf:"service_auth_401_redirect,omitempty"`

	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`. Defaults to `24h`.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// Option to skip the authorization interstitial when using the CLI. Defaults to `false`.
	SkipInterstitial *bool `json:"skipInterstitial,omitempty" tf:"skip_interstitial,omitempty"`

	// The itags associated with the application.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The application type. Available values: `app_launcher`, `bookmark`, `biso`, `dash_sso`, `saas`, `self_hosted`, `ssh`, `vnc`, `warp`. Defaults to `self_hosted`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The zone identifier to target for the resource. Conflicts with `account_id`.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ApplicationParameters struct {

	// The account identifier to target for the resource. Conflicts with `zone_id`.
	// +crossplane:generate:reference:type=github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// The identity providers selected for the application.
	// +kubebuilder:validation:Optional
	AllowedIdps []*string `json:"allowedIdps,omitempty" tf:"allowed_idps,omitempty"`

	// Option to show/hide applications in App Launcher. Defaults to `true`.
	// +kubebuilder:validation:Optional
	AppLauncherVisible *bool `json:"appLauncherVisible,omitempty" tf:"app_launcher_visible,omitempty"`

	// Option to skip identity provider selection if only one is configured in `allowed_idps`. Defaults to `false`.
	// +kubebuilder:validation:Optional
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// CORS configuration for the Access Application. See below for reference structure.
	// +kubebuilder:validation:Optional
	CorsHeaders []CorsHeadersParameters `json:"corsHeaders,omitempty" tf:"cors_headers,omitempty"`

	// Option that returns a custom error message when a user is denied access to the application.
	// +kubebuilder:validation:Optional
	CustomDenyMessage *string `json:"customDenyMessage,omitempty" tf:"custom_deny_message,omitempty"`

	// Option that redirects to a custom URL when a user is denied access to the application via identity based rules.
	// +kubebuilder:validation:Optional
	CustomDenyURL *string `json:"customDenyUrl,omitempty" tf:"custom_deny_url,omitempty"`

	// Option that redirects to a custom URL when a user is denied access to the application via non identity rules.
	// +kubebuilder:validation:Optional
	CustomNonIdentityDenyURL *string `json:"customNonIdentityDenyUrl,omitempty" tf:"custom_non_identity_deny_url,omitempty"`

	// The custom pages selected for the application.
	// +kubebuilder:validation:Optional
	CustomPages []*string `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// The primary hostname and path that Access will secure. If the app is visible in the App Launcher dashboard, this is the domain that will be displayed.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional "binding" cookie on requests. Defaults to `false`.
	// +kubebuilder:validation:Optional
	EnableBindingCookie *bool `json:"enableBindingCookie,omitempty" tf:"enable_binding_cookie,omitempty"`

	// Option to add the `HttpOnly` cookie flag to access tokens.
	// +kubebuilder:validation:Optional
	HTTPOnlyCookieAttribute *bool `json:"httpOnlyCookieAttribute,omitempty" tf:"http_only_cookie_attribute,omitempty"`

	// Image URL for the logo shown in the app launcher dashboard.
	// +kubebuilder:validation:Optional
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`

	// Friendly name of the Access Application.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// SaaS configuration for the Access Application.
	// +kubebuilder:validation:Optional
	SaasApp []SaasAppParameters `json:"saasApp,omitempty" tf:"saas_app,omitempty"`

	// Defines the same-site cookie setting for access tokens. Available values: `none`, `lax`, `strict`.
	// +kubebuilder:validation:Optional
	SameSiteCookieAttribute *string `json:"sameSiteCookieAttribute,omitempty" tf:"same_site_cookie_attribute,omitempty"`

	// List of domains that access will secure. Only present for self_hosted, vnc, and ssh applications. Always includes the value set as `domain`.
	// +kubebuilder:validation:Optional
	SelfHostedDomains []*string `json:"selfHostedDomains,omitempty" tf:"self_hosted_domains,omitempty"`

	// Option to return a 401 status code in service authentication rules on failed requests. Defaults to `false`.
	// +kubebuilder:validation:Optional
	ServiceAuth401Redirect *bool `json:"serviceAuth401Redirect,omitempty" tf:"service_auth_401_redirect,omitempty"`

	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`. Defaults to `24h`.
	// +kubebuilder:validation:Optional
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// Option to skip the authorization interstitial when using the CLI. Defaults to `false`.
	// +kubebuilder:validation:Optional
	SkipInterstitial *bool `json:"skipInterstitial,omitempty" tf:"skip_interstitial,omitempty"`

	// The itags associated with the application.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The application type. Available values: `app_launcher`, `bookmark`, `biso`, `dash_sso`, `saas`, `self_hosted`, `ssh`, `vnc`, `warp`. Defaults to `self_hosted`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The zone identifier to target for the resource. Conflicts with `account_id`.
	// +crossplane:generate:reference:type=github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type CorsHeadersInitParameters struct {

	// Value to determine whether all HTTP headers are exposed.
	AllowAllHeaders *bool `json:"allowAllHeaders,omitempty" tf:"allow_all_headers,omitempty"`

	// Value to determine whether all methods are exposed.
	AllowAllMethods *bool `json:"allowAllMethods,omitempty" tf:"allow_all_methods,omitempty"`

	// Value to determine whether all origins are permitted to make CORS requests.
	AllowAllOrigins *bool `json:"allowAllOrigins,omitempty" tf:"allow_all_origins,omitempty"`

	// Value to determine if credentials (cookies, authorization headers, or TLS client certificates) are included with requests.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// List of HTTP headers to expose via CORS.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// List of methods to expose via CORS.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// List of origins permitted to make CORS requests.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The maximum time a preflight request will be cached.
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type CorsHeadersObservation struct {

	// Value to determine whether all HTTP headers are exposed.
	AllowAllHeaders *bool `json:"allowAllHeaders,omitempty" tf:"allow_all_headers,omitempty"`

	// Value to determine whether all methods are exposed.
	AllowAllMethods *bool `json:"allowAllMethods,omitempty" tf:"allow_all_methods,omitempty"`

	// Value to determine whether all origins are permitted to make CORS requests.
	AllowAllOrigins *bool `json:"allowAllOrigins,omitempty" tf:"allow_all_origins,omitempty"`

	// Value to determine if credentials (cookies, authorization headers, or TLS client certificates) are included with requests.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// List of HTTP headers to expose via CORS.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// List of methods to expose via CORS.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// List of origins permitted to make CORS requests.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The maximum time a preflight request will be cached.
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type CorsHeadersParameters struct {

	// Value to determine whether all HTTP headers are exposed.
	// +kubebuilder:validation:Optional
	AllowAllHeaders *bool `json:"allowAllHeaders,omitempty" tf:"allow_all_headers,omitempty"`

	// Value to determine whether all methods are exposed.
	// +kubebuilder:validation:Optional
	AllowAllMethods *bool `json:"allowAllMethods,omitempty" tf:"allow_all_methods,omitempty"`

	// Value to determine whether all origins are permitted to make CORS requests.
	// +kubebuilder:validation:Optional
	AllowAllOrigins *bool `json:"allowAllOrigins,omitempty" tf:"allow_all_origins,omitempty"`

	// Value to determine if credentials (cookies, authorization headers, or TLS client certificates) are included with requests.
	// +kubebuilder:validation:Optional
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// List of HTTP headers to expose via CORS.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// List of methods to expose via CORS.
	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// List of origins permitted to make CORS requests.
	// +kubebuilder:validation:Optional
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The maximum time a preflight request will be cached.
	// +kubebuilder:validation:Optional
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type CustomAttributeInitParameters struct {

	// A friendly name for the attribute as provided to the SaaS app.
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// The name of the attribute as provided to the SaaS app.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A globally unique name for an identity or service provider.
	NameFormat *string `json:"nameFormat,omitempty" tf:"name_format,omitempty"`

	// True if the attribute must be always present.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type CustomAttributeObservation struct {

	// A friendly name for the attribute as provided to the SaaS app.
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// The name of the attribute as provided to the SaaS app.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A globally unique name for an identity or service provider.
	NameFormat *string `json:"nameFormat,omitempty" tf:"name_format,omitempty"`

	// True if the attribute must be always present.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`
}

type CustomAttributeParameters struct {

	// A friendly name for the attribute as provided to the SaaS app.
	// +kubebuilder:validation:Optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// The name of the attribute as provided to the SaaS app.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A globally unique name for an identity or service provider.
	// +kubebuilder:validation:Optional
	NameFormat *string `json:"nameFormat,omitempty" tf:"name_format,omitempty"`

	// True if the attribute must be always present.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type SaasAppInitParameters struct {

	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	ConsumerServiceURL *string `json:"consumerServiceUrl,omitempty" tf:"consumer_service_url,omitempty"`

	// Custom attribute mapped from IDPs.
	CustomAttribute []CustomAttributeInitParameters `json:"customAttribute,omitempty" tf:"custom_attribute,omitempty"`

	// The format of the name identifier sent to the SaaS application. Defaults to `email`.
	NameIDFormat *string `json:"nameIdFormat,omitempty" tf:"name_id_format,omitempty"`

	// A globally unique name for an identity or service provider.
	SpEntityID *string `json:"spEntityId,omitempty" tf:"sp_entity_id,omitempty"`
}

type SaasAppObservation struct {

	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	ConsumerServiceURL *string `json:"consumerServiceUrl,omitempty" tf:"consumer_service_url,omitempty"`

	// Custom attribute mapped from IDPs.
	CustomAttribute []CustomAttributeObservation `json:"customAttribute,omitempty" tf:"custom_attribute,omitempty"`

	// The unique identifier for the SaaS application.
	IdpEntityID *string `json:"idpEntityId,omitempty" tf:"idp_entity_id,omitempty"`

	// The format of the name identifier sent to the SaaS application. Defaults to `email`.
	NameIDFormat *string `json:"nameIdFormat,omitempty" tf:"name_id_format,omitempty"`

	// The public certificate that will be used to verify identities.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// A globally unique name for an identity or service provider.
	SpEntityID *string `json:"spEntityId,omitempty" tf:"sp_entity_id,omitempty"`

	// The endpoint where the SaaS application will send login requests.
	SsoEndpoint *string `json:"ssoEndpoint,omitempty" tf:"sso_endpoint,omitempty"`
}

type SaasAppParameters struct {

	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	// +kubebuilder:validation:Optional
	ConsumerServiceURL *string `json:"consumerServiceUrl,omitempty" tf:"consumer_service_url,omitempty"`

	// Custom attribute mapped from IDPs.
	// +kubebuilder:validation:Optional
	CustomAttribute []CustomAttributeParameters `json:"customAttribute,omitempty" tf:"custom_attribute,omitempty"`

	// The format of the name identifier sent to the SaaS application. Defaults to `email`.
	// +kubebuilder:validation:Optional
	NameIDFormat *string `json:"nameIdFormat,omitempty" tf:"name_id_format,omitempty"`

	// A globally unique name for an identity or service provider.
	// +kubebuilder:validation:Optional
	SpEntityID *string `json:"spEntityId,omitempty" tf:"sp_entity_id,omitempty"`
}

type SourceInitParameters struct {

	// The name of the attribute as provided by the IDP.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SourceObservation struct {

	// The name of the attribute as provided by the IDP.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SourceParameters struct {

	// The name of the attribute as provided by the IDP.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
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
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
