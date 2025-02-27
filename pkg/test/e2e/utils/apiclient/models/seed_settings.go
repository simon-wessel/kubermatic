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

// SeedSettings SeedSettings represents settings for a Seed cluster
//
// swagger:model SeedSettings
type SeedSettings struct {

	// the Seed level seed dns overwrite
	SeedDNSOverwrite string `json:"seedDNSOverwrite,omitempty"`

	// metering
	Metering *MeteringConfiguration `json:"metering,omitempty"`

	// mla
	Mla *MLA `json:"mla,omitempty"`
}

// Validate validates this seed settings
func (m *SeedSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMla(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SeedSettings) validateMetering(formats strfmt.Registry) error {
	if swag.IsZero(m.Metering) { // not required
		return nil
	}

	if m.Metering != nil {
		if err := m.Metering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metering")
			}
			return err
		}
	}

	return nil
}

func (m *SeedSettings) validateMla(formats strfmt.Registry) error {
	if swag.IsZero(m.Mla) { // not required
		return nil
	}

	if m.Mla != nil {
		if err := m.Mla.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mla")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this seed settings based on the context it is used
func (m *SeedSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMla(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SeedSettings) contextValidateMetering(ctx context.Context, formats strfmt.Registry) error {

	if m.Metering != nil {
		if err := m.Metering.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metering")
			}
			return err
		}
	}

	return nil
}

func (m *SeedSettings) contextValidateMla(ctx context.Context, formats strfmt.Registry) error {

	if m.Mla != nil {
		if err := m.Mla.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mla")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SeedSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeedSettings) UnmarshalBinary(b []byte) error {
	var res SeedSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
