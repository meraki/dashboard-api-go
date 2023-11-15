/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys{}

// GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys Details for API-only IP restrictions.
type GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys struct {
	// Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
	Enabled *bool `json:"enabled,omitempty"`
	// List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	Ranges []string `json:"ranges,omitempty"`
}

// NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys instantiates a new GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys() *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys {
	this := GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys{}
	return &this
}

// NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeysWithDefaults instantiates a new GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeysWithDefaults() *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys {
	this := GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetRanges() []string {
	if o == nil || IsNil(o.Ranges) {
		var ret []string
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) HasRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []string and assigns it to the Ranges field.
func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) SetRanges(v []string) {
	o.Ranges = v
}

func (o GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	return toSerialize, nil
}

type NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys struct {
	value *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys
	isSet bool
}

func (v NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) Get() *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys {
	return v.value
}

func (v *NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) Set(val *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys(val *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) *NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys {
	return &NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys{value: val, isSet: true}
}

func (v NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


