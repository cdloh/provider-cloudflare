package useragent

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for useragent group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_user_agent_blocking_rule", func(r *config.Resource) {
		r.ShortGroup = "UserAgent"
		r.Kind = "BlockingRule"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
