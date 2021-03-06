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

// DeleteAppsIDReader is a Reader for the DeleteAppsID structure.
type DeleteAppsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAppsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAppsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAppsIDNoContent creates a DeleteAppsIDNoContent with default headers values
func NewDeleteAppsIDNoContent() *DeleteAppsIDNoContent {
	return &DeleteAppsIDNoContent{}
}

/*DeleteAppsIDNoContent handles this case with default header values.

successful
*/
type DeleteAppsIDNoContent struct {
}

func (o *DeleteAppsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /apps/{id}][%d] deleteAppsIdNoContent ", 204)
}

func (o *DeleteAppsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppsIDDefault creates a DeleteAppsIDDefault with default headers values
func NewDeleteAppsIDDefault(code int) *DeleteAppsIDDefault {
	return &DeleteAppsIDDefault{
		_statusCode: code,
	}
}

/*DeleteAppsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type DeleteAppsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the delete apps ID default response
func (o *DeleteAppsIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAppsIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /apps/{id}][%d] DeleteAppsID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAppsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *DeleteAppsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
