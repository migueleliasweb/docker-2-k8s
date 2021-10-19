// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ContainerListOKCode is the HTTP code returned for type ContainerListOK
const ContainerListOKCode int = 200

/*ContainerListOK no error

swagger:response containerListOK
*/
type ContainerListOK struct {

	/*
	  In: Body
	*/
	Payload models.ContainerSummary `json:"body,omitempty"`
}

// NewContainerListOK creates ContainerListOK with default headers values
func NewContainerListOK() *ContainerListOK {

	return &ContainerListOK{}
}

// WithPayload adds the payload to the container list o k response
func (o *ContainerListOK) WithPayload(payload models.ContainerSummary) *ContainerListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container list o k response
func (o *ContainerListOK) SetPayload(payload models.ContainerSummary) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ContainerSummary{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ContainerListBadRequestCode is the HTTP code returned for type ContainerListBadRequest
const ContainerListBadRequestCode int = 400

/*ContainerListBadRequest bad parameter

swagger:response containerListBadRequest
*/
type ContainerListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerListBadRequest creates ContainerListBadRequest with default headers values
func NewContainerListBadRequest() *ContainerListBadRequest {

	return &ContainerListBadRequest{}
}

// WithPayload adds the payload to the container list bad request response
func (o *ContainerListBadRequest) WithPayload(payload *models.ErrorResponse) *ContainerListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container list bad request response
func (o *ContainerListBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerListInternalServerErrorCode is the HTTP code returned for type ContainerListInternalServerError
const ContainerListInternalServerErrorCode int = 500

/*ContainerListInternalServerError server error

swagger:response containerListInternalServerError
*/
type ContainerListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerListInternalServerError creates ContainerListInternalServerError with default headers values
func NewContainerListInternalServerError() *ContainerListInternalServerError {

	return &ContainerListInternalServerError{}
}

// WithPayload adds the payload to the container list internal server error response
func (o *ContainerListInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container list internal server error response
func (o *ContainerListInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
