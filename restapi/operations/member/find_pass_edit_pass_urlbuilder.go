// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// FindPassEditPassURL generates an URL for the find pass edit pass operation
type FindPassEditPassURL struct {
	Client      *string
	ExtraInfo1  *string
	Imei        *string
	NewPass     *string
	OldPass     *string
	PhoneNumber *string
	Ts          *int64
	Type        *int64
	Val         *string
	Version     *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindPassEditPassURL) WithBasePath(bp string) *FindPassEditPassURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindPassEditPassURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindPassEditPassURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/member/findPass/editPass"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/joe9724/jinlinonline/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var client string
	if o.Client != nil {
		client = *o.Client
	}
	if client != "" {
		qs.Set("client", client)
	}

	var extraInfo1 string
	if o.ExtraInfo1 != nil {
		extraInfo1 = *o.ExtraInfo1
	}
	if extraInfo1 != "" {
		qs.Set("extraInfo1", extraInfo1)
	}

	var imei string
	if o.Imei != nil {
		imei = *o.Imei
	}
	if imei != "" {
		qs.Set("imei", imei)
	}

	var newPass string
	if o.NewPass != nil {
		newPass = *o.NewPass
	}
	if newPass != "" {
		qs.Set("newPass", newPass)
	}

	var oldPass string
	if o.OldPass != nil {
		oldPass = *o.OldPass
	}
	if oldPass != "" {
		qs.Set("oldPass", oldPass)
	}

	var phoneNumber string
	if o.PhoneNumber != nil {
		phoneNumber = *o.PhoneNumber
	}
	if phoneNumber != "" {
		qs.Set("phoneNumber", phoneNumber)
	}

	var ts string
	if o.Ts != nil {
		ts = swag.FormatInt64(*o.Ts)
	}
	if ts != "" {
		qs.Set("ts", ts)
	}

	var typeVar string
	if o.Type != nil {
		typeVar = swag.FormatInt64(*o.Type)
	}
	if typeVar != "" {
		qs.Set("type", typeVar)
	}

	var val string
	if o.Val != nil {
		val = *o.Val
	}
	if val != "" {
		qs.Set("val", val)
	}

	var version string
	if o.Version != nil {
		version = *o.Version
	}
	if version != "" {
		qs.Set("version", version)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindPassEditPassURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindPassEditPassURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindPassEditPassURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindPassEditPassURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindPassEditPassURL")
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
func (o *FindPassEditPassURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}