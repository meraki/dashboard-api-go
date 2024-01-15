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

// checks if the UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner{}

// UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner struct for UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner
type UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner struct {
	// Seconds since Sunday at midnight when the outage range starts.
	Start int32 `json:"start"`
	// Seconds since Sunday at midnight when that outage range ends.
	End int32 `json:"end"`
}

// NewUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner instantiates a new UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner(start int32, end int32) *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner {
	this := UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner{}
	this.Start = start
	this.End = end
	return &this
}

// NewUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInnerWithDefaults() *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner {
	this := UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner{}
	return &this
}

// GetStart returns the Start field value
func (o *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) SetStart(v int32) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) SetEnd(v int32) {
	o.End = v
}

func (o UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner struct {
	value *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) Get() *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) Set(val *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner(val *UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) *NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner {
	return &NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


