// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser PluginConfigUser plugin config user
//
// swagger:model PluginConfigUser
type PluginConfigUser struct {

	// g ID
	GID uint32 `json:"GID,omitempty"`

	// UID
	UID uint32 `json:"UID,omitempty"`
}

// Validate validates this plugin config user
func (m *PluginConfigUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plugin config user based on context it is used
func (m *PluginConfigUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigUser) UnmarshalBinary(b []byte) error {
	var res PluginConfigUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
