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

// GetNetworkApplianceRfProfiles200Response struct for GetNetworkApplianceRfProfiles200Response
type GetNetworkApplianceRfProfiles200Response struct {
	// RF Profiles
	Assigned []GetNetworkApplianceRfProfiles200ResponseAssignedInner `json:"assigned,omitempty"`
}

// NewGetNetworkApplianceRfProfiles200Response instantiates a new GetNetworkApplianceRfProfiles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceRfProfiles200Response() *GetNetworkApplianceRfProfiles200Response {
	this := GetNetworkApplianceRfProfiles200Response{}
	return &this
}

// NewGetNetworkApplianceRfProfiles200ResponseWithDefaults instantiates a new GetNetworkApplianceRfProfiles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceRfProfiles200ResponseWithDefaults() *GetNetworkApplianceRfProfiles200Response {
	this := GetNetworkApplianceRfProfiles200Response{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200Response) GetAssigned() []GetNetworkApplianceRfProfiles200ResponseAssignedInner {
	if o == nil || isNil(o.Assigned) {
		var ret []GetNetworkApplianceRfProfiles200ResponseAssignedInner
		return ret
	}
	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200Response) GetAssignedOk() ([]GetNetworkApplianceRfProfiles200ResponseAssignedInner, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200Response) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given []GetNetworkApplianceRfProfiles200ResponseAssignedInner and assigns it to the Assigned field.
func (o *GetNetworkApplianceRfProfiles200Response) SetAssigned(v []GetNetworkApplianceRfProfiles200ResponseAssignedInner) {
	o.Assigned = v
}

func (o GetNetworkApplianceRfProfiles200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkApplianceRfProfiles200Response struct {
	value *GetNetworkApplianceRfProfiles200Response
	isSet bool
}

func (v NullableGetNetworkApplianceRfProfiles200Response) Get() *GetNetworkApplianceRfProfiles200Response {
	return v.value
}

func (v *NullableGetNetworkApplianceRfProfiles200Response) Set(val *GetNetworkApplianceRfProfiles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceRfProfiles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceRfProfiles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceRfProfiles200Response(val *GetNetworkApplianceRfProfiles200Response) *NullableGetNetworkApplianceRfProfiles200Response {
	return &NullableGetNetworkApplianceRfProfiles200Response{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceRfProfiles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceRfProfiles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

