// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2003EmbeddedApps inline response 200 3 embedded apps
// swagger:model inline_response_200_3__embedded_apps
type InlineResponse2003EmbeddedApps struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// embedded
	Embedded *InlineResponse2003EmbeddedEmbedded `json:"_embedded,omitempty"`

	// links
	Links *InlineResponse2003EmbeddedLinks `json:"_links,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// git repo
	GitRepo string `json:"git_repo,omitempty"`

	// handle
	Handle string `json:"handle,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this inline response 200 3 embedded apps
func (m *InlineResponse2003EmbeddedApps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2003EmbeddedApps) validateEmbedded(formats strfmt.Registry) error {

	if swag.IsZero(m.Embedded) { // not required
		return nil
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedApps) validateLinks(formats strfmt.Registry) error {

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
func (m *InlineResponse2003EmbeddedApps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedApps) UnmarshalBinary(b []byte) error {
	var res InlineResponse2003EmbeddedApps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}