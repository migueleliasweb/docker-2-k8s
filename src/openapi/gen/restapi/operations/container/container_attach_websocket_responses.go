// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ContainerAttachWebsocketSwitchingProtocolsCode is the HTTP code returned for type ContainerAttachWebsocketSwitchingProtocols
const ContainerAttachWebsocketSwitchingProtocolsCode int = 101

/*ContainerAttachWebsocketSwitchingProtocols no error, hints proxy about hijacking

swagger:response containerAttachWebsocketSwitchingProtocols
*/
type ContainerAttachWebsocketSwitchingProtocols struct {
}

// NewContainerAttachWebsocketSwitchingProtocols creates ContainerAttachWebsocketSwitchingProtocols with default headers values
func NewContainerAttachWebsocketSwitchingProtocols() *ContainerAttachWebsocketSwitchingProtocols {

	return &ContainerAttachWebsocketSwitchingProtocols{}
}

// WriteResponse to the client
func (o *ContainerAttachWebsocketSwitchingProtocols) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(101)
}

// ContainerAttachWebsocketOKCode is the HTTP code returned for type ContainerAttachWebsocketOK
const ContainerAttachWebsocketOKCode int = 200

/*ContainerAttachWebsocketOK no error, no upgrade header found

swagger:response containerAttachWebsocketOK
*/
type ContainerAttachWebsocketOK struct {
}

// NewContainerAttachWebsocketOK creates ContainerAttachWebsocketOK with default headers values
func NewContainerAttachWebsocketOK() *ContainerAttachWebsocketOK {

	return &ContainerAttachWebsocketOK{}
}

// WriteResponse to the client
func (o *ContainerAttachWebsocketOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ContainerAttachWebsocketBadRequestCode is the HTTP code returned for type ContainerAttachWebsocketBadRequest
const ContainerAttachWebsocketBadRequestCode int = 400

/*ContainerAttachWebsocketBadRequest bad parameter

swagger:response containerAttachWebsocketBadRequest
*/
type ContainerAttachWebsocketBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerAttachWebsocketBadRequest creates ContainerAttachWebsocketBadRequest with default headers values
func NewContainerAttachWebsocketBadRequest() *ContainerAttachWebsocketBadRequest {

	return &ContainerAttachWebsocketBadRequest{}
}

// WithPayload adds the payload to the container attach websocket bad request response
func (o *ContainerAttachWebsocketBadRequest) WithPayload(payload *models.ErrorResponse) *ContainerAttachWebsocketBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container attach websocket bad request response
func (o *ContainerAttachWebsocketBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerAttachWebsocketBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerAttachWebsocketNotFoundCode is the HTTP code returned for type ContainerAttachWebsocketNotFound
const ContainerAttachWebsocketNotFoundCode int = 404

/*ContainerAttachWebsocketNotFound no such container

swagger:response containerAttachWebsocketNotFound
*/
type ContainerAttachWebsocketNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerAttachWebsocketNotFound creates ContainerAttachWebsocketNotFound with default headers values
func NewContainerAttachWebsocketNotFound() *ContainerAttachWebsocketNotFound {

	return &ContainerAttachWebsocketNotFound{}
}

// WithPayload adds the payload to the container attach websocket not found response
func (o *ContainerAttachWebsocketNotFound) WithPayload(payload *models.ErrorResponse) *ContainerAttachWebsocketNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container attach websocket not found response
func (o *ContainerAttachWebsocketNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerAttachWebsocketNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerAttachWebsocketInternalServerErrorCode is the HTTP code returned for type ContainerAttachWebsocketInternalServerError
const ContainerAttachWebsocketInternalServerErrorCode int = 500

/*ContainerAttachWebsocketInternalServerError server error

swagger:response containerAttachWebsocketInternalServerError
*/
type ContainerAttachWebsocketInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerAttachWebsocketInternalServerError creates ContainerAttachWebsocketInternalServerError with default headers values
func NewContainerAttachWebsocketInternalServerError() *ContainerAttachWebsocketInternalServerError {

	return &ContainerAttachWebsocketInternalServerError{}
}

// WithPayload adds the payload to the container attach websocket internal server error response
func (o *ContainerAttachWebsocketInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerAttachWebsocketInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container attach websocket internal server error response
func (o *ContainerAttachWebsocketInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerAttachWebsocketInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}