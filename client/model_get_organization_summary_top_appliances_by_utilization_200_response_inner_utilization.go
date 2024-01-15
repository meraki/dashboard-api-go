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

// checks if the GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization{}

// GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization Utilization of the appliance
type GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization struct {
	Average *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationAverage `json:"average,omitempty"`
}

// NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization instantiates a new GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization() *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization {
	this := GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization{}
	return &this
}

// NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationWithDefaults instantiates a new GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationWithDefaults() *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization {
	this := GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) GetAverage() GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationAverage {
	if o == nil || IsNil(o.Average) {
		var ret GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationAverage
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) GetAverageOk() (*GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationAverage, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationAverage and assigns it to the Average field.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) SetAverage(v GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilizationAverage) {
	o.Average = &v
}

func (o GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization struct {
	value *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization
	isSet bool
}

func (v NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) Get() *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) Set(val *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization(val *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization {
	return &NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


