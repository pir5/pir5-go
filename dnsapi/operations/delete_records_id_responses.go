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

// DeleteRecordsIDReader is a Reader for the DeleteRecordsID structure.
type DeleteRecordsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRecordsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRecordsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecordsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRecordsIDNoContent creates a DeleteRecordsIDNoContent with default headers values
func NewDeleteRecordsIDNoContent() *DeleteRecordsIDNoContent {
	return &DeleteRecordsIDNoContent{}
}

/*DeleteRecordsIDNoContent handles this case with default header values.

No Content
*/
type DeleteRecordsIDNoContent struct {
	Payload *model.ModelRecord
}

func (o *DeleteRecordsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /records/{id}][%d] deleteRecordsIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRecordsIDNoContent) GetPayload() *model.ModelRecord {
	return o.Payload
}

func (o *DeleteRecordsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ModelRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordsIDForbidden creates a DeleteRecordsIDForbidden with default headers values
func NewDeleteRecordsIDForbidden() *DeleteRecordsIDForbidden {
	return &DeleteRecordsIDForbidden{}
}

/*DeleteRecordsIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteRecordsIDForbidden struct {
	Payload model.PdnsAPIHTTPError
}

func (o *DeleteRecordsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /records/{id}][%d] deleteRecordsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRecordsIDForbidden) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *DeleteRecordsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordsIDNotFound creates a DeleteRecordsIDNotFound with default headers values
func NewDeleteRecordsIDNotFound() *DeleteRecordsIDNotFound {
	return &DeleteRecordsIDNotFound{}
}

/*DeleteRecordsIDNotFound handles this case with default header values.

Not Found
*/
type DeleteRecordsIDNotFound struct {
	Payload model.PdnsAPIHTTPError
}

func (o *DeleteRecordsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /records/{id}][%d] deleteRecordsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRecordsIDNotFound) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *DeleteRecordsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordsIDInternalServerError creates a DeleteRecordsIDInternalServerError with default headers values
func NewDeleteRecordsIDInternalServerError() *DeleteRecordsIDInternalServerError {
	return &DeleteRecordsIDInternalServerError{}
}

/*DeleteRecordsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteRecordsIDInternalServerError struct {
	Payload model.PdnsAPIHTTPError
}

func (o *DeleteRecordsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /records/{id}][%d] deleteRecordsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRecordsIDInternalServerError) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *DeleteRecordsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
