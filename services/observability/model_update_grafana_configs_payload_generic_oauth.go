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

// checks if the UpdateGrafanaConfigsPayloadGenericOauth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGrafanaConfigsPayloadGenericOauth{}

/*
	types and functions for apiUrl
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlRetType = string

/*
	types and functions for authUrl
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlRetType = string

/*
	types and functions for enabled
*/

// isBoolean
type UpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeType = *bool
type UpdateGrafanaConfigsPayloadGenericOauthgetEnabledArgType = bool
type UpdateGrafanaConfigsPayloadGenericOauthgetEnabledRetType = bool

func getUpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthgetEnabledRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthgetEnabledRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetNameRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetNameArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetNameRetType = string

/*
	types and functions for oauthClientId
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdRetType = string

/*
	types and functions for oauthClientSecret
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretRetType = string

/*
	types and functions for roleAttributePath
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathRetType = string

/*
	types and functions for roleAttributeStrict
*/

// isBoolean
type UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeType = *bool
type UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictArgType = bool
type UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictRetType = bool

func getUpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictRetType) {
	*arg = &val
}

/*
	types and functions for scopes
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetScopesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetScopesRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetScopesArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetScopesRetType = string

/*
	types and functions for tokenUrl
*/

// isNotNullableString
type UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeType = *string

func getUpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlRetType) {
	*arg = &val
}

type UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlArgType = string
type UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlRetType = string

/*
	types and functions for usePkce
*/

// isBoolean
type UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeType = *bool
type UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceArgType = bool
type UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceRetType = bool

func getUpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeTypeOk(arg UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeType) (ret UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeType(arg *UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeType, val UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceRetType) {
	*arg = &val
}

// UpdateGrafanaConfigsPayloadGenericOauth struct for UpdateGrafanaConfigsPayloadGenericOauth
type UpdateGrafanaConfigsPayloadGenericOauth struct {
	// Set api_url to the resource that returns OpenID UserInfo compatible information.
	// REQUIRED
	ApiUrl UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeType `json:"apiUrl" required:"true"`
	// Authentication endpoint of idp.
	// REQUIRED
	AuthUrl UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeType `json:"authUrl" required:"true"`
	// enable or disable generic oauth login
	// REQUIRED
	Enabled UpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeType `json:"enabled" required:"true"`
	// Display name for the oAuth provider
	Name UpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeType `json:"name,omitempty"`
	// Oauth client id for auth endpoint.
	// REQUIRED
	OauthClientId UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeType `json:"oauthClientId" required:"true"`
	// Oauth client secret for auth endpoint.
	// REQUIRED
	OauthClientSecret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeType `json:"oauthClientSecret" required:"true"`
	// Grafana checks for the presence of a role using the JMESPath specified via the role_attribute_path configuration option. The JMESPath is applied to the id_token first. If there is no match, then the UserInfo endpoint specified via the api_url configuration option is tried next. The result after evaluation of the role_attribute_path JMESPath expression should be a valid Grafana role, for example, Viewer, Editor or Admin For example: contains(roles[\\*], 'grafana-admin') && 'Admin' || contains(roles[\\*], 'grafana-editor') && 'Editor' || contains(roles[\\*], 'grafana-viewer') && 'Viewer'
	// REQUIRED
	RoleAttributePath UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeType `json:"roleAttributePath" required:"true"`
	// If  therole_attribute_path property does not return a role, then the user is assigned the Viewer role by default. You can disable the role assignment by setting role_attribute_strict = true. It denies user access if no role or an invalid role is returned.
	RoleAttributeStrict UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeType `json:"roleAttributeStrict,omitempty"`
	// Space seperated list of scopes of the token
	Scopes UpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeType `json:"scopes,omitempty"`
	// Token endpoint of the idp.
	// REQUIRED
	TokenUrl UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeType `json:"tokenUrl" required:"true"`
	// enable or disable Proof Key for Code Exchange
	UsePkce UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeType `json:"usePkce,omitempty"`
}

type _UpdateGrafanaConfigsPayloadGenericOauth UpdateGrafanaConfigsPayloadGenericOauth

// NewUpdateGrafanaConfigsPayloadGenericOauth instantiates a new UpdateGrafanaConfigsPayloadGenericOauth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGrafanaConfigsPayloadGenericOauth(apiUrl UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlArgType, authUrl UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlArgType, enabled UpdateGrafanaConfigsPayloadGenericOauthgetEnabledArgType, oauthClientId UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdArgType, oauthClientSecret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretArgType, roleAttributePath UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathArgType, tokenUrl UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlArgType) *UpdateGrafanaConfigsPayloadGenericOauth {
	this := UpdateGrafanaConfigsPayloadGenericOauth{}
	setUpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeType(&this.ApiUrl, apiUrl)
	setUpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeType(&this.AuthUrl, authUrl)
	setUpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeType(&this.Enabled, enabled)
	setUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeType(&this.OauthClientId, oauthClientId)
	setUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeType(&this.OauthClientSecret, oauthClientSecret)
	setUpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeType(&this.RoleAttributePath, roleAttributePath)
	setUpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeType(&this.TokenUrl, tokenUrl)
	return &this
}

