/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/clementblaise/provider-cloudflare/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal cloudflare credentials as JSON"

	// Configuration options
	accountID         = "account_id"
	apiBasePath       = "api_base_path"
	apiClientLogging  = "api_client_logging"
	apiHostname       = "api_hostname"
	apiKey            = "api_key"
	apiToken          = "api_token"
	apiUserServiceKey = "api_user_service_key"
	email             = "email"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if v, ok := creds[accountID]; ok {
			ps.Configuration[accountID] = v
		}
		if v, ok := creds[apiBasePath]; ok {
			ps.Configuration[apiBasePath] = v
		}
		if v, ok := creds[apiClientLogging]; ok {
			ps.Configuration[apiClientLogging] = v
		}
		if v, ok := creds[apiHostname]; ok {
			ps.Configuration[apiHostname] = v
		}
		if v, ok := creds[apiKey]; ok {
			ps.Configuration[apiKey] = v
		}
		if v, ok := creds[apiToken]; ok {
			ps.Configuration[apiToken] = v
		}
		if v, ok := creds[apiUserServiceKey]; ok {
			ps.Configuration[apiUserServiceKey] = v
		}
		if v, ok := creds[email]; ok {
			ps.Configuration[email] = v
		}
		return ps, nil
	}
}
