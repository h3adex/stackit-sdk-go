/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the PostgresDatabaseParameterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostgresDatabaseParameterResponse{}

/*
	types and functions for parameter
*/

// isArray
type PostgresDatabaseParameterResponseGetParameterAttributeType = *[]PostgresDatabaseParameter
type PostgresDatabaseParameterResponseGetParameterArgType = []PostgresDatabaseParameter
type PostgresDatabaseParameterResponseGetParameterRetType = []PostgresDatabaseParameter

func getPostgresDatabaseParameterResponseGetParameterAttributeTypeOk(arg PostgresDatabaseParameterResponseGetParameterAttributeType) (ret PostgresDatabaseParameterResponseGetParameterRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPostgresDatabaseParameterResponseGetParameterAttributeType(arg *PostgresDatabaseParameterResponseGetParameterAttributeType, val PostgresDatabaseParameterResponseGetParameterRetType) {
	*arg = &val
}

// PostgresDatabaseParameterResponse struct for PostgresDatabaseParameterResponse
type PostgresDatabaseParameterResponse struct {
	// List of the parameter
	Parameter PostgresDatabaseParameterResponseGetParameterAttributeType `json:"parameter,omitempty"`
}

// NewPostgresDatabaseParameterResponse instantiates a new PostgresDatabaseParameterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresDatabaseParameterResponse() *PostgresDatabaseParameterResponse {
	this := PostgresDatabaseParameterResponse{}
	return &this
}

// NewPostgresDatabaseParameterResponseWithDefaults instantiates a new PostgresDatabaseParameterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresDatabaseParameterResponseWithDefaults() *PostgresDatabaseParameterResponse {
	this := PostgresDatabaseParameterResponse{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *PostgresDatabaseParameterResponse) GetParameter() (res PostgresDatabaseParameterResponseGetParameterRetType) {
	res, _ = o.GetParameterOk()
	return
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresDatabaseParameterResponse) GetParameterOk() (ret PostgresDatabaseParameterResponseGetParameterRetType, ok bool) {
	return getPostgresDatabaseParameterResponseGetParameterAttributeTypeOk(o.Parameter)
}

// HasParameter returns a boolean if a field has been set.
func (o *PostgresDatabaseParameterResponse) HasParameter() bool {
	_, ok := o.GetParameterOk()
	return ok
}

// SetParameter gets a reference to the given []PostgresDatabaseParameter and assigns it to the Parameter field.
func (o *PostgresDatabaseParameterResponse) SetParameter(v PostgresDatabaseParameterResponseGetParameterRetType) {
	setPostgresDatabaseParameterResponseGetParameterAttributeType(&o.Parameter, v)
}

func (o PostgresDatabaseParameterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPostgresDatabaseParameterResponseGetParameterAttributeTypeOk(o.Parameter); ok {
		toSerialize["Parameter"] = val
	}
	return toSerialize, nil
}

type NullablePostgresDatabaseParameterResponse struct {
	value *PostgresDatabaseParameterResponse
	isSet bool
}

func (v NullablePostgresDatabaseParameterResponse) Get() *PostgresDatabaseParameterResponse {
	return v.value
}

func (v *NullablePostgresDatabaseParameterResponse) Set(val *PostgresDatabaseParameterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresDatabaseParameterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresDatabaseParameterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresDatabaseParameterResponse(val *PostgresDatabaseParameterResponse) *NullablePostgresDatabaseParameterResponse {
	return &NullablePostgresDatabaseParameterResponse{value: val, isSet: true}
}

func (v NullablePostgresDatabaseParameterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresDatabaseParameterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
