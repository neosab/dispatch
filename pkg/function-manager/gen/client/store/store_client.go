///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new store API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for store API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddFunction adds a new function
*/
func (a *Client) AddFunction(params *AddFunctionParams, authInfo runtime.ClientAuthInfoWriter) (*AddFunctionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddFunctionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addFunction",
		Method:             "POST",
		PathPattern:        "/function",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddFunctionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddFunctionOK), nil

}

/*
DeleteFunction deletes a function
*/
func (a *Client) DeleteFunction(params *DeleteFunctionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFunctionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFunctionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFunction",
		Method:             "DELETE",
		PathPattern:        "/function/{functionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFunctionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFunctionOK), nil

}

/*
GetFunction finds function by name

Returns a single function
*/
func (a *Client) GetFunction(params *GetFunctionParams, authInfo runtime.ClientAuthInfoWriter) (*GetFunctionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFunctionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFunction",
		Method:             "GET",
		PathPattern:        "/function/{functionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFunctionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFunctionOK), nil

}

/*
GetFunctions lists all existing functions
*/
func (a *Client) GetFunctions(params *GetFunctionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFunctionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFunctionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFunctions",
		Method:             "GET",
		PathPattern:        "/function",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFunctionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFunctionsOK), nil

}

/*
UpdateFunction updates a function
*/
func (a *Client) UpdateFunction(params *UpdateFunctionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFunctionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFunctionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFunction",
		Method:             "PUT",
		PathPattern:        "/function/{functionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateFunctionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateFunctionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
