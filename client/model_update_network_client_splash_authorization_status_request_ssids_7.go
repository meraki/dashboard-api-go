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

// UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 Splash authorization for SSID 7
type UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids7 instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids7() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids7{}
	return &this
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids7WithDefaults instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids7WithDefaults() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids7{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) GetIsAuthorized() bool {
	if o == nil || isNil(o.IsAuthorized) {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || isNil(o.IsAuthorized) {
    return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) HasIsAuthorized() bool {
	if o != nil && !isNil(o.IsAuthorized) {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsAuthorized) {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7 struct {
	value *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7
	isSet bool
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7) Get() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 {
	return v.value
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7) Set(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7 {
	return &NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


