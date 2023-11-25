package addressmap

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for address group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_address_map", func(r *config.Resource) {
		r.ShortGroup = "AddressMap"
		r.Kind = "AddressMap"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"account_id"},
		}
		config.MoveToStatus(r.TerraformResource,
			"ips.ip",
			"memberships.identifier",
			"memberships.kind",
		)

	})
}
