/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CronTrigger.
func (mg *CronTrigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScriptName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScriptNameRefs,
		Selector:     mg.Spec.ForProvider.ScriptNameSelector,
		To: reference.To{
			List:    &ScriptList{},
			Managed: &Script{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScriptName")
	}
	mg.Spec.ForProvider.ScriptName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScriptNameRefs = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Kv.
func (mg *Kv) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &KvNamespaceList{},
			Managed: &KvNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScriptName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScriptNameRefs,
		Selector:     mg.Spec.ForProvider.ScriptNameSelector,
		To: reference.To{
			List:    &ScriptList{},
			Managed: &Script{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScriptName")
	}
	mg.Spec.ForProvider.ScriptName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScriptNameRefs = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRefs,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha1.ZoneList{},
			Managed: &v1alpha1.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRefs = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Script.
func (mg *Script) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.KvNamespaceBinding); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceIDRef,
			Selector:     mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceIDSelector,
			To: reference.To{
				List:    &KvNamespaceList{},
				Managed: &KvNamespace{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceID")
		}
		mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceIDRef = rsp.ResolvedReference

	}

	return nil
}
