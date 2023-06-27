/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkApplianceSsidRequestRadiusServersInner struct for UpdateNetworkApplianceSsidRequestRadiusServersInner
type UpdateNetworkApplianceSsidRequestRadiusServersInner struct {
	// The IP address of your RADIUS server.
	Host *string `json:"host,omitempty"`
	// The UDP port your RADIUS servers listens on for Access-requests.
	Port *int32 `json:"port,omitempty"`
	// The RADIUS client shared secret.
	Secret *string `json:"secret,omitempty"`
}

// NewUpdateNetworkApplianceSsidRequestRadiusServersInner instantiates a new UpdateNetworkApplianceSsidRequestRadiusServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSsidRequestRadiusServersInner() *UpdateNetworkApplianceSsidRequestRadiusServersInner {
	this := UpdateNetworkApplianceSsidRequestRadiusServersInner{}
	return &this
}

// NewUpdateNetworkApplianceSsidRequestRadiusServersInnerWithDefaults instantiates a new UpdateNetworkApplianceSsidRequestRadiusServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSsidRequestRadiusServersInnerWithDefaults() *UpdateNetworkApplianceSsidRequestRadiusServersInner {
	this := UpdateNetworkApplianceSsidRequestRadiusServersInner{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) SetSecret(v string) {
	o.Secret = &v
}

func (o UpdateNetworkApplianceSsidRequestRadiusServersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceSsidRequestRadiusServersInner struct {
	value *UpdateNetworkApplianceSsidRequestRadiusServersInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceSsidRequestRadiusServersInner) Get() *UpdateNetworkApplianceSsidRequestRadiusServersInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSsidRequestRadiusServersInner) Set(val *UpdateNetworkApplianceSsidRequestRadiusServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSsidRequestRadiusServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSsidRequestRadiusServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSsidRequestRadiusServersInner(val *UpdateNetworkApplianceSsidRequestRadiusServersInner) *NullableUpdateNetworkApplianceSsidRequestRadiusServersInner {
	return &NullableUpdateNetworkApplianceSsidRequestRadiusServersInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSsidRequestRadiusServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSsidRequestRadiusServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


