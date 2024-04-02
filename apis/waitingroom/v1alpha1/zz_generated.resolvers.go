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

// ResolveReferences of this Event.
func (mg *Event) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WaitingRoomID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.WaitingRoomIDRef,
		Selector:     mg.Spec.ForProvider.WaitingRoomIDSelector,
		To: reference.To{
			List:    &RoomList{},
			Managed: &Room{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WaitingRoomID")
	}
	mg.Spec.ForProvider.WaitingRoomID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WaitingRoomIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
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
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WaitingRoomID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.WaitingRoomIDRef,
		Selector:     mg.Spec.InitProvider.WaitingRoomIDSelector,
		To: reference.To{
			List:    &RoomList{},
			Managed: &Room{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WaitingRoomID")
	}
	mg.Spec.InitProvider.WaitingRoomID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WaitingRoomIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha1.ZoneList{},
			Managed: &v1alpha1.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Room.
func (mg *Room) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
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
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha1.ZoneList{},
			Managed: &v1alpha1.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rules.
func (mg *Rules) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WaitingRoomID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.WaitingRoomIDRef,
		Selector:     mg.Spec.ForProvider.WaitingRoomIDSelector,
		To: reference.To{
			List:    &RoomList{},
			Managed: &Room{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WaitingRoomID")
	}
	mg.Spec.ForProvider.WaitingRoomID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WaitingRoomIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
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
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WaitingRoomID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.WaitingRoomIDRef,
		Selector:     mg.Spec.InitProvider.WaitingRoomIDSelector,
		To: reference.To{
			List:    &RoomList{},
			Managed: &Room{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WaitingRoomID")
	}
	mg.Spec.InitProvider.WaitingRoomID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WaitingRoomIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha1.ZoneList{},
			Managed: &v1alpha1.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}
