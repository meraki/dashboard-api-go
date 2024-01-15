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

// checks if the CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst{}

// CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst Version information for the switch network being upgraded
type CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst struct {
	NextUpgrade *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst{}
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystWithDefaults() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) GetNextUpgrade() CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystNextUpgrade {
	if o == nil || IsNil(o.NextUpgrade) {
		var ret CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) GetNextUpgradeOk() (*CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystNextUpgrade, bool) {
	if o == nil || IsNil(o.NextUpgrade) {
		return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) HasNextUpgrade() bool {
	if o != nil && !IsNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystNextUpgrade and assigns it to the NextUpgrade field.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) SetNextUpgrade(v CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalystNextUpgrade) {
	o.NextUpgrade = &v
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst struct {
	value *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) Get() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) Set(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst {
	return &NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchCatalyst) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


