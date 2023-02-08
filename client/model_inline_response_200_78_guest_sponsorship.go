/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20078GuestSponsorship Details associated with guest sponsored splash
type InlineResponse20078GuestSponsorship struct {
	// Duration in minutes of sponsored guest authorization.
	DurationInMinutes *int32 `json:"durationInMinutes,omitempty"`
	// Whether or not guests can specify how much time they are requesting.
	GuestCanRequestTimeframe *bool `json:"guestCanRequestTimeframe,omitempty"`
}

// NewInlineResponse20078GuestSponsorship instantiates a new InlineResponse20078GuestSponsorship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20078GuestSponsorship() *InlineResponse20078GuestSponsorship {
	this := InlineResponse20078GuestSponsorship{}
	return &this
}

// NewInlineResponse20078GuestSponsorshipWithDefaults instantiates a new InlineResponse20078GuestSponsorship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20078GuestSponsorshipWithDefaults() *InlineResponse20078GuestSponsorship {
	this := InlineResponse20078GuestSponsorship{}
	return &this
}

// GetDurationInMinutes returns the DurationInMinutes field value if set, zero value otherwise.
func (o *InlineResponse20078GuestSponsorship) GetDurationInMinutes() int32 {
	if o == nil || isNil(o.DurationInMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationInMinutes
}

// GetDurationInMinutesOk returns a tuple with the DurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20078GuestSponsorship) GetDurationInMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.DurationInMinutes) {
    return nil, false
	}
	return o.DurationInMinutes, true
}

// HasDurationInMinutes returns a boolean if a field has been set.
func (o *InlineResponse20078GuestSponsorship) HasDurationInMinutes() bool {
	if o != nil && !isNil(o.DurationInMinutes) {
		return true
	}

	return false
}

// SetDurationInMinutes gets a reference to the given int32 and assigns it to the DurationInMinutes field.
func (o *InlineResponse20078GuestSponsorship) SetDurationInMinutes(v int32) {
	o.DurationInMinutes = &v
}

// GetGuestCanRequestTimeframe returns the GuestCanRequestTimeframe field value if set, zero value otherwise.
func (o *InlineResponse20078GuestSponsorship) GetGuestCanRequestTimeframe() bool {
	if o == nil || isNil(o.GuestCanRequestTimeframe) {
		var ret bool
		return ret
	}
	return *o.GuestCanRequestTimeframe
}

// GetGuestCanRequestTimeframeOk returns a tuple with the GuestCanRequestTimeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20078GuestSponsorship) GetGuestCanRequestTimeframeOk() (*bool, bool) {
	if o == nil || isNil(o.GuestCanRequestTimeframe) {
    return nil, false
	}
	return o.GuestCanRequestTimeframe, true
}

// HasGuestCanRequestTimeframe returns a boolean if a field has been set.
func (o *InlineResponse20078GuestSponsorship) HasGuestCanRequestTimeframe() bool {
	if o != nil && !isNil(o.GuestCanRequestTimeframe) {
		return true
	}

	return false
}

// SetGuestCanRequestTimeframe gets a reference to the given bool and assigns it to the GuestCanRequestTimeframe field.
func (o *InlineResponse20078GuestSponsorship) SetGuestCanRequestTimeframe(v bool) {
	o.GuestCanRequestTimeframe = &v
}

func (o InlineResponse20078GuestSponsorship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DurationInMinutes) {
		toSerialize["durationInMinutes"] = o.DurationInMinutes
	}
	if !isNil(o.GuestCanRequestTimeframe) {
		toSerialize["guestCanRequestTimeframe"] = o.GuestCanRequestTimeframe
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20078GuestSponsorship struct {
	value *InlineResponse20078GuestSponsorship
	isSet bool
}

func (v NullableInlineResponse20078GuestSponsorship) Get() *InlineResponse20078GuestSponsorship {
	return v.value
}

func (v *NullableInlineResponse20078GuestSponsorship) Set(val *InlineResponse20078GuestSponsorship) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20078GuestSponsorship) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20078GuestSponsorship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20078GuestSponsorship(val *InlineResponse20078GuestSponsorship) *NullableInlineResponse20078GuestSponsorship {
	return &NullableInlineResponse20078GuestSponsorship{value: val, isSet: true}
}

func (v NullableInlineResponse20078GuestSponsorship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20078GuestSponsorship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

