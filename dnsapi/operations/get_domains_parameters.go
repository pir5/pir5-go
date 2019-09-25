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

// NewGetDomainsParams creates a new GetDomainsParams object
// with the default values initialized.
func NewGetDomainsParams() *GetDomainsParams {
	var ()
	return &GetDomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainsParamsWithTimeout creates a new GetDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainsParamsWithTimeout(timeout time.Duration) *GetDomainsParams {
	var ()
	return &GetDomainsParams{

		timeout: timeout,
	}
}

// NewGetDomainsParamsWithContext creates a new GetDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainsParamsWithContext(ctx context.Context) *GetDomainsParams {
	var ()
	return &GetDomainsParams{

		Context: ctx,
	}
}

// NewGetDomainsParamsWithHTTPClient creates a new GetDomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainsParamsWithHTTPClient(client *http.Client) *GetDomainsParams {
	var ()
	return &GetDomainsParams{
		HTTPClient: client,
	}
}

/*GetDomainsParams contains all the parameters to send to the API endpoint
for the get domains operation typically these are written to a http.Request
*/
type GetDomainsParams struct {

	/*ID
	  Domain ID

	*/
	ID *int64
	/*Name
	  Name

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domains params
func (o *GetDomainsParams) WithTimeout(timeout time.Duration) *GetDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domains params
func (o *GetDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domains params
func (o *GetDomainsParams) WithContext(ctx context.Context) *GetDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domains params
func (o *GetDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domains params
func (o *GetDomainsParams) WithHTTPClient(client *http.Client) *GetDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domains params
func (o *GetDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get domains params
func (o *GetDomainsParams) WithID(id *int64) *GetDomainsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get domains params
func (o *GetDomainsParams) SetID(id *int64) {
	o.ID = id
}

// WithName adds the name to the get domains params
func (o *GetDomainsParams) WithName(name *string) *GetDomainsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get domains params
func (o *GetDomainsParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
