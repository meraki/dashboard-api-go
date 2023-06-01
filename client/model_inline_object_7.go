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

// InlineObject7 struct for InlineObject7
type InlineObject7 struct {
	// Boolean indicating if external rtsp stream is exposed
	ExternalRtspEnabled *bool `json:"externalRtspEnabled,omitempty"`
}

// NewInlineObject7 instantiates a new InlineObject7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject7() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// NewInlineObject7WithDefaults instantiates a new InlineObject7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject7WithDefaults() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// GetExternalRtspEnabled returns the ExternalRtspEnabled field value if set, zero value otherwise.
func (o *InlineObject7) GetExternalRtspEnabled() bool {
	if o == nil || isNil(o.ExternalRtspEnabled) {
		var ret bool
		return ret
	}
	return *o.ExternalRtspEnabled
}

// GetExternalRtspEnabledOk returns a tuple with the ExternalRtspEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetExternalRtspEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ExternalRtspEnabled) {
    return nil, false
	}
	return o.ExternalRtspEnabled, true
}

// HasExternalRtspEnabled returns a boolean if a field has been set.
func (o *InlineObject7) HasExternalRtspEnabled() bool {
	if o != nil && !isNil(o.ExternalRtspEnabled) {
		return true
	}

	return false
}

// SetExternalRtspEnabled gets a reference to the given bool and assigns it to the ExternalRtspEnabled field.
func (o *InlineObject7) SetExternalRtspEnabled(v bool) {
	o.ExternalRtspEnabled = &v
}

func (o InlineObject7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalRtspEnabled) {
		toSerialize["externalRtspEnabled"] = o.ExternalRtspEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject7 struct {
	value *InlineObject7
	isSet bool
}

func (v NullableInlineObject7) Get() *InlineObject7 {
	return v.value
}

func (v *NullableInlineObject7) Set(val *InlineObject7) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject7) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject7(val *InlineObject7) *NullableInlineObject7 {
	return &NullableInlineObject7{value: val, isSet: true}
}

func (v NullableInlineObject7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

