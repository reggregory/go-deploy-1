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

// InlineResponse2012 inline response 201 2
// swagger:model inline_response_201_2
type InlineResponse2012 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// links
	// Required: true
	Links *InlineResponse2006EmbeddedLinks `json:"_links"`

	// acme
	// Required: true
	Acme *bool `json:"acme"`

	// certificate arn
	// Required: true
	CertificateArn *string `json:"certificate_arn"`

	// certificate body
	// Required: true
	CertificateBody *string `json:"certificate_body"`

	// certificate chain
	// Required: true
	CertificateChain *string `json:"certificate_chain"`

	// common name
	// Required: true
	CommonName *string `json:"common_name"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// issuer common name
	// Required: true
	IssuerCommonName *string `json:"issuer_common_name"`

	// issuer country
	// Required: true
	IssuerCountry *string `json:"issuer_country"`

	// issuer organization
	// Required: true
	IssuerOrganization *string `json:"issuer_organization"`

	// issuer website
	// Required: true
	IssuerWebsite *string `json:"issuer_website"`

	// leaf certificate
	// Required: true
	LeafCertificate *string `json:"leaf_certificate"`

	// not after
	// Required: true
	NotAfter *string `json:"not_after"`

	// not before
	// Required: true
	NotBefore *string `json:"not_before"`

	// private key
	// Required: true
	PrivateKey *string `json:"private_key"`

	// private key algorithm
	// Required: true
	PrivateKeyAlgorithm *string `json:"private_key_algorithm"`

	// self signed
	// Required: true
	SelfSigned *bool `json:"self_signed"`

	// sha256 fingerprint
	// Required: true
	Sha256Fingerprint *string `json:"sha256_fingerprint"`

	// subject alternative names
	// Required: true
	SubjectAlternativeNames []string `json:"subject_alternative_names"`

	// subject country
	// Required: true
	SubjectCountry *string `json:"subject_country"`

	// subject locale
	// Required: true
	SubjectLocale *string `json:"subject_locale"`

	// subject organization
	// Required: true
	SubjectOrganization *string `json:"subject_organization"`

	// subject state
	// Required: true
	SubjectState *string `json:"subject_state"`

	// trusted
	// Required: true
	Trusted *bool `json:"trusted"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this inline response 201 2
func (m *InlineResponse2012) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommonName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuerCommonName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuerCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuerOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuerWebsite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeafCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKeyAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfSigned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha256Fingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectAlternativeNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrusted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2012) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
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

func (m *InlineResponse2012) validateAcme(formats strfmt.Registry) error {

	if err := validate.Required("acme", "body", m.Acme); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateCertificateArn(formats strfmt.Registry) error {

	if err := validate.Required("certificate_arn", "body", m.CertificateArn); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateCertificateBody(formats strfmt.Registry) error {

	if err := validate.Required("certificate_body", "body", m.CertificateBody); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateCertificateChain(formats strfmt.Registry) error {

	if err := validate.Required("certificate_chain", "body", m.CertificateChain); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateCommonName(formats strfmt.Registry) error {

	if err := validate.Required("common_name", "body", m.CommonName); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateIssuerCommonName(formats strfmt.Registry) error {

	if err := validate.Required("issuer_common_name", "body", m.IssuerCommonName); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateIssuerCountry(formats strfmt.Registry) error {

	if err := validate.Required("issuer_country", "body", m.IssuerCountry); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateIssuerOrganization(formats strfmt.Registry) error {

	if err := validate.Required("issuer_organization", "body", m.IssuerOrganization); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateIssuerWebsite(formats strfmt.Registry) error {

	if err := validate.Required("issuer_website", "body", m.IssuerWebsite); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateLeafCertificate(formats strfmt.Registry) error {

	if err := validate.Required("leaf_certificate", "body", m.LeafCertificate); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateNotAfter(formats strfmt.Registry) error {

	if err := validate.Required("not_after", "body", m.NotAfter); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateNotBefore(formats strfmt.Registry) error {

	if err := validate.Required("not_before", "body", m.NotBefore); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("private_key", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validatePrivateKeyAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("private_key_algorithm", "body", m.PrivateKeyAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateSelfSigned(formats strfmt.Registry) error {

	if err := validate.Required("self_signed", "body", m.SelfSigned); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateSha256Fingerprint(formats strfmt.Registry) error {

	if err := validate.Required("sha256_fingerprint", "body", m.Sha256Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateSubjectAlternativeNames(formats strfmt.Registry) error {

	if err := validate.Required("subject_alternative_names", "body", m.SubjectAlternativeNames); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateSubjectCountry(formats strfmt.Registry) error {

	if err := validate.Required("subject_country", "body", m.SubjectCountry); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateSubjectLocale(formats strfmt.Registry) error {

	if err := validate.Required("subject_locale", "body", m.SubjectLocale); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateSubjectOrganization(formats strfmt.Registry) error {

	if err := validate.Required("subject_organization", "body", m.SubjectOrganization); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateSubjectState(formats strfmt.Registry) error {

	if err := validate.Required("subject_state", "body", m.SubjectState); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateTrusted(formats strfmt.Registry) error {

	if err := validate.Required("trusted", "body", m.Trusted); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2012) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2012) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2012) UnmarshalBinary(b []byte) error {
	var res InlineResponse2012
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
