///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetOrganizationHandlerFunc turns a function with the right signature into a get organization handler
type GetOrganizationHandlerFunc func(GetOrganizationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOrganizationHandlerFunc) Handle(params GetOrganizationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetOrganizationHandler interface for that can handle valid get organization params
type GetOrganizationHandler interface {
	Handle(GetOrganizationParams, interface{}) middleware.Responder
}

// NewGetOrganization creates a new http.Handler for the get organization operation
func NewGetOrganization(ctx *middleware.Context, handler GetOrganizationHandler) *GetOrganization {
	return &GetOrganization{Context: ctx, Handler: handler}
}

/*GetOrganization swagger:route GET /v1/iam/organization/{organizationName} organization getOrganization

Find Organization by name

get an Organization by name

*/
type GetOrganization struct {
	Context *middleware.Context
	Handler GetOrganizationHandler
}

func (o *GetOrganization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOrganizationParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
