/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the CreateAlertConfigRoutePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlertConfigRoutePayload{}

/*
	types and functions for continue
*/

// isBoolean
type CreateAlertConfigRoutePayloadgetContinueAttributeType = *bool
type CreateAlertConfigRoutePayloadgetContinueArgType = bool
type CreateAlertConfigRoutePayloadgetContinueRetType = bool

func getCreateAlertConfigRoutePayloadgetContinueAttributeTypeOk(arg CreateAlertConfigRoutePayloadgetContinueAttributeType) (ret CreateAlertConfigRoutePayloadgetContinueRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadgetContinueAttributeType(arg *CreateAlertConfigRoutePayloadgetContinueAttributeType, val CreateAlertConfigRoutePayloadgetContinueRetType) {
	*arg = &val
}

/*
	types and functions for groupBy
*/

// isArray
type CreateAlertConfigRoutePayloadGetGroupByAttributeType = *[]string
type CreateAlertConfigRoutePayloadGetGroupByArgType = []string
type CreateAlertConfigRoutePayloadGetGroupByRetType = []string

func getCreateAlertConfigRoutePayloadGetGroupByAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetGroupByAttributeType) (ret CreateAlertConfigRoutePayloadGetGroupByRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetGroupByAttributeType(arg *CreateAlertConfigRoutePayloadGetGroupByAttributeType, val CreateAlertConfigRoutePayloadGetGroupByRetType) {
	*arg = &val
}

/*
	types and functions for groupInterval
*/

// isNotNullableString
type CreateAlertConfigRoutePayloadGetGroupIntervalAttributeType = *string

func getCreateAlertConfigRoutePayloadGetGroupIntervalAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetGroupIntervalAttributeType) (ret CreateAlertConfigRoutePayloadGetGroupIntervalRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetGroupIntervalAttributeType(arg *CreateAlertConfigRoutePayloadGetGroupIntervalAttributeType, val CreateAlertConfigRoutePayloadGetGroupIntervalRetType) {
	*arg = &val
}

type CreateAlertConfigRoutePayloadGetGroupIntervalArgType = string
type CreateAlertConfigRoutePayloadGetGroupIntervalRetType = string

/*
	types and functions for groupWait
*/

// isNotNullableString
type CreateAlertConfigRoutePayloadGetGroupWaitAttributeType = *string

func getCreateAlertConfigRoutePayloadGetGroupWaitAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetGroupWaitAttributeType) (ret CreateAlertConfigRoutePayloadGetGroupWaitRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetGroupWaitAttributeType(arg *CreateAlertConfigRoutePayloadGetGroupWaitAttributeType, val CreateAlertConfigRoutePayloadGetGroupWaitRetType) {
	*arg = &val
}

type CreateAlertConfigRoutePayloadGetGroupWaitArgType = string
type CreateAlertConfigRoutePayloadGetGroupWaitRetType = string

/*
	types and functions for match
*/

// isFreeform
type CreateAlertConfigRoutePayloadGetMatchAttributeType = *map[string]interface{}
type CreateAlertConfigRoutePayloadGetMatchArgType = map[string]interface{}
type CreateAlertConfigRoutePayloadGetMatchRetType = map[string]interface{}

func getCreateAlertConfigRoutePayloadGetMatchAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetMatchAttributeType) (ret CreateAlertConfigRoutePayloadGetMatchRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetMatchAttributeType(arg *CreateAlertConfigRoutePayloadGetMatchAttributeType, val CreateAlertConfigRoutePayloadGetMatchRetType) {
	*arg = &val
}

/*
	types and functions for matchRe
*/

// isFreeform
type CreateAlertConfigRoutePayloadGetMatchReAttributeType = *map[string]interface{}
type CreateAlertConfigRoutePayloadGetMatchReArgType = map[string]interface{}
type CreateAlertConfigRoutePayloadGetMatchReRetType = map[string]interface{}

func getCreateAlertConfigRoutePayloadGetMatchReAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetMatchReAttributeType) (ret CreateAlertConfigRoutePayloadGetMatchReRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetMatchReAttributeType(arg *CreateAlertConfigRoutePayloadGetMatchReAttributeType, val CreateAlertConfigRoutePayloadGetMatchReRetType) {
	*arg = &val
}

/*
	types and functions for matchers
*/

// isArray
type CreateAlertConfigRoutePayloadGetMatchersAttributeType = *[]string
type CreateAlertConfigRoutePayloadGetMatchersArgType = []string
type CreateAlertConfigRoutePayloadGetMatchersRetType = []string

func getCreateAlertConfigRoutePayloadGetMatchersAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetMatchersAttributeType) (ret CreateAlertConfigRoutePayloadGetMatchersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetMatchersAttributeType(arg *CreateAlertConfigRoutePayloadGetMatchersAttributeType, val CreateAlertConfigRoutePayloadGetMatchersRetType) {
	*arg = &val
}

/*
	types and functions for receiver
*/

// isNotNullableString
type CreateAlertConfigRoutePayloadGetReceiverAttributeType = *string

func getCreateAlertConfigRoutePayloadGetReceiverAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetReceiverAttributeType) (ret CreateAlertConfigRoutePayloadGetReceiverRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetReceiverAttributeType(arg *CreateAlertConfigRoutePayloadGetReceiverAttributeType, val CreateAlertConfigRoutePayloadGetReceiverRetType) {
	*arg = &val
}

type CreateAlertConfigRoutePayloadGetReceiverArgType = string
type CreateAlertConfigRoutePayloadGetReceiverRetType = string

/*
	types and functions for repeatInterval
*/

// isNotNullableString
type CreateAlertConfigRoutePayloadGetRepeatIntervalAttributeType = *string

func getCreateAlertConfigRoutePayloadGetRepeatIntervalAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetRepeatIntervalAttributeType) (ret CreateAlertConfigRoutePayloadGetRepeatIntervalRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetRepeatIntervalAttributeType(arg *CreateAlertConfigRoutePayloadGetRepeatIntervalAttributeType, val CreateAlertConfigRoutePayloadGetRepeatIntervalRetType) {
	*arg = &val
}

type CreateAlertConfigRoutePayloadGetRepeatIntervalArgType = string
type CreateAlertConfigRoutePayloadGetRepeatIntervalRetType = string

/*
	types and functions for routes
*/

// isArray
type CreateAlertConfigRoutePayloadGetRoutesAttributeType = *[]CreateAlertConfigRoutePayloadRoutesInner
type CreateAlertConfigRoutePayloadGetRoutesArgType = []CreateAlertConfigRoutePayloadRoutesInner
type CreateAlertConfigRoutePayloadGetRoutesRetType = []CreateAlertConfigRoutePayloadRoutesInner

func getCreateAlertConfigRoutePayloadGetRoutesAttributeTypeOk(arg CreateAlertConfigRoutePayloadGetRoutesAttributeType) (ret CreateAlertConfigRoutePayloadGetRoutesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigRoutePayloadGetRoutesAttributeType(arg *CreateAlertConfigRoutePayloadGetRoutesAttributeType, val CreateAlertConfigRoutePayloadGetRoutesRetType) {
	*arg = &val
}

// CreateAlertConfigRoutePayload The root node of the routing tree.
type CreateAlertConfigRoutePayload struct {
	// Whether an alert should continue matching subsequent sibling nodes.
	Continue CreateAlertConfigRoutePayloadgetContinueAttributeType `json:"continue,omitempty"`
	// The labels by which incoming alerts are grouped together. For example, multiple alerts coming in for cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels use the special value '...' as the sole label name, for example: group_by: ['...']. This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
	GroupBy CreateAlertConfigRoutePayloadGetGroupByAttributeType `json:"groupBy,omitempty"`
	// How long to wait before sending a notification about new alerts that are added to a group of alerts for which an initial notification has already been sent. (Usually ~5m or more.) `Additional Validators:` * must be a valid time format
	GroupInterval CreateAlertConfigRoutePayloadGetGroupIntervalAttributeType `json:"groupInterval,omitempty"`
	// How long to initially wait to send a notification for a group of alerts. Allows to wait for an inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.) `Additional Validators:` * must be a valid time format
	GroupWait CreateAlertConfigRoutePayloadGetGroupWaitAttributeType `json:"groupWait,omitempty"`
	//  Deprecated: map of key:value. A set of equality matchers an alert has to fulfill to match the node.  `Additional Validators:` * should not contain more than 5 keys * each key and value should not be longer than 200 characters * key and values should only include the characters: a-zA-Z0-9_./@&?:-
	Match CreateAlertConfigRoutePayloadGetMatchAttributeType `json:"match,omitempty"`
	//  Deprecated: map of key:value. A set of regex-matchers an alert has to fulfill to match the node.  `Additional Validators:` * should not contain more than 5 keys * each key and value should not be longer than 200 characters
	MatchRe CreateAlertConfigRoutePayloadGetMatchReAttributeType `json:"matchRe,omitempty"`
	// A list of matchers that an alert has to fulfill to match the node. A matcher is a string with a syntax inspired by PromQL and OpenMetrics. The syntax of a matcher consists of three tokens: * A valid Prometheus label name. * One of =, !=, =~, or !~. = means equals, != means that the strings are not equal, =~ is used for equality of regex expressions and !~ is used for un-equality of regex expressions. They have the same meaning as known from PromQL selectors. * A UTF-8 string, which may be enclosed in double quotes. Before or after each token, there may be any amount of whitespace. `Additional Validators:` * should not contain more than 5 keys * each key and value should not be longer than 200 characters
	Matchers CreateAlertConfigRoutePayloadGetMatchersAttributeType `json:"matchers,omitempty"`
	// Receiver that should be one item of receivers `Additional Validators:` * must be a in name of receivers
	// REQUIRED
	Receiver CreateAlertConfigRoutePayloadGetReceiverAttributeType `json:"receiver" required:"true"`
	// How long to wait before sending a notification again if it has already been sent successfully for an alert. (Usually ~3h or more). `Additional Validators:` * must be a valid time format
	RepeatInterval CreateAlertConfigRoutePayloadGetRepeatIntervalAttributeType `json:"repeatInterval,omitempty"`
	// Zero or more child routes.
	Routes CreateAlertConfigRoutePayloadGetRoutesAttributeType `json:"routes,omitempty"`
}

type _CreateAlertConfigRoutePayload CreateAlertConfigRoutePayload

// NewCreateAlertConfigRoutePayload instantiates a new CreateAlertConfigRoutePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlertConfigRoutePayload(receiver CreateAlertConfigRoutePayloadGetReceiverArgType) *CreateAlertConfigRoutePayload {
	this := CreateAlertConfigRoutePayload{}
	setCreateAlertConfigRoutePayloadGetReceiverAttributeType(&this.Receiver, receiver)
	return &this
}

// NewCreateAlertConfigRoutePayloadWithDefaults instantiates a new CreateAlertConfigRoutePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlertConfigRoutePayloadWithDefaults() *CreateAlertConfigRoutePayload {
	this := CreateAlertConfigRoutePayload{}
	var continue_ bool = false
	this.Continue = &continue_
	var groupInterval string = "5m"
	this.GroupInterval = &groupInterval
	var groupWait string = "30s"
	this.GroupWait = &groupWait
	var repeatInterval string = "4h"
	this.RepeatInterval = &repeatInterval
	return &this
}

