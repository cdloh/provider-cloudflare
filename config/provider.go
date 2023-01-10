/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/cdloh/provider-cloudflare/config/account"
	"github.com/cdloh/provider-cloudflare/config/apishield"
	"github.com/cdloh/provider-cloudflare/config/argo"
	"github.com/cdloh/provider-cloudflare/config/authenticatedoriginpulls"
	"github.com/cdloh/provider-cloudflare/config/byoip"
	"github.com/cdloh/provider-cloudflare/config/certificate"
	"github.com/cdloh/provider-cloudflare/config/custom"
	"github.com/cdloh/provider-cloudflare/config/customhostname"
	"github.com/cdloh/provider-cloudflare/config/dns"
	"github.com/cdloh/provider-cloudflare/config/emailrouting"
	"github.com/cdloh/provider-cloudflare/config/loadbalancer"
	"github.com/cdloh/provider-cloudflare/config/logpush"
	"github.com/cdloh/provider-cloudflare/config/notification"
	"github.com/cdloh/provider-cloudflare/config/originca"
	"github.com/cdloh/provider-cloudflare/config/page"
	"github.com/cdloh/provider-cloudflare/config/pages"
	"github.com/cdloh/provider-cloudflare/config/spectrum"
	"github.com/cdloh/provider-cloudflare/config/teams"
	"github.com/cdloh/provider-cloudflare/config/urlnormalization"
	"github.com/cdloh/provider-cloudflare/config/useragent"
	"github.com/cdloh/provider-cloudflare/config/waf"
	"github.com/cdloh/provider-cloudflare/config/waitingroom"
	"github.com/cdloh/provider-cloudflare/config/warp"
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
		account.Configure,
		apishield.Configure,
		argo.Configure,
		authenticatedoriginpulls.Configure,
		byoip.Configure,
		certificate.Configure,
		custom.Configure,
		customhostname.Configure,
		dns.Configure,
		emailrouting.Configure,
		loadbalancer.Configure,
		logpush.Configure,
		notification.Configure,
		originca.Configure,
		page.Configure,
		pages.Configure,
		spectrum.Configure,
		teams.Configure,
		urlnormalization.Configure,
		useragent.Configure,
		waf.Configure,
		waitingroom.Configure,
		warp.Configure,
		worker.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
