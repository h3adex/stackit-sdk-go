/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ZoneObservabilityExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneObservabilityExtension{}

/*
	types and functions for observabilityInstanceId
*/

// isNotNullableString
type ZoneObservabilityExtensionGetObservabilityInstanceIdAttributeType = *string

func getZoneObservabilityExtensionGetObservabilityInstanceIdAttributeTypeOk(arg ZoneObservabilityExtensionGetObservabilityInstanceIdAttributeType) (ret ZoneObservabilityExtensionGetObservabilityInstanceIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneObservabilityExtensionGetObservabilityInstanceIdAttributeType(arg *ZoneObservabilityExtensionGetObservabilityInstanceIdAttributeType, val ZoneObservabilityExtensionGetObservabilityInstanceIdRetType) {
	*arg = &val
}

type ZoneObservabilityExtensionGetObservabilityInstanceIdArgType = string
type ZoneObservabilityExtensionGetObservabilityInstanceIdRetType = string

/*
	types and functions for state
*/

// isNotNullableString
type ZoneObservabilityExtensionGetStateAttributeType = *string

func getZoneObservabilityExtensionGetStateAttributeTypeOk(arg ZoneObservabilityExtensionGetStateAttributeType) (ret ZoneObservabilityExtensionGetStateRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneObservabilityExtensionGetStateAttributeType(arg *ZoneObservabilityExtensionGetStateAttributeType, val ZoneObservabilityExtensionGetStateRetType) {
	*arg = &val
}

type ZoneObservabilityExtensionGetStateArgType = string
type ZoneObservabilityExtensionGetStateRetType = string

// ZoneObservabilityExtension struct for ZoneObservabilityExtension
type ZoneObservabilityExtension struct {
	// REQUIRED
	ObservabilityInstanceId ZoneObservabilityExtensionGetObservabilityInstanceIdAttributeType `json:"observabilityInstanceId" required:"true"`
	State                   ZoneObservabilityExtensionGetStateAttributeType                   `json:"state,omitempty"`
}

type _ZoneObservabilityExtension ZoneObservabilityExtension

// NewZoneObservabilityExtension instantiates a new ZoneObservabilityExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneObservabilityExtension(observabilityInstanceId ZoneObservabilityExtensionGetObservabilityInstanceIdArgType) *ZoneObservabilityExtension {
	this := ZoneObservabilityExtension{}
	setZoneObservabilityExtensionGetObservabilityInstanceIdAttributeType(&this.ObservabilityInstanceId, observabilityInstanceId)
	return &this
}

// NewZoneObservabilityExtensionWithDefaults instantiates a new ZoneObservabilityExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneObservabilityExtensionWithDefaults() *ZoneObservabilityExtension {
	this := ZoneObservabilityExtension{}
	return &this
}

// GetObservabilityInstanceId returns the ObservabilityInstanceId field value
func (o *ZoneObservabilityExtension) GetObservabilityInstanceId() (ret ZoneObservabilityExtensionGetObservabilityInstanceIdRetType) {
	ret, _ = o.GetObservabilityInstanceIdOk()
	return ret
}

// GetObservabilityInstanceIdOk returns a tuple with the ObservabilityInstanceId field value
// and a boolean to check if the value has been set.
func (o *ZoneObservabilityExtension) GetObservabilityInstanceIdOk() (ret ZoneObservabilityExtensionGetObservabilityInstanceIdRetType, ok bool) {
	return getZoneObservabilityExtensionGetObservabilityInstanceIdAttributeTypeOk(o.ObservabilityInstanceId)
}

// SetObservabilityInstanceId sets field value
func (o *ZoneObservabilityExtension) SetObservabilityInstanceId(v ZoneObservabilityExtensionGetObservabilityInstanceIdRetType) {
	setZoneObservabilityExtensionGetObservabilityInstanceIdAttributeType(&o.ObservabilityInstanceId, v)
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ZoneObservabilityExtension) GetState() (res ZoneObservabilityExtensionGetStateRetType) {
	res, _ = o.GetStateOk()
	return
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneObservabilityExtension) GetStateOk() (ret ZoneObservabilityExtensionGetStateRetType, ok bool) {
	return getZoneObservabilityExtensionGetStateAttributeTypeOk(o.State)
}

// HasState returns a boolean if a field has been set.
func (o *ZoneObservabilityExtension) HasState() bool {
	_, ok := o.GetStateOk()
	return ok
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ZoneObservabilityExtension) SetState(v ZoneObservabilityExtensionGetStateRetType) {
	setZoneObservabilityExtensionGetStateAttributeType(&o.State, v)
}

func (o ZoneObservabilityExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getZoneObservabilityExtensionGetObservabilityInstanceIdAttributeTypeOk(o.ObservabilityInstanceId); ok {
		toSerialize["ObservabilityInstanceId"] = val
	}
	if val, ok := getZoneObservabilityExtensionGetStateAttributeTypeOk(o.State); ok {
		toSerialize["State"] = val
	}
	return toSerialize, nil
}

type NullableZoneObservabilityExtension struct {
	value *ZoneObservabilityExtension
	isSet bool
}

func (v NullableZoneObservabilityExtension) Get() *ZoneObservabilityExtension {
	return v.value
}

func (v *NullableZoneObservabilityExtension) Set(val *ZoneObservabilityExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneObservabilityExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneObservabilityExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneObservabilityExtension(val *ZoneObservabilityExtension) *NullableZoneObservabilityExtension {
	return &NullableZoneObservabilityExtension{value: val, isSet: true}
}

func (v NullableZoneObservabilityExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneObservabilityExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
