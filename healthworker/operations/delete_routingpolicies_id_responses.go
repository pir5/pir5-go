// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/pir5/pir5-go/healthworker/model"
)

// DeleteRoutingpoliciesIDReader is a Reader for the DeleteRoutingpoliciesID structure.
type DeleteRoutingpoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingpoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingpoliciesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRoutingpoliciesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingpoliciesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingpoliciesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRoutingpoliciesIDNoContent creates a DeleteRoutingpoliciesIDNoContent with default headers values
func NewDeleteRoutingpoliciesIDNoContent() *DeleteRoutingpoliciesIDNoContent {
	return &DeleteRoutingpoliciesIDNoContent{}
}

/*DeleteRoutingpoliciesIDNoContent handles this case with default header values.

No Content
*/
type DeleteRoutingpoliciesIDNoContent struct {
	Payload *model.ModelRoutingPolicy
}

func (o *DeleteRoutingpoliciesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /routingpolicies/{id}][%d] deleteRoutingpoliciesIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRoutingpoliciesIDNoContent) GetPayload() *model.ModelRoutingPolicy {
	return o.Payload
}

func (o *DeleteRoutingpoliciesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ModelRoutingPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingpoliciesIDForbidden creates a DeleteRoutingpoliciesIDForbidden with default headers values
func NewDeleteRoutingpoliciesIDForbidden() *DeleteRoutingpoliciesIDForbidden {
	return &DeleteRoutingpoliciesIDForbidden{}
}

/*DeleteRoutingpoliciesIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteRoutingpoliciesIDForbidden struct {
	Payload model.HealthWorkerHTTPError
}

func (o *DeleteRoutingpoliciesIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /routingpolicies/{id}][%d] deleteRoutingpoliciesIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingpoliciesIDForbidden) GetPayload() model.HealthWorkerHTTPError {
	return o.Payload
}

func (o *DeleteRoutingpoliciesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingpoliciesIDNotFound creates a DeleteRoutingpoliciesIDNotFound with default headers values
func NewDeleteRoutingpoliciesIDNotFound() *DeleteRoutingpoliciesIDNotFound {
	return &DeleteRoutingpoliciesIDNotFound{}
}

/*DeleteRoutingpoliciesIDNotFound handles this case with default header values.

Not Found
*/
type DeleteRoutingpoliciesIDNotFound struct {
	Payload model.HealthWorkerHTTPError
}

func (o *DeleteRoutingpoliciesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /routingpolicies/{id}][%d] deleteRoutingpoliciesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingpoliciesIDNotFound) GetPayload() model.HealthWorkerHTTPError {
	return o.Payload
}

func (o *DeleteRoutingpoliciesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingpoliciesIDInternalServerError creates a DeleteRoutingpoliciesIDInternalServerError with default headers values
func NewDeleteRoutingpoliciesIDInternalServerError() *DeleteRoutingpoliciesIDInternalServerError {
	return &DeleteRoutingpoliciesIDInternalServerError{}
}

/*DeleteRoutingpoliciesIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteRoutingpoliciesIDInternalServerError struct {
	Payload model.HealthWorkerHTTPError
}

func (o *DeleteRoutingpoliciesIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /routingpolicies/{id}][%d] deleteRoutingpoliciesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingpoliciesIDInternalServerError) GetPayload() model.HealthWorkerHTTPError {
	return o.Payload
}

func (o *DeleteRoutingpoliciesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
