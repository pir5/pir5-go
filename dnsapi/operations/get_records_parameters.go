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
)

// NewGetRecordsParams creates a new GetRecordsParams object
// with the default values initialized.
func NewGetRecordsParams() *GetRecordsParams {
	var ()
	return &GetRecordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordsParamsWithTimeout creates a new GetRecordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordsParamsWithTimeout(timeout time.Duration) *GetRecordsParams {
	var ()
	return &GetRecordsParams{

		timeout: timeout,
	}
}

// NewGetRecordsParamsWithContext creates a new GetRecordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordsParamsWithContext(ctx context.Context) *GetRecordsParams {
	var ()
	return &GetRecordsParams{

		Context: ctx,
	}
}

// NewGetRecordsParamsWithHTTPClient creates a new GetRecordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordsParamsWithHTTPClient(client *http.Client) *GetRecordsParams {
	var ()
	return &GetRecordsParams{
		HTTPClient: client,
	}
}

/*GetRecordsParams contains all the parameters to send to the API endpoint
for the get records operation typically these are written to a http.Request
*/
type GetRecordsParams struct {

	/*DomainID
	  Domain ID

	*/
	DomainID *int64
	/*ID
	  Record ID

	*/
	ID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get records params
func (o *GetRecordsParams) WithTimeout(timeout time.Duration) *GetRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get records params
func (o *GetRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get records params
func (o *GetRecordsParams) WithContext(ctx context.Context) *GetRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get records params
func (o *GetRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get records params
func (o *GetRecordsParams) WithHTTPClient(client *http.Client) *GetRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get records params
func (o *GetRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get records params
func (o *GetRecordsParams) WithDomainID(domainID *int64) *GetRecordsParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get records params
func (o *GetRecordsParams) SetDomainID(domainID *int64) {
	o.DomainID = domainID
}

// WithID adds the id to the get records params
func (o *GetRecordsParams) WithID(id *int64) *GetRecordsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get records params
func (o *GetRecordsParams) SetID(id *int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID int64
		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := swag.FormatInt64(qrDomainID)
		if qDomainID != "" {
			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID int64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
