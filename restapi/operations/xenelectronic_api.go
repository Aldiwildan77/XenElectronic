// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/gifff/xenelectronic/restapi/operations/carts"
	"github.com/gifff/xenelectronic/restapi/operations/categories"
	"github.com/gifff/xenelectronic/restapi/operations/orders"
)

// NewXenelectronicAPI creates a new Xenelectronic instance
func NewXenelectronicAPI(spec *loads.Document) *XenelectronicAPI {
	return &XenelectronicAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CartsAddOneProductIntoCartHandler: carts.AddOneProductIntoCartHandlerFunc(func(params carts.AddOneProductIntoCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.AddOneProductIntoCart has not yet been implemented")
		}),
		OrdersCheckoutHandler: orders.CheckoutHandlerFunc(func(params orders.CheckoutParams) middleware.Responder {
			return middleware.NotImplemented("operation orders.Checkout has not yet been implemented")
		}),
		CartsCreateCartHandler: carts.CreateCartHandlerFunc(func(params carts.CreateCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.CreateCart has not yet been implemented")
		}),
		CategoriesListCategoriesHandler: categories.ListCategoriesHandlerFunc(func(params categories.ListCategoriesParams) middleware.Responder {
			return middleware.NotImplemented("operation categories.ListCategories has not yet been implemented")
		}),
		CartsListProductsInCartHandler: carts.ListProductsInCartHandlerFunc(func(params carts.ListProductsInCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.ListProductsInCart has not yet been implemented")
		}),
		CategoriesListProductsOfCategoryHandler: categories.ListProductsOfCategoryHandlerFunc(func(params categories.ListProductsOfCategoryParams) middleware.Responder {
			return middleware.NotImplemented("operation categories.ListProductsOfCategory has not yet been implemented")
		}),
		CartsRemoveOneProductFromCartHandler: carts.RemoveOneProductFromCartHandlerFunc(func(params carts.RemoveOneProductFromCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.RemoveOneProductFromCart has not yet been implemented")
		}),
		OrdersViewOneOrderHandler: orders.ViewOneOrderHandlerFunc(func(params orders.ViewOneOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation orders.ViewOneOrder has not yet been implemented")
		}),
	}
}

/*XenelectronicAPI e-Commerce application for XenElectronic */
type XenelectronicAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/xenelectronic.v1+json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/xenelectronic.v1+json
	JSONProducer runtime.Producer

	// CartsAddOneProductIntoCartHandler sets the operation handler for the add one product into cart operation
	CartsAddOneProductIntoCartHandler carts.AddOneProductIntoCartHandler
	// OrdersCheckoutHandler sets the operation handler for the checkout operation
	OrdersCheckoutHandler orders.CheckoutHandler
	// CartsCreateCartHandler sets the operation handler for the create cart operation
	CartsCreateCartHandler carts.CreateCartHandler
	// CategoriesListCategoriesHandler sets the operation handler for the list categories operation
	CategoriesListCategoriesHandler categories.ListCategoriesHandler
	// CartsListProductsInCartHandler sets the operation handler for the list products in cart operation
	CartsListProductsInCartHandler carts.ListProductsInCartHandler
	// CategoriesListProductsOfCategoryHandler sets the operation handler for the list products of category operation
	CategoriesListProductsOfCategoryHandler categories.ListProductsOfCategoryHandler
	// CartsRemoveOneProductFromCartHandler sets the operation handler for the remove one product from cart operation
	CartsRemoveOneProductFromCartHandler carts.RemoveOneProductFromCartHandler
	// OrdersViewOneOrderHandler sets the operation handler for the view one order operation
	OrdersViewOneOrderHandler orders.ViewOneOrderHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *XenelectronicAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *XenelectronicAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *XenelectronicAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *XenelectronicAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *XenelectronicAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *XenelectronicAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *XenelectronicAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the XenelectronicAPI
func (o *XenelectronicAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CartsAddOneProductIntoCartHandler == nil {
		unregistered = append(unregistered, "carts.AddOneProductIntoCartHandler")
	}
	if o.OrdersCheckoutHandler == nil {
		unregistered = append(unregistered, "orders.CheckoutHandler")
	}
	if o.CartsCreateCartHandler == nil {
		unregistered = append(unregistered, "carts.CreateCartHandler")
	}
	if o.CategoriesListCategoriesHandler == nil {
		unregistered = append(unregistered, "categories.ListCategoriesHandler")
	}
	if o.CartsListProductsInCartHandler == nil {
		unregistered = append(unregistered, "carts.ListProductsInCartHandler")
	}
	if o.CategoriesListProductsOfCategoryHandler == nil {
		unregistered = append(unregistered, "categories.ListProductsOfCategoryHandler")
	}
	if o.CartsRemoveOneProductFromCartHandler == nil {
		unregistered = append(unregistered, "carts.RemoveOneProductFromCartHandler")
	}
	if o.OrdersViewOneOrderHandler == nil {
		unregistered = append(unregistered, "orders.ViewOneOrderHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *XenelectronicAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *XenelectronicAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *XenelectronicAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *XenelectronicAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/xenelectronic.v1+json":
			result["application/xenelectronic.v1+json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *XenelectronicAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/xenelectronic.v1+json":
			result["application/xenelectronic.v1+json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *XenelectronicAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the xenelectronic API
func (o *XenelectronicAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *XenelectronicAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/carts/{cart_id}/products"] = carts.NewAddOneProductIntoCart(o.context, o.CartsAddOneProductIntoCartHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/orders"] = orders.NewCheckout(o.context, o.OrdersCheckoutHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/carts"] = carts.NewCreateCart(o.context, o.CartsCreateCartHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/categories"] = categories.NewListCategories(o.context, o.CategoriesListCategoriesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/carts/{cart_id}/products"] = carts.NewListProductsInCart(o.context, o.CartsListProductsInCartHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/categories/{category_id}/products"] = categories.NewListProductsOfCategory(o.context, o.CategoriesListProductsOfCategoryHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/carts/{cart_id}/products/{product_id}"] = carts.NewRemoveOneProductFromCart(o.context, o.CartsRemoveOneProductFromCartHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/orders/{orderId}"] = orders.NewViewOneOrder(o.context, o.OrdersViewOneOrderHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *XenelectronicAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *XenelectronicAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *XenelectronicAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *XenelectronicAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *XenelectronicAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
