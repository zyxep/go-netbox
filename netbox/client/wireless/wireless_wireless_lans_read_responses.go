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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zyxep/go-netbox/v3/netbox/models"
)

// WirelessWirelessLansReadReader is a Reader for the WirelessWirelessLansRead structure.
type WirelessWirelessLansReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLansReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWirelessWirelessLansReadOK creates a WirelessWirelessLansReadOK with default headers values
func NewWirelessWirelessLansReadOK() *WirelessWirelessLansReadOK {
	return &WirelessWirelessLansReadOK{}
}

/*
WirelessWirelessLansReadOK describes a response with status code 200, with default header values.

WirelessWirelessLansReadOK wireless wireless lans read o k
*/
type WirelessWirelessLansReadOK struct {
	Payload *models.WirelessLAN
}

// IsSuccess returns true when this wireless wireless lans read o k response has a 2xx status code
func (o *WirelessWirelessLansReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans read o k response has a 3xx status code
func (o *WirelessWirelessLansReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans read o k response has a 4xx status code
func (o *WirelessWirelessLansReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans read o k response has a 5xx status code
func (o *WirelessWirelessLansReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans read o k response a status code equal to that given
func (o *WirelessWirelessLansReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *WirelessWirelessLansReadOK) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansReadOK) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansReadOK) GetPayload() *models.WirelessLAN {
	return o.Payload
}

func (o *WirelessWirelessLansReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
