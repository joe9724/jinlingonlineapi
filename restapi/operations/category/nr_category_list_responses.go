// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "jinlingonlineapi/models"
)

// CategoryListOKCode is the HTTP code returned for type CategoryListOK
const CategoryListOKCode int = 200

/*CategoryListOK 查询成功

swagger:response categoryListOK
*/
type CategoryListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2008 `json:"body,omitempty"`
}

// NewCategoryListOK creates CategoryListOK with default headers values
func NewCategoryListOK() *CategoryListOK {
	return &CategoryListOK{}
}

// WithPayload adds the payload to the category list o k response
func (o *CategoryListOK) WithPayload(payload *models.InlineResponse2008) *CategoryListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the category list o k response
func (o *CategoryListOK) SetPayload(payload *models.InlineResponse2008) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CategoryListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
