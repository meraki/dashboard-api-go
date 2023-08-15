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

// checks if the GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork{}

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork Network details object
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork struct {
	// The network ID the AP is associated to
	Id *string `json:"id,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetworkWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetworkWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) SetId(v string) {
	o.Id = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


