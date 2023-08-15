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

// checks if the GetNetworkSettings200ResponseLocalStatusPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSettings200ResponseLocalStatusPage{}

// GetNetworkSettings200ResponseLocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type GetNetworkSettings200ResponseLocalStatusPage struct {
	Authentication *GetNetworkSettings200ResponseLocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewGetNetworkSettings200ResponseLocalStatusPage instantiates a new GetNetworkSettings200ResponseLocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSettings200ResponseLocalStatusPage() *GetNetworkSettings200ResponseLocalStatusPage {
	this := GetNetworkSettings200ResponseLocalStatusPage{}
	return &this
}

// NewGetNetworkSettings200ResponseLocalStatusPageWithDefaults instantiates a new GetNetworkSettings200ResponseLocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSettings200ResponseLocalStatusPageWithDefaults() *GetNetworkSettings200ResponseLocalStatusPage {
	this := GetNetworkSettings200ResponseLocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *GetNetworkSettings200ResponseLocalStatusPage) GetAuthentication() GetNetworkSettings200ResponseLocalStatusPageAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret GetNetworkSettings200ResponseLocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseLocalStatusPage) GetAuthenticationOk() (*GetNetworkSettings200ResponseLocalStatusPageAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *GetNetworkSettings200ResponseLocalStatusPage) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given GetNetworkSettings200ResponseLocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *GetNetworkSettings200ResponseLocalStatusPage) SetAuthentication(v GetNetworkSettings200ResponseLocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o GetNetworkSettings200ResponseLocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSettings200ResponseLocalStatusPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return toSerialize, nil
}

type NullableGetNetworkSettings200ResponseLocalStatusPage struct {
	value *GetNetworkSettings200ResponseLocalStatusPage
	isSet bool
}

func (v NullableGetNetworkSettings200ResponseLocalStatusPage) Get() *GetNetworkSettings200ResponseLocalStatusPage {
	return v.value
}

func (v *NullableGetNetworkSettings200ResponseLocalStatusPage) Set(val *GetNetworkSettings200ResponseLocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSettings200ResponseLocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSettings200ResponseLocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSettings200ResponseLocalStatusPage(val *GetNetworkSettings200ResponseLocalStatusPage) *NullableGetNetworkSettings200ResponseLocalStatusPage {
	return &NullableGetNetworkSettings200ResponseLocalStatusPage{value: val, isSet: true}
}

func (v NullableGetNetworkSettings200ResponseLocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSettings200ResponseLocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


