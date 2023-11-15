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

// checks if the UpdateNetworkSettingsRequestLocalStatusPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSettingsRequestLocalStatusPage{}

// UpdateNetworkSettingsRequestLocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type UpdateNetworkSettingsRequestLocalStatusPage struct {
	Authentication *UpdateNetworkSettingsRequestLocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewUpdateNetworkSettingsRequestLocalStatusPage instantiates a new UpdateNetworkSettingsRequestLocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSettingsRequestLocalStatusPage() *UpdateNetworkSettingsRequestLocalStatusPage {
	this := UpdateNetworkSettingsRequestLocalStatusPage{}
	return &this
}

// NewUpdateNetworkSettingsRequestLocalStatusPageWithDefaults instantiates a new UpdateNetworkSettingsRequestLocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSettingsRequestLocalStatusPageWithDefaults() *UpdateNetworkSettingsRequestLocalStatusPage {
	this := UpdateNetworkSettingsRequestLocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) GetAuthentication() UpdateNetworkSettingsRequestLocalStatusPageAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret UpdateNetworkSettingsRequestLocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) GetAuthenticationOk() (*UpdateNetworkSettingsRequestLocalStatusPageAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given UpdateNetworkSettingsRequestLocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) SetAuthentication(v UpdateNetworkSettingsRequestLocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o UpdateNetworkSettingsRequestLocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSettingsRequestLocalStatusPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSettingsRequestLocalStatusPage struct {
	value *UpdateNetworkSettingsRequestLocalStatusPage
	isSet bool
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPage) Get() *UpdateNetworkSettingsRequestLocalStatusPage {
	return v.value
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPage) Set(val *UpdateNetworkSettingsRequestLocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSettingsRequestLocalStatusPage(val *UpdateNetworkSettingsRequestLocalStatusPage) *NullableUpdateNetworkSettingsRequestLocalStatusPage {
	return &NullableUpdateNetworkSettingsRequestLocalStatusPage{value: val, isSet: true}
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


