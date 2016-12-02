package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJGroupFetchModeratorsIDParams creates a new PostRemoteAPIJGroupFetchModeratorsIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupFetchModeratorsIDParams() *PostRemoteAPIJGroupFetchModeratorsIDParams {
	var ()
	return &PostRemoteAPIJGroupFetchModeratorsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupFetchModeratorsIDParamsWithTimeout creates a new PostRemoteAPIJGroupFetchModeratorsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupFetchModeratorsIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupFetchModeratorsIDParams {
	var ()
	return &PostRemoteAPIJGroupFetchModeratorsIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupFetchModeratorsIDParamsWithContext creates a new PostRemoteAPIJGroupFetchModeratorsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupFetchModeratorsIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupFetchModeratorsIDParams {
	var ()
	return &PostRemoteAPIJGroupFetchModeratorsIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupFetchModeratorsIDParams contains all the parameters to send to the API endpoint
for the post remote API j group fetch moderators ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupFetchModeratorsIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j group fetch moderators ID params
func (o *PostRemoteAPIJGroupFetchModeratorsIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupFetchModeratorsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group fetch moderators ID params
func (o *PostRemoteAPIJGroupFetchModeratorsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group fetch moderators ID params
func (o *PostRemoteAPIJGroupFetchModeratorsIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupFetchModeratorsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group fetch moderators ID params
func (o *PostRemoteAPIJGroupFetchModeratorsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j group fetch moderators ID params
func (o *PostRemoteAPIJGroupFetchModeratorsIDParams) WithID(id string) *PostRemoteAPIJGroupFetchModeratorsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group fetch moderators ID params
func (o *PostRemoteAPIJGroupFetchModeratorsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupFetchModeratorsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
