// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DocumentParam Locale_param
// swagger:model Document_param
type DocumentParam struct {

	// Идентификатор записи
	// Required: true
	Code *string `json:"code"`

	// Локаль справочника ISO 639-2
	// Required: true
	Locale *string `json:"locale"`

	// Содержание записи
	// Required: true
	Text *string `json:"text"`

	// Тип справочника
	// Required: true
	Type *string `json:"type"`

	// document param additional properties
	DocumentParamAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *DocumentParam) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Идентификатор записи
		// Required: true
		Code *string `json:"code"`

		// Локаль справочника ISO 639-2
		// Required: true
		Locale *string `json:"locale"`

		// Содержание записи
		// Required: true
		Text *string `json:"text"`

		// Тип справочника
		// Required: true
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DocumentParam

	rcv.Code = stage1.Code

	rcv.Locale = stage1.Locale

	rcv.Text = stage1.Text

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "code")

	delete(stage2, "locale")

	delete(stage2, "text")

	delete(stage2, "type")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.DocumentParamAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m DocumentParam) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Идентификатор записи
		// Required: true
		Code *string `json:"code"`

		// Локаль справочника ISO 639-2
		// Required: true
		Locale *string `json:"locale"`

		// Содержание записи
		// Required: true
		Text *string `json:"text"`

		// Тип справочника
		// Required: true
		Type *string `json:"type"`
	}

	stage1.Code = m.Code

	stage1.Locale = m.Locale

	stage1.Text = m.Text

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.DocumentParamAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.DocumentParamAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this document param
func (m *DocumentParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocumentParam) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *DocumentParam) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *DocumentParam) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *DocumentParam) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocumentParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentParam) UnmarshalBinary(b []byte) error {
	var res DocumentParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}