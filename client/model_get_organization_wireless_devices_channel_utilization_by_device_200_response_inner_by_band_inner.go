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

// checks if the GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner{}

// GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner struct for GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner
type GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner struct {
	// The band for the given metrics.
	Band *string `json:"band,omitempty"`
	Wifi *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWifi `json:"wifi,omitempty"`
	NonWifi *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi `json:"nonWifi,omitempty"`
	Total *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal `json:"total,omitempty"`
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner{}
	return &this
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner{}
	return &this
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetBand() string {
	if o == nil || IsNil(o.Band) {
		var ret string
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetBandOk() (*string, bool) {
	if o == nil || IsNil(o.Band) {
		return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) HasBand() bool {
	if o != nil && !IsNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given string and assigns it to the Band field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) SetBand(v string) {
	o.Band = &v
}

// GetWifi returns the Wifi field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetWifi() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWifi {
	if o == nil || IsNil(o.Wifi) {
		var ret GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWifi
		return ret
	}
	return *o.Wifi
}

// GetWifiOk returns a tuple with the Wifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetWifiOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWifi, bool) {
	if o == nil || IsNil(o.Wifi) {
		return nil, false
	}
	return o.Wifi, true
}

// HasWifi returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) HasWifi() bool {
	if o != nil && !IsNil(o.Wifi) {
		return true
	}

	return false
}

// SetWifi gets a reference to the given GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWifi and assigns it to the Wifi field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) SetWifi(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerWifi) {
	o.Wifi = &v
}

// GetNonWifi returns the NonWifi field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetNonWifi() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi {
	if o == nil || IsNil(o.NonWifi) {
		var ret GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi
		return ret
	}
	return *o.NonWifi
}

// GetNonWifiOk returns a tuple with the NonWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetNonWifiOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi, bool) {
	if o == nil || IsNil(o.NonWifi) {
		return nil, false
	}
	return o.NonWifi, true
}

// HasNonWifi returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) HasNonWifi() bool {
	if o != nil && !IsNil(o.NonWifi) {
		return true
	}

	return false
}

// SetNonWifi gets a reference to the given GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi and assigns it to the NonWifi field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) SetNonWifi(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerNonWifi) {
	o.NonWifi = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetTotal() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal {
	if o == nil || IsNil(o.Total) {
		var ret GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) GetTotalOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal and assigns it to the Total field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) SetTotal(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInnerTotal) {
	o.Total = &v
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !IsNil(o.Wifi) {
		toSerialize["wifi"] = o.Wifi
	}
	if !IsNil(o.NonWifi) {
		toSerialize["nonWifi"] = o.NonWifi
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner struct {
	value *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) Get() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) Set(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner {
	return &NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


