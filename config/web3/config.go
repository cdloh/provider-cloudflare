package web3

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for web3 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_web3_hostname", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
