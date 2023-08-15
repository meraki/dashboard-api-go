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

// checks if the ClaimIntoOrganizationInventoryRequestLicensesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimIntoOrganizationInventoryRequestLicensesInner{}

// ClaimIntoOrganizationInventoryRequestLicensesInner struct for ClaimIntoOrganizationInventoryRequestLicensesInner
type ClaimIntoOrganizationInventoryRequestLicensesInner struct {
	// The key of the license
	Key string `json:"key"`
	// Co-term licensing only: either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. Does not apply to organizations using per-device licensing model.
	Mode *string `json:"mode,omitempty"`
}

// NewClaimIntoOrganizationInventoryRequestLicensesInner instantiates a new ClaimIntoOrganizationInventoryRequestLicensesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimIntoOrganizationInventoryRequestLicensesInner(key string) *ClaimIntoOrganizationInventoryRequestLicensesInner {
	this := ClaimIntoOrganizationInventoryRequestLicensesInner{}
	this.Key = key
	return &this
}

// NewClaimIntoOrganizationInventoryRequestLicensesInnerWithDefaults instantiates a new ClaimIntoOrganizationInventoryRequestLicensesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimIntoOrganizationInventoryRequestLicensesInnerWithDefaults() *ClaimIntoOrganizationInventoryRequestLicensesInner {
	this := ClaimIntoOrganizationInventoryRequestLicensesInner{}
	return &this
}

// GetKey returns the Key field value
func (o *ClaimIntoOrganizationInventoryRequestLicensesInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationInventoryRequestLicensesInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ClaimIntoOrganizationInventoryRequestLicensesInner) SetKey(v string) {
	o.Key = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationInventoryRequestLicensesInner) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationInventoryRequestLicensesInner) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationInventoryRequestLicensesInner) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ClaimIntoOrganizationInventoryRequestLicensesInner) SetMode(v string) {
	o.Mode = &v
}

func (o ClaimIntoOrganizationInventoryRequestLicensesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimIntoOrganizationInventoryRequestLicensesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableClaimIntoOrganizationInventoryRequestLicensesInner struct {
	value *ClaimIntoOrganizationInventoryRequestLicensesInner
	isSet bool
}

func (v NullableClaimIntoOrganizationInventoryRequestLicensesInner) Get() *ClaimIntoOrganizationInventoryRequestLicensesInner {
	return v.value
}

func (v *NullableClaimIntoOrganizationInventoryRequestLicensesInner) Set(val *ClaimIntoOrganizationInventoryRequestLicensesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimIntoOrganizationInventoryRequestLicensesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimIntoOrganizationInventoryRequestLicensesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimIntoOrganizationInventoryRequestLicensesInner(val *ClaimIntoOrganizationInventoryRequestLicensesInner) *NullableClaimIntoOrganizationInventoryRequestLicensesInner {
	return &NullableClaimIntoOrganizationInventoryRequestLicensesInner{value: val, isSet: true}
}

func (v NullableClaimIntoOrganizationInventoryRequestLicensesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimIntoOrganizationInventoryRequestLicensesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


