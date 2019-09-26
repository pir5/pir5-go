// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ModelHealthCheckParams model health check params
// swagger:model model.HealthCheckParams
type ModelHealthCheckParams struct {

	// addr
	Addr string `json:"addr,omitempty"`

	// host name
	HostName string `json:"hostName,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// protocol
	Protocol int64 `json:"protocol,omitempty"`

	// search word
	SearchWord string `json:"searchWord,omitempty"`

	// timeout
	Timeout string `json:"timeout,omitempty"`
}

// Validate validates this model health check params
func (m *ModelHealthCheckParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelHealthCheckParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelHealthCheckParams) UnmarshalBinary(b []byte) error {
	var res ModelHealthCheckParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}