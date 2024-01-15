/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner{}

// GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner struct for GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner
type GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner struct {
	// The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
	Dscp *int32 `json:"dscp,omitempty"`
	// The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
	Cos *int32 `json:"cos,omitempty"`
	// Label for the mapping (optional).
	Title *string `json:"title,omitempty"`
}

// NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner instantiates a new GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner() *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner {
	this := GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner{}
	return &this
}

// NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInnerWithDefaults instantiates a new GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInnerWithDefaults() *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner {
	this := GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner{}
	return &this
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetDscp() int32 {
	if o == nil || IsNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetDscpOk() (*int32, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) SetDscp(v int32) {
	o.Dscp = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetCos() int32 {
	if o == nil || IsNil(o.Cos) {
		var ret int32
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetCosOk() (*int32, bool) {
	if o == nil || IsNil(o.Cos) {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) HasCos() bool {
	if o != nil && !IsNil(o.Cos) {
		return true
	}

	return false
}

// SetCos gets a reference to the given int32 and assigns it to the Cos field.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) SetCos(v int32) {
	o.Cos = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) SetTitle(v string) {
	o.Title = &v
}

func (o GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !IsNil(o.Cos) {
		toSerialize["cos"] = o.Cos
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner struct {
	value *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner
	isSet bool
}

func (v NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) Get() *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner {
	return v.value
}

func (v *NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) Set(val *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner(val *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) *NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner {
	return &NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


