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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zyxep/go-netbox/v3/netbox/models"
)

// DcimPowerPortTemplatesReadReader is a Reader for the DcimPowerPortTemplatesRead structure.
type DcimPowerPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesReadOK creates a DcimPowerPortTemplatesReadOK with default headers values
func NewDcimPowerPortTemplatesReadOK() *DcimPowerPortTemplatesReadOK {
	return &DcimPowerPortTemplatesReadOK{}
}

/*
DcimPowerPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesReadOK dcim power port templates read o k
*/
type DcimPowerPortTemplatesReadOK struct {
	Payload *models.PowerPortTemplate
}

// IsSuccess returns true when this dcim power port templates read o k response has a 2xx status code
func (o *DcimPowerPortTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates read o k response has a 3xx status code
func (o *DcimPowerPortTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates read o k response has a 4xx status code
func (o *DcimPowerPortTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates read o k response has a 5xx status code
func (o *DcimPowerPortTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates read o k response a status code equal to that given
func (o *DcimPowerPortTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power port templates read o k response
func (o *DcimPowerPortTemplatesReadOK) Code() int {
	return 200
}

func (o *DcimPowerPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesReadOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
