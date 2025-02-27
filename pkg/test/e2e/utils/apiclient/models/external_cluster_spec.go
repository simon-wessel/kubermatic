// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExternalClusterSpec ExternalClusterSpec defines the external cluster specification.
//
// swagger:model ExternalClusterSpec
type ExternalClusterSpec struct {

	// version
	Version Semver `json:"version,omitempty"`
}

// Validate validates this external cluster spec
func (m *ExternalClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClusterSpec) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := m.Version.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("version")
		}
		return err
	}

	return nil
}

// ContextValidate validate this external cluster spec based on the context it is used
func (m *ExternalClusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClusterSpec) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Version.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("version")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalClusterSpec) UnmarshalBinary(b []byte) error {
	var res ExternalClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
