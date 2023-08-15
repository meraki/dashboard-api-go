/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerHumidity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerHumidity{}

// GetOrganizationSensorReadingsHistory200ResponseInnerHumidity Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'.
type GetOrganizationSensorReadingsHistory200ResponseInnerHumidity struct {
	// Humidity reading in %RH.
	RelativePercentage *int32 `json:"relativePercentage,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerHumidity instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerHumidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerHumidity() *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerHumidity{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerHumidityWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerHumidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerHumidityWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerHumidity{}
	return &this
}

// GetRelativePercentage returns the RelativePercentage field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) GetRelativePercentage() int32 {
	if o == nil || IsNil(o.RelativePercentage) {
		var ret int32
		return ret
	}
	return *o.RelativePercentage
}

// GetRelativePercentageOk returns a tuple with the RelativePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) GetRelativePercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.RelativePercentage) {
		return nil, false
	}
	return o.RelativePercentage, true
}

// HasRelativePercentage returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) HasRelativePercentage() bool {
	if o != nil && !IsNil(o.RelativePercentage) {
		return true
	}

	return false
}

// SetRelativePercentage gets a reference to the given int32 and assigns it to the RelativePercentage field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) SetRelativePercentage(v int32) {
	o.RelativePercentage = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelativePercentage) {
		toSerialize["relativePercentage"] = o.RelativePercentage
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity(val *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerHumidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