// GetContinue returns the Continue field value if set, zero value otherwise.
func (o *CreateAlertConfigRoutePayload) GetContinue() (res CreateAlertConfigRoutePayloadgetContinueRetType) {
	res, _ = o.GetContinueOk()
	return
}

// GetContinueOk returns a tuple with the Continue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetContinueOk() (ret CreateAlertConfigRoutePayloadgetContinueRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadgetContinueAttributeTypeOk(o.Continue)
}

// HasContinue returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasContinue() bool {
	_, ok := o.GetContinueOk()
	return ok
}

// SetContinue gets a reference to the given bool and assigns it to the Continue field.
func (o *CreateAlertConfigRoutePayload) SetContinue(v CreateAlertConfigRoutePayloadgetContinueRetType) {
	setCreateAlertConfigRoutePayloadgetContinueAttributeType(&o.Continue, v)
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *CreateAlertConfigRoutePayload) GetGroupBy() (res CreateAlertConfigRoutePayloadGetGroupByRetType) {
	res, _ = o.GetGroupByOk()
	return
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetGroupByOk() (ret CreateAlertConfigRoutePayloadGetGroupByRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetGroupByAttributeTypeOk(o.GroupBy)
}

// HasGroupBy returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasGroupBy() bool {
	_, ok := o.GetGroupByOk()
	return ok
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *CreateAlertConfigRoutePayload) SetGroupBy(v CreateAlertConfigRoutePayloadGetGroupByRetType) {
	setCreateAlertConfigRoutePayloadGetGroupByAttributeType(&o.GroupBy, v)
}

