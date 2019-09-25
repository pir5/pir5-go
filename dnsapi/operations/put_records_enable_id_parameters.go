// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pir5/pdns-api/models"
)

// NewPutRecordsEnableIDParams creates a new PutRecordsEnableIDParams object
// with the default values initialized.
func NewPutRecordsEnableIDParams() *PutRecordsEnableIDParams {
	var ()
	return &PutRecordsEnableIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRecordsEnableIDParamsWithTimeout creates a new PutRecordsEnableIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRecordsEnableIDParamsWithTimeout(timeout time.Duration) *PutRecordsEnableIDParams {
	var ()
	return &PutRecordsEnableIDParams{

		timeout: timeout,
	}
}

// NewPutRecordsEnableIDParamsWithContext creates a new PutRecordsEnableIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRecordsEnableIDParamsWithContext(ctx context.Context) *PutRecordsEnableIDParams {
	var ()
	return &PutRecordsEnableIDParams{

		Context: ctx,
	}
}

// NewPutRecordsEnableIDParamsWithHTTPClient creates a new PutRecordsEnableIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRecordsEnableIDParamsWithHTTPClient(client *http.Client) *PutRecordsEnableIDParams {
	var ()
	return &PutRecordsEnableIDParams{
		HTTPClient: client,
	}
}

/*PutRecordsEnableIDParams contains all the parameters to send to the API endpoint
for the put records enable ID operation typically these are written to a http.Request
*/
type PutRecordsEnableIDParams struct {

	/*ID
	  Record ID

	*/
	ID int64
	/*Record
	  Record Object

	*/
	Record *models.ModelRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put records enable ID params
func (o *PutRecordsEnableIDParams) WithTimeout(timeout time.Duration) *PutRecordsEnableIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put records enable ID params
func (o *PutRecordsEnableIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put records enable ID params
func (o *PutRecordsEnableIDParams) WithContext(ctx context.Context) *PutRecordsEnableIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put records enable ID params
func (o *PutRecordsEnableIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put records enable ID params
func (o *PutRecordsEnableIDParams) WithHTTPClient(client *http.Client) *PutRecordsEnableIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put records enable ID params
func (o *PutRecordsEnableIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put records enable ID params
func (o *PutRecordsEnableIDParams) WithID(id int64) *PutRecordsEnableIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put records enable ID params
func (o *PutRecordsEnableIDParams) SetID(id int64) {
	o.ID = id
}

// WithRecord adds the record to the put records enable ID params
func (o *PutRecordsEnableIDParams) WithRecord(record *models.ModelRecord) *PutRecordsEnableIDParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put records enable ID params
func (o *PutRecordsEnableIDParams) SetRecord(record *models.ModelRecord) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutRecordsEnableIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
