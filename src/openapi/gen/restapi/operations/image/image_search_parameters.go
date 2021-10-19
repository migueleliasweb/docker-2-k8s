// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewImageSearchParams creates a new ImageSearchParams object
//
// There are no default values defined in the spec.
func NewImageSearchParams() ImageSearchParams {

	return ImageSearchParams{}
}

// ImageSearchParams contains all the bound params for the image search operation
// typically these are obtained from a http.Request
//
// swagger:parameters ImageSearch
type ImageSearchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A JSON encoded value of the filters (a `map[string][]string`) to process on the images list. Available filters:

	- `is-automated=(true|false)`
	- `is-official=(true|false)`
	- `stars=<number>` Matches images that has at least 'number' stars.

	  In: query
	*/
	Filters *string
	/*Maximum number of results to return
	  In: query
	*/
	Limit *int64
	/*Term to search
	  Required: true
	  In: query
	*/
	Term string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewImageSearchParams() beforehand.
func (o *ImageSearchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFilters, qhkFilters, _ := qs.GetOK("filters")
	if err := o.bindFilters(qFilters, qhkFilters, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qTerm, qhkTerm, _ := qs.GetOK("term")
	if err := o.bindTerm(qTerm, qhkTerm, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFilters binds and validates parameter Filters from query.
func (o *ImageSearchParams) bindFilters(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Filters = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *ImageSearchParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindTerm binds and validates parameter Term from query.
func (o *ImageSearchParams) bindTerm(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("term", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("term", "query", raw); err != nil {
		return err
	}
	o.Term = raw

	return nil
}
