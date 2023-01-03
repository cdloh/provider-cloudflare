package byoip

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for byoip group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_byo_ip_prefix", func(r *config.Resource) {
		r.ShortGroup = "BYOIP"
		r.Kind = "IPPrefix"
	})
}
