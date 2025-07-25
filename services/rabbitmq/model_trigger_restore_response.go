/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

import (
	"encoding/json"
)

// checks if the TriggerRestoreResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerRestoreResponse{}

/*
	types and functions for id
*/

// isInteger
type TriggerRestoreResponseGetIdAttributeType = *int64
type TriggerRestoreResponseGetIdArgType = int64
type TriggerRestoreResponseGetIdRetType = int64

func getTriggerRestoreResponseGetIdAttributeTypeOk(arg TriggerRestoreResponseGetIdAttributeType) (ret TriggerRestoreResponseGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setTriggerRestoreResponseGetIdAttributeType(arg *TriggerRestoreResponseGetIdAttributeType, val TriggerRestoreResponseGetIdRetType) {
	*arg = &val
}

// TriggerRestoreResponse struct for TriggerRestoreResponse
type TriggerRestoreResponse struct {
	// REQUIRED
	Id TriggerRestoreResponseGetIdAttributeType `json:"id" required:"true"`
}

type _TriggerRestoreResponse TriggerRestoreResponse

// NewTriggerRestoreResponse instantiates a new TriggerRestoreResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerRestoreResponse(id TriggerRestoreResponseGetIdArgType) *TriggerRestoreResponse {
	this := TriggerRestoreResponse{}
	setTriggerRestoreResponseGetIdAttributeType(&this.Id, id)
	return &this
}

// NewTriggerRestoreResponseWithDefaults instantiates a new TriggerRestoreResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerRestoreResponseWithDefaults() *TriggerRestoreResponse {
	this := TriggerRestoreResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TriggerRestoreResponse) GetId() (ret TriggerRestoreResponseGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TriggerRestoreResponse) GetIdOk() (ret TriggerRestoreResponseGetIdRetType, ok bool) {
	return getTriggerRestoreResponseGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *TriggerRestoreResponse) SetId(v TriggerRestoreResponseGetIdRetType) {
	setTriggerRestoreResponseGetIdAttributeType(&o.Id, v)
}

func (o TriggerRestoreResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getTriggerRestoreResponseGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	return toSerialize, nil
}

type NullableTriggerRestoreResponse struct {
	value *TriggerRestoreResponse
	isSet bool
}

func (v NullableTriggerRestoreResponse) Get() *TriggerRestoreResponse {
	return v.value
}

func (v *NullableTriggerRestoreResponse) Set(val *TriggerRestoreResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerRestoreResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerRestoreResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerRestoreResponse(val *TriggerRestoreResponse) *NullableTriggerRestoreResponse {
	return &NullableTriggerRestoreResponse{value: val, isSet: true}
}

func (v NullableTriggerRestoreResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerRestoreResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