// GetGroupInterval returns the GroupInterval field value if set, zero value otherwise.
func (o *CreateAlertConfigRoutePayload) GetGroupInterval() (res CreateAlertConfigRoutePayloadGetGroupIntervalRetType) {
	res, _ = o.GetGroupIntervalOk()
	return
}

// GetGroupIntervalOk returns a tuple with the GroupInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetGroupIntervalOk() (ret CreateAlertConfigRoutePayloadGetGroupIntervalRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetGroupIntervalAttributeTypeOk(o.GroupInterval)
}

// HasGroupInterval returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasGroupInterval() bool {
	_, ok := o.GetGroupIntervalOk()
	return ok
}

// SetGroupInterval gets a reference to the given string and assigns it to the GroupInterval field.
func (o *CreateAlertConfigRoutePayload) SetGroupInterval(v CreateAlertConfigRoutePayloadGetGroupIntervalRetType) {
	setCreateAlertConfigRoutePayloadGetGroupIntervalAttributeType(&o.GroupInterval, v)
}

// GetGroupWait returns the GroupWait field value if set, zero value otherwise.
func (o *CreateAlertConfigRoutePayload) GetGroupWait() (res CreateAlertConfigRoutePayloadGetGroupWaitRetType) {
	res, _ = o.GetGroupWaitOk()
	return
}

