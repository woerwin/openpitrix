// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAppVersion openpitrix app version
// swagger:model openpitrixAppVersion
type OpenpitrixAppVersion struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// package name
	PackageName string `json:"package_name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix app version
func (m *OpenpitrixAppVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAppVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAppVersion) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAppVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
