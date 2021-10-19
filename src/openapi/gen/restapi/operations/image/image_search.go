// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageSearchHandlerFunc turns a function with the right signature into a image search handler
type ImageSearchHandlerFunc func(ImageSearchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ImageSearchHandlerFunc) Handle(params ImageSearchParams) middleware.Responder {
	return fn(params)
}

// ImageSearchHandler interface for that can handle valid image search params
type ImageSearchHandler interface {
	Handle(ImageSearchParams) middleware.Responder
}

// NewImageSearch creates a new http.Handler for the image search operation
func NewImageSearch(ctx *middleware.Context, handler ImageSearchHandler) *ImageSearch {
	return &ImageSearch{Context: ctx, Handler: handler}
}

/* ImageSearch swagger:route GET /images/search Image imageSearch

Search images

Search for an image on Docker Hub.

*/
type ImageSearch struct {
	Context *middleware.Context
	Handler ImageSearchHandler
}

func (o *ImageSearch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewImageSearchParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ImageSearchOKBodyItems0 ImageSearchResponseItem
//
// swagger:model ImageSearchOKBodyItems0
type ImageSearchOKBodyItems0 struct {

	// description
	Description string `json:"description,omitempty"`

	// is automated
	IsAutomated bool `json:"is_automated,omitempty"`

	// is official
	IsOfficial bool `json:"is_official,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// star count
	StarCount int64 `json:"star_count,omitempty"`
}

// Validate validates this image search o k body items0
func (o *ImageSearchOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image search o k body items0 based on context it is used
func (o *ImageSearchOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageSearchOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageSearchOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ImageSearchOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}