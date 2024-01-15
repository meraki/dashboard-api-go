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

// checks if the GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday{}

// GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday Wednesday schedule
type GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday struct {
	// Whether the schedule is active or inactive
	Active *bool `json:"active,omitempty"`
	// The time, from '00:00' to '24:00'
	From *string `json:"from,omitempty"`
	// The time, from '00:00' to '24:00'
	To *string `json:"to,omitempty"`
}

// NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday instantiates a new GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday() *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday {
	this := GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday{}
	return &this
}

// NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesdayWithDefaults instantiates a new GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesdayWithDefaults() *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday {
	this := GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) SetActive(v bool) {
	o.Active = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) SetTo(v string) {
	o.To = &v
}

func (o GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday struct {
	value *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday
	isSet bool
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) Get() *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday {
	return v.value
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) Set(val *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday(val *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday {
	return &NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


