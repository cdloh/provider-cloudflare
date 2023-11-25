package waitingroom

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroupName = "waitingroom"
)

// Configure adds configurations for waitingroom group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_waiting_room", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Room"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_waiting_room_event", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Event"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.References["waiting_room_id"] = config.Reference{
			Type: "Room",
		}
	})
	p.AddResourceConfigurator("cloudflare_waiting_room_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Rules"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
		r.References["waiting_room_id"] = config.Reference{
			Type: "Room",
		}
	})
	p.AddResourceConfigurator("cloudflare_waiting_room_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroupName
		r.Kind = "Settings"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
