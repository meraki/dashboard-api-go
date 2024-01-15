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

// checks if the GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging{}

// GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging VLAN tagging settings.
type GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging struct {
	// Whether VLAN tagging is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the VLAN to use for VLAN tagging.
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging{}
	return &this
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTaggingWithDefaults instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTaggingWithDefaults() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return toSerialize, nil
}

type NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging struct {
	value *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging
	isSet bool
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) Get() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging {
	return v.value
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) Set(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging {
	return &NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


