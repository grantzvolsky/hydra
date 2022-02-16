// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HealthStatus HealthStatus HealthStatus health status
//
// swagger:model healthStatus
type HealthStatus struct {

	// Status always contains "ok".
	Status string `json:"status,omitempty"`
}

// Validate validates this health status
func (m *HealthStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health status based on context it is used
func (m *HealthStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthStatus) UnmarshalBinary(b []byte) error {
	var res HealthStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
