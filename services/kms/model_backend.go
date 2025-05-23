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
)

// Backend The backend that is responsible for maintaining this key.
type Backend string

// List of backend
const (
	BACKEND_SOFTWARE Backend = "software"
)

// All allowed values of Backend enum
var AllowedBackendEnumValues = []Backend{
	"software",
}

func (v *Backend) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// Allow unmarshalling zero value for testing purposes
	var zeroValue string
	if value == zeroValue {
		return nil
	}
	enumTypeValue := Backend(value)
	for _, existing := range AllowedBackendEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Backend", value)
}

// NewBackendFromValue returns a pointer to a valid Backend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackendFromValue(v string) (*Backend, error) {
	ev := Backend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Backend: valid values are %v", v, AllowedBackendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Backend) IsValid() bool {
	for _, existing := range AllowedBackendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to backend value
func (v Backend) Ptr() *Backend {
	return &v
}

type NullableBackend struct {
	value *Backend
	isSet bool
}

func (v NullableBackend) Get() *Backend {
	return v.value
}

func (v *NullableBackend) Set(val *Backend) {
	v.value = val
	v.isSet = true
}

func (v NullableBackend) IsSet() bool {
	return v.isSet
}

func (v *NullableBackend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackend(val *Backend) *NullableBackend {
	return &NullableBackend{value: val, isSet: true}
}

func (v NullableBackend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
