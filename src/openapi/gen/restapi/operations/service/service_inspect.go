// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ServiceInspectHandlerFunc turns a function with the right signature into a service inspect handler
type ServiceInspectHandlerFunc func(ServiceInspectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceInspectHandlerFunc) Handle(params ServiceInspectParams) middleware.Responder {
	return fn(params)
}

// ServiceInspectHandler interface for that can handle valid service inspect params
type ServiceInspectHandler interface {
	Handle(ServiceInspectParams) middleware.Responder
}

// NewServiceInspect creates a new http.Handler for the service inspect operation
func NewServiceInspect(ctx *middleware.Context, handler ServiceInspectHandler) *ServiceInspect {
	return &ServiceInspect{Context: ctx, Handler: handler}
}

/* ServiceInspect swagger:route GET /services/{id} Service serviceInspect

Inspect a service

*/
type ServiceInspect struct {
	Context *middleware.Context
	Handler ServiceInspectHandler
}

func (o *ServiceInspect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceInspectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}