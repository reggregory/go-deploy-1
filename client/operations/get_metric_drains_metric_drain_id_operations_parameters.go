// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMetricDrainsMetricDrainIDOperationsParams creates a new GetMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized.
func NewGetMetricDrainsMetricDrainIDOperationsParams() *GetMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &GetMetricDrainsMetricDrainIDOperationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetricDrainsMetricDrainIDOperationsParamsWithTimeout creates a new GetMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMetricDrainsMetricDrainIDOperationsParamsWithTimeout(timeout time.Duration) *GetMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &GetMetricDrainsMetricDrainIDOperationsParams{

		timeout: timeout,
	}
}

// NewGetMetricDrainsMetricDrainIDOperationsParamsWithContext creates a new GetMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMetricDrainsMetricDrainIDOperationsParamsWithContext(ctx context.Context) *GetMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &GetMetricDrainsMetricDrainIDOperationsParams{

		Context: ctx,
	}
}

// NewGetMetricDrainsMetricDrainIDOperationsParamsWithHTTPClient creates a new GetMetricDrainsMetricDrainIDOperationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMetricDrainsMetricDrainIDOperationsParamsWithHTTPClient(client *http.Client) *GetMetricDrainsMetricDrainIDOperationsParams {
	var ()
	return &GetMetricDrainsMetricDrainIDOperationsParams{
		HTTPClient: client,
	}
}

/*GetMetricDrainsMetricDrainIDOperationsParams contains all the parameters to send to the API endpoint
for the get metric drains metric drain ID operations operation typically these are written to a http.Request
*/
type GetMetricDrainsMetricDrainIDOperationsParams struct {

	/*MetricDrainID
	  metric_drain_id

	*/
	MetricDrainID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) WithTimeout(timeout time.Duration) *GetMetricDrainsMetricDrainIDOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) WithContext(ctx context.Context) *GetMetricDrainsMetricDrainIDOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) WithHTTPClient(client *http.Client) *GetMetricDrainsMetricDrainIDOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetricDrainID adds the metricDrainID to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) WithMetricDrainID(metricDrainID int64) *GetMetricDrainsMetricDrainIDOperationsParams {
	o.SetMetricDrainID(metricDrainID)
	return o
}

// SetMetricDrainID adds the metricDrainId to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) SetMetricDrainID(metricDrainID int64) {
	o.MetricDrainID = metricDrainID
}

// WithPage adds the page to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) WithPage(page *int64) *GetMetricDrainsMetricDrainIDOperationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get metric drains metric drain ID operations params
func (o *GetMetricDrainsMetricDrainIDOperationsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetricDrainsMetricDrainIDOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metric_drain_id
	if err := r.SetPathParam("metric_drain_id", swag.FormatInt64(o.MetricDrainID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
