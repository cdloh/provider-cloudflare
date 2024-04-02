package filters

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for filters group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_filter", func(r *config.Resource) {
		r.ShortGroup = "Filters"
		r.Kind = "Filter"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
