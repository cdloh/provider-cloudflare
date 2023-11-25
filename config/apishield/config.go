package apishield

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for apishield group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_api_shield", func(r *config.Resource) {
		r.ShortGroup = "APIShield"
		r.Kind = "APIShield"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_api_shield_operation", func(r *config.Resource) {
		r.ShortGroup = "APIShield"
		r.Kind = "Operation"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	// TODO: Operation Ids
	p.AddResourceConfigurator("cloudflare_api_shield_operation_schema_validation_settings", func(r *config.Resource) {
		r.ShortGroup = "APIShield"
		r.Kind = "OperationSchemaValidationSettings"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		//r.References["operation_id"] = config.Reference{
		//	Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		//}
	})
	p.AddResourceConfigurator("cloudflare_api_shield_schema", func(r *config.Resource) {
		r.ShortGroup = "APIShield"
		r.Kind = "Schema"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_api_shield_schema_validation_settings", func(r *config.Resource) {
		r.ShortGroup = "APIShield"
		r.Kind = "SchemaValidationSettings"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})

}
