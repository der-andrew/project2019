// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// DocumentCollectionURL generates an URL for the document collection operation
type DocumentCollectionURL struct {
	Code   *string
	Limit  *int64
	Locale string
	Offset *int64
	Type   string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DocumentCollectionURL) WithBasePath(bp string) *DocumentCollectionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DocumentCollectionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DocumentCollectionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/documents"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var code string
	if o.Code != nil {
		code = *o.Code
	}
	if code != "" {
		qs.Set("code", code)
	}

	var limit string
	if o.Limit != nil {
		limit = swag.FormatInt64(*o.Limit)
	}
	if limit != "" {
		qs.Set("limit", limit)
	}

	locale := o.Locale
	if locale != "" {
		qs.Set("locale", locale)
	}

	var offset string
	if o.Offset != nil {
		offset = swag.FormatInt64(*o.Offset)
	}
	if offset != "" {
		qs.Set("offset", offset)
	}

	typeVar := o.Type
	if typeVar != "" {
		qs.Set("type", typeVar)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DocumentCollectionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DocumentCollectionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DocumentCollectionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DocumentCollectionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DocumentCollectionURL")
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
func (o *DocumentCollectionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
