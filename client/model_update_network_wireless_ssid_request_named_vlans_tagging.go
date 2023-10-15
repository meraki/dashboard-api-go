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

// checks if the UpdateNetworkWirelessSsidRequestNamedVlansTagging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestNamedVlansTagging{}

// UpdateNetworkWirelessSsidRequestNamedVlansTagging VLAN tagging settings. This param is only valid when ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'.
type UpdateNetworkWirelessSsidRequestNamedVlansTagging struct {
	// Whether or not traffic should be directed to use specific VLAN names.
	Enabled *bool `json:"enabled,omitempty"`
	// The default VLAN name used to tag traffic in the absence of a matching AP tag.
	DefaultVlanName *string `json:"defaultVlanName,omitempty"`
	// The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag.
	ByApTags []UpdateNetworkWirelessSsidRequestNamedVlansTaggingByApTagsInner `json:"byApTags,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestNamedVlansTagging instantiates a new UpdateNetworkWirelessSsidRequestNamedVlansTagging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestNamedVlansTagging() *UpdateNetworkWirelessSsidRequestNamedVlansTagging {
	this := UpdateNetworkWirelessSsidRequestNamedVlansTagging{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestNamedVlansTaggingWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestNamedVlansTagging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestNamedVlansTaggingWithDefaults() *UpdateNetworkWirelessSsidRequestNamedVlansTagging {
	this := UpdateNetworkWirelessSsidRequestNamedVlansTagging{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefaultVlanName returns the DefaultVlanName field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) GetDefaultVlanName() string {
	if o == nil || IsNil(o.DefaultVlanName) {
		var ret string
		return ret
	}
	return *o.DefaultVlanName
}

// GetDefaultVlanNameOk returns a tuple with the DefaultVlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) GetDefaultVlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultVlanName) {
		return nil, false
	}
	return o.DefaultVlanName, true
}

// HasDefaultVlanName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) HasDefaultVlanName() bool {
	if o != nil && !IsNil(o.DefaultVlanName) {
		return true
	}

	return false
}

// SetDefaultVlanName gets a reference to the given string and assigns it to the DefaultVlanName field.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) SetDefaultVlanName(v string) {
	o.DefaultVlanName = &v
}

// GetByApTags returns the ByApTags field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) GetByApTags() []UpdateNetworkWirelessSsidRequestNamedVlansTaggingByApTagsInner {
	if o == nil || IsNil(o.ByApTags) {
		var ret []UpdateNetworkWirelessSsidRequestNamedVlansTaggingByApTagsInner
		return ret
	}
	return o.ByApTags
}

// GetByApTagsOk returns a tuple with the ByApTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) GetByApTagsOk() ([]UpdateNetworkWirelessSsidRequestNamedVlansTaggingByApTagsInner, bool) {
	if o == nil || IsNil(o.ByApTags) {
		return nil, false
	}
	return o.ByApTags, true
}

// HasByApTags returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) HasByApTags() bool {
	if o != nil && !IsNil(o.ByApTags) {
		return true
	}

	return false
}

// SetByApTags gets a reference to the given []UpdateNetworkWirelessSsidRequestNamedVlansTaggingByApTagsInner and assigns it to the ByApTags field.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansTagging) SetByApTags(v []UpdateNetworkWirelessSsidRequestNamedVlansTaggingByApTagsInner) {
	o.ByApTags = v
}

func (o UpdateNetworkWirelessSsidRequestNamedVlansTagging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestNamedVlansTagging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DefaultVlanName) {
		toSerialize["defaultVlanName"] = o.DefaultVlanName
	}
	if !IsNil(o.ByApTags) {
		toSerialize["byApTags"] = o.ByApTags
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging struct {
	value *UpdateNetworkWirelessSsidRequestNamedVlansTagging
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging) Get() *UpdateNetworkWirelessSsidRequestNamedVlansTagging {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging) Set(val *UpdateNetworkWirelessSsidRequestNamedVlansTagging) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestNamedVlansTagging(val *UpdateNetworkWirelessSsidRequestNamedVlansTagging) *NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging {
	return &NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestNamedVlansTagging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


