package authenticatedoriginpulls

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for authenticatedorigins group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_authenticated_origin_pulls", func(r *config.Resource) {
		r.ShortGroup = "AuthenticatedOriginPulls"
		r.Kind = "AuthenticatedOriginsPulls"
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
		r.References["authenticated_origin_pulls_certificate"] = config.Reference{
			Type: "Certificate",
		}
	})
	p.AddResourceConfigurator("cloudflare_authenticated_origin_pulls_certificate", func(r *config.Resource) {
		r.ShortGroup = "AuthenticatedOriginPulls"
		r.Kind = "Certificate"
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
	})
}
