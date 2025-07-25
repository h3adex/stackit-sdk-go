/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the AlertGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertGroupsResponse{}

/*
	types and functions for data
*/

// isArray
type AlertGroupsResponseGetDataAttributeType = *[]AlertGroup
type AlertGroupsResponseGetDataArgType = []AlertGroup
type AlertGroupsResponseGetDataRetType = []AlertGroup

func getAlertGroupsResponseGetDataAttributeTypeOk(arg AlertGroupsResponseGetDataAttributeType) (ret AlertGroupsResponseGetDataRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertGroupsResponseGetDataAttributeType(arg *AlertGroupsResponseGetDataAttributeType, val AlertGroupsResponseGetDataRetType) {
	*arg = &val
}

/*
	types and functions for message
*/

// isNotNullableString
type AlertGroupsResponseGetMessageAttributeType = *string

func getAlertGroupsResponseGetMessageAttributeTypeOk(arg AlertGroupsResponseGetMessageAttributeType) (ret AlertGroupsResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertGroupsResponseGetMessageAttributeType(arg *AlertGroupsResponseGetMessageAttributeType, val AlertGroupsResponseGetMessageRetType) {
	*arg = &val
}

type AlertGroupsResponseGetMessageArgType = string
type AlertGroupsResponseGetMessageRetType = string

// AlertGroupsResponse struct for AlertGroupsResponse
type AlertGroupsResponse struct {
	// REQUIRED
	Data AlertGroupsResponseGetDataAttributeType `json:"data" required:"true"`
	// REQUIRED
	Message AlertGroupsResponseGetMessageAttributeType `json:"message" required:"true"`
}

type _AlertGroupsResponse AlertGroupsResponse

// NewAlertGroupsResponse instantiates a new AlertGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertGroupsResponse(data AlertGroupsResponseGetDataArgType, message AlertGroupsResponseGetMessageArgType) *AlertGroupsResponse {
	this := AlertGroupsResponse{}
	setAlertGroupsResponseGetDataAttributeType(&this.Data, data)
	setAlertGroupsResponseGetMessageAttributeType(&this.Message, message)
	return &this
}

// NewAlertGroupsResponseWithDefaults instantiates a new AlertGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertGroupsResponseWithDefaults() *AlertGroupsResponse {
	this := AlertGroupsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AlertGroupsResponse) GetData() (ret AlertGroupsResponseGetDataRetType) {
	ret, _ = o.GetDataOk()
	return ret
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AlertGroupsResponse) GetDataOk() (ret AlertGroupsResponseGetDataRetType, ok bool) {
	return getAlertGroupsResponseGetDataAttributeTypeOk(o.Data)
}

// SetData sets field value
func (o *AlertGroupsResponse) SetData(v AlertGroupsResponseGetDataRetType) {
	setAlertGroupsResponseGetDataAttributeType(&o.Data, v)
}

// GetMessage returns the Message field value
func (o *AlertGroupsResponse) GetMessage() (ret AlertGroupsResponseGetMessageRetType) {
	ret, _ = o.GetMessageOk()
	return ret
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlertGroupsResponse) GetMessageOk() (ret AlertGroupsResponseGetMessageRetType, ok bool) {
	return getAlertGroupsResponseGetMessageAttributeTypeOk(o.Message)
}

// SetMessage sets field value
func (o *AlertGroupsResponse) SetMessage(v AlertGroupsResponseGetMessageRetType) {
	setAlertGroupsResponseGetMessageAttributeType(&o.Message, v)
}

func (o AlertGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getAlertGroupsResponseGetDataAttributeTypeOk(o.Data); ok {
		toSerialize["Data"] = val
	}
	if val, ok := getAlertGroupsResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	return toSerialize, nil
}

type NullableAlertGroupsResponse struct {
	value *AlertGroupsResponse
	isSet bool
}

func (v NullableAlertGroupsResponse) Get() *AlertGroupsResponse {
	return v.value
}

func (v *NullableAlertGroupsResponse) Set(val *AlertGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertGroupsResponse(val *AlertGroupsResponse) *NullableAlertGroupsResponse {
	return &NullableAlertGroupsResponse{value: val, isSet: true}
}

func (v NullableAlertGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
