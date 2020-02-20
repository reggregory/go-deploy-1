// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20010EmbeddedConfigurations inline response 200 10 embedded configurations
// swagger:model inline_response_200_10__embedded_configurations
type InlineResponse20010EmbeddedConfigurations struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse20010EmbeddedLinks `json:"_links,omitempty"`

	// env
	Env interface{} `json:"env,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this inline response 200 10 embedded configurations
func (m *InlineResponse20010EmbeddedConfigurations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20010EmbeddedConfigurations) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20010EmbeddedConfigurations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20010EmbeddedConfigurations) UnmarshalBinary(b []byte) error {
	var res InlineResponse20010EmbeddedConfigurations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}