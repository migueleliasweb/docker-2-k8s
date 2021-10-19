// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NetworkInspectHandlerFunc turns a function with the right signature into a network inspect handler
type NetworkInspectHandlerFunc func(NetworkInspectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NetworkInspectHandlerFunc) Handle(params NetworkInspectParams) middleware.Responder {
	return fn(params)
}

// NetworkInspectHandler interface for that can handle valid network inspect params
type NetworkInspectHandler interface {
	Handle(NetworkInspectParams) middleware.Responder
}

// NewNetworkInspect creates a new http.Handler for the network inspect operation
func NewNetworkInspect(ctx *middleware.Context, handler NetworkInspectHandler) *NetworkInspect {
	return &NetworkInspect{Context: ctx, Handler: handler}
}

/* NetworkInspect swagger:route GET /networks/{id} Network networkInspect

Inspect a network

*/
type NetworkInspect struct {
	Context *middleware.Context
	Handler NetworkInspectHandler
}

func (o *NetworkInspect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewNetworkInspectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}