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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Hook Hook definition
//
// swagger:model v1Hook
type V1Hook struct {

	// An optional template containing conditions to check before starting the run
	Conditions string `json:"conditions,omitempty"`

	// Optional connection
	Connection string `json:"connection,omitempty"`

	// Optional preset to use for running this operation
	DisableDefaults bool `json:"disableDefaults,omitempty"`

	// Component name to use
	HubRef string `json:"hub_ref,omitempty"`

	// Optional dict of params
	Params map[string]V1Param `json:"params,omitempty"`

	// Optional preset to use for running this operation
	Presets []string `json:"presets"`

	// Optional trigger policy, default done
	Trigger V1Statuses `json:"trigger,omitempty"`
}

// Validate validates this v1 hook
func (m *V1Hook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Hook) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(m.Params) { // not required
		return nil
	}

	for k := range m.Params {

		if err := validate.Required("params"+"."+k, "body", m.Params[k]); err != nil {
			return err
		}
		if val, ok := m.Params[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V1Hook) validateTrigger(formats strfmt.Registry) error {

	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if err := m.Trigger.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trigger")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Hook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Hook) UnmarshalBinary(b []byte) error {
	var res V1Hook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
