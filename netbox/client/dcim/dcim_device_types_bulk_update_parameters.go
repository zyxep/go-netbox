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

// NewDcimDeviceTypesBulkUpdateParams creates a new DcimDeviceTypesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceTypesBulkUpdateParams() *DcimDeviceTypesBulkUpdateParams {
	return &DcimDeviceTypesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesBulkUpdateParamsWithTimeout creates a new DcimDeviceTypesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceTypesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesBulkUpdateParams {
	return &DcimDeviceTypesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceTypesBulkUpdateParamsWithContext creates a new DcimDeviceTypesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceTypesBulkUpdateParamsWithContext(ctx context.Context) *DcimDeviceTypesBulkUpdateParams {
	return &DcimDeviceTypesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceTypesBulkUpdateParamsWithHTTPClient creates a new DcimDeviceTypesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceTypesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesBulkUpdateParams {
	return &DcimDeviceTypesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimDeviceTypesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim device types bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimDeviceTypesBulkUpdateParams struct {

	// Data.
	Data *models.WritableDeviceType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device types bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesBulkUpdateParams) WithDefaults() *DcimDeviceTypesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device types bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) WithContext(ctx context.Context) *DcimDeviceTypesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) WithData(data *models.WritableDeviceType) *DcimDeviceTypesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device types bulk update params
func (o *DcimDeviceTypesBulkUpdateParams) SetData(data *models.WritableDeviceType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
