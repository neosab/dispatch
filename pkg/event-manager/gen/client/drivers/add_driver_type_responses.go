///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// AddDriverTypeReader is a Reader for the AddDriverType structure.
type AddDriverTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDriverTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddDriverTypeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddDriverTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddDriverTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddDriverTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddDriverTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDriverTypeCreated creates a AddDriverTypeCreated with default headers values
func NewAddDriverTypeCreated() *AddDriverTypeCreated {
	return &AddDriverTypeCreated{}
}

/*AddDriverTypeCreated handles this case with default header values.

Driver Type created
*/
type AddDriverTypeCreated struct {
	Payload *models.DriverType
}

func (o *AddDriverTypeCreated) Error() string {
	return fmt.Sprintf("[POST /drivertypes][%d] addDriverTypeCreated  %+v", 201, o.Payload)
}

func (o *AddDriverTypeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriverType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverTypeBadRequest creates a AddDriverTypeBadRequest with default headers values
func NewAddDriverTypeBadRequest() *AddDriverTypeBadRequest {
	return &AddDriverTypeBadRequest{}
}

/*AddDriverTypeBadRequest handles this case with default header values.

Invalid input
*/
type AddDriverTypeBadRequest struct {
	Payload *models.Error
}

func (o *AddDriverTypeBadRequest) Error() string {
	return fmt.Sprintf("[POST /drivertypes][%d] addDriverTypeBadRequest  %+v", 400, o.Payload)
}

func (o *AddDriverTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverTypeUnauthorized creates a AddDriverTypeUnauthorized with default headers values
func NewAddDriverTypeUnauthorized() *AddDriverTypeUnauthorized {
	return &AddDriverTypeUnauthorized{}
}

/*AddDriverTypeUnauthorized handles this case with default header values.

Unauthorized Request
*/
type AddDriverTypeUnauthorized struct {
	Payload *models.Error
}

func (o *AddDriverTypeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /drivertypes][%d] addDriverTypeUnauthorized  %+v", 401, o.Payload)
}

func (o *AddDriverTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverTypeInternalServerError creates a AddDriverTypeInternalServerError with default headers values
func NewAddDriverTypeInternalServerError() *AddDriverTypeInternalServerError {
	return &AddDriverTypeInternalServerError{}
}

/*AddDriverTypeInternalServerError handles this case with default header values.

Internal server error
*/
type AddDriverTypeInternalServerError struct {
	Payload *models.Error
}

func (o *AddDriverTypeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /drivertypes][%d] addDriverTypeInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDriverTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverTypeDefault creates a AddDriverTypeDefault with default headers values
func NewAddDriverTypeDefault(code int) *AddDriverTypeDefault {
	return &AddDriverTypeDefault{
		_statusCode: code,
	}
}

/*AddDriverTypeDefault handles this case with default header values.

Unknown error
*/
type AddDriverTypeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add driver type default response
func (o *AddDriverTypeDefault) Code() int {
	return o._statusCode
}

func (o *AddDriverTypeDefault) Error() string {
	return fmt.Sprintf("[POST /drivertypes][%d] addDriverType default  %+v", o._statusCode, o.Payload)
}

func (o *AddDriverTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
