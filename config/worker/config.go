package worker

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "worker"
)

// Configure adds configurations for zone group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_worker_cron_trigger", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
	})
	p.AddResourceConfigurator("cloudflare_worker_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
		r.References["zone_id"] = config.Reference{
			Type: "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone",
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
	})
	p.AddResourceConfigurator("cloudflare_workers_kv", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.References["namespace_id"] = config.Reference{
			Type: "KVNamespace",
		}
	})
}
