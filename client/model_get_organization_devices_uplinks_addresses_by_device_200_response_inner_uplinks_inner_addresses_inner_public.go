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

// checks if the GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic{}

// GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic Public interface information.
type GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic struct {
	// The device uplink public IP address.
	Address *string `json:"address,omitempty"`
}

// NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic {
	this := GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic{}
	return &this
}

// NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublicWithDefaults instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublicWithDefaults() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic {
	this := GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) SetAddress(v string) {
	o.Address = &v
}

func (o GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic struct {
	value *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic
	isSet bool
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) Get() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic {
	return v.value
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) Set(val *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic(val *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic {
	return &NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


