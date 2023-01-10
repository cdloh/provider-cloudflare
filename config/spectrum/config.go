package spectrum

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for spectrum group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_spectrum_application", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
