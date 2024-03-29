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

// GetRoutingpoliciesReader is a Reader for the GetRoutingpolicies structure.
type GetRoutingpoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingpoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingpoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoutingpoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingpoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRoutingpoliciesOK creates a GetRoutingpoliciesOK with default headers values
func NewGetRoutingpoliciesOK() *GetRoutingpoliciesOK {
	return &GetRoutingpoliciesOK{}
}

/*GetRoutingpoliciesOK handles this case with default header values.

OK
*/
type GetRoutingpoliciesOK struct {
	Payload []*model.ModelRoutingPolicy
}

func (o *GetRoutingpoliciesOK) Error() string {
	return fmt.Sprintf("[GET /routingpolicies][%d] getRoutingpoliciesOK  %+v", 200, o.Payload)
}

func (o *GetRoutingpoliciesOK) GetPayload() []*model.ModelRoutingPolicy {
	return o.Payload
}

func (o *GetRoutingpoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingpoliciesNotFound creates a GetRoutingpoliciesNotFound with default headers values
func NewGetRoutingpoliciesNotFound() *GetRoutingpoliciesNotFound {
	return &GetRoutingpoliciesNotFound{}
}

/*GetRoutingpoliciesNotFound handles this case with default header values.

Not Found
*/
type GetRoutingpoliciesNotFound struct {
	Payload model.HealthWorkerHTTPError
}

func (o *GetRoutingpoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /routingpolicies][%d] getRoutingpoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingpoliciesNotFound) GetPayload() model.HealthWorkerHTTPError {
	return o.Payload
}

func (o *GetRoutingpoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingpoliciesInternalServerError creates a GetRoutingpoliciesInternalServerError with default headers values
func NewGetRoutingpoliciesInternalServerError() *GetRoutingpoliciesInternalServerError {
	return &GetRoutingpoliciesInternalServerError{}
}

/*GetRoutingpoliciesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetRoutingpoliciesInternalServerError struct {
	Payload model.HealthWorkerHTTPError
}

func (o *GetRoutingpoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /routingpolicies][%d] getRoutingpoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingpoliciesInternalServerError) GetPayload() model.HealthWorkerHTTPError {
	return o.Payload
}

func (o *GetRoutingpoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
