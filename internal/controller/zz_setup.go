/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/cdloh/provider-cloudflare/internal/controller/account/account"
	apitoken "github.com/cdloh/provider-cloudflare/internal/controller/account/apitoken"
	member "github.com/cdloh/provider-cloudflare/internal/controller/account/member"
	apishield "github.com/cdloh/provider-cloudflare/internal/controller/apishield/apishield"
	argo "github.com/cdloh/provider-cloudflare/internal/controller/argo/argo"
	tunnel "github.com/cdloh/provider-cloudflare/internal/controller/argo/tunnel"
	authenticatedoriginspulls "github.com/cdloh/provider-cloudflare/internal/controller/authenticatedoriginpulls/authenticatedoriginspulls"
	certificate "github.com/cdloh/provider-cloudflare/internal/controller/authenticatedoriginpulls/certificate"
	ipprefix "github.com/cdloh/provider-cloudflare/internal/controller/byoip/ipprefix"
	pack "github.com/cdloh/provider-cloudflare/internal/controller/certificate/pack"
	fallbackorigin "github.com/cdloh/provider-cloudflare/internal/controller/customhostname/fallbackorigin"
	hostname "github.com/cdloh/provider-cloudflare/internal/controller/customhostname/hostname"
	record "github.com/cdloh/provider-cloudflare/internal/controller/dns/record"
	rule "github.com/cdloh/provider-cloudflare/internal/controller/page/rule"
	providerconfig "github.com/cdloh/provider-cloudflare/internal/controller/providerconfig"
	group "github.com/cdloh/provider-cloudflare/internal/controller/waf/group"
	override "github.com/cdloh/provider-cloudflare/internal/controller/waf/override"
	rulewaf "github.com/cdloh/provider-cloudflare/internal/controller/waf/rule"
	wafpackage "github.com/cdloh/provider-cloudflare/internal/controller/waf/wafpackage"
	crontrigger "github.com/cdloh/provider-cloudflare/internal/controller/worker/crontrigger"
	kv "github.com/cdloh/provider-cloudflare/internal/controller/worker/kv"
	kvnamespace "github.com/cdloh/provider-cloudflare/internal/controller/worker/kvnamespace"
	route "github.com/cdloh/provider-cloudflare/internal/controller/worker/route"
	script "github.com/cdloh/provider-cloudflare/internal/controller/worker/script"
	dnssec "github.com/cdloh/provider-cloudflare/internal/controller/zone/dnssec"
	settingsoverride "github.com/cdloh/provider-cloudflare/internal/controller/zone/settingsoverride"
	zone "github.com/cdloh/provider-cloudflare/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		apitoken.Setup,
		member.Setup,
		apishield.Setup,
		argo.Setup,
		tunnel.Setup,
		authenticatedoriginspulls.Setup,
		certificate.Setup,
		ipprefix.Setup,
		pack.Setup,
		fallbackorigin.Setup,
		hostname.Setup,
		record.Setup,
		rule.Setup,
		providerconfig.Setup,
		group.Setup,
		override.Setup,
		rulewaf.Setup,
		wafpackage.Setup,
		crontrigger.Setup,
		kv.Setup,
		kvnamespace.Setup,
		route.Setup,
		script.Setup,
		dnssec.Setup,
		settingsoverride.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
