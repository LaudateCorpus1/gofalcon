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

// MessagesCase messages case
//
// swagger:model messages.Case
type MessagesCase struct {

	// aids
	// Required: true
	Aids []string `json:"aids"`

	// assigner
	// Required: true
	Assigner *MessagesAuthor `json:"assigner"`

	// attachments
	// Required: true
	Attachments []*MessagesAttachment `json:"attachments"`

	// body
	// Required: true
	Body *string `json:"body"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// created time
	// Required: true
	CreatedTime *string `json:"created_time"`

	// detections
	// Required: true
	Detections []*MessagesDetection `json:"detections"`

	// hosts
	// Required: true
	Hosts []string `json:"hosts"`

	// id
	// Required: true
	ID *string `json:"id"`

	// incidents
	// Required: true
	Incidents []*MessagesIncident `json:"incidents"`

	// ip addresses
	// Required: true
	IPAddresses []string `json:"ip_addresses"`

	// key
	// Required: true
	Key *string `json:"key"`

	// last modified time
	// Required: true
	LastModifiedTime *string `json:"last_modified_time"`

	// status
	// Required: true
	Status *string `json:"status"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this messages case
func (m *MessagesCase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssigner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagesCase) validateAids(formats strfmt.Registry) error {

	if err := validate.Required("aids", "body", m.Aids); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateAssigner(formats strfmt.Registry) error {

	if err := validate.Required("assigner", "body", m.Assigner); err != nil {
		return err
	}

	if m.Assigner != nil {
		if err := m.Assigner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assigner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assigner")
			}
			return err
		}
	}

	return nil
}

func (m *MessagesCase) validateAttachments(formats strfmt.Registry) error {

	if err := validate.Required("attachments", "body", m.Attachments); err != nil {
		return err
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagesCase) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateCreatedTime(formats strfmt.Registry) error {

	if err := validate.Required("created_time", "body", m.CreatedTime); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateDetections(formats strfmt.Registry) error {

	if err := validate.Required("detections", "body", m.Detections); err != nil {
		return err
	}

	for i := 0; i < len(m.Detections); i++ {
		if swag.IsZero(m.Detections[i]) { // not required
			continue
		}

		if m.Detections[i] != nil {
			if err := m.Detections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagesCase) validateHosts(formats strfmt.Registry) error {

	if err := validate.Required("hosts", "body", m.Hosts); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateIncidents(formats strfmt.Registry) error {

	if err := validate.Required("incidents", "body", m.Incidents); err != nil {
		return err
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagesCase) validateIPAddresses(formats strfmt.Registry) error {

	if err := validate.Required("ip_addresses", "body", m.IPAddresses); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateLastModifiedTime(formats strfmt.Registry) error {

	if err := validate.Required("last_modified_time", "body", m.LastModifiedTime); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *MessagesCase) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this messages case based on the context it is used
func (m *MessagesCase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssigner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDetections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncidents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagesCase) contextValidateAssigner(ctx context.Context, formats strfmt.Registry) error {

	if m.Assigner != nil {
		if err := m.Assigner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assigner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assigner")
			}
			return err
		}
	}

	return nil
}

func (m *MessagesCase) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagesCase) contextValidateDetections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Detections); i++ {

		if m.Detections[i] != nil {
			if err := m.Detections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagesCase) contextValidateIncidents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Incidents); i++ {

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessagesCase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessagesCase) UnmarshalBinary(b []byte) error {
	var res MessagesCase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
