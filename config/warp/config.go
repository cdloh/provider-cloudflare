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
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_posture_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DevicePostureIntegration"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_posture_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DevicePostureRule"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_settings_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "DeviceSettingsPolicy"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_fallback_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "FallbackDomain"
		r.References["policy_id"] = config.Reference{
			Type: "DeviceSettingsPolicy",
		}
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_split_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "SplitTunnel"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["policy_id"] = config.Reference{
			Type: "DeviceSettingsPolicy",
		}
	})
}
