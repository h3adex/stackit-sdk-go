/*
STACKIT Key Management Service API

This API provides endpoints for managing keys and key rings.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kms

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the KeyRing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyRing{}

/*
	types and functions for createdAt
*/

// isDateTime
type KeyRingGetCreatedAtAttributeType = *time.Time
type KeyRingGetCreatedAtArgType = time.Time
type KeyRingGetCreatedAtRetType = time.Time

func getKeyRingGetCreatedAtAttributeTypeOk(arg KeyRingGetCreatedAtAttributeType) (ret KeyRingGetCreatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setKeyRingGetCreatedAtAttributeType(arg *KeyRingGetCreatedAtAttributeType, val KeyRingGetCreatedAtRetType) {
	*arg = &val
}

/*
	types and functions for description
*/

// isNotNullableString
type KeyRingGetDescriptionAttributeType = *string

func getKeyRingGetDescriptionAttributeTypeOk(arg KeyRingGetDescriptionAttributeType) (ret KeyRingGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setKeyRingGetDescriptionAttributeType(arg *KeyRingGetDescriptionAttributeType, val KeyRingGetDescriptionRetType) {
	*arg = &val
}

type KeyRingGetDescriptionArgType = string
type KeyRingGetDescriptionRetType = string

/*
	types and functions for displayName
*/

// isNotNullableString
type KeyRingGetDisplayNameAttributeType = *string

func getKeyRingGetDisplayNameAttributeTypeOk(arg KeyRingGetDisplayNameAttributeType) (ret KeyRingGetDisplayNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setKeyRingGetDisplayNameAttributeType(arg *KeyRingGetDisplayNameAttributeType, val KeyRingGetDisplayNameRetType) {
	*arg = &val
}

type KeyRingGetDisplayNameArgType = string
type KeyRingGetDisplayNameRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type KeyRingGetIdAttributeType = *string

func getKeyRingGetIdAttributeTypeOk(arg KeyRingGetIdAttributeType) (ret KeyRingGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setKeyRingGetIdAttributeType(arg *KeyRingGetIdAttributeType, val KeyRingGetIdRetType) {
	*arg = &val
}

type KeyRingGetIdArgType = string
type KeyRingGetIdRetType = string

/*
	types and functions for state
*/

// isEnum

// KeyRingState The current state of the key ring.
// value type for enums
type KeyRingState string

// List of State
const (
	KEYRINGSTATE_CREATING KeyRingState = "creating"
	KEYRINGSTATE_ACTIVE   KeyRingState = "active"
	KEYRINGSTATE_DELETED  KeyRingState = "deleted"
)

// All allowed values of KeyRing enum
var AllowedKeyRingStateEnumValues = []KeyRingState{
	"creating",
	"active",
	"deleted",
}

func (v *KeyRingState) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson KeyRingState
	var value TmpJson
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// Allow unmarshalling zero value for testing purposes
	var zeroValue TmpJson
	if value == zeroValue {
		return nil
	}
	enumTypeValue := KeyRingState(value)
	for _, existing := range AllowedKeyRingStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyRing", value)
}

// NewKeyRingStateFromValue returns a pointer to a valid KeyRingState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyRingStateFromValue(v KeyRingState) (*KeyRingState, error) {
	ev := KeyRingState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyRingState: valid values are %v", v, AllowedKeyRingStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyRingState) IsValid() bool {
	for _, existing := range AllowedKeyRingStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateState value
func (v KeyRingState) Ptr() *KeyRingState {
	return &v
}

type NullableKeyRingState struct {
	value *KeyRingState
	isSet bool
}

func (v NullableKeyRingState) Get() *KeyRingState {
	return v.value
}

func (v *NullableKeyRingState) Set(val *KeyRingState) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyRingState) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyRingState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyRingState(val *KeyRingState) *NullableKeyRingState {
	return &NullableKeyRingState{value: val, isSet: true}
}

func (v NullableKeyRingState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyRingState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type KeyRingGetStateAttributeType = *KeyRingState
type KeyRingGetStateArgType = KeyRingState
type KeyRingGetStateRetType = KeyRingState

func getKeyRingGetStateAttributeTypeOk(arg KeyRingGetStateAttributeType) (ret KeyRingGetStateRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setKeyRingGetStateAttributeType(arg *KeyRingGetStateAttributeType, val KeyRingGetStateRetType) {
	*arg = &val
}

// KeyRing struct for KeyRing
type KeyRing struct {
	// The date and time the creation of the key ring was triggered.
	// REQUIRED
	CreatedAt KeyRingGetCreatedAtAttributeType `json:"createdAt" required:"true"`
	// A user chosen description to distinguish multiple key rings.
	Description KeyRingGetDescriptionAttributeType `json:"description,omitempty"`
	// The display name to distinguish multiple key rings.
	// REQUIRED
	DisplayName KeyRingGetDisplayNameAttributeType `json:"displayName" required:"true"`
	// A auto generated unique id which identifies the key ring.
	// REQUIRED
	Id KeyRingGetIdAttributeType `json:"id" required:"true"`
	// The current state of the key ring.
	// REQUIRED
	State KeyRingGetStateAttributeType `json:"state" required:"true"`
}

type _KeyRing KeyRing

// NewKeyRing instantiates a new KeyRing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyRing(createdAt KeyRingGetCreatedAtArgType, displayName KeyRingGetDisplayNameArgType, id KeyRingGetIdArgType, state KeyRingGetStateArgType) *KeyRing {
	this := KeyRing{}
	setKeyRingGetCreatedAtAttributeType(&this.CreatedAt, createdAt)
	setKeyRingGetDisplayNameAttributeType(&this.DisplayName, displayName)
	setKeyRingGetIdAttributeType(&this.Id, id)
	setKeyRingGetStateAttributeType(&this.State, state)
	return &this
}

// NewKeyRingWithDefaults instantiates a new KeyRing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyRingWithDefaults() *KeyRing {
	this := KeyRing{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *KeyRing) GetCreatedAt() (ret KeyRingGetCreatedAtRetType) {
	ret, _ = o.GetCreatedAtOk()
	return ret
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *KeyRing) GetCreatedAtOk() (ret KeyRingGetCreatedAtRetType, ok bool) {
	return getKeyRingGetCreatedAtAttributeTypeOk(o.CreatedAt)
}

// SetCreatedAt sets field value
func (o *KeyRing) SetCreatedAt(v KeyRingGetCreatedAtRetType) {
	setKeyRingGetCreatedAtAttributeType(&o.CreatedAt, v)
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KeyRing) GetDescription() (res KeyRingGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRing) GetDescriptionOk() (ret KeyRingGetDescriptionRetType, ok bool) {
	return getKeyRingGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *KeyRing) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KeyRing) SetDescription(v KeyRingGetDescriptionRetType) {
	setKeyRingGetDescriptionAttributeType(&o.Description, v)
}

// GetDisplayName returns the DisplayName field value
func (o *KeyRing) GetDisplayName() (ret KeyRingGetDisplayNameRetType) {
	ret, _ = o.GetDisplayNameOk()
	return ret
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *KeyRing) GetDisplayNameOk() (ret KeyRingGetDisplayNameRetType, ok bool) {
	return getKeyRingGetDisplayNameAttributeTypeOk(o.DisplayName)
}

// SetDisplayName sets field value
func (o *KeyRing) SetDisplayName(v KeyRingGetDisplayNameRetType) {
	setKeyRingGetDisplayNameAttributeType(&o.DisplayName, v)
}

// GetId returns the Id field value
func (o *KeyRing) GetId() (ret KeyRingGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KeyRing) GetIdOk() (ret KeyRingGetIdRetType, ok bool) {
	return getKeyRingGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *KeyRing) SetId(v KeyRingGetIdRetType) {
	setKeyRingGetIdAttributeType(&o.Id, v)
}

// GetState returns the State field value
func (o *KeyRing) GetState() (ret KeyRingGetStateRetType) {
	ret, _ = o.GetStateOk()
	return ret
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *KeyRing) GetStateOk() (ret KeyRingGetStateRetType, ok bool) {
	return getKeyRingGetStateAttributeTypeOk(o.State)
}

// SetState sets field value
func (o *KeyRing) SetState(v KeyRingGetStateRetType) {
	setKeyRingGetStateAttributeType(&o.State, v)
}

func (o KeyRing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getKeyRingGetCreatedAtAttributeTypeOk(o.CreatedAt); ok {
		toSerialize["CreatedAt"] = val
	}
	if val, ok := getKeyRingGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getKeyRingGetDisplayNameAttributeTypeOk(o.DisplayName); ok {
		toSerialize["DisplayName"] = val
	}
	if val, ok := getKeyRingGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getKeyRingGetStateAttributeTypeOk(o.State); ok {
		toSerialize["State"] = val
	}
	return toSerialize, nil
}

type NullableKeyRing struct {
	value *KeyRing
	isSet bool
}

func (v NullableKeyRing) Get() *KeyRing {
	return v.value
}

func (v *NullableKeyRing) Set(val *KeyRing) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyRing) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyRing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyRing(val *KeyRing) *NullableKeyRing {
	return &NullableKeyRing{value: val, isSet: true}
}

func (v NullableKeyRing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyRing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
