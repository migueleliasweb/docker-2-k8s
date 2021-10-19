// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// NetworkDisconnectOKCode is the HTTP code returned for type NetworkDisconnectOK
const NetworkDisconnectOKCode int = 200

/*NetworkDisconnectOK No error

swagger:response networkDisconnectOK
*/
type NetworkDisconnectOK struct {
}

// NewNetworkDisconnectOK creates NetworkDisconnectOK with default headers values
func NewNetworkDisconnectOK() *NetworkDisconnectOK {

	return &NetworkDisconnectOK{}
}

// WriteResponse to the client
func (o *NetworkDisconnectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// NetworkDisconnectForbiddenCode is the HTTP code returned for type NetworkDisconnectForbidden
const NetworkDisconnectForbiddenCode int = 403

/*NetworkDisconnectForbidden Operation not supported for swarm scoped networks

swagger:response networkDisconnectForbidden
*/
type NetworkDisconnectForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkDisconnectForbidden creates NetworkDisconnectForbidden with default headers values
func NewNetworkDisconnectForbidden() *NetworkDisconnectForbidden {

	return &NetworkDisconnectForbidden{}
}

// WithPayload adds the payload to the network disconnect forbidden response
func (o *NetworkDisconnectForbidden) WithPayload(payload *models.ErrorResponse) *NetworkDisconnectForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network disconnect forbidden response
func (o *NetworkDisconnectForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkDisconnectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkDisconnectNotFoundCode is the HTTP code returned for type NetworkDisconnectNotFound
const NetworkDisconnectNotFoundCode int = 404

/*NetworkDisconnectNotFound Network or container not found

swagger:response networkDisconnectNotFound
*/
type NetworkDisconnectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkDisconnectNotFound creates NetworkDisconnectNotFound with default headers values
func NewNetworkDisconnectNotFound() *NetworkDisconnectNotFound {

	return &NetworkDisconnectNotFound{}
}

// WithPayload adds the payload to the network disconnect not found response
func (o *NetworkDisconnectNotFound) WithPayload(payload *models.ErrorResponse) *NetworkDisconnectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network disconnect not found response
func (o *NetworkDisconnectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkDisconnectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkDisconnectInternalServerErrorCode is the HTTP code returned for type NetworkDisconnectInternalServerError
const NetworkDisconnectInternalServerErrorCode int = 500

/*NetworkDisconnectInternalServerError Server error

swagger:response networkDisconnectInternalServerError
*/
type NetworkDisconnectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkDisconnectInternalServerError creates NetworkDisconnectInternalServerError with default headers values
func NewNetworkDisconnectInternalServerError() *NetworkDisconnectInternalServerError {

	return &NetworkDisconnectInternalServerError{}
}

// WithPayload adds the payload to the network disconnect internal server error response
func (o *NetworkDisconnectInternalServerError) WithPayload(payload *models.ErrorResponse) *NetworkDisconnectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network disconnect internal server error response
func (o *NetworkDisconnectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkDisconnectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
