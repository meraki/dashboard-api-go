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

// checks if the GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients{}

// GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients Clients info of the SSID
type GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients struct {
	Counts *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts `json:"counts,omitempty"`
}

// NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients instantiates a new GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients {
	this := GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients{}
	return &this
}

// NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsWithDefaults instantiates a new GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsWithDefaults() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients {
	this := GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) GetCounts() GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts {
	if o == nil || IsNil(o.Counts) {
		var ret GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) GetCountsOk() (*GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts and assigns it to the Counts field.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) SetCounts(v GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClientsCounts) {
	o.Counts = &v
}

func (o GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients struct {
	value *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients
	isSet bool
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) Get() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) Set(val *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients(val *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients {
	return &NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


