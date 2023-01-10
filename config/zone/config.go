package zone

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "Zone"
)

// Configure adds configurations for zone group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *config.Resource) {
		r.ShortGroup = "zone"
	})
	p.AddResourceConfigurator("cloudflare_zone_settings_override", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["zone_id"] = config.Reference{
			Type:              "Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}

	})
	p.AddResourceConfigurator("cloudflare_zone_dnssec", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["zone_id"] = config.Reference{
			Type:              "Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}

	})
	p.AddResourceConfigurator("cloudflare_total_tls", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TotalTLS"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_tiered_cache", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TieredCache"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
}
