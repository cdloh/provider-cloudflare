package zone

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "Zone"
)

// Configure adds configurations for zone group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *config.Resource) {
		r.ShortGroup = "zone"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
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
	p.AddResourceConfigurator("cloudflare_hostname_tls_setting", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TLSSetting"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_hostname_tls_setting_ciphers", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TLSSettingCiphers"
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
	p.AddResourceConfigurator("cloudflare_url_normalization_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "URLNormalizationSettings"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_user_agent_blocking_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "UserAgentBlockingRule"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_user_agent_blocking_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "UserAgentBlockingRule"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_keyless_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "KeylessCertificate"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_observatory_scheduled_test", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "ObservatoryScheduledTest"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_regional_tiered_cache", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "RegionalTieredCache"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_turnstile_widget", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TurnstileWidget"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_zone_cache_reserve", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "CacheServer"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_zone_cache_variants", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "ZoneCacheVariants"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_zone_hold", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Hold"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_zone_lockdown", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Lockdown"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})

}
