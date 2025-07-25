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

// checks if the Credentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credentials{}

/*
	types and functions for password
*/

// isNotNullableString
type CredentialsGetPasswordAttributeType = *string

func getCredentialsGetPasswordAttributeTypeOk(arg CredentialsGetPasswordAttributeType) (ret CredentialsGetPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetPasswordAttributeType(arg *CredentialsGetPasswordAttributeType, val CredentialsGetPasswordRetType) {
	*arg = &val
}

type CredentialsGetPasswordArgType = string
type CredentialsGetPasswordRetType = string

/*
	types and functions for username
*/

// isNotNullableString
type CredentialsGetUsernameAttributeType = *string

func getCredentialsGetUsernameAttributeTypeOk(arg CredentialsGetUsernameAttributeType) (ret CredentialsGetUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetUsernameAttributeType(arg *CredentialsGetUsernameAttributeType, val CredentialsGetUsernameRetType) {
	*arg = &val
}

type CredentialsGetUsernameArgType = string
type CredentialsGetUsernameRetType = string

// Credentials struct for Credentials
type Credentials struct {
	// REQUIRED
	Password CredentialsGetPasswordAttributeType `json:"password" required:"true"`
	// REQUIRED
	Username CredentialsGetUsernameAttributeType `json:"username" required:"true"`
}

type _Credentials Credentials

// NewCredentials instantiates a new Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentials(password CredentialsGetPasswordArgType, username CredentialsGetUsernameArgType) *Credentials {
	this := Credentials{}
	setCredentialsGetPasswordAttributeType(&this.Password, password)
	setCredentialsGetUsernameAttributeType(&this.Username, username)
	return &this
}

// NewCredentialsWithDefaults instantiates a new Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsWithDefaults() *Credentials {
	this := Credentials{}
	return &this
}

// GetPassword returns the Password field value
func (o *Credentials) GetPassword() (ret CredentialsGetPasswordRetType) {
	ret, _ = o.GetPasswordOk()
	return ret
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetPasswordOk() (ret CredentialsGetPasswordRetType, ok bool) {
	return getCredentialsGetPasswordAttributeTypeOk(o.Password)
}

// SetPassword sets field value
func (o *Credentials) SetPassword(v CredentialsGetPasswordRetType) {
	setCredentialsGetPasswordAttributeType(&o.Password, v)
}

// GetUsername returns the Username field value
func (o *Credentials) GetUsername() (ret CredentialsGetUsernameRetType) {
	ret, _ = o.GetUsernameOk()
	return ret
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetUsernameOk() (ret CredentialsGetUsernameRetType, ok bool) {
	return getCredentialsGetUsernameAttributeTypeOk(o.Username)
}

// SetUsername sets field value
func (o *Credentials) SetUsername(v CredentialsGetUsernameRetType) {
	setCredentialsGetUsernameAttributeType(&o.Username, v)
}

func (o Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCredentialsGetPasswordAttributeTypeOk(o.Password); ok {
		toSerialize["Password"] = val
	}
	if val, ok := getCredentialsGetUsernameAttributeTypeOk(o.Username); ok {
		toSerialize["Username"] = val
	}
	return toSerialize, nil
}

type NullableCredentials struct {
	value *Credentials
	isSet bool
}

func (v NullableCredentials) Get() *Credentials {
	return v.value
}

func (v *NullableCredentials) Set(val *Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentials(val *Credentials) *NullableCredentials {
	return &NullableCredentials{value: val, isSet: true}
}

func (v NullableCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
