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

// checks if the GetNetworkSwitchStackRoutingInterfaceDhcp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchStackRoutingInterfaceDhcp200Response{}

// GetNetworkSwitchStackRoutingInterfaceDhcp200Response struct for GetNetworkSwitchStackRoutingInterfaceDhcp200Response
type GetNetworkSwitchStackRoutingInterfaceDhcp200Response struct {
	// The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpMode *string `json:"dhcpMode,omitempty"`
	// The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	DnsNameserversOption *string `json:"dnsNameserversOption,omitempty"`
	// The DHCP name server IPs when DHCP name server option is 'custom'
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
	// Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	BootOptionsEnabled *bool `json:"bootOptionsEnabled,omitempty"`
	// The PXE boot server IP for the DHCP server running on the switch stack interface
	BootNextServer *string `json:"bootNextServer,omitempty"`
	// The PXE boot server file name for the DHCP server running on the switch stack interface
	BootFileName *string `json:"bootFileName,omitempty"`
	// Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpOptions []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner `json:"dhcpOptions,omitempty"`
	// Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIpRanges []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner `json:"reservedIpRanges,omitempty"`
	// Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	FixedIpAssignments []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner `json:"fixedIpAssignments,omitempty"`
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcp200Response instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchStackRoutingInterfaceDhcp200Response() *GetNetworkSwitchStackRoutingInterfaceDhcp200Response {
	this := GetNetworkSwitchStackRoutingInterfaceDhcp200Response{}
	return &this
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseWithDefaults instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseWithDefaults() *GetNetworkSwitchStackRoutingInterfaceDhcp200Response {
	this := GetNetworkSwitchStackRoutingInterfaceDhcp200Response{}
	return &this
}

// GetDhcpMode returns the DhcpMode field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpMode() string {
	if o == nil || IsNil(o.DhcpMode) {
		var ret string
		return ret
	}
	return *o.DhcpMode
}

// GetDhcpModeOk returns a tuple with the DhcpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpModeOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpMode) {
		return nil, false
	}
	return o.DhcpMode, true
}

// HasDhcpMode returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDhcpMode() bool {
	if o != nil && !IsNil(o.DhcpMode) {
		return true
	}

	return false
}

// SetDhcpMode gets a reference to the given string and assigns it to the DhcpMode field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDhcpMode(v string) {
	o.DhcpMode = &v
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpLeaseTime() string {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDhcpLeaseTime() bool {
	if o != nil && !IsNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameserversOption returns the DnsNameserversOption field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsNameserversOption() string {
	if o == nil || IsNil(o.DnsNameserversOption) {
		var ret string
		return ret
	}
	return *o.DnsNameserversOption
}

// GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsNameserversOptionOk() (*string, bool) {
	if o == nil || IsNil(o.DnsNameserversOption) {
		return nil, false
	}
	return o.DnsNameserversOption, true
}

// HasDnsNameserversOption returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDnsNameserversOption() bool {
	if o != nil && !IsNil(o.DnsNameserversOption) {
		return true
	}

	return false
}

// SetDnsNameserversOption gets a reference to the given string and assigns it to the DnsNameserversOption field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDnsNameserversOption(v string) {
	o.DnsNameserversOption = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsCustomNameservers() []string {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDnsCustomNameservers() bool {
	if o != nil && !IsNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

// GetBootOptionsEnabled returns the BootOptionsEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootOptionsEnabled() bool {
	if o == nil || IsNil(o.BootOptionsEnabled) {
		var ret bool
		return ret
	}
	return *o.BootOptionsEnabled
}

// GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootOptionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BootOptionsEnabled) {
		return nil, false
	}
	return o.BootOptionsEnabled, true
}

// HasBootOptionsEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasBootOptionsEnabled() bool {
	if o != nil && !IsNil(o.BootOptionsEnabled) {
		return true
	}

	return false
}

// SetBootOptionsEnabled gets a reference to the given bool and assigns it to the BootOptionsEnabled field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetBootOptionsEnabled(v bool) {
	o.BootOptionsEnabled = &v
}

// GetBootNextServer returns the BootNextServer field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootNextServer() string {
	if o == nil || IsNil(o.BootNextServer) {
		var ret string
		return ret
	}
	return *o.BootNextServer
}

// GetBootNextServerOk returns a tuple with the BootNextServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootNextServerOk() (*string, bool) {
	if o == nil || IsNil(o.BootNextServer) {
		return nil, false
	}
	return o.BootNextServer, true
}

// HasBootNextServer returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasBootNextServer() bool {
	if o != nil && !IsNil(o.BootNextServer) {
		return true
	}

	return false
}

// SetBootNextServer gets a reference to the given string and assigns it to the BootNextServer field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetBootNextServer(v string) {
	o.BootNextServer = &v
}

// GetBootFileName returns the BootFileName field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootFileName() string {
	if o == nil || IsNil(o.BootFileName) {
		var ret string
		return ret
	}
	return *o.BootFileName
}

// GetBootFileNameOk returns a tuple with the BootFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.BootFileName) {
		return nil, false
	}
	return o.BootFileName, true
}

// HasBootFileName returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasBootFileName() bool {
	if o != nil && !IsNil(o.BootFileName) {
		return true
	}

	return false
}

// SetBootFileName gets a reference to the given string and assigns it to the BootFileName field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetBootFileName(v string) {
	o.BootFileName = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpOptions() []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpOptionsOk() ([]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner and assigns it to the DhcpOptions field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDhcpOptions(v []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner) {
	o.DhcpOptions = v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetReservedIpRanges() []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner {
	if o == nil || IsNil(o.ReservedIpRanges) {
		var ret []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetReservedIpRangesOk() ([]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner, bool) {
	if o == nil || IsNil(o.ReservedIpRanges) {
		return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasReservedIpRanges() bool {
	if o != nil && !IsNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner and assigns it to the ReservedIpRanges field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetReservedIpRanges(v []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) {
	o.ReservedIpRanges = v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetFixedIpAssignments() []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner {
	if o == nil || IsNil(o.FixedIpAssignments) {
		var ret []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetFixedIpAssignmentsOk() ([]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner, bool) {
	if o == nil || IsNil(o.FixedIpAssignments) {
		return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasFixedIpAssignments() bool {
	if o != nil && !IsNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner and assigns it to the FixedIpAssignments field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetFixedIpAssignments(v []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) {
	o.FixedIpAssignments = v
}

func (o GetNetworkSwitchStackRoutingInterfaceDhcp200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchStackRoutingInterfaceDhcp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DhcpMode) {
		toSerialize["dhcpMode"] = o.DhcpMode
	}
	if !IsNil(o.DhcpLeaseTime) {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if !IsNil(o.DnsNameserversOption) {
		toSerialize["dnsNameserversOption"] = o.DnsNameserversOption
	}
	if !IsNil(o.DnsCustomNameservers) {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	if !IsNil(o.BootOptionsEnabled) {
		toSerialize["bootOptionsEnabled"] = o.BootOptionsEnabled
	}
	if !IsNil(o.BootNextServer) {
		toSerialize["bootNextServer"] = o.BootNextServer
	}
	if !IsNil(o.BootFileName) {
		toSerialize["bootFileName"] = o.BootFileName
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcpOptions"] = o.DhcpOptions
	}
	if !IsNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !IsNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response struct {
	value *GetNetworkSwitchStackRoutingInterfaceDhcp200Response
	isSet bool
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response) Get() *GetNetworkSwitchStackRoutingInterfaceDhcp200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response) Set(val *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response(val *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response {
	return &NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


