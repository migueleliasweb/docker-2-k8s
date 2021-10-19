// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ServiceInspectOKCode is the HTTP code returned for type ServiceInspectOK
const ServiceInspectOKCode int = 200

/*ServiceInspectOK no error

swagger:response serviceInspectOK
*/
type ServiceInspectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Service `json:"body,omitempty"`
}

// NewServiceInspectOK creates ServiceInspectOK with default headers values
func NewServiceInspectOK() *ServiceInspectOK {

	return &ServiceInspectOK{}
}

// WithPayload adds the payload to the service inspect o k response
func (o *ServiceInspectOK) WithPayload(payload *models.Service) *ServiceInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service inspect o k response
func (o *ServiceInspectOK) SetPayload(payload *models.Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInspectNotFoundCode is the HTTP code returned for type ServiceInspectNotFound
const ServiceInspectNotFoundCode int = 404

/*ServiceInspectNotFound no such service

swagger:response serviceInspectNotFound
*/
type ServiceInspectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewServiceInspectNotFound creates ServiceInspectNotFound with default headers values
func NewServiceInspectNotFound() *ServiceInspectNotFound {

	return &ServiceInspectNotFound{}
}

// WithPayload adds the payload to the service inspect not found response
func (o *ServiceInspectNotFound) WithPayload(payload *models.ErrorResponse) *ServiceInspectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service inspect not found response
func (o *ServiceInspectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInspectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInspectInternalServerErrorCode is the HTTP code returned for type ServiceInspectInternalServerError
const ServiceInspectInternalServerErrorCode int = 500

/*ServiceInspectInternalServerError server error

swagger:response serviceInspectInternalServerError
*/
type ServiceInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewServiceInspectInternalServerError creates ServiceInspectInternalServerError with default headers values
func NewServiceInspectInternalServerError() *ServiceInspectInternalServerError {

	return &ServiceInspectInternalServerError{}
}

// WithPayload adds the payload to the service inspect internal server error response
func (o *ServiceInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *ServiceInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service inspect internal server error response
func (o *ServiceInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInspectServiceUnavailableCode is the HTTP code returned for type ServiceInspectServiceUnavailable
const ServiceInspectServiceUnavailableCode int = 503

/*ServiceInspectServiceUnavailable node is not part of a swarm

swagger:response serviceInspectServiceUnavailable
*/
type ServiceInspectServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewServiceInspectServiceUnavailable creates ServiceInspectServiceUnavailable with default headers values
func NewServiceInspectServiceUnavailable() *ServiceInspectServiceUnavailable {

	return &ServiceInspectServiceUnavailable{}
}

// WithPayload adds the payload to the service inspect service unavailable response
func (o *ServiceInspectServiceUnavailable) WithPayload(payload *models.ErrorResponse) *ServiceInspectServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service inspect service unavailable response
func (o *ServiceInspectServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInspectServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}