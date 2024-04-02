package account

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for account group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_account", func(r *config.Resource) {
		r.ShortGroup = "account"

	})
	p.AddResourceConfigurator("cloudflare_account_member", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.References["account_id"] = config.Reference{
			Type: "Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_api_token", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "APIToken"
	})
}
