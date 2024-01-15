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

// checks if the GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts{}

// GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts The network devices to be updated
type GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts struct {
	Switch *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch `json:"switch,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts{}
	return &this
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) GetSwitch() GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch {
	if o == nil || IsNil(o.Switch) {
		var ret GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) GetSwitchOk() (*GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch, bool) {
	if o == nil || IsNil(o.Switch) {
		return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) HasSwitch() bool {
	if o != nil && !IsNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch and assigns it to the Switch field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) SetSwitch(v GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) {
	o.Switch = &v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) Get() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) Set(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


