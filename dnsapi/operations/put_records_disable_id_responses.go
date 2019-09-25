// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pir5/pdns-api/models"
)

// PutRecordsDisableIDReader is a Reader for the PutRecordsDisableID structure.
type PutRecordsDisableIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRecordsDisableIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRecordsDisableIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPutRecordsDisableIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRecordsDisableIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRecordsDisableIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutRecordsDisableIDOK creates a PutRecordsDisableIDOK with default headers values
func NewPutRecordsDisableIDOK() *PutRecordsDisableIDOK {
	return &PutRecordsDisableIDOK{}
}

/*PutRecordsDisableIDOK handles this case with default header values.

OK
*/
type PutRecordsDisableIDOK struct {
	Payload *models.ModelRecord
}

func (o *PutRecordsDisableIDOK) Error() string {
	return fmt.Sprintf("[PUT /records/disable/{id}][%d] putRecordsDisableIdOK  %+v", 200, o.Payload)
}

func (o *PutRecordsDisableIDOK) GetPayload() *models.ModelRecord {
	return o.Payload
}

func (o *PutRecordsDisableIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordsDisableIDForbidden creates a PutRecordsDisableIDForbidden with default headers values
func NewPutRecordsDisableIDForbidden() *PutRecordsDisableIDForbidden {
	return &PutRecordsDisableIDForbidden{}
}

/*PutRecordsDisableIDForbidden handles this case with default header values.

Forbidden
*/
type PutRecordsDisableIDForbidden struct {
	Payload models.PdnsAPIHTTPError
}

func (o *PutRecordsDisableIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /records/disable/{id}][%d] putRecordsDisableIdForbidden  %+v", 403, o.Payload)
}

func (o *PutRecordsDisableIDForbidden) GetPayload() models.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutRecordsDisableIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordsDisableIDNotFound creates a PutRecordsDisableIDNotFound with default headers values
func NewPutRecordsDisableIDNotFound() *PutRecordsDisableIDNotFound {
	return &PutRecordsDisableIDNotFound{}
}

/*PutRecordsDisableIDNotFound handles this case with default header values.

Not Found
*/
type PutRecordsDisableIDNotFound struct {
	Payload models.PdnsAPIHTTPError
}

func (o *PutRecordsDisableIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /records/disable/{id}][%d] putRecordsDisableIdNotFound  %+v", 404, o.Payload)
}

func (o *PutRecordsDisableIDNotFound) GetPayload() models.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutRecordsDisableIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordsDisableIDInternalServerError creates a PutRecordsDisableIDInternalServerError with default headers values
func NewPutRecordsDisableIDInternalServerError() *PutRecordsDisableIDInternalServerError {
	return &PutRecordsDisableIDInternalServerError{}
}

/*PutRecordsDisableIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutRecordsDisableIDInternalServerError struct {
	Payload models.PdnsAPIHTTPError
}

func (o *PutRecordsDisableIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /records/disable/{id}][%d] putRecordsDisableIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRecordsDisableIDInternalServerError) GetPayload() models.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PutRecordsDisableIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
