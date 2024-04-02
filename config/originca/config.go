package originca

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for originca group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_origin_ca_certificate", func(r *config.Resource) {
		r.ShortGroup = "OriginCA"
		r.Kind = "Certificate"
	})
}
