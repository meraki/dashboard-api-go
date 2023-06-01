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

// InlineResponse20017ValueDestination Destination of 'custom' type traffic filter
type InlineResponse20017ValueDestination struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address (e.g.\"192.168.10.1\", which is the same as \"192.168.10.1/32\"), or \"any\"
	Cidr *string `json:"cidr,omitempty"`
}

// NewInlineResponse20017ValueDestination instantiates a new InlineResponse20017ValueDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017ValueDestination() *InlineResponse20017ValueDestination {
	this := InlineResponse20017ValueDestination{}
	return &this
}

// NewInlineResponse20017ValueDestinationWithDefaults instantiates a new InlineResponse20017ValueDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017ValueDestinationWithDefaults() *InlineResponse20017ValueDestination {
	this := InlineResponse20017ValueDestination{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse20017ValueDestination) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017ValueDestination) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse20017ValueDestination) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *InlineResponse20017ValueDestination) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineResponse20017ValueDestination) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017ValueDestination) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse20017ValueDestination) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineResponse20017ValueDestination) SetCidr(v string) {
	o.Cidr = &v
}

func (o InlineResponse20017ValueDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017ValueDestination struct {
	value *InlineResponse20017ValueDestination
	isSet bool
}

func (v NullableInlineResponse20017ValueDestination) Get() *InlineResponse20017ValueDestination {
	return v.value
}

func (v *NullableInlineResponse20017ValueDestination) Set(val *InlineResponse20017ValueDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017ValueDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017ValueDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017ValueDestination(val *InlineResponse20017ValueDestination) *NullableInlineResponse20017ValueDestination {
	return &NullableInlineResponse20017ValueDestination{value: val, isSet: true}
}

func (v NullableInlineResponse20017ValueDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017ValueDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

