// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceapiGroupMeta deviceapi group meta
//
// swagger:model deviceapi.GroupMeta
type DeviceapiGroupMeta struct {

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this deviceapi group meta
func (m *DeviceapiGroupMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this deviceapi group meta based on context it is used
func (m *DeviceapiGroupMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceapiGroupMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceapiGroupMeta) UnmarshalBinary(b []byte) error {
	var res DeviceapiGroupMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
