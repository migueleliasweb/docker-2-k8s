// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPluginPullParams creates a new PluginPullParams object
//
// There are no default values defined in the spec.
func NewPluginPullParams() PluginPullParams {

	return PluginPullParams{}
}

// PluginPullParams contains all the bound params for the plugin pull operation
// typically these are obtained from a http.Request
//
// swagger:parameters PluginPull
type PluginPullParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A base64url-encoded auth configuration to use when pulling a plugin
	from a registry.

	Refer to the [authentication section](#section/Authentication) for
	details.

	  In: header
	*/
	XRegistryAuth *string
	/*
	  In: body
	*/
	Body []*PluginPullParamsBodyItems0
	/*Local name for the pulled plugin.

	The `:latest` tag is optional, and is used as the default if omitted.

	  In: query
	*/
	Name *string
	/*Remote reference for plugin to install.

	The `:latest` tag is optional, and is used as the default if omitted.

	  Required: true
	  In: query
	*/
	Remote string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPluginPullParams() beforehand.
func (o *PluginPullParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXRegistryAuth(r.Header[http.CanonicalHeaderKey("X-Registry-Auth")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*PluginPullParamsBodyItems0
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {

			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qRemote, qhkRemote, _ := qs.GetOK("remote")
	if err := o.bindRemote(qRemote, qhkRemote, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRegistryAuth binds and validates parameter XRegistryAuth from header.
func (o *PluginPullParams) bindXRegistryAuth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XRegistryAuth = &raw

	return nil
}

// bindName binds and validates parameter Name from query.
func (o *PluginPullParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Name = &raw

	return nil
}

// bindRemote binds and validates parameter Remote from query.
func (o *PluginPullParams) bindRemote(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("remote", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("remote", "query", raw); err != nil {
		return err
	}
	o.Remote = raw

	return nil
}
