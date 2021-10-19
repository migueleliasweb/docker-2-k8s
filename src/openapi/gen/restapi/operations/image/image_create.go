// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ImageCreateHandlerFunc turns a function with the right signature into a image create handler
type ImageCreateHandlerFunc func(ImageCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ImageCreateHandlerFunc) Handle(params ImageCreateParams) middleware.Responder {
	return fn(params)
}

// ImageCreateHandler interface for that can handle valid image create params
type ImageCreateHandler interface {
	Handle(ImageCreateParams) middleware.Responder
}

// NewImageCreate creates a new http.Handler for the image create operation
func NewImageCreate(ctx *middleware.Context, handler ImageCreateHandler) *ImageCreate {
	return &ImageCreate{Context: ctx, Handler: handler}
}

/* ImageCreate swagger:route POST /images/create Image imageCreate

Create an image

Create an image by either pulling it from a registry or importing it.

*/
type ImageCreate struct {
	Context *middleware.Context
	Handler ImageCreateHandler
}

func (o *ImageCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewImageCreateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
