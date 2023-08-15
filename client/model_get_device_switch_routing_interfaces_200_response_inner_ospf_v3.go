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

// checks if the GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3{}

// GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 IPv6 OSPF Settings
type GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 struct {
	// Area id
	Area *string `json:"area,omitempty"`
	// OSPF Cost
	Cost *int32 `json:"cost,omitempty"`
	// Disable sending Hello packets on this interface's IPv6 area
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3{}
	return &this
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3WithDefaults instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3WithDefaults() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) GetArea() string {
	if o == nil || IsNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) GetAreaOk() (*string, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) GetCost() int32 {
	if o == nil || IsNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) GetCostOk() (*int32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) GetIsPassiveEnabled() bool {
	if o == nil || IsNil(o.IsPassiveEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPassiveEnabled) {
		return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) HasIsPassiveEnabled() bool {
	if o != nil && !IsNil(o.IsPassiveEnabled) {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) ToMap() (map[string]interface{}, error) {
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

type NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 struct {
	value *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3
	isSet bool
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) Get() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 {
	return v.value
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) Set(val *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3(val *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 {
	return &NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


