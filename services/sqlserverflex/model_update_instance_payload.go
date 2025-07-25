/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the UpdateInstancePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstancePayload{}

/*
	types and functions for acl
*/

// isModel
type UpdateInstancePayloadGetAclAttributeType = *CreateInstancePayloadAcl
type UpdateInstancePayloadGetAclArgType = CreateInstancePayloadAcl
type UpdateInstancePayloadGetAclRetType = CreateInstancePayloadAcl

func getUpdateInstancePayloadGetAclAttributeTypeOk(arg UpdateInstancePayloadGetAclAttributeType) (ret UpdateInstancePayloadGetAclRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateInstancePayloadGetAclAttributeType(arg *UpdateInstancePayloadGetAclAttributeType, val UpdateInstancePayloadGetAclRetType) {
	*arg = &val
}

/*
	types and functions for backupSchedule
*/

// isNotNullableString
type UpdateInstancePayloadGetBackupScheduleAttributeType = *string

func getUpdateInstancePayloadGetBackupScheduleAttributeTypeOk(arg UpdateInstancePayloadGetBackupScheduleAttributeType) (ret UpdateInstancePayloadGetBackupScheduleRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateInstancePayloadGetBackupScheduleAttributeType(arg *UpdateInstancePayloadGetBackupScheduleAttributeType, val UpdateInstancePayloadGetBackupScheduleRetType) {
	*arg = &val
}

type UpdateInstancePayloadGetBackupScheduleArgType = string
type UpdateInstancePayloadGetBackupScheduleRetType = string

/*
	types and functions for flavorId
*/

// isNotNullableString
type UpdateInstancePayloadGetFlavorIdAttributeType = *string

func getUpdateInstancePayloadGetFlavorIdAttributeTypeOk(arg UpdateInstancePayloadGetFlavorIdAttributeType) (ret UpdateInstancePayloadGetFlavorIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateInstancePayloadGetFlavorIdAttributeType(arg *UpdateInstancePayloadGetFlavorIdAttributeType, val UpdateInstancePayloadGetFlavorIdRetType) {
	*arg = &val
}

type UpdateInstancePayloadGetFlavorIdArgType = string
type UpdateInstancePayloadGetFlavorIdRetType = string

/*
	types and functions for labels
*/

// isFreeform
type UpdateInstancePayloadGetLabelsAttributeType = *map[string]interface{}
type UpdateInstancePayloadGetLabelsArgType = map[string]interface{}
type UpdateInstancePayloadGetLabelsRetType = map[string]interface{}

func getUpdateInstancePayloadGetLabelsAttributeTypeOk(arg UpdateInstancePayloadGetLabelsAttributeType) (ret UpdateInstancePayloadGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateInstancePayloadGetLabelsAttributeType(arg *UpdateInstancePayloadGetLabelsAttributeType, val UpdateInstancePayloadGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type UpdateInstancePayloadGetNameAttributeType = *string

func getUpdateInstancePayloadGetNameAttributeTypeOk(arg UpdateInstancePayloadGetNameAttributeType) (ret UpdateInstancePayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateInstancePayloadGetNameAttributeType(arg *UpdateInstancePayloadGetNameAttributeType, val UpdateInstancePayloadGetNameRetType) {
	*arg = &val
}

type UpdateInstancePayloadGetNameArgType = string
type UpdateInstancePayloadGetNameRetType = string

/*
	types and functions for version
*/

// isNotNullableString
type UpdateInstancePayloadGetVersionAttributeType = *string

func getUpdateInstancePayloadGetVersionAttributeTypeOk(arg UpdateInstancePayloadGetVersionAttributeType) (ret UpdateInstancePayloadGetVersionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateInstancePayloadGetVersionAttributeType(arg *UpdateInstancePayloadGetVersionAttributeType, val UpdateInstancePayloadGetVersionRetType) {
	*arg = &val
}

type UpdateInstancePayloadGetVersionArgType = string
type UpdateInstancePayloadGetVersionRetType = string

// UpdateInstancePayload struct for UpdateInstancePayload
type UpdateInstancePayload struct {
	// REQUIRED
	Acl UpdateInstancePayloadGetAclAttributeType `json:"acl" required:"true"`
	// Cronjob for the daily full backup if not provided a job will generated between 00:00 and 04:59
	// REQUIRED
	BackupSchedule UpdateInstancePayloadGetBackupScheduleAttributeType `json:"backupSchedule" required:"true"`
	// Id of the selected flavor
	// REQUIRED
	FlavorId UpdateInstancePayloadGetFlavorIdAttributeType `json:"flavorId" required:"true"`
	// REQUIRED
	Labels UpdateInstancePayloadGetLabelsAttributeType `json:"labels" required:"true"`
	// Name of the instance
	// REQUIRED
	Name UpdateInstancePayloadGetNameAttributeType `json:"name" required:"true"`
	// Version of the MSSQL Server
	// REQUIRED
	Version UpdateInstancePayloadGetVersionAttributeType `json:"version" required:"true"`
}

type _UpdateInstancePayload UpdateInstancePayload

// NewUpdateInstancePayload instantiates a new UpdateInstancePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstancePayload(acl UpdateInstancePayloadGetAclArgType, backupSchedule UpdateInstancePayloadGetBackupScheduleArgType, flavorId UpdateInstancePayloadGetFlavorIdArgType, labels UpdateInstancePayloadGetLabelsArgType, name UpdateInstancePayloadGetNameArgType, version UpdateInstancePayloadGetVersionArgType) *UpdateInstancePayload {
	this := UpdateInstancePayload{}
	setUpdateInstancePayloadGetAclAttributeType(&this.Acl, acl)
	setUpdateInstancePayloadGetBackupScheduleAttributeType(&this.BackupSchedule, backupSchedule)
	setUpdateInstancePayloadGetFlavorIdAttributeType(&this.FlavorId, flavorId)
	setUpdateInstancePayloadGetLabelsAttributeType(&this.Labels, labels)
	setUpdateInstancePayloadGetNameAttributeType(&this.Name, name)
	setUpdateInstancePayloadGetVersionAttributeType(&this.Version, version)
	return &this
}

// NewUpdateInstancePayloadWithDefaults instantiates a new UpdateInstancePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstancePayloadWithDefaults() *UpdateInstancePayload {
	this := UpdateInstancePayload{}
	var version string = "2022"
	this.Version = &version
	return &this
}

// GetAcl returns the Acl field value
func (o *UpdateInstancePayload) GetAcl() (ret UpdateInstancePayloadGetAclRetType) {
	ret, _ = o.GetAclOk()
	return ret
}

// GetAclOk returns a tuple with the Acl field value
// and a boolean to check if the value has been set.
func (o *UpdateInstancePayload) GetAclOk() (ret UpdateInstancePayloadGetAclRetType, ok bool) {
	return getUpdateInstancePayloadGetAclAttributeTypeOk(o.Acl)
}

// SetAcl sets field value
func (o *UpdateInstancePayload) SetAcl(v UpdateInstancePayloadGetAclRetType) {
	setUpdateInstancePayloadGetAclAttributeType(&o.Acl, v)
}

// GetBackupSchedule returns the BackupSchedule field value
func (o *UpdateInstancePayload) GetBackupSchedule() (ret UpdateInstancePayloadGetBackupScheduleRetType) {
	ret, _ = o.GetBackupScheduleOk()
	return ret
}

// GetBackupScheduleOk returns a tuple with the BackupSchedule field value
// and a boolean to check if the value has been set.
func (o *UpdateInstancePayload) GetBackupScheduleOk() (ret UpdateInstancePayloadGetBackupScheduleRetType, ok bool) {
	return getUpdateInstancePayloadGetBackupScheduleAttributeTypeOk(o.BackupSchedule)
}

// SetBackupSchedule sets field value
func (o *UpdateInstancePayload) SetBackupSchedule(v UpdateInstancePayloadGetBackupScheduleRetType) {
	setUpdateInstancePayloadGetBackupScheduleAttributeType(&o.BackupSchedule, v)
}

// GetFlavorId returns the FlavorId field value
func (o *UpdateInstancePayload) GetFlavorId() (ret UpdateInstancePayloadGetFlavorIdRetType) {
	ret, _ = o.GetFlavorIdOk()
	return ret
}

// GetFlavorIdOk returns a tuple with the FlavorId field value
// and a boolean to check if the value has been set.
func (o *UpdateInstancePayload) GetFlavorIdOk() (ret UpdateInstancePayloadGetFlavorIdRetType, ok bool) {
	return getUpdateInstancePayloadGetFlavorIdAttributeTypeOk(o.FlavorId)
}

// SetFlavorId sets field value
func (o *UpdateInstancePayload) SetFlavorId(v UpdateInstancePayloadGetFlavorIdRetType) {
	setUpdateInstancePayloadGetFlavorIdAttributeType(&o.FlavorId, v)
}

// GetLabels returns the Labels field value
func (o *UpdateInstancePayload) GetLabels() (ret UpdateInstancePayloadGetLabelsRetType) {
	ret, _ = o.GetLabelsOk()
	return ret
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *UpdateInstancePayload) GetLabelsOk() (ret UpdateInstancePayloadGetLabelsRetType, ok bool) {
	return getUpdateInstancePayloadGetLabelsAttributeTypeOk(o.Labels)
}

// SetLabels sets field value
func (o *UpdateInstancePayload) SetLabels(v UpdateInstancePayloadGetLabelsRetType) {
	setUpdateInstancePayloadGetLabelsAttributeType(&o.Labels, v)
}

// GetName returns the Name field value
func (o *UpdateInstancePayload) GetName() (ret UpdateInstancePayloadGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateInstancePayload) GetNameOk() (ret UpdateInstancePayloadGetNameRetType, ok bool) {
	return getUpdateInstancePayloadGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *UpdateInstancePayload) SetName(v UpdateInstancePayloadGetNameRetType) {
	setUpdateInstancePayloadGetNameAttributeType(&o.Name, v)
}

// GetVersion returns the Version field value
func (o *UpdateInstancePayload) GetVersion() (ret UpdateInstancePayloadGetVersionRetType) {
	ret, _ = o.GetVersionOk()
	return ret
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateInstancePayload) GetVersionOk() (ret UpdateInstancePayloadGetVersionRetType, ok bool) {
	return getUpdateInstancePayloadGetVersionAttributeTypeOk(o.Version)
}

// SetVersion sets field value
func (o *UpdateInstancePayload) SetVersion(v UpdateInstancePayloadGetVersionRetType) {
	setUpdateInstancePayloadGetVersionAttributeType(&o.Version, v)
}

func (o UpdateInstancePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateInstancePayloadGetAclAttributeTypeOk(o.Acl); ok {
		toSerialize["Acl"] = val
	}
	if val, ok := getUpdateInstancePayloadGetBackupScheduleAttributeTypeOk(o.BackupSchedule); ok {
		toSerialize["BackupSchedule"] = val
	}
	if val, ok := getUpdateInstancePayloadGetFlavorIdAttributeTypeOk(o.FlavorId); ok {
		toSerialize["FlavorId"] = val
	}
	if val, ok := getUpdateInstancePayloadGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getUpdateInstancePayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getUpdateInstancePayloadGetVersionAttributeTypeOk(o.Version); ok {
		toSerialize["Version"] = val
	}
	return toSerialize, nil
}

type NullableUpdateInstancePayload struct {
	value *UpdateInstancePayload
	isSet bool
}

func (v NullableUpdateInstancePayload) Get() *UpdateInstancePayload {
	return v.value
}

func (v *NullableUpdateInstancePayload) Set(val *UpdateInstancePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInstancePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInstancePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInstancePayload(val *UpdateInstancePayload) *NullableUpdateInstancePayload {
	return &NullableUpdateInstancePayload{value: val, isSet: true}
}

func (v NullableUpdateInstancePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInstancePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
