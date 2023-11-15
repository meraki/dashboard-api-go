/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkHealthAlerts200ResponseInnerScopePeersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkHealthAlerts200ResponseInnerScopePeersInner{}

// GetNetworkHealthAlerts200ResponseInnerScopePeersInner struct for GetNetworkHealthAlerts200ResponseInnerScopePeersInner
type GetNetworkHealthAlerts200ResponseInnerScopePeersInner struct {
	// URL to the peer
	Url *string `json:"url,omitempty"`
	Network *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork `json:"network,omitempty"`
}

// NewGetNetworkHealthAlerts200ResponseInnerScopePeersInner instantiates a new GetNetworkHealthAlerts200ResponseInnerScopePeersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkHealthAlerts200ResponseInnerScopePeersInner() *GetNetworkHealthAlerts200ResponseInnerScopePeersInner {
	this := GetNetworkHealthAlerts200ResponseInnerScopePeersInner{}
	return &this
}

// NewGetNetworkHealthAlerts200ResponseInnerScopePeersInnerWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInnerScopePeersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkHealthAlerts200ResponseInnerScopePeersInnerWithDefaults() *GetNetworkHealthAlerts200ResponseInnerScopePeersInner {
	this := GetNetworkHealthAlerts200ResponseInnerScopePeersInner{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) SetUrl(v string) {
	o.Url = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) GetNetwork() GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) GetNetworkOk() (*GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork and assigns it to the Network field.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) SetNetwork(v GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) {
	o.Network = &v
}

func (o GetNetworkHealthAlerts200ResponseInnerScopePeersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkHealthAlerts200ResponseInnerScopePeersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return toSerialize, nil
}

type NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner struct {
	value *GetNetworkHealthAlerts200ResponseInnerScopePeersInner
	isSet bool
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner) Get() *GetNetworkHealthAlerts200ResponseInnerScopePeersInner {
	return v.value
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner) Set(val *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner(val *GetNetworkHealthAlerts200ResponseInnerScopePeersInner) *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner {
	return &NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner{value: val, isSet: true}
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


