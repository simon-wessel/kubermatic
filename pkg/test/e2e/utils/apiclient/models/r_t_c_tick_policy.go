// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// RTCTickPolicy RTCTickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.
//
// swagger:model RTCTickPolicy
type RTCTickPolicy string

// Validate validates this r t c tick policy
func (m RTCTickPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r t c tick policy based on context it is used
func (m RTCTickPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
