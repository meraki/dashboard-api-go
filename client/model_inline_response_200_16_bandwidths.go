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

// InlineResponse20016Bandwidths struct for InlineResponse20016Bandwidths
type InlineResponse20016Bandwidths struct {
	// name of uplink
	Interface *string `json:"interface,omitempty"`
	// limit in bytes (null indicates unlimited)
	Limit *int32 `json:"limit,omitempty"`
}

// NewInlineResponse20016Bandwidths instantiates a new InlineResponse20016Bandwidths object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016Bandwidths() *InlineResponse20016Bandwidths {
	this := InlineResponse20016Bandwidths{}
	return &this
}

// NewInlineResponse20016BandwidthsWithDefaults instantiates a new InlineResponse20016Bandwidths object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016BandwidthsWithDefaults() *InlineResponse20016Bandwidths {
	this := InlineResponse20016Bandwidths{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *InlineResponse20016Bandwidths) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Bandwidths) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *InlineResponse20016Bandwidths) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *InlineResponse20016Bandwidths) SetInterface(v string) {
	o.Interface = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *InlineResponse20016Bandwidths) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Bandwidths) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *InlineResponse20016Bandwidths) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *InlineResponse20016Bandwidths) SetLimit(v int32) {
	o.Limit = &v
}

func (o InlineResponse20016Bandwidths) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016Bandwidths struct {
	value *InlineResponse20016Bandwidths
	isSet bool
}

func (v NullableInlineResponse20016Bandwidths) Get() *InlineResponse20016Bandwidths {
	return v.value
}

func (v *NullableInlineResponse20016Bandwidths) Set(val *InlineResponse20016Bandwidths) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016Bandwidths) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016Bandwidths) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016Bandwidths(val *InlineResponse20016Bandwidths) *NullableInlineResponse20016Bandwidths {
	return &NullableInlineResponse20016Bandwidths{value: val, isSet: true}
}

func (v NullableInlineResponse20016Bandwidths) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016Bandwidths) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

