/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidRequestActiveDirectory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestActiveDirectory{}

// UpdateNetworkWirelessSsidRequestActiveDirectory The current setting for Active Directory. Only valid if splashPage is 'Password-protected with Active Directory'
type UpdateNetworkWirelessSsidRequestActiveDirectory struct {
	// The Active Directory servers to be used for authentication.
	Servers []UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner `json:"servers,omitempty"`
	Credentials *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials `json:"credentials,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestActiveDirectory instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestActiveDirectory() *UpdateNetworkWirelessSsidRequestActiveDirectory {
	this := UpdateNetworkWirelessSsidRequestActiveDirectory{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestActiveDirectoryWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestActiveDirectoryWithDefaults() *UpdateNetworkWirelessSsidRequestActiveDirectory {
	this := UpdateNetworkWirelessSsidRequestActiveDirectory{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetServers() []UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner {
	if o == nil || IsNil(o.Servers) {
		var ret []UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetServersOk() ([]UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner and assigns it to the Servers field.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) SetServers(v []UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) {
	o.Servers = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetCredentials() UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetCredentialsOk() (*UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials and assigns it to the Credentials field.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) SetCredentials(v UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) {
	o.Credentials = &v
}

func (o UpdateNetworkWirelessSsidRequestActiveDirectory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestActiveDirectory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestActiveDirectory struct {
	value *UpdateNetworkWirelessSsidRequestActiveDirectory
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectory) Get() *UpdateNetworkWirelessSsidRequestActiveDirectory {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectory) Set(val *UpdateNetworkWirelessSsidRequestActiveDirectory) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectory) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestActiveDirectory(val *UpdateNetworkWirelessSsidRequestActiveDirectory) *NullableUpdateNetworkWirelessSsidRequestActiveDirectory {
	return &NullableUpdateNetworkWirelessSsidRequestActiveDirectory{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


