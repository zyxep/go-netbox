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

	"github.com/zyxep/go-netbox/v3/netbox/models"
)

// NewDcimCableTerminationsBulkUpdateParams creates a new DcimCableTerminationsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCableTerminationsBulkUpdateParams() *DcimCableTerminationsBulkUpdateParams {
	return &DcimCableTerminationsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCableTerminationsBulkUpdateParamsWithTimeout creates a new DcimCableTerminationsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimCableTerminationsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimCableTerminationsBulkUpdateParams {
	return &DcimCableTerminationsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimCableTerminationsBulkUpdateParamsWithContext creates a new DcimCableTerminationsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimCableTerminationsBulkUpdateParamsWithContext(ctx context.Context) *DcimCableTerminationsBulkUpdateParams {
	return &DcimCableTerminationsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimCableTerminationsBulkUpdateParamsWithHTTPClient creates a new DcimCableTerminationsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCableTerminationsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimCableTerminationsBulkUpdateParams {
	return &DcimCableTerminationsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimCableTerminationsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim cable terminations bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimCableTerminationsBulkUpdateParams struct {

	// Data.
	Data *models.CableTermination

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cable terminations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCableTerminationsBulkUpdateParams) WithDefaults() *DcimCableTerminationsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cable terminations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCableTerminationsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimCableTerminationsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) WithContext(ctx context.Context) *DcimCableTerminationsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimCableTerminationsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) WithData(data *models.CableTermination) *DcimCableTerminationsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim cable terminations bulk update params
func (o *DcimCableTerminationsBulkUpdateParams) SetData(data *models.CableTermination) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCableTerminationsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
