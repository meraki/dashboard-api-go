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

// checks if the GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey{}

// GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey API key
type GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey struct {
	// If API key is created for this user
	Created *bool `json:"created,omitempty"`
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey() *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey{}
	return &this
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationApiKeyWithDefaults instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationApiKeyWithDefaults() *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) GetCreated() bool {
	if o == nil || IsNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) GetCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) SetCreated(v bool) {
	o.Created = &v
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	return toSerialize, nil
}

type NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey struct {
	value *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey
	isSet bool
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) Get() *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey {
	return v.value
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) Set(val *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey(val *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey {
	return &NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey{value: val, isSet: true}
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


