///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Driver driver
// swagger:model Driver

type Driver struct {

	// config
	Config DriverConfig `json:"config"`

	// created time
	// Read Only: true
	CreatedTime int64 `json:"created-time,omitempty"`

	// id
	// Read Only: true
	ID strfmt.UUID `json:"id,omitempty"`

	// modified time
	// Read Only: true
	ModifiedTime int64 `json:"modified-time,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// secrets
	Secrets []string `json:"secrets"`

	// status
	// Read Only: true
	Status Status `json:"status,omitempty"`

	// tags
	Tags DriverTags `json:"tags"`

	// type
	// Required: true
	// Max Length: 32
	Type *string `json:"type"`
}

/* polymorph Driver config false */

/* polymorph Driver created-time false */

/* polymorph Driver id false */

/* polymorph Driver modified-time false */

/* polymorph Driver name false */

/* polymorph Driver secrets false */

/* polymorph Driver status false */

/* polymorph Driver tags false */

/* polymorph Driver type false */

// Validate validates this driver
func (m *Driver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecrets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Driver) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Driver) validateSecrets(formats strfmt.Registry) error {

	if swag.IsZero(m.Secrets) { // not required
		return nil
	}

	return nil
}

func (m *Driver) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Driver) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", string(*m.Type), 32); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Driver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Driver) UnmarshalBinary(b []byte) error {
	var res Driver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
