package lists

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for lists group.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("cloudflare_list", func(r *config.Resource) {
		r.ShortGroup = "lists"
		r.Kind = "List"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_list_item", func(r *config.Resource) {
		r.ShortGroup = "lists"
		r.Kind = "Item"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["list_id"] = config.Reference{
			Type: "List",
		}
	})
}
