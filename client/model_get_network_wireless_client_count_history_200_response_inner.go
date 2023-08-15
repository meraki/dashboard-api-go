/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkWirelessClientCountHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessClientCountHistory200ResponseInner{}

// GetNetworkWirelessClientCountHistory200ResponseInner struct for GetNetworkWirelessClientCountHistory200ResponseInner
type GetNetworkWirelessClientCountHistory200ResponseInner struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Number of connected clients
	ClientCount *int32 `json:"clientCount,omitempty"`
}

// NewGetNetworkWirelessClientCountHistory200ResponseInner instantiates a new GetNetworkWirelessClientCountHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessClientCountHistory200ResponseInner() *GetNetworkWirelessClientCountHistory200ResponseInner {
	this := GetNetworkWirelessClientCountHistory200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessClientCountHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessClientCountHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessClientCountHistory200ResponseInnerWithDefaults() *GetNetworkWirelessClientCountHistory200ResponseInner {
	this := GetNetworkWirelessClientCountHistory200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetClientCount returns the ClientCount field value if set, zero value otherwise.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) GetClientCount() int32 {
	if o == nil || IsNil(o.ClientCount) {
		var ret int32
		return ret
	}
	return *o.ClientCount
}

// GetClientCountOk returns a tuple with the ClientCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) GetClientCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ClientCount) {
		return nil, false
	}
	return o.ClientCount, true
}

// HasClientCount returns a boolean if a field has been set.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) HasClientCount() bool {
	if o != nil && !IsNil(o.ClientCount) {
		return true
	}

	return false
}

// SetClientCount gets a reference to the given int32 and assigns it to the ClientCount field.
func (o *GetNetworkWirelessClientCountHistory200ResponseInner) SetClientCount(v int32) {
	o.ClientCount = &v
}

func (o GetNetworkWirelessClientCountHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessClientCountHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.ClientCount) {
		toSerialize["clientCount"] = o.ClientCount
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessClientCountHistory200ResponseInner struct {
	value *GetNetworkWirelessClientCountHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessClientCountHistory200ResponseInner) Get() *GetNetworkWirelessClientCountHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessClientCountHistory200ResponseInner) Set(val *GetNetworkWirelessClientCountHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessClientCountHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessClientCountHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessClientCountHistory200ResponseInner(val *GetNetworkWirelessClientCountHistory200ResponseInner) *NullableGetNetworkWirelessClientCountHistory200ResponseInner {
	return &NullableGetNetworkWirelessClientCountHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessClientCountHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessClientCountHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


