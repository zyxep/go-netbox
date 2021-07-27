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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCable writable cable
//
// swagger:model WritableCable
type WritableCable struct {

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	// Max Length: 100
	Label string `json:"label,omitempty"`

	// Length
	// Maximum: 32767
	// Minimum: 0
	Length *int64 `json:"length,omitempty"`

	// Length unit
	// Enum: [m cm ft in]
	LengthUnit string `json:"length_unit,omitempty"`

	// Status
	// Enum: [connected planned decommissioning]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Termination a
	// Read Only: true
	Terminationa map[string]*string `json:"termination_a,omitempty"`

	// Termination a id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationaID *int64 `json:"termination_a_id"`

	// Termination a type
	// Required: true
	TerminationaType *string `json:"termination_a_type"`

	// Termination b
	// Read Only: true
	Terminationb map[string]*string `json:"termination_b,omitempty"`

	// Termination b id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationbID *int64 `json:"termination_b_id"`

	// Termination b type
	// Required: true
	TerminationbType *string `json:"termination_b_type"`

	// Type
	// Enum: [cat3 cat5 cat5e cat6 cat6a cat7 cat7a cat8 dac-active dac-passive mrj21-trunk coaxial mmf mmf-om1 mmf-om2 mmf-om3 mmf-om4 mmf-om5 smf smf-os1 smf-os2 aoc power]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable cable
func (m *WritableCable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLengthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCable) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", m.Color, 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateLength(formats strfmt.Registry) error {
	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if err := validate.MinimumInt("length", "body", *m.Length, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "body", *m.Length, 32767, false); err != nil {
		return err
	}

	return nil
}

var writableCableTypeLengthUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["m","cm","ft","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeLengthUnitPropEnum = append(writableCableTypeLengthUnitPropEnum, v)
	}
}

const (

	// WritableCableLengthUnitM captures enum value "m"
	WritableCableLengthUnitM string = "m"

	// WritableCableLengthUnitCm captures enum value "cm"
	WritableCableLengthUnitCm string = "cm"

	// WritableCableLengthUnitFt captures enum value "ft"
	WritableCableLengthUnitFt string = "ft"

	// WritableCableLengthUnitIn captures enum value "in"
	WritableCableLengthUnitIn string = "in"
)

// prop value enum
func (m *WritableCable) validateLengthUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableCableTypeLengthUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateLengthUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.LengthUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateLengthUnitEnum("length_unit", "body", m.LengthUnit); err != nil {
		return err
	}

	return nil
}

var writableCableTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","planned","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeStatusPropEnum = append(writableCableTypeStatusPropEnum, v)
	}
}

const (

	// WritableCableStatusConnected captures enum value "connected"
	WritableCableStatusConnected string = "connected"

	// WritableCableStatusPlanned captures enum value "planned"
	WritableCableStatusPlanned string = "planned"

	// WritableCableStatusDecommissioning captures enum value "decommissioning"
	WritableCableStatusDecommissioning string = "decommissioning"
)

// prop value enum
func (m *WritableCable) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableCableTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableCable) validateTerminationaID(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_id", "body", m.TerminationaID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_a_id", "body", *m.TerminationaID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_a_id", "body", *m.TerminationaID, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationaType(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_type", "body", m.TerminationaType); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationbID(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_id", "body", m.TerminationbID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_b_id", "body", *m.TerminationbID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_b_id", "body", *m.TerminationbID, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationbType(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_type", "body", m.TerminationbType); err != nil {
		return err
	}

	return nil
}

var writableCableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cat3","cat5","cat5e","cat6","cat6a","cat7","cat7a","cat8","dac-active","dac-passive","mrj21-trunk","coaxial","mmf","mmf-om1","mmf-om2","mmf-om3","mmf-om4","mmf-om5","smf","smf-os1","smf-os2","aoc","power"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeTypePropEnum = append(writableCableTypeTypePropEnum, v)
	}
}

const (

	// WritableCableTypeCat3 captures enum value "cat3"
	WritableCableTypeCat3 string = "cat3"

	// WritableCableTypeCat5 captures enum value "cat5"
	WritableCableTypeCat5 string = "cat5"

	// WritableCableTypeCat5e captures enum value "cat5e"
	WritableCableTypeCat5e string = "cat5e"

	// WritableCableTypeCat6 captures enum value "cat6"
	WritableCableTypeCat6 string = "cat6"

	// WritableCableTypeCat6a captures enum value "cat6a"
	WritableCableTypeCat6a string = "cat6a"

	// WritableCableTypeCat7 captures enum value "cat7"
	WritableCableTypeCat7 string = "cat7"

	// WritableCableTypeCat7a captures enum value "cat7a"
	WritableCableTypeCat7a string = "cat7a"

	// WritableCableTypeCat8 captures enum value "cat8"
	WritableCableTypeCat8 string = "cat8"

	// WritableCableTypeDacDashActive captures enum value "dac-active"
	WritableCableTypeDacDashActive string = "dac-active"

	// WritableCableTypeDacDashPassive captures enum value "dac-passive"
	WritableCableTypeDacDashPassive string = "dac-passive"

	// WritableCableTypeMrj21DashTrunk captures enum value "mrj21-trunk"
	WritableCableTypeMrj21DashTrunk string = "mrj21-trunk"

	// WritableCableTypeCoaxial captures enum value "coaxial"
	WritableCableTypeCoaxial string = "coaxial"

	// WritableCableTypeMmf captures enum value "mmf"
	WritableCableTypeMmf string = "mmf"

	// WritableCableTypeMmfDashOm1 captures enum value "mmf-om1"
	WritableCableTypeMmfDashOm1 string = "mmf-om1"

	// WritableCableTypeMmfDashOm2 captures enum value "mmf-om2"
	WritableCableTypeMmfDashOm2 string = "mmf-om2"

	// WritableCableTypeMmfDashOm3 captures enum value "mmf-om3"
	WritableCableTypeMmfDashOm3 string = "mmf-om3"

	// WritableCableTypeMmfDashOm4 captures enum value "mmf-om4"
	WritableCableTypeMmfDashOm4 string = "mmf-om4"

	// WritableCableTypeMmfDashOm5 captures enum value "mmf-om5"
	WritableCableTypeMmfDashOm5 string = "mmf-om5"

	// WritableCableTypeSmf captures enum value "smf"
	WritableCableTypeSmf string = "smf"

	// WritableCableTypeSmfDashOs1 captures enum value "smf-os1"
	WritableCableTypeSmfDashOs1 string = "smf-os1"

	// WritableCableTypeSmfDashOs2 captures enum value "smf-os2"
	WritableCableTypeSmfDashOs2 string = "smf-os2"

	// WritableCableTypeAoc captures enum value "aoc"
	WritableCableTypeAoc string = "aoc"

	// WritableCableTypePower captures enum value "power"
	WritableCableTypePower string = "power"
)

// prop value enum
func (m *WritableCable) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableCableTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable cable based on the context it is used
func (m *WritableCable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerminationa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerminationb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCable) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableCable) contextValidateTerminationa(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableCable) contextValidateTerminationb(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableCable) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCable) UnmarshalBinary(b []byte) error {
	var res WritableCable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
