// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	ujconversion "github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/crossplane/upjet/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this Datasource to the hub type.
func (tr *Datasource) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Datasource type.
func (tr *Datasource) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this Function to the hub type.
func (tr *Function) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Function type.
func (tr *Function) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this GraphQLAPI to the hub type.
func (tr *GraphQLAPI) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the GraphQLAPI type.
func (tr *GraphQLAPI) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this Resolver to the hub type.
func (tr *Resolver) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Resolver type.
func (tr *Resolver) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}
