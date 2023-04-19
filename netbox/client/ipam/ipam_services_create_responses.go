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

// IpamServicesCreateReader is a Reader for the IpamServicesCreate structure.
type IpamServicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamServicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamServicesCreateCreated creates a IpamServicesCreateCreated with default headers values
func NewIpamServicesCreateCreated() *IpamServicesCreateCreated {
	return &IpamServicesCreateCreated{}
}

/*
IpamServicesCreateCreated describes a response with status code 201, with default header values.

IpamServicesCreateCreated ipam services create created
*/
type IpamServicesCreateCreated struct {
	Payload *models.Service
}

// IsSuccess returns true when this ipam services create created response has a 2xx status code
func (o *IpamServicesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam services create created response has a 3xx status code
func (o *IpamServicesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services create created response has a 4xx status code
func (o *IpamServicesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam services create created response has a 5xx status code
func (o *IpamServicesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services create created response a status code equal to that given
func (o *IpamServicesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam services create created response
func (o *IpamServicesCreateCreated) Code() int {
	return 201
}

func (o *IpamServicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipamServicesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamServicesCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipamServicesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamServicesCreateCreated) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
