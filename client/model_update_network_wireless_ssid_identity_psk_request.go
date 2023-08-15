/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the UpdateNetworkWirelessSsidIdentityPskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidIdentityPskRequest{}

// UpdateNetworkWirelessSsidIdentityPskRequest struct for UpdateNetworkWirelessSsidIdentityPskRequest
type UpdateNetworkWirelessSsidIdentityPskRequest struct {
	// The name of the Identity PSK
	Name *string `json:"name,omitempty"`
	// The passphrase for client authentication
	Passphrase *string `json:"passphrase,omitempty"`
	// The group policy to be applied to clients
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// Timestamp for when the Identity PSK expires, or 'null' to never expire
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewUpdateNetworkWirelessSsidIdentityPskRequest instantiates a new UpdateNetworkWirelessSsidIdentityPskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidIdentityPskRequest() *UpdateNetworkWirelessSsidIdentityPskRequest {
	this := UpdateNetworkWirelessSsidIdentityPskRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidIdentityPskRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidIdentityPskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidIdentityPskRequestWithDefaults() *UpdateNetworkWirelessSsidIdentityPskRequest {
	this := UpdateNetworkWirelessSsidIdentityPskRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) SetName(v string) {
	o.Name = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetGroupPolicyId() string {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *UpdateNetworkWirelessSsidIdentityPskRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o UpdateNetworkWirelessSsidIdentityPskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidIdentityPskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidIdentityPskRequest struct {
	value *UpdateNetworkWirelessSsidIdentityPskRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidIdentityPskRequest) Get() *UpdateNetworkWirelessSsidIdentityPskRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidIdentityPskRequest) Set(val *UpdateNetworkWirelessSsidIdentityPskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidIdentityPskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidIdentityPskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidIdentityPskRequest(val *UpdateNetworkWirelessSsidIdentityPskRequest) *NullableUpdateNetworkWirelessSsidIdentityPskRequest {
	return &NullableUpdateNetworkWirelessSsidIdentityPskRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidIdentityPskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidIdentityPskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


