// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ContainerRenameNoContentCode is the HTTP code returned for type ContainerRenameNoContent
const ContainerRenameNoContentCode int = 204

/*ContainerRenameNoContent no error

swagger:response containerRenameNoContent
*/
type ContainerRenameNoContent struct {
}

// NewContainerRenameNoContent creates ContainerRenameNoContent with default headers values
func NewContainerRenameNoContent() *ContainerRenameNoContent {

	return &ContainerRenameNoContent{}
}

// WriteResponse to the client
func (o *ContainerRenameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ContainerRenameNotFoundCode is the HTTP code returned for type ContainerRenameNotFound
const ContainerRenameNotFoundCode int = 404

/*ContainerRenameNotFound no such container

swagger:response containerRenameNotFound
*/
type ContainerRenameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerRenameNotFound creates ContainerRenameNotFound with default headers values
func NewContainerRenameNotFound() *ContainerRenameNotFound {

	return &ContainerRenameNotFound{}
}

// WithPayload adds the payload to the container rename not found response
func (o *ContainerRenameNotFound) WithPayload(payload *models.ErrorResponse) *ContainerRenameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container rename not found response
func (o *ContainerRenameNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerRenameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerRenameConflictCode is the HTTP code returned for type ContainerRenameConflict
const ContainerRenameConflictCode int = 409

/*ContainerRenameConflict name already in use

swagger:response containerRenameConflict
*/
type ContainerRenameConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerRenameConflict creates ContainerRenameConflict with default headers values
func NewContainerRenameConflict() *ContainerRenameConflict {

	return &ContainerRenameConflict{}
}

// WithPayload adds the payload to the container rename conflict response
func (o *ContainerRenameConflict) WithPayload(payload *models.ErrorResponse) *ContainerRenameConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container rename conflict response
func (o *ContainerRenameConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerRenameConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerRenameInternalServerErrorCode is the HTTP code returned for type ContainerRenameInternalServerError
const ContainerRenameInternalServerErrorCode int = 500

/*ContainerRenameInternalServerError server error

swagger:response containerRenameInternalServerError
*/
type ContainerRenameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerRenameInternalServerError creates ContainerRenameInternalServerError with default headers values
func NewContainerRenameInternalServerError() *ContainerRenameInternalServerError {

	return &ContainerRenameInternalServerError{}
}

// WithPayload adds the payload to the container rename internal server error response
func (o *ContainerRenameInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerRenameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container rename internal server error response
func (o *ContainerRenameInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerRenameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
