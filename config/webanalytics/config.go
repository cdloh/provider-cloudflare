package webanalytics

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "webanalytics"
)

// Configure adds configurations for waitingroom group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_web_analytics_site", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Site"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_web_analytics_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Rule"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		// TODO: Add ruleset_id ref
	})
}
