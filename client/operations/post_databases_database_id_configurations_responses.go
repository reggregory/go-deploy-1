// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/reggregory/go-deploy/models"
)

// PostDatabasesDatabaseIDConfigurationsReader is a Reader for the PostDatabasesDatabaseIDConfigurations structure.
type PostDatabasesDatabaseIDConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDatabasesDatabaseIDConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostDatabasesDatabaseIDConfigurationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostDatabasesDatabaseIDConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostDatabasesDatabaseIDConfigurationsCreated creates a PostDatabasesDatabaseIDConfigurationsCreated with default headers values
func NewPostDatabasesDatabaseIDConfigurationsCreated() *PostDatabasesDatabaseIDConfigurationsCreated {
	return &PostDatabasesDatabaseIDConfigurationsCreated{}
}

/*PostDatabasesDatabaseIDConfigurationsCreated handles this case with default header values.

successful
*/
type PostDatabasesDatabaseIDConfigurationsCreated struct {
	Payload *models.InlineResponse2013
}

func (o *PostDatabasesDatabaseIDConfigurationsCreated) Error() string {
	return fmt.Sprintf("[POST /databases/{database_id}/configurations][%d] postDatabasesDatabaseIdConfigurationsCreated  %+v", 201, o.Payload)
}

func (o *PostDatabasesDatabaseIDConfigurationsCreated) GetPayload() *models.InlineResponse2013 {
	return o.Payload
}

func (o *PostDatabasesDatabaseIDConfigurationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2013)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatabasesDatabaseIDConfigurationsDefault creates a PostDatabasesDatabaseIDConfigurationsDefault with default headers values
func NewPostDatabasesDatabaseIDConfigurationsDefault(code int) *PostDatabasesDatabaseIDConfigurationsDefault {
	return &PostDatabasesDatabaseIDConfigurationsDefault{
		_statusCode: code,
	}
}

/*PostDatabasesDatabaseIDConfigurationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostDatabasesDatabaseIDConfigurationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post databases database ID configurations default response
func (o *PostDatabasesDatabaseIDConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *PostDatabasesDatabaseIDConfigurationsDefault) Error() string {
	return fmt.Sprintf("[POST /databases/{database_id}/configurations][%d] PostDatabasesDatabaseIDConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *PostDatabasesDatabaseIDConfigurationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostDatabasesDatabaseIDConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}