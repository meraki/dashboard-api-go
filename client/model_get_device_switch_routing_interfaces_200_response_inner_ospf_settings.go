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

// checks if the GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings{}

// GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings IPv4 OSPF Settings
type GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings struct {
	// Area id
	Area *string `json:"area,omitempty"`
	// OSPF Cost
	Cost *int32 `json:"cost,omitempty"`
	// Disable sending Hello packets on this interface's IPv4 area
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings{}
	return &this
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettingsWithDefaults instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettingsWithDefaults() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetArea() string {
	if o == nil || IsNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetAreaOk() (*string, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetCost() int32 {
	if o == nil || IsNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetCostOk() (*int32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetIsPassiveEnabled() bool {
	if o == nil || IsNil(o.IsPassiveEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPassiveEnabled) {
		return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) HasIsPassiveEnabled() bool {
	if o != nil && !IsNil(o.IsPassiveEnabled) {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.IsPassiveEnabled) {
		toSerialize["isPassiveEnabled"] = o.IsPassiveEnabled
	}
	return toSerialize, nil
}

type NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings struct {
	value *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings
	isSet bool
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) Get() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings {
	return v.value
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) Set(val *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings(val *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings {
	return &NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


