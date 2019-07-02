// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DocumentDeleteOKCode is the HTTP code returned for type DocumentDeleteOK
const DocumentDeleteOKCode int = 200

/*DocumentDeleteOK SUCCESS

swagger:response documentDeleteOK
*/
type DocumentDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *DocumentDeleteOKBody `json:"body,omitempty"`
}

// NewDocumentDeleteOK creates DocumentDeleteOK with default headers values
func NewDocumentDeleteOK() *DocumentDeleteOK {

	return &DocumentDeleteOK{}
}

// WithPayload adds the payload to the document delete o k response
func (o *DocumentDeleteOK) WithPayload(payload *DocumentDeleteOKBody) *DocumentDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document delete o k response
func (o *DocumentDeleteOK) SetPayload(payload *DocumentDeleteOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentDeleteNotFoundCode is the HTTP code returned for type DocumentDeleteNotFound
const DocumentDeleteNotFoundCode int = 404

/*DocumentDeleteNotFound Not found

swagger:response documentDeleteNotFound
*/
type DocumentDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *DocumentDeleteNotFoundBody `json:"body,omitempty"`
}

// NewDocumentDeleteNotFound creates DocumentDeleteNotFound with default headers values
func NewDocumentDeleteNotFound() *DocumentDeleteNotFound {

	return &DocumentDeleteNotFound{}
}

// WithPayload adds the payload to the document delete not found response
func (o *DocumentDeleteNotFound) WithPayload(payload *DocumentDeleteNotFoundBody) *DocumentDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document delete not found response
func (o *DocumentDeleteNotFound) SetPayload(payload *DocumentDeleteNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentDeleteMethodNotAllowedCode is the HTTP code returned for type DocumentDeleteMethodNotAllowed
const DocumentDeleteMethodNotAllowedCode int = 405

/*DocumentDeleteMethodNotAllowed Invalid Method

swagger:response documentDeleteMethodNotAllowed
*/
type DocumentDeleteMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *DocumentDeleteMethodNotAllowedBody `json:"body,omitempty"`
}

// NewDocumentDeleteMethodNotAllowed creates DocumentDeleteMethodNotAllowed with default headers values
func NewDocumentDeleteMethodNotAllowed() *DocumentDeleteMethodNotAllowed {

	return &DocumentDeleteMethodNotAllowed{}
}

// WithPayload adds the payload to the document delete method not allowed response
func (o *DocumentDeleteMethodNotAllowed) WithPayload(payload *DocumentDeleteMethodNotAllowedBody) *DocumentDeleteMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document delete method not allowed response
func (o *DocumentDeleteMethodNotAllowed) SetPayload(payload *DocumentDeleteMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentDeleteMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentDeleteInternalServerErrorCode is the HTTP code returned for type DocumentDeleteInternalServerError
const DocumentDeleteInternalServerErrorCode int = 500

/*DocumentDeleteInternalServerError Internal server error

swagger:response documentDeleteInternalServerError
*/
type DocumentDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *DocumentDeleteInternalServerErrorBody `json:"body,omitempty"`
}

// NewDocumentDeleteInternalServerError creates DocumentDeleteInternalServerError with default headers values
func NewDocumentDeleteInternalServerError() *DocumentDeleteInternalServerError {

	return &DocumentDeleteInternalServerError{}
}

// WithPayload adds the payload to the document delete internal server error response
func (o *DocumentDeleteInternalServerError) WithPayload(payload *DocumentDeleteInternalServerErrorBody) *DocumentDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document delete internal server error response
func (o *DocumentDeleteInternalServerError) SetPayload(payload *DocumentDeleteInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}