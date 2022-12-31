/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/cdloh/provider-cloudflare/internal/controller/providerconfig"
	group "github.com/cdloh/provider-cloudflare/internal/controller/waf/group"
	override "github.com/cdloh/provider-cloudflare/internal/controller/waf/override"
	rule "github.com/cdloh/provider-cloudflare/internal/controller/waf/rule"
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
		providerconfig.Setup,
		group.Setup,
		override.Setup,
		rule.Setup,
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
