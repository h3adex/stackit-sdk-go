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

// checks if the QuotaList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaList{}

/*
	types and functions for backupGigabytes
*/

// isModel
type QuotaListGetBackupGigabytesAttributeType = *QuotaListBackupGigabytes
type QuotaListGetBackupGigabytesArgType = QuotaListBackupGigabytes
type QuotaListGetBackupGigabytesRetType = QuotaListBackupGigabytes

func getQuotaListGetBackupGigabytesAttributeTypeOk(arg QuotaListGetBackupGigabytesAttributeType) (ret QuotaListGetBackupGigabytesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetBackupGigabytesAttributeType(arg *QuotaListGetBackupGigabytesAttributeType, val QuotaListGetBackupGigabytesRetType) {
	*arg = &val
}

/*
	types and functions for backups
*/

// isModel
type QuotaListGetBackupsAttributeType = *QuotaListBackups
type QuotaListGetBackupsArgType = QuotaListBackups
type QuotaListGetBackupsRetType = QuotaListBackups

func getQuotaListGetBackupsAttributeTypeOk(arg QuotaListGetBackupsAttributeType) (ret QuotaListGetBackupsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetBackupsAttributeType(arg *QuotaListGetBackupsAttributeType, val QuotaListGetBackupsRetType) {
	*arg = &val
}

/*
	types and functions for gigabytes
*/

// isModel
type QuotaListGetGigabytesAttributeType = *QuotaListGigabytes
type QuotaListGetGigabytesArgType = QuotaListGigabytes
type QuotaListGetGigabytesRetType = QuotaListGigabytes

func getQuotaListGetGigabytesAttributeTypeOk(arg QuotaListGetGigabytesAttributeType) (ret QuotaListGetGigabytesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetGigabytesAttributeType(arg *QuotaListGetGigabytesAttributeType, val QuotaListGetGigabytesRetType) {
	*arg = &val
}

/*
	types and functions for networks
*/

// isModel
type QuotaListGetNetworksAttributeType = *QuotaListNetworks
type QuotaListGetNetworksArgType = QuotaListNetworks
type QuotaListGetNetworksRetType = QuotaListNetworks

func getQuotaListGetNetworksAttributeTypeOk(arg QuotaListGetNetworksAttributeType) (ret QuotaListGetNetworksRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetNetworksAttributeType(arg *QuotaListGetNetworksAttributeType, val QuotaListGetNetworksRetType) {
	*arg = &val
}

/*
	types and functions for nics
*/

// isModel
type QuotaListGetNicsAttributeType = *QuotaListNics
type QuotaListGetNicsArgType = QuotaListNics
type QuotaListGetNicsRetType = QuotaListNics

func getQuotaListGetNicsAttributeTypeOk(arg QuotaListGetNicsAttributeType) (ret QuotaListGetNicsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetNicsAttributeType(arg *QuotaListGetNicsAttributeType, val QuotaListGetNicsRetType) {
	*arg = &val
}

/*
	types and functions for publicIps
*/

// isModel
type QuotaListGetPublicIpsAttributeType = *QuotaListPublicIps
type QuotaListGetPublicIpsArgType = QuotaListPublicIps
type QuotaListGetPublicIpsRetType = QuotaListPublicIps

func getQuotaListGetPublicIpsAttributeTypeOk(arg QuotaListGetPublicIpsAttributeType) (ret QuotaListGetPublicIpsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetPublicIpsAttributeType(arg *QuotaListGetPublicIpsAttributeType, val QuotaListGetPublicIpsRetType) {
	*arg = &val
}

/*
	types and functions for ram
*/

// isModel
type QuotaListGetRamAttributeType = *QuotaListRam
type QuotaListGetRamArgType = QuotaListRam
type QuotaListGetRamRetType = QuotaListRam

func getQuotaListGetRamAttributeTypeOk(arg QuotaListGetRamAttributeType) (ret QuotaListGetRamRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetRamAttributeType(arg *QuotaListGetRamAttributeType, val QuotaListGetRamRetType) {
	*arg = &val
}

/*
	types and functions for securityGroupRules
*/

// isModel
type QuotaListGetSecurityGroupRulesAttributeType = *QuotaListSecurityGroupRules
type QuotaListGetSecurityGroupRulesArgType = QuotaListSecurityGroupRules
type QuotaListGetSecurityGroupRulesRetType = QuotaListSecurityGroupRules

func getQuotaListGetSecurityGroupRulesAttributeTypeOk(arg QuotaListGetSecurityGroupRulesAttributeType) (ret QuotaListGetSecurityGroupRulesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetSecurityGroupRulesAttributeType(arg *QuotaListGetSecurityGroupRulesAttributeType, val QuotaListGetSecurityGroupRulesRetType) {
	*arg = &val
}

/*
	types and functions for securityGroups
*/

// isModel
type QuotaListGetSecurityGroupsAttributeType = *QuotaListSecurityGroups
type QuotaListGetSecurityGroupsArgType = QuotaListSecurityGroups
type QuotaListGetSecurityGroupsRetType = QuotaListSecurityGroups

func getQuotaListGetSecurityGroupsAttributeTypeOk(arg QuotaListGetSecurityGroupsAttributeType) (ret QuotaListGetSecurityGroupsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetSecurityGroupsAttributeType(arg *QuotaListGetSecurityGroupsAttributeType, val QuotaListGetSecurityGroupsRetType) {
	*arg = &val
}

/*
	types and functions for snapshots
*/

// isModel
type QuotaListGetSnapshotsAttributeType = *QuotaListSnapshots
type QuotaListGetSnapshotsArgType = QuotaListSnapshots
type QuotaListGetSnapshotsRetType = QuotaListSnapshots

func getQuotaListGetSnapshotsAttributeTypeOk(arg QuotaListGetSnapshotsAttributeType) (ret QuotaListGetSnapshotsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetSnapshotsAttributeType(arg *QuotaListGetSnapshotsAttributeType, val QuotaListGetSnapshotsRetType) {
	*arg = &val
}

/*
	types and functions for vcpu
*/

// isModel
type QuotaListGetVcpuAttributeType = *QuotaListVcpu
type QuotaListGetVcpuArgType = QuotaListVcpu
type QuotaListGetVcpuRetType = QuotaListVcpu

func getQuotaListGetVcpuAttributeTypeOk(arg QuotaListGetVcpuAttributeType) (ret QuotaListGetVcpuRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetVcpuAttributeType(arg *QuotaListGetVcpuAttributeType, val QuotaListGetVcpuRetType) {
	*arg = &val
}

/*
	types and functions for volumes
*/

// isModel
type QuotaListGetVolumesAttributeType = *QuotaListVolumes
type QuotaListGetVolumesArgType = QuotaListVolumes
type QuotaListGetVolumesRetType = QuotaListVolumes

func getQuotaListGetVolumesAttributeTypeOk(arg QuotaListGetVolumesAttributeType) (ret QuotaListGetVolumesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListGetVolumesAttributeType(arg *QuotaListGetVolumesAttributeType, val QuotaListGetVolumesRetType) {
	*arg = &val
}

// QuotaList Object that represents the quotas for a project.
type QuotaList struct {
	// REQUIRED
	BackupGigabytes QuotaListGetBackupGigabytesAttributeType `json:"backupGigabytes" required:"true"`
	// REQUIRED
	Backups QuotaListGetBackupsAttributeType `json:"backups" required:"true"`
	// REQUIRED
	Gigabytes QuotaListGetGigabytesAttributeType `json:"gigabytes" required:"true"`
	// REQUIRED
	Networks QuotaListGetNetworksAttributeType `json:"networks" required:"true"`
	// REQUIRED
	Nics QuotaListGetNicsAttributeType `json:"nics" required:"true"`
	// REQUIRED
	PublicIps QuotaListGetPublicIpsAttributeType `json:"publicIps" required:"true"`
	// REQUIRED
	Ram QuotaListGetRamAttributeType `json:"ram" required:"true"`
	// REQUIRED
	SecurityGroupRules QuotaListGetSecurityGroupRulesAttributeType `json:"securityGroupRules" required:"true"`
	// REQUIRED
	SecurityGroups QuotaListGetSecurityGroupsAttributeType `json:"securityGroups" required:"true"`
	// REQUIRED
	Snapshots QuotaListGetSnapshotsAttributeType `json:"snapshots" required:"true"`
	// REQUIRED
	Vcpu QuotaListGetVcpuAttributeType `json:"vcpu" required:"true"`
	// REQUIRED
	Volumes QuotaListGetVolumesAttributeType `json:"volumes" required:"true"`
}

type _QuotaList QuotaList

// NewQuotaList instantiates a new QuotaList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaList(backupGigabytes QuotaListGetBackupGigabytesArgType, backups QuotaListGetBackupsArgType, gigabytes QuotaListGetGigabytesArgType, networks QuotaListGetNetworksArgType, nics QuotaListGetNicsArgType, publicIps QuotaListGetPublicIpsArgType, ram QuotaListGetRamArgType, securityGroupRules QuotaListGetSecurityGroupRulesArgType, securityGroups QuotaListGetSecurityGroupsArgType, snapshots QuotaListGetSnapshotsArgType, vcpu QuotaListGetVcpuArgType, volumes QuotaListGetVolumesArgType) *QuotaList {
	this := QuotaList{}
	setQuotaListGetBackupGigabytesAttributeType(&this.BackupGigabytes, backupGigabytes)
	setQuotaListGetBackupsAttributeType(&this.Backups, backups)
	setQuotaListGetGigabytesAttributeType(&this.Gigabytes, gigabytes)
	setQuotaListGetNetworksAttributeType(&this.Networks, networks)
	setQuotaListGetNicsAttributeType(&this.Nics, nics)
	setQuotaListGetPublicIpsAttributeType(&this.PublicIps, publicIps)
	setQuotaListGetRamAttributeType(&this.Ram, ram)
	setQuotaListGetSecurityGroupRulesAttributeType(&this.SecurityGroupRules, securityGroupRules)
	setQuotaListGetSecurityGroupsAttributeType(&this.SecurityGroups, securityGroups)
	setQuotaListGetSnapshotsAttributeType(&this.Snapshots, snapshots)
	setQuotaListGetVcpuAttributeType(&this.Vcpu, vcpu)
	setQuotaListGetVolumesAttributeType(&this.Volumes, volumes)
	return &this
}

// NewQuotaListWithDefaults instantiates a new QuotaList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListWithDefaults() *QuotaList {
	this := QuotaList{}
	return &this
}

// GetBackupGigabytes returns the BackupGigabytes field value
func (o *QuotaList) GetBackupGigabytes() (ret QuotaListGetBackupGigabytesRetType) {
	ret, _ = o.GetBackupGigabytesOk()
	return ret
}

// GetBackupGigabytesOk returns a tuple with the BackupGigabytes field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetBackupGigabytesOk() (ret QuotaListGetBackupGigabytesRetType, ok bool) {
	return getQuotaListGetBackupGigabytesAttributeTypeOk(o.BackupGigabytes)
}

// SetBackupGigabytes sets field value
func (o *QuotaList) SetBackupGigabytes(v QuotaListGetBackupGigabytesRetType) {
	setQuotaListGetBackupGigabytesAttributeType(&o.BackupGigabytes, v)
}

// GetBackups returns the Backups field value
func (o *QuotaList) GetBackups() (ret QuotaListGetBackupsRetType) {
	ret, _ = o.GetBackupsOk()
	return ret
}

// GetBackupsOk returns a tuple with the Backups field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetBackupsOk() (ret QuotaListGetBackupsRetType, ok bool) {
	return getQuotaListGetBackupsAttributeTypeOk(o.Backups)
}

// SetBackups sets field value
func (o *QuotaList) SetBackups(v QuotaListGetBackupsRetType) {
	setQuotaListGetBackupsAttributeType(&o.Backups, v)
}

// GetGigabytes returns the Gigabytes field value
func (o *QuotaList) GetGigabytes() (ret QuotaListGetGigabytesRetType) {
	ret, _ = o.GetGigabytesOk()
	return ret
}

// GetGigabytesOk returns a tuple with the Gigabytes field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetGigabytesOk() (ret QuotaListGetGigabytesRetType, ok bool) {
	return getQuotaListGetGigabytesAttributeTypeOk(o.Gigabytes)
}

// SetGigabytes sets field value
func (o *QuotaList) SetGigabytes(v QuotaListGetGigabytesRetType) {
	setQuotaListGetGigabytesAttributeType(&o.Gigabytes, v)
}

// GetNetworks returns the Networks field value
func (o *QuotaList) GetNetworks() (ret QuotaListGetNetworksRetType) {
	ret, _ = o.GetNetworksOk()
	return ret
}

// GetNetworksOk returns a tuple with the Networks field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetNetworksOk() (ret QuotaListGetNetworksRetType, ok bool) {
	return getQuotaListGetNetworksAttributeTypeOk(o.Networks)
}

// SetNetworks sets field value
func (o *QuotaList) SetNetworks(v QuotaListGetNetworksRetType) {
	setQuotaListGetNetworksAttributeType(&o.Networks, v)
}

// GetNics returns the Nics field value
func (o *QuotaList) GetNics() (ret QuotaListGetNicsRetType) {
	ret, _ = o.GetNicsOk()
	return ret
}

// GetNicsOk returns a tuple with the Nics field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetNicsOk() (ret QuotaListGetNicsRetType, ok bool) {
	return getQuotaListGetNicsAttributeTypeOk(o.Nics)
}

// SetNics sets field value
func (o *QuotaList) SetNics(v QuotaListGetNicsRetType) {
	setQuotaListGetNicsAttributeType(&o.Nics, v)
}

// GetPublicIps returns the PublicIps field value
func (o *QuotaList) GetPublicIps() (ret QuotaListGetPublicIpsRetType) {
	ret, _ = o.GetPublicIpsOk()
	return ret
}

// GetPublicIpsOk returns a tuple with the PublicIps field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetPublicIpsOk() (ret QuotaListGetPublicIpsRetType, ok bool) {
	return getQuotaListGetPublicIpsAttributeTypeOk(o.PublicIps)
}

// SetPublicIps sets field value
func (o *QuotaList) SetPublicIps(v QuotaListGetPublicIpsRetType) {
	setQuotaListGetPublicIpsAttributeType(&o.PublicIps, v)
}

// GetRam returns the Ram field value
func (o *QuotaList) GetRam() (ret QuotaListGetRamRetType) {
	ret, _ = o.GetRamOk()
	return ret
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetRamOk() (ret QuotaListGetRamRetType, ok bool) {
	return getQuotaListGetRamAttributeTypeOk(o.Ram)
}

// SetRam sets field value
func (o *QuotaList) SetRam(v QuotaListGetRamRetType) {
	setQuotaListGetRamAttributeType(&o.Ram, v)
}

// GetSecurityGroupRules returns the SecurityGroupRules field value
func (o *QuotaList) GetSecurityGroupRules() (ret QuotaListGetSecurityGroupRulesRetType) {
	ret, _ = o.GetSecurityGroupRulesOk()
	return ret
}

// GetSecurityGroupRulesOk returns a tuple with the SecurityGroupRules field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetSecurityGroupRulesOk() (ret QuotaListGetSecurityGroupRulesRetType, ok bool) {
	return getQuotaListGetSecurityGroupRulesAttributeTypeOk(o.SecurityGroupRules)
}

// SetSecurityGroupRules sets field value
func (o *QuotaList) SetSecurityGroupRules(v QuotaListGetSecurityGroupRulesRetType) {
	setQuotaListGetSecurityGroupRulesAttributeType(&o.SecurityGroupRules, v)
}

// GetSecurityGroups returns the SecurityGroups field value
func (o *QuotaList) GetSecurityGroups() (ret QuotaListGetSecurityGroupsRetType) {
	ret, _ = o.GetSecurityGroupsOk()
	return ret
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetSecurityGroupsOk() (ret QuotaListGetSecurityGroupsRetType, ok bool) {
	return getQuotaListGetSecurityGroupsAttributeTypeOk(o.SecurityGroups)
}

// SetSecurityGroups sets field value
func (o *QuotaList) SetSecurityGroups(v QuotaListGetSecurityGroupsRetType) {
	setQuotaListGetSecurityGroupsAttributeType(&o.SecurityGroups, v)
}

// GetSnapshots returns the Snapshots field value
func (o *QuotaList) GetSnapshots() (ret QuotaListGetSnapshotsRetType) {
	ret, _ = o.GetSnapshotsOk()
	return ret
}

// GetSnapshotsOk returns a tuple with the Snapshots field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetSnapshotsOk() (ret QuotaListGetSnapshotsRetType, ok bool) {
	return getQuotaListGetSnapshotsAttributeTypeOk(o.Snapshots)
}

// SetSnapshots sets field value
func (o *QuotaList) SetSnapshots(v QuotaListGetSnapshotsRetType) {
	setQuotaListGetSnapshotsAttributeType(&o.Snapshots, v)
}

// GetVcpu returns the Vcpu field value
func (o *QuotaList) GetVcpu() (ret QuotaListGetVcpuRetType) {
	ret, _ = o.GetVcpuOk()
	return ret
}

// GetVcpuOk returns a tuple with the Vcpu field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetVcpuOk() (ret QuotaListGetVcpuRetType, ok bool) {
	return getQuotaListGetVcpuAttributeTypeOk(o.Vcpu)
}

// SetVcpu sets field value
func (o *QuotaList) SetVcpu(v QuotaListGetVcpuRetType) {
	setQuotaListGetVcpuAttributeType(&o.Vcpu, v)
}

// GetVolumes returns the Volumes field value
func (o *QuotaList) GetVolumes() (ret QuotaListGetVolumesRetType) {
	ret, _ = o.GetVolumesOk()
	return ret
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *QuotaList) GetVolumesOk() (ret QuotaListGetVolumesRetType, ok bool) {
	return getQuotaListGetVolumesAttributeTypeOk(o.Volumes)
}

// SetVolumes sets field value
func (o *QuotaList) SetVolumes(v QuotaListGetVolumesRetType) {
	setQuotaListGetVolumesAttributeType(&o.Volumes, v)
}

func (o QuotaList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getQuotaListGetBackupGigabytesAttributeTypeOk(o.BackupGigabytes); ok {
		toSerialize["BackupGigabytes"] = val
	}
	if val, ok := getQuotaListGetBackupsAttributeTypeOk(o.Backups); ok {
		toSerialize["Backups"] = val
	}
	if val, ok := getQuotaListGetGigabytesAttributeTypeOk(o.Gigabytes); ok {
		toSerialize["Gigabytes"] = val
	}
	if val, ok := getQuotaListGetNetworksAttributeTypeOk(o.Networks); ok {
		toSerialize["Networks"] = val
	}
	if val, ok := getQuotaListGetNicsAttributeTypeOk(o.Nics); ok {
		toSerialize["Nics"] = val
	}
	if val, ok := getQuotaListGetPublicIpsAttributeTypeOk(o.PublicIps); ok {
		toSerialize["PublicIps"] = val
	}
	if val, ok := getQuotaListGetRamAttributeTypeOk(o.Ram); ok {
		toSerialize["Ram"] = val
	}
	if val, ok := getQuotaListGetSecurityGroupRulesAttributeTypeOk(o.SecurityGroupRules); ok {
		toSerialize["SecurityGroupRules"] = val
	}
	if val, ok := getQuotaListGetSecurityGroupsAttributeTypeOk(o.SecurityGroups); ok {
		toSerialize["SecurityGroups"] = val
	}
	if val, ok := getQuotaListGetSnapshotsAttributeTypeOk(o.Snapshots); ok {
		toSerialize["Snapshots"] = val
	}
	if val, ok := getQuotaListGetVcpuAttributeTypeOk(o.Vcpu); ok {
		toSerialize["Vcpu"] = val
	}
	if val, ok := getQuotaListGetVolumesAttributeTypeOk(o.Volumes); ok {
		toSerialize["Volumes"] = val
	}
	return toSerialize, nil
}

type NullableQuotaList struct {
	value *QuotaList
	isSet bool
}

func (v NullableQuotaList) Get() *QuotaList {
	return v.value
}

func (v *NullableQuotaList) Set(val *QuotaList) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaList) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaList(val *QuotaList) *NullableQuotaList {
	return &NullableQuotaList{value: val, isSet: true}
}

func (v NullableQuotaList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
