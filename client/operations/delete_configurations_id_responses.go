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

// DeleteConfigurationsIDReader is a Reader for the DeleteConfigurationsID structure.
type DeleteConfigurationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConfigurationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConfigurationsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteConfigurationsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteConfigurationsIDNoContent creates a DeleteConfigurationsIDNoContent with default headers values
func NewDeleteConfigurationsIDNoContent() *DeleteConfigurationsIDNoContent {
	return &DeleteConfigurationsIDNoContent{}
}

/*DeleteConfigurationsIDNoContent handles this case with default header values.

successful
*/
type DeleteConfigurationsIDNoContent struct {
}

func (o *DeleteConfigurationsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /configurations/{id}][%d] deleteConfigurationsIdNoContent ", 204)
}

func (o *DeleteConfigurationsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConfigurationsIDDefault creates a DeleteConfigurationsIDDefault with default headers values
func NewDeleteConfigurationsIDDefault(code int) *DeleteConfigurationsIDDefault {
	return &DeleteConfigurationsIDDefault{
		_statusCode: code,
	}
}

/*DeleteConfigurationsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type DeleteConfigurationsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the delete configurations ID default response
func (o *DeleteConfigurationsIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteConfigurationsIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /configurations/{id}][%d] DeleteConfigurationsID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteConfigurationsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *DeleteConfigurationsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
