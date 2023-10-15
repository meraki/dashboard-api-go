/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkClient200ResponseClientVpnConnectionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkClient200ResponseClientVpnConnectionsInner{}

// GetNetworkClient200ResponseClientVpnConnectionsInner struct for GetNetworkClient200ResponseClientVpnConnectionsInner
type GetNetworkClient200ResponseClientVpnConnectionsInner struct {
	// The IP address of the VPN the client last connected to
	RemoteIp *string `json:"remoteIp,omitempty"`
	// The time the client last connected to the VPN
	ConnectedAt *int32 `json:"connectedAt,omitempty"`
	// The time the client last disconnectd from the VPN
	DisconnectedAt *int32 `json:"disconnectedAt,omitempty"`
}

// NewGetNetworkClient200ResponseClientVpnConnectionsInner instantiates a new GetNetworkClient200ResponseClientVpnConnectionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkClient200ResponseClientVpnConnectionsInner() *GetNetworkClient200ResponseClientVpnConnectionsInner {
	this := GetNetworkClient200ResponseClientVpnConnectionsInner{}
	return &this
}

// NewGetNetworkClient200ResponseClientVpnConnectionsInnerWithDefaults instantiates a new GetNetworkClient200ResponseClientVpnConnectionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkClient200ResponseClientVpnConnectionsInnerWithDefaults() *GetNetworkClient200ResponseClientVpnConnectionsInner {
	this := GetNetworkClient200ResponseClientVpnConnectionsInner{}
	return &this
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) GetRemoteIp() string {
	if o == nil || IsNil(o.RemoteIp) {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) GetRemoteIpOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteIp) {
		return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) HasRemoteIp() bool {
	if o != nil && !IsNil(o.RemoteIp) {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

// GetConnectedAt returns the ConnectedAt field value if set, zero value otherwise.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) GetConnectedAt() int32 {
	if o == nil || IsNil(o.ConnectedAt) {
		var ret int32
		return ret
	}
	return *o.ConnectedAt
}

// GetConnectedAtOk returns a tuple with the ConnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) GetConnectedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectedAt) {
		return nil, false
	}
	return o.ConnectedAt, true
}

// HasConnectedAt returns a boolean if a field has been set.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) HasConnectedAt() bool {
	if o != nil && !IsNil(o.ConnectedAt) {
		return true
	}

	return false
}

// SetConnectedAt gets a reference to the given int32 and assigns it to the ConnectedAt field.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) SetConnectedAt(v int32) {
	o.ConnectedAt = &v
}

// GetDisconnectedAt returns the DisconnectedAt field value if set, zero value otherwise.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) GetDisconnectedAt() int32 {
	if o == nil || IsNil(o.DisconnectedAt) {
		var ret int32
		return ret
	}
	return *o.DisconnectedAt
}

// GetDisconnectedAtOk returns a tuple with the DisconnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) GetDisconnectedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.DisconnectedAt) {
		return nil, false
	}
	return o.DisconnectedAt, true
}

// HasDisconnectedAt returns a boolean if a field has been set.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) HasDisconnectedAt() bool {
	if o != nil && !IsNil(o.DisconnectedAt) {
		return true
	}

	return false
}

// SetDisconnectedAt gets a reference to the given int32 and assigns it to the DisconnectedAt field.
func (o *GetNetworkClient200ResponseClientVpnConnectionsInner) SetDisconnectedAt(v int32) {
	o.DisconnectedAt = &v
}

func (o GetNetworkClient200ResponseClientVpnConnectionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkClient200ResponseClientVpnConnectionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemoteIp) {
		toSerialize["remoteIp"] = o.RemoteIp
	}
	if !IsNil(o.ConnectedAt) {
		toSerialize["connectedAt"] = o.ConnectedAt
	}
	if !IsNil(o.DisconnectedAt) {
		toSerialize["disconnectedAt"] = o.DisconnectedAt
	}
	return toSerialize, nil
}

type NullableGetNetworkClient200ResponseClientVpnConnectionsInner struct {
	value *GetNetworkClient200ResponseClientVpnConnectionsInner
	isSet bool
}

func (v NullableGetNetworkClient200ResponseClientVpnConnectionsInner) Get() *GetNetworkClient200ResponseClientVpnConnectionsInner {
	return v.value
}

func (v *NullableGetNetworkClient200ResponseClientVpnConnectionsInner) Set(val *GetNetworkClient200ResponseClientVpnConnectionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkClient200ResponseClientVpnConnectionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkClient200ResponseClientVpnConnectionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkClient200ResponseClientVpnConnectionsInner(val *GetNetworkClient200ResponseClientVpnConnectionsInner) *NullableGetNetworkClient200ResponseClientVpnConnectionsInner {
	return &NullableGetNetworkClient200ResponseClientVpnConnectionsInner{value: val, isSet: true}
}

func (v NullableGetNetworkClient200ResponseClientVpnConnectionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkClient200ResponseClientVpnConnectionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


