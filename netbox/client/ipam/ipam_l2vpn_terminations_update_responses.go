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

// IpamL2vpnTerminationsUpdateReader is a Reader for the IpamL2vpnTerminationsUpdate structure.
type IpamL2vpnTerminationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnTerminationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnTerminationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamL2vpnTerminationsUpdateOK creates a IpamL2vpnTerminationsUpdateOK with default headers values
func NewIpamL2vpnTerminationsUpdateOK() *IpamL2vpnTerminationsUpdateOK {
	return &IpamL2vpnTerminationsUpdateOK{}
}

/*
IpamL2vpnTerminationsUpdateOK describes a response with status code 200, with default header values.

IpamL2vpnTerminationsUpdateOK ipam l2vpn terminations update o k
*/
type IpamL2vpnTerminationsUpdateOK struct {
	Payload *models.L2VPNTermination
}

// IsSuccess returns true when this ipam l2vpn terminations update o k response has a 2xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpn terminations update o k response has a 3xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpn terminations update o k response has a 4xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpn terminations update o k response has a 5xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpn terminations update o k response a status code equal to that given
func (o *IpamL2vpnTerminationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamL2vpnTerminationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/l2vpn-terminations/{id}/][%d] ipamL2vpnTerminationsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/l2vpn-terminations/{id}/][%d] ipamL2vpnTerminationsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsUpdateOK) GetPayload() *models.L2VPNTermination {
	return o.Payload
}

func (o *IpamL2vpnTerminationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPNTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
