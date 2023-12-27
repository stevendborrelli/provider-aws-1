/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Table.
func (mg *Table) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyspaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyspaceNameRef,
		Selector:     mg.Spec.ForProvider.KeyspaceNameSelector,
		To: reference.To{
			List:    &KeyspaceList{},
			Managed: &Keyspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyspaceName")
	}
	mg.Spec.ForProvider.KeyspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyspaceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyspaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KeyspaceNameRef,
		Selector:     mg.Spec.InitProvider.KeyspaceNameSelector,
		To: reference.To{
			List:    &KeyspaceList{},
			Managed: &Keyspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyspaceName")
	}
	mg.Spec.InitProvider.KeyspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyspaceNameRef = rsp.ResolvedReference

	return nil
}
