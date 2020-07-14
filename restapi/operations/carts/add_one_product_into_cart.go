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

// AddOneProductIntoCartHandlerFunc turns a function with the right signature into a add one product into cart handler
type AddOneProductIntoCartHandlerFunc func(AddOneProductIntoCartParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOneProductIntoCartHandlerFunc) Handle(params AddOneProductIntoCartParams) middleware.Responder {
	return fn(params)
}

// AddOneProductIntoCartHandler interface for that can handle valid add one product into cart params
type AddOneProductIntoCartHandler interface {
	Handle(AddOneProductIntoCartParams) middleware.Responder
}

// NewAddOneProductIntoCart creates a new http.Handler for the add one product into cart operation
func NewAddOneProductIntoCart(ctx *middleware.Context, handler AddOneProductIntoCartHandler) *AddOneProductIntoCart {
	return &AddOneProductIntoCart{Context: ctx, Handler: handler}
}

/*AddOneProductIntoCart swagger:route POST /carts/{cart_id}/products carts addOneProductIntoCart

Add a Product into cart

*/
type AddOneProductIntoCart struct {
	Context *middleware.Context
	Handler AddOneProductIntoCartHandler
}

func (o *AddOneProductIntoCart) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddOneProductIntoCartParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddOneProductIntoCartBody add one product into cart body
//
// swagger:model AddOneProductIntoCartBody
type AddOneProductIntoCartBody struct {

	// product id
	// Required: true
	ProductID *int64 `json:"product_id"`
}

// Validate validates this add one product into cart body
func (o *AddOneProductIntoCartBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddOneProductIntoCartBody) validateProductID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"product_id", "body", o.ProductID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddOneProductIntoCartBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOneProductIntoCartBody) UnmarshalBinary(b []byte) error {
	var res AddOneProductIntoCartBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
