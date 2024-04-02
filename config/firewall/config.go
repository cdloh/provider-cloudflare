package firewall

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for firewall group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = "Firewall"
		r.Kind = "Rule"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.References["filter_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/filters/v1alpha1.Filter",
		}
	})
}
