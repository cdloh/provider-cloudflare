package argo

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroupName = "Argo"
)

// Configure adds configurations for argo group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_argo", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Argo"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_argo_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Tunnel"
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_tunnel_config", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TunnelConfig"
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["tunnel_id"] = config.Reference{
			Type: "Tunnel",
		}
	})
	p.AddResourceConfigurator("cloudflare_tunnel_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TunnelRoute"
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["tunnel_id"] = config.Reference{
			Type: "Tunnel",
		}
		r.References["virtual_network_id"] = config.Reference{
			Type: "TunnelVirtualNetwork",
		}
	})
	p.AddResourceConfigurator("cloudflare_tunnel_virtual_network", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "TunnelVirtualNetwork"
		r.References["account_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
