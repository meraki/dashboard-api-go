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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality{}

// GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'.
type GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality struct {
	// Indoor air quality score between 0 and 100.
	Score *int32 `json:"score,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality() *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQualityWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQualityWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) SetScore(v int32) {
	o.Score = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality(val *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


