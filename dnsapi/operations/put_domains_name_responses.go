// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/pir5/pir5-go/dnsapi/model"
)

// PutDomainsNameReader is a Reader for the PutDomainsName structure.
type PutDomainsNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDomainsNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDomainsNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPutDomainsNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDomainsNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDomainsNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDomainsNameOK creates a PutDomainsNameOK with default headers values
func NewPutDomainsNameOK() *PutDomainsNameOK {
	return &PutDomainsNameOK{}
}

/*PutDomainsNameOK handles this case with default header values.

OK
*/
type PutDomainsNameOK struct {
	Payload *model.ModelDomain
}

func (o *PutDomainsNameOK) Error() string {
	return fmt.Sprintf("[PUT /domains/{name}][%d] putDomainsNameOK  %+v", 200, o.Payload)
}

func (o *PutDomainsNameOK) GetPayload() *model.ModelDomain {
	return o.Payload
}

func (o *PutDomainsNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ModelDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDomainsNameForbidden creates a PutDomainsNameForbidden with default headers values
func NewPutDomainsNameForbidden() *PutDomainsNameForbidden {
	return &PutDomainsNameForbidden{}
}

/*PutDomainsNameForbidden handles this case with default header values.

Forbidden
*/
type PutDomainsNameForbidden struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PutDomainsNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /domains/{name}][%d] putDomainsNameForbidden  %+v", 403, o.Payload)
}

func (o *PutDomainsNameForbidden) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutDomainsNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDomainsNameNotFound creates a PutDomainsNameNotFound with default headers values
func NewPutDomainsNameNotFound() *PutDomainsNameNotFound {
	return &PutDomainsNameNotFound{}
}

/*PutDomainsNameNotFound handles this case with default header values.

Not Found
*/
type PutDomainsNameNotFound struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PutDomainsNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /domains/{name}][%d] putDomainsNameNotFound  %+v", 404, o.Payload)
}

func (o *PutDomainsNameNotFound) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutDomainsNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDomainsNameInternalServerError creates a PutDomainsNameInternalServerError with default headers values
func NewPutDomainsNameInternalServerError() *PutDomainsNameInternalServerError {
	return &PutDomainsNameInternalServerError{}
}

/*PutDomainsNameInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutDomainsNameInternalServerError struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PutDomainsNameInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /domains/{name}][%d] putDomainsNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PutDomainsNameInternalServerError) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutDomainsNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
