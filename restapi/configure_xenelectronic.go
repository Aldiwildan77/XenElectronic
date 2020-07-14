// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gifff/xenelectronic/restapi/operations"
	"github.com/gifff/xenelectronic/restapi/operations/carts"
	"github.com/gifff/xenelectronic/restapi/operations/categories"
	"github.com/gifff/xenelectronic/restapi/operations/orders"
)

//go:generate swagger generate server --target ../../xenelectronic --name Xenelectronic --spec ../swagger.yml

func configureFlags(api *operations.XenelectronicAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.XenelectronicAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CartsAddOneProductIntoCartHandler == nil {
		api.CartsAddOneProductIntoCartHandler = carts.AddOneProductIntoCartHandlerFunc(func(params carts.AddOneProductIntoCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.AddOneProductIntoCart has not yet been implemented")
		})
	}
	if api.OrdersCheckoutHandler == nil {
		api.OrdersCheckoutHandler = orders.CheckoutHandlerFunc(func(params orders.CheckoutParams) middleware.Responder {
			return middleware.NotImplemented("operation orders.Checkout has not yet been implemented")
		})
	}
	if api.CartsCreateCartHandler == nil {
		api.CartsCreateCartHandler = carts.CreateCartHandlerFunc(func(params carts.CreateCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.CreateCart has not yet been implemented")
		})
	}
	if api.CategoriesListCategoriesHandler == nil {
		api.CategoriesListCategoriesHandler = categories.ListCategoriesHandlerFunc(func(params categories.ListCategoriesParams) middleware.Responder {
			return middleware.NotImplemented("operation categories.ListCategories has not yet been implemented")
		})
	}
	if api.CartsListProductsInCartHandler == nil {
		api.CartsListProductsInCartHandler = carts.ListProductsInCartHandlerFunc(func(params carts.ListProductsInCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.ListProductsInCart has not yet been implemented")
		})
	}
	if api.CategoriesListProductsOfCategoryHandler == nil {
		api.CategoriesListProductsOfCategoryHandler = categories.ListProductsOfCategoryHandlerFunc(func(params categories.ListProductsOfCategoryParams) middleware.Responder {
			return middleware.NotImplemented("operation categories.ListProductsOfCategory has not yet been implemented")
		})
	}
	if api.CartsRemoveOneProductFromCartHandler == nil {
		api.CartsRemoveOneProductFromCartHandler = carts.RemoveOneProductFromCartHandlerFunc(func(params carts.RemoveOneProductFromCartParams) middleware.Responder {
			return middleware.NotImplemented("operation carts.RemoveOneProductFromCart has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
