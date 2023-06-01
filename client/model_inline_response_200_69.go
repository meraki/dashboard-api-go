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

// InlineResponse20069 struct for InlineResponse20069
type InlineResponse20069 struct {
	// List of the syslog servers for this network
	Servers []InlineResponse20069Servers `json:"servers,omitempty"`
}

// NewInlineResponse20069 instantiates a new InlineResponse20069 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069() *InlineResponse20069 {
	this := InlineResponse20069{}
	return &this
}

// NewInlineResponse20069WithDefaults instantiates a new InlineResponse20069 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069WithDefaults() *InlineResponse20069 {
	this := InlineResponse20069{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *InlineResponse20069) GetServers() []InlineResponse20069Servers {
	if o == nil || isNil(o.Servers) {
		var ret []InlineResponse20069Servers
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069) GetServersOk() ([]InlineResponse20069Servers, bool) {
	if o == nil || isNil(o.Servers) {
    return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *InlineResponse20069) HasServers() bool {
	if o != nil && !isNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []InlineResponse20069Servers and assigns it to the Servers field.
func (o *InlineResponse20069) SetServers(v []InlineResponse20069Servers) {
	o.Servers = v
}

func (o InlineResponse20069) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20069 struct {
	value *InlineResponse20069
	isSet bool
}

func (v NullableInlineResponse20069) Get() *InlineResponse20069 {
	return v.value
}

func (v *NullableInlineResponse20069) Set(val *InlineResponse20069) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069(val *InlineResponse20069) *NullableInlineResponse20069 {
	return &NullableInlineResponse20069{value: val, isSet: true}
}

func (v NullableInlineResponse20069) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

