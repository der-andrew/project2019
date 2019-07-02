// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DocumentUpdateOKCode is the HTTP code returned for type DocumentUpdateOK
const DocumentUpdateOKCode int = 200

/*DocumentUpdateOK Справочник

swagger:response documentUpdateOK
*/
type DocumentUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *DocumentUpdateOKBody `json:"body,omitempty"`
}

// NewDocumentUpdateOK creates DocumentUpdateOK with default headers values
func NewDocumentUpdateOK() *DocumentUpdateOK {

	return &DocumentUpdateOK{}
}

// WithPayload adds the payload to the document update o k response
func (o *DocumentUpdateOK) WithPayload(payload *DocumentUpdateOKBody) *DocumentUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document update o k response
func (o *DocumentUpdateOK) SetPayload(payload *DocumentUpdateOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentUpdateBadRequestCode is the HTTP code returned for type DocumentUpdateBadRequest
const DocumentUpdateBadRequestCode int = 400

/*DocumentUpdateBadRequest Validation error

swagger:response documentUpdateBadRequest
*/
type DocumentUpdateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *DocumentUpdateBadRequestBody `json:"body,omitempty"`
}

// NewDocumentUpdateBadRequest creates DocumentUpdateBadRequest with default headers values
func NewDocumentUpdateBadRequest() *DocumentUpdateBadRequest {

	return &DocumentUpdateBadRequest{}
}

// WithPayload adds the payload to the document update bad request response
func (o *DocumentUpdateBadRequest) WithPayload(payload *DocumentUpdateBadRequestBody) *DocumentUpdateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document update bad request response
func (o *DocumentUpdateBadRequest) SetPayload(payload *DocumentUpdateBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentUpdateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentUpdateNotFoundCode is the HTTP code returned for type DocumentUpdateNotFound
const DocumentUpdateNotFoundCode int = 404

/*DocumentUpdateNotFound Not found

swagger:response documentUpdateNotFound
*/
type DocumentUpdateNotFound struct {

	/*
	  In: Body
	*/
	Payload *DocumentUpdateNotFoundBody `json:"body,omitempty"`
}

// NewDocumentUpdateNotFound creates DocumentUpdateNotFound with default headers values
func NewDocumentUpdateNotFound() *DocumentUpdateNotFound {

	return &DocumentUpdateNotFound{}
}

// WithPayload adds the payload to the document update not found response
func (o *DocumentUpdateNotFound) WithPayload(payload *DocumentUpdateNotFoundBody) *DocumentUpdateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document update not found response
func (o *DocumentUpdateNotFound) SetPayload(payload *DocumentUpdateNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentUpdateMethodNotAllowedCode is the HTTP code returned for type DocumentUpdateMethodNotAllowed
const DocumentUpdateMethodNotAllowedCode int = 405

/*DocumentUpdateMethodNotAllowed Invalid Method

swagger:response documentUpdateMethodNotAllowed
*/
type DocumentUpdateMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *DocumentUpdateMethodNotAllowedBody `json:"body,omitempty"`
}

// NewDocumentUpdateMethodNotAllowed creates DocumentUpdateMethodNotAllowed with default headers values
func NewDocumentUpdateMethodNotAllowed() *DocumentUpdateMethodNotAllowed {

	return &DocumentUpdateMethodNotAllowed{}
}

// WithPayload adds the payload to the document update method not allowed response
func (o *DocumentUpdateMethodNotAllowed) WithPayload(payload *DocumentUpdateMethodNotAllowedBody) *DocumentUpdateMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document update method not allowed response
func (o *DocumentUpdateMethodNotAllowed) SetPayload(payload *DocumentUpdateMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentUpdateMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentUpdateInternalServerErrorCode is the HTTP code returned for type DocumentUpdateInternalServerError
const DocumentUpdateInternalServerErrorCode int = 500

/*DocumentUpdateInternalServerError Internal server error

swagger:response documentUpdateInternalServerError
*/
type DocumentUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *DocumentUpdateInternalServerErrorBody `json:"body,omitempty"`
}

// NewDocumentUpdateInternalServerError creates DocumentUpdateInternalServerError with default headers values
func NewDocumentUpdateInternalServerError() *DocumentUpdateInternalServerError {

	return &DocumentUpdateInternalServerError{}
}

// WithPayload adds the payload to the document update internal server error response
func (o *DocumentUpdateInternalServerError) WithPayload(payload *DocumentUpdateInternalServerErrorBody) *DocumentUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document update internal server error response
func (o *DocumentUpdateInternalServerError) SetPayload(payload *DocumentUpdateInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}