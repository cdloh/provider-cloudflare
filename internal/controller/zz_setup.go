/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/clementblaise/provider-cloudflare/internal/controller/access/application"
	cacertificate "github.com/clementblaise/provider-cloudflare/internal/controller/access/cacertificate"
	group "github.com/clementblaise/provider-cloudflare/internal/controller/access/group"
	identityprovider "github.com/clementblaise/provider-cloudflare/internal/controller/access/identityprovider"
	keysconfiguration "github.com/clementblaise/provider-cloudflare/internal/controller/access/keysconfiguration"
	mutualtlscertificate "github.com/clementblaise/provider-cloudflare/internal/controller/access/mutualtlscertificate"
	organization "github.com/clementblaise/provider-cloudflare/internal/controller/access/organization"
	policy "github.com/clementblaise/provider-cloudflare/internal/controller/access/policy"
	rule "github.com/clementblaise/provider-cloudflare/internal/controller/access/rule"
	servicetoken "github.com/clementblaise/provider-cloudflare/internal/controller/access/servicetoken"
	tag "github.com/clementblaise/provider-cloudflare/internal/controller/access/tag"
	account "github.com/clementblaise/provider-cloudflare/internal/controller/account/account"
	apitoken "github.com/clementblaise/provider-cloudflare/internal/controller/account/apitoken"
	member "github.com/clementblaise/provider-cloudflare/internal/controller/account/member"
	addressmap "github.com/clementblaise/provider-cloudflare/internal/controller/addressmap/addressmap"
	apishield "github.com/clementblaise/provider-cloudflare/internal/controller/apishield/apishield"
	authenticatedoriginspulls "github.com/clementblaise/provider-cloudflare/internal/controller/authenticatedoriginpulls/authenticatedoriginspulls"
	certificate "github.com/clementblaise/provider-cloudflare/internal/controller/authenticatedoriginpulls/certificate"
	management "github.com/clementblaise/provider-cloudflare/internal/controller/bot/management"
	ipprefix "github.com/clementblaise/provider-cloudflare/internal/controller/byoip/ipprefix"
	pack "github.com/clementblaise/provider-cloudflare/internal/controller/certificate/pack"
	queue "github.com/clementblaise/provider-cloudflare/internal/controller/cloudflare/queue"
	pages "github.com/clementblaise/provider-cloudflare/internal/controller/custom/pages"
	ssl "github.com/clementblaise/provider-cloudflare/internal/controller/custom/ssl"
	fallbackorigin "github.com/clementblaise/provider-cloudflare/internal/controller/customhostname/fallbackorigin"
	hostname "github.com/clementblaise/provider-cloudflare/internal/controller/customhostname/hostname"
	database "github.com/clementblaise/provider-cloudflare/internal/controller/d1/database"
	dextext "github.com/clementblaise/provider-cloudflare/internal/controller/device/dextext"
	managednetwork "github.com/clementblaise/provider-cloudflare/internal/controller/device/managednetwork"
	profile "github.com/clementblaise/provider-cloudflare/internal/controller/dlp/profile"
	record "github.com/clementblaise/provider-cloudflare/internal/controller/dns/record"
	regionalhostname "github.com/clementblaise/provider-cloudflare/internal/controller/dns/regionalhostname"
	address "github.com/clementblaise/provider-cloudflare/internal/controller/emailrouting/address"
	catchall "github.com/clementblaise/provider-cloudflare/internal/controller/emailrouting/catchall"
	ruleemailrouting "github.com/clementblaise/provider-cloudflare/internal/controller/emailrouting/rule"
	settings "github.com/clementblaise/provider-cloudflare/internal/controller/emailrouting/settings"
	filter "github.com/clementblaise/provider-cloudflare/internal/controller/filters/filter"
	rulefirewall "github.com/clementblaise/provider-cloudflare/internal/controller/firewall/rule"
	item "github.com/clementblaise/provider-cloudflare/internal/controller/lists/item"
	list "github.com/clementblaise/provider-cloudflare/internal/controller/lists/list"
	loadbalancer "github.com/clementblaise/provider-cloudflare/internal/controller/loadbalancer/loadbalancer"
	monitor "github.com/clementblaise/provider-cloudflare/internal/controller/loadbalancer/monitor"
	pool "github.com/clementblaise/provider-cloudflare/internal/controller/loadbalancer/pool"
	job "github.com/clementblaise/provider-cloudflare/internal/controller/logpush/job"
	ownershipchallenge "github.com/clementblaise/provider-cloudflare/internal/controller/logpush/ownershipchallenge"
	firewallruleset "github.com/clementblaise/provider-cloudflare/internal/controller/magic/firewallruleset"
	gretunnel "github.com/clementblaise/provider-cloudflare/internal/controller/magic/gretunnel"
	ipsectunnel "github.com/clementblaise/provider-cloudflare/internal/controller/magic/ipsectunnel"
	staticroute "github.com/clementblaise/provider-cloudflare/internal/controller/magic/staticroute"
	certificatemtls "github.com/clementblaise/provider-cloudflare/internal/controller/mtls/certificate"
	policynotification "github.com/clementblaise/provider-cloudflare/internal/controller/notification/policy"
	policywebhooks "github.com/clementblaise/provider-cloudflare/internal/controller/notification/policywebhooks"
	certificateoriginca "github.com/clementblaise/provider-cloudflare/internal/controller/originca/certificate"
	rulepage "github.com/clementblaise/provider-cloudflare/internal/controller/page/rule"
	domain "github.com/clementblaise/provider-cloudflare/internal/controller/pages/domain"
	project "github.com/clementblaise/provider-cloudflare/internal/controller/pages/project"
	providerconfig "github.com/clementblaise/provider-cloudflare/internal/controller/providerconfig"
	bucket "github.com/clementblaise/provider-cloudflare/internal/controller/r2/bucket"
	ruleset "github.com/clementblaise/provider-cloudflare/internal/controller/ruleset/ruleset"
	applicationspectrum "github.com/clementblaise/provider-cloudflare/internal/controller/spectrum/application"
	accountteams "github.com/clementblaise/provider-cloudflare/internal/controller/teams/account"
	listteams "github.com/clementblaise/provider-cloudflare/internal/controller/teams/list"
	location "github.com/clementblaise/provider-cloudflare/internal/controller/teams/location"
	proxyendpoint "github.com/clementblaise/provider-cloudflare/internal/controller/teams/proxyendpoint"
	ruleteams "github.com/clementblaise/provider-cloudflare/internal/controller/teams/rule"
	argo "github.com/clementblaise/provider-cloudflare/internal/controller/tunnel/argo"
	config "github.com/clementblaise/provider-cloudflare/internal/controller/tunnel/config"
	route "github.com/clementblaise/provider-cloudflare/internal/controller/tunnel/route"
	tunnel "github.com/clementblaise/provider-cloudflare/internal/controller/tunnel/tunnel"
	virtualnetwork "github.com/clementblaise/provider-cloudflare/internal/controller/tunnel/virtualnetwork"
	event "github.com/clementblaise/provider-cloudflare/internal/controller/waitingroom/event"
	room "github.com/clementblaise/provider-cloudflare/internal/controller/waitingroom/room"
	rules "github.com/clementblaise/provider-cloudflare/internal/controller/waitingroom/rules"
	settingswaitingroom "github.com/clementblaise/provider-cloudflare/internal/controller/waitingroom/settings"
	devicepolicycertificates "github.com/clementblaise/provider-cloudflare/internal/controller/warp/devicepolicycertificates"
	devicepostureintegration "github.com/clementblaise/provider-cloudflare/internal/controller/warp/devicepostureintegration"
	deviceposturerule "github.com/clementblaise/provider-cloudflare/internal/controller/warp/deviceposturerule"
	devicesettingspolicy "github.com/clementblaise/provider-cloudflare/internal/controller/warp/devicesettingspolicy"
	fallbackdomain "github.com/clementblaise/provider-cloudflare/internal/controller/warp/fallbackdomain"
	splittunnel "github.com/clementblaise/provider-cloudflare/internal/controller/warp/splittunnel"
	hostnameweb3 "github.com/clementblaise/provider-cloudflare/internal/controller/web3/hostname"
	rulewebanalytics "github.com/clementblaise/provider-cloudflare/internal/controller/webanalytics/rule"
	site "github.com/clementblaise/provider-cloudflare/internal/controller/webanalytics/site"
	crontrigger "github.com/clementblaise/provider-cloudflare/internal/controller/worker/crontrigger"
	domainworker "github.com/clementblaise/provider-cloudflare/internal/controller/worker/domain"
	kv "github.com/clementblaise/provider-cloudflare/internal/controller/worker/kv"
	kvnamespace "github.com/clementblaise/provider-cloudflare/internal/controller/worker/kvnamespace"
	routeworker "github.com/clementblaise/provider-cloudflare/internal/controller/worker/route"
	script "github.com/clementblaise/provider-cloudflare/internal/controller/worker/script"
	cacheserver "github.com/clementblaise/provider-cloudflare/internal/controller/zone/cacheserver"
	dnssec "github.com/clementblaise/provider-cloudflare/internal/controller/zone/dnssec"
	healthcheck "github.com/clementblaise/provider-cloudflare/internal/controller/zone/healthcheck"
	hold "github.com/clementblaise/provider-cloudflare/internal/controller/zone/hold"
	keylesscertificate "github.com/clementblaise/provider-cloudflare/internal/controller/zone/keylesscertificate"
	lockdown "github.com/clementblaise/provider-cloudflare/internal/controller/zone/lockdown"
	logpullretention "github.com/clementblaise/provider-cloudflare/internal/controller/zone/logpullretention"
	managedheaders "github.com/clementblaise/provider-cloudflare/internal/controller/zone/managedheaders"
	observatoryscheduledtest "github.com/clementblaise/provider-cloudflare/internal/controller/zone/observatoryscheduledtest"
	ratelimit "github.com/clementblaise/provider-cloudflare/internal/controller/zone/ratelimit"
	regionaltieredcache "github.com/clementblaise/provider-cloudflare/internal/controller/zone/regionaltieredcache"
	settingsoverride "github.com/clementblaise/provider-cloudflare/internal/controller/zone/settingsoverride"
	tieredcache "github.com/clementblaise/provider-cloudflare/internal/controller/zone/tieredcache"
	tlssetting "github.com/clementblaise/provider-cloudflare/internal/controller/zone/tlssetting"
	tlssettingciphers "github.com/clementblaise/provider-cloudflare/internal/controller/zone/tlssettingciphers"
	totaltls "github.com/clementblaise/provider-cloudflare/internal/controller/zone/totaltls"
	urlnormalizationsettings "github.com/clementblaise/provider-cloudflare/internal/controller/zone/urlnormalizationsettings"
	useragentblockingrule "github.com/clementblaise/provider-cloudflare/internal/controller/zone/useragentblockingrule"
	zone "github.com/clementblaise/provider-cloudflare/internal/controller/zone/zone"
	zonecachevariants "github.com/clementblaise/provider-cloudflare/internal/controller/zone/zonecachevariants"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		cacertificate.Setup,
		group.Setup,
		identityprovider.Setup,
		keysconfiguration.Setup,
		mutualtlscertificate.Setup,
		organization.Setup,
		policy.Setup,
		rule.Setup,
		servicetoken.Setup,
		tag.Setup,
		account.Setup,
		apitoken.Setup,
		member.Setup,
		addressmap.Setup,
		apishield.Setup,
		authenticatedoriginspulls.Setup,
		certificate.Setup,
		management.Setup,
		ipprefix.Setup,
		pack.Setup,
		queue.Setup,
		pages.Setup,
		ssl.Setup,
		fallbackorigin.Setup,
		hostname.Setup,
		database.Setup,
		dextext.Setup,
		managednetwork.Setup,
		profile.Setup,
		record.Setup,
		regionalhostname.Setup,
		address.Setup,
		catchall.Setup,
		ruleemailrouting.Setup,
		settings.Setup,
		filter.Setup,
		rulefirewall.Setup,
		item.Setup,
		list.Setup,
		loadbalancer.Setup,
		monitor.Setup,
		pool.Setup,
		job.Setup,
		ownershipchallenge.Setup,
		firewallruleset.Setup,
		gretunnel.Setup,
		ipsectunnel.Setup,
		staticroute.Setup,
		certificatemtls.Setup,
		policynotification.Setup,
		policywebhooks.Setup,
		certificateoriginca.Setup,
		rulepage.Setup,
		domain.Setup,
		project.Setup,
		providerconfig.Setup,
		bucket.Setup,
		ruleset.Setup,
		applicationspectrum.Setup,
		accountteams.Setup,
		listteams.Setup,
		location.Setup,
		proxyendpoint.Setup,
		ruleteams.Setup,
		argo.Setup,
		config.Setup,
		route.Setup,
		tunnel.Setup,
		virtualnetwork.Setup,
		event.Setup,
		room.Setup,
		rules.Setup,
		settingswaitingroom.Setup,
		devicepolicycertificates.Setup,
		devicepostureintegration.Setup,
		deviceposturerule.Setup,
		devicesettingspolicy.Setup,
		fallbackdomain.Setup,
		splittunnel.Setup,
		hostnameweb3.Setup,
		rulewebanalytics.Setup,
		site.Setup,
		crontrigger.Setup,
		domainworker.Setup,
		kv.Setup,
		kvnamespace.Setup,
		routeworker.Setup,
		script.Setup,
		cacheserver.Setup,
		dnssec.Setup,
		healthcheck.Setup,
		hold.Setup,
		keylesscertificate.Setup,
		lockdown.Setup,
		logpullretention.Setup,
		managedheaders.Setup,
		observatoryscheduledtest.Setup,
		ratelimit.Setup,
		regionaltieredcache.Setup,
		settingsoverride.Setup,
		tieredcache.Setup,
		tlssetting.Setup,
		tlssettingciphers.Setup,
		totaltls.Setup,
		urlnormalizationsettings.Setup,
		useragentblockingrule.Setup,
		zone.Setup,
		zonecachevariants.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
