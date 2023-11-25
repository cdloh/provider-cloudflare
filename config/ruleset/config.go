package ruleset

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for ruleset group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_ruleset", func(r *config.Resource) {
		r.ShortGroup = "Ruleset"
		r.Kind = "Ruleset"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		config.MoveToStatus(r.TerraformResource,
			"rules.action_parameters.overrides.enabled",
			"rules.action_parameters.overrides.categories.enabled",
			"rules.action_parameters.overrides.rules.enabled",
			"rules.logging.enabled",
		)
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"account_id",
				"zone_id",
				"rules.action_parameters.cache_key.cache_by_device_type",
				"rules.action_parameters.cache_key.custom_key.user.device_type",
				"rules.action_parameters.cache_key.custom_key.query_string.exclude",
				"rules.action_parameters.cache_key.custom_key.query_string.include",
				"rules.action_parameters.edge_ttl.status_code_ttl.status_code",
				"rules.action_parameters.edge_ttl.status_code_ttl.status_code_range",
				"rules.action_parameters.from_value.target_url.expression",
				"rules.action_parameters.from_value.target_url.value",
				"rules.action_parameters.headers.expression",
				"rules.action_parameters.headers.value",
			},
		}
	})
}
