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

// checks if the ProjectStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectStatus{}

/*
	types and functions for project
*/

// isNotNullableString
type ProjectStatusGetProjectAttributeType = *string

func getProjectStatusGetProjectAttributeTypeOk(arg ProjectStatusGetProjectAttributeType) (ret ProjectStatusGetProjectRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectStatusGetProjectAttributeType(arg *ProjectStatusGetProjectAttributeType, val ProjectStatusGetProjectRetType) {
	*arg = &val
}

type ProjectStatusGetProjectArgType = string
type ProjectStatusGetProjectRetType = string

/*
	types and functions for scope
*/

// isEnumRef
type ProjectStatusGetScopeAttributeType = *ProjectScope
type ProjectStatusGetScopeArgType = ProjectScope
type ProjectStatusGetScopeRetType = ProjectScope

func getProjectStatusGetScopeAttributeTypeOk(arg ProjectStatusGetScopeAttributeType) (ret ProjectStatusGetScopeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectStatusGetScopeAttributeType(arg *ProjectStatusGetScopeAttributeType, val ProjectStatusGetScopeRetType) {
	*arg = &val
}

// ProjectStatus struct for ProjectStatus
type ProjectStatus struct {
	// Project ID
	// REQUIRED
	Project ProjectStatusGetProjectAttributeType `json:"project" required:"true"`
	// REQUIRED
	Scope ProjectStatusGetScopeAttributeType `json:"scope" required:"true"`
}

type _ProjectStatus ProjectStatus

// NewProjectStatus instantiates a new ProjectStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectStatus(project ProjectStatusGetProjectArgType, scope ProjectStatusGetScopeArgType) *ProjectStatus {
	this := ProjectStatus{}
	setProjectStatusGetProjectAttributeType(&this.Project, project)
	setProjectStatusGetScopeAttributeType(&this.Scope, scope)
	return &this
}

// NewProjectStatusWithDefaults instantiates a new ProjectStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectStatusWithDefaults() *ProjectStatus {
	this := ProjectStatus{}
	return &this
}

// GetProject returns the Project field value
func (o *ProjectStatus) GetProject() (ret ProjectStatusGetProjectRetType) {
	ret, _ = o.GetProjectOk()
	return ret
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProjectStatus) GetProjectOk() (ret ProjectStatusGetProjectRetType, ok bool) {
	return getProjectStatusGetProjectAttributeTypeOk(o.Project)
}

// SetProject sets field value
func (o *ProjectStatus) SetProject(v ProjectStatusGetProjectRetType) {
	setProjectStatusGetProjectAttributeType(&o.Project, v)
}

// GetScope returns the Scope field value
func (o *ProjectStatus) GetScope() (ret ProjectStatusGetScopeRetType) {
	ret, _ = o.GetScopeOk()
	return ret
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ProjectStatus) GetScopeOk() (ret ProjectStatusGetScopeRetType, ok bool) {
	return getProjectStatusGetScopeAttributeTypeOk(o.Scope)
}

// SetScope sets field value
func (o *ProjectStatus) SetScope(v ProjectStatusGetScopeRetType) {
	setProjectStatusGetScopeAttributeType(&o.Scope, v)
}

func (o ProjectStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getProjectStatusGetProjectAttributeTypeOk(o.Project); ok {
		toSerialize["Project"] = val
	}
	if val, ok := getProjectStatusGetScopeAttributeTypeOk(o.Scope); ok {
		toSerialize["Scope"] = val
	}
	return toSerialize, nil
}

type NullableProjectStatus struct {
	value *ProjectStatus
	isSet bool
}

func (v NullableProjectStatus) Get() *ProjectStatus {
	return v.value
}

func (v *NullableProjectStatus) Set(val *ProjectStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectStatus(val *ProjectStatus) *NullableProjectStatus {
	return &NullableProjectStatus{value: val, isSet: true}
}

func (v NullableProjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
