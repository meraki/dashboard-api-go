/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SetNetworkWirelessEthernetPortsProfilesDefault200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetNetworkWirelessEthernetPortsProfilesDefault200Response{}

// SetNetworkWirelessEthernetPortsProfilesDefault200Response struct for SetNetworkWirelessEthernetPortsProfilesDefault200Response
type SetNetworkWirelessEthernetPortsProfilesDefault200Response struct {
	// AP profile ID
	ProfileId *string `json:"profileId,omitempty"`
}

// NewSetNetworkWirelessEthernetPortsProfilesDefault200Response instantiates a new SetNetworkWirelessEthernetPortsProfilesDefault200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetNetworkWirelessEthernetPortsProfilesDefault200Response() *SetNetworkWirelessEthernetPortsProfilesDefault200Response {
	this := SetNetworkWirelessEthernetPortsProfilesDefault200Response{}
	return &this
}

// NewSetNetworkWirelessEthernetPortsProfilesDefault200ResponseWithDefaults instantiates a new SetNetworkWirelessEthernetPortsProfilesDefault200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetNetworkWirelessEthernetPortsProfilesDefault200ResponseWithDefaults() *SetNetworkWirelessEthernetPortsProfilesDefault200Response {
	this := SetNetworkWirelessEthernetPortsProfilesDefault200Response{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *SetNetworkWirelessEthernetPortsProfilesDefault200Response) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetNetworkWirelessEthernetPortsProfilesDefault200Response) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *SetNetworkWirelessEthernetPortsProfilesDefault200Response) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *SetNetworkWirelessEthernetPortsProfilesDefault200Response) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o SetNetworkWirelessEthernetPortsProfilesDefault200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetNetworkWirelessEthernetPortsProfilesDefault200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

type NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response struct {
	value *SetNetworkWirelessEthernetPortsProfilesDefault200Response
	isSet bool
}

func (v NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response) Get() *SetNetworkWirelessEthernetPortsProfilesDefault200Response {
	return v.value
}

func (v *NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response) Set(val *SetNetworkWirelessEthernetPortsProfilesDefault200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetNetworkWirelessEthernetPortsProfilesDefault200Response(val *SetNetworkWirelessEthernetPortsProfilesDefault200Response) *NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response {
	return &NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response{value: val, isSet: true}
}

func (v NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetNetworkWirelessEthernetPortsProfilesDefault200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

