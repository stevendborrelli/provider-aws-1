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
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this App.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *App) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.IAMServiceRoleArnRef,
			Selector:     mg.Spec.ForProvider.IAMServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMServiceRoleArn")
	}
	mg.Spec.ForProvider.IAMServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMServiceRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IAMServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.IAMServiceRoleArnRef,
			Selector:     mg.Spec.InitProvider.IAMServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMServiceRoleArn")
	}
	mg.Spec.InitProvider.IAMServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IAMServiceRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackendEnvironment.
func (mg *BackendEnvironment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("amplify.aws.upbound.io", "v1beta2", "App", "AppList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AppIDRef,
			Selector:     mg.Spec.ForProvider.AppIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AppID")
	}
	mg.Spec.ForProvider.AppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AppIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Branch.
func (mg *Branch) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("amplify.aws.upbound.io", "v1beta2", "App", "AppList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AppIDRef,
			Selector:     mg.Spec.ForProvider.AppIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AppID")
	}
	mg.Spec.ForProvider.AppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AppIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Webhook.
func (mg *Webhook) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("amplify.aws.upbound.io", "v1beta2", "App", "AppList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AppIDRef,
			Selector:     mg.Spec.ForProvider.AppIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AppID")
	}
	mg.Spec.ForProvider.AppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AppIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("amplify.aws.upbound.io", "v1beta1", "Branch", "BranchList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BranchName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.BranchNameRef,
			Selector:     mg.Spec.ForProvider.BranchNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BranchName")
	}
	mg.Spec.ForProvider.BranchName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BranchNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("amplify.aws.upbound.io", "v1beta2", "App", "AppList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AppID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.AppIDRef,
			Selector:     mg.Spec.InitProvider.AppIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AppID")
	}
	mg.Spec.InitProvider.AppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AppIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("amplify.aws.upbound.io", "v1beta1", "Branch", "BranchList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BranchName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.BranchNameRef,
			Selector:     mg.Spec.InitProvider.BranchNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BranchName")
	}
	mg.Spec.InitProvider.BranchName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BranchNameRef = rsp.ResolvedReference

	return nil
}
