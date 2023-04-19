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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// VirtualizationVirtualMachinesUpdateReader is a Reader for the VirtualizationVirtualMachinesUpdate structure.
type VirtualizationVirtualMachinesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesUpdateOK creates a VirtualizationVirtualMachinesUpdateOK with default headers values
func NewVirtualizationVirtualMachinesUpdateOK() *VirtualizationVirtualMachinesUpdateOK {
	return &VirtualizationVirtualMachinesUpdateOK{}
}

/*
VirtualizationVirtualMachinesUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesUpdateOK virtualization virtual machines update o k
*/
type VirtualizationVirtualMachinesUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

// IsSuccess returns true when this virtualization virtual machines update o k response has a 2xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual machines update o k response has a 3xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines update o k response has a 4xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual machines update o k response has a 5xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines update o k response a status code equal to that given
func (o *VirtualizationVirtualMachinesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization virtual machines update o k response
func (o *VirtualizationVirtualMachinesUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationVirtualMachinesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
