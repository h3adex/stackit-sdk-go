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

// checks if the ImageShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageShare{}

/*
	types and functions for parentOrganization
*/

// isBoolean
type ImageSharegetParentOrganizationAttributeType = *bool
type ImageSharegetParentOrganizationArgType = bool
type ImageSharegetParentOrganizationRetType = bool

func getImageSharegetParentOrganizationAttributeTypeOk(arg ImageSharegetParentOrganizationAttributeType) (ret ImageSharegetParentOrganizationRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setImageSharegetParentOrganizationAttributeType(arg *ImageSharegetParentOrganizationAttributeType, val ImageSharegetParentOrganizationRetType) {
	*arg = &val
}

/*
	types and functions for projects
*/

// isArray
type ImageShareGetProjectsAttributeType = *[]string
type ImageShareGetProjectsArgType = []string
type ImageShareGetProjectsRetType = []string

func getImageShareGetProjectsAttributeTypeOk(arg ImageShareGetProjectsAttributeType) (ret ImageShareGetProjectsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setImageShareGetProjectsAttributeType(arg *ImageShareGetProjectsAttributeType, val ImageShareGetProjectsRetType) {
	*arg = &val
}

// ImageShare Share details of an Image. For requests ParentOrganization and Projects are mutually exclusive.
type ImageShare struct {
	// Image is shared with all projects inside the image owners organization.
	ParentOrganization ImageSharegetParentOrganizationAttributeType `json:"parentOrganization,omitempty"`
	// List of all projects the Image is shared with.
	Projects ImageShareGetProjectsAttributeType `json:"projects,omitempty"`
}

// NewImageShare instantiates a new ImageShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageShare() *ImageShare {
	this := ImageShare{}
	return &this
}

// NewImageShareWithDefaults instantiates a new ImageShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageShareWithDefaults() *ImageShare {
	this := ImageShare{}
	return &this
}

// GetParentOrganization returns the ParentOrganization field value if set, zero value otherwise.
func (o *ImageShare) GetParentOrganization() (res ImageSharegetParentOrganizationRetType) {
	res, _ = o.GetParentOrganizationOk()
	return
}

// GetParentOrganizationOk returns a tuple with the ParentOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShare) GetParentOrganizationOk() (ret ImageSharegetParentOrganizationRetType, ok bool) {
	return getImageSharegetParentOrganizationAttributeTypeOk(o.ParentOrganization)
}

// HasParentOrganization returns a boolean if a field has been set.
func (o *ImageShare) HasParentOrganization() bool {
	_, ok := o.GetParentOrganizationOk()
	return ok
}

// SetParentOrganization gets a reference to the given bool and assigns it to the ParentOrganization field.
func (o *ImageShare) SetParentOrganization(v ImageSharegetParentOrganizationRetType) {
	setImageSharegetParentOrganizationAttributeType(&o.ParentOrganization, v)
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *ImageShare) GetProjects() (res ImageShareGetProjectsRetType) {
	res, _ = o.GetProjectsOk()
	return
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShare) GetProjectsOk() (ret ImageShareGetProjectsRetType, ok bool) {
	return getImageShareGetProjectsAttributeTypeOk(o.Projects)
}

// HasProjects returns a boolean if a field has been set.
func (o *ImageShare) HasProjects() bool {
	_, ok := o.GetProjectsOk()
	return ok
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *ImageShare) SetProjects(v ImageShareGetProjectsRetType) {
	setImageShareGetProjectsAttributeType(&o.Projects, v)
}

func (o ImageShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getImageSharegetParentOrganizationAttributeTypeOk(o.ParentOrganization); ok {
		toSerialize["ParentOrganization"] = val
	}
	if val, ok := getImageShareGetProjectsAttributeTypeOk(o.Projects); ok {
		toSerialize["Projects"] = val
	}
	return toSerialize, nil
}

type NullableImageShare struct {
	value *ImageShare
	isSet bool
}

func (v NullableImageShare) Get() *ImageShare {
	return v.value
}

func (v *NullableImageShare) Set(val *ImageShare) {
	v.value = val
	v.isSet = true
}

func (v NullableImageShare) IsSet() bool {
	return v.isSet
}

func (v *NullableImageShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageShare(val *ImageShare) *NullableImageShare {
	return &NullableImageShare{value: val, isSet: true}
}

func (v NullableImageShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
