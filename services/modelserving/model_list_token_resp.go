/*
STACKIT Model Serving API

This API provides endpoints for the model serving api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package modelserving

import (
	"encoding/json"
)

// checks if the ListTokenResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTokenResp{}

/*
	types and functions for message
*/

// isNotNullableString
type ListTokenRespGetMessageAttributeType = *string

func getListTokenRespGetMessageAttributeTypeOk(arg ListTokenRespGetMessageAttributeType) (ret ListTokenRespGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListTokenRespGetMessageAttributeType(arg *ListTokenRespGetMessageAttributeType, val ListTokenRespGetMessageRetType) {
	*arg = &val
}

type ListTokenRespGetMessageArgType = string
type ListTokenRespGetMessageRetType = string

/*
	types and functions for tokens
*/

// isArray
type ListTokenRespGetTokensAttributeType = *[]Token
type ListTokenRespGetTokensArgType = []Token
type ListTokenRespGetTokensRetType = []Token

func getListTokenRespGetTokensAttributeTypeOk(arg ListTokenRespGetTokensAttributeType) (ret ListTokenRespGetTokensRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListTokenRespGetTokensAttributeType(arg *ListTokenRespGetTokensAttributeType, val ListTokenRespGetTokensRetType) {
	*arg = &val
}

// ListTokenResp struct for ListTokenResp
type ListTokenResp struct {
	Message ListTokenRespGetMessageAttributeType `json:"message,omitempty"`
	// REQUIRED
	Tokens ListTokenRespGetTokensAttributeType `json:"tokens" required:"true"`
}

type _ListTokenResp ListTokenResp

// NewListTokenResp instantiates a new ListTokenResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTokenResp(tokens ListTokenRespGetTokensArgType) *ListTokenResp {
	this := ListTokenResp{}
	setListTokenRespGetTokensAttributeType(&this.Tokens, tokens)
	return &this
}

// NewListTokenRespWithDefaults instantiates a new ListTokenResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTokenRespWithDefaults() *ListTokenResp {
	this := ListTokenResp{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ListTokenResp) GetMessage() (res ListTokenRespGetMessageRetType) {
	res, _ = o.GetMessageOk()
	return
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokenResp) GetMessageOk() (ret ListTokenRespGetMessageRetType, ok bool) {
	return getListTokenRespGetMessageAttributeTypeOk(o.Message)
}

// HasMessage returns a boolean if a field has been set.
func (o *ListTokenResp) HasMessage() bool {
	_, ok := o.GetMessageOk()
	return ok
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ListTokenResp) SetMessage(v ListTokenRespGetMessageRetType) {
	setListTokenRespGetMessageAttributeType(&o.Message, v)
}

// GetTokens returns the Tokens field value
func (o *ListTokenResp) GetTokens() (ret ListTokenRespGetTokensRetType) {
	ret, _ = o.GetTokensOk()
	return ret
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *ListTokenResp) GetTokensOk() (ret ListTokenRespGetTokensRetType, ok bool) {
	return getListTokenRespGetTokensAttributeTypeOk(o.Tokens)
}

// SetTokens sets field value
func (o *ListTokenResp) SetTokens(v ListTokenRespGetTokensRetType) {
	setListTokenRespGetTokensAttributeType(&o.Tokens, v)
}

func (o ListTokenResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListTokenRespGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	if val, ok := getListTokenRespGetTokensAttributeTypeOk(o.Tokens); ok {
		toSerialize["Tokens"] = val
	}
	return toSerialize, nil
}

type NullableListTokenResp struct {
	value *ListTokenResp
	isSet bool
}

func (v NullableListTokenResp) Get() *ListTokenResp {
	return v.value
}

func (v *NullableListTokenResp) Set(val *ListTokenResp) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokenResp) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokenResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokenResp(val *ListTokenResp) *NullableListTokenResp {
	return &NullableListTokenResp{value: val, isSet: true}
}

func (v NullableListTokenResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokenResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
