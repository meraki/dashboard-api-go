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

// checks if the UpdateNetworkApplianceVpnBgpRequestNeighborsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVpnBgpRequestNeighborsInner{}

// UpdateNetworkApplianceVpnBgpRequestNeighborsInner struct for UpdateNetworkApplianceVpnBgpRequestNeighborsInner
type UpdateNetworkApplianceVpnBgpRequestNeighborsInner struct {
	// The IPv4 address of the neighbor
	Ip *string `json:"ip,omitempty"`
	Ipv6 *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 `json:"ipv6,omitempty"`
	// Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295.
	RemoteAsNumber int32 `json:"remoteAsNumber"`
	// The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 4294967295. When absent, it defaults to 0.
	ReceiveLimit *int32 `json:"receiveLimit,omitempty"`
	// When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false.
	AllowTransit *bool `json:"allowTransit,omitempty"`
	// The EBGP hold timer in seconds for each neighbor. The EBGP hold timer must be an integer between 12 and 240.
	EbgpHoldTimer int32 `json:"ebgpHoldTimer"`
	// Configure this if the neighbor is not adjacent. The EBGP multi-hop must be an integer between 1 and 255.
	EbgpMultihop int32 `json:"ebgpMultihop"`
	// The output interface for peering with the remote BGP peer. Valid values are: 'wired0', 'wired1' or 'vlan{VLAN ID}'(e.g. 'vlan123').
	SourceInterface *string `json:"sourceInterface,omitempty"`
	// The IPv4 address of the remote BGP peer that will establish a TCP session with the local MX.
	NextHopIp *string `json:"nextHopIp,omitempty"`
	TtlSecurity *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity `json:"ttlSecurity,omitempty"`
	Authentication *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication `json:"authentication,omitempty"`
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInner instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInner(remoteAsNumber int32, ebgpHoldTimer int32, ebgpMultihop int32) *UpdateNetworkApplianceVpnBgpRequestNeighborsInner {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInner{}
	this.RemoteAsNumber = remoteAsNumber
	this.EbgpHoldTimer = ebgpHoldTimer
	this.EbgpMultihop = ebgpMultihop
	return &this
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerWithDefaults instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerWithDefaults() *UpdateNetworkApplianceVpnBgpRequestNeighborsInner {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInner{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetIp(v string) {
	o.Ip = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIpv6() UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 {
	if o == nil || IsNil(o.Ipv6) {
		var ret UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIpv6Ok() (*UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 and assigns it to the Ipv6 field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetIpv6(v UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) {
	o.Ipv6 = &v
}

// GetRemoteAsNumber returns the RemoteAsNumber field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetRemoteAsNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemoteAsNumber
}

// GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetRemoteAsNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAsNumber, true
}

// SetRemoteAsNumber sets field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetRemoteAsNumber(v int32) {
	o.RemoteAsNumber = v
}

// GetReceiveLimit returns the ReceiveLimit field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetReceiveLimit() int32 {
	if o == nil || IsNil(o.ReceiveLimit) {
		var ret int32
		return ret
	}
	return *o.ReceiveLimit
}

// GetReceiveLimitOk returns a tuple with the ReceiveLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetReceiveLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ReceiveLimit) {
		return nil, false
	}
	return o.ReceiveLimit, true
}

// HasReceiveLimit returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasReceiveLimit() bool {
	if o != nil && !IsNil(o.ReceiveLimit) {
		return true
	}

	return false
}

// SetReceiveLimit gets a reference to the given int32 and assigns it to the ReceiveLimit field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetReceiveLimit(v int32) {
	o.ReceiveLimit = &v
}

// GetAllowTransit returns the AllowTransit field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetAllowTransit() bool {
	if o == nil || IsNil(o.AllowTransit) {
		var ret bool
		return ret
	}
	return *o.AllowTransit
}

// GetAllowTransitOk returns a tuple with the AllowTransit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetAllowTransitOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowTransit) {
		return nil, false
	}
	return o.AllowTransit, true
}

// HasAllowTransit returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasAllowTransit() bool {
	if o != nil && !IsNil(o.AllowTransit) {
		return true
	}

	return false
}

