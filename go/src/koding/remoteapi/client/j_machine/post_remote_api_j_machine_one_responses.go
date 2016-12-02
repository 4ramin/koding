package j_machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJMachineOneReader is a Reader for the PostRemoteAPIJMachineOne structure.
type PostRemoteAPIJMachineOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJMachineOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJMachineOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJMachineOneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJMachineOneOK creates a PostRemoteAPIJMachineOneOK with default headers values
func NewPostRemoteAPIJMachineOneOK() *PostRemoteAPIJMachineOneOK {
	return &PostRemoteAPIJMachineOneOK{}
}

/*PostRemoteAPIJMachineOneOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJMachineOneOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJMachineOneOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JMachine.one][%d] postRemoteApiJMachineOneOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJMachineOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJMachineOneUnauthorized creates a PostRemoteAPIJMachineOneUnauthorized with default headers values
func NewPostRemoteAPIJMachineOneUnauthorized() *PostRemoteAPIJMachineOneUnauthorized {
	return &PostRemoteAPIJMachineOneUnauthorized{}
}

/*PostRemoteAPIJMachineOneUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJMachineOneUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJMachineOneUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JMachine.one][%d] postRemoteApiJMachineOneUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJMachineOneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
