package authenticatedoriginpulls

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for authenticatedorigins group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_authenticated_origin_pulls", func(r *config.Resource) {
		r.ShortGroup = "AuthenticatedOriginPulls"
		r.Kind = "AuthenticatedOriginsPulls"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.References["authenticated_origin_pulls_certificate"] = config.Reference{
			Type: "Certificate",
		}
	})
	p.AddResourceConfigurator("cloudflare_authenticated_origin_pulls_certificate", func(r *config.Resource) {
		r.ShortGroup = "AuthenticatedOriginPulls"
		r.Kind = "Certificate"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
