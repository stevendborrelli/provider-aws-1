// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Application) ResolveReferences( // ResolveReferences of this Application.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AppversionLifecycle); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppversionLifecycle[i3].ServiceRole),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.AppversionLifecycle[i3].ServiceRoleRef,
				Selector:     mg.Spec.ForProvider.AppversionLifecycle[i3].ServiceRoleSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AppversionLifecycle[i3].ServiceRole")
		}
		mg.Spec.ForProvider.AppversionLifecycle[i3].ServiceRole = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AppversionLifecycle[i3].ServiceRoleRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.AppversionLifecycle); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AppversionLifecycle[i3].ServiceRole),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.AppversionLifecycle[i3].ServiceRoleRef,
				Selector:     mg.Spec.InitProvider.AppversionLifecycle[i3].ServiceRoleSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AppversionLifecycle[i3].ServiceRole")
		}
		mg.Spec.InitProvider.AppversionLifecycle[i3].ServiceRole = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.AppversionLifecycle[i3].ServiceRoleRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ApplicationVersion.
func (mg *ApplicationVersion) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.BucketRef,
			Selector:     mg.Spec.ForProvider.BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Object", "ObjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Key),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyRef,
			Selector:     mg.Spec.ForProvider.KeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Key")
	}
	mg.Spec.ForProvider.Key = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.BucketRef,
			Selector:     mg.Spec.InitProvider.BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Object", "ObjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Key),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyRef,
			Selector:     mg.Spec.InitProvider.KeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Key")
	}
	mg.Spec.InitProvider.Key = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ConfigurationTemplate.
func (mg *ConfigurationTemplate) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("elasticbeanstalk.aws.upbound.io", "v1beta2", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Application),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ApplicationRef,
			Selector:     mg.Spec.ForProvider.ApplicationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Application")
	}
	mg.Spec.ForProvider.Application = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elasticbeanstalk.aws.upbound.io", "v1beta2", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Application),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ApplicationRef,
			Selector:     mg.Spec.InitProvider.ApplicationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Application")
	}
	mg.Spec.InitProvider.Application = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationRef = rsp.ResolvedReference

	return nil
}
