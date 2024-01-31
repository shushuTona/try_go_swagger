// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// HelloHandlerFunc turns a function with the right signature into a hello handler
type HelloHandlerFunc func(HelloParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HelloHandlerFunc) Handle(params HelloParams) middleware.Responder {
	return fn(params)
}

// HelloHandler interface for that can handle valid hello params
type HelloHandler interface {
	Handle(HelloParams) middleware.Responder
}

// NewHello creates a new http.Handler for the hello operation
func NewHello(ctx *middleware.Context, handler HelloHandler) *Hello {
	return &Hello{Context: ctx, Handler: handler}
}

/*
	Hello swagger:route GET /hello hello

Hello hello API
*/
type Hello struct {
	Context *middleware.Context
	Handler HelloHandler
}

func (o *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewHelloParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