// SetAllowTransit gets a reference to the given bool and assigns it to the AllowTransit field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetAllowTransit(v bool) {
	o.AllowTransit = &v
}

// GetEbgpHoldTimer returns the EbgpHoldTimer field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpHoldTimer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EbgpHoldTimer
}

// GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpHoldTimerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EbgpHoldTimer, true
}

// SetEbgpHoldTimer sets field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetEbgpHoldTimer(v int32) {
	o.EbgpHoldTimer = v
}

// GetEbgpMultihop returns the EbgpMultihop field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpMultihop() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EbgpMultihop
}

// GetEbgpMultihopOk returns a tuple with the EbgpMultihop field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpMultihopOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EbgpMultihop, true
}

// SetEbgpMultihop sets field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetEbgpMultihop(v int32) {
	o.EbgpMultihop = v
}

// GetSourceInterface returns the SourceInterface field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetSourceInterface() string {
	if o == nil || IsNil(o.SourceInterface) {
		var ret string
		return ret
	}
	return *o.SourceInterface
}

// GetSourceInterfaceOk returns a tuple with the SourceInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetSourceInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.SourceInterface) {
		return nil, false
	}
	return o.SourceInterface, true
}

// HasSourceInterface returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasSourceInterface() bool {
	if o != nil && !IsNil(o.SourceInterface) {
		return true
	}

	return false
}

// SetSourceInterface gets a reference to the given string and assigns it to the SourceInterface field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetSourceInterface(v string) {
	o.SourceInterface = &v
}

// GetNextHopIp returns the NextHopIp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetNextHopIp() string {
	if o == nil || IsNil(o.NextHopIp) {
		var ret string
		return ret
	}
	return *o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetNextHopIpOk() (*string, bool) {
	if o == nil || IsNil(o.NextHopIp) {
		return nil, false
	}
	return o.NextHopIp, true
}

// HasNextHopIp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasNextHopIp() bool {
	if o != nil && !IsNil(o.NextHopIp) {
		return true
	}

	return false
}

// SetNextHopIp gets a reference to the given string and assigns it to the NextHopIp field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetNextHopIp(v string) {
	o.NextHopIp = &v
}

// GetTtlSecurity returns the TtlSecurity field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetTtlSecurity() UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity {
	if o == nil || IsNil(o.TtlSecurity) {
		var ret UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity
		return ret
	}
	return *o.TtlSecurity
}

// GetTtlSecurityOk returns a tuple with the TtlSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetTtlSecurityOk() (*UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity, bool) {
	if o == nil || IsNil(o.TtlSecurity) {
		return nil, false
	}
	return o.TtlSecurity, true
}

// HasTtlSecurity returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasTtlSecurity() bool {
	if o != nil && !IsNil(o.TtlSecurity) {
		return true
	}

	return false
}

// SetTtlSecurity gets a reference to the given UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity and assigns it to the TtlSecurity field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetTtlSecurity(v UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) {
	o.TtlSecurity = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetAuthentication() UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetAuthenticationOk() (*UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication and assigns it to the Authentication field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetAuthentication(v UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) {
	o.Authentication = &v
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	toSerialize["remoteAsNumber"] = o.RemoteAsNumber
	if !IsNil(o.ReceiveLimit) {
		toSerialize["receiveLimit"] = o.ReceiveLimit
	}
	if !IsNil(o.AllowTransit) {
		toSerialize["allowTransit"] = o.AllowTransit
	}
	toSerialize["ebgpHoldTimer"] = o.EbgpHoldTimer
	toSerialize["ebgpMultihop"] = o.EbgpMultihop
	if !IsNil(o.SourceInterface) {
		toSerialize["sourceInterface"] = o.SourceInterface
	}
	if !IsNil(o.NextHopIp) {
		toSerialize["nextHopIp"] = o.NextHopIp
	}
	if !IsNil(o.TtlSecurity) {
		toSerialize["ttlSecurity"] = o.TtlSecurity
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner struct {
	value *UpdateNetworkApplianceVpnBgpRequestNeighborsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner) Get() *UpdateNetworkApplianceVpnBgpRequestNeighborsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner) Set(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner {
	return &NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


