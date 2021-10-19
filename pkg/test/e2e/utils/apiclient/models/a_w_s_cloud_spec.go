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

// AWSCloudSpec AWSCloudSpec specifies access data to Amazon Web Services.
//
// swagger:model AWSCloudSpec
type AWSCloudSpec struct {

	// access key ID
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// assume role a r n
	AssumeRoleARN string `json:"assumeRoleARN,omitempty"`

	// assume role external ID
	AssumeRoleExternalID string `json:"assumeRoleExternalID,omitempty"`

	// The IAM role, the control plane will use. The control plane will perform an assume-role
	ControlPlaneRoleARN string `json:"roleARN,omitempty"`

	// instance profile name
	InstanceProfileName string `json:"instanceProfileName,omitempty"`

	// DEPRECATED. Don't care for the role name. We only require the ControlPlaneRoleARN to be set so the control plane
	// can perform the assume-role.
	// We keep it for backwards compatibility (We use this name for cleanup purpose).
	RoleName string `json:"roleName,omitempty"`

	// route table ID
	RouteTableID string `json:"routeTableId,omitempty"`

	// secret access key
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// security group ID
	SecurityGroupID string `json:"securityGroupID,omitempty"`

	// v p c ID
	VPCID string `json:"vpcId,omitempty"`

	// credentials reference
	CredentialsReference *GlobalSecretKeySelector `json:"credentialsReference,omitempty"`
}

// Validate validates this a w s cloud spec
func (m *AWSCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AWSCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsReference) { // not required
		return nil
	}

	if m.CredentialsReference != nil {
		if err := m.CredentialsReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentialsReference")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a w s cloud spec based on the context it is used
func (m *AWSCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentialsReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AWSCloudSpec) contextValidateCredentialsReference(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsReference != nil {
		if err := m.CredentialsReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentialsReference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AWSCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSCloudSpec) UnmarshalBinary(b []byte) error {
	var res AWSCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
