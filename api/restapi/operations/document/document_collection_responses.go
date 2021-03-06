// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DocumentCollectionOKCode is the HTTP code returned for type DocumentCollectionOK
const DocumentCollectionOKCode int = 200

/*DocumentCollectionOK Справочник

swagger:response documentCollectionOK
*/
type DocumentCollectionOK struct {

	/*
	  In: Body
	*/
	Payload *DocumentCollectionOKBody `json:"body,omitempty"`
}

// NewDocumentCollectionOK creates DocumentCollectionOK with default headers values
func NewDocumentCollectionOK() *DocumentCollectionOK {

	return &DocumentCollectionOK{}
}

// WithPayload adds the payload to the document collection o k response
func (o *DocumentCollectionOK) WithPayload(payload *DocumentCollectionOKBody) *DocumentCollectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document collection o k response
func (o *DocumentCollectionOK) SetPayload(payload *DocumentCollectionOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentCollectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentCollectionNotFoundCode is the HTTP code returned for type DocumentCollectionNotFound
const DocumentCollectionNotFoundCode int = 404

/*DocumentCollectionNotFound Not found

swagger:response documentCollectionNotFound
*/
type DocumentCollectionNotFound struct {

	/*
	  In: Body
	*/
	Payload *DocumentCollectionNotFoundBody `json:"body,omitempty"`
}

// NewDocumentCollectionNotFound creates DocumentCollectionNotFound with default headers values
func NewDocumentCollectionNotFound() *DocumentCollectionNotFound {

	return &DocumentCollectionNotFound{}
}

// WithPayload adds the payload to the document collection not found response
func (o *DocumentCollectionNotFound) WithPayload(payload *DocumentCollectionNotFoundBody) *DocumentCollectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document collection not found response
func (o *DocumentCollectionNotFound) SetPayload(payload *DocumentCollectionNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentCollectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentCollectionMethodNotAllowedCode is the HTTP code returned for type DocumentCollectionMethodNotAllowed
const DocumentCollectionMethodNotAllowedCode int = 405

/*DocumentCollectionMethodNotAllowed Invalid Method

swagger:response documentCollectionMethodNotAllowed
*/
type DocumentCollectionMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *DocumentCollectionMethodNotAllowedBody `json:"body,omitempty"`
}

// NewDocumentCollectionMethodNotAllowed creates DocumentCollectionMethodNotAllowed with default headers values
func NewDocumentCollectionMethodNotAllowed() *DocumentCollectionMethodNotAllowed {

	return &DocumentCollectionMethodNotAllowed{}
}

// WithPayload adds the payload to the document collection method not allowed response
func (o *DocumentCollectionMethodNotAllowed) WithPayload(payload *DocumentCollectionMethodNotAllowedBody) *DocumentCollectionMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document collection method not allowed response
func (o *DocumentCollectionMethodNotAllowed) SetPayload(payload *DocumentCollectionMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentCollectionMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DocumentCollectionInternalServerErrorCode is the HTTP code returned for type DocumentCollectionInternalServerError
const DocumentCollectionInternalServerErrorCode int = 500

/*DocumentCollectionInternalServerError Internal server error

swagger:response documentCollectionInternalServerError
*/
type DocumentCollectionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *DocumentCollectionInternalServerErrorBody `json:"body,omitempty"`
}

// NewDocumentCollectionInternalServerError creates DocumentCollectionInternalServerError with default headers values
func NewDocumentCollectionInternalServerError() *DocumentCollectionInternalServerError {

	return &DocumentCollectionInternalServerError{}
}

// WithPayload adds the payload to the document collection internal server error response
func (o *DocumentCollectionInternalServerError) WithPayload(payload *DocumentCollectionInternalServerErrorBody) *DocumentCollectionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the document collection internal server error response
func (o *DocumentCollectionInternalServerError) SetPayload(payload *DocumentCollectionInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DocumentCollectionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