// NewUpdateGrafanaConfigsPayloadGenericOauthWithDefaults instantiates a new UpdateGrafanaConfigsPayloadGenericOauth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGrafanaConfigsPayloadGenericOauthWithDefaults() *UpdateGrafanaConfigsPayloadGenericOauth {
	this := UpdateGrafanaConfigsPayloadGenericOauth{}
	var roleAttributeStrict bool = true
	this.RoleAttributeStrict = &roleAttributeStrict
	var scopes string = "openid profile email"
	this.Scopes = &scopes
	return &this
}

// GetApiUrl returns the ApiUrl field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetApiUrl() (ret UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlRetType) {
	ret, _ = o.GetApiUrlOk()
	return ret
}

// GetApiUrlOk returns a tuple with the ApiUrl field value
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetApiUrlOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeTypeOk(o.ApiUrl)
}

// SetApiUrl sets field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetApiUrl(v UpdateGrafanaConfigsPayloadGenericOauthGetApiUrlRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeType(&o.ApiUrl, v)
}

// GetAuthUrl returns the AuthUrl field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetAuthUrl() (ret UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlRetType) {
	ret, _ = o.GetAuthUrlOk()
	return ret
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetAuthUrlOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeTypeOk(o.AuthUrl)
}

// SetAuthUrl sets field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetAuthUrl(v UpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeType(&o.AuthUrl, v)
}

// GetEnabled returns the Enabled field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetEnabled() (ret UpdateGrafanaConfigsPayloadGenericOauthgetEnabledRetType) {
	ret, _ = o.GetEnabledOk()
	return ret
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetEnabledOk() (ret UpdateGrafanaConfigsPayloadGenericOauthgetEnabledRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeTypeOk(o.Enabled)
}

// SetEnabled sets field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetEnabled(v UpdateGrafanaConfigsPayloadGenericOauthgetEnabledRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeType(&o.Enabled, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetName() (res UpdateGrafanaConfigsPayloadGenericOauthGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetNameOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetNameRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetName(v UpdateGrafanaConfigsPayloadGenericOauthGetNameRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeType(&o.Name, v)
}

// GetOauthClientId returns the OauthClientId field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetOauthClientId() (ret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdRetType) {
	ret, _ = o.GetOauthClientIdOk()
	return ret
}

// GetOauthClientIdOk returns a tuple with the OauthClientId field value
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetOauthClientIdOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeTypeOk(o.OauthClientId)
}

// SetOauthClientId sets field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetOauthClientId(v UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeType(&o.OauthClientId, v)
}

// GetOauthClientSecret returns the OauthClientSecret field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetOauthClientSecret() (ret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretRetType) {
	ret, _ = o.GetOauthClientSecretOk()
	return ret
}

// GetOauthClientSecretOk returns a tuple with the OauthClientSecret field value
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetOauthClientSecretOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeTypeOk(o.OauthClientSecret)
}

// SetOauthClientSecret sets field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetOauthClientSecret(v UpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeType(&o.OauthClientSecret, v)
}

// GetRoleAttributePath returns the RoleAttributePath field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetRoleAttributePath() (ret UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathRetType) {
	ret, _ = o.GetRoleAttributePathOk()
	return ret
}

// GetRoleAttributePathOk returns a tuple with the RoleAttributePath field value
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetRoleAttributePathOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeTypeOk(o.RoleAttributePath)
}

// SetRoleAttributePath sets field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetRoleAttributePath(v UpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeType(&o.RoleAttributePath, v)
}

// GetRoleAttributeStrict returns the RoleAttributeStrict field value if set, zero value otherwise.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetRoleAttributeStrict() (res UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictRetType) {
	res, _ = o.GetRoleAttributeStrictOk()
	return
}

// GetRoleAttributeStrictOk returns a tuple with the RoleAttributeStrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetRoleAttributeStrictOk() (ret UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeTypeOk(o.RoleAttributeStrict)
}

