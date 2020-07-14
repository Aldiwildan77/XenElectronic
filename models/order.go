// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order order
//
// swagger:model Order
type Order struct {

	// cart id
	// Required: true
	// Format: uuid
	CartID *strfmt.UUID `json:"cart_id"`

	// customer address
	// Required: true
	// Min Length: 1
	CustomerAddress *string `json:"customer_address"`

	// customer email
	// Required: true
	// Min Length: 1
	// Format: email
	CustomerEmail *strfmt.Email `json:"customer_email"`

	// customer name
	// Required: true
	// Min Length: 1
	CustomerName *string `json:"customer_name"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCartID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateCartID(formats strfmt.Registry) error {

	if err := validate.Required("cart_id", "body", m.CartID); err != nil {
		return err
	}

	if err := validate.FormatOf("cart_id", "body", "uuid", m.CartID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCustomerAddress(formats strfmt.Registry) error {

	if err := validate.Required("customer_address", "body", m.CustomerAddress); err != nil {
		return err
	}

	if err := validate.MinLength("customer_address", "body", string(*m.CustomerAddress), 1); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCustomerEmail(formats strfmt.Registry) error {

	if err := validate.Required("customer_email", "body", m.CustomerEmail); err != nil {
		return err
	}

	if err := validate.MinLength("customer_email", "body", string(*m.CustomerEmail), 1); err != nil {
		return err
	}

	if err := validate.FormatOf("customer_email", "body", "email", m.CustomerEmail.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCustomerName(formats strfmt.Registry) error {

	if err := validate.Required("customer_name", "body", m.CustomerName); err != nil {
		return err
	}

	if err := validate.MinLength("customer_name", "body", string(*m.CustomerName), 1); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
