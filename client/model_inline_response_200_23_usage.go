/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023Usage Usage, sent and received
type InlineResponse20023Usage struct {
	// Usage sent by the client
	Sent *float32 `json:"sent,omitempty"`
	// Usage received by the client
	Recv *float32 `json:"recv,omitempty"`
}

// NewInlineResponse20023Usage instantiates a new InlineResponse20023Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023Usage() *InlineResponse20023Usage {
	this := InlineResponse20023Usage{}
	return &this
}

// NewInlineResponse20023UsageWithDefaults instantiates a new InlineResponse20023Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023UsageWithDefaults() *InlineResponse20023Usage {
	this := InlineResponse20023Usage{}
	return &this
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse20023Usage) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023Usage) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse20023Usage) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *InlineResponse20023Usage) SetSent(v float32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *InlineResponse20023Usage) GetRecv() float32 {
	if o == nil || isNil(o.Recv) {
		var ret float32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023Usage) GetRecvOk() (*float32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *InlineResponse20023Usage) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given float32 and assigns it to the Recv field.
func (o *InlineResponse20023Usage) SetRecv(v float32) {
	o.Recv = &v
}

func (o InlineResponse20023Usage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023Usage struct {
	value *InlineResponse20023Usage
	isSet bool
}

func (v NullableInlineResponse20023Usage) Get() *InlineResponse20023Usage {
	return v.value
}

func (v *NullableInlineResponse20023Usage) Set(val *InlineResponse20023Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023Usage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023Usage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023Usage(val *InlineResponse20023Usage) *NullableInlineResponse20023Usage {
	return &NullableInlineResponse20023Usage{value: val, isSet: true}
}

func (v NullableInlineResponse20023Usage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023Usage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

