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

// InlineObject28 struct for InlineObject28
type InlineObject28 struct {
	// The list of connectivity monitoring destinations
	Destinations []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"`
}

// NewInlineObject28 instantiates a new InlineObject28 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject28() *InlineObject28 {
	this := InlineObject28{}
	return &this
}

// NewInlineObject28WithDefaults instantiates a new InlineObject28 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject28WithDefaults() *InlineObject28 {
	this := InlineObject28{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *InlineObject28) GetDestinations() []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations {
	if o == nil || isNil(o.Destinations) {
		var ret []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject28) GetDestinationsOk() ([]NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations, bool) {
	if o == nil || isNil(o.Destinations) {
    return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *InlineObject28) HasDestinations() bool {
	if o != nil && !isNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations and assigns it to the Destinations field.
func (o *InlineObject28) SetDestinations(v []NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) {
	o.Destinations = v
}

func (o InlineObject28) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject28 struct {
	value *InlineObject28
	isSet bool
}

func (v NullableInlineObject28) Get() *InlineObject28 {
	return v.value
}

func (v *NullableInlineObject28) Set(val *InlineObject28) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject28) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject28) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject28(val *InlineObject28) *NullableInlineObject28 {
	return &NullableInlineObject28{value: val, isSet: true}
}

func (v NullableInlineObject28) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject28) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

