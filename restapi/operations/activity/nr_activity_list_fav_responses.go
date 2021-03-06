// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "jinlingonlineapi/models"
)

// ActivityListFavOKCode is the HTTP code returned for type ActivityListFavOK
const ActivityListFavOKCode int = 200

/*ActivityListFavOK 获取列表

swagger:response activityListFavOK
*/
type ActivityListFavOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewActivityListFavOK creates ActivityListFavOK with default headers values
func NewActivityListFavOK() *ActivityListFavOK {
	return &ActivityListFavOK{}
}

// WithPayload adds the payload to the activity list fav o k response
func (o *ActivityListFavOK) WithPayload(payload *models.InlineResponse2004) *ActivityListFavOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity list fav o k response
func (o *ActivityListFavOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityListFavOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrActivityListFavDefault error

swagger:response activityListFavDefault
*/
type NrActivityListFavDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrActivityListFavDefault creates NrActivityListFavDefault with default headers values
func NewNrActivityListFavDefault(code int) *NrActivityListFavDefault {
	if code <= 0 {
		code = 500
	}

	return &NrActivityListFavDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity list fav default response
func (o *NrActivityListFavDefault) WithStatusCode(code int) *NrActivityListFavDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity list fav default response
func (o *NrActivityListFavDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity list fav default response
func (o *NrActivityListFavDefault) WithPayload(payload *models.Error) *NrActivityListFavDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity list fav default response
func (o *NrActivityListFavDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrActivityListFavDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
