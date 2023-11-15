/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationApiRequests200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApiRequests200ResponseInner{}

// GetOrganizationApiRequests200ResponseInner struct for GetOrganizationApiRequests200ResponseInner
type GetOrganizationApiRequests200ResponseInner struct {
	// Database ID for the admin user who made the API request.
	AdminId *string `json:"adminId,omitempty"`
	// HTTP method used in the API request.
	Method *string `json:"method,omitempty"`
	// The host which the API request was directed at.
	Host *string `json:"host,omitempty"`
	// The API request path.
	Path *string `json:"path,omitempty"`
	// The query string sent with the API request.
	QueryString *string `json:"queryString,omitempty"`
	// The API request user agent.
	UserAgent *string `json:"userAgent,omitempty"`
	// Timestamp, in iso8601 format, indicating when the API request was made.
	Ts *time.Time `json:"ts,omitempty"`
	// API request response code.
	ResponseCode *int32 `json:"responseCode,omitempty"`
	// Public IP address from which the API request was made.
	SourceIp *string `json:"sourceIp,omitempty"`
	// API version of the endpoint.
	Version *int32 `json:"version,omitempty"`
	// Operation ID for the endpoint.
	OperationId *string `json:"operationId,omitempty"`
}

// NewGetOrganizationApiRequests200ResponseInner instantiates a new GetOrganizationApiRequests200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApiRequests200ResponseInner() *GetOrganizationApiRequests200ResponseInner {
	this := GetOrganizationApiRequests200ResponseInner{}
	return &this
}

// NewGetOrganizationApiRequests200ResponseInnerWithDefaults instantiates a new GetOrganizationApiRequests200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApiRequests200ResponseInnerWithDefaults() *GetOrganizationApiRequests200ResponseInner {
	this := GetOrganizationApiRequests200ResponseInner{}
	return &this
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetAdminId() string {
	if o == nil || IsNil(o.AdminId) {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetAdminIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdminId) {
		return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasAdminId() bool {
	if o != nil && !IsNil(o.AdminId) {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *GetOrganizationApiRequests200ResponseInner) SetAdminId(v string) {
	o.AdminId = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *GetOrganizationApiRequests200ResponseInner) SetMethod(v string) {
	o.Method = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetOrganizationApiRequests200ResponseInner) SetHost(v string) {
	o.Host = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GetOrganizationApiRequests200ResponseInner) SetPath(v string) {
	o.Path = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetQueryString() string {
	if o == nil || IsNil(o.QueryString) {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *GetOrganizationApiRequests200ResponseInner) SetQueryString(v string) {
	o.QueryString = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *GetOrganizationApiRequests200ResponseInner) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetOrganizationApiRequests200ResponseInner) SetTs(v time.Time) {
	o.Ts = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetResponseCode() int32 {
	if o == nil || IsNil(o.ResponseCode) {
		var ret int32
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetResponseCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasResponseCode() bool {
	if o != nil && !IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int32 and assigns it to the ResponseCode field.
func (o *GetOrganizationApiRequests200ResponseInner) SetResponseCode(v int32) {
	o.ResponseCode = &v
}

// GetSourceIp returns the SourceIp field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetSourceIp() string {
	if o == nil || IsNil(o.SourceIp) {
		var ret string
		return ret
	}
	return *o.SourceIp
}

// GetSourceIpOk returns a tuple with the SourceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetSourceIpOk() (*string, bool) {
	if o == nil || IsNil(o.SourceIp) {
		return nil, false
	}
	return o.SourceIp, true
}

// HasSourceIp returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasSourceIp() bool {
	if o != nil && !IsNil(o.SourceIp) {
		return true
	}

	return false
}

// SetSourceIp gets a reference to the given string and assigns it to the SourceIp field.
func (o *GetOrganizationApiRequests200ResponseInner) SetSourceIp(v string) {
	o.SourceIp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *GetOrganizationApiRequests200ResponseInner) SetVersion(v int32) {
	o.Version = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *GetOrganizationApiRequests200ResponseInner) GetOperationId() string {
	if o == nil || IsNil(o.OperationId) {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequests200ResponseInner) GetOperationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperationId) {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *GetOrganizationApiRequests200ResponseInner) HasOperationId() bool {
	if o != nil && !IsNil(o.OperationId) {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *GetOrganizationApiRequests200ResponseInner) SetOperationId(v string) {
	o.OperationId = &v
}

func (o GetOrganizationApiRequests200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApiRequests200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminId) {
		toSerialize["adminId"] = o.AdminId
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.QueryString) {
		toSerialize["queryString"] = o.QueryString
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !IsNil(o.SourceIp) {
		toSerialize["sourceIp"] = o.SourceIp
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.OperationId) {
		toSerialize["operationId"] = o.OperationId
	}
	return toSerialize, nil
}

type NullableGetOrganizationApiRequests200ResponseInner struct {
	value *GetOrganizationApiRequests200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationApiRequests200ResponseInner) Get() *GetOrganizationApiRequests200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationApiRequests200ResponseInner) Set(val *GetOrganizationApiRequests200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApiRequests200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApiRequests200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApiRequests200ResponseInner(val *GetOrganizationApiRequests200ResponseInner) *NullableGetOrganizationApiRequests200ResponseInner {
	return &NullableGetOrganizationApiRequests200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApiRequests200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApiRequests200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


