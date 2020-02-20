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

// PostEphemeralSessionsEphemeralSessionIDOperationsReader is a Reader for the PostEphemeralSessionsEphemeralSessionIDOperations structure.
type PostEphemeralSessionsEphemeralSessionIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostEphemeralSessionsEphemeralSessionIDOperationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostEphemeralSessionsEphemeralSessionIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostEphemeralSessionsEphemeralSessionIDOperationsCreated creates a PostEphemeralSessionsEphemeralSessionIDOperationsCreated with default headers values
func NewPostEphemeralSessionsEphemeralSessionIDOperationsCreated() *PostEphemeralSessionsEphemeralSessionIDOperationsCreated {
	return &PostEphemeralSessionsEphemeralSessionIDOperationsCreated{}
}

/*PostEphemeralSessionsEphemeralSessionIDOperationsCreated handles this case with default header values.

successful
*/
type PostEphemeralSessionsEphemeralSessionIDOperationsCreated struct {
	Payload *models.InlineResponse20028
}

func (o *PostEphemeralSessionsEphemeralSessionIDOperationsCreated) Error() string {
	return fmt.Sprintf("[POST /ephemeral_sessions/{ephemeral_session_id}/operations][%d] postEphemeralSessionsEphemeralSessionIdOperationsCreated  %+v", 201, o.Payload)
}

func (o *PostEphemeralSessionsEphemeralSessionIDOperationsCreated) GetPayload() *models.InlineResponse20028 {
	return o.Payload
}

func (o *PostEphemeralSessionsEphemeralSessionIDOperationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20028)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEphemeralSessionsEphemeralSessionIDOperationsDefault creates a PostEphemeralSessionsEphemeralSessionIDOperationsDefault with default headers values
func NewPostEphemeralSessionsEphemeralSessionIDOperationsDefault(code int) *PostEphemeralSessionsEphemeralSessionIDOperationsDefault {
	return &PostEphemeralSessionsEphemeralSessionIDOperationsDefault{
		_statusCode: code,
	}
}

/*PostEphemeralSessionsEphemeralSessionIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostEphemeralSessionsEphemeralSessionIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post ephemeral sessions ephemeral session ID operations default response
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *PostEphemeralSessionsEphemeralSessionIDOperationsDefault) Error() string {
	return fmt.Sprintf("[POST /ephemeral_sessions/{ephemeral_session_id}/operations][%d] PostEphemeralSessionsEphemeralSessionIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *PostEphemeralSessionsEphemeralSessionIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostEphemeralSessionsEphemeralSessionIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}