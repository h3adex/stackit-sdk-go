/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

import (
	"encoding/json"
)

// checks if the RolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesResponse{}

/*
	types and functions for resourceId
*/

// isNotNullableString
type RolesResponseGetResourceIdAttributeType = *string

func getRolesResponseGetResourceIdAttributeTypeOk(arg RolesResponseGetResourceIdAttributeType) (ret RolesResponseGetResourceIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRolesResponseGetResourceIdAttributeType(arg *RolesResponseGetResourceIdAttributeType, val RolesResponseGetResourceIdRetType) {
	*arg = &val
}

type RolesResponseGetResourceIdArgType = string
type RolesResponseGetResourceIdRetType = string

/*
	types and functions for resourceType
*/

// isNotNullableString
type RolesResponseGetResourceTypeAttributeType = *string

func getRolesResponseGetResourceTypeAttributeTypeOk(arg RolesResponseGetResourceTypeAttributeType) (ret RolesResponseGetResourceTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRolesResponseGetResourceTypeAttributeType(arg *RolesResponseGetResourceTypeAttributeType, val RolesResponseGetResourceTypeRetType) {
	*arg = &val
}

type RolesResponseGetResourceTypeArgType = string
type RolesResponseGetResourceTypeRetType = string

/*
	types and functions for roles
*/

// isArray
type RolesResponseGetRolesAttributeType = *[]Role
type RolesResponseGetRolesArgType = []Role
type RolesResponseGetRolesRetType = []Role

func getRolesResponseGetRolesAttributeTypeOk(arg RolesResponseGetRolesAttributeType) (ret RolesResponseGetRolesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRolesResponseGetRolesAttributeType(arg *RolesResponseGetRolesAttributeType, val RolesResponseGetRolesRetType) {
	*arg = &val
}

// RolesResponse struct for RolesResponse
type RolesResponse struct {
	// REQUIRED
	ResourceId RolesResponseGetResourceIdAttributeType `json:"resourceId" required:"true"`
	// REQUIRED
	ResourceType RolesResponseGetResourceTypeAttributeType `json:"resourceType" required:"true"`
	// REQUIRED
	Roles RolesResponseGetRolesAttributeType `json:"roles" required:"true"`
}

type _RolesResponse RolesResponse

// NewRolesResponse instantiates a new RolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesResponse(resourceId RolesResponseGetResourceIdArgType, resourceType RolesResponseGetResourceTypeArgType, roles RolesResponseGetRolesArgType) *RolesResponse {
	this := RolesResponse{}
	setRolesResponseGetResourceIdAttributeType(&this.ResourceId, resourceId)
	setRolesResponseGetResourceTypeAttributeType(&this.ResourceType, resourceType)
	setRolesResponseGetRolesAttributeType(&this.Roles, roles)
	return &this
}

// NewRolesResponseWithDefaults instantiates a new RolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesResponseWithDefaults() *RolesResponse {
	this := RolesResponse{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *RolesResponse) GetResourceId() (ret RolesResponseGetResourceIdRetType) {
	ret, _ = o.GetResourceIdOk()
	return ret
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *RolesResponse) GetResourceIdOk() (ret RolesResponseGetResourceIdRetType, ok bool) {
	return getRolesResponseGetResourceIdAttributeTypeOk(o.ResourceId)
}

// SetResourceId sets field value
func (o *RolesResponse) SetResourceId(v RolesResponseGetResourceIdRetType) {
	setRolesResponseGetResourceIdAttributeType(&o.ResourceId, v)
}

// GetResourceType returns the ResourceType field value
func (o *RolesResponse) GetResourceType() (ret RolesResponseGetResourceTypeRetType) {
	ret, _ = o.GetResourceTypeOk()
	return ret
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RolesResponse) GetResourceTypeOk() (ret RolesResponseGetResourceTypeRetType, ok bool) {
	return getRolesResponseGetResourceTypeAttributeTypeOk(o.ResourceType)
}

// SetResourceType sets field value
func (o *RolesResponse) SetResourceType(v RolesResponseGetResourceTypeRetType) {
	setRolesResponseGetResourceTypeAttributeType(&o.ResourceType, v)
}

// GetRoles returns the Roles field value
func (o *RolesResponse) GetRoles() (ret RolesResponseGetRolesRetType) {
	ret, _ = o.GetRolesOk()
	return ret
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RolesResponse) GetRolesOk() (ret RolesResponseGetRolesRetType, ok bool) {
	return getRolesResponseGetRolesAttributeTypeOk(o.Roles)
}

// SetRoles sets field value
func (o *RolesResponse) SetRoles(v RolesResponseGetRolesRetType) {
	setRolesResponseGetRolesAttributeType(&o.Roles, v)
}

func (o RolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRolesResponseGetResourceIdAttributeTypeOk(o.ResourceId); ok {
		toSerialize["ResourceId"] = val
	}
	if val, ok := getRolesResponseGetResourceTypeAttributeTypeOk(o.ResourceType); ok {
		toSerialize["ResourceType"] = val
	}
	if val, ok := getRolesResponseGetRolesAttributeTypeOk(o.Roles); ok {
		toSerialize["Roles"] = val
	}
	return toSerialize, nil
}

type NullableRolesResponse struct {
	value *RolesResponse
	isSet bool
}

func (v NullableRolesResponse) Get() *RolesResponse {
	return v.value
}

func (v *NullableRolesResponse) Set(val *RolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesResponse(val *RolesResponse) *NullableRolesResponse {
	return &NullableRolesResponse{value: val, isSet: true}
}

func (v NullableRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
