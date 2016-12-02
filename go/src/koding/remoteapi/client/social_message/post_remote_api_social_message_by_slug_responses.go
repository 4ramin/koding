package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialMessageBySlugReader is a Reader for the PostRemoteAPISocialMessageBySlug structure.
type PostRemoteAPISocialMessageBySlugReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialMessageBySlugReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialMessageBySlugOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialMessageBySlugUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialMessageBySlugOK creates a PostRemoteAPISocialMessageBySlugOK with default headers values
func NewPostRemoteAPISocialMessageBySlugOK() *PostRemoteAPISocialMessageBySlugOK {
	return &PostRemoteAPISocialMessageBySlugOK{}
}

/*PostRemoteAPISocialMessageBySlugOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPISocialMessageBySlugOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialMessageBySlugOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.bySlug][%d] postRemoteApiSocialMessageBySlugOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialMessageBySlugOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialMessageBySlugUnauthorized creates a PostRemoteAPISocialMessageBySlugUnauthorized with default headers values
func NewPostRemoteAPISocialMessageBySlugUnauthorized() *PostRemoteAPISocialMessageBySlugUnauthorized {
	return &PostRemoteAPISocialMessageBySlugUnauthorized{}
}

/*PostRemoteAPISocialMessageBySlugUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialMessageBySlugUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialMessageBySlugUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.bySlug][%d] postRemoteApiSocialMessageBySlugUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialMessageBySlugUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
