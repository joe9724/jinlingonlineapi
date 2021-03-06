// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "jinlingonlineapi/models"
)

// ActivityListSponsorOKCode is the HTTP code returned for type ActivityListSponsorOK
const ActivityListSponsorOKCode int = 200

/*ActivityListSponsorOK 获取列表

swagger:response activityListSponsorOK
*/
type ActivityListSponsorOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2006 `json:"body,omitempty"`
}

// NewActivityListSponsorOK creates ActivityListSponsorOK with default headers values
func NewActivityListSponsorOK() *ActivityListSponsorOK {
	return &ActivityListSponsorOK{}
}

// WithPayload adds the payload to the activity list sponsor o k response
func (o *ActivityListSponsorOK) WithPayload(payload *models.InlineResponse2006) *ActivityListSponsorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity list sponsor o k response
func (o *ActivityListSponsorOK) SetPayload(payload *models.InlineResponse2006) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityListSponsorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrActivityListSponsorDefault error

swagger:response activityListSponsorDefault
*/
type NrActivityListSponsorDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrActivityListSponsorDefault creates NrActivityListSponsorDefault with default headers values
func NewNrActivityListSponsorDefault(code int) *NrActivityListSponsorDefault {
	if code <= 0 {
		code = 500
	}

	return &NrActivityListSponsorDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity list sponsor default response
func (o *NrActivityListSponsorDefault) WithStatusCode(code int) *NrActivityListSponsorDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity list sponsor default response
func (o *NrActivityListSponsorDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity list sponsor default response
func (o *NrActivityListSponsorDefault) WithPayload(payload *models.Error) *NrActivityListSponsorDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity list sponsor default response
func (o *NrActivityListSponsorDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrActivityListSponsorDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
