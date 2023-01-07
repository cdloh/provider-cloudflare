package warp

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "WARP"
)

// Configure adds configurations for warp group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_device_policy_certificates", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DevicePolicyCertificates"
		r.References["zone_id"] = config.Reference{
			Type:              "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
			RefFieldName:      "ZoneIDRefs",
			SelectorFieldName: "ZoneIDSelector",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_posture_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DevicePostureIntegration"
	})
	p.AddResourceConfigurator("cloudflare_device_posture_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DevicePostureRule"
	})
	p.AddResourceConfigurator("cloudflare_device_settings_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DeviceSettingsPolicy"
	})
	p.AddResourceConfigurator("cloudflare_split_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "SplitTunnel"
		r.References["policy_id"] = config.Reference{
			Type: "DeviceSettingsPolicy",
		}
	})
}
