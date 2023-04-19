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
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// DcimCableTerminationsListReader is a Reader for the DcimCableTerminationsList structure.
type DcimCableTerminationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCableTerminationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimCableTerminationsListOK creates a DcimCableTerminationsListOK with default headers values
func NewDcimCableTerminationsListOK() *DcimCableTerminationsListOK {
	return &DcimCableTerminationsListOK{}
}

/*
DcimCableTerminationsListOK describes a response with status code 200, with default header values.

DcimCableTerminationsListOK dcim cable terminations list o k
*/
type DcimCableTerminationsListOK struct {
	Payload *DcimCableTerminationsListOKBody
}

// IsSuccess returns true when this dcim cable terminations list o k response has a 2xx status code
func (o *DcimCableTerminationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations list o k response has a 3xx status code
func (o *DcimCableTerminationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations list o k response has a 4xx status code
func (o *DcimCableTerminationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations list o k response has a 5xx status code
func (o *DcimCableTerminationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations list o k response a status code equal to that given
func (o *DcimCableTerminationsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim cable terminations list o k response
func (o *DcimCableTerminationsListOK) Code() int {
	return 200
}

func (o *DcimCableTerminationsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/cable-terminations/][%d] dcimCableTerminationsListOK  %+v", 200, o.Payload)
}

func (o *DcimCableTerminationsListOK) String() string {
	return fmt.Sprintf("[GET /dcim/cable-terminations/][%d] dcimCableTerminationsListOK  %+v", 200, o.Payload)
}

func (o *DcimCableTerminationsListOK) GetPayload() *DcimCableTerminationsListOKBody {
	return o.Payload
}

func (o *DcimCableTerminationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimCableTerminationsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimCableTerminationsListOKBody dcim cable terminations list o k body
swagger:model DcimCableTerminationsListOKBody
*/
type DcimCableTerminationsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.CableTermination `json:"results"`
}

// Validate validates this dcim cable terminations list o k body
func (o *DcimCableTerminationsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimCableTerminationsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimCableTerminationsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimCableTerminationsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCableTerminationsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCableTerminationsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCableTerminationsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCableTerminationsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimCableTerminationsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim cable terminations list o k body based on the context it is used
func (o *DcimCableTerminationsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimCableTerminationsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimCableTerminationsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimCableTerminationsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimCableTerminationsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
