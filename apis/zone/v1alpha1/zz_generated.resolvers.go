/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DNSSEC.
func (mg *DNSSEC) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Healthcheck.
func (mg *Healthcheck) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LogpullRetention.
func (mg *LogpullRetention) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ManagedHeaders.
func (mg *ManagedHeaders) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RateLimit.
func (mg *RateLimit) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SettingsOverride.
func (mg *SettingsOverride) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TieredCache.
func (mg *TieredCache) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TotalTLS.
func (mg *TotalTLS) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this URLNormalizationSettings.
func (mg *URLNormalizationSettings) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserAgentBlockingRule.
func (mg *UserAgentBlockingRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Zone.
func (mg *Zone) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountIDRef,
		Selector:     mg.Spec.ForProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	return nil
}
