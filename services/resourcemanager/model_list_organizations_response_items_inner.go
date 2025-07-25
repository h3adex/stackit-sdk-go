/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"encoding/json"
	"time"
)

// checks if the ListOrganizationsResponseItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationsResponseItemsInner{}

/*
	types and functions for containerId
*/

// isNotNullableString
type ListOrganizationsResponseItemsInnerGetContainerIdAttributeType = *string

func getListOrganizationsResponseItemsInnerGetContainerIdAttributeTypeOk(arg ListOrganizationsResponseItemsInnerGetContainerIdAttributeType) (ret ListOrganizationsResponseItemsInnerGetContainerIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOrganizationsResponseItemsInnerGetContainerIdAttributeType(arg *ListOrganizationsResponseItemsInnerGetContainerIdAttributeType, val ListOrganizationsResponseItemsInnerGetContainerIdRetType) {
	*arg = &val
}

type ListOrganizationsResponseItemsInnerGetContainerIdArgType = string
type ListOrganizationsResponseItemsInnerGetContainerIdRetType = string

/*
	types and functions for creationTime
*/

// isDateTime
type ListOrganizationsResponseItemsInnerGetCreationTimeAttributeType = *time.Time
type ListOrganizationsResponseItemsInnerGetCreationTimeArgType = time.Time
type ListOrganizationsResponseItemsInnerGetCreationTimeRetType = time.Time

func getListOrganizationsResponseItemsInnerGetCreationTimeAttributeTypeOk(arg ListOrganizationsResponseItemsInnerGetCreationTimeAttributeType) (ret ListOrganizationsResponseItemsInnerGetCreationTimeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOrganizationsResponseItemsInnerGetCreationTimeAttributeType(arg *ListOrganizationsResponseItemsInnerGetCreationTimeAttributeType, val ListOrganizationsResponseItemsInnerGetCreationTimeRetType) {
	*arg = &val
}

/*
	types and functions for labels
*/

// isContainer
type ListOrganizationsResponseItemsInnerGetLabelsAttributeType = *map[string]string
type ListOrganizationsResponseItemsInnerGetLabelsArgType = map[string]string
type ListOrganizationsResponseItemsInnerGetLabelsRetType = map[string]string

func getListOrganizationsResponseItemsInnerGetLabelsAttributeTypeOk(arg ListOrganizationsResponseItemsInnerGetLabelsAttributeType) (ret ListOrganizationsResponseItemsInnerGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOrganizationsResponseItemsInnerGetLabelsAttributeType(arg *ListOrganizationsResponseItemsInnerGetLabelsAttributeType, val ListOrganizationsResponseItemsInnerGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for lifecycleState
*/

// isEnumRef
type ListOrganizationsResponseItemsInnerGetLifecycleStateAttributeType = *LifecycleState
type ListOrganizationsResponseItemsInnerGetLifecycleStateArgType = LifecycleState
type ListOrganizationsResponseItemsInnerGetLifecycleStateRetType = LifecycleState

func getListOrganizationsResponseItemsInnerGetLifecycleStateAttributeTypeOk(arg ListOrganizationsResponseItemsInnerGetLifecycleStateAttributeType) (ret ListOrganizationsResponseItemsInnerGetLifecycleStateRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOrganizationsResponseItemsInnerGetLifecycleStateAttributeType(arg *ListOrganizationsResponseItemsInnerGetLifecycleStateAttributeType, val ListOrganizationsResponseItemsInnerGetLifecycleStateRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type ListOrganizationsResponseItemsInnerGetNameAttributeType = *string

func getListOrganizationsResponseItemsInnerGetNameAttributeTypeOk(arg ListOrganizationsResponseItemsInnerGetNameAttributeType) (ret ListOrganizationsResponseItemsInnerGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOrganizationsResponseItemsInnerGetNameAttributeType(arg *ListOrganizationsResponseItemsInnerGetNameAttributeType, val ListOrganizationsResponseItemsInnerGetNameRetType) {
	*arg = &val
}

type ListOrganizationsResponseItemsInnerGetNameArgType = string
type ListOrganizationsResponseItemsInnerGetNameRetType = string

/*
	types and functions for organizationId
*/

// isNotNullableString
type ListOrganizationsResponseItemsInnerGetOrganizationIdAttributeType = *string

func getListOrganizationsResponseItemsInnerGetOrganizationIdAttributeTypeOk(arg ListOrganizationsResponseItemsInnerGetOrganizationIdAttributeType) (ret ListOrganizationsResponseItemsInnerGetOrganizationIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOrganizationsResponseItemsInnerGetOrganizationIdAttributeType(arg *ListOrganizationsResponseItemsInnerGetOrganizationIdAttributeType, val ListOrganizationsResponseItemsInnerGetOrganizationIdRetType) {
	*arg = &val
}

type ListOrganizationsResponseItemsInnerGetOrganizationIdArgType = string
type ListOrganizationsResponseItemsInnerGetOrganizationIdRetType = string

/*
	types and functions for updateTime
*/

// isDateTime
type ListOrganizationsResponseItemsInnerGetUpdateTimeAttributeType = *time.Time
type ListOrganizationsResponseItemsInnerGetUpdateTimeArgType = time.Time
type ListOrganizationsResponseItemsInnerGetUpdateTimeRetType = time.Time

func getListOrganizationsResponseItemsInnerGetUpdateTimeAttributeTypeOk(arg ListOrganizationsResponseItemsInnerGetUpdateTimeAttributeType) (ret ListOrganizationsResponseItemsInnerGetUpdateTimeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOrganizationsResponseItemsInnerGetUpdateTimeAttributeType(arg *ListOrganizationsResponseItemsInnerGetUpdateTimeAttributeType, val ListOrganizationsResponseItemsInnerGetUpdateTimeRetType) {
	*arg = &val
}

// ListOrganizationsResponseItemsInner struct for ListOrganizationsResponseItemsInner
type ListOrganizationsResponseItemsInner struct {
	// Globally unique, user-friendly identifier.
	// REQUIRED
	ContainerId ListOrganizationsResponseItemsInnerGetContainerIdAttributeType `json:"containerId" required:"true"`
	// Timestamp at which the organization was created.
	// REQUIRED
	CreationTime ListOrganizationsResponseItemsInnerGetCreationTimeAttributeType `json:"creationTime" required:"true"`
	// Labels are key-value string pairs that can be attached to a resource container. Some labels may be enforced via policies.  - A label key must match the regex `[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`. - A label value must match the regex `^$|[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`.
	Labels ListOrganizationsResponseItemsInnerGetLabelsAttributeType `json:"labels,omitempty"`
	// REQUIRED
	LifecycleState ListOrganizationsResponseItemsInnerGetLifecycleStateAttributeType `json:"lifecycleState" required:"true"`
	// Name of the organization.
	// REQUIRED
	Name ListOrganizationsResponseItemsInnerGetNameAttributeType `json:"name" required:"true"`
	// Globally unique, organization identifier.
	// REQUIRED
	OrganizationId ListOrganizationsResponseItemsInnerGetOrganizationIdAttributeType `json:"organizationId" required:"true"`
	// Timestamp at which the organization was last modified.
	// REQUIRED
	UpdateTime ListOrganizationsResponseItemsInnerGetUpdateTimeAttributeType `json:"updateTime" required:"true"`
}

type _ListOrganizationsResponseItemsInner ListOrganizationsResponseItemsInner

// NewListOrganizationsResponseItemsInner instantiates a new ListOrganizationsResponseItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationsResponseItemsInner(containerId ListOrganizationsResponseItemsInnerGetContainerIdArgType, creationTime ListOrganizationsResponseItemsInnerGetCreationTimeArgType, lifecycleState ListOrganizationsResponseItemsInnerGetLifecycleStateArgType, name ListOrganizationsResponseItemsInnerGetNameArgType, organizationId ListOrganizationsResponseItemsInnerGetOrganizationIdArgType, updateTime ListOrganizationsResponseItemsInnerGetUpdateTimeArgType) *ListOrganizationsResponseItemsInner {
	this := ListOrganizationsResponseItemsInner{}
	setListOrganizationsResponseItemsInnerGetContainerIdAttributeType(&this.ContainerId, containerId)
	setListOrganizationsResponseItemsInnerGetCreationTimeAttributeType(&this.CreationTime, creationTime)
	setListOrganizationsResponseItemsInnerGetLifecycleStateAttributeType(&this.LifecycleState, lifecycleState)
	setListOrganizationsResponseItemsInnerGetNameAttributeType(&this.Name, name)
	setListOrganizationsResponseItemsInnerGetOrganizationIdAttributeType(&this.OrganizationId, organizationId)
	setListOrganizationsResponseItemsInnerGetUpdateTimeAttributeType(&this.UpdateTime, updateTime)
	return &this
}

// NewListOrganizationsResponseItemsInnerWithDefaults instantiates a new ListOrganizationsResponseItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationsResponseItemsInnerWithDefaults() *ListOrganizationsResponseItemsInner {
	this := ListOrganizationsResponseItemsInner{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *ListOrganizationsResponseItemsInner) GetContainerId() (ret ListOrganizationsResponseItemsInnerGetContainerIdRetType) {
	ret, _ = o.GetContainerIdOk()
	return ret
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponseItemsInner) GetContainerIdOk() (ret ListOrganizationsResponseItemsInnerGetContainerIdRetType, ok bool) {
	return getListOrganizationsResponseItemsInnerGetContainerIdAttributeTypeOk(o.ContainerId)
}

// SetContainerId sets field value
func (o *ListOrganizationsResponseItemsInner) SetContainerId(v ListOrganizationsResponseItemsInnerGetContainerIdRetType) {
	setListOrganizationsResponseItemsInnerGetContainerIdAttributeType(&o.ContainerId, v)
}

// GetCreationTime returns the CreationTime field value
func (o *ListOrganizationsResponseItemsInner) GetCreationTime() (ret ListOrganizationsResponseItemsInnerGetCreationTimeRetType) {
	ret, _ = o.GetCreationTimeOk()
	return ret
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponseItemsInner) GetCreationTimeOk() (ret ListOrganizationsResponseItemsInnerGetCreationTimeRetType, ok bool) {
	return getListOrganizationsResponseItemsInnerGetCreationTimeAttributeTypeOk(o.CreationTime)
}

// SetCreationTime sets field value
func (o *ListOrganizationsResponseItemsInner) SetCreationTime(v ListOrganizationsResponseItemsInnerGetCreationTimeRetType) {
	setListOrganizationsResponseItemsInnerGetCreationTimeAttributeType(&o.CreationTime, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ListOrganizationsResponseItemsInner) GetLabels() (res ListOrganizationsResponseItemsInnerGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponseItemsInner) GetLabelsOk() (ret ListOrganizationsResponseItemsInnerGetLabelsRetType, ok bool) {
	return getListOrganizationsResponseItemsInnerGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *ListOrganizationsResponseItemsInner) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ListOrganizationsResponseItemsInner) SetLabels(v ListOrganizationsResponseItemsInnerGetLabelsRetType) {
	setListOrganizationsResponseItemsInnerGetLabelsAttributeType(&o.Labels, v)
}

// GetLifecycleState returns the LifecycleState field value
func (o *ListOrganizationsResponseItemsInner) GetLifecycleState() (ret ListOrganizationsResponseItemsInnerGetLifecycleStateRetType) {
	ret, _ = o.GetLifecycleStateOk()
	return ret
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponseItemsInner) GetLifecycleStateOk() (ret ListOrganizationsResponseItemsInnerGetLifecycleStateRetType, ok bool) {
	return getListOrganizationsResponseItemsInnerGetLifecycleStateAttributeTypeOk(o.LifecycleState)
}

// SetLifecycleState sets field value
func (o *ListOrganizationsResponseItemsInner) SetLifecycleState(v ListOrganizationsResponseItemsInnerGetLifecycleStateRetType) {
	setListOrganizationsResponseItemsInnerGetLifecycleStateAttributeType(&o.LifecycleState, v)
}

// GetName returns the Name field value
func (o *ListOrganizationsResponseItemsInner) GetName() (ret ListOrganizationsResponseItemsInnerGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponseItemsInner) GetNameOk() (ret ListOrganizationsResponseItemsInnerGetNameRetType, ok bool) {
	return getListOrganizationsResponseItemsInnerGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *ListOrganizationsResponseItemsInner) SetName(v ListOrganizationsResponseItemsInnerGetNameRetType) {
	setListOrganizationsResponseItemsInnerGetNameAttributeType(&o.Name, v)
}

// GetOrganizationId returns the OrganizationId field value
func (o *ListOrganizationsResponseItemsInner) GetOrganizationId() (ret ListOrganizationsResponseItemsInnerGetOrganizationIdRetType) {
	ret, _ = o.GetOrganizationIdOk()
	return ret
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponseItemsInner) GetOrganizationIdOk() (ret ListOrganizationsResponseItemsInnerGetOrganizationIdRetType, ok bool) {
	return getListOrganizationsResponseItemsInnerGetOrganizationIdAttributeTypeOk(o.OrganizationId)
}

// SetOrganizationId sets field value
func (o *ListOrganizationsResponseItemsInner) SetOrganizationId(v ListOrganizationsResponseItemsInnerGetOrganizationIdRetType) {
	setListOrganizationsResponseItemsInnerGetOrganizationIdAttributeType(&o.OrganizationId, v)
}

// GetUpdateTime returns the UpdateTime field value
func (o *ListOrganizationsResponseItemsInner) GetUpdateTime() (ret ListOrganizationsResponseItemsInnerGetUpdateTimeRetType) {
	ret, _ = o.GetUpdateTimeOk()
	return ret
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponseItemsInner) GetUpdateTimeOk() (ret ListOrganizationsResponseItemsInnerGetUpdateTimeRetType, ok bool) {
	return getListOrganizationsResponseItemsInnerGetUpdateTimeAttributeTypeOk(o.UpdateTime)
}

// SetUpdateTime sets field value
func (o *ListOrganizationsResponseItemsInner) SetUpdateTime(v ListOrganizationsResponseItemsInnerGetUpdateTimeRetType) {
	setListOrganizationsResponseItemsInnerGetUpdateTimeAttributeType(&o.UpdateTime, v)
}

func (o ListOrganizationsResponseItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListOrganizationsResponseItemsInnerGetContainerIdAttributeTypeOk(o.ContainerId); ok {
		toSerialize["ContainerId"] = val
	}
	if val, ok := getListOrganizationsResponseItemsInnerGetCreationTimeAttributeTypeOk(o.CreationTime); ok {
		toSerialize["CreationTime"] = val
	}
	if val, ok := getListOrganizationsResponseItemsInnerGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getListOrganizationsResponseItemsInnerGetLifecycleStateAttributeTypeOk(o.LifecycleState); ok {
		toSerialize["LifecycleState"] = val
	}
	if val, ok := getListOrganizationsResponseItemsInnerGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getListOrganizationsResponseItemsInnerGetOrganizationIdAttributeTypeOk(o.OrganizationId); ok {
		toSerialize["OrganizationId"] = val
	}
	if val, ok := getListOrganizationsResponseItemsInnerGetUpdateTimeAttributeTypeOk(o.UpdateTime); ok {
		toSerialize["UpdateTime"] = val
	}
	return toSerialize, nil
}

type NullableListOrganizationsResponseItemsInner struct {
	value *ListOrganizationsResponseItemsInner
	isSet bool
}

func (v NullableListOrganizationsResponseItemsInner) Get() *ListOrganizationsResponseItemsInner {
	return v.value
}

func (v *NullableListOrganizationsResponseItemsInner) Set(val *ListOrganizationsResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationsResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationsResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationsResponseItemsInner(val *ListOrganizationsResponseItemsInner) *NullableListOrganizationsResponseItemsInner {
	return &NullableListOrganizationsResponseItemsInner{value: val, isSet: true}
}

func (v NullableListOrganizationsResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationsResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
