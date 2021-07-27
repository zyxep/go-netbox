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
	"github.com/go-openapi/swag"
)

// NewIpamRirsReadParams creates a new IpamRirsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRirsReadParams() *IpamRirsReadParams {
	return &IpamRirsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsReadParamsWithTimeout creates a new IpamRirsReadParams object
// with the ability to set a timeout on a request.
func NewIpamRirsReadParamsWithTimeout(timeout time.Duration) *IpamRirsReadParams {
	return &IpamRirsReadParams{
		timeout: timeout,
	}
}

// NewIpamRirsReadParamsWithContext creates a new IpamRirsReadParams object
// with the ability to set a context for a request.
func NewIpamRirsReadParamsWithContext(ctx context.Context) *IpamRirsReadParams {
	return &IpamRirsReadParams{
		Context: ctx,
	}
}

// NewIpamRirsReadParamsWithHTTPClient creates a new IpamRirsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRirsReadParamsWithHTTPClient(client *http.Client) *IpamRirsReadParams {
	return &IpamRirsReadParams{
		HTTPClient: client,
	}
}

/* IpamRirsReadParams contains all the parameters to send to the API endpoint
   for the ipam rirs read operation.

   Typically these are written to a http.Request.
*/
type IpamRirsReadParams struct {

	/* ID.

	   A unique integer value identifying this RIR.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam rirs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsReadParams) WithDefaults() *IpamRirsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam rirs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam rirs read params
func (o *IpamRirsReadParams) WithTimeout(timeout time.Duration) *IpamRirsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs read params
func (o *IpamRirsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs read params
func (o *IpamRirsReadParams) WithContext(ctx context.Context) *IpamRirsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs read params
func (o *IpamRirsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs read params
func (o *IpamRirsReadParams) WithHTTPClient(client *http.Client) *IpamRirsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs read params
func (o *IpamRirsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam rirs read params
func (o *IpamRirsReadParams) WithID(id int64) *IpamRirsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam rirs read params
func (o *IpamRirsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
