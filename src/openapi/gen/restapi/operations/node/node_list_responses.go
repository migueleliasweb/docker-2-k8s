// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// NodeListOKCode is the HTTP code returned for type NodeListOK
const NodeListOKCode int = 200

/*NodeListOK no error

swagger:response nodeListOK
*/
type NodeListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Node `json:"body,omitempty"`
}

// NewNodeListOK creates NodeListOK with default headers values
func NewNodeListOK() *NodeListOK {

	return &NodeListOK{}
}

// WithPayload adds the payload to the node list o k response
func (o *NodeListOK) WithPayload(payload []*models.Node) *NodeListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the node list o k response
func (o *NodeListOK) SetPayload(payload []*models.Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodeListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Node, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// NodeListInternalServerErrorCode is the HTTP code returned for type NodeListInternalServerError
const NodeListInternalServerErrorCode int = 500

/*NodeListInternalServerError server error

swagger:response nodeListInternalServerError
*/
type NodeListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodeListInternalServerError creates NodeListInternalServerError with default headers values
func NewNodeListInternalServerError() *NodeListInternalServerError {

	return &NodeListInternalServerError{}
}

// WithPayload adds the payload to the node list internal server error response
func (o *NodeListInternalServerError) WithPayload(payload *models.ErrorResponse) *NodeListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the node list internal server error response
func (o *NodeListInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodeListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodeListServiceUnavailableCode is the HTTP code returned for type NodeListServiceUnavailable
const NodeListServiceUnavailableCode int = 503

/*NodeListServiceUnavailable node is not part of a swarm

swagger:response nodeListServiceUnavailable
*/
type NodeListServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodeListServiceUnavailable creates NodeListServiceUnavailable with default headers values
func NewNodeListServiceUnavailable() *NodeListServiceUnavailable {

	return &NodeListServiceUnavailable{}
}

// WithPayload adds the payload to the node list service unavailable response
func (o *NodeListServiceUnavailable) WithPayload(payload *models.ErrorResponse) *NodeListServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the node list service unavailable response
func (o *NodeListServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodeListServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
