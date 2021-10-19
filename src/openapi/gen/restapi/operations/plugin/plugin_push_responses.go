// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// PluginPushOKCode is the HTTP code returned for type PluginPushOK
const PluginPushOKCode int = 200

/*PluginPushOK no error

swagger:response pluginPushOK
*/
type PluginPushOK struct {
}

// NewPluginPushOK creates PluginPushOK with default headers values
func NewPluginPushOK() *PluginPushOK {

	return &PluginPushOK{}
}

// WriteResponse to the client
func (o *PluginPushOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PluginPushNotFoundCode is the HTTP code returned for type PluginPushNotFound
const PluginPushNotFoundCode int = 404

/*PluginPushNotFound plugin not installed

swagger:response pluginPushNotFound
*/
type PluginPushNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPluginPushNotFound creates PluginPushNotFound with default headers values
func NewPluginPushNotFound() *PluginPushNotFound {

	return &PluginPushNotFound{}
}

// WithPayload adds the payload to the plugin push not found response
func (o *PluginPushNotFound) WithPayload(payload *models.ErrorResponse) *PluginPushNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin push not found response
func (o *PluginPushNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginPushNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginPushInternalServerErrorCode is the HTTP code returned for type PluginPushInternalServerError
const PluginPushInternalServerErrorCode int = 500

/*PluginPushInternalServerError server error

swagger:response pluginPushInternalServerError
*/
type PluginPushInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPluginPushInternalServerError creates PluginPushInternalServerError with default headers values
func NewPluginPushInternalServerError() *PluginPushInternalServerError {

	return &PluginPushInternalServerError{}
}

// WithPayload adds the payload to the plugin push internal server error response
func (o *PluginPushInternalServerError) WithPayload(payload *models.ErrorResponse) *PluginPushInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin push internal server error response
func (o *PluginPushInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginPushInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}