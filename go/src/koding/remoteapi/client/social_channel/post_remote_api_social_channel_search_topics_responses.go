package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialChannelSearchTopicsReader is a Reader for the PostRemoteAPISocialChannelSearchTopics structure.
type PostRemoteAPISocialChannelSearchTopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialChannelSearchTopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialChannelSearchTopicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialChannelSearchTopicsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialChannelSearchTopicsOK creates a PostRemoteAPISocialChannelSearchTopicsOK with default headers values
func NewPostRemoteAPISocialChannelSearchTopicsOK() *PostRemoteAPISocialChannelSearchTopicsOK {
	return &PostRemoteAPISocialChannelSearchTopicsOK{}
}

/*PostRemoteAPISocialChannelSearchTopicsOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPISocialChannelSearchTopicsOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialChannelSearchTopicsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.searchTopics][%d] postRemoteApiSocialChannelSearchTopicsOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialChannelSearchTopicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialChannelSearchTopicsUnauthorized creates a PostRemoteAPISocialChannelSearchTopicsUnauthorized with default headers values
func NewPostRemoteAPISocialChannelSearchTopicsUnauthorized() *PostRemoteAPISocialChannelSearchTopicsUnauthorized {
	return &PostRemoteAPISocialChannelSearchTopicsUnauthorized{}
}

/*PostRemoteAPISocialChannelSearchTopicsUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialChannelSearchTopicsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialChannelSearchTopicsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.searchTopics][%d] postRemoteApiSocialChannelSearchTopicsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialChannelSearchTopicsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
