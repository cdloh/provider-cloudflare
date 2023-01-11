/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/cdloh/provider-cloudflare/config/access"
	"github.com/cdloh/provider-cloudflare/config/account"
	"github.com/cdloh/provider-cloudflare/config/apishield"
	"github.com/cdloh/provider-cloudflare/config/argo"
	"github.com/cdloh/provider-cloudflare/config/authenticatedoriginpulls"
	"github.com/cdloh/provider-cloudflare/config/byoip"
	"github.com/cdloh/provider-cloudflare/config/certificate"
	"github.com/cdloh/provider-cloudflare/config/custom"
	"github.com/cdloh/provider-cloudflare/config/customhostname"
	"github.com/cdloh/provider-cloudflare/config/dlp"
	"github.com/cdloh/provider-cloudflare/config/dns"
	"github.com/cdloh/provider-cloudflare/config/emailrouting"
	"github.com/cdloh/provider-cloudflare/config/filters"
	"github.com/cdloh/provider-cloudflare/config/firewall"
	"github.com/cdloh/provider-cloudflare/config/lists"
	"github.com/cdloh/provider-cloudflare/config/loadbalancer"
	"github.com/cdloh/provider-cloudflare/config/logpush"
	"github.com/cdloh/provider-cloudflare/config/magic"
	"github.com/cdloh/provider-cloudflare/config/notification"
	"github.com/cdloh/provider-cloudflare/config/originca"
	"github.com/cdloh/provider-cloudflare/config/page"
	"github.com/cdloh/provider-cloudflare/config/pages"
	"github.com/cdloh/provider-cloudflare/config/ruleset"
	"github.com/cdloh/provider-cloudflare/config/spectrum"
	"github.com/cdloh/provider-cloudflare/config/teams"
	"github.com/cdloh/provider-cloudflare/config/urlnormalization"
	"github.com/cdloh/provider-cloudflare/config/useragent"
	"github.com/cdloh/provider-cloudflare/config/waf"
	"github.com/cdloh/provider-cloudflare/config/waitingroom"
	"github.com/cdloh/provider-cloudflare/config/warp"
	"github.com/cdloh/provider-cloudflare/config/web3"
	"github.com/cdloh/provider-cloudflare/config/worker"
	"github.com/cdloh/provider-cloudflare/config/zone"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/cdloh/provider-cloudflare"
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
		apishield.Configure,
		argo.Configure,
		authenticatedoriginpulls.Configure,
		byoip.Configure,
		certificate.Configure,
		custom.Configure,
		customhostname.Configure,
		dlp.Configure,
		dns.Configure,
		emailrouting.Configure,
		filters.Configure,
		firewall.Configure,
		lists.Configure,
		loadbalancer.Configure,
		logpush.Configure,
		magic.Configure,
		notification.Configure,
		originca.Configure,
		page.Configure,
		pages.Configure,
		ruleset.Configure,
		spectrum.Configure,
		teams.Configure,
		urlnormalization.Configure,
		useragent.Configure,
		waf.Configure,
		waitingroom.Configure,
		warp.Configure,
		web3.Configure,
		worker.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
