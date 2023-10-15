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

// checks if the GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings{}

// GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings Per-SSID radio settings by number.
type GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings struct {
	Var1 *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings1 `json:"1,omitempty"`
	Var2 *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings2 `json:"2,omitempty"`
	Var3 *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings3 `json:"3,omitempty"`
	Var4 *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 `json:"4,omitempty"`
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings{}
	return &this
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettingsWithDefaults instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettingsWithDefaults() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings{}
	return &this
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar1() GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings1 {
	if o == nil || IsNil(o.Var1) {
		var ret GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar1Ok() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings1, bool) {
	if o == nil || IsNil(o.Var1) {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) HasVar1() bool {
	if o != nil && !IsNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings1 and assigns it to the Var1 field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) SetVar1(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar2() GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings2 {
	if o == nil || IsNil(o.Var2) {
		var ret GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar2Ok() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings2, bool) {
	if o == nil || IsNil(o.Var2) {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) HasVar2() bool {
	if o != nil && !IsNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings2 and assigns it to the Var2 field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) SetVar2(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar3() GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings3 {
	if o == nil || IsNil(o.Var3) {
		var ret GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar3Ok() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings3, bool) {
	if o == nil || IsNil(o.Var3) {
		return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) HasVar3() bool {
	if o != nil && !IsNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings3 and assigns it to the Var3 field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) SetVar3(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar4() GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 {
	if o == nil || IsNil(o.Var4) {
		var ret GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) GetVar4Ok() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4, bool) {
	if o == nil || IsNil(o.Var4) {
		return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) HasVar4() bool {
	if o != nil && !IsNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4 and assigns it to the Var4 field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) SetVar4(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings4) {
	o.Var4 = &v
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var1) {
		toSerialize["1"] = o.Var1
	}
	if !IsNil(o.Var2) {
		toSerialize["2"] = o.Var2
	}
	if !IsNil(o.Var3) {
		toSerialize["3"] = o.Var3
	}
	if !IsNil(o.Var4) {
		toSerialize["4"] = o.Var4
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings struct {
	value *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings
	isSet bool
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) Get() *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings {
	return v.value
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) Set(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings(val *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings {
	return &NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