// GetGroupWaitOk returns a tuple with the GroupWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetGroupWaitOk() (ret CreateAlertConfigRoutePayloadGetGroupWaitRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetGroupWaitAttributeTypeOk(o.GroupWait)
}

// HasGroupWait returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasGroupWait() bool {
	_, ok := o.GetGroupWaitOk()
	return ok
}

// SetGroupWait gets a reference to the given string and assigns it to the GroupWait field.
func (o *CreateAlertConfigRoutePayload) SetGroupWait(v CreateAlertConfigRoutePayloadGetGroupWaitRetType) {
	setCreateAlertConfigRoutePayloadGetGroupWaitAttributeType(&o.GroupWait, v)
}

// GetMatch returns the Match field value if set, zero value otherwise.
// Deprecated
func (o *CreateAlertConfigRoutePayload) GetMatch() (res CreateAlertConfigRoutePayloadGetMatchRetType) {
	res, _ = o.GetMatchOk()
	return
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateAlertConfigRoutePayload) GetMatchOk() (ret CreateAlertConfigRoutePayloadGetMatchRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetMatchAttributeTypeOk(o.Match)
}

// HasMatch returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasMatch() bool {
	_, ok := o.GetMatchOk()
	return ok
}

// SetMatch gets a reference to the given map[string]interface{} and assigns it to the Match field.
// Deprecated
func (o *CreateAlertConfigRoutePayload) SetMatch(v CreateAlertConfigRoutePayloadGetMatchRetType) {
	setCreateAlertConfigRoutePayloadGetMatchAttributeType(&o.Match, v)
}

// GetMatchRe returns the MatchRe field value if set, zero value otherwise.
// Deprecated
func (o *CreateAlertConfigRoutePayload) GetMatchRe() (res CreateAlertConfigRoutePayloadGetMatchReRetType) {
	res, _ = o.GetMatchReOk()
	return
}

// GetMatchReOk returns a tuple with the MatchRe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateAlertConfigRoutePayload) GetMatchReOk() (ret CreateAlertConfigRoutePayloadGetMatchReRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetMatchReAttributeTypeOk(o.MatchRe)
}

// HasMatchRe returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasMatchRe() bool {
	_, ok := o.GetMatchReOk()
	return ok
}

// SetMatchRe gets a reference to the given map[string]interface{} and assigns it to the MatchRe field.
// Deprecated
func (o *CreateAlertConfigRoutePayload) SetMatchRe(v CreateAlertConfigRoutePayloadGetMatchReRetType) {
	setCreateAlertConfigRoutePayloadGetMatchReAttributeType(&o.MatchRe, v)
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *CreateAlertConfigRoutePayload) GetMatchers() (res CreateAlertConfigRoutePayloadGetMatchersRetType) {
	res, _ = o.GetMatchersOk()
	return
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetMatchersOk() (ret CreateAlertConfigRoutePayloadGetMatchersRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetMatchersAttributeTypeOk(o.Matchers)
}

// HasMatchers returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasMatchers() bool {
	_, ok := o.GetMatchersOk()
	return ok
}

// SetMatchers gets a reference to the given []string and assigns it to the Matchers field.
func (o *CreateAlertConfigRoutePayload) SetMatchers(v CreateAlertConfigRoutePayloadGetMatchersRetType) {
	setCreateAlertConfigRoutePayloadGetMatchersAttributeType(&o.Matchers, v)
}

