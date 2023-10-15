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

// checks if the GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise{}

// GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise Object containing the number of sensor alerts that occurred due to noise readings
type GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise struct {
	// Number of sensor alerts that occurred due to ambient noise readings
	Ambient *int32 `json:"ambient,omitempty"`
}

// NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise {
	this := GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise{}
	return &this
}

// NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoiseWithDefaults instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoiseWithDefaults() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise {
	this := GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise{}
	return &this
}

// GetAmbient returns the Ambient field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) GetAmbient() int32 {
	if o == nil || IsNil(o.Ambient) {
		var ret int32
		return ret
	}
	return *o.Ambient
}

// GetAmbientOk returns a tuple with the Ambient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) GetAmbientOk() (*int32, bool) {
	if o == nil || IsNil(o.Ambient) {
		return nil, false
	}
	return o.Ambient, true
}

// HasAmbient returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) HasAmbient() bool {
	if o != nil && !IsNil(o.Ambient) {
		return true
	}

	return false
}

// SetAmbient gets a reference to the given int32 and assigns it to the Ambient field.
func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) SetAmbient(v int32) {
	o.Ambient = &v
}

func (o GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ambient) {
		toSerialize["ambient"] = o.Ambient
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise struct {
	value *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise
	isSet bool
}

func (v NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) Get() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) Set(val *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise(val *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise {
	return &NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


