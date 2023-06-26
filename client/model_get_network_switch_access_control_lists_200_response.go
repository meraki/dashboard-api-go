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

// GetNetworkSwitchAccessControlLists200Response struct for GetNetworkSwitchAccessControlLists200Response
type GetNetworkSwitchAccessControlLists200Response struct {
	// An ordered array of the access control list rules
	Rules []GetNetworkSwitchAccessControlLists200ResponseRulesInner `json:"rules,omitempty"`
}

// NewGetNetworkSwitchAccessControlLists200Response instantiates a new GetNetworkSwitchAccessControlLists200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAccessControlLists200Response() *GetNetworkSwitchAccessControlLists200Response {
	this := GetNetworkSwitchAccessControlLists200Response{}
	return &this
}

// NewGetNetworkSwitchAccessControlLists200ResponseWithDefaults instantiates a new GetNetworkSwitchAccessControlLists200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAccessControlLists200ResponseWithDefaults() *GetNetworkSwitchAccessControlLists200Response {
	this := GetNetworkSwitchAccessControlLists200Response{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessControlLists200Response) GetRules() []GetNetworkSwitchAccessControlLists200ResponseRulesInner {
	if o == nil || isNil(o.Rules) {
		var ret []GetNetworkSwitchAccessControlLists200ResponseRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessControlLists200Response) GetRulesOk() ([]GetNetworkSwitchAccessControlLists200ResponseRulesInner, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessControlLists200Response) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetNetworkSwitchAccessControlLists200ResponseRulesInner and assigns it to the Rules field.
func (o *GetNetworkSwitchAccessControlLists200Response) SetRules(v []GetNetworkSwitchAccessControlLists200ResponseRulesInner) {
	o.Rules = v
}

func (o GetNetworkSwitchAccessControlLists200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchAccessControlLists200Response struct {
	value *GetNetworkSwitchAccessControlLists200Response
	isSet bool
}

func (v NullableGetNetworkSwitchAccessControlLists200Response) Get() *GetNetworkSwitchAccessControlLists200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchAccessControlLists200Response) Set(val *GetNetworkSwitchAccessControlLists200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAccessControlLists200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAccessControlLists200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAccessControlLists200Response(val *GetNetworkSwitchAccessControlLists200Response) *NullableGetNetworkSwitchAccessControlLists200Response {
	return &NullableGetNetworkSwitchAccessControlLists200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAccessControlLists200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAccessControlLists200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


