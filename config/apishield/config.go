package apishield

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for apishield group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_api_shield", func(r *config.Resource) {
		r.ShortGroup = "APIShield"
		r.Kind = "APIShield"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
