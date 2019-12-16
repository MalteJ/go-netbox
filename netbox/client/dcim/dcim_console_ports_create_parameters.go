// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netbox-community/go-netbox/netbox/models"
)

// NewDcimConsolePortsCreateParams creates a new DcimConsolePortsCreateParams object
// with the default values initialized.
func NewDcimConsolePortsCreateParams() *DcimConsolePortsCreateParams {
	var ()
	return &DcimConsolePortsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsCreateParamsWithTimeout creates a new DcimConsolePortsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortsCreateParamsWithTimeout(timeout time.Duration) *DcimConsolePortsCreateParams {
	var ()
	return &DcimConsolePortsCreateParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortsCreateParamsWithContext creates a new DcimConsolePortsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortsCreateParamsWithContext(ctx context.Context) *DcimConsolePortsCreateParams {
	var ()
	return &DcimConsolePortsCreateParams{

		Context: ctx,
	}
}

// NewDcimConsolePortsCreateParamsWithHTTPClient creates a new DcimConsolePortsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortsCreateParamsWithHTTPClient(client *http.Client) *DcimConsolePortsCreateParams {
	var ()
	return &DcimConsolePortsCreateParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortsCreateParams contains all the parameters to send to the API endpoint
for the dcim console ports create operation typically these are written to a http.Request
*/
type DcimConsolePortsCreateParams struct {

	/*Data*/
	Data *models.WritableConsolePort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithTimeout(timeout time.Duration) *DcimConsolePortsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithContext(ctx context.Context) *DcimConsolePortsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithHTTPClient(client *http.Client) *DcimConsolePortsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithData(data *models.WritableConsolePort) *DcimConsolePortsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetData(data *models.WritableConsolePort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
