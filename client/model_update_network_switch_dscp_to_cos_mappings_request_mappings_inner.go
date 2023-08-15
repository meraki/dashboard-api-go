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

// checks if the UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner{}

// UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner struct for UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner
type UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner struct {
	// The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
	Dscp int32 `json:"dscp"`
	// The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
	Cos int32 `json:"cos"`
	// Label for the mapping (optional).
	Title *string `json:"title,omitempty"`
}

// NewUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner instantiates a new UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner(dscp int32, cos int32) *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner {
	this := UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner{}
	this.Dscp = dscp
	this.Cos = cos
	return &this
}

// NewUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInnerWithDefaults instantiates a new UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInnerWithDefaults() *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner {
	this := UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner{}
	return &this
}

// GetDscp returns the Dscp field value
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) GetDscp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) GetDscpOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dscp, true
}

// SetDscp sets field value
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) SetDscp(v int32) {
	o.Dscp = v
}

// GetCos returns the Cos field value
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) GetCos() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cos
}

// GetCosOk returns a tuple with the Cos field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) GetCosOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cos, true
}

// SetCos sets field value
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) SetCos(v int32) {
	o.Cos = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) SetTitle(v string) {
	o.Title = &v
}

func (o UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dscp"] = o.Dscp
	toSerialize["cos"] = o.Cos
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner struct {
	value *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner
	isSet bool
}

func (v NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) Get() *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner {
	return v.value
}

func (v *NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) Set(val *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner(val *UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) *NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner {
	return &NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


