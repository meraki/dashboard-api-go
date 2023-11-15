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

// checks if the ClaimIntoOrganizationRequestLicensesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimIntoOrganizationRequestLicensesInner{}

// ClaimIntoOrganizationRequestLicensesInner struct for ClaimIntoOrganizationRequestLicensesInner
type ClaimIntoOrganizationRequestLicensesInner struct {
	// The key of the license
	Key string `json:"key"`
	// Either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. This parameter is legacy and does not apply to organizations with per-device licensing enabled.
	Mode *string `json:"mode,omitempty"`
}

// NewClaimIntoOrganizationRequestLicensesInner instantiates a new ClaimIntoOrganizationRequestLicensesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimIntoOrganizationRequestLicensesInner(key string) *ClaimIntoOrganizationRequestLicensesInner {
	this := ClaimIntoOrganizationRequestLicensesInner{}
	this.Key = key
	return &this
}

// NewClaimIntoOrganizationRequestLicensesInnerWithDefaults instantiates a new ClaimIntoOrganizationRequestLicensesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimIntoOrganizationRequestLicensesInnerWithDefaults() *ClaimIntoOrganizationRequestLicensesInner {
	this := ClaimIntoOrganizationRequestLicensesInner{}
	return &this
}

// GetKey returns the Key field value
func (o *ClaimIntoOrganizationRequestLicensesInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationRequestLicensesInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ClaimIntoOrganizationRequestLicensesInner) SetKey(v string) {
	o.Key = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationRequestLicensesInner) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationRequestLicensesInner) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationRequestLicensesInner) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ClaimIntoOrganizationRequestLicensesInner) SetMode(v string) {
	o.Mode = &v
}

func (o ClaimIntoOrganizationRequestLicensesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimIntoOrganizationRequestLicensesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableClaimIntoOrganizationRequestLicensesInner struct {
	value *ClaimIntoOrganizationRequestLicensesInner
	isSet bool
}

func (v NullableClaimIntoOrganizationRequestLicensesInner) Get() *ClaimIntoOrganizationRequestLicensesInner {
	return v.value
}

func (v *NullableClaimIntoOrganizationRequestLicensesInner) Set(val *ClaimIntoOrganizationRequestLicensesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimIntoOrganizationRequestLicensesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimIntoOrganizationRequestLicensesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimIntoOrganizationRequestLicensesInner(val *ClaimIntoOrganizationRequestLicensesInner) *NullableClaimIntoOrganizationRequestLicensesInner {
	return &NullableClaimIntoOrganizationRequestLicensesInner{value: val, isSet: true}
}

func (v NullableClaimIntoOrganizationRequestLicensesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimIntoOrganizationRequestLicensesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


