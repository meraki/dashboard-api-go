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

// checks if the UpdateNetworkWirelessSsidRequestLocalRadius type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestLocalRadius{}

// UpdateNetworkWirelessSsidRequestLocalRadius The current setting for Local Authentication, a built-in RADIUS server on the access point. Only valid if authMode is '8021x-localradius'.
type UpdateNetworkWirelessSsidRequestLocalRadius struct {
	// The duration (in seconds) for which LDAP and OCSP lookups are cached.
	CacheTimeout *int32 `json:"cacheTimeout,omitempty"`
	PasswordAuthentication *UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication `json:"passwordAuthentication,omitempty"`
	CertificateAuthentication *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication `json:"certificateAuthentication,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestLocalRadius instantiates a new UpdateNetworkWirelessSsidRequestLocalRadius object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestLocalRadius() *UpdateNetworkWirelessSsidRequestLocalRadius {
	this := UpdateNetworkWirelessSsidRequestLocalRadius{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestLocalRadiusWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLocalRadius object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestLocalRadiusWithDefaults() *UpdateNetworkWirelessSsidRequestLocalRadius {
	this := UpdateNetworkWirelessSsidRequestLocalRadius{}
	return &this
}

// GetCacheTimeout returns the CacheTimeout field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCacheTimeout() int32 {
	if o == nil || IsNil(o.CacheTimeout) {
		var ret int32
		return ret
	}
	return *o.CacheTimeout
}

// GetCacheTimeoutOk returns a tuple with the CacheTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCacheTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheTimeout) {
		return nil, false
	}
	return o.CacheTimeout, true
}

// HasCacheTimeout returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) HasCacheTimeout() bool {
	if o != nil && !IsNil(o.CacheTimeout) {
		return true
	}

	return false
}

// SetCacheTimeout gets a reference to the given int32 and assigns it to the CacheTimeout field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) SetCacheTimeout(v int32) {
	o.CacheTimeout = &v
}

// GetPasswordAuthentication returns the PasswordAuthentication field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetPasswordAuthentication() UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication {
	if o == nil || IsNil(o.PasswordAuthentication) {
		var ret UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication
		return ret
	}
	return *o.PasswordAuthentication
}

// GetPasswordAuthenticationOk returns a tuple with the PasswordAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetPasswordAuthenticationOk() (*UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication, bool) {
	if o == nil || IsNil(o.PasswordAuthentication) {
		return nil, false
	}
	return o.PasswordAuthentication, true
}

// HasPasswordAuthentication returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) HasPasswordAuthentication() bool {
	if o != nil && !IsNil(o.PasswordAuthentication) {
		return true
	}

	return false
}

// SetPasswordAuthentication gets a reference to the given UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication and assigns it to the PasswordAuthentication field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) SetPasswordAuthentication(v UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication) {
	o.PasswordAuthentication = &v
}

// GetCertificateAuthentication returns the CertificateAuthentication field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCertificateAuthentication() UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication {
	if o == nil || IsNil(o.CertificateAuthentication) {
		var ret UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication
		return ret
	}
	return *o.CertificateAuthentication
}

// GetCertificateAuthenticationOk returns a tuple with the CertificateAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCertificateAuthenticationOk() (*UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication, bool) {
	if o == nil || IsNil(o.CertificateAuthentication) {
		return nil, false
	}
	return o.CertificateAuthentication, true
}

// HasCertificateAuthentication returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) HasCertificateAuthentication() bool {
	if o != nil && !IsNil(o.CertificateAuthentication) {
		return true
	}

	return false
}

// SetCertificateAuthentication gets a reference to the given UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication and assigns it to the CertificateAuthentication field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadius) SetCertificateAuthentication(v UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) {
	o.CertificateAuthentication = &v
}

func (o UpdateNetworkWirelessSsidRequestLocalRadius) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestLocalRadius) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CacheTimeout) {
		toSerialize["cacheTimeout"] = o.CacheTimeout
	}
	if !IsNil(o.PasswordAuthentication) {
		toSerialize["passwordAuthentication"] = o.PasswordAuthentication
	}
	if !IsNil(o.CertificateAuthentication) {
		toSerialize["certificateAuthentication"] = o.CertificateAuthentication
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestLocalRadius struct {
	value *UpdateNetworkWirelessSsidRequestLocalRadius
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadius) Get() *UpdateNetworkWirelessSsidRequestLocalRadius {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadius) Set(val *UpdateNetworkWirelessSsidRequestLocalRadius) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadius) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadius) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestLocalRadius(val *UpdateNetworkWirelessSsidRequestLocalRadius) *NullableUpdateNetworkWirelessSsidRequestLocalRadius {
	return &NullableUpdateNetworkWirelessSsidRequestLocalRadius{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadius) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadius) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


