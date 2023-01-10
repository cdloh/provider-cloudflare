package lists

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for lists group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_ip_list", func(r *config.Resource) {
		r.ShortGroup = "lists"
		r.Kind = "IPList"
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_list", func(r *config.Resource) {
		r.ShortGroup = "lists"
		r.Kind = "List"
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
