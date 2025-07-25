/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the ListMetricsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetricsResponse{}

/*
	types and functions for hosts
*/

// isArray
type ListMetricsResponseGetHostsAttributeType = *[]Host
type ListMetricsResponseGetHostsArgType = []Host
type ListMetricsResponseGetHostsRetType = []Host

func getListMetricsResponseGetHostsAttributeTypeOk(arg ListMetricsResponseGetHostsAttributeType) (ret ListMetricsResponseGetHostsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListMetricsResponseGetHostsAttributeType(arg *ListMetricsResponseGetHostsAttributeType, val ListMetricsResponseGetHostsRetType) {
	*arg = &val
}

// ListMetricsResponse struct for ListMetricsResponse
type ListMetricsResponse struct {
	Hosts ListMetricsResponseGetHostsAttributeType `json:"hosts,omitempty"`
}

// NewListMetricsResponse instantiates a new ListMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetricsResponse() *ListMetricsResponse {
	this := ListMetricsResponse{}
	return &this
}

// NewListMetricsResponseWithDefaults instantiates a new ListMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetricsResponseWithDefaults() *ListMetricsResponse {
	this := ListMetricsResponse{}
	return &this
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *ListMetricsResponse) GetHosts() (res ListMetricsResponseGetHostsRetType) {
	res, _ = o.GetHostsOk()
	return
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricsResponse) GetHostsOk() (ret ListMetricsResponseGetHostsRetType, ok bool) {
	return getListMetricsResponseGetHostsAttributeTypeOk(o.Hosts)
}

// HasHosts returns a boolean if a field has been set.
func (o *ListMetricsResponse) HasHosts() bool {
	_, ok := o.GetHostsOk()
	return ok
}

// SetHosts gets a reference to the given []Host and assigns it to the Hosts field.
func (o *ListMetricsResponse) SetHosts(v ListMetricsResponseGetHostsRetType) {
	setListMetricsResponseGetHostsAttributeType(&o.Hosts, v)
}

func (o ListMetricsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListMetricsResponseGetHostsAttributeTypeOk(o.Hosts); ok {
		toSerialize["Hosts"] = val
	}
	return toSerialize, nil
}

type NullableListMetricsResponse struct {
	value *ListMetricsResponse
	isSet bool
}

func (v NullableListMetricsResponse) Get() *ListMetricsResponse {
	return v.value
}

func (v *NullableListMetricsResponse) Set(val *ListMetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetricsResponse(val *ListMetricsResponse) *NullableListMetricsResponse {
	return &NullableListMetricsResponse{value: val, isSet: true}
}

func (v NullableListMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
