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

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc struct {
	// Alerting threshold as TVOC micrograms per cubic meter.
	Concentration *int32 `json:"concentration,omitempty"`
	// Alerting threshold as a qualitative TVOC level.
	Quality *string `json:"quality,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvocWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvocWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) SetConcentration(v int32) {
	o.Concentration = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) SetQuality(v string) {
	o.Quality = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTvoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


