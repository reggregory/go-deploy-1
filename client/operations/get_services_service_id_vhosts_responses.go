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

// GetServicesServiceIDVhostsReader is a Reader for the GetServicesServiceIDVhosts structure.
type GetServicesServiceIDVhostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesServiceIDVhostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesServiceIDVhostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServicesServiceIDVhostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesServiceIDVhostsOK creates a GetServicesServiceIDVhostsOK with default headers values
func NewGetServicesServiceIDVhostsOK() *GetServicesServiceIDVhostsOK {
	return &GetServicesServiceIDVhostsOK{}
}

/*GetServicesServiceIDVhostsOK handles this case with default header values.

successful
*/
type GetServicesServiceIDVhostsOK struct {
	Payload *models.InlineResponse20038
}

func (o *GetServicesServiceIDVhostsOK) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}/vhosts][%d] getServicesServiceIdVhostsOK  %+v", 200, o.Payload)
}

func (o *GetServicesServiceIDVhostsOK) GetPayload() *models.InlineResponse20038 {
	return o.Payload
}

func (o *GetServicesServiceIDVhostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20038)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesServiceIDVhostsDefault creates a GetServicesServiceIDVhostsDefault with default headers values
func NewGetServicesServiceIDVhostsDefault(code int) *GetServicesServiceIDVhostsDefault {
	return &GetServicesServiceIDVhostsDefault{
		_statusCode: code,
	}
}

/*GetServicesServiceIDVhostsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetServicesServiceIDVhostsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get services service ID vhosts default response
func (o *GetServicesServiceIDVhostsDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesServiceIDVhostsDefault) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}/vhosts][%d] GetServicesServiceIDVhosts default  %+v", o._statusCode, o.Payload)
}

func (o *GetServicesServiceIDVhostsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetServicesServiceIDVhostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
