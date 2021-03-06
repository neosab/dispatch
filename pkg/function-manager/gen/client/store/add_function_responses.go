///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// AddFunctionReader is a Reader for the AddFunction structure.
type AddFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddFunctionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddFunctionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddFunctionOK creates a AddFunctionOK with default headers values
func NewAddFunctionOK() *AddFunctionOK {
	return &AddFunctionOK{}
}

/*AddFunctionOK handles this case with default header values.

Function created
*/
type AddFunctionOK struct {
	Payload *models.Function
}

func (o *AddFunctionOK) Error() string {
	return fmt.Sprintf("[POST /][%d] addFunctionOK  %+v", 200, o.Payload)
}

func (o *AddFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddFunctionBadRequest creates a AddFunctionBadRequest with default headers values
func NewAddFunctionBadRequest() *AddFunctionBadRequest {
	return &AddFunctionBadRequest{}
}

/*AddFunctionBadRequest handles this case with default header values.

Invalid input
*/
type AddFunctionBadRequest struct {
	Payload *models.Error
}

func (o *AddFunctionBadRequest) Error() string {
	return fmt.Sprintf("[POST /][%d] addFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *AddFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddFunctionUnauthorized creates a AddFunctionUnauthorized with default headers values
func NewAddFunctionUnauthorized() *AddFunctionUnauthorized {
	return &AddFunctionUnauthorized{}
}

/*AddFunctionUnauthorized handles this case with default header values.

Unauthorized Request
*/
type AddFunctionUnauthorized struct {
	Payload *models.Error
}

func (o *AddFunctionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /][%d] addFunctionUnauthorized  %+v", 401, o.Payload)
}

func (o *AddFunctionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddFunctionInternalServerError creates a AddFunctionInternalServerError with default headers values
func NewAddFunctionInternalServerError() *AddFunctionInternalServerError {
	return &AddFunctionInternalServerError{}
}

/*AddFunctionInternalServerError handles this case with default header values.

Internal error
*/
type AddFunctionInternalServerError struct {
	Payload *models.Error
}

func (o *AddFunctionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /][%d] addFunctionInternalServerError  %+v", 500, o.Payload)
}

func (o *AddFunctionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
