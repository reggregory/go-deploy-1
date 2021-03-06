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

// GetStacksStackIDVpnTunnelsReader is a Reader for the GetStacksStackIDVpnTunnels structure.
type GetStacksStackIDVpnTunnelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStacksStackIDVpnTunnelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStacksStackIDVpnTunnelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStacksStackIDVpnTunnelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStacksStackIDVpnTunnelsOK creates a GetStacksStackIDVpnTunnelsOK with default headers values
func NewGetStacksStackIDVpnTunnelsOK() *GetStacksStackIDVpnTunnelsOK {
	return &GetStacksStackIDVpnTunnelsOK{}
}

/*GetStacksStackIDVpnTunnelsOK handles this case with default header values.

successful
*/
type GetStacksStackIDVpnTunnelsOK struct {
	Payload *models.InlineResponse20041
}

func (o *GetStacksStackIDVpnTunnelsOK) Error() string {
	return fmt.Sprintf("[GET /stacks/{stack_id}/vpn_tunnels][%d] getStacksStackIdVpnTunnelsOK  %+v", 200, o.Payload)
}

func (o *GetStacksStackIDVpnTunnelsOK) GetPayload() *models.InlineResponse20041 {
	return o.Payload
}

func (o *GetStacksStackIDVpnTunnelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20041)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStacksStackIDVpnTunnelsDefault creates a GetStacksStackIDVpnTunnelsDefault with default headers values
func NewGetStacksStackIDVpnTunnelsDefault(code int) *GetStacksStackIDVpnTunnelsDefault {
	return &GetStacksStackIDVpnTunnelsDefault{
		_statusCode: code,
	}
}

/*GetStacksStackIDVpnTunnelsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetStacksStackIDVpnTunnelsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get stacks stack ID vpn tunnels default response
func (o *GetStacksStackIDVpnTunnelsDefault) Code() int {
	return o._statusCode
}

func (o *GetStacksStackIDVpnTunnelsDefault) Error() string {
	return fmt.Sprintf("[GET /stacks/{stack_id}/vpn_tunnels][%d] GetStacksStackIDVpnTunnels default  %+v", o._statusCode, o.Payload)
}

func (o *GetStacksStackIDVpnTunnelsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetStacksStackIDVpnTunnelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
