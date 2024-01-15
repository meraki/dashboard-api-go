/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkAppliancePortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkAppliancePortRequest{}

// UpdateNetworkAppliancePortRequest struct for UpdateNetworkAppliancePortRequest
type UpdateNetworkAppliancePortRequest struct {
	// The status of the port
	Enabled *bool `json:"enabled,omitempty"`
	// Trunk port can Drop all Untagged traffic. When true, no VLAN is required. Access ports cannot have dropUntaggedTraffic set to true.
	DropUntaggedTraffic *bool `json:"dropUntaggedTraffic,omitempty"`
	// The type of the port: 'access' or 'trunk'.
	Type *string `json:"type,omitempty"`
	// Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
	Vlan *int32 `json:"vlan,omitempty"`
	// Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The name of the policy. Only applicable to Access ports. Valid values are: 'open', '8021x-radius', 'mac-radius', 'hybris-radius' for MX64 or Z3 or any MX supporting the per port authentication feature. Otherwise, 'open' is the only valid value and 'open' is the default value if the field is missing.
	AccessPolicy *string `json:"accessPolicy,omitempty"`
}

// NewUpdateNetworkAppliancePortRequest instantiates a new UpdateNetworkAppliancePortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkAppliancePortRequest() *UpdateNetworkAppliancePortRequest {
	this := UpdateNetworkAppliancePortRequest{}
	return &this
}

// NewUpdateNetworkAppliancePortRequestWithDefaults instantiates a new UpdateNetworkAppliancePortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkAppliancePortRequestWithDefaults() *UpdateNetworkAppliancePortRequest {
	this := UpdateNetworkAppliancePortRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePortRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePortRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePortRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkAppliancePortRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDropUntaggedTraffic returns the DropUntaggedTraffic field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePortRequest) GetDropUntaggedTraffic() bool {
	if o == nil || IsNil(o.DropUntaggedTraffic) {
		var ret bool
		return ret
	}
	return *o.DropUntaggedTraffic
}

// GetDropUntaggedTrafficOk returns a tuple with the DropUntaggedTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePortRequest) GetDropUntaggedTrafficOk() (*bool, bool) {
	if o == nil || IsNil(o.DropUntaggedTraffic) {
		return nil, false
	}
	return o.DropUntaggedTraffic, true
}

// HasDropUntaggedTraffic returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePortRequest) HasDropUntaggedTraffic() bool {
	if o != nil && !IsNil(o.DropUntaggedTraffic) {
		return true
	}

	return false
}

// SetDropUntaggedTraffic gets a reference to the given bool and assigns it to the DropUntaggedTraffic field.
func (o *UpdateNetworkAppliancePortRequest) SetDropUntaggedTraffic(v bool) {
	o.DropUntaggedTraffic = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePortRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePortRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePortRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateNetworkAppliancePortRequest) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePortRequest) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePortRequest) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePortRequest) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *UpdateNetworkAppliancePortRequest) SetVlan(v int32) {
	o.Vlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePortRequest) GetAllowedVlans() string {
	if o == nil || IsNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePortRequest) GetAllowedVlansOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedVlans) {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePortRequest) HasAllowedVlans() bool {
	if o != nil && !IsNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *UpdateNetworkAppliancePortRequest) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePortRequest) GetAccessPolicy() string {
	if o == nil || IsNil(o.AccessPolicy) {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePortRequest) GetAccessPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicy) {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePortRequest) HasAccessPolicy() bool {
	if o != nil && !IsNil(o.AccessPolicy) {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *UpdateNetworkAppliancePortRequest) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

func (o UpdateNetworkAppliancePortRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkAppliancePortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DropUntaggedTraffic) {
		toSerialize["dropUntaggedTraffic"] = o.DropUntaggedTraffic
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	if !IsNil(o.AccessPolicy) {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	return toSerialize, nil
}

type NullableUpdateNetworkAppliancePortRequest struct {
	value *UpdateNetworkAppliancePortRequest
	isSet bool
}

func (v NullableUpdateNetworkAppliancePortRequest) Get() *UpdateNetworkAppliancePortRequest {
	return v.value
}

func (v *NullableUpdateNetworkAppliancePortRequest) Set(val *UpdateNetworkAppliancePortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkAppliancePortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkAppliancePortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkAppliancePortRequest(val *UpdateNetworkAppliancePortRequest) *NullableUpdateNetworkAppliancePortRequest {
	return &NullableUpdateNetworkAppliancePortRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkAppliancePortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkAppliancePortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


