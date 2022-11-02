/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner struct for UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner
type UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner struct {
	// The starting IP address of the reserved IP range
	Start string `json:"start"`
	// The ending IP address of the reserved IP range
	End string `json:"end"`
	// The comment for the reserved IP range
	Comment *string `json:"comment,omitempty"`
}

// NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner(start string, end string) *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner {
	this := UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner{}
	this.Start = start
	this.End = end
	return &this
}

// NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInnerWithDefaults instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInnerWithDefaults() *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner {
	this := UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner{}
	return &this
}

// GetStart returns the Start field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) GetStartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) GetEndOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) SetEnd(v string) {
	o.End = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) SetComment(v string) {
	o.Comment = &v
}

func (o UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner struct {
	value *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner
	isSet bool
}

func (v NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) Get() *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner {
	return v.value
}

func (v *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) Set(val *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner(val *UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner {
	return &NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


