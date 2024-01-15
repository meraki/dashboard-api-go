/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkWirelessSsidEapOverride200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidEapOverride200Response{}

// GetNetworkWirelessSsidEapOverride200Response struct for GetNetworkWirelessSsidEapOverride200Response
type GetNetworkWirelessSsidEapOverride200Response struct {
	// General EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
	// Maximum number of general EAP retries.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	Identity *GetNetworkWirelessSsidEapOverride200ResponseIdentity `json:"identity,omitempty"`
	EapolKey *GetNetworkWirelessSsidEapOverride200ResponseEapolKey `json:"eapolKey,omitempty"`
}

// NewGetNetworkWirelessSsidEapOverride200Response instantiates a new GetNetworkWirelessSsidEapOverride200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidEapOverride200Response() *GetNetworkWirelessSsidEapOverride200Response {
	this := GetNetworkWirelessSsidEapOverride200Response{}
	return &this
}

// NewGetNetworkWirelessSsidEapOverride200ResponseWithDefaults instantiates a new GetNetworkWirelessSsidEapOverride200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidEapOverride200ResponseWithDefaults() *GetNetworkWirelessSsidEapOverride200Response {
	this := GetNetworkWirelessSsidEapOverride200Response{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *GetNetworkWirelessSsidEapOverride200Response) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetMaxRetries() int32 {
	if o == nil || IsNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRetries) {
		return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) HasMaxRetries() bool {
	if o != nil && !IsNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *GetNetworkWirelessSsidEapOverride200Response) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetIdentity() GetNetworkWirelessSsidEapOverride200ResponseIdentity {
	if o == nil || IsNil(o.Identity) {
		var ret GetNetworkWirelessSsidEapOverride200ResponseIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetIdentityOk() (*GetNetworkWirelessSsidEapOverride200ResponseIdentity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given GetNetworkWirelessSsidEapOverride200ResponseIdentity and assigns it to the Identity field.
func (o *GetNetworkWirelessSsidEapOverride200Response) SetIdentity(v GetNetworkWirelessSsidEapOverride200ResponseIdentity) {
	o.Identity = &v
}

// GetEapolKey returns the EapolKey field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetEapolKey() GetNetworkWirelessSsidEapOverride200ResponseEapolKey {
	if o == nil || IsNil(o.EapolKey) {
		var ret GetNetworkWirelessSsidEapOverride200ResponseEapolKey
		return ret
	}
	return *o.EapolKey
}

// GetEapolKeyOk returns a tuple with the EapolKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) GetEapolKeyOk() (*GetNetworkWirelessSsidEapOverride200ResponseEapolKey, bool) {
	if o == nil || IsNil(o.EapolKey) {
		return nil, false
	}
	return o.EapolKey, true
}

// HasEapolKey returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidEapOverride200Response) HasEapolKey() bool {
	if o != nil && !IsNil(o.EapolKey) {
		return true
	}

	return false
}

// SetEapolKey gets a reference to the given GetNetworkWirelessSsidEapOverride200ResponseEapolKey and assigns it to the EapolKey field.
func (o *GetNetworkWirelessSsidEapOverride200Response) SetEapolKey(v GetNetworkWirelessSsidEapOverride200ResponseEapolKey) {
	o.EapolKey = &v
}

func (o GetNetworkWirelessSsidEapOverride200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidEapOverride200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.MaxRetries) {
		toSerialize["maxRetries"] = o.MaxRetries
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.EapolKey) {
		toSerialize["eapolKey"] = o.EapolKey
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidEapOverride200Response struct {
	value *GetNetworkWirelessSsidEapOverride200Response
	isSet bool
}

func (v NullableGetNetworkWirelessSsidEapOverride200Response) Get() *GetNetworkWirelessSsidEapOverride200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidEapOverride200Response) Set(val *GetNetworkWirelessSsidEapOverride200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidEapOverride200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidEapOverride200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidEapOverride200Response(val *GetNetworkWirelessSsidEapOverride200Response) *NullableGetNetworkWirelessSsidEapOverride200Response {
	return &NullableGetNetworkWirelessSsidEapOverride200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidEapOverride200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidEapOverride200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


