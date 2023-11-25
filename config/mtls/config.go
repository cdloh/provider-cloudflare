package mtls

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for magic group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_mtls_certificate", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
