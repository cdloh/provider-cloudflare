package access

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for access group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_access_application", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_ca_certificate", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["application_id"] = config.Reference{
			Type: "Application",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_group", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_identity_provider", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_keys_configuration", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_access_mutual_tls_certificate", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_organization", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_policy", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["application_id"] = config.Reference{
			Type: "Application",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_rule", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_service_token", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id", "zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_custom_page", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"zone_id"},
		}
	})
	p.AddResourceConfigurator("cloudflare_access_tag", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
