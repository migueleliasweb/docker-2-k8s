// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ConfigCreateCreatedCode is the HTTP code returned for type ConfigCreateCreated
const ConfigCreateCreatedCode int = 201

/*ConfigCreateCreated no error

swagger:response configCreateCreated
*/
type ConfigCreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IDResponse `json:"body,omitempty"`
}

// NewConfigCreateCreated creates ConfigCreateCreated with default headers values
func NewConfigCreateCreated() *ConfigCreateCreated {

	return &ConfigCreateCreated{}
}

// WithPayload adds the payload to the config create created response
func (o *ConfigCreateCreated) WithPayload(payload *models.IDResponse) *ConfigCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config create created response
func (o *ConfigCreateCreated) SetPayload(payload *models.IDResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConfigCreateConflictCode is the HTTP code returned for type ConfigCreateConflict
const ConfigCreateConflictCode int = 409

/*ConfigCreateConflict name conflicts with an existing object

swagger:response configCreateConflict
*/
type ConfigCreateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewConfigCreateConflict creates ConfigCreateConflict with default headers values
func NewConfigCreateConflict() *ConfigCreateConflict {

	return &ConfigCreateConflict{}
}

// WithPayload adds the payload to the config create conflict response
func (o *ConfigCreateConflict) WithPayload(payload *models.ErrorResponse) *ConfigCreateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config create conflict response
func (o *ConfigCreateConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigCreateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConfigCreateInternalServerErrorCode is the HTTP code returned for type ConfigCreateInternalServerError
const ConfigCreateInternalServerErrorCode int = 500

/*ConfigCreateInternalServerError server error

swagger:response configCreateInternalServerError
*/
type ConfigCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewConfigCreateInternalServerError creates ConfigCreateInternalServerError with default headers values
func NewConfigCreateInternalServerError() *ConfigCreateInternalServerError {

	return &ConfigCreateInternalServerError{}
}

// WithPayload adds the payload to the config create internal server error response
func (o *ConfigCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *ConfigCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config create internal server error response
func (o *ConfigCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConfigCreateServiceUnavailableCode is the HTTP code returned for type ConfigCreateServiceUnavailable
const ConfigCreateServiceUnavailableCode int = 503

/*ConfigCreateServiceUnavailable node is not part of a swarm

swagger:response configCreateServiceUnavailable
*/
type ConfigCreateServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewConfigCreateServiceUnavailable creates ConfigCreateServiceUnavailable with default headers values
func NewConfigCreateServiceUnavailable() *ConfigCreateServiceUnavailable {

	return &ConfigCreateServiceUnavailable{}
}

// WithPayload adds the payload to the config create service unavailable response
func (o *ConfigCreateServiceUnavailable) WithPayload(payload *models.ErrorResponse) *ConfigCreateServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config create service unavailable response
func (o *ConfigCreateServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigCreateServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
