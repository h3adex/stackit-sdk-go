/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// checks if the RecordSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordSet{}

/*
	types and functions for active
*/

// isBoolean
type RecordSetgetActiveAttributeType = *bool
type RecordSetgetActiveArgType = bool
type RecordSetgetActiveRetType = bool

func getRecordSetgetActiveAttributeTypeOk(arg RecordSetgetActiveAttributeType) (ret RecordSetgetActiveRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetgetActiveAttributeType(arg *RecordSetgetActiveAttributeType, val RecordSetgetActiveRetType) {
	*arg = &val
}

/*
	types and functions for comment
*/

// isNotNullableString
type RecordSetGetCommentAttributeType = *string

func getRecordSetGetCommentAttributeTypeOk(arg RecordSetGetCommentAttributeType) (ret RecordSetGetCommentRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetCommentAttributeType(arg *RecordSetGetCommentAttributeType, val RecordSetGetCommentRetType) {
	*arg = &val
}

type RecordSetGetCommentArgType = string
type RecordSetGetCommentRetType = string

/*
	types and functions for creationFinished
*/

// isNotNullableString
type RecordSetGetCreationFinishedAttributeType = *string

func getRecordSetGetCreationFinishedAttributeTypeOk(arg RecordSetGetCreationFinishedAttributeType) (ret RecordSetGetCreationFinishedRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetCreationFinishedAttributeType(arg *RecordSetGetCreationFinishedAttributeType, val RecordSetGetCreationFinishedRetType) {
	*arg = &val
}

type RecordSetGetCreationFinishedArgType = string
type RecordSetGetCreationFinishedRetType = string

/*
	types and functions for creationStarted
*/

// isNotNullableString
type RecordSetGetCreationStartedAttributeType = *string

func getRecordSetGetCreationStartedAttributeTypeOk(arg RecordSetGetCreationStartedAttributeType) (ret RecordSetGetCreationStartedRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetCreationStartedAttributeType(arg *RecordSetGetCreationStartedAttributeType, val RecordSetGetCreationStartedRetType) {
	*arg = &val
}

type RecordSetGetCreationStartedArgType = string
type RecordSetGetCreationStartedRetType = string

/*
	types and functions for error
*/

// isNotNullableString
type RecordSetGetErrorAttributeType = *string

func getRecordSetGetErrorAttributeTypeOk(arg RecordSetGetErrorAttributeType) (ret RecordSetGetErrorRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetErrorAttributeType(arg *RecordSetGetErrorAttributeType, val RecordSetGetErrorRetType) {
	*arg = &val
}

type RecordSetGetErrorArgType = string
type RecordSetGetErrorRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type RecordSetGetIdAttributeType = *string

func getRecordSetGetIdAttributeTypeOk(arg RecordSetGetIdAttributeType) (ret RecordSetGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetIdAttributeType(arg *RecordSetGetIdAttributeType, val RecordSetGetIdRetType) {
	*arg = &val
}

type RecordSetGetIdArgType = string
type RecordSetGetIdRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type RecordSetGetNameAttributeType = *string

func getRecordSetGetNameAttributeTypeOk(arg RecordSetGetNameAttributeType) (ret RecordSetGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetNameAttributeType(arg *RecordSetGetNameAttributeType, val RecordSetGetNameRetType) {
	*arg = &val
}

type RecordSetGetNameArgType = string
type RecordSetGetNameRetType = string

/*
	types and functions for records
*/

// isArray
type RecordSetGetRecordsAttributeType = *[]Record
type RecordSetGetRecordsArgType = []Record
type RecordSetGetRecordsRetType = []Record

func getRecordSetGetRecordsAttributeTypeOk(arg RecordSetGetRecordsAttributeType) (ret RecordSetGetRecordsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetRecordsAttributeType(arg *RecordSetGetRecordsAttributeType, val RecordSetGetRecordsRetType) {
	*arg = &val
}

/*
	types and functions for state
*/

// isEnum

// RecordSetState record set state
// value type for enums
type RecordSetState string

// List of State
const (
	RECORDSETSTATE_CREATING         RecordSetState = "CREATING"
	RECORDSETSTATE_CREATE_SUCCEEDED RecordSetState = "CREATE_SUCCEEDED"
	RECORDSETSTATE_CREATE_FAILED    RecordSetState = "CREATE_FAILED"
	RECORDSETSTATE_DELETING         RecordSetState = "DELETING"
	RECORDSETSTATE_DELETE_SUCCEEDED RecordSetState = "DELETE_SUCCEEDED"
	RECORDSETSTATE_DELETE_FAILED    RecordSetState = "DELETE_FAILED"
	RECORDSETSTATE_UPDATING         RecordSetState = "UPDATING"
	RECORDSETSTATE_UPDATE_SUCCEEDED RecordSetState = "UPDATE_SUCCEEDED"
	RECORDSETSTATE_UPDATE_FAILED    RecordSetState = "UPDATE_FAILED"
)

// All allowed values of RecordSet enum
var AllowedRecordSetStateEnumValues = []RecordSetState{
	"CREATING",
	"CREATE_SUCCEEDED",
	"CREATE_FAILED",
	"DELETING",
	"DELETE_SUCCEEDED",
	"DELETE_FAILED",
	"UPDATING",
	"UPDATE_SUCCEEDED",
	"UPDATE_FAILED",
}

func (v *RecordSetState) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson RecordSetState
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
	enumTypeValue := RecordSetState(value)
	for _, existing := range AllowedRecordSetStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordSet", value)
}

// NewRecordSetStateFromValue returns a pointer to a valid RecordSetState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecordSetStateFromValue(v RecordSetState) (*RecordSetState, error) {
	ev := RecordSetState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecordSetState: valid values are %v", v, AllowedRecordSetStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecordSetState) IsValid() bool {
	for _, existing := range AllowedRecordSetStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateState value
func (v RecordSetState) Ptr() *RecordSetState {
	return &v
}

type NullableRecordSetState struct {
	value *RecordSetState
	isSet bool
}

func (v NullableRecordSetState) Get() *RecordSetState {
	return v.value
}

func (v *NullableRecordSetState) Set(val *RecordSetState) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordSetState) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordSetState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordSetState(val *RecordSetState) *NullableRecordSetState {
	return &NullableRecordSetState{value: val, isSet: true}
}

func (v NullableRecordSetState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordSetState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type RecordSetGetStateAttributeType = *RecordSetState
type RecordSetGetStateArgType = RecordSetState
type RecordSetGetStateRetType = RecordSetState

func getRecordSetGetStateAttributeTypeOk(arg RecordSetGetStateAttributeType) (ret RecordSetGetStateRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetStateAttributeType(arg *RecordSetGetStateAttributeType, val RecordSetGetStateRetType) {
	*arg = &val
}

/*
	types and functions for ttl
*/

// isInteger
type RecordSetGetTtlAttributeType = *int64
type RecordSetGetTtlArgType = int64
type RecordSetGetTtlRetType = int64

func getRecordSetGetTtlAttributeTypeOk(arg RecordSetGetTtlAttributeType) (ret RecordSetGetTtlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetTtlAttributeType(arg *RecordSetGetTtlAttributeType, val RecordSetGetTtlRetType) {
	*arg = &val
}

/*
	types and functions for type
*/

// isEnum

// RecordSetTypes record set type
// value type for enums
type RecordSetTypes string

// List of Type
const (
	RECORDSETTYPE_A      RecordSetTypes = "A"
	RECORDSETTYPE_AAAA   RecordSetTypes = "AAAA"
	RECORDSETTYPE_SOA    RecordSetTypes = "SOA"
	RECORDSETTYPE_CNAME  RecordSetTypes = "CNAME"
	RECORDSETTYPE_NS     RecordSetTypes = "NS"
	RECORDSETTYPE_MX     RecordSetTypes = "MX"
	RECORDSETTYPE_TXT    RecordSetTypes = "TXT"
	RECORDSETTYPE_SRV    RecordSetTypes = "SRV"
	RECORDSETTYPE_PTR    RecordSetTypes = "PTR"
	RECORDSETTYPE_ALIAS  RecordSetTypes = "ALIAS"
	RECORDSETTYPE_DNAME  RecordSetTypes = "DNAME"
	RECORDSETTYPE_CAA    RecordSetTypes = "CAA"
	RECORDSETTYPE_DNSKEY RecordSetTypes = "DNSKEY"
	RECORDSETTYPE_DS     RecordSetTypes = "DS"
	RECORDSETTYPE_LOC    RecordSetTypes = "LOC"
	RECORDSETTYPE_NAPTR  RecordSetTypes = "NAPTR"
	RECORDSETTYPE_SSHFP  RecordSetTypes = "SSHFP"
	RECORDSETTYPE_TLSA   RecordSetTypes = "TLSA"
	RECORDSETTYPE_URI    RecordSetTypes = "URI"
	RECORDSETTYPE_CERT   RecordSetTypes = "CERT"
	RECORDSETTYPE_SVCB   RecordSetTypes = "SVCB"
	RECORDSETTYPE_TYPE   RecordSetTypes = "TYPE"
	RECORDSETTYPE_CSYNC  RecordSetTypes = "CSYNC"
	RECORDSETTYPE_HINFO  RecordSetTypes = "HINFO"
	RECORDSETTYPE_HTTPS  RecordSetTypes = "HTTPS"
)

// All allowed values of RecordSet enum
var AllowedRecordSetTypesEnumValues = []RecordSetTypes{
	"A",
	"AAAA",
	"SOA",
	"CNAME",
	"NS",
	"MX",
	"TXT",
	"SRV",
	"PTR",
	"ALIAS",
	"DNAME",
	"CAA",
	"DNSKEY",
	"DS",
	"LOC",
	"NAPTR",
	"SSHFP",
	"TLSA",
	"URI",
	"CERT",
	"SVCB",
	"TYPE",
	"CSYNC",
	"HINFO",
	"HTTPS",
}

func (v *RecordSetTypes) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson RecordSetTypes
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
	enumTypeValue := RecordSetTypes(value)
	for _, existing := range AllowedRecordSetTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordSet", value)
}

// NewRecordSetTypesFromValue returns a pointer to a valid RecordSetTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecordSetTypesFromValue(v RecordSetTypes) (*RecordSetTypes, error) {
	ev := RecordSetTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecordSetTypes: valid values are %v", v, AllowedRecordSetTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecordSetTypes) IsValid() bool {
	for _, existing := range AllowedRecordSetTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypeTypes value
func (v RecordSetTypes) Ptr() *RecordSetTypes {
	return &v
}

type NullableRecordSetTypes struct {
	value *RecordSetTypes
	isSet bool
}

func (v NullableRecordSetTypes) Get() *RecordSetTypes {
	return v.value
}

func (v *NullableRecordSetTypes) Set(val *RecordSetTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordSetTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordSetTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordSetTypes(val *RecordSetTypes) *NullableRecordSetTypes {
	return &NullableRecordSetTypes{value: val, isSet: true}
}

func (v NullableRecordSetTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordSetTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type RecordSetGetTypeAttributeType = *RecordSetTypes
type RecordSetGetTypeArgType = RecordSetTypes
type RecordSetGetTypeRetType = RecordSetTypes

func getRecordSetGetTypeAttributeTypeOk(arg RecordSetGetTypeAttributeType) (ret RecordSetGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetTypeAttributeType(arg *RecordSetGetTypeAttributeType, val RecordSetGetTypeRetType) {
	*arg = &val
}

/*
	types and functions for updateFinished
*/

// isNotNullableString
type RecordSetGetUpdateFinishedAttributeType = *string

func getRecordSetGetUpdateFinishedAttributeTypeOk(arg RecordSetGetUpdateFinishedAttributeType) (ret RecordSetGetUpdateFinishedRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetUpdateFinishedAttributeType(arg *RecordSetGetUpdateFinishedAttributeType, val RecordSetGetUpdateFinishedRetType) {
	*arg = &val
}

type RecordSetGetUpdateFinishedArgType = string
type RecordSetGetUpdateFinishedRetType = string

/*
	types and functions for updateStarted
*/

// isNotNullableString
type RecordSetGetUpdateStartedAttributeType = *string

func getRecordSetGetUpdateStartedAttributeTypeOk(arg RecordSetGetUpdateStartedAttributeType) (ret RecordSetGetUpdateStartedRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetGetUpdateStartedAttributeType(arg *RecordSetGetUpdateStartedAttributeType, val RecordSetGetUpdateStartedRetType) {
	*arg = &val
}

type RecordSetGetUpdateStartedArgType = string
type RecordSetGetUpdateStartedRetType = string

// RecordSet RRSet.
type RecordSet struct {
	// if the record set is active or not
	Active RecordSetgetActiveAttributeType `json:"active,omitempty"`
	// comment
	Comment RecordSetGetCommentAttributeType `json:"comment,omitempty"`
	// when record set creation finished
	// REQUIRED
	CreationFinished RecordSetGetCreationFinishedAttributeType `json:"creationFinished" required:"true"`
	// when record set creation started
	// REQUIRED
	CreationStarted RecordSetGetCreationStartedAttributeType `json:"creationStarted" required:"true"`
	// Error shows error in case create/update/delete failed
	Error RecordSetGetErrorAttributeType `json:"error,omitempty"`
	// rr set id
	// REQUIRED
	Id RecordSetGetIdAttributeType `json:"id" required:"true"`
	// name of the record which should be a valid domain according to rfc1035 Section 2.3.4. For APEX records (same as zone name), the zone name itself has to be put in here.
	// REQUIRED
	Name RecordSetGetNameAttributeType `json:"name" required:"true"`
	// records
	// REQUIRED
	Records RecordSetGetRecordsAttributeType `json:"records" required:"true"`
	// record set state
	// REQUIRED
	State RecordSetGetStateAttributeType `json:"state" required:"true"`
	// time to live
	// Can be cast to int32 without loss of precision.
	// REQUIRED
	Ttl RecordSetGetTtlAttributeType `json:"ttl" required:"true"`
	// record set type
	// REQUIRED
	Type RecordSetGetTypeAttributeType `json:"type" required:"true"`
	// when record set update/deletion finished
	// REQUIRED
	UpdateFinished RecordSetGetUpdateFinishedAttributeType `json:"updateFinished" required:"true"`
	// when record set update/deletion started
	// REQUIRED
	UpdateStarted RecordSetGetUpdateStartedAttributeType `json:"updateStarted" required:"true"`
}

type _RecordSet RecordSet

// NewRecordSet instantiates a new RecordSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordSet(creationFinished RecordSetGetCreationFinishedArgType, creationStarted RecordSetGetCreationStartedArgType, id RecordSetGetIdArgType, name RecordSetGetNameArgType, records RecordSetGetRecordsArgType, state RecordSetGetStateArgType, ttl RecordSetGetTtlArgType, types RecordSetGetTypeArgType, updateFinished RecordSetGetUpdateFinishedArgType, updateStarted RecordSetGetUpdateStartedArgType) *RecordSet {
	this := RecordSet{}
	setRecordSetGetCreationFinishedAttributeType(&this.CreationFinished, creationFinished)
	setRecordSetGetCreationStartedAttributeType(&this.CreationStarted, creationStarted)
	setRecordSetGetIdAttributeType(&this.Id, id)
	setRecordSetGetNameAttributeType(&this.Name, name)
	setRecordSetGetRecordsAttributeType(&this.Records, records)
	setRecordSetGetStateAttributeType(&this.State, state)
	setRecordSetGetTtlAttributeType(&this.Ttl, ttl)
	setRecordSetGetTypeAttributeType(&this.Type, types)
	setRecordSetGetUpdateFinishedAttributeType(&this.UpdateFinished, updateFinished)
	setRecordSetGetUpdateStartedAttributeType(&this.UpdateStarted, updateStarted)
	return &this
}

// NewRecordSetWithDefaults instantiates a new RecordSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordSetWithDefaults() *RecordSet {
	this := RecordSet{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RecordSet) GetActive() (res RecordSetgetActiveRetType) {
	res, _ = o.GetActiveOk()
	return
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordSet) GetActiveOk() (ret RecordSetgetActiveRetType, ok bool) {
	return getRecordSetgetActiveAttributeTypeOk(o.Active)
}

// HasActive returns a boolean if a field has been set.
func (o *RecordSet) HasActive() bool {
	_, ok := o.GetActiveOk()
	return ok
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RecordSet) SetActive(v RecordSetgetActiveRetType) {
	setRecordSetgetActiveAttributeType(&o.Active, v)
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RecordSet) GetComment() (res RecordSetGetCommentRetType) {
	res, _ = o.GetCommentOk()
	return
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordSet) GetCommentOk() (ret RecordSetGetCommentRetType, ok bool) {
	return getRecordSetGetCommentAttributeTypeOk(o.Comment)
}

// HasComment returns a boolean if a field has been set.
func (o *RecordSet) HasComment() bool {
	_, ok := o.GetCommentOk()
	return ok
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RecordSet) SetComment(v RecordSetGetCommentRetType) {
	setRecordSetGetCommentAttributeType(&o.Comment, v)
}

// GetCreationFinished returns the CreationFinished field value
func (o *RecordSet) GetCreationFinished() (ret RecordSetGetCreationFinishedRetType) {
	ret, _ = o.GetCreationFinishedOk()
	return ret
}

// GetCreationFinishedOk returns a tuple with the CreationFinished field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetCreationFinishedOk() (ret RecordSetGetCreationFinishedRetType, ok bool) {
	return getRecordSetGetCreationFinishedAttributeTypeOk(o.CreationFinished)
}

// SetCreationFinished sets field value
func (o *RecordSet) SetCreationFinished(v RecordSetGetCreationFinishedRetType) {
	setRecordSetGetCreationFinishedAttributeType(&o.CreationFinished, v)
}

// GetCreationStarted returns the CreationStarted field value
func (o *RecordSet) GetCreationStarted() (ret RecordSetGetCreationStartedRetType) {
	ret, _ = o.GetCreationStartedOk()
	return ret
}

// GetCreationStartedOk returns a tuple with the CreationStarted field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetCreationStartedOk() (ret RecordSetGetCreationStartedRetType, ok bool) {
	return getRecordSetGetCreationStartedAttributeTypeOk(o.CreationStarted)
}

// SetCreationStarted sets field value
func (o *RecordSet) SetCreationStarted(v RecordSetGetCreationStartedRetType) {
	setRecordSetGetCreationStartedAttributeType(&o.CreationStarted, v)
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RecordSet) GetError() (res RecordSetGetErrorRetType) {
	res, _ = o.GetErrorOk()
	return
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordSet) GetErrorOk() (ret RecordSetGetErrorRetType, ok bool) {
	return getRecordSetGetErrorAttributeTypeOk(o.Error)
}

// HasError returns a boolean if a field has been set.
func (o *RecordSet) HasError() bool {
	_, ok := o.GetErrorOk()
	return ok
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RecordSet) SetError(v RecordSetGetErrorRetType) {
	setRecordSetGetErrorAttributeType(&o.Error, v)
}

// GetId returns the Id field value
func (o *RecordSet) GetId() (ret RecordSetGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetIdOk() (ret RecordSetGetIdRetType, ok bool) {
	return getRecordSetGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *RecordSet) SetId(v RecordSetGetIdRetType) {
	setRecordSetGetIdAttributeType(&o.Id, v)
}

// GetName returns the Name field value
func (o *RecordSet) GetName() (ret RecordSetGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetNameOk() (ret RecordSetGetNameRetType, ok bool) {
	return getRecordSetGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *RecordSet) SetName(v RecordSetGetNameRetType) {
	setRecordSetGetNameAttributeType(&o.Name, v)
}

// GetRecords returns the Records field value
func (o *RecordSet) GetRecords() (ret RecordSetGetRecordsRetType) {
	ret, _ = o.GetRecordsOk()
	return ret
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetRecordsOk() (ret RecordSetGetRecordsRetType, ok bool) {
	return getRecordSetGetRecordsAttributeTypeOk(o.Records)
}

// SetRecords sets field value
func (o *RecordSet) SetRecords(v RecordSetGetRecordsRetType) {
	setRecordSetGetRecordsAttributeType(&o.Records, v)
}

// GetState returns the State field value
func (o *RecordSet) GetState() (ret RecordSetGetStateRetType) {
	ret, _ = o.GetStateOk()
	return ret
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetStateOk() (ret RecordSetGetStateRetType, ok bool) {
	return getRecordSetGetStateAttributeTypeOk(o.State)
}

// SetState sets field value
func (o *RecordSet) SetState(v RecordSetGetStateRetType) {
	setRecordSetGetStateAttributeType(&o.State, v)
}

// GetTtl returns the Ttl field value
func (o *RecordSet) GetTtl() (ret RecordSetGetTtlRetType) {
	ret, _ = o.GetTtlOk()
	return ret
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetTtlOk() (ret RecordSetGetTtlRetType, ok bool) {
	return getRecordSetGetTtlAttributeTypeOk(o.Ttl)
}

// SetTtl sets field value
func (o *RecordSet) SetTtl(v RecordSetGetTtlRetType) {
	setRecordSetGetTtlAttributeType(&o.Ttl, v)
}

// GetType returns the Type field value
func (o *RecordSet) GetType() (ret RecordSetGetTypeRetType) {
	ret, _ = o.GetTypeOk()
	return ret
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetTypeOk() (ret RecordSetGetTypeRetType, ok bool) {
	return getRecordSetGetTypeAttributeTypeOk(o.Type)
}

// SetType sets field value
func (o *RecordSet) SetType(v RecordSetGetTypeRetType) {
	setRecordSetGetTypeAttributeType(&o.Type, v)
}

// GetUpdateFinished returns the UpdateFinished field value
func (o *RecordSet) GetUpdateFinished() (ret RecordSetGetUpdateFinishedRetType) {
	ret, _ = o.GetUpdateFinishedOk()
	return ret
}

// GetUpdateFinishedOk returns a tuple with the UpdateFinished field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetUpdateFinishedOk() (ret RecordSetGetUpdateFinishedRetType, ok bool) {
	return getRecordSetGetUpdateFinishedAttributeTypeOk(o.UpdateFinished)
}

// SetUpdateFinished sets field value
func (o *RecordSet) SetUpdateFinished(v RecordSetGetUpdateFinishedRetType) {
	setRecordSetGetUpdateFinishedAttributeType(&o.UpdateFinished, v)
}

// GetUpdateStarted returns the UpdateStarted field value
func (o *RecordSet) GetUpdateStarted() (ret RecordSetGetUpdateStartedRetType) {
	ret, _ = o.GetUpdateStartedOk()
	return ret
}

// GetUpdateStartedOk returns a tuple with the UpdateStarted field value
// and a boolean to check if the value has been set.
func (o *RecordSet) GetUpdateStartedOk() (ret RecordSetGetUpdateStartedRetType, ok bool) {
	return getRecordSetGetUpdateStartedAttributeTypeOk(o.UpdateStarted)
}

// SetUpdateStarted sets field value
func (o *RecordSet) SetUpdateStarted(v RecordSetGetUpdateStartedRetType) {
	setRecordSetGetUpdateStartedAttributeType(&o.UpdateStarted, v)
}

func (o RecordSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRecordSetgetActiveAttributeTypeOk(o.Active); ok {
		toSerialize["Active"] = val
	}
	if val, ok := getRecordSetGetCommentAttributeTypeOk(o.Comment); ok {
		toSerialize["Comment"] = val
	}
	if val, ok := getRecordSetGetCreationFinishedAttributeTypeOk(o.CreationFinished); ok {
		toSerialize["CreationFinished"] = val
	}
	if val, ok := getRecordSetGetCreationStartedAttributeTypeOk(o.CreationStarted); ok {
		toSerialize["CreationStarted"] = val
	}
	if val, ok := getRecordSetGetErrorAttributeTypeOk(o.Error); ok {
		toSerialize["Error"] = val
	}
	if val, ok := getRecordSetGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getRecordSetGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getRecordSetGetRecordsAttributeTypeOk(o.Records); ok {
		toSerialize["Records"] = val
	}
	if val, ok := getRecordSetGetStateAttributeTypeOk(o.State); ok {
		toSerialize["State"] = val
	}
	if val, ok := getRecordSetGetTtlAttributeTypeOk(o.Ttl); ok {
		toSerialize["Ttl"] = val
	}
	if val, ok := getRecordSetGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	if val, ok := getRecordSetGetUpdateFinishedAttributeTypeOk(o.UpdateFinished); ok {
		toSerialize["UpdateFinished"] = val
	}
	if val, ok := getRecordSetGetUpdateStartedAttributeTypeOk(o.UpdateStarted); ok {
		toSerialize["UpdateStarted"] = val
	}
	return toSerialize, nil
}

type NullableRecordSet struct {
	value *RecordSet
	isSet bool
}

func (v NullableRecordSet) Get() *RecordSet {
	return v.value
}

func (v *NullableRecordSet) Set(val *RecordSet) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordSet) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordSet(val *RecordSet) *NullableRecordSet {
	return &NullableRecordSet{value: val, isSet: true}
}

func (v NullableRecordSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
