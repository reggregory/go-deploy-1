// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2008Links inline response 200 8 links
// swagger:model inline_response_200_8__links
type InlineResponse2008Links struct {

	// log drain
	LogDrain *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"log_drain,omitempty"`

	// next
	Next *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"next,omitempty"`

	// prev
	Prev *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"prev,omitempty"`

	// release
	Release *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"release,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"self,omitempty"`
}

// Validate validates this inline response 200 8 links
func (m *InlineResponse2008Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogDrain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2008Links) validateLogDrain(formats strfmt.Registry) error {

	if swag.IsZero(m.LogDrain) { // not required
		return nil
	}

	if m.LogDrain != nil {
		if err := m.LogDrain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_drain")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008Links) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008Links) validatePrev(formats strfmt.Registry) error {

	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if m.Prev != nil {
		if err := m.Prev.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008Links) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008Links) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2008Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2008Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse2008Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
