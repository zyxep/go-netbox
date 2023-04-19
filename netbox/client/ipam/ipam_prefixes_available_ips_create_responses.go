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

// IpamPrefixesAvailableIpsCreateReader is a Reader for the IpamPrefixesAvailableIpsCreate structure.
type IpamPrefixesAvailableIpsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesAvailableIpsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesAvailableIpsCreateCreated creates a IpamPrefixesAvailableIpsCreateCreated with default headers values
func NewIpamPrefixesAvailableIpsCreateCreated() *IpamPrefixesAvailableIpsCreateCreated {
	return &IpamPrefixesAvailableIpsCreateCreated{}
}

/*
IpamPrefixesAvailableIpsCreateCreated describes a response with status code 201, with default header values.

IpamPrefixesAvailableIpsCreateCreated ipam prefixes available ips create created
*/
type IpamPrefixesAvailableIpsCreateCreated struct {
	Payload []*models.IPAddress
}

// IsSuccess returns true when this ipam prefixes available ips create created response has a 2xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes available ips create created response has a 3xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes available ips create created response has a 4xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes available ips create created response has a 5xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes available ips create created response a status code equal to that given
func (o *IpamPrefixesAvailableIpsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam prefixes available ips create created response
func (o *IpamPrefixesAvailableIpsCreateCreated) Code() int {
	return 201
}

func (o *IpamPrefixesAvailableIpsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamPrefixesAvailableIpsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamPrefixesAvailableIpsCreateCreated) GetPayload() []*models.IPAddress {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
