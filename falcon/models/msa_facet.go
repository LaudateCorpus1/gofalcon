// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsaFacet msa facet
//
// swagger:model msa.Facet
type MsaFacet struct {

	// by
	By []*MsaFacet `json:"by"`

	// count
	// Required: true
	Count *int64 `json:"count"`

	// facet
	Facet string `json:"facet,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// term
	// Required: true
	Term *string `json:"term"`
}

// Validate validates this msa facet
func (m *MsaFacet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaFacet) validateBy(formats strfmt.Registry) error {
	if swag.IsZero(m.By) { // not required
		return nil
	}

	for i := 0; i < len(m.By); i++ {
		if swag.IsZero(m.By[i]) { // not required
			continue
		}

		if m.By[i] != nil {
			if err := m.By[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("by" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaFacet) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *MsaFacet) validateTerm(formats strfmt.Registry) error {

	if err := validate.Required("term", "body", m.Term); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this msa facet based on the context it is used
func (m *MsaFacet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaFacet) contextValidateBy(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.By); i++ {

		if m.By[i] != nil {
			if err := m.By[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("by" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsaFacet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsaFacet) UnmarshalBinary(b []byte) error {
	var res MsaFacet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}