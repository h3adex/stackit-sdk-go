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

// checks if the PlanModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanModel{}

/*
	types and functions for alertMatchers
*/

// isInteger
type PlanModelGetAlertMatchersAttributeType = *int64
type PlanModelGetAlertMatchersArgType = int64
type PlanModelGetAlertMatchersRetType = int64

func getPlanModelGetAlertMatchersAttributeTypeOk(arg PlanModelGetAlertMatchersAttributeType) (ret PlanModelGetAlertMatchersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetAlertMatchersAttributeType(arg *PlanModelGetAlertMatchersAttributeType, val PlanModelGetAlertMatchersRetType) {
	*arg = &val
}

/*
	types and functions for alertReceivers
*/

// isInteger
type PlanModelGetAlertReceiversAttributeType = *int64
type PlanModelGetAlertReceiversArgType = int64
type PlanModelGetAlertReceiversRetType = int64

func getPlanModelGetAlertReceiversAttributeTypeOk(arg PlanModelGetAlertReceiversAttributeType) (ret PlanModelGetAlertReceiversRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetAlertReceiversAttributeType(arg *PlanModelGetAlertReceiversAttributeType, val PlanModelGetAlertReceiversRetType) {
	*arg = &val
}

/*
	types and functions for alertRules
*/

// isInteger
type PlanModelGetAlertRulesAttributeType = *int64
type PlanModelGetAlertRulesArgType = int64
type PlanModelGetAlertRulesRetType = int64

func getPlanModelGetAlertRulesAttributeTypeOk(arg PlanModelGetAlertRulesAttributeType) (ret PlanModelGetAlertRulesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetAlertRulesAttributeType(arg *PlanModelGetAlertRulesAttributeType, val PlanModelGetAlertRulesRetType) {
	*arg = &val
}

/*
	types and functions for amount
*/

// isNumber
type PlanModelGetAmountAttributeType = *float64
type PlanModelGetAmountArgType = float64
type PlanModelGetAmountRetType = float64

func getPlanModelGetAmountAttributeTypeOk(arg PlanModelGetAmountAttributeType) (ret PlanModelGetAmountRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetAmountAttributeType(arg *PlanModelGetAmountAttributeType, val PlanModelGetAmountRetType) {
	*arg = &val
}

/*
	types and functions for bucketSize
*/

// isInteger
type PlanModelGetBucketSizeAttributeType = *int64
type PlanModelGetBucketSizeArgType = int64
type PlanModelGetBucketSizeRetType = int64

func getPlanModelGetBucketSizeAttributeTypeOk(arg PlanModelGetBucketSizeAttributeType) (ret PlanModelGetBucketSizeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetBucketSizeAttributeType(arg *PlanModelGetBucketSizeAttributeType, val PlanModelGetBucketSizeRetType) {
	*arg = &val
}

/*
	types and functions for description
*/

// isNotNullableString
type PlanModelGetDescriptionAttributeType = *string

func getPlanModelGetDescriptionAttributeTypeOk(arg PlanModelGetDescriptionAttributeType) (ret PlanModelGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetDescriptionAttributeType(arg *PlanModelGetDescriptionAttributeType, val PlanModelGetDescriptionRetType) {
	*arg = &val
}

type PlanModelGetDescriptionArgType = string
type PlanModelGetDescriptionRetType = string

/*
	types and functions for grafanaGlobalDashboards
*/

// isInteger
type PlanModelGetGrafanaGlobalDashboardsAttributeType = *int64
type PlanModelGetGrafanaGlobalDashboardsArgType = int64
type PlanModelGetGrafanaGlobalDashboardsRetType = int64

func getPlanModelGetGrafanaGlobalDashboardsAttributeTypeOk(arg PlanModelGetGrafanaGlobalDashboardsAttributeType) (ret PlanModelGetGrafanaGlobalDashboardsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetGrafanaGlobalDashboardsAttributeType(arg *PlanModelGetGrafanaGlobalDashboardsAttributeType, val PlanModelGetGrafanaGlobalDashboardsRetType) {
	*arg = &val
}

/*
	types and functions for grafanaGlobalOrgs
*/

// isInteger
type PlanModelGetGrafanaGlobalOrgsAttributeType = *int64
type PlanModelGetGrafanaGlobalOrgsArgType = int64
type PlanModelGetGrafanaGlobalOrgsRetType = int64

func getPlanModelGetGrafanaGlobalOrgsAttributeTypeOk(arg PlanModelGetGrafanaGlobalOrgsAttributeType) (ret PlanModelGetGrafanaGlobalOrgsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetGrafanaGlobalOrgsAttributeType(arg *PlanModelGetGrafanaGlobalOrgsAttributeType, val PlanModelGetGrafanaGlobalOrgsRetType) {
	*arg = &val
}

/*
	types and functions for grafanaGlobalSessions
*/

// isInteger
type PlanModelGetGrafanaGlobalSessionsAttributeType = *int64
type PlanModelGetGrafanaGlobalSessionsArgType = int64
type PlanModelGetGrafanaGlobalSessionsRetType = int64

func getPlanModelGetGrafanaGlobalSessionsAttributeTypeOk(arg PlanModelGetGrafanaGlobalSessionsAttributeType) (ret PlanModelGetGrafanaGlobalSessionsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetGrafanaGlobalSessionsAttributeType(arg *PlanModelGetGrafanaGlobalSessionsAttributeType, val PlanModelGetGrafanaGlobalSessionsRetType) {
	*arg = &val
}

/*
	types and functions for grafanaGlobalUsers
*/

// isInteger
type PlanModelGetGrafanaGlobalUsersAttributeType = *int64
type PlanModelGetGrafanaGlobalUsersArgType = int64
type PlanModelGetGrafanaGlobalUsersRetType = int64

func getPlanModelGetGrafanaGlobalUsersAttributeTypeOk(arg PlanModelGetGrafanaGlobalUsersAttributeType) (ret PlanModelGetGrafanaGlobalUsersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetGrafanaGlobalUsersAttributeType(arg *PlanModelGetGrafanaGlobalUsersAttributeType, val PlanModelGetGrafanaGlobalUsersRetType) {
	*arg = &val
}

/*
	types and functions for id
*/

// isNotNullableString
type PlanModelGetIdAttributeType = *string

func getPlanModelGetIdAttributeTypeOk(arg PlanModelGetIdAttributeType) (ret PlanModelGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetIdAttributeType(arg *PlanModelGetIdAttributeType, val PlanModelGetIdRetType) {
	*arg = &val
}

type PlanModelGetIdArgType = string
type PlanModelGetIdRetType = string

/*
	types and functions for logsAlert
*/

// isInteger
type PlanModelGetLogsAlertAttributeType = *int64
type PlanModelGetLogsAlertArgType = int64
type PlanModelGetLogsAlertRetType = int64

func getPlanModelGetLogsAlertAttributeTypeOk(arg PlanModelGetLogsAlertAttributeType) (ret PlanModelGetLogsAlertRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetLogsAlertAttributeType(arg *PlanModelGetLogsAlertAttributeType, val PlanModelGetLogsAlertRetType) {
	*arg = &val
}

/*
	types and functions for logsStorage
*/

// isInteger
type PlanModelGetLogsStorageAttributeType = *int64
type PlanModelGetLogsStorageArgType = int64
type PlanModelGetLogsStorageRetType = int64

func getPlanModelGetLogsStorageAttributeTypeOk(arg PlanModelGetLogsStorageAttributeType) (ret PlanModelGetLogsStorageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetLogsStorageAttributeType(arg *PlanModelGetLogsStorageAttributeType, val PlanModelGetLogsStorageRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type PlanModelGetNameAttributeType = *string

func getPlanModelGetNameAttributeTypeOk(arg PlanModelGetNameAttributeType) (ret PlanModelGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetNameAttributeType(arg *PlanModelGetNameAttributeType, val PlanModelGetNameRetType) {
	*arg = &val
}

type PlanModelGetNameArgType = string
type PlanModelGetNameRetType = string

/*
	types and functions for planId
*/

// isNotNullableString
type PlanModelGetPlanIdAttributeType = *string

func getPlanModelGetPlanIdAttributeTypeOk(arg PlanModelGetPlanIdAttributeType) (ret PlanModelGetPlanIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetPlanIdAttributeType(arg *PlanModelGetPlanIdAttributeType, val PlanModelGetPlanIdRetType) {
	*arg = &val
}

type PlanModelGetPlanIdArgType = string
type PlanModelGetPlanIdRetType = string

/*
	types and functions for samplesPerScrape
*/

// isInteger
type PlanModelGetSamplesPerScrapeAttributeType = *int64
type PlanModelGetSamplesPerScrapeArgType = int64
type PlanModelGetSamplesPerScrapeRetType = int64

func getPlanModelGetSamplesPerScrapeAttributeTypeOk(arg PlanModelGetSamplesPerScrapeAttributeType) (ret PlanModelGetSamplesPerScrapeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetSamplesPerScrapeAttributeType(arg *PlanModelGetSamplesPerScrapeAttributeType, val PlanModelGetSamplesPerScrapeRetType) {
	*arg = &val
}

/*
	types and functions for targetNumber
*/

// isInteger
type PlanModelGetTargetNumberAttributeType = *int64
type PlanModelGetTargetNumberArgType = int64
type PlanModelGetTargetNumberRetType = int64

func getPlanModelGetTargetNumberAttributeTypeOk(arg PlanModelGetTargetNumberAttributeType) (ret PlanModelGetTargetNumberRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetTargetNumberAttributeType(arg *PlanModelGetTargetNumberAttributeType, val PlanModelGetTargetNumberRetType) {
	*arg = &val
}

/*
	types and functions for totalMetricSamples
*/

// isInteger
type PlanModelGetTotalMetricSamplesAttributeType = *int64
type PlanModelGetTotalMetricSamplesArgType = int64
type PlanModelGetTotalMetricSamplesRetType = int64

func getPlanModelGetTotalMetricSamplesAttributeTypeOk(arg PlanModelGetTotalMetricSamplesAttributeType) (ret PlanModelGetTotalMetricSamplesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetTotalMetricSamplesAttributeType(arg *PlanModelGetTotalMetricSamplesAttributeType, val PlanModelGetTotalMetricSamplesRetType) {
	*arg = &val
}

/*
	types and functions for tracesStorage
*/

// isInteger
type PlanModelGetTracesStorageAttributeType = *int64
type PlanModelGetTracesStorageArgType = int64
type PlanModelGetTracesStorageRetType = int64

func getPlanModelGetTracesStorageAttributeTypeOk(arg PlanModelGetTracesStorageAttributeType) (ret PlanModelGetTracesStorageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlanModelGetTracesStorageAttributeType(arg *PlanModelGetTracesStorageAttributeType, val PlanModelGetTracesStorageRetType) {
	*arg = &val
}

// PlanModel struct for PlanModel
type PlanModel struct {
	// REQUIRED
	AlertMatchers PlanModelGetAlertMatchersAttributeType `json:"alertMatchers" required:"true"`
	// REQUIRED
	AlertReceivers PlanModelGetAlertReceiversAttributeType `json:"alertReceivers" required:"true"`
	// REQUIRED
	AlertRules PlanModelGetAlertRulesAttributeType `json:"alertRules" required:"true"`
	Amount     PlanModelGetAmountAttributeType     `json:"amount,omitempty"`
	// REQUIRED
	BucketSize  PlanModelGetBucketSizeAttributeType  `json:"bucketSize" required:"true"`
	Description PlanModelGetDescriptionAttributeType `json:"description,omitempty"`
	// REQUIRED
	GrafanaGlobalDashboards PlanModelGetGrafanaGlobalDashboardsAttributeType `json:"grafanaGlobalDashboards" required:"true"`
	// REQUIRED
	GrafanaGlobalOrgs PlanModelGetGrafanaGlobalOrgsAttributeType `json:"grafanaGlobalOrgs" required:"true"`
	// REQUIRED
	GrafanaGlobalSessions PlanModelGetGrafanaGlobalSessionsAttributeType `json:"grafanaGlobalSessions" required:"true"`
	// REQUIRED
	GrafanaGlobalUsers PlanModelGetGrafanaGlobalUsersAttributeType `json:"grafanaGlobalUsers" required:"true"`
	// REQUIRED
	Id PlanModelGetIdAttributeType `json:"id" required:"true"`
	// REQUIRED
	LogsAlert PlanModelGetLogsAlertAttributeType `json:"logsAlert" required:"true"`
	// REQUIRED
	LogsStorage PlanModelGetLogsStorageAttributeType `json:"logsStorage" required:"true"`
	Name        PlanModelGetNameAttributeType        `json:"name,omitempty"`
	// REQUIRED
	PlanId PlanModelGetPlanIdAttributeType `json:"planId" required:"true"`
	// REQUIRED
	SamplesPerScrape PlanModelGetSamplesPerScrapeAttributeType `json:"samplesPerScrape" required:"true"`
	// REQUIRED
	TargetNumber PlanModelGetTargetNumberAttributeType `json:"targetNumber" required:"true"`
	// REQUIRED
	TotalMetricSamples PlanModelGetTotalMetricSamplesAttributeType `json:"totalMetricSamples" required:"true"`
	// REQUIRED
	TracesStorage PlanModelGetTracesStorageAttributeType `json:"tracesStorage" required:"true"`
}

type _PlanModel PlanModel

// NewPlanModel instantiates a new PlanModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanModel(alertMatchers PlanModelGetAlertMatchersArgType, alertReceivers PlanModelGetAlertReceiversArgType, alertRules PlanModelGetAlertRulesArgType, bucketSize PlanModelGetBucketSizeArgType, grafanaGlobalDashboards PlanModelGetGrafanaGlobalDashboardsArgType, grafanaGlobalOrgs PlanModelGetGrafanaGlobalOrgsArgType, grafanaGlobalSessions PlanModelGetGrafanaGlobalSessionsArgType, grafanaGlobalUsers PlanModelGetGrafanaGlobalUsersArgType, id PlanModelGetIdArgType, logsAlert PlanModelGetLogsAlertArgType, logsStorage PlanModelGetLogsStorageArgType, planId PlanModelGetPlanIdArgType, samplesPerScrape PlanModelGetSamplesPerScrapeArgType, targetNumber PlanModelGetTargetNumberArgType, totalMetricSamples PlanModelGetTotalMetricSamplesArgType, tracesStorage PlanModelGetTracesStorageArgType) *PlanModel {
	this := PlanModel{}
	setPlanModelGetAlertMatchersAttributeType(&this.AlertMatchers, alertMatchers)
	setPlanModelGetAlertReceiversAttributeType(&this.AlertReceivers, alertReceivers)
	setPlanModelGetAlertRulesAttributeType(&this.AlertRules, alertRules)
	setPlanModelGetBucketSizeAttributeType(&this.BucketSize, bucketSize)
	setPlanModelGetGrafanaGlobalDashboardsAttributeType(&this.GrafanaGlobalDashboards, grafanaGlobalDashboards)
	setPlanModelGetGrafanaGlobalOrgsAttributeType(&this.GrafanaGlobalOrgs, grafanaGlobalOrgs)
	setPlanModelGetGrafanaGlobalSessionsAttributeType(&this.GrafanaGlobalSessions, grafanaGlobalSessions)
	setPlanModelGetGrafanaGlobalUsersAttributeType(&this.GrafanaGlobalUsers, grafanaGlobalUsers)
	setPlanModelGetIdAttributeType(&this.Id, id)
	setPlanModelGetLogsAlertAttributeType(&this.LogsAlert, logsAlert)
	setPlanModelGetLogsStorageAttributeType(&this.LogsStorage, logsStorage)
	setPlanModelGetPlanIdAttributeType(&this.PlanId, planId)
	setPlanModelGetSamplesPerScrapeAttributeType(&this.SamplesPerScrape, samplesPerScrape)
	setPlanModelGetTargetNumberAttributeType(&this.TargetNumber, targetNumber)
	setPlanModelGetTotalMetricSamplesAttributeType(&this.TotalMetricSamples, totalMetricSamples)
	setPlanModelGetTracesStorageAttributeType(&this.TracesStorage, tracesStorage)
	return &this
}

// NewPlanModelWithDefaults instantiates a new PlanModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanModelWithDefaults() *PlanModel {
	this := PlanModel{}
	return &this
}

// GetAlertMatchers returns the AlertMatchers field value
func (o *PlanModel) GetAlertMatchers() (ret PlanModelGetAlertMatchersRetType) {
	ret, _ = o.GetAlertMatchersOk()
	return ret
}

// GetAlertMatchersOk returns a tuple with the AlertMatchers field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetAlertMatchersOk() (ret PlanModelGetAlertMatchersRetType, ok bool) {
	return getPlanModelGetAlertMatchersAttributeTypeOk(o.AlertMatchers)
}

// SetAlertMatchers sets field value
func (o *PlanModel) SetAlertMatchers(v PlanModelGetAlertMatchersRetType) {
	setPlanModelGetAlertMatchersAttributeType(&o.AlertMatchers, v)
}

// GetAlertReceivers returns the AlertReceivers field value
func (o *PlanModel) GetAlertReceivers() (ret PlanModelGetAlertReceiversRetType) {
	ret, _ = o.GetAlertReceiversOk()
	return ret
}

// GetAlertReceiversOk returns a tuple with the AlertReceivers field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetAlertReceiversOk() (ret PlanModelGetAlertReceiversRetType, ok bool) {
	return getPlanModelGetAlertReceiversAttributeTypeOk(o.AlertReceivers)
}

// SetAlertReceivers sets field value
func (o *PlanModel) SetAlertReceivers(v PlanModelGetAlertReceiversRetType) {
	setPlanModelGetAlertReceiversAttributeType(&o.AlertReceivers, v)
}

// GetAlertRules returns the AlertRules field value
func (o *PlanModel) GetAlertRules() (ret PlanModelGetAlertRulesRetType) {
	ret, _ = o.GetAlertRulesOk()
	return ret
}

// GetAlertRulesOk returns a tuple with the AlertRules field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetAlertRulesOk() (ret PlanModelGetAlertRulesRetType, ok bool) {
	return getPlanModelGetAlertRulesAttributeTypeOk(o.AlertRules)
}

// SetAlertRules sets field value
func (o *PlanModel) SetAlertRules(v PlanModelGetAlertRulesRetType) {
	setPlanModelGetAlertRulesAttributeType(&o.AlertRules, v)
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PlanModel) GetAmount() (res PlanModelGetAmountRetType) {
	res, _ = o.GetAmountOk()
	return
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanModel) GetAmountOk() (ret PlanModelGetAmountRetType, ok bool) {
	return getPlanModelGetAmountAttributeTypeOk(o.Amount)
}

// HasAmount returns a boolean if a field has been set.
func (o *PlanModel) HasAmount() bool {
	_, ok := o.GetAmountOk()
	return ok
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *PlanModel) SetAmount(v PlanModelGetAmountRetType) {
	setPlanModelGetAmountAttributeType(&o.Amount, v)
}

// GetBucketSize returns the BucketSize field value
func (o *PlanModel) GetBucketSize() (ret PlanModelGetBucketSizeRetType) {
	ret, _ = o.GetBucketSizeOk()
	return ret
}

// GetBucketSizeOk returns a tuple with the BucketSize field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetBucketSizeOk() (ret PlanModelGetBucketSizeRetType, ok bool) {
	return getPlanModelGetBucketSizeAttributeTypeOk(o.BucketSize)
}

// SetBucketSize sets field value
func (o *PlanModel) SetBucketSize(v PlanModelGetBucketSizeRetType) {
	setPlanModelGetBucketSizeAttributeType(&o.BucketSize, v)
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlanModel) GetDescription() (res PlanModelGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanModel) GetDescriptionOk() (ret PlanModelGetDescriptionRetType, ok bool) {
	return getPlanModelGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *PlanModel) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlanModel) SetDescription(v PlanModelGetDescriptionRetType) {
	setPlanModelGetDescriptionAttributeType(&o.Description, v)
}

// GetGrafanaGlobalDashboards returns the GrafanaGlobalDashboards field value
func (o *PlanModel) GetGrafanaGlobalDashboards() (ret PlanModelGetGrafanaGlobalDashboardsRetType) {
	ret, _ = o.GetGrafanaGlobalDashboardsOk()
	return ret
}

// GetGrafanaGlobalDashboardsOk returns a tuple with the GrafanaGlobalDashboards field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetGrafanaGlobalDashboardsOk() (ret PlanModelGetGrafanaGlobalDashboardsRetType, ok bool) {
	return getPlanModelGetGrafanaGlobalDashboardsAttributeTypeOk(o.GrafanaGlobalDashboards)
}

// SetGrafanaGlobalDashboards sets field value
func (o *PlanModel) SetGrafanaGlobalDashboards(v PlanModelGetGrafanaGlobalDashboardsRetType) {
	setPlanModelGetGrafanaGlobalDashboardsAttributeType(&o.GrafanaGlobalDashboards, v)
}

// GetGrafanaGlobalOrgs returns the GrafanaGlobalOrgs field value
func (o *PlanModel) GetGrafanaGlobalOrgs() (ret PlanModelGetGrafanaGlobalOrgsRetType) {
	ret, _ = o.GetGrafanaGlobalOrgsOk()
	return ret
}

// GetGrafanaGlobalOrgsOk returns a tuple with the GrafanaGlobalOrgs field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetGrafanaGlobalOrgsOk() (ret PlanModelGetGrafanaGlobalOrgsRetType, ok bool) {
	return getPlanModelGetGrafanaGlobalOrgsAttributeTypeOk(o.GrafanaGlobalOrgs)
}

// SetGrafanaGlobalOrgs sets field value
func (o *PlanModel) SetGrafanaGlobalOrgs(v PlanModelGetGrafanaGlobalOrgsRetType) {
	setPlanModelGetGrafanaGlobalOrgsAttributeType(&o.GrafanaGlobalOrgs, v)
}

// GetGrafanaGlobalSessions returns the GrafanaGlobalSessions field value
func (o *PlanModel) GetGrafanaGlobalSessions() (ret PlanModelGetGrafanaGlobalSessionsRetType) {
	ret, _ = o.GetGrafanaGlobalSessionsOk()
	return ret
}

// GetGrafanaGlobalSessionsOk returns a tuple with the GrafanaGlobalSessions field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetGrafanaGlobalSessionsOk() (ret PlanModelGetGrafanaGlobalSessionsRetType, ok bool) {
	return getPlanModelGetGrafanaGlobalSessionsAttributeTypeOk(o.GrafanaGlobalSessions)
}

// SetGrafanaGlobalSessions sets field value
func (o *PlanModel) SetGrafanaGlobalSessions(v PlanModelGetGrafanaGlobalSessionsRetType) {
	setPlanModelGetGrafanaGlobalSessionsAttributeType(&o.GrafanaGlobalSessions, v)
}

// GetGrafanaGlobalUsers returns the GrafanaGlobalUsers field value
func (o *PlanModel) GetGrafanaGlobalUsers() (ret PlanModelGetGrafanaGlobalUsersRetType) {
	ret, _ = o.GetGrafanaGlobalUsersOk()
	return ret
}

// GetGrafanaGlobalUsersOk returns a tuple with the GrafanaGlobalUsers field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetGrafanaGlobalUsersOk() (ret PlanModelGetGrafanaGlobalUsersRetType, ok bool) {
	return getPlanModelGetGrafanaGlobalUsersAttributeTypeOk(o.GrafanaGlobalUsers)
}

// SetGrafanaGlobalUsers sets field value
func (o *PlanModel) SetGrafanaGlobalUsers(v PlanModelGetGrafanaGlobalUsersRetType) {
	setPlanModelGetGrafanaGlobalUsersAttributeType(&o.GrafanaGlobalUsers, v)
}

// GetId returns the Id field value
func (o *PlanModel) GetId() (ret PlanModelGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetIdOk() (ret PlanModelGetIdRetType, ok bool) {
	return getPlanModelGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *PlanModel) SetId(v PlanModelGetIdRetType) {
	setPlanModelGetIdAttributeType(&o.Id, v)
}

// GetLogsAlert returns the LogsAlert field value
func (o *PlanModel) GetLogsAlert() (ret PlanModelGetLogsAlertRetType) {
	ret, _ = o.GetLogsAlertOk()
	return ret
}

// GetLogsAlertOk returns a tuple with the LogsAlert field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetLogsAlertOk() (ret PlanModelGetLogsAlertRetType, ok bool) {
	return getPlanModelGetLogsAlertAttributeTypeOk(o.LogsAlert)
}

// SetLogsAlert sets field value
func (o *PlanModel) SetLogsAlert(v PlanModelGetLogsAlertRetType) {
	setPlanModelGetLogsAlertAttributeType(&o.LogsAlert, v)
}

// GetLogsStorage returns the LogsStorage field value
func (o *PlanModel) GetLogsStorage() (ret PlanModelGetLogsStorageRetType) {
	ret, _ = o.GetLogsStorageOk()
	return ret
}

// GetLogsStorageOk returns a tuple with the LogsStorage field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetLogsStorageOk() (ret PlanModelGetLogsStorageRetType, ok bool) {
	return getPlanModelGetLogsStorageAttributeTypeOk(o.LogsStorage)
}

// SetLogsStorage sets field value
func (o *PlanModel) SetLogsStorage(v PlanModelGetLogsStorageRetType) {
	setPlanModelGetLogsStorageAttributeType(&o.LogsStorage, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlanModel) GetName() (res PlanModelGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanModel) GetNameOk() (ret PlanModelGetNameRetType, ok bool) {
	return getPlanModelGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *PlanModel) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlanModel) SetName(v PlanModelGetNameRetType) {
	setPlanModelGetNameAttributeType(&o.Name, v)
}

// GetPlanId returns the PlanId field value
func (o *PlanModel) GetPlanId() (ret PlanModelGetPlanIdRetType) {
	ret, _ = o.GetPlanIdOk()
	return ret
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetPlanIdOk() (ret PlanModelGetPlanIdRetType, ok bool) {
	return getPlanModelGetPlanIdAttributeTypeOk(o.PlanId)
}

// SetPlanId sets field value
func (o *PlanModel) SetPlanId(v PlanModelGetPlanIdRetType) {
	setPlanModelGetPlanIdAttributeType(&o.PlanId, v)
}

// GetSamplesPerScrape returns the SamplesPerScrape field value
func (o *PlanModel) GetSamplesPerScrape() (ret PlanModelGetSamplesPerScrapeRetType) {
	ret, _ = o.GetSamplesPerScrapeOk()
	return ret
}

// GetSamplesPerScrapeOk returns a tuple with the SamplesPerScrape field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetSamplesPerScrapeOk() (ret PlanModelGetSamplesPerScrapeRetType, ok bool) {
	return getPlanModelGetSamplesPerScrapeAttributeTypeOk(o.SamplesPerScrape)
}

// SetSamplesPerScrape sets field value
func (o *PlanModel) SetSamplesPerScrape(v PlanModelGetSamplesPerScrapeRetType) {
	setPlanModelGetSamplesPerScrapeAttributeType(&o.SamplesPerScrape, v)
}

// GetTargetNumber returns the TargetNumber field value
func (o *PlanModel) GetTargetNumber() (ret PlanModelGetTargetNumberRetType) {
	ret, _ = o.GetTargetNumberOk()
	return ret
}

// GetTargetNumberOk returns a tuple with the TargetNumber field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetTargetNumberOk() (ret PlanModelGetTargetNumberRetType, ok bool) {
	return getPlanModelGetTargetNumberAttributeTypeOk(o.TargetNumber)
}

// SetTargetNumber sets field value
func (o *PlanModel) SetTargetNumber(v PlanModelGetTargetNumberRetType) {
	setPlanModelGetTargetNumberAttributeType(&o.TargetNumber, v)
}

// GetTotalMetricSamples returns the TotalMetricSamples field value
func (o *PlanModel) GetTotalMetricSamples() (ret PlanModelGetTotalMetricSamplesRetType) {
	ret, _ = o.GetTotalMetricSamplesOk()
	return ret
}

// GetTotalMetricSamplesOk returns a tuple with the TotalMetricSamples field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetTotalMetricSamplesOk() (ret PlanModelGetTotalMetricSamplesRetType, ok bool) {
	return getPlanModelGetTotalMetricSamplesAttributeTypeOk(o.TotalMetricSamples)
}

// SetTotalMetricSamples sets field value
func (o *PlanModel) SetTotalMetricSamples(v PlanModelGetTotalMetricSamplesRetType) {
	setPlanModelGetTotalMetricSamplesAttributeType(&o.TotalMetricSamples, v)
}

// GetTracesStorage returns the TracesStorage field value
func (o *PlanModel) GetTracesStorage() (ret PlanModelGetTracesStorageRetType) {
	ret, _ = o.GetTracesStorageOk()
	return ret
}

// GetTracesStorageOk returns a tuple with the TracesStorage field value
// and a boolean to check if the value has been set.
func (o *PlanModel) GetTracesStorageOk() (ret PlanModelGetTracesStorageRetType, ok bool) {
	return getPlanModelGetTracesStorageAttributeTypeOk(o.TracesStorage)
}

// SetTracesStorage sets field value
func (o *PlanModel) SetTracesStorage(v PlanModelGetTracesStorageRetType) {
	setPlanModelGetTracesStorageAttributeType(&o.TracesStorage, v)
}

func (o PlanModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPlanModelGetAlertMatchersAttributeTypeOk(o.AlertMatchers); ok {
		toSerialize["AlertMatchers"] = val
	}
	if val, ok := getPlanModelGetAlertReceiversAttributeTypeOk(o.AlertReceivers); ok {
		toSerialize["AlertReceivers"] = val
	}
	if val, ok := getPlanModelGetAlertRulesAttributeTypeOk(o.AlertRules); ok {
		toSerialize["AlertRules"] = val
	}
	if val, ok := getPlanModelGetAmountAttributeTypeOk(o.Amount); ok {
		toSerialize["Amount"] = val
	}
	if val, ok := getPlanModelGetBucketSizeAttributeTypeOk(o.BucketSize); ok {
		toSerialize["BucketSize"] = val
	}
	if val, ok := getPlanModelGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getPlanModelGetGrafanaGlobalDashboardsAttributeTypeOk(o.GrafanaGlobalDashboards); ok {
		toSerialize["GrafanaGlobalDashboards"] = val
	}
	if val, ok := getPlanModelGetGrafanaGlobalOrgsAttributeTypeOk(o.GrafanaGlobalOrgs); ok {
		toSerialize["GrafanaGlobalOrgs"] = val
	}
	if val, ok := getPlanModelGetGrafanaGlobalSessionsAttributeTypeOk(o.GrafanaGlobalSessions); ok {
		toSerialize["GrafanaGlobalSessions"] = val
	}
	if val, ok := getPlanModelGetGrafanaGlobalUsersAttributeTypeOk(o.GrafanaGlobalUsers); ok {
		toSerialize["GrafanaGlobalUsers"] = val
	}
	if val, ok := getPlanModelGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getPlanModelGetLogsAlertAttributeTypeOk(o.LogsAlert); ok {
		toSerialize["LogsAlert"] = val
	}
	if val, ok := getPlanModelGetLogsStorageAttributeTypeOk(o.LogsStorage); ok {
		toSerialize["LogsStorage"] = val
	}
	if val, ok := getPlanModelGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getPlanModelGetPlanIdAttributeTypeOk(o.PlanId); ok {
		toSerialize["PlanId"] = val
	}
	if val, ok := getPlanModelGetSamplesPerScrapeAttributeTypeOk(o.SamplesPerScrape); ok {
		toSerialize["SamplesPerScrape"] = val
	}
	if val, ok := getPlanModelGetTargetNumberAttributeTypeOk(o.TargetNumber); ok {
		toSerialize["TargetNumber"] = val
	}
	if val, ok := getPlanModelGetTotalMetricSamplesAttributeTypeOk(o.TotalMetricSamples); ok {
		toSerialize["TotalMetricSamples"] = val
	}
	if val, ok := getPlanModelGetTracesStorageAttributeTypeOk(o.TracesStorage); ok {
		toSerialize["TracesStorage"] = val
	}
	return toSerialize, nil
}

type NullablePlanModel struct {
	value *PlanModel
	isSet bool
}

func (v NullablePlanModel) Get() *PlanModel {
	return v.value
}

func (v *NullablePlanModel) Set(val *PlanModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanModel(val *PlanModel) *NullablePlanModel {
	return &NullablePlanModel{value: val, isSet: true}
}

func (v NullablePlanModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
