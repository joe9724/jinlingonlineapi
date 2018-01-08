// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Activity 活动
// swagger:model Activity
type Activity struct {

	// 二维码生成地址
	QRcodeURL string `json:"QRcodeUrl,omitempty"`

	// activity Id
	ActivityID int64 `json:"activityId,omitempty"`

	// 活动地点
	Area string `json:"area,omitempty"`

	// 审核状态
	CheckStatus int64 `json:"checkStatus,omitempty"`

	// checked time
	CheckedTime int64 `json:"checkedTime,omitempty"`

	// contact addr
	ContactAddr string `json:"contactAddr,omitempty"`

	// contact people
	ContactPeople string `json:"contactPeople,omitempty"`

	// contact phone
	ContactPhone string `json:"contactPhone,omitempty"`

	// cover
	Cover string `json:"cover,omitempty"`

	// end time
	EndTime int64 `json:"endTime,omitempty"`

	// 备用字段
	ExtraInfo1 string `json:"extraInfo1,omitempty"`

	// 备用字段
	ExtraInfo2 string `json:"extraInfo2,omitempty"`

	// 备用字段
	ExtraInfo3 int64 `json:"extraInfo3,omitempty"`

	// extra Url
	ExtraURL string `json:"extraUrl,omitempty"`

	// 关注人数
	FollowCount int64 `json:"followCount,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// 是否关注(针对会员)
	IsFollow int64 `json:"isFollow,omitempty"`

	// 参加人数即报名人数
	JoinedPeople int64 `json:"joinedPeople,omitempty"`

	// lat
	Lat float64 `json:"lat,omitempty"`

	// limit people
	LimitPeople int64 `json:"limitPeople,omitempty"`

	// lon
	Lon float64 `json:"lon,omitempty"`

	// 获得赞助资金
	Money float64 `json:"money,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pc Url
	PcURL string `json:"pcUrl,omitempty"`

	// 点赞数
	PriseCount int64 `json:"priseCount,omitempty"`

	// publish time
	PublishTime int64 `json:"publishTime,omitempty"`

	// 是否由后台发起还是app用户发起
	PublishType int64 `json:"publishType,omitempty"`

	// publish user Id
	PublishUserID int64 `json:"publishUserId,omitempty"`

	// 免费0 付费1 全部2
	ServiceType int64 `json:"serviceType,omitempty"`

	// showicon
	Showicon bool `json:"showicon,omitempty"`

	// start time
	StartTime int64 `json:"startTime,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`

	// wap Url
	WapURL string `json:"wapUrl,omitempty"`
}

// Validate validates this activity
func (m *Activity) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Activity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Activity) UnmarshalBinary(b []byte) error {
	var res Activity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}