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

// checks if the GetNetworkFloorPlans200ResponseInnerCenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFloorPlans200ResponseInnerCenter{}

// GetNetworkFloorPlans200ResponseInnerCenter The longitude and latitude of the center of your floor plan. The 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.
type GetNetworkFloorPlans200ResponseInnerCenter struct {
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewGetNetworkFloorPlans200ResponseInnerCenter instantiates a new GetNetworkFloorPlans200ResponseInnerCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFloorPlans200ResponseInnerCenter() *GetNetworkFloorPlans200ResponseInnerCenter {
	this := GetNetworkFloorPlans200ResponseInnerCenter{}
	return &this
}

// NewGetNetworkFloorPlans200ResponseInnerCenterWithDefaults instantiates a new GetNetworkFloorPlans200ResponseInnerCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFloorPlans200ResponseInnerCenterWithDefaults() *GetNetworkFloorPlans200ResponseInnerCenter {
	this := GetNetworkFloorPlans200ResponseInnerCenter{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) GetLat() float32 {
	if o == nil || IsNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) GetLatOk() (*float32, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) GetLng() float32 {
	if o == nil || IsNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) GetLngOk() (*float32, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *GetNetworkFloorPlans200ResponseInnerCenter) SetLng(v float32) {
	o.Lng = &v
}

func (o GetNetworkFloorPlans200ResponseInnerCenter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFloorPlans200ResponseInnerCenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	return toSerialize, nil
}

type NullableGetNetworkFloorPlans200ResponseInnerCenter struct {
	value *GetNetworkFloorPlans200ResponseInnerCenter
	isSet bool
}

func (v NullableGetNetworkFloorPlans200ResponseInnerCenter) Get() *GetNetworkFloorPlans200ResponseInnerCenter {
	return v.value
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerCenter) Set(val *GetNetworkFloorPlans200ResponseInnerCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFloorPlans200ResponseInnerCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFloorPlans200ResponseInnerCenter(val *GetNetworkFloorPlans200ResponseInnerCenter) *NullableGetNetworkFloorPlans200ResponseInnerCenter {
	return &NullableGetNetworkFloorPlans200ResponseInnerCenter{value: val, isSet: true}
}

func (v NullableGetNetworkFloorPlans200ResponseInnerCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


