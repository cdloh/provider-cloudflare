/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1"
	v1alpha11 "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Address.
func (mg *Address) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.AccountIDRef,
		Selector:     mg.Spec.InitProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AccountID")
	}
	mg.Spec.InitProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CatchAll.
func (mg *CatchAll) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Settings.
func (mg *Settings) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}
