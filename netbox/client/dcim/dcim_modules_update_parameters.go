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

package dcim

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

// NewDcimModulesUpdateParams creates a new DcimModulesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModulesUpdateParams() *DcimModulesUpdateParams {
	return &DcimModulesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModulesUpdateParamsWithTimeout creates a new DcimModulesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModulesUpdateParamsWithTimeout(timeout time.Duration) *DcimModulesUpdateParams {
	return &DcimModulesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModulesUpdateParamsWithContext creates a new DcimModulesUpdateParams object
// with the ability to set a context for a request.
func NewDcimModulesUpdateParamsWithContext(ctx context.Context) *DcimModulesUpdateParams {
	return &DcimModulesUpdateParams{
		Context: ctx,
	}
}

// NewDcimModulesUpdateParamsWithHTTPClient creates a new DcimModulesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModulesUpdateParamsWithHTTPClient(client *http.Client) *DcimModulesUpdateParams {
	return &DcimModulesUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimModulesUpdateParams contains all the parameters to send to the API endpoint

	for the dcim modules update operation.

	Typically these are written to a http.Request.
*/
type DcimModulesUpdateParams struct {

	// Data.
	Data *models.WritableModule

	/* ID.

	   A unique integer value identifying this module.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim modules update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesUpdateParams) WithDefaults() *DcimModulesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim modules update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim modules update params
func (o *DcimModulesUpdateParams) WithTimeout(timeout time.Duration) *DcimModulesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim modules update params
func (o *DcimModulesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim modules update params
func (o *DcimModulesUpdateParams) WithContext(ctx context.Context) *DcimModulesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim modules update params
func (o *DcimModulesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim modules update params
func (o *DcimModulesUpdateParams) WithHTTPClient(client *http.Client) *DcimModulesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim modules update params
func (o *DcimModulesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim modules update params
func (o *DcimModulesUpdateParams) WithData(data *models.WritableModule) *DcimModulesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim modules update params
func (o *DcimModulesUpdateParams) SetData(data *models.WritableModule) {
	o.Data = data
}

// WithID adds the id to the dcim modules update params
func (o *DcimModulesUpdateParams) WithID(id int64) *DcimModulesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim modules update params
func (o *DcimModulesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModulesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
