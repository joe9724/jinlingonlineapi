// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "jinlingonlineapi/models"
)

// SearchOKCode is the HTTP code returned for type SearchOK
const SearchOKCode int = 200

/*SearchOK 获取列表

swagger:response searchOK
*/
type SearchOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewSearchOK creates SearchOK with default headers values
func NewSearchOK() *SearchOK {
	return &SearchOK{}
}

// WithPayload adds the payload to the search o k response
func (o *SearchOK) WithPayload(payload *models.InlineResponse2004) *SearchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search o k response
func (o *SearchOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrSearchDefault error

swagger:response searchDefault
*/
type NrSearchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrSearchDefault creates NrSearchDefault with default headers values
func NewNrSearchDefault(code int) *NrSearchDefault {
	if code <= 0 {
		code = 500
	}

	return &NrSearchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the search default response
func (o *NrSearchDefault) WithStatusCode(code int) *NrSearchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the search default response
func (o *NrSearchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the search default response
func (o *NrSearchDefault) WithPayload(payload *models.Error) *NrSearchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search default response
func (o *NrSearchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrSearchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
