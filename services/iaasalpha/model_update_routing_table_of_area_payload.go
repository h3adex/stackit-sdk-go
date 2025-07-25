/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 2alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the UpdateRoutingTableOfAreaPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoutingTableOfAreaPayload{}

/*
	types and functions for description
*/

// isNotNullableString
type UpdateRoutingTableOfAreaPayloadGetDescriptionAttributeType = *string

func getUpdateRoutingTableOfAreaPayloadGetDescriptionAttributeTypeOk(arg UpdateRoutingTableOfAreaPayloadGetDescriptionAttributeType) (ret UpdateRoutingTableOfAreaPayloadGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateRoutingTableOfAreaPayloadGetDescriptionAttributeType(arg *UpdateRoutingTableOfAreaPayloadGetDescriptionAttributeType, val UpdateRoutingTableOfAreaPayloadGetDescriptionRetType) {
	*arg = &val
}

type UpdateRoutingTableOfAreaPayloadGetDescriptionArgType = string
type UpdateRoutingTableOfAreaPayloadGetDescriptionRetType = string

/*
	types and functions for labels
*/

// isFreeform
type UpdateRoutingTableOfAreaPayloadGetLabelsAttributeType = *map[string]interface{}
type UpdateRoutingTableOfAreaPayloadGetLabelsArgType = map[string]interface{}
type UpdateRoutingTableOfAreaPayloadGetLabelsRetType = map[string]interface{}

func getUpdateRoutingTableOfAreaPayloadGetLabelsAttributeTypeOk(arg UpdateRoutingTableOfAreaPayloadGetLabelsAttributeType) (ret UpdateRoutingTableOfAreaPayloadGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateRoutingTableOfAreaPayloadGetLabelsAttributeType(arg *UpdateRoutingTableOfAreaPayloadGetLabelsAttributeType, val UpdateRoutingTableOfAreaPayloadGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type UpdateRoutingTableOfAreaPayloadGetNameAttributeType = *string

func getUpdateRoutingTableOfAreaPayloadGetNameAttributeTypeOk(arg UpdateRoutingTableOfAreaPayloadGetNameAttributeType) (ret UpdateRoutingTableOfAreaPayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateRoutingTableOfAreaPayloadGetNameAttributeType(arg *UpdateRoutingTableOfAreaPayloadGetNameAttributeType, val UpdateRoutingTableOfAreaPayloadGetNameRetType) {
	*arg = &val
}

type UpdateRoutingTableOfAreaPayloadGetNameArgType = string
type UpdateRoutingTableOfAreaPayloadGetNameRetType = string

// UpdateRoutingTableOfAreaPayload Object that represents the request body for a routing table update.
type UpdateRoutingTableOfAreaPayload struct {
	// Description Object. Allows string up to 255 Characters.
	Description UpdateRoutingTableOfAreaPayloadGetDescriptionAttributeType `json:"description,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels UpdateRoutingTableOfAreaPayloadGetLabelsAttributeType `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name UpdateRoutingTableOfAreaPayloadGetNameAttributeType `json:"name,omitempty"`
}

// NewUpdateRoutingTableOfAreaPayload instantiates a new UpdateRoutingTableOfAreaPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoutingTableOfAreaPayload() *UpdateRoutingTableOfAreaPayload {
	this := UpdateRoutingTableOfAreaPayload{}
	return &this
}

// NewUpdateRoutingTableOfAreaPayloadWithDefaults instantiates a new UpdateRoutingTableOfAreaPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoutingTableOfAreaPayloadWithDefaults() *UpdateRoutingTableOfAreaPayload {
	this := UpdateRoutingTableOfAreaPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateRoutingTableOfAreaPayload) GetDescription() (res UpdateRoutingTableOfAreaPayloadGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoutingTableOfAreaPayload) GetDescriptionOk() (ret UpdateRoutingTableOfAreaPayloadGetDescriptionRetType, ok bool) {
	return getUpdateRoutingTableOfAreaPayloadGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateRoutingTableOfAreaPayload) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateRoutingTableOfAreaPayload) SetDescription(v UpdateRoutingTableOfAreaPayloadGetDescriptionRetType) {
	setUpdateRoutingTableOfAreaPayloadGetDescriptionAttributeType(&o.Description, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateRoutingTableOfAreaPayload) GetLabels() (res UpdateRoutingTableOfAreaPayloadGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoutingTableOfAreaPayload) GetLabelsOk() (ret UpdateRoutingTableOfAreaPayloadGetLabelsRetType, ok bool) {
	return getUpdateRoutingTableOfAreaPayloadGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateRoutingTableOfAreaPayload) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateRoutingTableOfAreaPayload) SetLabels(v UpdateRoutingTableOfAreaPayloadGetLabelsRetType) {
	setUpdateRoutingTableOfAreaPayloadGetLabelsAttributeType(&o.Labels, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRoutingTableOfAreaPayload) GetName() (res UpdateRoutingTableOfAreaPayloadGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoutingTableOfAreaPayload) GetNameOk() (ret UpdateRoutingTableOfAreaPayloadGetNameRetType, ok bool) {
	return getUpdateRoutingTableOfAreaPayloadGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRoutingTableOfAreaPayload) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRoutingTableOfAreaPayload) SetName(v UpdateRoutingTableOfAreaPayloadGetNameRetType) {
	setUpdateRoutingTableOfAreaPayloadGetNameAttributeType(&o.Name, v)
}

func (o UpdateRoutingTableOfAreaPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateRoutingTableOfAreaPayloadGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getUpdateRoutingTableOfAreaPayloadGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getUpdateRoutingTableOfAreaPayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	return toSerialize, nil
}

type NullableUpdateRoutingTableOfAreaPayload struct {
	value *UpdateRoutingTableOfAreaPayload
	isSet bool
}

func (v NullableUpdateRoutingTableOfAreaPayload) Get() *UpdateRoutingTableOfAreaPayload {
	return v.value
}

func (v *NullableUpdateRoutingTableOfAreaPayload) Set(val *UpdateRoutingTableOfAreaPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoutingTableOfAreaPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoutingTableOfAreaPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoutingTableOfAreaPayload(val *UpdateRoutingTableOfAreaPayload) *NullableUpdateRoutingTableOfAreaPayload {
	return &NullableUpdateRoutingTableOfAreaPayload{value: val, isSet: true}
}

func (v NullableUpdateRoutingTableOfAreaPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoutingTableOfAreaPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
