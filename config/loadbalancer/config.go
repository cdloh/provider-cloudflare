package loadbalancer

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "loadbalancer"
)

// Configure adds configurations for loadbalancer group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_load_balancer", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "LoadBalancer"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.References["default_pool_ids"] = config.Reference{
			Type: "Pool",
		}
		r.References["fallback_pool_id"] = config.Reference{
			Type: "Pool",
		}
		r.References["fallback_pool_id"] = config.Reference{
			Type: "Pool",
		}
		r.References["pop_pools.pool_ids"] = config.Reference{
			Type: "Pool",
		}
		r.References["country_pools.pool_ids"] = config.Reference{
			Type: "Pool",
		}
		r.References["region_pools.pool_ids"] = config.Reference{
			Type: "Pool",
		}
	})
	p.AddResourceConfigurator("cloudflare_load_balancer_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Monitor"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_load_balancer_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Pool"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.References["monitor"] = config.Reference{
			Type: "Monitor",
		}
	})

}
