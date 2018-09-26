///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewAddEndpointParams creates a new AddEndpointParams object
// with the default values initialized.
func NewAddEndpointParams() *AddEndpointParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &AddEndpointParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAddEndpointParamsWithTimeout creates a new AddEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddEndpointParamsWithTimeout(timeout time.Duration) *AddEndpointParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &AddEndpointParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewAddEndpointParamsWithContext creates a new AddEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddEndpointParamsWithContext(ctx context.Context) *AddEndpointParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &AddEndpointParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewAddEndpointParamsWithHTTPClient creates a new AddEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddEndpointParamsWithHTTPClient(client *http.Client) *AddEndpointParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &AddEndpointParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*AddEndpointParams contains all the parameters to send to the API endpoint
for the add endpoint operation typically these are written to a http.Request
*/
type AddEndpointParams struct {

	/*XDispatchOrg*/
	XDispatchOrg *string
	/*XDispatchProject*/
	XDispatchProject *string
	/*Body
	  Endpoint object

	*/
	Body *v1.Endpoint

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add endpoint params
func (o *AddEndpointParams) WithTimeout(timeout time.Duration) *AddEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add endpoint params
func (o *AddEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add endpoint params
func (o *AddEndpointParams) WithContext(ctx context.Context) *AddEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add endpoint params
func (o *AddEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add endpoint params
func (o *AddEndpointParams) WithHTTPClient(client *http.Client) *AddEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add endpoint params
func (o *AddEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the add endpoint params
func (o *AddEndpointParams) WithXDispatchOrg(xDispatchOrg *string) *AddEndpointParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the add endpoint params
func (o *AddEndpointParams) SetXDispatchOrg(xDispatchOrg *string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the add endpoint params
func (o *AddEndpointParams) WithXDispatchProject(xDispatchProject *string) *AddEndpointParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the add endpoint params
func (o *AddEndpointParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithBody adds the body to the add endpoint params
func (o *AddEndpointParams) WithBody(body *v1.Endpoint) *AddEndpointParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add endpoint params
func (o *AddEndpointParams) SetBody(body *v1.Endpoint) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XDispatchOrg != nil {

		// header param X-Dispatch-Org
		if err := r.SetHeaderParam("X-Dispatch-Org", *o.XDispatchOrg); err != nil {
			return err
		}

	}

	if o.XDispatchProject != nil {

		// header param X-Dispatch-Project
		if err := r.SetHeaderParam("X-Dispatch-Project", *o.XDispatchProject); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}