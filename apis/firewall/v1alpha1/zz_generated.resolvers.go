/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/cdloh/provider-cloudflare/apis/filters/v1alpha1"
	v1alpha11 "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FilterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FilterIDRef,
		Selector:     mg.Spec.ForProvider.FilterIDSelector,
		To: reference.To{
			List:    &v1alpha1.FilterList{},
			Managed: &v1alpha1.Filter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FilterID")
	}
	mg.Spec.ForProvider.FilterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FilterIDRef = rsp.ResolvedReference

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
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FilterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FilterIDRef,
		Selector:     mg.Spec.InitProvider.FilterIDSelector,
		To: reference.To{
			List:    &v1alpha1.FilterList{},
			Managed: &v1alpha1.Filter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FilterID")
	}
	mg.Spec.InitProvider.FilterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FilterIDRef = rsp.ResolvedReference

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
