// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDocumentCollectionParams creates a new DocumentCollectionParams object
// with the default values initialized.
func NewDocumentCollectionParams() DocumentCollectionParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(0)

		offsetDefault = int64(0)
	)

	return DocumentCollectionParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// DocumentCollectionParams contains all the bound params for the document collection operation
// typically these are obtained from a http.Request
//
// swagger:parameters documentCollection
type DocumentCollectionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентфикатор документа
	  In: query
	*/
	Code *string
	/*Лимит
	  In: query
	  Default: 0
	*/
	Limit *int64
	/*Локаль справочника
	  Required: true
	  In: query
	*/
	Locale string
	/*Шаг
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*Категория документов
	  Required: true
	  In: query
	*/
	Type string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDocumentCollectionParams() beforehand.
func (o *DocumentCollectionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCode, qhkCode, _ := qs.GetOK("code")
	if err := o.bindCode(qCode, qhkCode, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qLocale, qhkLocale, _ := qs.GetOK("locale")
	if err := o.bindLocale(qLocale, qhkLocale, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qType, qhkType, _ := qs.GetOK("type")
	if err := o.bindType(qType, qhkType, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCode binds and validates parameter Code from query.
func (o *DocumentCollectionParams) bindCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Code = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *DocumentCollectionParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewDocumentCollectionParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindLocale binds and validates parameter Locale from query.
func (o *DocumentCollectionParams) bindLocale(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("locale", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("locale", "query", raw); err != nil {
		return err
	}

	o.Locale = raw

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *DocumentCollectionParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewDocumentCollectionParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	return nil
}

// bindType binds and validates parameter Type from query.
func (o *DocumentCollectionParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("type", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("type", "query", raw); err != nil {
		return err
	}

	o.Type = raw

	return nil
}
