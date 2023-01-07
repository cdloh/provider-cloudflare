package custom

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for custom group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_custom_pages", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
	})
	p.AddResourceConfigurator("cloudflare_custom_ssl", func(r *config.Resource) {
		r.Kind = "SSL"
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
	})
}
