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

// GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts Counts of currently alerting sensors, aggregated by alerting metric
type GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts struct {
	// Number of sensors that are currently alerting due to an open door
	Door *int32 `json:"door,omitempty"`
	// Number of sensors that are currently alerting due to humidity readings
	Humidity *int32 `json:"humidity,omitempty"`
	// Number of sensors that are currently alerting due to indoor air quality readings
	IndoorAirQuality *int32 `json:"indoorAirQuality,omitempty"`
	Noise *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise `json:"noise,omitempty"`
	// Number of sensors that are currently alerting due to PM2.5 readings
	Pm25 *int32 `json:"pm25,omitempty"`
	// Number of sensors that are currently alerting due to temperature readings
	Temperature *int32 `json:"temperature,omitempty"`
	// Number of sensors that are currently alerting due to TVOC readings
	Tvoc *int32 `json:"tvoc,omitempty"`
	// Number of sensors that are currently alerting due to the presence of water
	Water *int32 `json:"water,omitempty"`
}

// NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts instantiates a new GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts() *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts {
	this := GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts{}
	return &this
}

// NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsWithDefaults instantiates a new GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsWithDefaults() *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts {
	this := GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts{}
	return &this
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetDoor() int32 {
	if o == nil || isNil(o.Door) {
		var ret int32
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetDoorOk() (*int32, bool) {
	if o == nil || isNil(o.Door) {
    return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasDoor() bool {
	if o != nil && !isNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given int32 and assigns it to the Door field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetDoor(v int32) {
	o.Door = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetHumidity() int32 {
	if o == nil || isNil(o.Humidity) {
		var ret int32
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetHumidityOk() (*int32, bool) {
	if o == nil || isNil(o.Humidity) {
    return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasHumidity() bool {
	if o != nil && !isNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given int32 and assigns it to the Humidity field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetHumidity(v int32) {
	o.Humidity = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetIndoorAirQuality() int32 {
	if o == nil || isNil(o.IndoorAirQuality) {
		var ret int32
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetIndoorAirQualityOk() (*int32, bool) {
	if o == nil || isNil(o.IndoorAirQuality) {
    return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasIndoorAirQuality() bool {
	if o != nil && !isNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given int32 and assigns it to the IndoorAirQuality field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetIndoorAirQuality(v int32) {
	o.IndoorAirQuality = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetNoise() GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise {
	if o == nil || isNil(o.Noise) {
		var ret GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetNoiseOk() (*GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise, bool) {
	if o == nil || isNil(o.Noise) {
    return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasNoise() bool {
	if o != nil && !isNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise and assigns it to the Noise field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetNoise(v GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise) {
	o.Noise = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetPm25() int32 {
	if o == nil || isNil(o.Pm25) {
		var ret int32
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetPm25Ok() (*int32, bool) {
	if o == nil || isNil(o.Pm25) {
    return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasPm25() bool {
	if o != nil && !isNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given int32 and assigns it to the Pm25 field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetPm25(v int32) {
	o.Pm25 = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTemperature() int32 {
	if o == nil || isNil(o.Temperature) {
		var ret int32
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTemperatureOk() (*int32, bool) {
	if o == nil || isNil(o.Temperature) {
    return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasTemperature() bool {
	if o != nil && !isNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given int32 and assigns it to the Temperature field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetTemperature(v int32) {
	o.Temperature = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTvoc() int32 {
	if o == nil || isNil(o.Tvoc) {
		var ret int32
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTvocOk() (*int32, bool) {
	if o == nil || isNil(o.Tvoc) {
    return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasTvoc() bool {
	if o != nil && !isNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given int32 and assigns it to the Tvoc field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetTvoc(v int32) {
	o.Tvoc = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetWater() int32 {
	if o == nil || isNil(o.Water) {
		var ret int32
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetWaterOk() (*int32, bool) {
	if o == nil || isNil(o.Water) {
    return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasWater() bool {
	if o != nil && !isNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given int32 and assigns it to the Water field.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetWater(v int32) {
	o.Water = &v
}

func (o GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !isNil(o.Humidity) {
		toSerialize["humidity"] = o.Humidity
	}
	if !isNil(o.IndoorAirQuality) {
		toSerialize["indoorAirQuality"] = o.IndoorAirQuality
	}
	if !isNil(o.Noise) {
		toSerialize["noise"] = o.Noise
	}
	if !isNil(o.Pm25) {
		toSerialize["pm25"] = o.Pm25
	}
	if !isNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !isNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !isNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts struct {
	value *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts
	isSet bool
}

func (v NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) Get() *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) Set(val *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts(val *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) *NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts {
	return &NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

