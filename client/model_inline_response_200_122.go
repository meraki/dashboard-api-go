/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200122 struct for InlineResponse200122
type InlineResponse200122 struct {
	// Organization APNS Certificate used by devices to communication with Apple
	Certificate *string `json:"certificate,omitempty"`
}

// NewInlineResponse200122 instantiates a new InlineResponse200122 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200122() *InlineResponse200122 {
	this := InlineResponse200122{}
	return &this
}

// NewInlineResponse200122WithDefaults instantiates a new InlineResponse200122 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200122WithDefaults() *InlineResponse200122 {
	this := InlineResponse200122{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *InlineResponse200122) GetCertificate() string {
	if o == nil || isNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122) GetCertificateOk() (*string, bool) {
	if o == nil || isNil(o.Certificate) {
    return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *InlineResponse200122) HasCertificate() bool {
	if o != nil && !isNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *InlineResponse200122) SetCertificate(v string) {
	o.Certificate = &v
}

func (o InlineResponse200122) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200122 struct {
	value *InlineResponse200122
	isSet bool
}

func (v NullableInlineResponse200122) Get() *InlineResponse200122 {
	return v.value
}

func (v *NullableInlineResponse200122) Set(val *InlineResponse200122) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200122) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200122) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200122(val *InlineResponse200122) *NullableInlineResponse200122 {
	return &NullableInlineResponse200122{value: val, isSet: true}
}

func (v NullableInlineResponse200122) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200122) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

