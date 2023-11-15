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

// checks if the UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkClientSplashAuthorizationStatusRequestSsids9{}

// UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 Splash authorization for SSID 9
type UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids9 instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids9() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids9{}
	return &this
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids9WithDefaults instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids9WithDefaults() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids9{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) GetIsAuthorized() bool {
	if o == nil || IsNil(o.IsAuthorized) {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthorized) {
		return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) HasIsAuthorized() bool {
	if o != nil && !IsNil(o.IsAuthorized) {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAuthorized) {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return toSerialize, nil
}

type NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9 struct {
	value *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9
	isSet bool
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9) Get() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 {
	return v.value
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9) Set(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9 {
	return &NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


