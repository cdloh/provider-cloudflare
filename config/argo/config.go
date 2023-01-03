package argo

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for argo group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_argo", func(r *config.Resource) {
		r.ShortGroup = "Argo"
		r.Kind = "Argo"
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
	})
	p.AddResourceConfigurator("cloudflare_argo_tunnel", func(r *config.Resource) {
		r.ShortGroup = "Argo"
		r.Kind = "Tunnel"
	})
}
