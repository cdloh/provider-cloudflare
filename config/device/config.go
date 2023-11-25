package device

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for customhostname group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_device_dex_test", func(r *config.Resource) {
		r.Kind = "DexText"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_managed_networks", func(r *config.Resource) {
		r.Kind = "ManagedNetwork"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_policy_certificates", func(r *config.Resource) {
		r.Kind = "PolicyCertificate"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_posture_integration", func(r *config.Resource) {
		r.Kind = "PostureIntegration"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_posture_rule", func(r *config.Resource) {
		r.Kind = "PostureRule"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_device_settings_policy", func(r *config.Resource) {
		r.Kind = "SettingsPolicy"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
