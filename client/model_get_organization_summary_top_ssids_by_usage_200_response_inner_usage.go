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

// checks if the GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage{}

// GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage Date usage of the SSID, in megabytes
type GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage struct {
	// Total usage of the SSID
	Total *float32 `json:"total,omitempty"`
	// Downstream usage of the SSID
	Downstream *float32 `json:"downstream,omitempty"`
	// Upstream usage of the SSID
	Upstream *float32 `json:"upstream,omitempty"`
	// Percentage usage of the SSID
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage instantiates a new GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage{}
	return &this
}

// NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsageWithDefaults instantiates a new GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsageWithDefaults() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetDownstream() float32 {
	if o == nil || IsNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetDownstreamOk() (*float32, bool) {
	if o == nil || IsNil(o.Downstream) {
		return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) HasDownstream() bool {
	if o != nil && !IsNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) SetDownstream(v float32) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetUpstream() float32 {
	if o == nil || IsNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetUpstreamOk() (*float32, bool) {
	if o == nil || IsNil(o.Upstream) {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) HasUpstream() bool {
	if o != nil && !IsNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !IsNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage struct {
	value *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage
	isSet bool
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) Get() *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) Set(val *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage(val *GetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage {
	return &NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopSsidsByUsage200ResponseInnerUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