// HasRoleAttributeStrict returns a boolean if a field has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) HasRoleAttributeStrict() bool {
	_, ok := o.GetRoleAttributeStrictOk()
	return ok
}

// SetRoleAttributeStrict gets a reference to the given bool and assigns it to the RoleAttributeStrict field.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetRoleAttributeStrict(v UpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeType(&o.RoleAttributeStrict, v)
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetScopes() (res UpdateGrafanaConfigsPayloadGenericOauthGetScopesRetType) {
	res, _ = o.GetScopesOk()
	return
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetScopesOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetScopesRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeTypeOk(o.Scopes)
}

// HasScopes returns a boolean if a field has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) HasScopes() bool {
	_, ok := o.GetScopesOk()
	return ok
}

// SetScopes gets a reference to the given string and assigns it to the Scopes field.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetScopes(v UpdateGrafanaConfigsPayloadGenericOauthGetScopesRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeType(&o.Scopes, v)
}

// GetTokenUrl returns the TokenUrl field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetTokenUrl() (ret UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlRetType) {
	ret, _ = o.GetTokenUrlOk()
	return ret
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetTokenUrlOk() (ret UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeTypeOk(o.TokenUrl)
}

// SetTokenUrl sets field value
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetTokenUrl(v UpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeType(&o.TokenUrl, v)
}

// GetUsePkce returns the UsePkce field value if set, zero value otherwise.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetUsePkce() (res UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceRetType) {
	res, _ = o.GetUsePkceOk()
	return
}

// GetUsePkceOk returns a tuple with the UsePkce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) GetUsePkceOk() (ret UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceRetType, ok bool) {
	return getUpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeTypeOk(o.UsePkce)
}

// HasUsePkce returns a boolean if a field has been set.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) HasUsePkce() bool {
	_, ok := o.GetUsePkceOk()
	return ok
}

// SetUsePkce gets a reference to the given bool and assigns it to the UsePkce field.
func (o *UpdateGrafanaConfigsPayloadGenericOauth) SetUsePkce(v UpdateGrafanaConfigsPayloadGenericOauthgetUsePkceRetType) {
	setUpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeType(&o.UsePkce, v)
}

func (o UpdateGrafanaConfigsPayloadGenericOauth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetApiUrlAttributeTypeOk(o.ApiUrl); ok {
		toSerialize["ApiUrl"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetAuthUrlAttributeTypeOk(o.AuthUrl); ok {
		toSerialize["AuthUrl"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthgetEnabledAttributeTypeOk(o.Enabled); ok {
		toSerialize["Enabled"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientIdAttributeTypeOk(o.OauthClientId); ok {
		toSerialize["OauthClientId"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetOauthClientSecretAttributeTypeOk(o.OauthClientSecret); ok {
		toSerialize["OauthClientSecret"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetRoleAttributePathAttributeTypeOk(o.RoleAttributePath); ok {
		toSerialize["RoleAttributePath"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthgetRoleAttributeStrictAttributeTypeOk(o.RoleAttributeStrict); ok {
		toSerialize["RoleAttributeStrict"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetScopesAttributeTypeOk(o.Scopes); ok {
		toSerialize["Scopes"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthGetTokenUrlAttributeTypeOk(o.TokenUrl); ok {
		toSerialize["TokenUrl"] = val
	}
	if val, ok := getUpdateGrafanaConfigsPayloadGenericOauthgetUsePkceAttributeTypeOk(o.UsePkce); ok {
		toSerialize["UsePkce"] = val
	}
	return toSerialize, nil
}

type NullableUpdateGrafanaConfigsPayloadGenericOauth struct {
	value *UpdateGrafanaConfigsPayloadGenericOauth
	isSet bool
}

func (v NullableUpdateGrafanaConfigsPayloadGenericOauth) Get() *UpdateGrafanaConfigsPayloadGenericOauth {
	return v.value
}

func (v *NullableUpdateGrafanaConfigsPayloadGenericOauth) Set(val *UpdateGrafanaConfigsPayloadGenericOauth) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGrafanaConfigsPayloadGenericOauth) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGrafanaConfigsPayloadGenericOauth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGrafanaConfigsPayloadGenericOauth(val *UpdateGrafanaConfigsPayloadGenericOauth) *NullableUpdateGrafanaConfigsPayloadGenericOauth {
	return &NullableUpdateGrafanaConfigsPayloadGenericOauth{value: val, isSet: true}
}

func (v NullableUpdateGrafanaConfigsPayloadGenericOauth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGrafanaConfigsPayloadGenericOauth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
