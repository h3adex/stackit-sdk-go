/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

import (
	"encoding/json"
	"fmt"
)

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

/*
	types and functions for cfGuid
*/

// isNotNullableString
type InstanceGetCfGuidAttributeType = *string

func getInstanceGetCfGuidAttributeTypeOk(arg InstanceGetCfGuidAttributeType) (ret InstanceGetCfGuidRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetCfGuidAttributeType(arg *InstanceGetCfGuidAttributeType, val InstanceGetCfGuidRetType) {
	*arg = &val
}

type InstanceGetCfGuidArgType = string
type InstanceGetCfGuidRetType = string

/*
	types and functions for cfOrganizationGuid
*/

// isNotNullableString
type InstanceGetCfOrganizationGuidAttributeType = *string

func getInstanceGetCfOrganizationGuidAttributeTypeOk(arg InstanceGetCfOrganizationGuidAttributeType) (ret InstanceGetCfOrganizationGuidRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetCfOrganizationGuidAttributeType(arg *InstanceGetCfOrganizationGuidAttributeType, val InstanceGetCfOrganizationGuidRetType) {
	*arg = &val
}

type InstanceGetCfOrganizationGuidArgType = string
type InstanceGetCfOrganizationGuidRetType = string

/*
	types and functions for cfSpaceGuid
*/

// isNotNullableString
type InstanceGetCfSpaceGuidAttributeType = *string

func getInstanceGetCfSpaceGuidAttributeTypeOk(arg InstanceGetCfSpaceGuidAttributeType) (ret InstanceGetCfSpaceGuidRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetCfSpaceGuidAttributeType(arg *InstanceGetCfSpaceGuidAttributeType, val InstanceGetCfSpaceGuidRetType) {
	*arg = &val
}

type InstanceGetCfSpaceGuidArgType = string
type InstanceGetCfSpaceGuidRetType = string

/*
	types and functions for dashboardUrl
*/

// isNotNullableString
type InstanceGetDashboardUrlAttributeType = *string

func getInstanceGetDashboardUrlAttributeTypeOk(arg InstanceGetDashboardUrlAttributeType) (ret InstanceGetDashboardUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetDashboardUrlAttributeType(arg *InstanceGetDashboardUrlAttributeType, val InstanceGetDashboardUrlRetType) {
	*arg = &val
}

type InstanceGetDashboardUrlArgType = string
type InstanceGetDashboardUrlRetType = string

/*
	types and functions for imageUrl
*/

// isNotNullableString
type InstanceGetImageUrlAttributeType = *string

func getInstanceGetImageUrlAttributeTypeOk(arg InstanceGetImageUrlAttributeType) (ret InstanceGetImageUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetImageUrlAttributeType(arg *InstanceGetImageUrlAttributeType, val InstanceGetImageUrlRetType) {
	*arg = &val
}

type InstanceGetImageUrlArgType = string
type InstanceGetImageUrlRetType = string

/*
	types and functions for instanceId
*/

// isNotNullableString
type InstanceGetInstanceIdAttributeType = *string

func getInstanceGetInstanceIdAttributeTypeOk(arg InstanceGetInstanceIdAttributeType) (ret InstanceGetInstanceIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetInstanceIdAttributeType(arg *InstanceGetInstanceIdAttributeType, val InstanceGetInstanceIdRetType) {
	*arg = &val
}

type InstanceGetInstanceIdArgType = string
type InstanceGetInstanceIdRetType = string

/*
	types and functions for lastOperation
*/

// isModel
type InstanceGetLastOperationAttributeType = *InstanceLastOperation
type InstanceGetLastOperationArgType = InstanceLastOperation
type InstanceGetLastOperationRetType = InstanceLastOperation

func getInstanceGetLastOperationAttributeTypeOk(arg InstanceGetLastOperationAttributeType) (ret InstanceGetLastOperationRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetLastOperationAttributeType(arg *InstanceGetLastOperationAttributeType, val InstanceGetLastOperationRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type InstanceGetNameAttributeType = *string

func getInstanceGetNameAttributeTypeOk(arg InstanceGetNameAttributeType) (ret InstanceGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetNameAttributeType(arg *InstanceGetNameAttributeType, val InstanceGetNameRetType) {
	*arg = &val
}

type InstanceGetNameArgType = string
type InstanceGetNameRetType = string

/*
	types and functions for offeringName
*/

// isNotNullableString
type InstanceGetOfferingNameAttributeType = *string

func getInstanceGetOfferingNameAttributeTypeOk(arg InstanceGetOfferingNameAttributeType) (ret InstanceGetOfferingNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetOfferingNameAttributeType(arg *InstanceGetOfferingNameAttributeType, val InstanceGetOfferingNameRetType) {
	*arg = &val
}

type InstanceGetOfferingNameArgType = string
type InstanceGetOfferingNameRetType = string

/*
	types and functions for offeringVersion
*/

// isNotNullableString
type InstanceGetOfferingVersionAttributeType = *string

func getInstanceGetOfferingVersionAttributeTypeOk(arg InstanceGetOfferingVersionAttributeType) (ret InstanceGetOfferingVersionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetOfferingVersionAttributeType(arg *InstanceGetOfferingVersionAttributeType, val InstanceGetOfferingVersionRetType) {
	*arg = &val
}

type InstanceGetOfferingVersionArgType = string
type InstanceGetOfferingVersionRetType = string

/*
	types and functions for parameters
*/

// isFreeform
type InstanceGetParametersAttributeType = *map[string]interface{}
type InstanceGetParametersArgType = map[string]interface{}
type InstanceGetParametersRetType = map[string]interface{}

func getInstanceGetParametersAttributeTypeOk(arg InstanceGetParametersAttributeType) (ret InstanceGetParametersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetParametersAttributeType(arg *InstanceGetParametersAttributeType, val InstanceGetParametersRetType) {
	*arg = &val
}

/*
	types and functions for planId
*/

// isNotNullableString
type InstanceGetPlanIdAttributeType = *string

func getInstanceGetPlanIdAttributeTypeOk(arg InstanceGetPlanIdAttributeType) (ret InstanceGetPlanIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetPlanIdAttributeType(arg *InstanceGetPlanIdAttributeType, val InstanceGetPlanIdRetType) {
	*arg = &val
}

type InstanceGetPlanIdArgType = string
type InstanceGetPlanIdRetType = string

/*
	types and functions for planName
*/

// isNotNullableString
type InstanceGetPlanNameAttributeType = *string

func getInstanceGetPlanNameAttributeTypeOk(arg InstanceGetPlanNameAttributeType) (ret InstanceGetPlanNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetPlanNameAttributeType(arg *InstanceGetPlanNameAttributeType, val InstanceGetPlanNameRetType) {
	*arg = &val
}

type InstanceGetPlanNameArgType = string
type InstanceGetPlanNameRetType = string

/*
	types and functions for status
*/

// isEnum

// InstanceStatus the model 'Instance'
// value type for enums
type InstanceStatus string

// List of Status
const (
	INSTANCESTATUS_ACTIVE   InstanceStatus = "active"
	INSTANCESTATUS_FAILED   InstanceStatus = "failed"
	INSTANCESTATUS_STOPPED  InstanceStatus = "stopped"
	INSTANCESTATUS_CREATING InstanceStatus = "creating"
	INSTANCESTATUS_DELETING InstanceStatus = "deleting"
	INSTANCESTATUS_UPDATING InstanceStatus = "updating"
)

// All allowed values of Instance enum
var AllowedInstanceStatusEnumValues = []InstanceStatus{
	"active",
	"failed",
	"stopped",
	"creating",
	"deleting",
	"updating",
}

func (v *InstanceStatus) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson InstanceStatus
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
	enumTypeValue := InstanceStatus(value)
	for _, existing := range AllowedInstanceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Instance", value)
}

// NewInstanceStatusFromValue returns a pointer to a valid InstanceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceStatusFromValue(v InstanceStatus) (*InstanceStatus, error) {
	ev := InstanceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceStatus: valid values are %v", v, AllowedInstanceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceStatus) IsValid() bool {
	for _, existing := range AllowedInstanceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusStatus value
func (v InstanceStatus) Ptr() *InstanceStatus {
	return &v
}

type NullableInstanceStatus struct {
	value *InstanceStatus
	isSet bool
}

func (v NullableInstanceStatus) Get() *InstanceStatus {
	return v.value
}

func (v *NullableInstanceStatus) Set(val *InstanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceStatus(val *InstanceStatus) *NullableInstanceStatus {
	return &NullableInstanceStatus{value: val, isSet: true}
}

func (v NullableInstanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type InstanceGetStatusAttributeType = *InstanceStatus
type InstanceGetStatusArgType = InstanceStatus
type InstanceGetStatusRetType = InstanceStatus

func getInstanceGetStatusAttributeTypeOk(arg InstanceGetStatusAttributeType) (ret InstanceGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceGetStatusAttributeType(arg *InstanceGetStatusAttributeType, val InstanceGetStatusRetType) {
	*arg = &val
}

// Instance struct for Instance
type Instance struct {
	// REQUIRED
	CfGuid InstanceGetCfGuidAttributeType `json:"cfGuid" required:"true"`
	// REQUIRED
	CfOrganizationGuid InstanceGetCfOrganizationGuidAttributeType `json:"cfOrganizationGuid" required:"true"`
	// REQUIRED
	CfSpaceGuid InstanceGetCfSpaceGuidAttributeType `json:"cfSpaceGuid" required:"true"`
	// REQUIRED
	DashboardUrl InstanceGetDashboardUrlAttributeType `json:"dashboardUrl" required:"true"`
	// REQUIRED
	ImageUrl   InstanceGetImageUrlAttributeType   `json:"imageUrl" required:"true"`
	InstanceId InstanceGetInstanceIdAttributeType `json:"instanceId,omitempty"`
	// REQUIRED
	LastOperation InstanceGetLastOperationAttributeType `json:"lastOperation" required:"true"`
	// REQUIRED
	Name InstanceGetNameAttributeType `json:"name" required:"true"`
	// Deprecated: Check the GitHub changelog for alternatives
	// REQUIRED
	OfferingName InstanceGetOfferingNameAttributeType `json:"offeringName" required:"true"`
	// REQUIRED
	OfferingVersion InstanceGetOfferingVersionAttributeType `json:"offeringVersion" required:"true"`
	// REQUIRED
	Parameters InstanceGetParametersAttributeType `json:"parameters" required:"true"`
	// REQUIRED
	PlanId InstanceGetPlanIdAttributeType `json:"planId" required:"true"`
	// REQUIRED
	PlanName InstanceGetPlanNameAttributeType `json:"planName" required:"true"`
	Status   InstanceGetStatusAttributeType   `json:"status,omitempty"`
}

type _Instance Instance

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(cfGuid InstanceGetCfGuidArgType, cfOrganizationGuid InstanceGetCfOrganizationGuidArgType, cfSpaceGuid InstanceGetCfSpaceGuidArgType, dashboardUrl InstanceGetDashboardUrlArgType, imageUrl InstanceGetImageUrlArgType, lastOperation InstanceGetLastOperationArgType, name InstanceGetNameArgType, offeringName InstanceGetOfferingNameArgType, offeringVersion InstanceGetOfferingVersionArgType, parameters InstanceGetParametersArgType, planId InstanceGetPlanIdArgType, planName InstanceGetPlanNameArgType) *Instance {
	this := Instance{}
	setInstanceGetCfGuidAttributeType(&this.CfGuid, cfGuid)
	setInstanceGetCfOrganizationGuidAttributeType(&this.CfOrganizationGuid, cfOrganizationGuid)
	setInstanceGetCfSpaceGuidAttributeType(&this.CfSpaceGuid, cfSpaceGuid)
	setInstanceGetDashboardUrlAttributeType(&this.DashboardUrl, dashboardUrl)
	setInstanceGetImageUrlAttributeType(&this.ImageUrl, imageUrl)
	setInstanceGetLastOperationAttributeType(&this.LastOperation, lastOperation)
	setInstanceGetNameAttributeType(&this.Name, name)
	setInstanceGetOfferingNameAttributeType(&this.OfferingName, offeringName)
	setInstanceGetOfferingVersionAttributeType(&this.OfferingVersion, offeringVersion)
	setInstanceGetParametersAttributeType(&this.Parameters, parameters)
	setInstanceGetPlanIdAttributeType(&this.PlanId, planId)
	setInstanceGetPlanNameAttributeType(&this.PlanName, planName)
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetCfGuid returns the CfGuid field value
func (o *Instance) GetCfGuid() (ret InstanceGetCfGuidRetType) {
	ret, _ = o.GetCfGuidOk()
	return ret
}

// GetCfGuidOk returns a tuple with the CfGuid field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCfGuidOk() (ret InstanceGetCfGuidRetType, ok bool) {
	return getInstanceGetCfGuidAttributeTypeOk(o.CfGuid)
}

// SetCfGuid sets field value
func (o *Instance) SetCfGuid(v InstanceGetCfGuidRetType) {
	setInstanceGetCfGuidAttributeType(&o.CfGuid, v)
}

// GetCfOrganizationGuid returns the CfOrganizationGuid field value
func (o *Instance) GetCfOrganizationGuid() (ret InstanceGetCfOrganizationGuidRetType) {
	ret, _ = o.GetCfOrganizationGuidOk()
	return ret
}

// GetCfOrganizationGuidOk returns a tuple with the CfOrganizationGuid field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCfOrganizationGuidOk() (ret InstanceGetCfOrganizationGuidRetType, ok bool) {
	return getInstanceGetCfOrganizationGuidAttributeTypeOk(o.CfOrganizationGuid)
}

// SetCfOrganizationGuid sets field value
func (o *Instance) SetCfOrganizationGuid(v InstanceGetCfOrganizationGuidRetType) {
	setInstanceGetCfOrganizationGuidAttributeType(&o.CfOrganizationGuid, v)
}

// GetCfSpaceGuid returns the CfSpaceGuid field value
func (o *Instance) GetCfSpaceGuid() (ret InstanceGetCfSpaceGuidRetType) {
	ret, _ = o.GetCfSpaceGuidOk()
	return ret
}

// GetCfSpaceGuidOk returns a tuple with the CfSpaceGuid field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCfSpaceGuidOk() (ret InstanceGetCfSpaceGuidRetType, ok bool) {
	return getInstanceGetCfSpaceGuidAttributeTypeOk(o.CfSpaceGuid)
}

// SetCfSpaceGuid sets field value
func (o *Instance) SetCfSpaceGuid(v InstanceGetCfSpaceGuidRetType) {
	setInstanceGetCfSpaceGuidAttributeType(&o.CfSpaceGuid, v)
}

// GetDashboardUrl returns the DashboardUrl field value
func (o *Instance) GetDashboardUrl() (ret InstanceGetDashboardUrlRetType) {
	ret, _ = o.GetDashboardUrlOk()
	return ret
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value
// and a boolean to check if the value has been set.
func (o *Instance) GetDashboardUrlOk() (ret InstanceGetDashboardUrlRetType, ok bool) {
	return getInstanceGetDashboardUrlAttributeTypeOk(o.DashboardUrl)
}

// SetDashboardUrl sets field value
func (o *Instance) SetDashboardUrl(v InstanceGetDashboardUrlRetType) {
	setInstanceGetDashboardUrlAttributeType(&o.DashboardUrl, v)
}

// GetImageUrl returns the ImageUrl field value
func (o *Instance) GetImageUrl() (ret InstanceGetImageUrlRetType) {
	ret, _ = o.GetImageUrlOk()
	return ret
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *Instance) GetImageUrlOk() (ret InstanceGetImageUrlRetType, ok bool) {
	return getInstanceGetImageUrlAttributeTypeOk(o.ImageUrl)
}

// SetImageUrl sets field value
func (o *Instance) SetImageUrl(v InstanceGetImageUrlRetType) {
	setInstanceGetImageUrlAttributeType(&o.ImageUrl, v)
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *Instance) GetInstanceId() (res InstanceGetInstanceIdRetType) {
	res, _ = o.GetInstanceIdOk()
	return
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetInstanceIdOk() (ret InstanceGetInstanceIdRetType, ok bool) {
	return getInstanceGetInstanceIdAttributeTypeOk(o.InstanceId)
}

// HasInstanceId returns a boolean if a field has been set.
func (o *Instance) HasInstanceId() bool {
	_, ok := o.GetInstanceIdOk()
	return ok
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *Instance) SetInstanceId(v InstanceGetInstanceIdRetType) {
	setInstanceGetInstanceIdAttributeType(&o.InstanceId, v)
}

// GetLastOperation returns the LastOperation field value
func (o *Instance) GetLastOperation() (ret InstanceGetLastOperationRetType) {
	ret, _ = o.GetLastOperationOk()
	return ret
}

// GetLastOperationOk returns a tuple with the LastOperation field value
// and a boolean to check if the value has been set.
func (o *Instance) GetLastOperationOk() (ret InstanceGetLastOperationRetType, ok bool) {
	return getInstanceGetLastOperationAttributeTypeOk(o.LastOperation)
}

// SetLastOperation sets field value
func (o *Instance) SetLastOperation(v InstanceGetLastOperationRetType) {
	setInstanceGetLastOperationAttributeType(&o.LastOperation, v)
}

// GetName returns the Name field value
func (o *Instance) GetName() (ret InstanceGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNameOk() (ret InstanceGetNameRetType, ok bool) {
	return getInstanceGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *Instance) SetName(v InstanceGetNameRetType) {
	setInstanceGetNameAttributeType(&o.Name, v)
}

// GetOfferingName returns the OfferingName field value
// Deprecated
func (o *Instance) GetOfferingName() (ret InstanceGetOfferingNameRetType) {
	ret, _ = o.GetOfferingNameOk()
	return ret
}

// GetOfferingNameOk returns a tuple with the OfferingName field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Instance) GetOfferingNameOk() (ret InstanceGetOfferingNameRetType, ok bool) {
	return getInstanceGetOfferingNameAttributeTypeOk(o.OfferingName)
}

// SetOfferingName sets field value
// Deprecated
func (o *Instance) SetOfferingName(v InstanceGetOfferingNameRetType) {
	setInstanceGetOfferingNameAttributeType(&o.OfferingName, v)
}

// GetOfferingVersion returns the OfferingVersion field value
func (o *Instance) GetOfferingVersion() (ret InstanceGetOfferingVersionRetType) {
	ret, _ = o.GetOfferingVersionOk()
	return ret
}

// GetOfferingVersionOk returns a tuple with the OfferingVersion field value
// and a boolean to check if the value has been set.
func (o *Instance) GetOfferingVersionOk() (ret InstanceGetOfferingVersionRetType, ok bool) {
	return getInstanceGetOfferingVersionAttributeTypeOk(o.OfferingVersion)
}

// SetOfferingVersion sets field value
func (o *Instance) SetOfferingVersion(v InstanceGetOfferingVersionRetType) {
	setInstanceGetOfferingVersionAttributeType(&o.OfferingVersion, v)
}

// GetParameters returns the Parameters field value
func (o *Instance) GetParameters() (ret InstanceGetParametersRetType) {
	ret, _ = o.GetParametersOk()
	return ret
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *Instance) GetParametersOk() (ret InstanceGetParametersRetType, ok bool) {
	return getInstanceGetParametersAttributeTypeOk(o.Parameters)
}

// SetParameters sets field value
func (o *Instance) SetParameters(v InstanceGetParametersRetType) {
	setInstanceGetParametersAttributeType(&o.Parameters, v)
}

// GetPlanId returns the PlanId field value
func (o *Instance) GetPlanId() (ret InstanceGetPlanIdRetType) {
	ret, _ = o.GetPlanIdOk()
	return ret
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPlanIdOk() (ret InstanceGetPlanIdRetType, ok bool) {
	return getInstanceGetPlanIdAttributeTypeOk(o.PlanId)
}

// SetPlanId sets field value
func (o *Instance) SetPlanId(v InstanceGetPlanIdRetType) {
	setInstanceGetPlanIdAttributeType(&o.PlanId, v)
}

// GetPlanName returns the PlanName field value
func (o *Instance) GetPlanName() (ret InstanceGetPlanNameRetType) {
	ret, _ = o.GetPlanNameOk()
	return ret
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPlanNameOk() (ret InstanceGetPlanNameRetType, ok bool) {
	return getInstanceGetPlanNameAttributeTypeOk(o.PlanName)
}

// SetPlanName sets field value
func (o *Instance) SetPlanName(v InstanceGetPlanNameRetType) {
	setInstanceGetPlanNameAttributeType(&o.PlanName, v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Instance) GetStatus() (res InstanceGetStatusRetType) {
	res, _ = o.GetStatusOk()
	return
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetStatusOk() (ret InstanceGetStatusRetType, ok bool) {
	return getInstanceGetStatusAttributeTypeOk(o.Status)
}

// HasStatus returns a boolean if a field has been set.
func (o *Instance) HasStatus() bool {
	_, ok := o.GetStatusOk()
	return ok
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Instance) SetStatus(v InstanceGetStatusRetType) {
	setInstanceGetStatusAttributeType(&o.Status, v)
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getInstanceGetCfGuidAttributeTypeOk(o.CfGuid); ok {
		toSerialize["CfGuid"] = val
	}
	if val, ok := getInstanceGetCfOrganizationGuidAttributeTypeOk(o.CfOrganizationGuid); ok {
		toSerialize["CfOrganizationGuid"] = val
	}
	if val, ok := getInstanceGetCfSpaceGuidAttributeTypeOk(o.CfSpaceGuid); ok {
		toSerialize["CfSpaceGuid"] = val
	}
	if val, ok := getInstanceGetDashboardUrlAttributeTypeOk(o.DashboardUrl); ok {
		toSerialize["DashboardUrl"] = val
	}
	if val, ok := getInstanceGetImageUrlAttributeTypeOk(o.ImageUrl); ok {
		toSerialize["ImageUrl"] = val
	}
	if val, ok := getInstanceGetInstanceIdAttributeTypeOk(o.InstanceId); ok {
		toSerialize["InstanceId"] = val
	}
	if val, ok := getInstanceGetLastOperationAttributeTypeOk(o.LastOperation); ok {
		toSerialize["LastOperation"] = val
	}
	if val, ok := getInstanceGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getInstanceGetOfferingNameAttributeTypeOk(o.OfferingName); ok {
		toSerialize["OfferingName"] = val
	}
	if val, ok := getInstanceGetOfferingVersionAttributeTypeOk(o.OfferingVersion); ok {
		toSerialize["OfferingVersion"] = val
	}
	if val, ok := getInstanceGetParametersAttributeTypeOk(o.Parameters); ok {
		toSerialize["Parameters"] = val
	}
	if val, ok := getInstanceGetPlanIdAttributeTypeOk(o.PlanId); ok {
		toSerialize["PlanId"] = val
	}
	if val, ok := getInstanceGetPlanNameAttributeTypeOk(o.PlanName); ok {
		toSerialize["PlanName"] = val
	}
	if val, ok := getInstanceGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	return toSerialize, nil
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
