/*
STACKIT Opensearch API

The STACKIT Opensearch API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

import (
	"encoding/json"
	"fmt"
)

// checks if the InstanceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceParameters{}

/*
	types and functions for enable_monitoring
*/

// isBoolean
type InstanceParametersgetEnableMonitoringAttributeType = *bool
type InstanceParametersgetEnableMonitoringArgType = bool
type InstanceParametersgetEnableMonitoringRetType = bool

func getInstanceParametersgetEnableMonitoringAttributeTypeOk(arg InstanceParametersgetEnableMonitoringAttributeType) (ret InstanceParametersgetEnableMonitoringRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersgetEnableMonitoringAttributeType(arg *InstanceParametersgetEnableMonitoringAttributeType, val InstanceParametersgetEnableMonitoringRetType) {
	*arg = &val
}

/*
	types and functions for graphite
*/

// isNotNullableString
type InstanceParametersGetGraphiteAttributeType = *string

func getInstanceParametersGetGraphiteAttributeTypeOk(arg InstanceParametersGetGraphiteAttributeType) (ret InstanceParametersGetGraphiteRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetGraphiteAttributeType(arg *InstanceParametersGetGraphiteAttributeType, val InstanceParametersGetGraphiteRetType) {
	*arg = &val
}

type InstanceParametersGetGraphiteArgType = string
type InstanceParametersGetGraphiteRetType = string

/*
	types and functions for java_garbage_collector
*/

// isEnum

// InstanceParametersJavaGarbageCollector the model 'InstanceParameters'
// value type for enums
type InstanceParametersJavaGarbageCollector string

// List of JavaGarbageCollector
const (
	INSTANCEPARAMETERSJAVA_GARBAGE_COLLECTOR_USE_SERIAL_GC       InstanceParametersJavaGarbageCollector = "UseSerialGC"
	INSTANCEPARAMETERSJAVA_GARBAGE_COLLECTOR_USE_PARALLEL_GC     InstanceParametersJavaGarbageCollector = "UseParallelGC"
	INSTANCEPARAMETERSJAVA_GARBAGE_COLLECTOR_USE_PARALLEL_OLD_GC InstanceParametersJavaGarbageCollector = "UseParallelOldGC"
	INSTANCEPARAMETERSJAVA_GARBAGE_COLLECTOR_USE_G1_GC           InstanceParametersJavaGarbageCollector = "UseG1GC"
)

// All allowed values of InstanceParameters enum
var AllowedInstanceParametersJavaGarbageCollectorEnumValues = []InstanceParametersJavaGarbageCollector{
	"UseSerialGC",
	"UseParallelGC",
	"UseParallelOldGC",
	"UseG1GC",
}

func (v *InstanceParametersJavaGarbageCollector) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson InstanceParametersJavaGarbageCollector
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
	enumTypeValue := InstanceParametersJavaGarbageCollector(value)
	for _, existing := range AllowedInstanceParametersJavaGarbageCollectorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceParameters", value)
}

// NewInstanceParametersJavaGarbageCollectorFromValue returns a pointer to a valid InstanceParametersJavaGarbageCollector
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceParametersJavaGarbageCollectorFromValue(v InstanceParametersJavaGarbageCollector) (*InstanceParametersJavaGarbageCollector, error) {
	ev := InstanceParametersJavaGarbageCollector(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceParametersJavaGarbageCollector: valid values are %v", v, AllowedInstanceParametersJavaGarbageCollectorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceParametersJavaGarbageCollector) IsValid() bool {
	for _, existing := range AllowedInstanceParametersJavaGarbageCollectorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JavaGarbageCollectorJavaGarbageCollector value
func (v InstanceParametersJavaGarbageCollector) Ptr() *InstanceParametersJavaGarbageCollector {
	return &v
}

type NullableInstanceParametersJavaGarbageCollector struct {
	value *InstanceParametersJavaGarbageCollector
	isSet bool
}

func (v NullableInstanceParametersJavaGarbageCollector) Get() *InstanceParametersJavaGarbageCollector {
	return v.value
}

func (v *NullableInstanceParametersJavaGarbageCollector) Set(val *InstanceParametersJavaGarbageCollector) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceParametersJavaGarbageCollector) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceParametersJavaGarbageCollector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceParametersJavaGarbageCollector(val *InstanceParametersJavaGarbageCollector) *NullableInstanceParametersJavaGarbageCollector {
	return &NullableInstanceParametersJavaGarbageCollector{value: val, isSet: true}
}

func (v NullableInstanceParametersJavaGarbageCollector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceParametersJavaGarbageCollector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type InstanceParametersGetJavaGarbageCollectorAttributeType = *InstanceParametersJavaGarbageCollector
type InstanceParametersGetJavaGarbageCollectorArgType = InstanceParametersJavaGarbageCollector
type InstanceParametersGetJavaGarbageCollectorRetType = InstanceParametersJavaGarbageCollector

func getInstanceParametersGetJavaGarbageCollectorAttributeTypeOk(arg InstanceParametersGetJavaGarbageCollectorAttributeType) (ret InstanceParametersGetJavaGarbageCollectorRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetJavaGarbageCollectorAttributeType(arg *InstanceParametersGetJavaGarbageCollectorAttributeType, val InstanceParametersGetJavaGarbageCollectorRetType) {
	*arg = &val
}

/*
	types and functions for java_heapspace
*/

// isInteger
type InstanceParametersGetJavaHeapspaceAttributeType = *int64
type InstanceParametersGetJavaHeapspaceArgType = int64
type InstanceParametersGetJavaHeapspaceRetType = int64

func getInstanceParametersGetJavaHeapspaceAttributeTypeOk(arg InstanceParametersGetJavaHeapspaceAttributeType) (ret InstanceParametersGetJavaHeapspaceRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetJavaHeapspaceAttributeType(arg *InstanceParametersGetJavaHeapspaceAttributeType, val InstanceParametersGetJavaHeapspaceRetType) {
	*arg = &val
}

/*
	types and functions for java_maxmetaspace
*/

// isInteger
type InstanceParametersGetJavaMaxmetaspaceAttributeType = *int64
type InstanceParametersGetJavaMaxmetaspaceArgType = int64
type InstanceParametersGetJavaMaxmetaspaceRetType = int64

func getInstanceParametersGetJavaMaxmetaspaceAttributeTypeOk(arg InstanceParametersGetJavaMaxmetaspaceAttributeType) (ret InstanceParametersGetJavaMaxmetaspaceRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetJavaMaxmetaspaceAttributeType(arg *InstanceParametersGetJavaMaxmetaspaceAttributeType, val InstanceParametersGetJavaMaxmetaspaceRetType) {
	*arg = &val
}

/*
	types and functions for max_disk_threshold
*/

// isInteger
type InstanceParametersGetMaxDiskThresholdAttributeType = *int64
type InstanceParametersGetMaxDiskThresholdArgType = int64
type InstanceParametersGetMaxDiskThresholdRetType = int64

func getInstanceParametersGetMaxDiskThresholdAttributeTypeOk(arg InstanceParametersGetMaxDiskThresholdAttributeType) (ret InstanceParametersGetMaxDiskThresholdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetMaxDiskThresholdAttributeType(arg *InstanceParametersGetMaxDiskThresholdAttributeType, val InstanceParametersGetMaxDiskThresholdRetType) {
	*arg = &val
}

/*
	types and functions for metrics_frequency
*/

// isInteger
type InstanceParametersGetMetricsFrequencyAttributeType = *int64
type InstanceParametersGetMetricsFrequencyArgType = int64
type InstanceParametersGetMetricsFrequencyRetType = int64

func getInstanceParametersGetMetricsFrequencyAttributeTypeOk(arg InstanceParametersGetMetricsFrequencyAttributeType) (ret InstanceParametersGetMetricsFrequencyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetMetricsFrequencyAttributeType(arg *InstanceParametersGetMetricsFrequencyAttributeType, val InstanceParametersGetMetricsFrequencyRetType) {
	*arg = &val
}

/*
	types and functions for metrics_prefix
*/

// isNotNullableString
type InstanceParametersGetMetricsPrefixAttributeType = *string

func getInstanceParametersGetMetricsPrefixAttributeTypeOk(arg InstanceParametersGetMetricsPrefixAttributeType) (ret InstanceParametersGetMetricsPrefixRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetMetricsPrefixAttributeType(arg *InstanceParametersGetMetricsPrefixAttributeType, val InstanceParametersGetMetricsPrefixRetType) {
	*arg = &val
}

type InstanceParametersGetMetricsPrefixArgType = string
type InstanceParametersGetMetricsPrefixRetType = string

/*
	types and functions for monitoring_instance_id
*/

// isNotNullableString
type InstanceParametersGetMonitoringInstanceIdAttributeType = *string

func getInstanceParametersGetMonitoringInstanceIdAttributeTypeOk(arg InstanceParametersGetMonitoringInstanceIdAttributeType) (ret InstanceParametersGetMonitoringInstanceIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetMonitoringInstanceIdAttributeType(arg *InstanceParametersGetMonitoringInstanceIdAttributeType, val InstanceParametersGetMonitoringInstanceIdRetType) {
	*arg = &val
}

type InstanceParametersGetMonitoringInstanceIdArgType = string
type InstanceParametersGetMonitoringInstanceIdRetType = string

/*
	types and functions for plugins
*/

// isArray
type InstanceParametersGetPluginsAttributeType = *[]string
type InstanceParametersGetPluginsArgType = []string
type InstanceParametersGetPluginsRetType = []string

func getInstanceParametersGetPluginsAttributeTypeOk(arg InstanceParametersGetPluginsAttributeType) (ret InstanceParametersGetPluginsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetPluginsAttributeType(arg *InstanceParametersGetPluginsAttributeType, val InstanceParametersGetPluginsRetType) {
	*arg = &val
}

/*
	types and functions for sgw_acl
*/

// isNotNullableString
type InstanceParametersGetSgwAclAttributeType = *string

func getInstanceParametersGetSgwAclAttributeTypeOk(arg InstanceParametersGetSgwAclAttributeType) (ret InstanceParametersGetSgwAclRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetSgwAclAttributeType(arg *InstanceParametersGetSgwAclAttributeType, val InstanceParametersGetSgwAclRetType) {
	*arg = &val
}

type InstanceParametersGetSgwAclArgType = string
type InstanceParametersGetSgwAclRetType = string

/*
	types and functions for syslog
*/

// isArray
type InstanceParametersGetSyslogAttributeType = *[]string
type InstanceParametersGetSyslogArgType = []string
type InstanceParametersGetSyslogRetType = []string

func getInstanceParametersGetSyslogAttributeTypeOk(arg InstanceParametersGetSyslogAttributeType) (ret InstanceParametersGetSyslogRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetSyslogAttributeType(arg *InstanceParametersGetSyslogAttributeType, val InstanceParametersGetSyslogRetType) {
	*arg = &val
}

/*
	types and functions for tls-ciphers
*/

// isArray
type InstanceParametersGetTlsCiphersAttributeType = *[]string
type InstanceParametersGetTlsCiphersArgType = []string
type InstanceParametersGetTlsCiphersRetType = []string

func getInstanceParametersGetTlsCiphersAttributeTypeOk(arg InstanceParametersGetTlsCiphersAttributeType) (ret InstanceParametersGetTlsCiphersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetTlsCiphersAttributeType(arg *InstanceParametersGetTlsCiphersAttributeType, val InstanceParametersGetTlsCiphersRetType) {
	*arg = &val
}

/*
	types and functions for tls-protocols
*/

// isArray
type InstanceParametersGetTlsProtocolsAttributeType = *[]string
type InstanceParametersGetTlsProtocolsArgType = []string
type InstanceParametersGetTlsProtocolsRetType = []string

func getInstanceParametersGetTlsProtocolsAttributeTypeOk(arg InstanceParametersGetTlsProtocolsAttributeType) (ret InstanceParametersGetTlsProtocolsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceParametersGetTlsProtocolsAttributeType(arg *InstanceParametersGetTlsProtocolsAttributeType, val InstanceParametersGetTlsProtocolsRetType) {
	*arg = &val
}

// InstanceParameters struct for InstanceParameters
type InstanceParameters struct {
	EnableMonitoring InstanceParametersgetEnableMonitoringAttributeType `json:"enable_monitoring,omitempty"`
	// If you want to monitor your service with Graphite, you can set the custom parameter graphite. It expects the host and port where the Graphite metrics should be sent to.
	Graphite             InstanceParametersGetGraphiteAttributeType             `json:"graphite,omitempty"`
	JavaGarbageCollector InstanceParametersGetJavaGarbageCollectorAttributeType `json:"java_garbage_collector,omitempty"`
	// Default: not set, 46% of available memory will be used. The amount of memory (in MB) allocated as heap by the JVM for OpenSearch.
	// Can be cast to int32 without loss of precision.
	JavaHeapspace InstanceParametersGetJavaHeapspaceAttributeType `json:"java_heapspace,omitempty"`
	// The amount of memory (in MB) used by the JVM to store metadata for OpenSearch.
	// Can be cast to int32 without loss of precision.
	JavaMaxmetaspace InstanceParametersGetJavaMaxmetaspaceAttributeType `json:"java_maxmetaspace,omitempty"`
	// This component monitors ephemeral and persistent disk usage. If one of these disk usages reaches the default configured threshold of 80%, the a9s Parachute stops all processes on that node.
	// Can be cast to int32 without loss of precision.
	MaxDiskThreshold InstanceParametersGetMaxDiskThresholdAttributeType `json:"max_disk_threshold,omitempty"`
	// Frequency of metrics being emitted in seconds
	// Can be cast to int32 without loss of precision.
	MetricsFrequency InstanceParametersGetMetricsFrequencyAttributeType `json:"metrics_frequency,omitempty"`
	// Depending on your graphite provider, you might need to prefix the metrics with a certain value, like an API key for example.
	MetricsPrefix        InstanceParametersGetMetricsPrefixAttributeType        `json:"metrics_prefix,omitempty"`
	MonitoringInstanceId InstanceParametersGetMonitoringInstanceIdAttributeType `json:"monitoring_instance_id,omitempty"`
	// The plugins repository-s3 and repository-azure are enabled by default and cannot be disabled.
	Plugins InstanceParametersGetPluginsAttributeType `json:"plugins,omitempty"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	SgwAcl InstanceParametersGetSgwAclAttributeType `json:"sgw_acl,omitempty"`
	Syslog InstanceParametersGetSyslogAttributeType `json:"syslog,omitempty"`
	// Only Java format is supported.
	TlsCiphers   InstanceParametersGetTlsCiphersAttributeType   `json:"tls-ciphers,omitempty"`
	TlsProtocols InstanceParametersGetTlsProtocolsAttributeType `json:"tls-protocols,omitempty"`
}

// NewInstanceParameters instantiates a new InstanceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceParameters() *InstanceParameters {
	this := InstanceParameters{}
	return &this
}

// NewInstanceParametersWithDefaults instantiates a new InstanceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceParametersWithDefaults() *InstanceParameters {
	this := InstanceParameters{}
	var enableMonitoring bool = false
	this.EnableMonitoring = &enableMonitoring
	var javaGarbageCollector InstanceParametersJavaGarbageCollector = "UseG1GC"
	this.JavaGarbageCollector = &javaGarbageCollector
	var javaMaxmetaspace int64 = 512
	this.JavaMaxmetaspace = &javaMaxmetaspace
	var maxDiskThreshold int64 = 80
	this.MaxDiskThreshold = &maxDiskThreshold
	var metricsFrequency int64 = 10
	this.MetricsFrequency = &metricsFrequency
	return &this
}

// GetEnableMonitoring returns the EnableMonitoring field value if set, zero value otherwise.
func (o *InstanceParameters) GetEnableMonitoring() (res InstanceParametersgetEnableMonitoringRetType) {
	res, _ = o.GetEnableMonitoringOk()
	return
}

// GetEnableMonitoringOk returns a tuple with the EnableMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetEnableMonitoringOk() (ret InstanceParametersgetEnableMonitoringRetType, ok bool) {
	return getInstanceParametersgetEnableMonitoringAttributeTypeOk(o.EnableMonitoring)
}

// HasEnableMonitoring returns a boolean if a field has been set.
func (o *InstanceParameters) HasEnableMonitoring() bool {
	_, ok := o.GetEnableMonitoringOk()
	return ok
}

// SetEnableMonitoring gets a reference to the given bool and assigns it to the EnableMonitoring field.
func (o *InstanceParameters) SetEnableMonitoring(v InstanceParametersgetEnableMonitoringRetType) {
	setInstanceParametersgetEnableMonitoringAttributeType(&o.EnableMonitoring, v)
}

// GetGraphite returns the Graphite field value if set, zero value otherwise.
func (o *InstanceParameters) GetGraphite() (res InstanceParametersGetGraphiteRetType) {
	res, _ = o.GetGraphiteOk()
	return
}

// GetGraphiteOk returns a tuple with the Graphite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetGraphiteOk() (ret InstanceParametersGetGraphiteRetType, ok bool) {
	return getInstanceParametersGetGraphiteAttributeTypeOk(o.Graphite)
}

// HasGraphite returns a boolean if a field has been set.
func (o *InstanceParameters) HasGraphite() bool {
	_, ok := o.GetGraphiteOk()
	return ok
}

// SetGraphite gets a reference to the given string and assigns it to the Graphite field.
func (o *InstanceParameters) SetGraphite(v InstanceParametersGetGraphiteRetType) {
	setInstanceParametersGetGraphiteAttributeType(&o.Graphite, v)
}

// GetJavaGarbageCollector returns the JavaGarbageCollector field value if set, zero value otherwise.
func (o *InstanceParameters) GetJavaGarbageCollector() (res InstanceParametersGetJavaGarbageCollectorRetType) {
	res, _ = o.GetJavaGarbageCollectorOk()
	return
}

// GetJavaGarbageCollectorOk returns a tuple with the JavaGarbageCollector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetJavaGarbageCollectorOk() (ret InstanceParametersGetJavaGarbageCollectorRetType, ok bool) {
	return getInstanceParametersGetJavaGarbageCollectorAttributeTypeOk(o.JavaGarbageCollector)
}

// HasJavaGarbageCollector returns a boolean if a field has been set.
func (o *InstanceParameters) HasJavaGarbageCollector() bool {
	_, ok := o.GetJavaGarbageCollectorOk()
	return ok
}

// SetJavaGarbageCollector gets a reference to the given string and assigns it to the JavaGarbageCollector field.
func (o *InstanceParameters) SetJavaGarbageCollector(v InstanceParametersGetJavaGarbageCollectorRetType) {
	setInstanceParametersGetJavaGarbageCollectorAttributeType(&o.JavaGarbageCollector, v)
}

// GetJavaHeapspace returns the JavaHeapspace field value if set, zero value otherwise.
func (o *InstanceParameters) GetJavaHeapspace() (res InstanceParametersGetJavaHeapspaceRetType) {
	res, _ = o.GetJavaHeapspaceOk()
	return
}

// GetJavaHeapspaceOk returns a tuple with the JavaHeapspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetJavaHeapspaceOk() (ret InstanceParametersGetJavaHeapspaceRetType, ok bool) {
	return getInstanceParametersGetJavaHeapspaceAttributeTypeOk(o.JavaHeapspace)
}

// HasJavaHeapspace returns a boolean if a field has been set.
func (o *InstanceParameters) HasJavaHeapspace() bool {
	_, ok := o.GetJavaHeapspaceOk()
	return ok
}

// SetJavaHeapspace gets a reference to the given int64 and assigns it to the JavaHeapspace field.
func (o *InstanceParameters) SetJavaHeapspace(v InstanceParametersGetJavaHeapspaceRetType) {
	setInstanceParametersGetJavaHeapspaceAttributeType(&o.JavaHeapspace, v)
}

// GetJavaMaxmetaspace returns the JavaMaxmetaspace field value if set, zero value otherwise.
func (o *InstanceParameters) GetJavaMaxmetaspace() (res InstanceParametersGetJavaMaxmetaspaceRetType) {
	res, _ = o.GetJavaMaxmetaspaceOk()
	return
}

// GetJavaMaxmetaspaceOk returns a tuple with the JavaMaxmetaspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetJavaMaxmetaspaceOk() (ret InstanceParametersGetJavaMaxmetaspaceRetType, ok bool) {
	return getInstanceParametersGetJavaMaxmetaspaceAttributeTypeOk(o.JavaMaxmetaspace)
}

// HasJavaMaxmetaspace returns a boolean if a field has been set.
func (o *InstanceParameters) HasJavaMaxmetaspace() bool {
	_, ok := o.GetJavaMaxmetaspaceOk()
	return ok
}

// SetJavaMaxmetaspace gets a reference to the given int64 and assigns it to the JavaMaxmetaspace field.
func (o *InstanceParameters) SetJavaMaxmetaspace(v InstanceParametersGetJavaMaxmetaspaceRetType) {
	setInstanceParametersGetJavaMaxmetaspaceAttributeType(&o.JavaMaxmetaspace, v)
}

// GetMaxDiskThreshold returns the MaxDiskThreshold field value if set, zero value otherwise.
func (o *InstanceParameters) GetMaxDiskThreshold() (res InstanceParametersGetMaxDiskThresholdRetType) {
	res, _ = o.GetMaxDiskThresholdOk()
	return
}

// GetMaxDiskThresholdOk returns a tuple with the MaxDiskThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMaxDiskThresholdOk() (ret InstanceParametersGetMaxDiskThresholdRetType, ok bool) {
	return getInstanceParametersGetMaxDiskThresholdAttributeTypeOk(o.MaxDiskThreshold)
}

// HasMaxDiskThreshold returns a boolean if a field has been set.
func (o *InstanceParameters) HasMaxDiskThreshold() bool {
	_, ok := o.GetMaxDiskThresholdOk()
	return ok
}

// SetMaxDiskThreshold gets a reference to the given int64 and assigns it to the MaxDiskThreshold field.
func (o *InstanceParameters) SetMaxDiskThreshold(v InstanceParametersGetMaxDiskThresholdRetType) {
	setInstanceParametersGetMaxDiskThresholdAttributeType(&o.MaxDiskThreshold, v)
}

// GetMetricsFrequency returns the MetricsFrequency field value if set, zero value otherwise.
func (o *InstanceParameters) GetMetricsFrequency() (res InstanceParametersGetMetricsFrequencyRetType) {
	res, _ = o.GetMetricsFrequencyOk()
	return
}

// GetMetricsFrequencyOk returns a tuple with the MetricsFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMetricsFrequencyOk() (ret InstanceParametersGetMetricsFrequencyRetType, ok bool) {
	return getInstanceParametersGetMetricsFrequencyAttributeTypeOk(o.MetricsFrequency)
}

// HasMetricsFrequency returns a boolean if a field has been set.
func (o *InstanceParameters) HasMetricsFrequency() bool {
	_, ok := o.GetMetricsFrequencyOk()
	return ok
}

// SetMetricsFrequency gets a reference to the given int64 and assigns it to the MetricsFrequency field.
func (o *InstanceParameters) SetMetricsFrequency(v InstanceParametersGetMetricsFrequencyRetType) {
	setInstanceParametersGetMetricsFrequencyAttributeType(&o.MetricsFrequency, v)
}

// GetMetricsPrefix returns the MetricsPrefix field value if set, zero value otherwise.
func (o *InstanceParameters) GetMetricsPrefix() (res InstanceParametersGetMetricsPrefixRetType) {
	res, _ = o.GetMetricsPrefixOk()
	return
}

// GetMetricsPrefixOk returns a tuple with the MetricsPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMetricsPrefixOk() (ret InstanceParametersGetMetricsPrefixRetType, ok bool) {
	return getInstanceParametersGetMetricsPrefixAttributeTypeOk(o.MetricsPrefix)
}

// HasMetricsPrefix returns a boolean if a field has been set.
func (o *InstanceParameters) HasMetricsPrefix() bool {
	_, ok := o.GetMetricsPrefixOk()
	return ok
}

// SetMetricsPrefix gets a reference to the given string and assigns it to the MetricsPrefix field.
func (o *InstanceParameters) SetMetricsPrefix(v InstanceParametersGetMetricsPrefixRetType) {
	setInstanceParametersGetMetricsPrefixAttributeType(&o.MetricsPrefix, v)
}

// GetMonitoringInstanceId returns the MonitoringInstanceId field value if set, zero value otherwise.
func (o *InstanceParameters) GetMonitoringInstanceId() (res InstanceParametersGetMonitoringInstanceIdRetType) {
	res, _ = o.GetMonitoringInstanceIdOk()
	return
}

// GetMonitoringInstanceIdOk returns a tuple with the MonitoringInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMonitoringInstanceIdOk() (ret InstanceParametersGetMonitoringInstanceIdRetType, ok bool) {
	return getInstanceParametersGetMonitoringInstanceIdAttributeTypeOk(o.MonitoringInstanceId)
}

// HasMonitoringInstanceId returns a boolean if a field has been set.
func (o *InstanceParameters) HasMonitoringInstanceId() bool {
	_, ok := o.GetMonitoringInstanceIdOk()
	return ok
}

// SetMonitoringInstanceId gets a reference to the given string and assigns it to the MonitoringInstanceId field.
func (o *InstanceParameters) SetMonitoringInstanceId(v InstanceParametersGetMonitoringInstanceIdRetType) {
	setInstanceParametersGetMonitoringInstanceIdAttributeType(&o.MonitoringInstanceId, v)
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *InstanceParameters) GetPlugins() (res InstanceParametersGetPluginsRetType) {
	res, _ = o.GetPluginsOk()
	return
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetPluginsOk() (ret InstanceParametersGetPluginsRetType, ok bool) {
	return getInstanceParametersGetPluginsAttributeTypeOk(o.Plugins)
}

// HasPlugins returns a boolean if a field has been set.
func (o *InstanceParameters) HasPlugins() bool {
	_, ok := o.GetPluginsOk()
	return ok
}

// SetPlugins gets a reference to the given []string and assigns it to the Plugins field.
func (o *InstanceParameters) SetPlugins(v InstanceParametersGetPluginsRetType) {
	setInstanceParametersGetPluginsAttributeType(&o.Plugins, v)
}

// GetSgwAcl returns the SgwAcl field value if set, zero value otherwise.
func (o *InstanceParameters) GetSgwAcl() (res InstanceParametersGetSgwAclRetType) {
	res, _ = o.GetSgwAclOk()
	return
}

// GetSgwAclOk returns a tuple with the SgwAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetSgwAclOk() (ret InstanceParametersGetSgwAclRetType, ok bool) {
	return getInstanceParametersGetSgwAclAttributeTypeOk(o.SgwAcl)
}

// HasSgwAcl returns a boolean if a field has been set.
func (o *InstanceParameters) HasSgwAcl() bool {
	_, ok := o.GetSgwAclOk()
	return ok
}

// SetSgwAcl gets a reference to the given string and assigns it to the SgwAcl field.
func (o *InstanceParameters) SetSgwAcl(v InstanceParametersGetSgwAclRetType) {
	setInstanceParametersGetSgwAclAttributeType(&o.SgwAcl, v)
}

// GetSyslog returns the Syslog field value if set, zero value otherwise.
func (o *InstanceParameters) GetSyslog() (res InstanceParametersGetSyslogRetType) {
	res, _ = o.GetSyslogOk()
	return
}

// GetSyslogOk returns a tuple with the Syslog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetSyslogOk() (ret InstanceParametersGetSyslogRetType, ok bool) {
	return getInstanceParametersGetSyslogAttributeTypeOk(o.Syslog)
}

// HasSyslog returns a boolean if a field has been set.
func (o *InstanceParameters) HasSyslog() bool {
	_, ok := o.GetSyslogOk()
	return ok
}

// SetSyslog gets a reference to the given []string and assigns it to the Syslog field.
func (o *InstanceParameters) SetSyslog(v InstanceParametersGetSyslogRetType) {
	setInstanceParametersGetSyslogAttributeType(&o.Syslog, v)
}

// GetTlsCiphers returns the TlsCiphers field value if set, zero value otherwise.
func (o *InstanceParameters) GetTlsCiphers() (res InstanceParametersGetTlsCiphersRetType) {
	res, _ = o.GetTlsCiphersOk()
	return
}

// GetTlsCiphersOk returns a tuple with the TlsCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetTlsCiphersOk() (ret InstanceParametersGetTlsCiphersRetType, ok bool) {
	return getInstanceParametersGetTlsCiphersAttributeTypeOk(o.TlsCiphers)
}

// HasTlsCiphers returns a boolean if a field has been set.
func (o *InstanceParameters) HasTlsCiphers() bool {
	_, ok := o.GetTlsCiphersOk()
	return ok
}

// SetTlsCiphers gets a reference to the given []string and assigns it to the TlsCiphers field.
func (o *InstanceParameters) SetTlsCiphers(v InstanceParametersGetTlsCiphersRetType) {
	setInstanceParametersGetTlsCiphersAttributeType(&o.TlsCiphers, v)
}

// GetTlsProtocols returns the TlsProtocols field value if set, zero value otherwise.
func (o *InstanceParameters) GetTlsProtocols() (res InstanceParametersGetTlsProtocolsRetType) {
	res, _ = o.GetTlsProtocolsOk()
	return
}

// GetTlsProtocolsOk returns a tuple with the TlsProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetTlsProtocolsOk() (ret InstanceParametersGetTlsProtocolsRetType, ok bool) {
	return getInstanceParametersGetTlsProtocolsAttributeTypeOk(o.TlsProtocols)
}

// HasTlsProtocols returns a boolean if a field has been set.
func (o *InstanceParameters) HasTlsProtocols() bool {
	_, ok := o.GetTlsProtocolsOk()
	return ok
}

// SetTlsProtocols gets a reference to the given []string and assigns it to the TlsProtocols field.
func (o *InstanceParameters) SetTlsProtocols(v InstanceParametersGetTlsProtocolsRetType) {
	setInstanceParametersGetTlsProtocolsAttributeType(&o.TlsProtocols, v)
}

func (o InstanceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getInstanceParametersgetEnableMonitoringAttributeTypeOk(o.EnableMonitoring); ok {
		toSerialize["EnableMonitoring"] = val
	}
	if val, ok := getInstanceParametersGetGraphiteAttributeTypeOk(o.Graphite); ok {
		toSerialize["Graphite"] = val
	}
	if val, ok := getInstanceParametersGetJavaGarbageCollectorAttributeTypeOk(o.JavaGarbageCollector); ok {
		toSerialize["JavaGarbageCollector"] = val
	}
	if val, ok := getInstanceParametersGetJavaHeapspaceAttributeTypeOk(o.JavaHeapspace); ok {
		toSerialize["JavaHeapspace"] = val
	}
	if val, ok := getInstanceParametersGetJavaMaxmetaspaceAttributeTypeOk(o.JavaMaxmetaspace); ok {
		toSerialize["JavaMaxmetaspace"] = val
	}
	if val, ok := getInstanceParametersGetMaxDiskThresholdAttributeTypeOk(o.MaxDiskThreshold); ok {
		toSerialize["MaxDiskThreshold"] = val
	}
	if val, ok := getInstanceParametersGetMetricsFrequencyAttributeTypeOk(o.MetricsFrequency); ok {
		toSerialize["MetricsFrequency"] = val
	}
	if val, ok := getInstanceParametersGetMetricsPrefixAttributeTypeOk(o.MetricsPrefix); ok {
		toSerialize["MetricsPrefix"] = val
	}
	if val, ok := getInstanceParametersGetMonitoringInstanceIdAttributeTypeOk(o.MonitoringInstanceId); ok {
		toSerialize["MonitoringInstanceId"] = val
	}
	if val, ok := getInstanceParametersGetPluginsAttributeTypeOk(o.Plugins); ok {
		toSerialize["Plugins"] = val
	}
	if val, ok := getInstanceParametersGetSgwAclAttributeTypeOk(o.SgwAcl); ok {
		toSerialize["SgwAcl"] = val
	}
	if val, ok := getInstanceParametersGetSyslogAttributeTypeOk(o.Syslog); ok {
		toSerialize["Syslog"] = val
	}
	if val, ok := getInstanceParametersGetTlsCiphersAttributeTypeOk(o.TlsCiphers); ok {
		toSerialize["TlsCiphers"] = val
	}
	if val, ok := getInstanceParametersGetTlsProtocolsAttributeTypeOk(o.TlsProtocols); ok {
		toSerialize["TlsProtocols"] = val
	}
	return toSerialize, nil
}

type NullableInstanceParameters struct {
	value *InstanceParameters
	isSet bool
}

func (v NullableInstanceParameters) Get() *InstanceParameters {
	return v.value
}

func (v *NullableInstanceParameters) Set(val *InstanceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceParameters(val *InstanceParameters) *NullableInstanceParameters {
	return &NullableInstanceParameters{value: val, isSet: true}
}

func (v NullableInstanceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
