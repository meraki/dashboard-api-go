/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower Power details object
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower struct {
	// The PoE power mode for the AP. Can be 'full' or 'low'
	Mode *string `json:"mode,omitempty"`
	Ac *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc `json:"ac,omitempty"`
	Poe *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe `json:"poe,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) SetMode(v string) {
	o.Mode = &v
}

// GetAc returns the Ac field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetAc() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc {
	if o == nil || isNil(o.Ac) {
		var ret GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc
		return ret
	}
	return *o.Ac
}

// GetAcOk returns a tuple with the Ac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetAcOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc, bool) {
	if o == nil || isNil(o.Ac) {
    return nil, false
	}
	return o.Ac, true
}

// HasAc returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) HasAc() bool {
	if o != nil && !isNil(o.Ac) {
		return true
	}

	return false
}

// SetAc gets a reference to the given GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc and assigns it to the Ac field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) SetAc(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) {
	o.Ac = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetPoe() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe {
	if o == nil || isNil(o.Poe) {
		var ret GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetPoeOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe, bool) {
	if o == nil || isNil(o.Poe) {
    return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) HasPoe() bool {
	if o != nil && !isNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe and assigns it to the Poe field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) SetPoe(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe) {
	o.Poe = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Ac) {
		toSerialize["ac"] = o.Ac
	}
	if !isNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

