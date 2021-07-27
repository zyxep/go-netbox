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

package tenancy

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

// NewTenancyTenantsListParams creates a new TenancyTenantsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsListParams() *TenancyTenantsListParams {
	return &TenancyTenantsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsListParamsWithTimeout creates a new TenancyTenantsListParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsListParamsWithTimeout(timeout time.Duration) *TenancyTenantsListParams {
	return &TenancyTenantsListParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsListParamsWithContext creates a new TenancyTenantsListParams object
// with the ability to set a context for a request.
func NewTenancyTenantsListParamsWithContext(ctx context.Context) *TenancyTenantsListParams {
	return &TenancyTenantsListParams{
		Context: ctx,
	}
}

// NewTenancyTenantsListParamsWithHTTPClient creates a new TenancyTenantsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsListParamsWithHTTPClient(client *http.Client) *TenancyTenantsListParams {
	return &TenancyTenantsListParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsListParams contains all the parameters to send to the API endpoint
   for the tenancy tenants list operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Group.
	Group *string

	// Groupn.
	Groupn *string

	// GroupID.
	GroupID *string

	// GroupIDn.
	GroupIDn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Slug.
	Slug *string

	// SlugEmpty.
	SlugEmpty *string

	// SlugIc.
	SlugIc *string

	// SlugIe.
	SlugIe *string

	// SlugIew.
	SlugIew *string

	// SlugIsw.
	SlugIsw *string

	// Slugn.
	Slugn *string

	// SlugNic.
	SlugNic *string

	// SlugNie.
	SlugNie *string

	// SlugNiew.
	SlugNiew *string

	// SlugNisw.
	SlugNisw *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsListParams) WithDefaults() *TenancyTenantsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithTimeout(timeout time.Duration) *TenancyTenantsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithContext(ctx context.Context) *TenancyTenantsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithHTTPClient(client *http.Client) *TenancyTenantsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithCreated(created *string) *TenancyTenantsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithCreatedGte(createdGte *string) *TenancyTenantsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithCreatedLte(createdLte *string) *TenancyTenantsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithGroup adds the group to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithGroup(group *string) *TenancyTenantsListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupn adds the groupn to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithGroupn(groupn *string) *TenancyTenantsListParams {
	o.SetGroupn(groupn)
	return o
}

// SetGroupn adds the groupN to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetGroupn(groupn *string) {
	o.Groupn = groupn
}

// WithGroupID adds the groupID to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithGroupID(groupID *string) *TenancyTenantsListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithGroupIDn adds the groupIDn to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithGroupIDn(groupIDn *string) *TenancyTenantsListParams {
	o.SetGroupIDn(groupIDn)
	return o
}

// SetGroupIDn adds the groupIdN to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetGroupIDn(groupIDn *string) {
	o.GroupIDn = groupIDn
}

// WithID adds the id to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithID(id *string) *TenancyTenantsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithIDGt(iDGt *string) *TenancyTenantsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithIDGte(iDGte *string) *TenancyTenantsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithIDLt(iDLt *string) *TenancyTenantsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithIDLte(iDLte *string) *TenancyTenantsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithIDn(iDn *string) *TenancyTenantsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithLastUpdated(lastUpdated *string) *TenancyTenantsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *TenancyTenantsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *TenancyTenantsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithLimit(limit *int64) *TenancyTenantsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithName(name *string) *TenancyTenantsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameEmpty(nameEmpty *string) *TenancyTenantsListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameIc(nameIc *string) *TenancyTenantsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameIe(nameIe *string) *TenancyTenantsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameIew(nameIew *string) *TenancyTenantsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameIsw(nameIsw *string) *TenancyTenantsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNamen(namen *string) *TenancyTenantsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameNic(nameNic *string) *TenancyTenantsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameNie(nameNie *string) *TenancyTenantsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameNiew(nameNiew *string) *TenancyTenantsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithNameNisw(nameNisw *string) *TenancyTenantsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithOffset(offset *int64) *TenancyTenantsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithQ(q *string) *TenancyTenantsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlug(slug *string) *TenancyTenantsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugEmpty adds the slugEmpty to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugEmpty(slugEmpty *string) *TenancyTenantsListParams {
	o.SetSlugEmpty(slugEmpty)
	return o
}

// SetSlugEmpty adds the slugEmpty to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugEmpty(slugEmpty *string) {
	o.SlugEmpty = slugEmpty
}

// WithSlugIc adds the slugIc to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugIc(slugIc *string) *TenancyTenantsListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugIe(slugIe *string) *TenancyTenantsListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugIew(slugIew *string) *TenancyTenantsListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugIsw(slugIsw *string) *TenancyTenantsListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugn(slugn *string) *TenancyTenantsListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugNic(slugNic *string) *TenancyTenantsListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugNie(slugNie *string) *TenancyTenantsListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugNiew(slugNiew *string) *TenancyTenantsListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithSlugNisw(slugNisw *string) *TenancyTenantsListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WithTag adds the tag to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithTag(tag *string) *TenancyTenantsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the tenancy tenants list params
func (o *TenancyTenantsListParams) WithTagn(tagn *string) *TenancyTenantsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the tenancy tenants list params
func (o *TenancyTenantsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.Groupn != nil {

		// query param group__n
		var qrGroupn string

		if o.Groupn != nil {
			qrGroupn = *o.Groupn
		}
		qGroupn := qrGroupn
		if qGroupn != "" {

			if err := r.SetQueryParam("group__n", qGroupn); err != nil {
				return err
			}
		}
	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string

		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}
	}

	if o.GroupIDn != nil {

		// query param group_id__n
		var qrGroupIDn string

		if o.GroupIDn != nil {
			qrGroupIDn = *o.GroupIDn
		}
		qGroupIDn := qrGroupIDn
		if qGroupIDn != "" {

			if err := r.SetQueryParam("group_id__n", qGroupIDn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string

		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {

			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}
	}

	if o.SlugEmpty != nil {

		// query param slug__empty
		var qrSlugEmpty string

		if o.SlugEmpty != nil {
			qrSlugEmpty = *o.SlugEmpty
		}
		qSlugEmpty := qrSlugEmpty
		if qSlugEmpty != "" {

			if err := r.SetQueryParam("slug__empty", qSlugEmpty); err != nil {
				return err
			}
		}
	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string

		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {

			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}
	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string

		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {

			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}
	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string

		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {

			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}
	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string

		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {

			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}
	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string

		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {

			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}
	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string

		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {

			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}
	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string

		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {

			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}
	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string

		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {

			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}
	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string

		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {

			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
