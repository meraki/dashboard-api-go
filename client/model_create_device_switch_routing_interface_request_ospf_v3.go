/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateDeviceSwitchRoutingInterfaceRequestOspfV3 The OSPFv3 routing settings of the interface.
type CreateDeviceSwitchRoutingInterfaceRequestOspfV3 struct {
	// The OSPFv3 area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPFv3 area. Defaults to 'disabled'.
	Area *string `json:"area,omitempty"`
	// The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	Cost *int32 `json:"cost,omitempty"`
	// When enabled, OSPFv3 will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3 instantiates a new CreateDeviceSwitchRoutingInterfaceRequestOspfV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3() *CreateDeviceSwitchRoutingInterfaceRequestOspfV3 {
	this := CreateDeviceSwitchRoutingInterfaceRequestOspfV3{}
	return &this
}

// NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3WithDefaults instantiates a new CreateDeviceSwitchRoutingInterfaceRequestOspfV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3WithDefaults() *CreateDeviceSwitchRoutingInterfaceRequestOspfV3 {
	this := CreateDeviceSwitchRoutingInterfaceRequestOspfV3{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetArea() string {
	if o == nil || isNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetAreaOk() (*string, bool) {
	if o == nil || isNil(o.Area) {
    return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) HasArea() bool {
	if o != nil && !isNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetCost() int32 {
	if o == nil || isNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetCostOk() (*int32, bool) {
	if o == nil || isNil(o.Cost) {
    return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) HasCost() bool {
	if o != nil && !isNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetIsPassiveEnabled() bool {
	if o == nil || isNil(o.IsPassiveEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsPassiveEnabled) {
    return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) HasIsPassiveEnabled() bool {
	if o != nil && !isNil(o.IsPassiveEnabled) {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o CreateDeviceSwitchRoutingInterfaceRequestOspfV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !isNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !isNil(o.IsPassiveEnabled) {
		toSerialize["isPassiveEnabled"] = o.IsPassiveEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3 struct {
	value *CreateDeviceSwitchRoutingInterfaceRequestOspfV3
	isSet bool
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3) Get() *CreateDeviceSwitchRoutingInterfaceRequestOspfV3 {
	return v.value
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3) Set(val *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3(val *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3 {
	return &NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3{value: val, isSet: true}
}

func (v NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceSwitchRoutingInterfaceRequestOspfV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


