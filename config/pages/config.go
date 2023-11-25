package pages

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for pages group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_pages_domain", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["project_name"] = config.Reference{
			Type: "Project",
		}
	})
	p.AddResourceConfigurator("cloudflare_pages_project", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
