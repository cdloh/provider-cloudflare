/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/clementblaise/provider-cloudflare/config/d1"
	"github.com/clementblaise/provider-cloudflare/config/device"
	"github.com/clementblaise/provider-cloudflare/config/mtls"
	"github.com/clementblaise/provider-cloudflare/config/r2"
	"github.com/clementblaise/provider-cloudflare/config/tunnel"
	"github.com/clementblaise/provider-cloudflare/config/webanalytics"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/clementblaise/provider-cloudflare/config/access"
	"github.com/clementblaise/provider-cloudflare/config/account"
	"github.com/clementblaise/provider-cloudflare/config/addressmap"
	"github.com/clementblaise/provider-cloudflare/config/apishield"
	"github.com/clementblaise/provider-cloudflare/config/authenticatedoriginpulls"
	"github.com/clementblaise/provider-cloudflare/config/bot"
	"github.com/clementblaise/provider-cloudflare/config/byoip"
	"github.com/clementblaise/provider-cloudflare/config/certificate"
	"github.com/clementblaise/provider-cloudflare/config/custom"
	"github.com/clementblaise/provider-cloudflare/config/customhostname"
	"github.com/clementblaise/provider-cloudflare/config/dlp"
	"github.com/clementblaise/provider-cloudflare/config/dns"
	"github.com/clementblaise/provider-cloudflare/config/emailrouting"
	"github.com/clementblaise/provider-cloudflare/config/filters"
	"github.com/clementblaise/provider-cloudflare/config/firewall"
	"github.com/clementblaise/provider-cloudflare/config/lists"
	"github.com/clementblaise/provider-cloudflare/config/loadbalancer"
	"github.com/clementblaise/provider-cloudflare/config/logpush"
	"github.com/clementblaise/provider-cloudflare/config/magic"
	"github.com/clementblaise/provider-cloudflare/config/notification"
	"github.com/clementblaise/provider-cloudflare/config/originca"
	"github.com/clementblaise/provider-cloudflare/config/page"
	"github.com/clementblaise/provider-cloudflare/config/pages"
	"github.com/clementblaise/provider-cloudflare/config/ruleset"
	"github.com/clementblaise/provider-cloudflare/config/spectrum"
	"github.com/clementblaise/provider-cloudflare/config/teams"
	"github.com/clementblaise/provider-cloudflare/config/waitingroom"
	"github.com/clementblaise/provider-cloudflare/config/warp"
	"github.com/clementblaise/provider-cloudflare/config/web3"
	"github.com/clementblaise/provider-cloudflare/config/worker"
	"github.com/clementblaise/provider-cloudflare/config/zone"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/clementblaise/provider-cloudflare"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		access.Configure,
		account.Configure,
		addressmap.Configure,
		apishield.Configure,
		authenticatedoriginpulls.Configure,
		bot.Configure,
		byoip.Configure,
		certificate.Configure,
		custom.Configure,
		customhostname.Configure,
		d1.Configure,
		dlp.Configure,
		device.Configure,
		dns.Configure,
		emailrouting.Configure,
		filters.Configure,
		firewall.Configure,
		lists.Configure,
		loadbalancer.Configure,
		logpush.Configure,
		magic.Configure,
		mtls.Configure,
		notification.Configure,
		originca.Configure,
		page.Configure,
		pages.Configure,
		r2.Configure,
		ruleset.Configure,
		spectrum.Configure,
		teams.Configure,
		tunnel.Configure,
		waitingroom.Configure,
		warp.Configure,
		webanalytics.Configure,
		web3.Configure,
		worker.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
