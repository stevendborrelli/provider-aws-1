/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/kinesis/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Destination.
func (mg *Destination) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetArn),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.ForProvider.TargetArnRef,
		Selector:     mg.Spec.ForProvider.TargetArnSelector,
		To: reference.To{
			List:    &v1beta11.StreamList{},
			Managed: &v1beta11.Stream{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetArn")
	}
	mg.Spec.ForProvider.TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TargetArn),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.InitProvider.TargetArnRef,
		Selector:     mg.Spec.InitProvider.TargetArnSelector,
		To: reference.To{
			List:    &v1beta11.StreamList{},
			Managed: &v1beta11.Stream{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TargetArn")
	}
	mg.Spec.InitProvider.TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TargetArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta12.KeyList{},
			Managed: &v1beta12.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KMSKeyIDRef,
		Selector:     mg.Spec.InitProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta12.KeyList{},
			Managed: &v1beta12.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyID")
	}
	mg.Spec.InitProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MetricFilter.
func (mg *MetricFilter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LogGroupNameRef,
		Selector:     mg.Spec.ForProvider.LogGroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogGroupName")
	}
	mg.Spec.ForProvider.LogGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LogGroupNameRef,
		Selector:     mg.Spec.InitProvider.LogGroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LogGroupName")
	}
	mg.Spec.InitProvider.LogGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LogGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Stream.
func (mg *Stream) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LogGroupNameRef,
		Selector:     mg.Spec.ForProvider.LogGroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogGroupName")
	}
	mg.Spec.ForProvider.LogGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LogGroupNameRef,
		Selector:     mg.Spec.InitProvider.LogGroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LogGroupName")
	}
	mg.Spec.InitProvider.LogGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LogGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionFilter.
func (mg *SubscriptionFilter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationArn),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.ForProvider.DestinationArnRef,
		Selector:     mg.Spec.ForProvider.DestinationArnSelector,
		To: reference.To{
			List:    &v1beta11.StreamList{},
			Managed: &v1beta11.Stream{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DestinationArn")
	}
	mg.Spec.ForProvider.DestinationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DestinationArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DestinationArn),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.InitProvider.DestinationArnRef,
		Selector:     mg.Spec.InitProvider.DestinationArnSelector,
		To: reference.To{
			List:    &v1beta11.StreamList{},
			Managed: &v1beta11.Stream{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DestinationArn")
	}
	mg.Spec.InitProvider.DestinationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DestinationArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}
