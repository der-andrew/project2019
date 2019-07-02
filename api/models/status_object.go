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

// StatusObject Status_object
// swagger:model Status_object
type StatusObject struct {

	// Коллекция доступных локализаций ISO 639-2
	Locales []string `json:"locales"`

	// Коллекция справочника
	Name string `json:"name,omitempty"`

	// Тип справочника
	// Enum: [STATIC DYNAMIC]
	Type string `json:"type,omitempty"`

	// ISO8601 Дата обновления спрвочника
	Updated string `json:"updated,omitempty"`
}

// Validate validates this status object
func (m *StatusObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statusObjectTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATIC","DYNAMIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusObjectTypeTypePropEnum = append(statusObjectTypeTypePropEnum, v)
	}
}

const (

	// StatusObjectTypeSTATIC captures enum value "STATIC"
	StatusObjectTypeSTATIC string = "STATIC"

	// StatusObjectTypeDYNAMIC captures enum value "DYNAMIC"
	StatusObjectTypeDYNAMIC string = "DYNAMIC"
)

// prop value enum
func (m *StatusObject) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, statusObjectTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StatusObject) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusObject) UnmarshalBinary(b []byte) error {
	var res StatusObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}