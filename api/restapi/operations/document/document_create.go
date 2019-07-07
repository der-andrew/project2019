// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "gitlab.com/project2019-02/thesaurus/api/models"
)

// DocumentCreateHandlerFunc turns a function with the right signature into a document create handler
type DocumentCreateHandlerFunc func(DocumentCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DocumentCreateHandlerFunc) Handle(params DocumentCreateParams) middleware.Responder {
	return fn(params)
}

// DocumentCreateHandler interface for that can handle valid document create params
type DocumentCreateHandler interface {
	Handle(DocumentCreateParams) middleware.Responder
}

// NewDocumentCreate creates a new http.Handler for the document create operation
func NewDocumentCreate(ctx *middleware.Context, handler DocumentCreateHandler) *DocumentCreate {
	return &DocumentCreate{Context: ctx, Handler: handler}
}

/*DocumentCreate swagger:route POST /documents Document documentCreate

Создание документа

*/
type DocumentCreate struct {
	Context *middleware.Context
	Handler DocumentCreateHandler
}

func (o *DocumentCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDocumentCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DocumentCreateBadRequestBody document create bad request body
// swagger:model DocumentCreateBadRequestBody
type DocumentCreateBadRequestBody struct {
	models.Document400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentCreateBadRequestBody) UnmarshalJSON(raw []byte) error {
	// DocumentCreateBadRequestBodyAO0
	var documentCreateBadRequestBodyAO0 models.Document400Data
	if err := swag.ReadJSON(raw, &documentCreateBadRequestBodyAO0); err != nil {
		return err
	}
	o.Document400Data = documentCreateBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentCreateBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentCreateBadRequestBodyAO0, err := swag.WriteJSON(o.Document400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentCreateBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document create bad request body
func (o *DocumentCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Document400Data
	if err := o.Document400Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DocumentCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DocumentCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentCreateInternalServerErrorBody document create internal server error body
// swagger:model DocumentCreateInternalServerErrorBody
type DocumentCreateInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentCreateInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// DocumentCreateInternalServerErrorBodyAO0
	var documentCreateInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &documentCreateInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = documentCreateInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentCreateInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentCreateInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentCreateInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document create internal server error body
func (o *DocumentCreateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error500Data
	if err := o.Error500Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DocumentCreateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentCreateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DocumentCreateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentCreateMethodNotAllowedBody document create method not allowed body
// swagger:model DocumentCreateMethodNotAllowedBody
type DocumentCreateMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentCreateMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// DocumentCreateMethodNotAllowedBodyAO0
	var documentCreateMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &documentCreateMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = documentCreateMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentCreateMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentCreateMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentCreateMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document create method not allowed body
func (o *DocumentCreateMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error405Data
	if err := o.Error405Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DocumentCreateMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentCreateMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res DocumentCreateMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentCreateNotFoundBody document create not found body
// swagger:model DocumentCreateNotFoundBody
type DocumentCreateNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentCreateNotFoundBody) UnmarshalJSON(raw []byte) error {
	// DocumentCreateNotFoundBodyAO0
	var documentCreateNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &documentCreateNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = documentCreateNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentCreateNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentCreateNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentCreateNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document create not found body
func (o *DocumentCreateNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error404Data
	if err := o.Error404Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DocumentCreateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentCreateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DocumentCreateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentCreateOKBody document create o k body
// swagger:model DocumentCreateOKBody
type DocumentCreateOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentCreateOKBody) UnmarshalJSON(raw []byte) error {
	// DocumentCreateOKBodyAO0
	var documentCreateOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &documentCreateOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = documentCreateOKBodyAO0

	// DocumentCreateOKBodyAO1
	var dataDocumentCreateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataDocumentCreateOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataDocumentCreateOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentCreateOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	documentCreateOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentCreateOKBodyAO0)

	var dataDocumentCreateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataDocumentCreateOKBodyAO1.Data = o.Data

	jsonDataDocumentCreateOKBodyAO1, errDocumentCreateOKBodyAO1 := swag.WriteJSON(dataDocumentCreateOKBodyAO1)
	if errDocumentCreateOKBodyAO1 != nil {
		return nil, errDocumentCreateOKBodyAO1
	}
	_parts = append(_parts, jsonDataDocumentCreateOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document create o k body
func (o *DocumentCreateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SuccessData
	if err := o.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DocumentCreateOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("documentCreateOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documentCreateOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DocumentCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentCreateOKBody) UnmarshalBinary(b []byte) error {
	var res DocumentCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
