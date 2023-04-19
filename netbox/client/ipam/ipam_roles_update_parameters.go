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

	"github.com/zyxep/go-netbox/v3/netbox/models"
)

// NewIpamRolesUpdateParams creates a new IpamRolesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRolesUpdateParams() *IpamRolesUpdateParams {
	return &IpamRolesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRolesUpdateParamsWithTimeout creates a new IpamRolesUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRolesUpdateParamsWithTimeout(timeout time.Duration) *IpamRolesUpdateParams {
	return &IpamRolesUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRolesUpdateParamsWithContext creates a new IpamRolesUpdateParams object
// with the ability to set a context for a request.
func NewIpamRolesUpdateParamsWithContext(ctx context.Context) *IpamRolesUpdateParams {
	return &IpamRolesUpdateParams{
		Context: ctx,
	}
}

// NewIpamRolesUpdateParamsWithHTTPClient creates a new IpamRolesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRolesUpdateParamsWithHTTPClient(client *http.Client) *IpamRolesUpdateParams {
	return &IpamRolesUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamRolesUpdateParams contains all the parameters to send to the API endpoint

	for the ipam roles update operation.

	Typically these are written to a http.Request.
*/
type IpamRolesUpdateParams struct {

	// Data.
	Data *models.Role

	/* ID.

	   A unique integer value identifying this role.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesUpdateParams) WithDefaults() *IpamRolesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam roles update params
func (o *IpamRolesUpdateParams) WithTimeout(timeout time.Duration) *IpamRolesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles update params
func (o *IpamRolesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles update params
func (o *IpamRolesUpdateParams) WithContext(ctx context.Context) *IpamRolesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles update params
func (o *IpamRolesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles update params
func (o *IpamRolesUpdateParams) WithHTTPClient(client *http.Client) *IpamRolesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles update params
func (o *IpamRolesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam roles update params
func (o *IpamRolesUpdateParams) WithData(data *models.Role) *IpamRolesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam roles update params
func (o *IpamRolesUpdateParams) SetData(data *models.Role) {
	o.Data = data
}

// WithID adds the id to the ipam roles update params
func (o *IpamRolesUpdateParams) WithID(id int64) *IpamRolesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam roles update params
func (o *IpamRolesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRolesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
