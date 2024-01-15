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

// checks if the GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner{}

// GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner struct for GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner
type GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Number of useful information bits delivered over a network per unit of time
	Goodput *int32 `json:"goodput,omitempty"`
	// Duration of the response, in milliseconds
	ResponseDuration *int32 `json:"responseDuration,omitempty"`
}

// NewGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner instantiates a new GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner() *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner {
	this := GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner{}
	return &this
}

// NewGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInnerWithDefaults instantiates a new GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInnerWithDefaults() *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner {
	this := GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetGoodput returns the Goodput field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) GetGoodput() int32 {
	if o == nil || IsNil(o.Goodput) {
		var ret int32
		return ret
	}
	return *o.Goodput
}

// GetGoodputOk returns a tuple with the Goodput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) GetGoodputOk() (*int32, bool) {
	if o == nil || IsNil(o.Goodput) {
		return nil, false
	}
	return o.Goodput, true
}

// HasGoodput returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) HasGoodput() bool {
	if o != nil && !IsNil(o.Goodput) {
		return true
	}

	return false
}

// SetGoodput gets a reference to the given int32 and assigns it to the Goodput field.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) SetGoodput(v int32) {
	o.Goodput = &v
}

// GetResponseDuration returns the ResponseDuration field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) GetResponseDuration() int32 {
	if o == nil || IsNil(o.ResponseDuration) {
		var ret int32
		return ret
	}
	return *o.ResponseDuration
}

// GetResponseDurationOk returns a tuple with the ResponseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) GetResponseDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponseDuration) {
		return nil, false
	}
	return o.ResponseDuration, true
}

// HasResponseDuration returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) HasResponseDuration() bool {
	if o != nil && !IsNil(o.ResponseDuration) {
		return true
	}

	return false
}

// SetResponseDuration gets a reference to the given int32 and assigns it to the ResponseDuration field.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) SetResponseDuration(v int32) {
	o.ResponseDuration = &v
}

func (o GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Goodput) {
		toSerialize["goodput"] = o.Goodput
	}
	if !IsNil(o.ResponseDuration) {
		toSerialize["responseDuration"] = o.ResponseDuration
	}
	return toSerialize, nil
}

type NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner struct {
	value *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner
	isSet bool
}

func (v NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) Get() *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner {
	return v.value
}

func (v *NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) Set(val *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner(val *GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) *NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner {
	return &NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner{value: val, isSet: true}
}

func (v NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


