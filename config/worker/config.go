package worker

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "worker"
)

// Configure adds configurations for worker group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_worker_cron_trigger", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_script", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["kv_namespace_binding.namespace_id"] = config.Reference{
			Type: "KVNamespace",
		}
	})
	p.AddResourceConfigurator("cloudflare_workers_kv_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "KVNamespace"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
	})
	p.AddResourceConfigurator("cloudflare_workers_kv", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "KV"
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["namespace_id"] = config.Reference{
			Type: "KVNamespace",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["account_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
