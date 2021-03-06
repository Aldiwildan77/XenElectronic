// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewRemoveOneProductFromCartParams creates a new RemoveOneProductFromCartParams object
// no default values defined in spec.
func NewRemoveOneProductFromCartParams() RemoveOneProductFromCartParams {

	return RemoveOneProductFromCartParams{}
}

// RemoveOneProductFromCartParams contains all the bound params for the remove one product from cart operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeOneProductFromCart
type RemoveOneProductFromCartParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	CartID strfmt.UUID
	/*
	  Required: true
	  In: path
	*/
	ProductID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveOneProductFromCartParams() beforehand.
func (o *RemoveOneProductFromCartParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCartID, rhkCartID, _ := route.Params.GetOK("cart_id")
	if err := o.bindCartID(rCartID, rhkCartID, route.Formats); err != nil {
		res = append(res, err)
	}

	rProductID, rhkProductID, _ := route.Params.GetOK("product_id")
	if err := o.bindProductID(rProductID, rhkProductID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCartID binds and validates parameter CartID from path.
func (o *RemoveOneProductFromCartParams) bindCartID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("cart_id", "path", "strfmt.UUID", raw)
	}
	o.CartID = *(value.(*strfmt.UUID))

	if err := o.validateCartID(formats); err != nil {
		return err
	}

	return nil
}

// validateCartID carries on validations for parameter CartID
func (o *RemoveOneProductFromCartParams) validateCartID(formats strfmt.Registry) error {

	if err := validate.FormatOf("cart_id", "path", "uuid", o.CartID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindProductID binds and validates parameter ProductID from path.
func (o *RemoveOneProductFromCartParams) bindProductID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("product_id", "path", "int64", raw)
	}
	o.ProductID = value

	return nil
}
