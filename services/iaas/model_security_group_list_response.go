/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the SecurityGroupListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupListResponse{}

/*
	types and functions for items
*/

// isArray
type SecurityGroupListResponseGetItemsAttributeType = *[]SecurityGroup
type SecurityGroupListResponseGetItemsArgType = []SecurityGroup
type SecurityGroupListResponseGetItemsRetType = []SecurityGroup

func getSecurityGroupListResponseGetItemsAttributeTypeOk(arg SecurityGroupListResponseGetItemsAttributeType) (ret SecurityGroupListResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupListResponseGetItemsAttributeType(arg *SecurityGroupListResponseGetItemsAttributeType, val SecurityGroupListResponseGetItemsRetType) {
	*arg = &val
}

// SecurityGroupListResponse Security group list response.
type SecurityGroupListResponse struct {
	// A list containing security group objects.
	// REQUIRED
	Items SecurityGroupListResponseGetItemsAttributeType `json:"items" required:"true"`
}

type _SecurityGroupListResponse SecurityGroupListResponse

// NewSecurityGroupListResponse instantiates a new SecurityGroupListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupListResponse(items SecurityGroupListResponseGetItemsArgType) *SecurityGroupListResponse {
	this := SecurityGroupListResponse{}
	setSecurityGroupListResponseGetItemsAttributeType(&this.Items, items)
	return &this
}

// NewSecurityGroupListResponseWithDefaults instantiates a new SecurityGroupListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupListResponseWithDefaults() *SecurityGroupListResponse {
	this := SecurityGroupListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SecurityGroupListResponse) GetItems() (ret SecurityGroupListResponseGetItemsRetType) {
	ret, _ = o.GetItemsOk()
	return ret
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupListResponse) GetItemsOk() (ret SecurityGroupListResponseGetItemsRetType, ok bool) {
	return getSecurityGroupListResponseGetItemsAttributeTypeOk(o.Items)
}

// SetItems sets field value
func (o *SecurityGroupListResponse) SetItems(v SecurityGroupListResponseGetItemsRetType) {
	setSecurityGroupListResponseGetItemsAttributeType(&o.Items, v)
}

func (o SecurityGroupListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getSecurityGroupListResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
	}
	return toSerialize, nil
}

type NullableSecurityGroupListResponse struct {
	value *SecurityGroupListResponse
	isSet bool
}

func (v NullableSecurityGroupListResponse) Get() *SecurityGroupListResponse {
	return v.value
}

func (v *NullableSecurityGroupListResponse) Set(val *SecurityGroupListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupListResponse(val *SecurityGroupListResponse) *NullableSecurityGroupListResponse {
	return &NullableSecurityGroupListResponse{value: val, isSet: true}
}

func (v NullableSecurityGroupListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
