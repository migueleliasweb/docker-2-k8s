// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ImagePushHandlerFunc turns a function with the right signature into a image push handler
type ImagePushHandlerFunc func(ImagePushParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ImagePushHandlerFunc) Handle(params ImagePushParams) middleware.Responder {
	return fn(params)
}

// ImagePushHandler interface for that can handle valid image push params
type ImagePushHandler interface {
	Handle(ImagePushParams) middleware.Responder
}

// NewImagePush creates a new http.Handler for the image push operation
func NewImagePush(ctx *middleware.Context, handler ImagePushHandler) *ImagePush {
	return &ImagePush{Context: ctx, Handler: handler}
}

/* ImagePush swagger:route POST /images/{name}/push Image imagePush

Push an image

Push an image to a registry.

If you wish to push an image on to a private registry, that image must
already have a tag which references the registry. For example,
`registry.example.com/myimage:latest`.

The push is cancelled if the HTTP connection is closed.


*/
type ImagePush struct {
	Context *middleware.Context
	Handler ImagePushHandler
}

func (o *ImagePush) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewImagePushParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}