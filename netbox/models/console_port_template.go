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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsolePortTemplate console port template
//
// swagger:model ConsolePortTemplate
type ConsolePortTemplate struct {

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// type
	Type *ConsolePortTemplateType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this console port template
func (m *ConsolePortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *ConsolePortTemplate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this console port template based on the context it is used
func (m *ConsolePortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
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

func (m *ConsolePortTemplate) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceType != nil {
		if err := m.DeviceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePortTemplate) UnmarshalBinary(b []byte) error {
	var res ConsolePortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsolePortTemplateType Type
//
// swagger:model ConsolePortTemplateType
type ConsolePortTemplateType struct {

	// label
	// Required: true
	// Enum: [DE-9 DB-25 RJ-11 RJ-12 RJ-45 Mini-DIN 8 USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B USB Micro AB Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 mini-din-8 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b usb-micro-ab other]
	Value *string `json:"value"`
}

// Validate validates this console port template type
func (m *ConsolePortTemplateType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consolePortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE-9","DB-25","RJ-11","RJ-12","RJ-45","Mini-DIN 8","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","USB Micro AB","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortTemplateTypeTypeLabelPropEnum = append(consolePortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// ConsolePortTemplateTypeLabelDEDash9 captures enum value "DE-9"
	ConsolePortTemplateTypeLabelDEDash9 string = "DE-9"

	// ConsolePortTemplateTypeLabelDBDash25 captures enum value "DB-25"
	ConsolePortTemplateTypeLabelDBDash25 string = "DB-25"

	// ConsolePortTemplateTypeLabelRJDash11 captures enum value "RJ-11"
	ConsolePortTemplateTypeLabelRJDash11 string = "RJ-11"

	// ConsolePortTemplateTypeLabelRJDash12 captures enum value "RJ-12"
	ConsolePortTemplateTypeLabelRJDash12 string = "RJ-12"

	// ConsolePortTemplateTypeLabelRJDash45 captures enum value "RJ-45"
	ConsolePortTemplateTypeLabelRJDash45 string = "RJ-45"

	// ConsolePortTemplateTypeLabelMiniDashDIN8 captures enum value "Mini-DIN 8"
	ConsolePortTemplateTypeLabelMiniDashDIN8 string = "Mini-DIN 8"

	// ConsolePortTemplateTypeLabelUSBTypeA captures enum value "USB Type A"
	ConsolePortTemplateTypeLabelUSBTypeA string = "USB Type A"

	// ConsolePortTemplateTypeLabelUSBTypeB captures enum value "USB Type B"
	ConsolePortTemplateTypeLabelUSBTypeB string = "USB Type B"

	// ConsolePortTemplateTypeLabelUSBTypeC captures enum value "USB Type C"
	ConsolePortTemplateTypeLabelUSBTypeC string = "USB Type C"

	// ConsolePortTemplateTypeLabelUSBMiniA captures enum value "USB Mini A"
	ConsolePortTemplateTypeLabelUSBMiniA string = "USB Mini A"

	// ConsolePortTemplateTypeLabelUSBMiniB captures enum value "USB Mini B"
	ConsolePortTemplateTypeLabelUSBMiniB string = "USB Mini B"

	// ConsolePortTemplateTypeLabelUSBMicroA captures enum value "USB Micro A"
	ConsolePortTemplateTypeLabelUSBMicroA string = "USB Micro A"

	// ConsolePortTemplateTypeLabelUSBMicroB captures enum value "USB Micro B"
	ConsolePortTemplateTypeLabelUSBMicroB string = "USB Micro B"

	// ConsolePortTemplateTypeLabelUSBMicroAB captures enum value "USB Micro AB"
	ConsolePortTemplateTypeLabelUSBMicroAB string = "USB Micro AB"

	// ConsolePortTemplateTypeLabelOther captures enum value "Other"
	ConsolePortTemplateTypeLabelOther string = "Other"
)

// prop value enum
func (m *ConsolePortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consolePortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consolePortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","mini-din-8","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-micro-ab","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortTemplateTypeTypeValuePropEnum = append(consolePortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// ConsolePortTemplateTypeValueDeDash9 captures enum value "de-9"
	ConsolePortTemplateTypeValueDeDash9 string = "de-9"

	// ConsolePortTemplateTypeValueDbDash25 captures enum value "db-25"
	ConsolePortTemplateTypeValueDbDash25 string = "db-25"

	// ConsolePortTemplateTypeValueRjDash11 captures enum value "rj-11"
	ConsolePortTemplateTypeValueRjDash11 string = "rj-11"

	// ConsolePortTemplateTypeValueRjDash12 captures enum value "rj-12"
	ConsolePortTemplateTypeValueRjDash12 string = "rj-12"

	// ConsolePortTemplateTypeValueRjDash45 captures enum value "rj-45"
	ConsolePortTemplateTypeValueRjDash45 string = "rj-45"

	// ConsolePortTemplateTypeValueMiniDashDinDash8 captures enum value "mini-din-8"
	ConsolePortTemplateTypeValueMiniDashDinDash8 string = "mini-din-8"

	// ConsolePortTemplateTypeValueUsbDasha captures enum value "usb-a"
	ConsolePortTemplateTypeValueUsbDasha string = "usb-a"

	// ConsolePortTemplateTypeValueUsbDashb captures enum value "usb-b"
	ConsolePortTemplateTypeValueUsbDashb string = "usb-b"

	// ConsolePortTemplateTypeValueUsbDashc captures enum value "usb-c"
	ConsolePortTemplateTypeValueUsbDashc string = "usb-c"

	// ConsolePortTemplateTypeValueUsbDashMiniDasha captures enum value "usb-mini-a"
	ConsolePortTemplateTypeValueUsbDashMiniDasha string = "usb-mini-a"

	// ConsolePortTemplateTypeValueUsbDashMiniDashb captures enum value "usb-mini-b"
	ConsolePortTemplateTypeValueUsbDashMiniDashb string = "usb-mini-b"

	// ConsolePortTemplateTypeValueUsbDashMicroDasha captures enum value "usb-micro-a"
	ConsolePortTemplateTypeValueUsbDashMicroDasha string = "usb-micro-a"

	// ConsolePortTemplateTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	ConsolePortTemplateTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// ConsolePortTemplateTypeValueUsbDashMicroDashAb captures enum value "usb-micro-ab"
	ConsolePortTemplateTypeValueUsbDashMicroDashAb string = "usb-micro-ab"

	// ConsolePortTemplateTypeValueOther captures enum value "other"
	ConsolePortTemplateTypeValueOther string = "other"
)

// prop value enum
func (m *ConsolePortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consolePortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console port template type based on context it is used
func (m *ConsolePortTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePortTemplateType) UnmarshalBinary(b []byte) error {
	var res ConsolePortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
