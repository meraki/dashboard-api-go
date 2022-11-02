/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkWirelessRfProfileRequestTransmission Settings related to radio transmission.
type CreateNetworkWirelessRfProfileRequestTransmission struct {
	// Toggle for radio transmission. When false, radios will not transmit at all.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestTransmission instantiates a new CreateNetworkWirelessRfProfileRequestTransmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestTransmission() *CreateNetworkWirelessRfProfileRequestTransmission {
	this := CreateNetworkWirelessRfProfileRequestTransmission{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestTransmissionWithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestTransmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestTransmissionWithDefaults() *CreateNetworkWirelessRfProfileRequestTransmission {
	this := CreateNetworkWirelessRfProfileRequestTransmission{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestTransmission) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestTransmission) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestTransmission) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateNetworkWirelessRfProfileRequestTransmission) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CreateNetworkWirelessRfProfileRequestTransmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkWirelessRfProfileRequestTransmission struct {
	value *CreateNetworkWirelessRfProfileRequestTransmission
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestTransmission) Get() *CreateNetworkWirelessRfProfileRequestTransmission {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestTransmission) Set(val *CreateNetworkWirelessRfProfileRequestTransmission) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestTransmission) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestTransmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestTransmission(val *CreateNetworkWirelessRfProfileRequestTransmission) *NullableCreateNetworkWirelessRfProfileRequestTransmission {
	return &NullableCreateNetworkWirelessRfProfileRequestTransmission{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestTransmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestTransmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


