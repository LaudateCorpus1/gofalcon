// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainBreachedItemV1 domain breached item v1
//
// swagger:model domain.BreachedItemV1
type DomainBreachedItemV1 struct {

	// The domain associated with the breached account.
	// Required: true
	Domain *string `json:"domain"`

	// The email of the breached account.
	// Required: true
	Email *string `json:"email"`

	// The original hashing algorithm applied to the breached password. Possible values: 'plain', 'unknown', 'base64', 'md5', 'sha1', 'bcrypt', etc. The value 'plain' means that the password was originally found as plaintext.
	// Required: true
	HashType *string `json:"hash_type"`

	// The username of the breached account.
	// Required: true
	LoginID *string `json:"login_id"`

	// The name of the person associated with the breached account.
	// Required: true
	Name *string `json:"name"`

	// The breached password. Passwords are returned as salted hashes, generated using the SHA256 algorithm and the CID as the salt.
	// Required: true
	Password *string `json:"password"`

	// The phone number of the person associated with the breached account.
	// Required: true
	Phone *string `json:"phone"`
}

// Validate validates this domain breached item v1
func (m *DomainBreachedItemV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoginID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainBreachedItemV1) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *DomainBreachedItemV1) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *DomainBreachedItemV1) validateHashType(formats strfmt.Registry) error {

	if err := validate.Required("hash_type", "body", m.HashType); err != nil {
		return err
	}

	return nil
}

func (m *DomainBreachedItemV1) validateLoginID(formats strfmt.Registry) error {

	if err := validate.Required("login_id", "body", m.LoginID); err != nil {
		return err
	}

	return nil
}

func (m *DomainBreachedItemV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainBreachedItemV1) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *DomainBreachedItemV1) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("phone", "body", m.Phone); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain breached item v1 based on context it is used
func (m *DomainBreachedItemV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainBreachedItemV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainBreachedItemV1) UnmarshalBinary(b []byte) error {
	var res DomainBreachedItemV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
