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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zyxep/go-netbox/v3/netbox/models"
)

// IpamIPRangesBulkUpdateReader is a Reader for the IpamIPRangesBulkUpdate structure.
type IpamIPRangesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamIPRangesBulkUpdateOK creates a IpamIPRangesBulkUpdateOK with default headers values
func NewIpamIPRangesBulkUpdateOK() *IpamIPRangesBulkUpdateOK {
	return &IpamIPRangesBulkUpdateOK{}
}

/*
IpamIPRangesBulkUpdateOK describes a response with status code 200, with default header values.

IpamIPRangesBulkUpdateOK ipam Ip ranges bulk update o k
*/
type IpamIPRangesBulkUpdateOK struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges bulk update o k response has a 2xx status code
func (o *IpamIPRangesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges bulk update o k response has a 3xx status code
func (o *IpamIPRangesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges bulk update o k response has a 4xx status code
func (o *IpamIPRangesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges bulk update o k response has a 5xx status code
func (o *IpamIPRangesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges bulk update o k response a status code equal to that given
func (o *IpamIPRangesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip ranges bulk update o k response
func (o *IpamIPRangesBulkUpdateOK) Code() int {
	return 200
}

func (o *IpamIPRangesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/][%d] ipamIpRangesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPRangesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/][%d] ipamIpRangesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPRangesBulkUpdateOK) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
