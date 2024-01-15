/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner{}

// GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner struct for GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner
type GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner struct {
	// Timestamp for this data point
	Ts *time.Time `json:"ts,omitempty"`
	// Loss percentage
	LossPercent *float32 `json:"lossPercent,omitempty"`
	// Latency in milliseconds
	LatencyMs *float32 `json:"latencyMs,omitempty"`
}

// NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner instantiates a new GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner {
	this := GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner{}
	return &this
}

// NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInnerWithDefaults() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner {
	this := GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) SetTs(v time.Time) {
	o.Ts = &v
}

// GetLossPercent returns the LossPercent field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) GetLossPercent() float32 {
	if o == nil || IsNil(o.LossPercent) {
		var ret float32
		return ret
	}
	return *o.LossPercent
}

// GetLossPercentOk returns a tuple with the LossPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) GetLossPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.LossPercent) {
		return nil, false
	}
	return o.LossPercent, true
}

// HasLossPercent returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) HasLossPercent() bool {
	if o != nil && !IsNil(o.LossPercent) {
		return true
	}

	return false
}

// SetLossPercent gets a reference to the given float32 and assigns it to the LossPercent field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) SetLossPercent(v float32) {
	o.LossPercent = &v
}

// GetLatencyMs returns the LatencyMs field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) GetLatencyMs() float32 {
	if o == nil || IsNil(o.LatencyMs) {
		var ret float32
		return ret
	}
	return *o.LatencyMs
}

// GetLatencyMsOk returns a tuple with the LatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) GetLatencyMsOk() (*float32, bool) {
	if o == nil || IsNil(o.LatencyMs) {
		return nil, false
	}
	return o.LatencyMs, true
}

// HasLatencyMs returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) HasLatencyMs() bool {
	if o != nil && !IsNil(o.LatencyMs) {
		return true
	}

	return false
}

// SetLatencyMs gets a reference to the given float32 and assigns it to the LatencyMs field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) SetLatencyMs(v float32) {
	o.LatencyMs = &v
}

func (o GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.LossPercent) {
		toSerialize["lossPercent"] = o.LossPercent
	}
	if !IsNil(o.LatencyMs) {
		toSerialize["latencyMs"] = o.LatencyMs
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner struct {
	value *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner
	isSet bool
}

func (v NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) Get() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) Set(val *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner(val *GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner {
	return &NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


