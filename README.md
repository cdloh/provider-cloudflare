# Provider Cloudflare

`provider-cloudflare` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Cloudflare API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/cdloh/provider-cloudflare):
```
up ctp provider install cdloh/provider-cloudflare:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudflare
spec:
  package: cdloh/provider-cloudflare:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/clementblaise/provider-cloudflare).

## Resource Refrence and Status

I've tried to keep the resources grouped in a logical fashion and tested as many as I have access to.

Table below shows the Terraform resources and maps it to it's k8s details, along with whether or not testing was completed for the resource.

If you're able to confirm one works please open a PR or issue to update the table.

| TF Resource | API Group | API Version | Kind | Tested |
| ---         | ---       | ---  | ---         | ---    |
| cloudflare_access_application | Application | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_bookmark | Bookmark | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_ca_certificate | CACertificate | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_group | Group | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_identity_provider | IdentityProvider | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_keys_configuration | Configuration | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_mutual_tls_certificate | MutualTLSCertificate | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_organization | Organization | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_policy | Policy | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_rule | Rule | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_access_service_token | ServiceToken | access.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_account  | Account | account.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_account_member | Member | account.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_api_shield | APIShield | apishield.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_api_token | Token | account.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_argo | Argo | argo.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_argo_tunnel | Tunnel | argo.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_authenticated_origin_pulls | AuthenticatedOriginsPulls | authenticatedoriginpulls.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_authenticated_origin_pulls_certificate | Certificate | authenticatedoriginpulls.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_byo_ip_prefix | IPPrefix | IPPrefix | v1alpha1 | ❌ |
| cloudflare_certificate_pack | Pack | certificate.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_custom_hostname | Hostname | customhostname.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_custom_hostname_fallback_origin | FallbackOrigin | customhostname.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_custom_pages | Pages | custom.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_custom_ssl | SSL | custom.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_device_policy_certificates | DevicePolicyCertificates | warp.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_device_posture_integration | DevicePostureIntegration | warp.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_device_posture_rule | DevicePostureRule | warp.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_device_settings_policy | DeviceSettingsPolicy | warp.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_dlp_profile | Profile | dlp.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_email_routing_address | Address | emailrouting.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_email_routing_catch_all | CatchAll | emailrouting.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_email_routing_rule | Rule | emailrouting.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_email_routing_settings | Settings | emailrouting.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_fallback_domain | FallbackDomain | warp.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_filter | Filter | filters.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_firewall_rule | Rule | firewall.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_gre_tunnel | GRETunnel | magic.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_healthcheck | Healthcheck | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_ip_list | IPList | lists.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_ipsec_tunnel | IPsecTunnel | magic.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_list | List | lists.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_load_balancer | LoadBalancer | loadbalancer.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_load_balancer_monitor | Monitor | loadbalancer.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_load_balancer_pool | Pool | loadbalancer.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_logpull_retention | LogpullRetention | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_logpush_job | Job | logpush.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_logpush_ownership_challenge | OwnershipChallenge | logpush.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_magic_firewall_ruleset | FirewallRuleset | magic.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_managed_headers | ManagedHeaders | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_notification_policy | Policy | notification.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_notification_policy_webhooks | PolicyWebhooks | notification.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_origin_ca_certificate | Certificate | originca.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_page_rule | Rule | page.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_rate_limit | RateLimit | zone.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_record | Record | dns.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_ruleset | Ruleset | ruleset.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_spectrum_application | Application | spectrum.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_split_tunnel | SplitTunnel | warp.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_static_route | StaticRoute | magic.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_teams_account | Account | teams.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_teams_list | List | teams.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_teams_location | Location | teams.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_teams_proxy_endpoint | ProxyEndpoint | teams.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_teams_rule | Rule | teams.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_tiered_cache | TieredCache | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_total_tls | TotalTLS | zone.cloudflare.upbound.io  | v1alpha1 | ✅ |
| cloudflare_tunnel_config | TunnelConfig | argo.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_tunnel_route | TunnelRoute | argo.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_tunnel_virtual_network | TunnelVirtualNetwork | argo.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_url_normalization_settings | URLNormalizationSettings | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_user_agent_blocking_rule | UserAgentBlockingRule | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_waf_group | Group | waf.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_waf_override | Override | waf.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_waf_package | Package | waf.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_waf_rule | Rule | waf.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_waiting_room | Room | waitingroom.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_waiting_room_event | Event | waitingroom.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_waiting_room_rules | Rules | waitingroom.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_web3_hostname | Hostname | web3.cloudflare.upbound.io | v1alpha1 | ❌ |
| cloudflare_worker_cron_trigger | CronTrigger | worker.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_worker_route | Route | worker.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_worker_script | Script | worker.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_workers_kv | KV | worker.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_workers_kv_namespace | KVNamespace | worker.cloudflare.upbound.io  | v1alpha1 | ✅ |
| cloudflare_zone | Zone | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_zone_cache_variants | CacheVariants | zone.cloudflare.upbound.io | v1alpha1 | ✅ | SHITE
| cloudflare_zone_dnssec | DNSSEC | zone.cloudflare.upbound.io | v1alpha1 | ✅ |
| cloudflare_zone_lockdown | Lockdown | zone.cloudflare.upbound.io | v1alpha1 | ✅ | SHITE
| cloudflare_zone_settings_override | SettingsOverride | zone.cloudflare.upbound.io | v1alpha1 | ✅ |


## Developing

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/clementblaise/provider-cloudflare/issues).
