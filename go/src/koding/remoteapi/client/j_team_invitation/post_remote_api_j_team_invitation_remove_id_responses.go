package j_team_invitation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTeamInvitationRemoveIDReader is a Reader for the PostRemoteAPIJTeamInvitationRemoveID structure.
type PostRemoteAPIJTeamInvitationRemoveIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTeamInvitationRemoveIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTeamInvitationRemoveIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTeamInvitationRemoveIDOK creates a PostRemoteAPIJTeamInvitationRemoveIDOK with default headers values
func NewPostRemoteAPIJTeamInvitationRemoveIDOK() *PostRemoteAPIJTeamInvitationRemoveIDOK {
	return &PostRemoteAPIJTeamInvitationRemoveIDOK{}
}

/*PostRemoteAPIJTeamInvitationRemoveIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJTeamInvitationRemoveIDOK struct {
	Payload *models.JTeamInvitation
}

func (o *PostRemoteAPIJTeamInvitationRemoveIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTeamInvitation.remove/{id}][%d] postRemoteApiJTeamInvitationRemoveIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTeamInvitationRemoveIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JTeamInvitation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
