/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkWirelessSsidVpnRequestFailover Secondary VPN concentrator settings. This is only used when two VPN concentrators are configured on the SSID.
type UpdateNetworkWirelessSsidVpnRequestFailover struct {
	// IP addressed reserved on DHCP server where SSID will terminate.
	RequestIp *string `json:"requestIp,omitempty"`
	// Idle timer interval in seconds.
	HeartbeatInterval *int32 `json:"heartbeatInterval,omitempty"`
	// Idle timer timeout in seconds.
	IdleTimeout *int32 `json:"idleTimeout,omitempty"`
}

// NewUpdateNetworkWirelessSsidVpnRequestFailover instantiates a new UpdateNetworkWirelessSsidVpnRequestFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidVpnRequestFailover() *UpdateNetworkWirelessSsidVpnRequestFailover {
	this := UpdateNetworkWirelessSsidVpnRequestFailover{}
	return &this
}

// NewUpdateNetworkWirelessSsidVpnRequestFailoverWithDefaults instantiates a new UpdateNetworkWirelessSsidVpnRequestFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidVpnRequestFailoverWithDefaults() *UpdateNetworkWirelessSsidVpnRequestFailover {
	this := UpdateNetworkWirelessSsidVpnRequestFailover{}
	return &this
}

// GetRequestIp returns the RequestIp field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) GetRequestIp() string {
	if o == nil || isNil(o.RequestIp) {
		var ret string
		return ret
	}
	return *o.RequestIp
}

// GetRequestIpOk returns a tuple with the RequestIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) GetRequestIpOk() (*string, bool) {
	if o == nil || isNil(o.RequestIp) {
    return nil, false
	}
	return o.RequestIp, true
}

// HasRequestIp returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) HasRequestIp() bool {
	if o != nil && !isNil(o.RequestIp) {
		return true
	}

	return false
}

// SetRequestIp gets a reference to the given string and assigns it to the RequestIp field.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) SetRequestIp(v string) {
	o.RequestIp = &v
}

// GetHeartbeatInterval returns the HeartbeatInterval field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) GetHeartbeatInterval() int32 {
	if o == nil || isNil(o.HeartbeatInterval) {
		var ret int32
		return ret
	}
	return *o.HeartbeatInterval
}

// GetHeartbeatIntervalOk returns a tuple with the HeartbeatInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) GetHeartbeatIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.HeartbeatInterval) {
    return nil, false
	}
	return o.HeartbeatInterval, true
}

// HasHeartbeatInterval returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) HasHeartbeatInterval() bool {
	if o != nil && !isNil(o.HeartbeatInterval) {
		return true
	}

	return false
}

// SetHeartbeatInterval gets a reference to the given int32 and assigns it to the HeartbeatInterval field.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) SetHeartbeatInterval(v int32) {
	o.HeartbeatInterval = &v
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) GetIdleTimeout() int32 {
	if o == nil || isNil(o.IdleTimeout) {
		var ret int32
		return ret
	}
	return *o.IdleTimeout
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) GetIdleTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.IdleTimeout) {
    return nil, false
	}
	return o.IdleTimeout, true
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) HasIdleTimeout() bool {
	if o != nil && !isNil(o.IdleTimeout) {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given int32 and assigns it to the IdleTimeout field.
func (o *UpdateNetworkWirelessSsidVpnRequestFailover) SetIdleTimeout(v int32) {
	o.IdleTimeout = &v
}

func (o UpdateNetworkWirelessSsidVpnRequestFailover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RequestIp) {
		toSerialize["requestIp"] = o.RequestIp
	}
	if !isNil(o.HeartbeatInterval) {
		toSerialize["heartbeatInterval"] = o.HeartbeatInterval
	}
	if !isNil(o.IdleTimeout) {
		toSerialize["idleTimeout"] = o.IdleTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidVpnRequestFailover struct {
	value *UpdateNetworkWirelessSsidVpnRequestFailover
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestFailover) Get() *UpdateNetworkWirelessSsidVpnRequestFailover {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestFailover) Set(val *UpdateNetworkWirelessSsidVpnRequestFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidVpnRequestFailover(val *UpdateNetworkWirelessSsidVpnRequestFailover) *NullableUpdateNetworkWirelessSsidVpnRequestFailover {
	return &NullableUpdateNetworkWirelessSsidVpnRequestFailover{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


