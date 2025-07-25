/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 2.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"encoding/json"
)

// checks if the CredentialsGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsGroup{}

/*
	types and functions for credentialsGroupId
*/

// isNotNullableString
type CredentialsGroupGetCredentialsGroupIdAttributeType = *string

func getCredentialsGroupGetCredentialsGroupIdAttributeTypeOk(arg CredentialsGroupGetCredentialsGroupIdAttributeType) (ret CredentialsGroupGetCredentialsGroupIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGroupGetCredentialsGroupIdAttributeType(arg *CredentialsGroupGetCredentialsGroupIdAttributeType, val CredentialsGroupGetCredentialsGroupIdRetType) {
	*arg = &val
}

type CredentialsGroupGetCredentialsGroupIdArgType = string
type CredentialsGroupGetCredentialsGroupIdRetType = string

/*
	types and functions for displayName
*/

// isNotNullableString
type CredentialsGroupGetDisplayNameAttributeType = *string

func getCredentialsGroupGetDisplayNameAttributeTypeOk(arg CredentialsGroupGetDisplayNameAttributeType) (ret CredentialsGroupGetDisplayNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGroupGetDisplayNameAttributeType(arg *CredentialsGroupGetDisplayNameAttributeType, val CredentialsGroupGetDisplayNameRetType) {
	*arg = &val
}

type CredentialsGroupGetDisplayNameArgType = string
type CredentialsGroupGetDisplayNameRetType = string

/*
	types and functions for urn
*/

// isNotNullableString
type CredentialsGroupGetUrnAttributeType = *string

func getCredentialsGroupGetUrnAttributeTypeOk(arg CredentialsGroupGetUrnAttributeType) (ret CredentialsGroupGetUrnRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGroupGetUrnAttributeType(arg *CredentialsGroupGetUrnAttributeType, val CredentialsGroupGetUrnRetType) {
	*arg = &val
}

type CredentialsGroupGetUrnArgType = string
type CredentialsGroupGetUrnRetType = string

// CredentialsGroup struct for CredentialsGroup
type CredentialsGroup struct {
	// The ID of the credentials group
	// REQUIRED
	CredentialsGroupId CredentialsGroupGetCredentialsGroupIdAttributeType `json:"credentialsGroupId" required:"true"`
	// Name of the group holding credentials
	// REQUIRED
	DisplayName CredentialsGroupGetDisplayNameAttributeType `json:"displayName" required:"true"`
	// Credentials group URN
	// REQUIRED
	Urn CredentialsGroupGetUrnAttributeType `json:"urn" required:"true"`
}

type _CredentialsGroup CredentialsGroup

// NewCredentialsGroup instantiates a new CredentialsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsGroup(credentialsGroupId CredentialsGroupGetCredentialsGroupIdArgType, displayName CredentialsGroupGetDisplayNameArgType, urn CredentialsGroupGetUrnArgType) *CredentialsGroup {
	this := CredentialsGroup{}
	setCredentialsGroupGetCredentialsGroupIdAttributeType(&this.CredentialsGroupId, credentialsGroupId)
	setCredentialsGroupGetDisplayNameAttributeType(&this.DisplayName, displayName)
	setCredentialsGroupGetUrnAttributeType(&this.Urn, urn)
	return &this
}

// NewCredentialsGroupWithDefaults instantiates a new CredentialsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsGroupWithDefaults() *CredentialsGroup {
	this := CredentialsGroup{}
	return &this
}

// GetCredentialsGroupId returns the CredentialsGroupId field value
func (o *CredentialsGroup) GetCredentialsGroupId() (ret CredentialsGroupGetCredentialsGroupIdRetType) {
	ret, _ = o.GetCredentialsGroupIdOk()
	return ret
}

// GetCredentialsGroupIdOk returns a tuple with the CredentialsGroupId field value
// and a boolean to check if the value has been set.
func (o *CredentialsGroup) GetCredentialsGroupIdOk() (ret CredentialsGroupGetCredentialsGroupIdRetType, ok bool) {
	return getCredentialsGroupGetCredentialsGroupIdAttributeTypeOk(o.CredentialsGroupId)
}

// SetCredentialsGroupId sets field value
func (o *CredentialsGroup) SetCredentialsGroupId(v CredentialsGroupGetCredentialsGroupIdRetType) {
	setCredentialsGroupGetCredentialsGroupIdAttributeType(&o.CredentialsGroupId, v)
}

// GetDisplayName returns the DisplayName field value
func (o *CredentialsGroup) GetDisplayName() (ret CredentialsGroupGetDisplayNameRetType) {
	ret, _ = o.GetDisplayNameOk()
	return ret
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CredentialsGroup) GetDisplayNameOk() (ret CredentialsGroupGetDisplayNameRetType, ok bool) {
	return getCredentialsGroupGetDisplayNameAttributeTypeOk(o.DisplayName)
}

// SetDisplayName sets field value
func (o *CredentialsGroup) SetDisplayName(v CredentialsGroupGetDisplayNameRetType) {
	setCredentialsGroupGetDisplayNameAttributeType(&o.DisplayName, v)
}

// GetUrn returns the Urn field value
func (o *CredentialsGroup) GetUrn() (ret CredentialsGroupGetUrnRetType) {
	ret, _ = o.GetUrnOk()
	return ret
}

// GetUrnOk returns a tuple with the Urn field value
// and a boolean to check if the value has been set.
func (o *CredentialsGroup) GetUrnOk() (ret CredentialsGroupGetUrnRetType, ok bool) {
	return getCredentialsGroupGetUrnAttributeTypeOk(o.Urn)
}

// SetUrn sets field value
func (o *CredentialsGroup) SetUrn(v CredentialsGroupGetUrnRetType) {
	setCredentialsGroupGetUrnAttributeType(&o.Urn, v)
}

func (o CredentialsGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCredentialsGroupGetCredentialsGroupIdAttributeTypeOk(o.CredentialsGroupId); ok {
		toSerialize["CredentialsGroupId"] = val
	}
	if val, ok := getCredentialsGroupGetDisplayNameAttributeTypeOk(o.DisplayName); ok {
		toSerialize["DisplayName"] = val
	}
	if val, ok := getCredentialsGroupGetUrnAttributeTypeOk(o.Urn); ok {
		toSerialize["Urn"] = val
	}
	return toSerialize, nil
}

type NullableCredentialsGroup struct {
	value *CredentialsGroup
	isSet bool
}

func (v NullableCredentialsGroup) Get() *CredentialsGroup {
	return v.value
}

func (v *NullableCredentialsGroup) Set(val *CredentialsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsGroup(val *CredentialsGroup) *NullableCredentialsGroup {
	return &NullableCredentialsGroup{value: val, isSet: true}
}

func (v NullableCredentialsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
