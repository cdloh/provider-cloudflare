package dlp

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for dlp group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_dlp_profile", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
