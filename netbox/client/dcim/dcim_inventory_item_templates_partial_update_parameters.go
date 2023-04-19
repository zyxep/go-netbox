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

// NewDcimInventoryItemTemplatesPartialUpdateParams creates a new DcimInventoryItemTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemTemplatesPartialUpdateParams() *DcimInventoryItemTemplatesPartialUpdateParams {
	return &DcimInventoryItemTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemTemplatesPartialUpdateParamsWithTimeout creates a new DcimInventoryItemTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemTemplatesPartialUpdateParams {
	return &DcimInventoryItemTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemTemplatesPartialUpdateParamsWithContext creates a new DcimInventoryItemTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimInventoryItemTemplatesPartialUpdateParams {
	return &DcimInventoryItemTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimInventoryItemTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemTemplatesPartialUpdateParams {
	return &DcimInventoryItemTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimInventoryItemTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim inventory item templates partial update operation.

	Typically these are written to a http.Request.
*/
type DcimInventoryItemTemplatesPartialUpdateParams struct {

	// Data.
	Data *models.WritableInventoryItemTemplate

	/* ID.

	   A unique integer value identifying this inventory item template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory item templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemTemplatesPartialUpdateParams) WithDefaults() *DcimInventoryItemTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory item templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimInventoryItemTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) WithData(data *models.WritableInventoryItemTemplate) *DcimInventoryItemTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) SetData(data *models.WritableInventoryItemTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) WithID(id int64) *DcimInventoryItemTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory item templates partial update params
func (o *DcimInventoryItemTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
