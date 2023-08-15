/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceVpnSiteToSiteVpn200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceVpnSiteToSiteVpn200Response{}

// GetNetworkApplianceVpnSiteToSiteVpn200Response struct for GetNetworkApplianceVpnSiteToSiteVpn200Response
type GetNetworkApplianceVpnSiteToSiteVpn200Response struct {
	// The site-to-site VPN mode.
	Mode *string `json:"mode,omitempty"`
	// The list of VPN hubs, in order of preference.
	Hubs []GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner `json:"hubs,omitempty"`
	// The list of subnets and their VPN presence.
	Subnets []GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner `json:"subnets,omitempty"`
}

// NewGetNetworkApplianceVpnSiteToSiteVpn200Response instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceVpnSiteToSiteVpn200Response() *GetNetworkApplianceVpnSiteToSiteVpn200Response {
	this := GetNetworkApplianceVpnSiteToSiteVpn200Response{}
	return &this
}

// NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseWithDefaults instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseWithDefaults() *GetNetworkApplianceVpnSiteToSiteVpn200Response {
	this := GetNetworkApplianceVpnSiteToSiteVpn200Response{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) SetMode(v string) {
	o.Mode = &v
}

// GetHubs returns the Hubs field value if set, zero value otherwise.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetHubs() []GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner {
	if o == nil || IsNil(o.Hubs) {
		var ret []GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner
		return ret
	}
	return o.Hubs
}

// GetHubsOk returns a tuple with the Hubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetHubsOk() ([]GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner, bool) {
	if o == nil || IsNil(o.Hubs) {
		return nil, false
	}
	return o.Hubs, true
}

// HasHubs returns a boolean if a field has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) HasHubs() bool {
	if o != nil && !IsNil(o.Hubs) {
		return true
	}

	return false
}

// SetHubs gets a reference to the given []GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner and assigns it to the Hubs field.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) SetHubs(v []GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) {
	o.Hubs = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetSubnets() []GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner {
	if o == nil || IsNil(o.Subnets) {
		var ret []GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetSubnetsOk() ([]GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner and assigns it to the Subnets field.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) SetSubnets(v []GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) {
	o.Subnets = v
}

func (o GetNetworkApplianceVpnSiteToSiteVpn200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceVpnSiteToSiteVpn200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Hubs) {
		toSerialize["hubs"] = o.Hubs
	}
	if !IsNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceVpnSiteToSiteVpn200Response struct {
	value *GetNetworkApplianceVpnSiteToSiteVpn200Response
	isSet bool
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200Response) Get() *GetNetworkApplianceVpnSiteToSiteVpn200Response {
	return v.value
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200Response) Set(val *GetNetworkApplianceVpnSiteToSiteVpn200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceVpnSiteToSiteVpn200Response(val *GetNetworkApplianceVpnSiteToSiteVpn200Response) *NullableGetNetworkApplianceVpnSiteToSiteVpn200Response {
	return &NullableGetNetworkApplianceVpnSiteToSiteVpn200Response{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


