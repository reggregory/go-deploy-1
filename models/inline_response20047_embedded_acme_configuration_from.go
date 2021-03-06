// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InlineResponse20047EmbeddedAcmeConfigurationFrom inline response 200 47 embedded acme configuration from
// swagger:model inline_response_200_47__embedded_acme_configuration_from
type InlineResponse20047EmbeddedAcmeConfigurationFrom struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this inline response 200 47 embedded acme configuration from
func (m *InlineResponse20047EmbeddedAcmeConfigurationFrom) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20047EmbeddedAcmeConfigurationFrom) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20047EmbeddedAcmeConfigurationFrom) UnmarshalBinary(b []byte) error {
	var res InlineResponse20047EmbeddedAcmeConfigurationFrom
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
