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

	models "repo.nefrosovet.ru/maximus-platform/thesaurus/api/models"
)

// DocumentUpdateHandlerFunc turns a function with the right signature into a document update handler
type DocumentUpdateHandlerFunc func(DocumentUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DocumentUpdateHandlerFunc) Handle(params DocumentUpdateParams) middleware.Responder {
	return fn(params)
}

// DocumentUpdateHandler interface for that can handle valid document update params
type DocumentUpdateHandler interface {
	Handle(DocumentUpdateParams) middleware.Responder
}

// NewDocumentUpdate creates a new http.Handler for the document update operation
func NewDocumentUpdate(ctx *middleware.Context, handler DocumentUpdateHandler) *DocumentUpdate {
	return &DocumentUpdate{Context: ctx, Handler: handler}
}

/*DocumentUpdate swagger:route PUT /documents Document documentUpdate

Редактирование документа

*/
type DocumentUpdate struct {
	Context *middleware.Context
	Handler DocumentUpdateHandler
}

func (o *DocumentUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDocumentUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DocumentUpdateBadRequestBody document update bad request body
// swagger:model DocumentUpdateBadRequestBody
type DocumentUpdateBadRequestBody struct {
	models.Document400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentUpdateBadRequestBody) UnmarshalJSON(raw []byte) error {
	// DocumentUpdateBadRequestBodyAO0
	var documentUpdateBadRequestBodyAO0 models.Document400Data
	if err := swag.ReadJSON(raw, &documentUpdateBadRequestBodyAO0); err != nil {
		return err
	}
	o.Document400Data = documentUpdateBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentUpdateBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentUpdateBadRequestBodyAO0, err := swag.WriteJSON(o.Document400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentUpdateBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document update bad request body
func (o *DocumentUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
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
func (o *DocumentUpdateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentUpdateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DocumentUpdateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentUpdateInternalServerErrorBody document update internal server error body
// swagger:model DocumentUpdateInternalServerErrorBody
type DocumentUpdateInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentUpdateInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// DocumentUpdateInternalServerErrorBodyAO0
	var documentUpdateInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &documentUpdateInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = documentUpdateInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentUpdateInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentUpdateInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentUpdateInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document update internal server error body
func (o *DocumentUpdateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *DocumentUpdateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentUpdateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DocumentUpdateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentUpdateMethodNotAllowedBody document update method not allowed body
// swagger:model DocumentUpdateMethodNotAllowedBody
type DocumentUpdateMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentUpdateMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// DocumentUpdateMethodNotAllowedBodyAO0
	var documentUpdateMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &documentUpdateMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = documentUpdateMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentUpdateMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentUpdateMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentUpdateMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document update method not allowed body
func (o *DocumentUpdateMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *DocumentUpdateMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentUpdateMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res DocumentUpdateMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentUpdateNotFoundBody document update not found body
// swagger:model DocumentUpdateNotFoundBody
type DocumentUpdateNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentUpdateNotFoundBody) UnmarshalJSON(raw []byte) error {
	// DocumentUpdateNotFoundBodyAO0
	var documentUpdateNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &documentUpdateNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = documentUpdateNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentUpdateNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	documentUpdateNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentUpdateNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document update not found body
func (o *DocumentUpdateNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *DocumentUpdateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentUpdateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DocumentUpdateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DocumentUpdateOKBody document update o k body
// swagger:model DocumentUpdateOKBody
type DocumentUpdateOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DocumentUpdateOKBody) UnmarshalJSON(raw []byte) error {
	// DocumentUpdateOKBodyAO0
	var documentUpdateOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &documentUpdateOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = documentUpdateOKBodyAO0

	// DocumentUpdateOKBodyAO1
	var dataDocumentUpdateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataDocumentUpdateOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataDocumentUpdateOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DocumentUpdateOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	documentUpdateOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, documentUpdateOKBodyAO0)

	var dataDocumentUpdateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataDocumentUpdateOKBodyAO1.Data = o.Data

	jsonDataDocumentUpdateOKBodyAO1, errDocumentUpdateOKBodyAO1 := swag.WriteJSON(dataDocumentUpdateOKBodyAO1)
	if errDocumentUpdateOKBodyAO1 != nil {
		return nil, errDocumentUpdateOKBodyAO1
	}
	_parts = append(_parts, jsonDataDocumentUpdateOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this document update o k body
func (o *DocumentUpdateOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DocumentUpdateOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("documentUpdateOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documentUpdateOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DocumentUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocumentUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res DocumentUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}