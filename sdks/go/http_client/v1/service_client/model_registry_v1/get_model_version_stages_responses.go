// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package model_registry_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetModelVersionStagesReader is a Reader for the GetModelVersionStages structure.
type GetModelVersionStagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelVersionStagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModelVersionStagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetModelVersionStagesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetModelVersionStagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetModelVersionStagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetModelVersionStagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetModelVersionStagesOK creates a GetModelVersionStagesOK with default headers values
func NewGetModelVersionStagesOK() *GetModelVersionStagesOK {
	return &GetModelVersionStagesOK{}
}

/* GetModelVersionStagesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetModelVersionStagesOK struct {
	Payload *service_model.V1Stage
}

func (o *GetModelVersionStagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}/stages][%d] getModelVersionStagesOK  %+v", 200, o.Payload)
}
func (o *GetModelVersionStagesOK) GetPayload() *service_model.V1Stage {
	return o.Payload
}

func (o *GetModelVersionStagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Stage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionStagesNoContent creates a GetModelVersionStagesNoContent with default headers values
func NewGetModelVersionStagesNoContent() *GetModelVersionStagesNoContent {
	return &GetModelVersionStagesNoContent{}
}

/* GetModelVersionStagesNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetModelVersionStagesNoContent struct {
	Payload interface{}
}

func (o *GetModelVersionStagesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}/stages][%d] getModelVersionStagesNoContent  %+v", 204, o.Payload)
}
func (o *GetModelVersionStagesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelVersionStagesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionStagesForbidden creates a GetModelVersionStagesForbidden with default headers values
func NewGetModelVersionStagesForbidden() *GetModelVersionStagesForbidden {
	return &GetModelVersionStagesForbidden{}
}

/* GetModelVersionStagesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetModelVersionStagesForbidden struct {
	Payload interface{}
}

func (o *GetModelVersionStagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}/stages][%d] getModelVersionStagesForbidden  %+v", 403, o.Payload)
}
func (o *GetModelVersionStagesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelVersionStagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionStagesNotFound creates a GetModelVersionStagesNotFound with default headers values
func NewGetModelVersionStagesNotFound() *GetModelVersionStagesNotFound {
	return &GetModelVersionStagesNotFound{}
}

/* GetModelVersionStagesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetModelVersionStagesNotFound struct {
	Payload interface{}
}

func (o *GetModelVersionStagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}/stages][%d] getModelVersionStagesNotFound  %+v", 404, o.Payload)
}
func (o *GetModelVersionStagesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelVersionStagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelVersionStagesDefault creates a GetModelVersionStagesDefault with default headers values
func NewGetModelVersionStagesDefault(code int) *GetModelVersionStagesDefault {
	return &GetModelVersionStagesDefault{
		_statusCode: code,
	}
}

/* GetModelVersionStagesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetModelVersionStagesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get model version stages default response
func (o *GetModelVersionStagesDefault) Code() int {
	return o._statusCode
}

func (o *GetModelVersionStagesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/registry/{entity}/versions/{name}/stages][%d] GetModelVersionStages default  %+v", o._statusCode, o.Payload)
}
func (o *GetModelVersionStagesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetModelVersionStagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
