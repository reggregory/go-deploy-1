// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineResponse20027EmbeddedDrainConfiguration inline response 200 27 embedded drain configuration
// swagger:model inline_response_200_27__embedded_drain_configuration
type InlineResponse20027EmbeddedDrainConfiguration struct {

	// address
	// Format: uri
	Address strfmt.URI `json:"address,omitempty"`

	// database
	Database string `json:"database,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this inline response 200 27 embedded drain configuration
func (m *InlineResponse20027EmbeddedDrainConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20027EmbeddedDrainConfiguration) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.FormatOf("address", "body", "uri", m.Address.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20027EmbeddedDrainConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20027EmbeddedDrainConfiguration) UnmarshalBinary(b []byte) error {
	var res InlineResponse20027EmbeddedDrainConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
