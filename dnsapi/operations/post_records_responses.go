// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/pir5/pdns-api/model"
)

// PostRecordsReader is a Reader for the PostRecords structure.
type PostRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRecordsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostRecordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRecordsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRecordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRecordsCreated creates a PostRecordsCreated with default headers values
func NewPostRecordsCreated() *PostRecordsCreated {
	return &PostRecordsCreated{}
}

/*PostRecordsCreated handles this case with default header values.

Created
*/
type PostRecordsCreated struct {
	Payload *model.ModelRecord
}

func (o *PostRecordsCreated) Error() string {
	return fmt.Sprintf("[POST /records][%d] postRecordsCreated  %+v", 201, o.Payload)
}

func (o *PostRecordsCreated) GetPayload() *model.ModelRecord {
	return o.Payload
}

func (o *PostRecordsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ModelRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordsForbidden creates a PostRecordsForbidden with default headers values
func NewPostRecordsForbidden() *PostRecordsForbidden {
	return &PostRecordsForbidden{}
}

/*PostRecordsForbidden handles this case with default header values.

Forbidden
*/
type PostRecordsForbidden struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PostRecordsForbidden) Error() string {
	return fmt.Sprintf("[POST /records][%d] postRecordsForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordsForbidden) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PostRecordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordsNotFound creates a PostRecordsNotFound with default headers values
func NewPostRecordsNotFound() *PostRecordsNotFound {
	return &PostRecordsNotFound{}
}

/*PostRecordsNotFound handles this case with default header values.

Not Found
*/
type PostRecordsNotFound struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PostRecordsNotFound) Error() string {
	return fmt.Sprintf("[POST /records][%d] postRecordsNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordsNotFound) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PostRecordsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordsInternalServerError creates a PostRecordsInternalServerError with default headers values
func NewPostRecordsInternalServerError() *PostRecordsInternalServerError {
	return &PostRecordsInternalServerError{}
}

/*PostRecordsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostRecordsInternalServerError struct {
	Payload model.PdnsAPIHTTPError
}

func (o *PostRecordsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /records][%d] postRecordsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordsInternalServerError) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *PostRecordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
