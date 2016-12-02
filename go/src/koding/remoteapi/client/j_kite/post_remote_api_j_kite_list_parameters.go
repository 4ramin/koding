package j_kite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJKiteListParams creates a new PostRemoteAPIJKiteListParams object
// with the default values initialized.
func NewPostRemoteAPIJKiteListParams() *PostRemoteAPIJKiteListParams {
	var ()
	return &PostRemoteAPIJKiteListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJKiteListParamsWithTimeout creates a new PostRemoteAPIJKiteListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJKiteListParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJKiteListParams {
	var ()
	return &PostRemoteAPIJKiteListParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJKiteListParamsWithContext creates a new PostRemoteAPIJKiteListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJKiteListParamsWithContext(ctx context.Context) *PostRemoteAPIJKiteListParams {
	var ()
	return &PostRemoteAPIJKiteListParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJKiteListParams contains all the parameters to send to the API endpoint
for the post remote API j kite list operation typically these are written to a http.Request
*/
type PostRemoteAPIJKiteListParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j kite list params
func (o *PostRemoteAPIJKiteListParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJKiteListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j kite list params
func (o *PostRemoteAPIJKiteListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j kite list params
func (o *PostRemoteAPIJKiteListParams) WithContext(ctx context.Context) *PostRemoteAPIJKiteListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j kite list params
func (o *PostRemoteAPIJKiteListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j kite list params
func (o *PostRemoteAPIJKiteListParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJKiteListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j kite list params
func (o *PostRemoteAPIJKiteListParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJKiteListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
