/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade Details of the last firmware upgrade on the device
type GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade struct {
	// Timestamp of the last successful firmware upgrade
	Time *time.Time `json:"time,omitempty"`
	FromVersion *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeFromVersion `json:"fromVersion,omitempty"`
	ToVersion *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade{}
	return &this
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeWithDefaults instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeWithDefaults() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) SetTime(v time.Time) {
	o.Time = &v
}

// GetFromVersion returns the FromVersion field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) GetFromVersion() GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeFromVersion {
	if o == nil || isNil(o.FromVersion) {
		var ret GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeFromVersion
		return ret
	}
	return *o.FromVersion
}

// GetFromVersionOk returns a tuple with the FromVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) GetFromVersionOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeFromVersion, bool) {
	if o == nil || isNil(o.FromVersion) {
    return nil, false
	}
	return o.FromVersion, true
}

// HasFromVersion returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) HasFromVersion() bool {
	if o != nil && !isNil(o.FromVersion) {
		return true
	}

	return false
}

// SetFromVersion gets a reference to the given GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeFromVersion and assigns it to the FromVersion field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) SetFromVersion(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeFromVersion) {
	o.FromVersion = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) GetToVersion() GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) GetToVersionOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion and assigns it to the ToVersion field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) SetToVersion(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) {
	o.ToVersion = &v
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.FromVersion) {
		toSerialize["fromVersion"] = o.FromVersion
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade struct {
	value *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) Get() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) Set(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade {
	return &NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


