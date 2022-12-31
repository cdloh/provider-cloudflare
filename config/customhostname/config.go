package customhostname

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for customhostnames group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_custom_hostname", func(r *config.Resource) {
		r.ShortGroup = "customhostname"
		r.Kind = "Hostname"
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
	})
	p.AddResourceConfigurator("cloudflare_custom_hostname_fallback_origin", func(r *config.Resource) {
		r.ShortGroup = "customhostname"
		r.Kind = "FallbackOrigin"
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
	})
}