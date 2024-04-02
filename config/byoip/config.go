package byoip

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for byoip group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_byo_ip_prefix", func(r *config.Resource) {
		r.ShortGroup = "BYOIP"
		r.Kind = "IPPrefix"
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
