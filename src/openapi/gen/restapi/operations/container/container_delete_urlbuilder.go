// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// ContainerDeleteURL generates an URL for the container delete operation
type ContainerDeleteURL struct {
	ID string

	Force *bool
	Link  *bool
	V     *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ContainerDeleteURL) WithBasePath(bp string) *ContainerDeleteURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ContainerDeleteURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ContainerDeleteURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/containers/{id}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on ContainerDeleteURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1.41"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var forceQ string
	if o.Force != nil {
		forceQ = swag.FormatBool(*o.Force)
	}
	if forceQ != "" {
		qs.Set("force", forceQ)
	}

	var linkQ string
	if o.Link != nil {
		linkQ = swag.FormatBool(*o.Link)
	}
	if linkQ != "" {
		qs.Set("link", linkQ)
	}

	var vQ string
	if o.V != nil {
		vQ = swag.FormatBool(*o.V)
	}
	if vQ != "" {
		qs.Set("v", vQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ContainerDeleteURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ContainerDeleteURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ContainerDeleteURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ContainerDeleteURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ContainerDeleteURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ContainerDeleteURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}