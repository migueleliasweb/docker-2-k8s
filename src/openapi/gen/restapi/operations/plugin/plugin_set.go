// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PluginSetHandlerFunc turns a function with the right signature into a plugin set handler
type PluginSetHandlerFunc func(PluginSetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PluginSetHandlerFunc) Handle(params PluginSetParams) middleware.Responder {
	return fn(params)
}

// PluginSetHandler interface for that can handle valid plugin set params
type PluginSetHandler interface {
	Handle(PluginSetParams) middleware.Responder
}

// NewPluginSet creates a new http.Handler for the plugin set operation
func NewPluginSet(ctx *middleware.Context, handler PluginSetHandler) *PluginSet {
	return &PluginSet{Context: ctx, Handler: handler}
}

/* PluginSet swagger:route POST /plugins/{name}/set Plugin pluginSet

Configure a plugin

*/
type PluginSet struct {
	Context *middleware.Context
	Handler PluginSetHandler
}

func (o *PluginSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPluginSetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
