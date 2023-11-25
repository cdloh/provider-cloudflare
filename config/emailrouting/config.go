package emailrouting

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for emailrouting group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_email_routing_address", func(r *config.Resource) {
		r.ShortGroup = "emailrouting"
		r.Kind = "Address"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_email_routing_catch_all", func(r *config.Resource) {
		r.ShortGroup = "emailrouting"
		r.Kind = "CatchAll"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_email_routing_rule", func(r *config.Resource) {
		r.ShortGroup = "emailrouting"
		r.Kind = "Rule"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_email_routing_settings", func(r *config.Resource) {
		r.ShortGroup = "emailrouting"
		r.Kind = "Settings"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
