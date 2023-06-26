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

// GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner struct for GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner
type GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner struct {
	// The first IP in the reserved range
	Start *string `json:"start,omitempty"`
	// The last IP in the reserved range
	End *string `json:"end,omitempty"`
	// A text comment for the reserved range
	Comment *string `json:"comment,omitempty"`
}

// NewGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner instantiates a new GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner() *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner {
	this := GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner{}
	return &this
}

// NewGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInnerWithDefaults instantiates a new GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInnerWithDefaults() *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner {
	this := GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) SetEnd(v string) {
	o.End = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) SetComment(v string) {
	o.Comment = &v
}

func (o GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner struct {
	value *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner
	isSet bool
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) Get() *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner {
	return v.value
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) Set(val *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner(val *GetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) *NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner {
	return &NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerReservedIpRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


