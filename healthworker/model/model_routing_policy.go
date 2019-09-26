// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ModelRoutingPolicy model routing policy
// swagger:model model.RoutingPolicy
type ModelRoutingPolicy struct {

	// client
	Client string `json:"client,omitempty"`

	// db
	Db string `json:"db,omitempty"`

	// health check ID
	HealthCheckID int64 `json:"healthCheckID,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// record ID
	RecordID int64 `json:"recordID,omitempty"`

	// type
	Type int64 `json:"type,omitempty"`
}

// Validate validates this model routing policy
func (m *ModelRoutingPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelRoutingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelRoutingPolicy) UnmarshalBinary(b []byte) error {
	var res ModelRoutingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}