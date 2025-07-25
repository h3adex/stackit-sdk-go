/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
)

// checks if the DNS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNS{}

/*
	types and functions for enabled
*/

// isBoolean
type DNSgetEnabledAttributeType = *bool
type DNSgetEnabledArgType = bool
type DNSgetEnabledRetType = bool

func getDNSgetEnabledAttributeTypeOk(arg DNSgetEnabledAttributeType) (ret DNSgetEnabledRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDNSgetEnabledAttributeType(arg *DNSgetEnabledAttributeType, val DNSgetEnabledRetType) {
	*arg = &val
}

/*
	types and functions for zones
*/

// isArray
type DNSGetZonesAttributeType = *[]string
type DNSGetZonesArgType = []string
type DNSGetZonesRetType = []string

func getDNSGetZonesAttributeTypeOk(arg DNSGetZonesAttributeType) (ret DNSGetZonesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDNSGetZonesAttributeType(arg *DNSGetZonesAttributeType, val DNSGetZonesRetType) {
	*arg = &val
}

// DNS struct for DNS
type DNS struct {
	// Enables the dns extension.
	// REQUIRED
	Enabled DNSgetEnabledAttributeType `json:"enabled" required:"true"`
	// Array of domain filters for externalDNS, e.g., *.runs.onstackit.cloud.
	Zones DNSGetZonesAttributeType `json:"zones,omitempty"`
}

type _DNS DNS

// NewDNS instantiates a new DNS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNS(enabled DNSgetEnabledArgType) *DNS {
	this := DNS{}
	setDNSgetEnabledAttributeType(&this.Enabled, enabled)
	return &this
}

// NewDNSWithDefaults instantiates a new DNS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSWithDefaults() *DNS {
	this := DNS{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DNS) GetEnabled() (ret DNSgetEnabledRetType) {
	ret, _ = o.GetEnabledOk()
	return ret
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DNS) GetEnabledOk() (ret DNSgetEnabledRetType, ok bool) {
	return getDNSgetEnabledAttributeTypeOk(o.Enabled)
}

// SetEnabled sets field value
func (o *DNS) SetEnabled(v DNSgetEnabledRetType) {
	setDNSgetEnabledAttributeType(&o.Enabled, v)
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *DNS) GetZones() (res DNSGetZonesRetType) {
	res, _ = o.GetZonesOk()
	return
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNS) GetZonesOk() (ret DNSGetZonesRetType, ok bool) {
	return getDNSGetZonesAttributeTypeOk(o.Zones)
}

// HasZones returns a boolean if a field has been set.
func (o *DNS) HasZones() bool {
	_, ok := o.GetZonesOk()
	return ok
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *DNS) SetZones(v DNSGetZonesRetType) {
	setDNSGetZonesAttributeType(&o.Zones, v)
}

func (o DNS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getDNSgetEnabledAttributeTypeOk(o.Enabled); ok {
		toSerialize["Enabled"] = val
	}
	if val, ok := getDNSGetZonesAttributeTypeOk(o.Zones); ok {
		toSerialize["Zones"] = val
	}
	return toSerialize, nil
}

type NullableDNS struct {
	value *DNS
	isSet bool
}

func (v NullableDNS) Get() *DNS {
	return v.value
}

func (v *NullableDNS) Set(val *DNS) {
	v.value = val
	v.isSet = true
}

func (v NullableDNS) IsSet() bool {
	return v.isSet
}

func (v *NullableDNS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNS(val *DNS) *NullableDNS {
	return &NullableDNS{value: val, isSet: true}
}

func (v NullableDNS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
