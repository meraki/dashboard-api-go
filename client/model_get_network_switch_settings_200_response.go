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

// checks if the GetNetworkSwitchSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchSettings200Response{}

// GetNetworkSwitchSettings200Response struct for GetNetworkSwitchSettings200Response
type GetNetworkSwitchSettings200Response struct {
	// Management VLAN
	Vlan *int32 `json:"vlan,omitempty"`
	// The use Combined Power as the default behavior of secondary power supplies on supported devices.
	UseCombinedPower *bool `json:"useCombinedPower,omitempty"`
	// Exceptions on a per switch basis to \"useCombinedPower\"
	PowerExceptions []GetNetworkSwitchSettings200ResponsePowerExceptionsInner `json:"powerExceptions,omitempty"`
	UplinkClientSampling *GetNetworkSwitchSettings200ResponseUplinkClientSampling `json:"uplinkClientSampling,omitempty"`
	MacBlocklist *GetNetworkSwitchSettings200ResponseMacBlocklist `json:"macBlocklist,omitempty"`
}

// NewGetNetworkSwitchSettings200Response instantiates a new GetNetworkSwitchSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchSettings200Response() *GetNetworkSwitchSettings200Response {
	this := GetNetworkSwitchSettings200Response{}
	return &this
}

// NewGetNetworkSwitchSettings200ResponseWithDefaults instantiates a new GetNetworkSwitchSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchSettings200ResponseWithDefaults() *GetNetworkSwitchSettings200Response {
	this := GetNetworkSwitchSettings200Response{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200Response) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200Response) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200Response) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetNetworkSwitchSettings200Response) SetVlan(v int32) {
	o.Vlan = &v
}

// GetUseCombinedPower returns the UseCombinedPower field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200Response) GetUseCombinedPower() bool {
	if o == nil || IsNil(o.UseCombinedPower) {
		var ret bool
		return ret
	}
	return *o.UseCombinedPower
}

// GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200Response) GetUseCombinedPowerOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCombinedPower) {
		return nil, false
	}
	return o.UseCombinedPower, true
}

// HasUseCombinedPower returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200Response) HasUseCombinedPower() bool {
	if o != nil && !IsNil(o.UseCombinedPower) {
		return true
	}

	return false
}

// SetUseCombinedPower gets a reference to the given bool and assigns it to the UseCombinedPower field.
func (o *GetNetworkSwitchSettings200Response) SetUseCombinedPower(v bool) {
	o.UseCombinedPower = &v
}

// GetPowerExceptions returns the PowerExceptions field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200Response) GetPowerExceptions() []GetNetworkSwitchSettings200ResponsePowerExceptionsInner {
	if o == nil || IsNil(o.PowerExceptions) {
		var ret []GetNetworkSwitchSettings200ResponsePowerExceptionsInner
		return ret
	}
	return o.PowerExceptions
}

// GetPowerExceptionsOk returns a tuple with the PowerExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200Response) GetPowerExceptionsOk() ([]GetNetworkSwitchSettings200ResponsePowerExceptionsInner, bool) {
	if o == nil || IsNil(o.PowerExceptions) {
		return nil, false
	}
	return o.PowerExceptions, true
}

// HasPowerExceptions returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200Response) HasPowerExceptions() bool {
	if o != nil && !IsNil(o.PowerExceptions) {
		return true
	}

	return false
}

// SetPowerExceptions gets a reference to the given []GetNetworkSwitchSettings200ResponsePowerExceptionsInner and assigns it to the PowerExceptions field.
func (o *GetNetworkSwitchSettings200Response) SetPowerExceptions(v []GetNetworkSwitchSettings200ResponsePowerExceptionsInner) {
	o.PowerExceptions = v
}

// GetUplinkClientSampling returns the UplinkClientSampling field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200Response) GetUplinkClientSampling() GetNetworkSwitchSettings200ResponseUplinkClientSampling {
	if o == nil || IsNil(o.UplinkClientSampling) {
		var ret GetNetworkSwitchSettings200ResponseUplinkClientSampling
		return ret
	}
	return *o.UplinkClientSampling
}

// GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200Response) GetUplinkClientSamplingOk() (*GetNetworkSwitchSettings200ResponseUplinkClientSampling, bool) {
	if o == nil || IsNil(o.UplinkClientSampling) {
		return nil, false
	}
	return o.UplinkClientSampling, true
}

// HasUplinkClientSampling returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200Response) HasUplinkClientSampling() bool {
	if o != nil && !IsNil(o.UplinkClientSampling) {
		return true
	}

	return false
}

// SetUplinkClientSampling gets a reference to the given GetNetworkSwitchSettings200ResponseUplinkClientSampling and assigns it to the UplinkClientSampling field.
func (o *GetNetworkSwitchSettings200Response) SetUplinkClientSampling(v GetNetworkSwitchSettings200ResponseUplinkClientSampling) {
	o.UplinkClientSampling = &v
}

// GetMacBlocklist returns the MacBlocklist field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200Response) GetMacBlocklist() GetNetworkSwitchSettings200ResponseMacBlocklist {
	if o == nil || IsNil(o.MacBlocklist) {
		var ret GetNetworkSwitchSettings200ResponseMacBlocklist
		return ret
	}
	return *o.MacBlocklist
}

// GetMacBlocklistOk returns a tuple with the MacBlocklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200Response) GetMacBlocklistOk() (*GetNetworkSwitchSettings200ResponseMacBlocklist, bool) {
	if o == nil || IsNil(o.MacBlocklist) {
		return nil, false
	}
	return o.MacBlocklist, true
}

// HasMacBlocklist returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200Response) HasMacBlocklist() bool {
	if o != nil && !IsNil(o.MacBlocklist) {
		return true
	}

	return false
}

// SetMacBlocklist gets a reference to the given GetNetworkSwitchSettings200ResponseMacBlocklist and assigns it to the MacBlocklist field.
func (o *GetNetworkSwitchSettings200Response) SetMacBlocklist(v GetNetworkSwitchSettings200ResponseMacBlocklist) {
	o.MacBlocklist = &v
}

func (o GetNetworkSwitchSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.UseCombinedPower) {
		toSerialize["useCombinedPower"] = o.UseCombinedPower
	}
	if !IsNil(o.PowerExceptions) {
		toSerialize["powerExceptions"] = o.PowerExceptions
	}
	if !IsNil(o.UplinkClientSampling) {
		toSerialize["uplinkClientSampling"] = o.UplinkClientSampling
	}
	if !IsNil(o.MacBlocklist) {
		toSerialize["macBlocklist"] = o.MacBlocklist
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchSettings200Response struct {
	value *GetNetworkSwitchSettings200Response
	isSet bool
}

func (v NullableGetNetworkSwitchSettings200Response) Get() *GetNetworkSwitchSettings200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchSettings200Response) Set(val *GetNetworkSwitchSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchSettings200Response(val *GetNetworkSwitchSettings200Response) *NullableGetNetworkSwitchSettings200Response {
	return &NullableGetNetworkSwitchSettings200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


