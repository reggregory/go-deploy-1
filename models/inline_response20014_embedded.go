// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20014Embedded inline response 200 14 embedded
// swagger:model inline_response_200_14__embedded
type InlineResponse20014Embedded struct {

	// databases
	Databases []*InlineResponse20014EmbeddedDatabases `json:"databases"`
}

// Validate validates this inline response 200 14 embedded
func (m *InlineResponse20014Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20014Embedded) validateDatabases(formats strfmt.Registry) error {

	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	for i := 0; i < len(m.Databases); i++ {
		if swag.IsZero(m.Databases[i]) { // not required
			continue
		}

		if m.Databases[i] != nil {
			if err := m.Databases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("databases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20014Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20014Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20014Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
