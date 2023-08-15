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

// checks if the UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkClientSplashAuthorizationStatusRequestSsids4{}

// UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 Splash authorization for SSID 4
type UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids4 instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids4() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids4{}
	return &this
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids4WithDefaults instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids4WithDefaults() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids4{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) GetIsAuthorized() bool {
	if o == nil || IsNil(o.IsAuthorized) {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthorized) {
		return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) HasIsAuthorized() bool {
	if o != nil && !IsNil(o.IsAuthorized) {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAuthorized) {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return toSerialize, nil
}

type NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4 struct {
	value *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4
	isSet bool
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4) Get() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 {
	return v.value
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4) Set(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4 {
	return &NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


