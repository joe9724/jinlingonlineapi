// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrCategoryListParams creates a new NrCategoryListParams object
// with the default values initialized.
func NewNrCategoryListParams() NrCategoryListParams {
	var ()
	return NrCategoryListParams{}
}

// NrCategoryListParams contains all the bound params for the category list operation
// typically these are obtained from a http.Request
//
// swagger:parameters /category/list
type NrCategoryListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*客户端类型
	  In: query
	*/
	Client *string
	/*结束时间戳
	  In: query
	*/
	EndTs *int64
	/*唯一识别号
	  In: query
	*/
	Imei *string
	/*用户id
	  In: query
	*/
	MemberID *string
	/*分页索引
	  In: query
	*/
	PageIndex *int64
	/*分页尺寸
	  In: query
	*/
	PageSize *int64
	/*-1表示根目录 其他表示获取子类
	  In: query
	*/
	ParentID *int64
	/*开始时间戳
	  In: query
	*/
	StartTs *int64
	/*分类
	  In: query
	*/
	Tags *string
	/*版本号
	  In: query
	*/
	Version *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrCategoryListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qClient, qhkClient, _ := qs.GetOK("client")
	if err := o.bindClient(qClient, qhkClient, route.Formats); err != nil {
		res = append(res, err)
	}

	qEndTs, qhkEndTs, _ := qs.GetOK("endTs")
	if err := o.bindEndTs(qEndTs, qhkEndTs, route.Formats); err != nil {
		res = append(res, err)
	}

	qImei, qhkImei, _ := qs.GetOK("imei")
	if err := o.bindImei(qImei, qhkImei, route.Formats); err != nil {
		res = append(res, err)
	}

	qMemberID, qhkMemberID, _ := qs.GetOK("memberId")
	if err := o.bindMemberID(qMemberID, qhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageIndex, qhkPageIndex, _ := qs.GetOK("pageIndex")
	if err := o.bindPageIndex(qPageIndex, qhkPageIndex, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qParentID, qhkParentID, _ := qs.GetOK("parentId")
	if err := o.bindParentID(qParentID, qhkParentID, route.Formats); err != nil {
		res = append(res, err)
	}

	qStartTs, qhkStartTs, _ := qs.GetOK("startTs")
	if err := o.bindStartTs(qStartTs, qhkStartTs, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrCategoryListParams) bindClient(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Client = &raw

	return nil
}

func (o *NrCategoryListParams) bindEndTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("endTs", "query", "int64", raw)
	}
	o.EndTs = &value

	return nil
}

func (o *NrCategoryListParams) bindImei(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Imei = &raw

	return nil
}

func (o *NrCategoryListParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MemberID = &raw

	return nil
}

func (o *NrCategoryListParams) bindPageIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageIndex", "query", "int64", raw)
	}
	o.PageIndex = &value

	return nil
}

func (o *NrCategoryListParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = &value

	return nil
}

func (o *NrCategoryListParams) bindParentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("parentId", "query", "int64", raw)
	}
	o.ParentID = &value

	return nil
}

func (o *NrCategoryListParams) bindStartTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("startTs", "query", "int64", raw)
	}
	o.StartTs = &value

	return nil
}

func (o *NrCategoryListParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Tags = &raw

	return nil
}

func (o *NrCategoryListParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Version = &raw

	return nil
}
