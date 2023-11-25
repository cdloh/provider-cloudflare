package notification

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for notification group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_notification_policy", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["webhooks_integration.id"] = config.Reference{
			Type: "PolicyWebhooks",
		}
	})
	p.AddResourceConfigurator("cloudflare_notification_policy_webhooks", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
}
