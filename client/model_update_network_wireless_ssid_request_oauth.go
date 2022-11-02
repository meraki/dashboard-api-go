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

// UpdateNetworkWirelessSsidRequestOauth The OAuth settings of this SSID. Only valid if splashPage is 'Google OAuth'.
type UpdateNetworkWirelessSsidRequestOauth struct {
	// (Optional) The list of domains allowed access to the network.
	AllowedDomains []string `json:"allowedDomains,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestOauth instantiates a new UpdateNetworkWirelessSsidRequestOauth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestOauth() *UpdateNetworkWirelessSsidRequestOauth {
	this := UpdateNetworkWirelessSsidRequestOauth{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestOauthWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestOauth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestOauthWithDefaults() *UpdateNetworkWirelessSsidRequestOauth {
	this := UpdateNetworkWirelessSsidRequestOauth{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestOauth) GetAllowedDomains() []string {
	if o == nil || isNil(o.AllowedDomains) {
		var ret []string
		return ret
	}
	return o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestOauth) GetAllowedDomainsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedDomains) {
    return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestOauth) HasAllowedDomains() bool {
	if o != nil && !isNil(o.AllowedDomains) {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *UpdateNetworkWirelessSsidRequestOauth) SetAllowedDomains(v []string) {
	o.AllowedDomains = v
}

func (o UpdateNetworkWirelessSsidRequestOauth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowedDomains) {
		toSerialize["allowedDomains"] = o.AllowedDomains
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidRequestOauth struct {
	value *UpdateNetworkWirelessSsidRequestOauth
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestOauth) Get() *UpdateNetworkWirelessSsidRequestOauth {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestOauth) Set(val *UpdateNetworkWirelessSsidRequestOauth) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestOauth) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestOauth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestOauth(val *UpdateNetworkWirelessSsidRequestOauth) *NullableUpdateNetworkWirelessSsidRequestOauth {
	return &NullableUpdateNetworkWirelessSsidRequestOauth{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestOauth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestOauth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


