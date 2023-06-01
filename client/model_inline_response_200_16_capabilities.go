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

// InlineResponse20016Capabilities struct for InlineResponse20016Capabilities
type InlineResponse20016Capabilities struct {
	// model number of appliance
	Model *string `json:"model,omitempty"`
	// array of uplink limits
	Bandwidths []InlineResponse20016Bandwidths `json:"bandwidths,omitempty"`
}

// NewInlineResponse20016Capabilities instantiates a new InlineResponse20016Capabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016Capabilities() *InlineResponse20016Capabilities {
	this := InlineResponse20016Capabilities{}
	return &this
}

// NewInlineResponse20016CapabilitiesWithDefaults instantiates a new InlineResponse20016Capabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016CapabilitiesWithDefaults() *InlineResponse20016Capabilities {
	this := InlineResponse20016Capabilities{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse20016Capabilities) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Capabilities) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse20016Capabilities) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse20016Capabilities) SetModel(v string) {
	o.Model = &v
}

// GetBandwidths returns the Bandwidths field value if set, zero value otherwise.
func (o *InlineResponse20016Capabilities) GetBandwidths() []InlineResponse20016Bandwidths {
	if o == nil || isNil(o.Bandwidths) {
		var ret []InlineResponse20016Bandwidths
		return ret
	}
	return o.Bandwidths
}

// GetBandwidthsOk returns a tuple with the Bandwidths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Capabilities) GetBandwidthsOk() ([]InlineResponse20016Bandwidths, bool) {
	if o == nil || isNil(o.Bandwidths) {
    return nil, false
	}
	return o.Bandwidths, true
}

// HasBandwidths returns a boolean if a field has been set.
func (o *InlineResponse20016Capabilities) HasBandwidths() bool {
	if o != nil && !isNil(o.Bandwidths) {
		return true
	}

	return false
}

// SetBandwidths gets a reference to the given []InlineResponse20016Bandwidths and assigns it to the Bandwidths field.
func (o *InlineResponse20016Capabilities) SetBandwidths(v []InlineResponse20016Bandwidths) {
	o.Bandwidths = v
}

func (o InlineResponse20016Capabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Bandwidths) {
		toSerialize["bandwidths"] = o.Bandwidths
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016Capabilities struct {
	value *InlineResponse20016Capabilities
	isSet bool
}

func (v NullableInlineResponse20016Capabilities) Get() *InlineResponse20016Capabilities {
	return v.value
}

func (v *NullableInlineResponse20016Capabilities) Set(val *InlineResponse20016Capabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016Capabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016Capabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016Capabilities(val *InlineResponse20016Capabilities) *NullableInlineResponse20016Capabilities {
	return &NullableInlineResponse20016Capabilities{value: val, isSet: true}
}

func (v NullableInlineResponse20016Capabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016Capabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

