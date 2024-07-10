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
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Connector) ResolveReferences( // ResolveReferences of this Connector.
	ctx context.Context, c client.Reader) error {
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
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessRole),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.AccessRoleRef,
			Selector:     mg.Spec.ForProvider.AccessRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessRole")
	}
	mg.Spec.ForProvider.AccessRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessRoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SftpConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("secretsmanager.aws.upbound.io", "v1beta1", "Secret", "SecretList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SftpConfig[i3].UserSecretID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SftpConfig[i3].UserSecretIDRef,
				Selector:     mg.Spec.ForProvider.SftpConfig[i3].UserSecretIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SftpConfig[i3].UserSecretID")
		}
		mg.Spec.ForProvider.SftpConfig[i3].UserSecretID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SftpConfig[i3].UserSecretIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccessRole),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.AccessRoleRef,
			Selector:     mg.Spec.InitProvider.AccessRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AccessRole")
	}
	mg.Spec.InitProvider.AccessRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AccessRoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.SftpConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("secretsmanager.aws.upbound.io", "v1beta1", "Secret", "SecretList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SftpConfig[i3].UserSecretID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.SftpConfig[i3].UserSecretIDRef,
				Selector:     mg.Spec.InitProvider.SftpConfig[i3].UserSecretIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SftpConfig[i3].UserSecretID")
		}
		mg.Spec.InitProvider.SftpConfig[i3].UserSecretID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SftpConfig[i3].UserSecretIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this SSHKey.
func (mg *SSHKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ServerIDRef,
			Selector:     mg.Spec.ForProvider.ServerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserNameRef,
			Selector:     mg.Spec.ForProvider.UserNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ServerIDRef,
			Selector:     mg.Spec.InitProvider.ServerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.UserName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserNameRef,
			Selector:     mg.Spec.InitProvider.UserNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.UserName")
	}
	mg.Spec.InitProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Certificate),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.CertificateRef,
			Selector:     mg.Spec.ForProvider.CertificateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Certificate")
	}
	mg.Spec.ForProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta1", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DirectoryIDRef,
			Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EndpointDetails); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointDetails[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.EndpointDetails[i3].VPCIDRef,
				Selector:     mg.Spec.ForProvider.EndpointDetails[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EndpointDetails[i3].VPCID")
		}
		mg.Spec.ForProvider.EndpointDetails[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EndpointDetails[i3].VPCIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingRole),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.LoggingRoleRef,
			Selector:     mg.Spec.ForProvider.LoggingRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoggingRole")
	}
	mg.Spec.ForProvider.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoggingRoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Certificate),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.CertificateRef,
			Selector:     mg.Spec.InitProvider.CertificateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Certificate")
	}
	mg.Spec.InitProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta1", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DirectoryIDRef,
			Selector:     mg.Spec.InitProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DirectoryID")
	}
	mg.Spec.InitProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DirectoryIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.EndpointDetails); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EndpointDetails[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.EndpointDetails[i3].VPCIDRef,
				Selector:     mg.Spec.InitProvider.EndpointDetails[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EndpointDetails[i3].VPCID")
		}
		mg.Spec.InitProvider.EndpointDetails[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EndpointDetails[i3].VPCIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingRole),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.LoggingRoleRef,
			Selector:     mg.Spec.InitProvider.LoggingRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoggingRole")
	}
	mg.Spec.InitProvider.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoggingRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Tag.
func (mg *Tag) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ResourceArnRef,
			Selector:     mg.Spec.ForProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ResourceArnRef,
			Selector:     mg.Spec.InitProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceArn")
	}
	mg.Spec.InitProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
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
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta1", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ServerIDRef,
			Selector:     mg.Spec.ForProvider.ServerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleRef,
			Selector:     mg.Spec.InitProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta1", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ServerIDRef,
			Selector:     mg.Spec.InitProvider.ServerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workflow.
func (mg *Workflow) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Steps); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Steps[i3].CustomStepDetails); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].Target),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].TargetRef,
					Selector:     mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].TargetSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].Target")
			}
			mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].Target = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Steps[i3].CustomStepDetails[i4].TargetRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Steps); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Steps[i3].CustomStepDetails); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].Target),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].TargetRef,
					Selector:     mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].TargetSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].Target")
			}
			mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].Target = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Steps[i3].CustomStepDetails[i4].TargetRef = rsp.ResolvedReference

		}
	}

	return nil
}
