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

// PutRecordsEnableIDReader is a Reader for the PutRecordsEnableID structure.
type PutRecordsEnableIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRecordsEnableIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRecordsEnableIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPutRecordsEnableIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRecordsEnableIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRecordsEnableIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutRecordsEnableIDOK creates a PutRecordsEnableIDOK with default headers values
func NewPutRecordsEnableIDOK() *PutRecordsEnableIDOK {
	return &PutRecordsEnableIDOK{}
}

/*PutRecordsEnableIDOK handles this case with default header values.

OK
*/
type PutRecordsEnableIDOK struct {
	Payload *model.ModelRecord
}

func (o *PutRecordsEnableIDOK) Error() string {
	return fmt.Sprintf("[PUT /records/enable/{id}][%d] putRecordsEnableIdOK  %+v", 200, o.Payload)
}

func (o *PutRecordsEnableIDOK) GetPayload() *model.ModelRecord {
	return o.Payload
}

func (o *PutRecordsEnableIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ModelRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordsEnableIDForbidden creates a PutRecordsEnableIDForbidden with default headers values
func NewPutRecordsEnableIDForbidden() *PutRecordsEnableIDForbidden {
	return &PutRecordsEnableIDForbidden{}
}

/*PutRecordsEnableIDForbidden handles this case with default header values.

Forbidden
*/
type PutRecordsEnableIDForbidden struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PutRecordsEnableIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /records/enable/{id}][%d] putRecordsEnableIdForbidden  %+v", 403, o.Payload)
}

func (o *PutRecordsEnableIDForbidden) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutRecordsEnableIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordsEnableIDNotFound creates a PutRecordsEnableIDNotFound with default headers values
func NewPutRecordsEnableIDNotFound() *PutRecordsEnableIDNotFound {
	return &PutRecordsEnableIDNotFound{}
}

/*PutRecordsEnableIDNotFound handles this case with default header values.

Not Found
*/
type PutRecordsEnableIDNotFound struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PutRecordsEnableIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /records/enable/{id}][%d] putRecordsEnableIdNotFound  %+v", 404, o.Payload)
}

func (o *PutRecordsEnableIDNotFound) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutRecordsEnableIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordsEnableIDInternalServerError creates a PutRecordsEnableIDInternalServerError with default headers values
func NewPutRecordsEnableIDInternalServerError() *PutRecordsEnableIDInternalServerError {
	return &PutRecordsEnableIDInternalServerError{}
}

/*PutRecordsEnableIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutRecordsEnableIDInternalServerError struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PutRecordsEnableIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /records/enable/{id}][%d] putRecordsEnableIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRecordsEnableIDInternalServerError) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutRecordsEnableIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
