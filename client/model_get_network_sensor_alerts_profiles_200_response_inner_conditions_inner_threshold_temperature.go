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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature struct {
	// Alerting threshold in degrees Celsius.
	Celsius *float32 `json:"celsius,omitempty"`
	// Alerting threshold in degrees Fahrenheit.
	Fahrenheit *float32 `json:"fahrenheit,omitempty"`
	// Alerting threshold as a qualitative temperature level.
	Quality *string `json:"quality,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperatureWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperatureWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature{}
	return &this
}

// GetCelsius returns the Celsius field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) GetCelsius() float32 {
	if o == nil || IsNil(o.Celsius) {
		var ret float32
		return ret
	}
	return *o.Celsius
}

// GetCelsiusOk returns a tuple with the Celsius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) GetCelsiusOk() (*float32, bool) {
	if o == nil || IsNil(o.Celsius) {
		return nil, false
	}
	return o.Celsius, true
}

// HasCelsius returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) HasCelsius() bool {
	if o != nil && !IsNil(o.Celsius) {
		return true
	}

	return false
}

// SetCelsius gets a reference to the given float32 and assigns it to the Celsius field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) SetCelsius(v float32) {
	o.Celsius = &v
}

// GetFahrenheit returns the Fahrenheit field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) GetFahrenheit() float32 {
	if o == nil || IsNil(o.Fahrenheit) {
		var ret float32
		return ret
	}
	return *o.Fahrenheit
}

// GetFahrenheitOk returns a tuple with the Fahrenheit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) GetFahrenheitOk() (*float32, bool) {
	if o == nil || IsNil(o.Fahrenheit) {
		return nil, false
	}
	return o.Fahrenheit, true
}

// HasFahrenheit returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) HasFahrenheit() bool {
	if o != nil && !IsNil(o.Fahrenheit) {
		return true
	}

	return false
}

// SetFahrenheit gets a reference to the given float32 and assigns it to the Fahrenheit field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) SetFahrenheit(v float32) {
	o.Fahrenheit = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) SetQuality(v string) {
	o.Quality = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Celsius) {
		toSerialize["celsius"] = o.Celsius
	}
	if !IsNil(o.Fahrenheit) {
		toSerialize["fahrenheit"] = o.Fahrenheit
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdTemperature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


