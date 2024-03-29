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

// GetDomainsReader is a Reader for the GetDomains structure.
type GetDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDomainsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDomainsOK creates a GetDomainsOK with default headers values
func NewGetDomainsOK() *GetDomainsOK {
	return &GetDomainsOK{}
}

/*GetDomainsOK handles this case with default header values.

OK
*/
type GetDomainsOK struct {
	Payload []*model.ModelDomain
}

func (o *GetDomainsOK) Error() string {
	return fmt.Sprintf("[GET /domains][%d] getDomainsOK  %+v", 200, o.Payload)
}

func (o *GetDomainsOK) GetPayload() []*model.ModelDomain {
	return o.Payload
}

func (o *GetDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsNotFound creates a GetDomainsNotFound with default headers values
func NewGetDomainsNotFound() *GetDomainsNotFound {
	return &GetDomainsNotFound{}
}

/*GetDomainsNotFound handles this case with default header values.

Not Found
*/
type GetDomainsNotFound struct {
	Payload model.PdnsAPIHTTPError
}

func (o *GetDomainsNotFound) Error() string {
	return fmt.Sprintf("[GET /domains][%d] getDomainsNotFound  %+v", 404, o.Payload)
}

func (o *GetDomainsNotFound) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *GetDomainsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsInternalServerError creates a GetDomainsInternalServerError with default headers values
func NewGetDomainsInternalServerError() *GetDomainsInternalServerError {
	return &GetDomainsInternalServerError{}
}

/*GetDomainsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDomainsInternalServerError struct {
	Payload model.PdnsAPIHTTPError
}

func (o *GetDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /domains][%d] getDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDomainsInternalServerError) GetPayload() model.PdnsAPIHTTPError {
	return o.Payload
}

func (o *GetDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
