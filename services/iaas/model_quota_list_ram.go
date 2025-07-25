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

// checks if the QuotaListRam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListRam{}

/*
	types and functions for limit
*/

// isLong
type QuotaListRamGetLimitAttributeType = *int64
type QuotaListRamGetLimitArgType = int64
type QuotaListRamGetLimitRetType = int64

func getQuotaListRamGetLimitAttributeTypeOk(arg QuotaListRamGetLimitAttributeType) (ret QuotaListRamGetLimitRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListRamGetLimitAttributeType(arg *QuotaListRamGetLimitAttributeType, val QuotaListRamGetLimitRetType) {
	*arg = &val
}

/*
	types and functions for usage
*/

// isLong
type QuotaListRamGetUsageAttributeType = *int64
type QuotaListRamGetUsageArgType = int64
type QuotaListRamGetUsageRetType = int64

func getQuotaListRamGetUsageAttributeTypeOk(arg QuotaListRamGetUsageAttributeType) (ret QuotaListRamGetUsageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListRamGetUsageAttributeType(arg *QuotaListRamGetUsageAttributeType, val QuotaListRamGetUsageRetType) {
	*arg = &val
}

// QuotaListRam Amount of server RAM in MiB.
type QuotaListRam struct {
	// REQUIRED
	Limit QuotaListRamGetLimitAttributeType `json:"limit" required:"true"`
	// REQUIRED
	Usage QuotaListRamGetUsageAttributeType `json:"usage" required:"true"`
}

type _QuotaListRam QuotaListRam

// NewQuotaListRam instantiates a new QuotaListRam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListRam(limit QuotaListRamGetLimitArgType, usage QuotaListRamGetUsageArgType) *QuotaListRam {
	this := QuotaListRam{}
	setQuotaListRamGetLimitAttributeType(&this.Limit, limit)
	setQuotaListRamGetUsageAttributeType(&this.Usage, usage)
	return &this
}

// NewQuotaListRamWithDefaults instantiates a new QuotaListRam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListRamWithDefaults() *QuotaListRam {
	this := QuotaListRam{}
	return &this
}

// GetLimit returns the Limit field value
func (o *QuotaListRam) GetLimit() (ret QuotaListRamGetLimitRetType) {
	ret, _ = o.GetLimitOk()
	return ret
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListRam) GetLimitOk() (ret QuotaListRamGetLimitRetType, ok bool) {
	return getQuotaListRamGetLimitAttributeTypeOk(o.Limit)
}

// SetLimit sets field value
func (o *QuotaListRam) SetLimit(v QuotaListRamGetLimitRetType) {
	setQuotaListRamGetLimitAttributeType(&o.Limit, v)
}

// GetUsage returns the Usage field value
func (o *QuotaListRam) GetUsage() (ret QuotaListRamGetUsageRetType) {
	ret, _ = o.GetUsageOk()
	return ret
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListRam) GetUsageOk() (ret QuotaListRamGetUsageRetType, ok bool) {
	return getQuotaListRamGetUsageAttributeTypeOk(o.Usage)
}

// SetUsage sets field value
func (o *QuotaListRam) SetUsage(v QuotaListRamGetUsageRetType) {
	setQuotaListRamGetUsageAttributeType(&o.Usage, v)
}

func (o QuotaListRam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getQuotaListRamGetLimitAttributeTypeOk(o.Limit); ok {
		toSerialize["Limit"] = val
	}
	if val, ok := getQuotaListRamGetUsageAttributeTypeOk(o.Usage); ok {
		toSerialize["Usage"] = val
	}
	return toSerialize, nil
}

type NullableQuotaListRam struct {
	value *QuotaListRam
	isSet bool
}

func (v NullableQuotaListRam) Get() *QuotaListRam {
	return v.value
}

func (v *NullableQuotaListRam) Set(val *QuotaListRam) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListRam) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListRam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListRam(val *QuotaListRam) *NullableQuotaListRam {
	return &NullableQuotaListRam{value: val, isSet: true}
}

func (v NullableQuotaListRam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListRam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
