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

package secrets

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

	"github.com/smutel/go-netbox/netbox/models"
)

// NewSecretsSecretsUpdateParams creates a new SecretsSecretsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretsSecretsUpdateParams() *SecretsSecretsUpdateParams {
	return &SecretsSecretsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretsUpdateParamsWithTimeout creates a new SecretsSecretsUpdateParams object
// with the ability to set a timeout on a request.
func NewSecretsSecretsUpdateParamsWithTimeout(timeout time.Duration) *SecretsSecretsUpdateParams {
	return &SecretsSecretsUpdateParams{
		timeout: timeout,
	}
}

// NewSecretsSecretsUpdateParamsWithContext creates a new SecretsSecretsUpdateParams object
// with the ability to set a context for a request.
func NewSecretsSecretsUpdateParamsWithContext(ctx context.Context) *SecretsSecretsUpdateParams {
	return &SecretsSecretsUpdateParams{
		Context: ctx,
	}
}

// NewSecretsSecretsUpdateParamsWithHTTPClient creates a new SecretsSecretsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretsSecretsUpdateParamsWithHTTPClient(client *http.Client) *SecretsSecretsUpdateParams {
	return &SecretsSecretsUpdateParams{
		HTTPClient: client,
	}
}

/* SecretsSecretsUpdateParams contains all the parameters to send to the API endpoint
   for the secrets secrets update operation.

   Typically these are written to a http.Request.
*/
type SecretsSecretsUpdateParams struct {

	// Data.
	Data *models.WritableSecret

	/* ID.

	   A unique integer value identifying this secret.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secrets secrets update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsSecretsUpdateParams) WithDefaults() *SecretsSecretsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secrets secrets update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsSecretsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithTimeout(timeout time.Duration) *SecretsSecretsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithContext(ctx context.Context) *SecretsSecretsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithHTTPClient(client *http.Client) *SecretsSecretsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithData(data *models.WritableSecret) *SecretsSecretsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetData(data *models.WritableSecret) {
	o.Data = data
}

// WithID adds the id to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithID(id int64) *SecretsSecretsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
