package j_proposed_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJProposedDomainFetchDomainsReader is a Reader for the PostRemoteAPIJProposedDomainFetchDomains structure.
type PostRemoteAPIJProposedDomainFetchDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJProposedDomainFetchDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJProposedDomainFetchDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJProposedDomainFetchDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJProposedDomainFetchDomainsOK creates a PostRemoteAPIJProposedDomainFetchDomainsOK with default headers values
func NewPostRemoteAPIJProposedDomainFetchDomainsOK() *PostRemoteAPIJProposedDomainFetchDomainsOK {
	return &PostRemoteAPIJProposedDomainFetchDomainsOK{}
}

/*PostRemoteAPIJProposedDomainFetchDomainsOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJProposedDomainFetchDomainsOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJProposedDomainFetchDomainsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.fetchDomains][%d] postRemoteApiJProposedDomainFetchDomainsOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJProposedDomainFetchDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJProposedDomainFetchDomainsUnauthorized creates a PostRemoteAPIJProposedDomainFetchDomainsUnauthorized with default headers values
func NewPostRemoteAPIJProposedDomainFetchDomainsUnauthorized() *PostRemoteAPIJProposedDomainFetchDomainsUnauthorized {
	return &PostRemoteAPIJProposedDomainFetchDomainsUnauthorized{}
}

/*PostRemoteAPIJProposedDomainFetchDomainsUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJProposedDomainFetchDomainsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJProposedDomainFetchDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.fetchDomains][%d] postRemoteApiJProposedDomainFetchDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJProposedDomainFetchDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
