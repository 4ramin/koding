package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTagCreateSynonymIDReader is a Reader for the PostRemoteAPIJTagCreateSynonymID structure.
type PostRemoteAPIJTagCreateSynonymIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagCreateSynonymIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagCreateSynonymIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagCreateSynonymIDOK creates a PostRemoteAPIJTagCreateSynonymIDOK with default headers values
func NewPostRemoteAPIJTagCreateSynonymIDOK() *PostRemoteAPIJTagCreateSynonymIDOK {
	return &PostRemoteAPIJTagCreateSynonymIDOK{}
}

/*PostRemoteAPIJTagCreateSynonymIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJTagCreateSynonymIDOK struct {
	Payload *models.JTag
}

func (o *PostRemoteAPIJTagCreateSynonymIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.createSynonym/{id}][%d] postRemoteApiJTagCreateSynonymIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagCreateSynonymIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
