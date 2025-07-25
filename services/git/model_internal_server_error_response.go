/*
STACKIT Git API

STACKIT Git management API.

API version: 1beta.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package git

import (
	"encoding/json"
)

// checks if the InternalServerErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalServerErrorResponse{}

/*
	types and functions for error
*/

// isNotNullableString
type InternalServerErrorResponseGetErrorAttributeType = *string

func getInternalServerErrorResponseGetErrorAttributeTypeOk(arg InternalServerErrorResponseGetErrorAttributeType) (ret InternalServerErrorResponseGetErrorRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInternalServerErrorResponseGetErrorAttributeType(arg *InternalServerErrorResponseGetErrorAttributeType, val InternalServerErrorResponseGetErrorRetType) {
	*arg = &val
}

type InternalServerErrorResponseGetErrorArgType = string
type InternalServerErrorResponseGetErrorRetType = string

// InternalServerErrorResponse Internal server error.
type InternalServerErrorResponse struct {
	Error InternalServerErrorResponseGetErrorAttributeType `json:"error,omitempty"`
}

// NewInternalServerErrorResponse instantiates a new InternalServerErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerErrorResponse() *InternalServerErrorResponse {
	this := InternalServerErrorResponse{}
	return &this
}

// NewInternalServerErrorResponseWithDefaults instantiates a new InternalServerErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerErrorResponseWithDefaults() *InternalServerErrorResponse {
	this := InternalServerErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InternalServerErrorResponse) GetError() (res InternalServerErrorResponseGetErrorRetType) {
	res, _ = o.GetErrorOk()
	return
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalServerErrorResponse) GetErrorOk() (ret InternalServerErrorResponseGetErrorRetType, ok bool) {
	return getInternalServerErrorResponseGetErrorAttributeTypeOk(o.Error)
}

// HasError returns a boolean if a field has been set.
func (o *InternalServerErrorResponse) HasError() bool {
	_, ok := o.GetErrorOk()
	return ok
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InternalServerErrorResponse) SetError(v InternalServerErrorResponseGetErrorRetType) {
	setInternalServerErrorResponseGetErrorAttributeType(&o.Error, v)
}

func (o InternalServerErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getInternalServerErrorResponseGetErrorAttributeTypeOk(o.Error); ok {
		toSerialize["Error"] = val
	}
	return toSerialize, nil
}

type NullableInternalServerErrorResponse struct {
	value *InternalServerErrorResponse
	isSet bool
}

func (v NullableInternalServerErrorResponse) Get() *InternalServerErrorResponse {
	return v.value
}

func (v *NullableInternalServerErrorResponse) Set(val *InternalServerErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerErrorResponse(val *InternalServerErrorResponse) *NullableInternalServerErrorResponse {
	return &NullableInternalServerErrorResponse{value: val, isSet: true}
}

func (v NullableInternalServerErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalServerErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
