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

// GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage Data usage of the device
type GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage struct {
	// Total data usage of the device
	Total *float32 `json:"total,omitempty"`
	// Data usage of the device by percentage
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage() *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage{}
	return &this
}

// NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsageWithDefaults instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsageWithDefaults() *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage struct {
	value *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage
	isSet bool
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) Get() *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) Set(val *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage(val *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage {
	return &NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

