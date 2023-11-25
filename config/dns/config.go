package dns

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for dns group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_record", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"data",
				"value",
			},
		}
	})
	p.AddResourceConfigurator("cloudflare_regional_hostname", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "RegionalHostname"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})

}
