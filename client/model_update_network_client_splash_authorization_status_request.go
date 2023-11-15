/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkClientSplashAuthorizationStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkClientSplashAuthorizationStatusRequest{}

// UpdateNetworkClientSplashAuthorizationStatusRequest struct for UpdateNetworkClientSplashAuthorizationStatusRequest
type UpdateNetworkClientSplashAuthorizationStatusRequest struct {
	Ssids UpdateNetworkClientSplashAuthorizationStatusRequestSsids `json:"ssids"`
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequest instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientSplashAuthorizationStatusRequest(ssids UpdateNetworkClientSplashAuthorizationStatusRequestSsids) *UpdateNetworkClientSplashAuthorizationStatusRequest {
	this := UpdateNetworkClientSplashAuthorizationStatusRequest{}
	this.Ssids = ssids
	return &this
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestWithDefaults instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientSplashAuthorizationStatusRequestWithDefaults() *UpdateNetworkClientSplashAuthorizationStatusRequest {
	this := UpdateNetworkClientSplashAuthorizationStatusRequest{}
	return &this
}

// GetSsids returns the Ssids field value
func (o *UpdateNetworkClientSplashAuthorizationStatusRequest) GetSsids() UpdateNetworkClientSplashAuthorizationStatusRequestSsids {
	if o == nil {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids
		return ret
	}

	return o.Ssids
}

// GetSsidsOk returns a tuple with the Ssids field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequest) GetSsidsOk() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssids, true
}

// SetSsids sets field value
func (o *UpdateNetworkClientSplashAuthorizationStatusRequest) SetSsids(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids) {
	o.Ssids = v
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ssids"] = o.Ssids
	return toSerialize, nil
}

type NullableUpdateNetworkClientSplashAuthorizationStatusRequest struct {
	value *UpdateNetworkClientSplashAuthorizationStatusRequest
	isSet bool
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequest) Get() *UpdateNetworkClientSplashAuthorizationStatusRequest {
	return v.value
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequest) Set(val *UpdateNetworkClientSplashAuthorizationStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientSplashAuthorizationStatusRequest(val *UpdateNetworkClientSplashAuthorizationStatusRequest) *NullableUpdateNetworkClientSplashAuthorizationStatusRequest {
	return &NullableUpdateNetworkClientSplashAuthorizationStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


