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

// checks if the CreateDeviceSwitchRoutingInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceSwitchRoutingInterfaceRequest{}

// CreateDeviceSwitchRoutingInterfaceRequest struct for CreateDeviceSwitchRoutingInterfaceRequest
type CreateDeviceSwitchRoutingInterfaceRequest struct {
	// A friendly name or description for the interface or VLAN.
	Name *string `json:"name,omitempty"`
	// The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet *string `json:"subnet,omitempty"`
	// The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same         as the switch's management IP.
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// Enable multicast support if, multicast routing between VLANs is required. Options are:         'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// The VLAN this routed interface is on. VLAN must be between 1 and 4094.
	VlanId *int32 `json:"vlanId,omitempty"`
	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a routed interface. Required if this is the first IPv4 interface.
	DefaultGateway *string `json:"defaultGateway,omitempty"`
	OspfSettings *CreateDeviceSwitchRoutingInterfaceRequestOspfSettings `json:"ospfSettings,omitempty"`
	OspfV3 *CreateDeviceSwitchRoutingInterfaceRequestOspfV3 `json:"ospfV3,omitempty"`
	Ipv6 *CreateDeviceSwitchRoutingInterfaceRequestIpv6 `json:"ipv6,omitempty"`
}

// NewCreateDeviceSwitchRoutingInterfaceRequest instantiates a new CreateDeviceSwitchRoutingInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceSwitchRoutingInterfaceRequest() *CreateDeviceSwitchRoutingInterfaceRequest {
	this := CreateDeviceSwitchRoutingInterfaceRequest{}
	return &this
}

// NewCreateDeviceSwitchRoutingInterfaceRequestWithDefaults instantiates a new CreateDeviceSwitchRoutingInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceSwitchRoutingInterfaceRequestWithDefaults() *CreateDeviceSwitchRoutingInterfaceRequest {
	this := CreateDeviceSwitchRoutingInterfaceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetInterfaceIp() string {
	if o == nil || IsNil(o.InterfaceIp) {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetInterfaceIpOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceIp) {
		return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasInterfaceIp() bool {
	if o != nil && !IsNil(o.InterfaceIp) {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetMulticastRouting() string {
	if o == nil || IsNil(o.MulticastRouting) {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.MulticastRouting) {
		return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasMulticastRouting() bool {
	if o != nil && !IsNil(o.MulticastRouting) {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetDefaultGateway() string {
	if o == nil || IsNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultGateway) {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasDefaultGateway() bool {
	if o != nil && !IsNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfSettings() CreateDeviceSwitchRoutingInterfaceRequestOspfSettings {
	if o == nil || IsNil(o.OspfSettings) {
		var ret CreateDeviceSwitchRoutingInterfaceRequestOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfSettingsOk() (*CreateDeviceSwitchRoutingInterfaceRequestOspfSettings, bool) {
	if o == nil || IsNil(o.OspfSettings) {
		return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasOspfSettings() bool {
	if o != nil && !IsNil(o.OspfSettings) {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given CreateDeviceSwitchRoutingInterfaceRequestOspfSettings and assigns it to the OspfSettings field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetOspfSettings(v CreateDeviceSwitchRoutingInterfaceRequestOspfSettings) {
	o.OspfSettings = &v
}

// GetOspfV3 returns the OspfV3 field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfV3() CreateDeviceSwitchRoutingInterfaceRequestOspfV3 {
	if o == nil || IsNil(o.OspfV3) {
		var ret CreateDeviceSwitchRoutingInterfaceRequestOspfV3
		return ret
	}
	return *o.OspfV3
}

// GetOspfV3Ok returns a tuple with the OspfV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfV3Ok() (*CreateDeviceSwitchRoutingInterfaceRequestOspfV3, bool) {
	if o == nil || IsNil(o.OspfV3) {
		return nil, false
	}
	return o.OspfV3, true
}

// HasOspfV3 returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasOspfV3() bool {
	if o != nil && !IsNil(o.OspfV3) {
		return true
	}

	return false
}

// SetOspfV3 gets a reference to the given CreateDeviceSwitchRoutingInterfaceRequestOspfV3 and assigns it to the OspfV3 field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetOspfV3(v CreateDeviceSwitchRoutingInterfaceRequestOspfV3) {
	o.OspfV3 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetIpv6() CreateDeviceSwitchRoutingInterfaceRequestIpv6 {
	if o == nil || IsNil(o.Ipv6) {
		var ret CreateDeviceSwitchRoutingInterfaceRequestIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetIpv6Ok() (*CreateDeviceSwitchRoutingInterfaceRequestIpv6, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given CreateDeviceSwitchRoutingInterfaceRequestIpv6 and assigns it to the Ipv6 field.
func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetIpv6(v CreateDeviceSwitchRoutingInterfaceRequestIpv6) {
	o.Ipv6 = &v
}

func (o CreateDeviceSwitchRoutingInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceSwitchRoutingInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.InterfaceIp) {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if !IsNil(o.MulticastRouting) {
		toSerialize["multicastRouting"] = o.MulticastRouting
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.DefaultGateway) {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if !IsNil(o.OspfSettings) {
		toSerialize["ospfSettings"] = o.OspfSettings
	}
	if !IsNil(o.OspfV3) {
		toSerialize["ospfV3"] = o.OspfV3
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return toSerialize, nil
}

type NullableCreateDeviceSwitchRoutingInterfaceRequest struct {
	value *CreateDeviceSwitchRoutingInterfaceRequest
	isSet bool
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequest) Get() *CreateDeviceSwitchRoutingInterfaceRequest {
	return v.value
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequest) Set(val *CreateDeviceSwitchRoutingInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceSwitchRoutingInterfaceRequest(val *CreateDeviceSwitchRoutingInterfaceRequest) *NullableCreateDeviceSwitchRoutingInterfaceRequest {
	return &NullableCreateDeviceSwitchRoutingInterfaceRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