// GetReceiver returns the Receiver field value
func (o *CreateAlertConfigRoutePayload) GetReceiver() (ret CreateAlertConfigRoutePayloadGetReceiverRetType) {
	ret, _ = o.GetReceiverOk()
	return ret
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetReceiverOk() (ret CreateAlertConfigRoutePayloadGetReceiverRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetReceiverAttributeTypeOk(o.Receiver)
}

// SetReceiver sets field value
func (o *CreateAlertConfigRoutePayload) SetReceiver(v CreateAlertConfigRoutePayloadGetReceiverRetType) {
	setCreateAlertConfigRoutePayloadGetReceiverAttributeType(&o.Receiver, v)
}

// GetRepeatInterval returns the RepeatInterval field value if set, zero value otherwise.
func (o *CreateAlertConfigRoutePayload) GetRepeatInterval() (res CreateAlertConfigRoutePayloadGetRepeatIntervalRetType) {
	res, _ = o.GetRepeatIntervalOk()
	return
}

// GetRepeatIntervalOk returns a tuple with the RepeatInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetRepeatIntervalOk() (ret CreateAlertConfigRoutePayloadGetRepeatIntervalRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetRepeatIntervalAttributeTypeOk(o.RepeatInterval)
}

// HasRepeatInterval returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasRepeatInterval() bool {
	_, ok := o.GetRepeatIntervalOk()
	return ok
}

// SetRepeatInterval gets a reference to the given string and assigns it to the RepeatInterval field.
func (o *CreateAlertConfigRoutePayload) SetRepeatInterval(v CreateAlertConfigRoutePayloadGetRepeatIntervalRetType) {
	setCreateAlertConfigRoutePayloadGetRepeatIntervalAttributeType(&o.RepeatInterval, v)
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *CreateAlertConfigRoutePayload) GetRoutes() (res CreateAlertConfigRoutePayloadGetRoutesRetType) {
	res, _ = o.GetRoutesOk()
	return
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigRoutePayload) GetRoutesOk() (ret CreateAlertConfigRoutePayloadGetRoutesRetType, ok bool) {
	return getCreateAlertConfigRoutePayloadGetRoutesAttributeTypeOk(o.Routes)
}

// HasRoutes returns a boolean if a field has been set.
func (o *CreateAlertConfigRoutePayload) HasRoutes() bool {
	_, ok := o.GetRoutesOk()
	return ok
}

// SetRoutes gets a reference to the given []CreateAlertConfigRoutePayloadRoutesInner and assigns it to the Routes field.
func (o *CreateAlertConfigRoutePayload) SetRoutes(v CreateAlertConfigRoutePayloadGetRoutesRetType) {
	setCreateAlertConfigRoutePayloadGetRoutesAttributeType(&o.Routes, v)
}

func (o CreateAlertConfigRoutePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateAlertConfigRoutePayloadgetContinueAttributeTypeOk(o.Continue); ok {
		toSerialize["Continue"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetGroupByAttributeTypeOk(o.GroupBy); ok {
		toSerialize["GroupBy"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetGroupIntervalAttributeTypeOk(o.GroupInterval); ok {
		toSerialize["GroupInterval"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetGroupWaitAttributeTypeOk(o.GroupWait); ok {
		toSerialize["GroupWait"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetMatchAttributeTypeOk(o.Match); ok {
		toSerialize["Match"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetMatchReAttributeTypeOk(o.MatchRe); ok {
		toSerialize["MatchRe"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetMatchersAttributeTypeOk(o.Matchers); ok {
		toSerialize["Matchers"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetReceiverAttributeTypeOk(o.Receiver); ok {
		toSerialize["Receiver"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetRepeatIntervalAttributeTypeOk(o.RepeatInterval); ok {
		toSerialize["RepeatInterval"] = val
	}
	if val, ok := getCreateAlertConfigRoutePayloadGetRoutesAttributeTypeOk(o.Routes); ok {
		toSerialize["Routes"] = val
	}
	return toSerialize, nil
}

type NullableCreateAlertConfigRoutePayload struct {
	value *CreateAlertConfigRoutePayload
	isSet bool
}

func (v NullableCreateAlertConfigRoutePayload) Get() *CreateAlertConfigRoutePayload {
	return v.value
}

func (v *NullableCreateAlertConfigRoutePayload) Set(val *CreateAlertConfigRoutePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlertConfigRoutePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlertConfigRoutePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlertConfigRoutePayload(val *CreateAlertConfigRoutePayload) *NullableCreateAlertConfigRoutePayload {
	return &NullableCreateAlertConfigRoutePayload{value: val, isSet: true}
}

func (v NullableCreateAlertConfigRoutePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlertConfigRoutePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
