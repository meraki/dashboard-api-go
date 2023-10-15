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

// checks if the GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal{}

// GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal An object containing total channel utilization.
type GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal struct {
	// Percentage of total channel utiliation for the given band.
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal{}
	return &this
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotalWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotalWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal struct {
	value *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) Get() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) Set(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal {
	return &NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


