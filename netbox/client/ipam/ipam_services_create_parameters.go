// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/zyxep/go-netbox/v3/netbox/models"
)

// NewIpamServicesCreateParams creates a new IpamServicesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesCreateParams() *IpamServicesCreateParams {
	return &IpamServicesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesCreateParamsWithTimeout creates a new IpamServicesCreateParams object
// with the ability to set a timeout on a request.
func NewIpamServicesCreateParamsWithTimeout(timeout time.Duration) *IpamServicesCreateParams {
	return &IpamServicesCreateParams{
		timeout: timeout,
	}
}

// NewIpamServicesCreateParamsWithContext creates a new IpamServicesCreateParams object
// with the ability to set a context for a request.
func NewIpamServicesCreateParamsWithContext(ctx context.Context) *IpamServicesCreateParams {
	return &IpamServicesCreateParams{
		Context: ctx,
	}
}

// NewIpamServicesCreateParamsWithHTTPClient creates a new IpamServicesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesCreateParamsWithHTTPClient(client *http.Client) *IpamServicesCreateParams {
	return &IpamServicesCreateParams{
		HTTPClient: client,
	}
}

/*
IpamServicesCreateParams contains all the parameters to send to the API endpoint

	for the ipam services create operation.

	Typically these are written to a http.Request.
*/
type IpamServicesCreateParams struct {

	// Data.
	Data *models.WritableService

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesCreateParams) WithDefaults() *IpamServicesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services create params
func (o *IpamServicesCreateParams) WithTimeout(timeout time.Duration) *IpamServicesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services create params
func (o *IpamServicesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services create params
func (o *IpamServicesCreateParams) WithContext(ctx context.Context) *IpamServicesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services create params
func (o *IpamServicesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services create params
func (o *IpamServicesCreateParams) WithHTTPClient(client *http.Client) *IpamServicesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services create params
func (o *IpamServicesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services create params
func (o *IpamServicesCreateParams) WithData(data *models.WritableService) *IpamServicesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services create params
func (o *IpamServicesCreateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
