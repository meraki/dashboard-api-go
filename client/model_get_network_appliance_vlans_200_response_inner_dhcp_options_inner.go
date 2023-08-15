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

// checks if the GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner{}

// GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner struct for GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner
type GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner struct {
	// The code for the DHCP option. This should be an integer between 2 and 254.
	Code string `json:"code"`
	// The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	Type string `json:"type"`
	// The value for the DHCP option
	Value string `json:"value"`
}

// NewGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner instantiates a new GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner(code string, type_ string, value string) *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner {
	this := GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner{}
	this.Code = code
	this.Type = type_
	this.Value = value
	return &this
}

// NewGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInnerWithDefaults instantiates a new GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInnerWithDefaults() *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner {
	this := GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner{}
	return &this
}

// GetCode returns the Code field value
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) SetValue(v string) {
	o.Value = v
}

func (o GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner struct {
	value *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner
	isSet bool
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) Get() *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner {
	return v.value
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) Set(val *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner(val *GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) *NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner {
	return &NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


