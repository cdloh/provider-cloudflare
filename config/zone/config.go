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
			Type: "Zone",
		}

	})
	p.AddResourceConfigurator("cloudflare_zone_dnssec", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
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
	p.AddResourceConfigurator("cloudflare_logpull_retention", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "LogpullRetention"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_rate_limit", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "RateLimit"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_managed_headers", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "ManagedHeaders"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_healthcheck", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Healthcheck"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
}
