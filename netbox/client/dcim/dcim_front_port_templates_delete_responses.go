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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimFrontPortTemplatesDeleteReader is a Reader for the DcimFrontPortTemplatesDelete structure.
type DcimFrontPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimFrontPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesDeleteNoContent creates a DcimFrontPortTemplatesDeleteNoContent with default headers values
func NewDcimFrontPortTemplatesDeleteNoContent() *DcimFrontPortTemplatesDeleteNoContent {
	return &DcimFrontPortTemplatesDeleteNoContent{}
}

/*
DcimFrontPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimFrontPortTemplatesDeleteNoContent dcim front port templates delete no content
*/
type DcimFrontPortTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim front port templates delete no content response has a 2xx status code
func (o *DcimFrontPortTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front port templates delete no content response has a 3xx status code
func (o *DcimFrontPortTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates delete no content response has a 4xx status code
func (o *DcimFrontPortTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front port templates delete no content response has a 5xx status code
func (o *DcimFrontPortTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates delete no content response a status code equal to that given
func (o *DcimFrontPortTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim front port templates delete no content response
func (o *DcimFrontPortTemplatesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimFrontPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimFrontPortTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimFrontPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
