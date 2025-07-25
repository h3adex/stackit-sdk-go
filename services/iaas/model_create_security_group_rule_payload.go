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

// checks if the CreateSecurityGroupRulePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSecurityGroupRulePayload{}

/*
	types and functions for description
*/

// isNotNullableString
type CreateSecurityGroupRulePayloadGetDescriptionAttributeType = *string

func getCreateSecurityGroupRulePayloadGetDescriptionAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetDescriptionAttributeType) (ret CreateSecurityGroupRulePayloadGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetDescriptionAttributeType(arg *CreateSecurityGroupRulePayloadGetDescriptionAttributeType, val CreateSecurityGroupRulePayloadGetDescriptionRetType) {
	*arg = &val
}

type CreateSecurityGroupRulePayloadGetDescriptionArgType = string
type CreateSecurityGroupRulePayloadGetDescriptionRetType = string

/*
	types and functions for direction
*/

// isNotNullableString
type CreateSecurityGroupRulePayloadGetDirectionAttributeType = *string

func getCreateSecurityGroupRulePayloadGetDirectionAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetDirectionAttributeType) (ret CreateSecurityGroupRulePayloadGetDirectionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetDirectionAttributeType(arg *CreateSecurityGroupRulePayloadGetDirectionAttributeType, val CreateSecurityGroupRulePayloadGetDirectionRetType) {
	*arg = &val
}

type CreateSecurityGroupRulePayloadGetDirectionArgType = string
type CreateSecurityGroupRulePayloadGetDirectionRetType = string

/*
	types and functions for ethertype
*/

// isNotNullableString
type CreateSecurityGroupRulePayloadGetEthertypeAttributeType = *string

func getCreateSecurityGroupRulePayloadGetEthertypeAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetEthertypeAttributeType) (ret CreateSecurityGroupRulePayloadGetEthertypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetEthertypeAttributeType(arg *CreateSecurityGroupRulePayloadGetEthertypeAttributeType, val CreateSecurityGroupRulePayloadGetEthertypeRetType) {
	*arg = &val
}

type CreateSecurityGroupRulePayloadGetEthertypeArgType = string
type CreateSecurityGroupRulePayloadGetEthertypeRetType = string

/*
	types and functions for icmpParameters
*/

// isModel
type CreateSecurityGroupRulePayloadGetIcmpParametersAttributeType = *ICMPParameters
type CreateSecurityGroupRulePayloadGetIcmpParametersArgType = ICMPParameters
type CreateSecurityGroupRulePayloadGetIcmpParametersRetType = ICMPParameters

func getCreateSecurityGroupRulePayloadGetIcmpParametersAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetIcmpParametersAttributeType) (ret CreateSecurityGroupRulePayloadGetIcmpParametersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetIcmpParametersAttributeType(arg *CreateSecurityGroupRulePayloadGetIcmpParametersAttributeType, val CreateSecurityGroupRulePayloadGetIcmpParametersRetType) {
	*arg = &val
}

/*
	types and functions for id
*/

// isNotNullableString
type CreateSecurityGroupRulePayloadGetIdAttributeType = *string

func getCreateSecurityGroupRulePayloadGetIdAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetIdAttributeType) (ret CreateSecurityGroupRulePayloadGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetIdAttributeType(arg *CreateSecurityGroupRulePayloadGetIdAttributeType, val CreateSecurityGroupRulePayloadGetIdRetType) {
	*arg = &val
}

type CreateSecurityGroupRulePayloadGetIdArgType = string
type CreateSecurityGroupRulePayloadGetIdRetType = string

/*
	types and functions for ipRange
*/

// isNotNullableString
type CreateSecurityGroupRulePayloadGetIpRangeAttributeType = *string

func getCreateSecurityGroupRulePayloadGetIpRangeAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetIpRangeAttributeType) (ret CreateSecurityGroupRulePayloadGetIpRangeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetIpRangeAttributeType(arg *CreateSecurityGroupRulePayloadGetIpRangeAttributeType, val CreateSecurityGroupRulePayloadGetIpRangeRetType) {
	*arg = &val
}

type CreateSecurityGroupRulePayloadGetIpRangeArgType = string
type CreateSecurityGroupRulePayloadGetIpRangeRetType = string

/*
	types and functions for portRange
*/

// isModel
type CreateSecurityGroupRulePayloadGetPortRangeAttributeType = *PortRange
type CreateSecurityGroupRulePayloadGetPortRangeArgType = PortRange
type CreateSecurityGroupRulePayloadGetPortRangeRetType = PortRange

func getCreateSecurityGroupRulePayloadGetPortRangeAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetPortRangeAttributeType) (ret CreateSecurityGroupRulePayloadGetPortRangeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetPortRangeAttributeType(arg *CreateSecurityGroupRulePayloadGetPortRangeAttributeType, val CreateSecurityGroupRulePayloadGetPortRangeRetType) {
	*arg = &val
}

/*
	types and functions for remoteSecurityGroupId
*/

// isNotNullableString
type CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeType = *string

func getCreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeType) (ret CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeType(arg *CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeType, val CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdRetType) {
	*arg = &val
}

type CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdArgType = string
type CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdRetType = string

/*
	types and functions for securityGroupId
*/

// isNotNullableString
type CreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeType = *string

func getCreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeType) (ret CreateSecurityGroupRulePayloadGetSecurityGroupIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeType(arg *CreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeType, val CreateSecurityGroupRulePayloadGetSecurityGroupIdRetType) {
	*arg = &val
}

type CreateSecurityGroupRulePayloadGetSecurityGroupIdArgType = string
type CreateSecurityGroupRulePayloadGetSecurityGroupIdRetType = string

/*
	types and functions for protocol
*/

// isModel
type CreateSecurityGroupRulePayloadGetProtocolAttributeType = *CreateProtocol
type CreateSecurityGroupRulePayloadGetProtocolArgType = CreateProtocol
type CreateSecurityGroupRulePayloadGetProtocolRetType = CreateProtocol

func getCreateSecurityGroupRulePayloadGetProtocolAttributeTypeOk(arg CreateSecurityGroupRulePayloadGetProtocolAttributeType) (ret CreateSecurityGroupRulePayloadGetProtocolRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateSecurityGroupRulePayloadGetProtocolAttributeType(arg *CreateSecurityGroupRulePayloadGetProtocolAttributeType, val CreateSecurityGroupRulePayloadGetProtocolRetType) {
	*arg = &val
}

// CreateSecurityGroupRulePayload Object that represents a request body for security group rule creation.
type CreateSecurityGroupRulePayload struct {
	// Description Object. Allows string up to 255 Characters.
	Description CreateSecurityGroupRulePayloadGetDescriptionAttributeType `json:"description,omitempty"`
	// The direction of the traffic which the rule should match.
	// REQUIRED
	Direction CreateSecurityGroupRulePayloadGetDirectionAttributeType `json:"direction" required:"true"`
	// The ethertype which the rule should match.
	Ethertype      CreateSecurityGroupRulePayloadGetEthertypeAttributeType      `json:"ethertype,omitempty"`
	IcmpParameters CreateSecurityGroupRulePayloadGetIcmpParametersAttributeType `json:"icmpParameters,omitempty"`
	// Universally Unique Identifier (UUID).
	Id CreateSecurityGroupRulePayloadGetIdAttributeType `json:"id,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	IpRange   CreateSecurityGroupRulePayloadGetIpRangeAttributeType   `json:"ipRange,omitempty"`
	PortRange CreateSecurityGroupRulePayloadGetPortRangeAttributeType `json:"portRange,omitempty"`
	// Universally Unique Identifier (UUID).
	RemoteSecurityGroupId CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeType `json:"remoteSecurityGroupId,omitempty"`
	// Universally Unique Identifier (UUID).
	SecurityGroupId CreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeType `json:"securityGroupId,omitempty"`
	Protocol        CreateSecurityGroupRulePayloadGetProtocolAttributeType        `json:"protocol,omitempty"`
}

type _CreateSecurityGroupRulePayload CreateSecurityGroupRulePayload

// NewCreateSecurityGroupRulePayload instantiates a new CreateSecurityGroupRulePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecurityGroupRulePayload(direction CreateSecurityGroupRulePayloadGetDirectionArgType) *CreateSecurityGroupRulePayload {
	this := CreateSecurityGroupRulePayload{}
	setCreateSecurityGroupRulePayloadGetDirectionAttributeType(&this.Direction, direction)
	return &this
}

// NewCreateSecurityGroupRulePayloadWithDefaults instantiates a new CreateSecurityGroupRulePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecurityGroupRulePayloadWithDefaults() *CreateSecurityGroupRulePayload {
	this := CreateSecurityGroupRulePayload{}
	var ethertype string = "IPv4"
	this.Ethertype = &ethertype
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetDescription() (res CreateSecurityGroupRulePayloadGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetDescriptionOk() (ret CreateSecurityGroupRulePayloadGetDescriptionRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateSecurityGroupRulePayload) SetDescription(v CreateSecurityGroupRulePayloadGetDescriptionRetType) {
	setCreateSecurityGroupRulePayloadGetDescriptionAttributeType(&o.Description, v)
}

// GetDirection returns the Direction field value
func (o *CreateSecurityGroupRulePayload) GetDirection() (ret CreateSecurityGroupRulePayloadGetDirectionRetType) {
	ret, _ = o.GetDirectionOk()
	return ret
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetDirectionOk() (ret CreateSecurityGroupRulePayloadGetDirectionRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetDirectionAttributeTypeOk(o.Direction)
}

// SetDirection sets field value
func (o *CreateSecurityGroupRulePayload) SetDirection(v CreateSecurityGroupRulePayloadGetDirectionRetType) {
	setCreateSecurityGroupRulePayloadGetDirectionAttributeType(&o.Direction, v)
}

// GetEthertype returns the Ethertype field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetEthertype() (res CreateSecurityGroupRulePayloadGetEthertypeRetType) {
	res, _ = o.GetEthertypeOk()
	return
}

// GetEthertypeOk returns a tuple with the Ethertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetEthertypeOk() (ret CreateSecurityGroupRulePayloadGetEthertypeRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetEthertypeAttributeTypeOk(o.Ethertype)
}

// HasEthertype returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasEthertype() bool {
	_, ok := o.GetEthertypeOk()
	return ok
}

// SetEthertype gets a reference to the given string and assigns it to the Ethertype field.
func (o *CreateSecurityGroupRulePayload) SetEthertype(v CreateSecurityGroupRulePayloadGetEthertypeRetType) {
	setCreateSecurityGroupRulePayloadGetEthertypeAttributeType(&o.Ethertype, v)
}

// GetIcmpParameters returns the IcmpParameters field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetIcmpParameters() (res CreateSecurityGroupRulePayloadGetIcmpParametersRetType) {
	res, _ = o.GetIcmpParametersOk()
	return
}

// GetIcmpParametersOk returns a tuple with the IcmpParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetIcmpParametersOk() (ret CreateSecurityGroupRulePayloadGetIcmpParametersRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetIcmpParametersAttributeTypeOk(o.IcmpParameters)
}

// HasIcmpParameters returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasIcmpParameters() bool {
	_, ok := o.GetIcmpParametersOk()
	return ok
}

// SetIcmpParameters gets a reference to the given ICMPParameters and assigns it to the IcmpParameters field.
func (o *CreateSecurityGroupRulePayload) SetIcmpParameters(v CreateSecurityGroupRulePayloadGetIcmpParametersRetType) {
	setCreateSecurityGroupRulePayloadGetIcmpParametersAttributeType(&o.IcmpParameters, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetId() (res CreateSecurityGroupRulePayloadGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetIdOk() (ret CreateSecurityGroupRulePayloadGetIdRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateSecurityGroupRulePayload) SetId(v CreateSecurityGroupRulePayloadGetIdRetType) {
	setCreateSecurityGroupRulePayloadGetIdAttributeType(&o.Id, v)
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetIpRange() (res CreateSecurityGroupRulePayloadGetIpRangeRetType) {
	res, _ = o.GetIpRangeOk()
	return
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetIpRangeOk() (ret CreateSecurityGroupRulePayloadGetIpRangeRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetIpRangeAttributeTypeOk(o.IpRange)
}

// HasIpRange returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasIpRange() bool {
	_, ok := o.GetIpRangeOk()
	return ok
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *CreateSecurityGroupRulePayload) SetIpRange(v CreateSecurityGroupRulePayloadGetIpRangeRetType) {
	setCreateSecurityGroupRulePayloadGetIpRangeAttributeType(&o.IpRange, v)
}

// GetPortRange returns the PortRange field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetPortRange() (res CreateSecurityGroupRulePayloadGetPortRangeRetType) {
	res, _ = o.GetPortRangeOk()
	return
}

// GetPortRangeOk returns a tuple with the PortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetPortRangeOk() (ret CreateSecurityGroupRulePayloadGetPortRangeRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetPortRangeAttributeTypeOk(o.PortRange)
}

// HasPortRange returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasPortRange() bool {
	_, ok := o.GetPortRangeOk()
	return ok
}

// SetPortRange gets a reference to the given PortRange and assigns it to the PortRange field.
func (o *CreateSecurityGroupRulePayload) SetPortRange(v CreateSecurityGroupRulePayloadGetPortRangeRetType) {
	setCreateSecurityGroupRulePayloadGetPortRangeAttributeType(&o.PortRange, v)
}

// GetRemoteSecurityGroupId returns the RemoteSecurityGroupId field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetRemoteSecurityGroupId() (res CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdRetType) {
	res, _ = o.GetRemoteSecurityGroupIdOk()
	return
}

// GetRemoteSecurityGroupIdOk returns a tuple with the RemoteSecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetRemoteSecurityGroupIdOk() (ret CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeTypeOk(o.RemoteSecurityGroupId)
}

// HasRemoteSecurityGroupId returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasRemoteSecurityGroupId() bool {
	_, ok := o.GetRemoteSecurityGroupIdOk()
	return ok
}

// SetRemoteSecurityGroupId gets a reference to the given string and assigns it to the RemoteSecurityGroupId field.
func (o *CreateSecurityGroupRulePayload) SetRemoteSecurityGroupId(v CreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdRetType) {
	setCreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeType(&o.RemoteSecurityGroupId, v)
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetSecurityGroupId() (res CreateSecurityGroupRulePayloadGetSecurityGroupIdRetType) {
	res, _ = o.GetSecurityGroupIdOk()
	return
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetSecurityGroupIdOk() (ret CreateSecurityGroupRulePayloadGetSecurityGroupIdRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeTypeOk(o.SecurityGroupId)
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasSecurityGroupId() bool {
	_, ok := o.GetSecurityGroupIdOk()
	return ok
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *CreateSecurityGroupRulePayload) SetSecurityGroupId(v CreateSecurityGroupRulePayloadGetSecurityGroupIdRetType) {
	setCreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeType(&o.SecurityGroupId, v)
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *CreateSecurityGroupRulePayload) GetProtocol() (res CreateSecurityGroupRulePayloadGetProtocolRetType) {
	res, _ = o.GetProtocolOk()
	return
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRulePayload) GetProtocolOk() (ret CreateSecurityGroupRulePayloadGetProtocolRetType, ok bool) {
	return getCreateSecurityGroupRulePayloadGetProtocolAttributeTypeOk(o.Protocol)
}

// HasProtocol returns a boolean if a field has been set.
func (o *CreateSecurityGroupRulePayload) HasProtocol() bool {
	_, ok := o.GetProtocolOk()
	return ok
}

// SetProtocol gets a reference to the given CreateProtocol and assigns it to the Protocol field.
func (o *CreateSecurityGroupRulePayload) SetProtocol(v CreateSecurityGroupRulePayloadGetProtocolRetType) {
	setCreateSecurityGroupRulePayloadGetProtocolAttributeType(&o.Protocol, v)
}

func (o CreateSecurityGroupRulePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateSecurityGroupRulePayloadGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetDirectionAttributeTypeOk(o.Direction); ok {
		toSerialize["Direction"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetEthertypeAttributeTypeOk(o.Ethertype); ok {
		toSerialize["Ethertype"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetIcmpParametersAttributeTypeOk(o.IcmpParameters); ok {
		toSerialize["IcmpParameters"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetIpRangeAttributeTypeOk(o.IpRange); ok {
		toSerialize["IpRange"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetPortRangeAttributeTypeOk(o.PortRange); ok {
		toSerialize["PortRange"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetRemoteSecurityGroupIdAttributeTypeOk(o.RemoteSecurityGroupId); ok {
		toSerialize["RemoteSecurityGroupId"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetSecurityGroupIdAttributeTypeOk(o.SecurityGroupId); ok {
		toSerialize["SecurityGroupId"] = val
	}
	if val, ok := getCreateSecurityGroupRulePayloadGetProtocolAttributeTypeOk(o.Protocol); ok {
		toSerialize["Protocol"] = val
	}
	return toSerialize, nil
}

type NullableCreateSecurityGroupRulePayload struct {
	value *CreateSecurityGroupRulePayload
	isSet bool
}

func (v NullableCreateSecurityGroupRulePayload) Get() *CreateSecurityGroupRulePayload {
	return v.value
}

func (v *NullableCreateSecurityGroupRulePayload) Set(val *CreateSecurityGroupRulePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecurityGroupRulePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecurityGroupRulePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecurityGroupRulePayload(val *CreateSecurityGroupRulePayload) *NullableCreateSecurityGroupRulePayload {
	return &NullableCreateSecurityGroupRulePayload{value: val, isSet: true}
}

func (v NullableCreateSecurityGroupRulePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecurityGroupRulePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
