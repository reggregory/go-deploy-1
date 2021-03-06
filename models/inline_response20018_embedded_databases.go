// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20018EmbeddedDatabases inline response 200 18 embedded databases
// swagger:model inline_response_200_18__embedded_databases
type InlineResponse20018EmbeddedDatabases struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// embedded
	Embedded *InlineResponse20018EmbeddedEmbedded `json:"_embedded,omitempty"`

	// links
	Links *InlineResponse20018EmbeddedLinks `json:"_links,omitempty"`

	// connection url
	ConnectionURL *string `json:"connection_url,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// docker repo
	DockerRepo *string `json:"docker_repo,omitempty"`

	// handle
	Handle string `json:"handle,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// initial container size
	InitialContainerSize *int64 `json:"initial_container_size,omitempty"`

	// initial disk size
	InitialDiskSize *int64 `json:"initial_disk_size,omitempty"`

	// passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	// port mapping
	PortMapping []int64 `json:"port_mapping"`

	// provisioned
	Provisioned bool `json:"provisioned,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type *string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this inline response 200 18 embedded databases
func (m *InlineResponse20018EmbeddedDatabases) Validate(formats strfmt.Registry) error {
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

func (m *InlineResponse20018EmbeddedDatabases) validateEmbedded(formats strfmt.Registry) error {

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

func (m *InlineResponse20018EmbeddedDatabases) validateLinks(formats strfmt.Registry) error {

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
func (m *InlineResponse20018EmbeddedDatabases) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20018EmbeddedDatabases) UnmarshalBinary(b []byte) error {
	var res InlineResponse20018EmbeddedDatabases
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
