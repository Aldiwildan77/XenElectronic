// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateCartHandlerFunc turns a function with the right signature into a create cart handler
type CreateCartHandlerFunc func(CreateCartParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCartHandlerFunc) Handle(params CreateCartParams) middleware.Responder {
	return fn(params)
}

// CreateCartHandler interface for that can handle valid create cart params
type CreateCartHandler interface {
	Handle(CreateCartParams) middleware.Responder
}

// NewCreateCart creates a new http.Handler for the create cart operation
func NewCreateCart(ctx *middleware.Context, handler CreateCartHandler) *CreateCart {
	return &CreateCart{Context: ctx, Handler: handler}
}

/*CreateCart swagger:route POST /carts carts createCart

Create new cart

*/
type CreateCart struct {
	Context *middleware.Context
	Handler CreateCartHandler
}

func (o *CreateCart) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCartParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateCartOKBody create cart o k body
//
// swagger:model CreateCartOKBody
type CreateCartOKBody struct {

	// cart id
	// Required: true
	// Format: uuid
	CartID *strfmt.UUID `json:"cart_id"`
}

// Validate validates this create cart o k body
func (o *CreateCartOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCartID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCartOKBody) validateCartID(formats strfmt.Registry) error {

	if err := validate.Required("createCartOK"+"."+"cart_id", "body", o.CartID); err != nil {
		return err
	}

	if err := validate.FormatOf("createCartOK"+"."+"cart_id", "body", "uuid", o.CartID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCartOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCartOKBody) UnmarshalBinary(b []byte) error {
	var res CreateCartOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
