// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ImageBuildURL generates an URL for the image build operation
type ImageBuildURL struct {
	Buildargs   *string
	Cachefrom   *string
	Cpuperiod   *int64
	Cpuquota    *int64
	Cpusetcpus  *string
	Cpushares   *int64
	Dockerfile  *string
	Extrahosts  *string
	Forcerm     *bool
	Labels      *string
	Memory      *int64
	Memswap     *int64
	Networkmode *string
	Nocache     *bool
	Outputs     *string
	Platform    *string
	Pull        *string
	Q           *bool
	Remote      *string
	Rm          *bool
	Shmsize     *int64
	Squash      *bool
	T           *string
	Target      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ImageBuildURL) WithBasePath(bp string) *ImageBuildURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ImageBuildURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ImageBuildURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/build"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1.41"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var buildargsQ string
	if o.Buildargs != nil {
		buildargsQ = *o.Buildargs
	}
	if buildargsQ != "" {
		qs.Set("buildargs", buildargsQ)
	}

	var cachefromQ string
	if o.Cachefrom != nil {
		cachefromQ = *o.Cachefrom
	}
	if cachefromQ != "" {
		qs.Set("cachefrom", cachefromQ)
	}

	var cpuperiodQ string
	if o.Cpuperiod != nil {
		cpuperiodQ = swag.FormatInt64(*o.Cpuperiod)
	}
	if cpuperiodQ != "" {
		qs.Set("cpuperiod", cpuperiodQ)
	}

	var cpuquotaQ string
	if o.Cpuquota != nil {
		cpuquotaQ = swag.FormatInt64(*o.Cpuquota)
	}
	if cpuquotaQ != "" {
		qs.Set("cpuquota", cpuquotaQ)
	}

	var cpusetcpusQ string
	if o.Cpusetcpus != nil {
		cpusetcpusQ = *o.Cpusetcpus
	}
	if cpusetcpusQ != "" {
		qs.Set("cpusetcpus", cpusetcpusQ)
	}

	var cpusharesQ string
	if o.Cpushares != nil {
		cpusharesQ = swag.FormatInt64(*o.Cpushares)
	}
	if cpusharesQ != "" {
		qs.Set("cpushares", cpusharesQ)
	}

	var dockerfileQ string
	if o.Dockerfile != nil {
		dockerfileQ = *o.Dockerfile
	}
	if dockerfileQ != "" {
		qs.Set("dockerfile", dockerfileQ)
	}

	var extrahostsQ string
	if o.Extrahosts != nil {
		extrahostsQ = *o.Extrahosts
	}
	if extrahostsQ != "" {
		qs.Set("extrahosts", extrahostsQ)
	}

	var forcermQ string
	if o.Forcerm != nil {
		forcermQ = swag.FormatBool(*o.Forcerm)
	}
	if forcermQ != "" {
		qs.Set("forcerm", forcermQ)
	}

	var labelsQ string
	if o.Labels != nil {
		labelsQ = *o.Labels
	}
	if labelsQ != "" {
		qs.Set("labels", labelsQ)
	}

	var memoryQ string
	if o.Memory != nil {
		memoryQ = swag.FormatInt64(*o.Memory)
	}
	if memoryQ != "" {
		qs.Set("memory", memoryQ)
	}

	var memswapQ string
	if o.Memswap != nil {
		memswapQ = swag.FormatInt64(*o.Memswap)
	}
	if memswapQ != "" {
		qs.Set("memswap", memswapQ)
	}

	var networkmodeQ string
	if o.Networkmode != nil {
		networkmodeQ = *o.Networkmode
	}
	if networkmodeQ != "" {
		qs.Set("networkmode", networkmodeQ)
	}

	var nocacheQ string
	if o.Nocache != nil {
		nocacheQ = swag.FormatBool(*o.Nocache)
	}
	if nocacheQ != "" {
		qs.Set("nocache", nocacheQ)
	}

	var outputsQ string
	if o.Outputs != nil {
		outputsQ = *o.Outputs
	}
	if outputsQ != "" {
		qs.Set("outputs", outputsQ)
	}

	var platformQ string
	if o.Platform != nil {
		platformQ = *o.Platform
	}
	if platformQ != "" {
		qs.Set("platform", platformQ)
	}

	var pullQ string
	if o.Pull != nil {
		pullQ = *o.Pull
	}
	if pullQ != "" {
		qs.Set("pull", pullQ)
	}

	var qQ string
	if o.Q != nil {
		qQ = swag.FormatBool(*o.Q)
	}
	if qQ != "" {
		qs.Set("q", qQ)
	}

	var remoteQ string
	if o.Remote != nil {
		remoteQ = *o.Remote
	}
	if remoteQ != "" {
		qs.Set("remote", remoteQ)
	}

	var rmQ string
	if o.Rm != nil {
		rmQ = swag.FormatBool(*o.Rm)
	}
	if rmQ != "" {
		qs.Set("rm", rmQ)
	}

	var shmsizeQ string
	if o.Shmsize != nil {
		shmsizeQ = swag.FormatInt64(*o.Shmsize)
	}
	if shmsizeQ != "" {
		qs.Set("shmsize", shmsizeQ)
	}

	var squashQ string
	if o.Squash != nil {
		squashQ = swag.FormatBool(*o.Squash)
	}
	if squashQ != "" {
		qs.Set("squash", squashQ)
	}

	var tQ string
	if o.T != nil {
		tQ = *o.T
	}
	if tQ != "" {
		qs.Set("t", tQ)
	}

	var targetQ string
	if o.Target != nil {
		targetQ = *o.Target
	}
	if targetQ != "" {
		qs.Set("target", targetQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ImageBuildURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ImageBuildURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ImageBuildURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ImageBuildURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ImageBuildURL")
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
func (o *ImageBuildURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
