/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200100CountsByStatus byStatus
type InlineResponse200100CountsByStatus struct {
	// online count
	Online *int32 `json:"online,omitempty"`
	// alerting count
	Alerting *int32 `json:"alerting,omitempty"`
	// offline count
	Offline *int32 `json:"offline,omitempty"`
	// dormant count
	Dormant *int32 `json:"dormant,omitempty"`
}

// NewInlineResponse200100CountsByStatus instantiates a new InlineResponse200100CountsByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200100CountsByStatus() *InlineResponse200100CountsByStatus {
	this := InlineResponse200100CountsByStatus{}
	return &this
}

// NewInlineResponse200100CountsByStatusWithDefaults instantiates a new InlineResponse200100CountsByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200100CountsByStatusWithDefaults() *InlineResponse200100CountsByStatus {
	this := InlineResponse200100CountsByStatus{}
	return &this
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *InlineResponse200100CountsByStatus) GetOnline() int32 {
	if o == nil || isNil(o.Online) {
		var ret int32
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100CountsByStatus) GetOnlineOk() (*int32, bool) {
	if o == nil || isNil(o.Online) {
    return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *InlineResponse200100CountsByStatus) HasOnline() bool {
	if o != nil && !isNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given int32 and assigns it to the Online field.
func (o *InlineResponse200100CountsByStatus) SetOnline(v int32) {
	o.Online = &v
}

// GetAlerting returns the Alerting field value if set, zero value otherwise.
func (o *InlineResponse200100CountsByStatus) GetAlerting() int32 {
	if o == nil || isNil(o.Alerting) {
		var ret int32
		return ret
	}
	return *o.Alerting
}

// GetAlertingOk returns a tuple with the Alerting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100CountsByStatus) GetAlertingOk() (*int32, bool) {
	if o == nil || isNil(o.Alerting) {
    return nil, false
	}
	return o.Alerting, true
}

// HasAlerting returns a boolean if a field has been set.
func (o *InlineResponse200100CountsByStatus) HasAlerting() bool {
	if o != nil && !isNil(o.Alerting) {
		return true
	}

	return false
}

// SetAlerting gets a reference to the given int32 and assigns it to the Alerting field.
func (o *InlineResponse200100CountsByStatus) SetAlerting(v int32) {
	o.Alerting = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *InlineResponse200100CountsByStatus) GetOffline() int32 {
	if o == nil || isNil(o.Offline) {
		var ret int32
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100CountsByStatus) GetOfflineOk() (*int32, bool) {
	if o == nil || isNil(o.Offline) {
    return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *InlineResponse200100CountsByStatus) HasOffline() bool {
	if o != nil && !isNil(o.Offline) {
		return true
	}

	return false
}

// SetOffline gets a reference to the given int32 and assigns it to the Offline field.
func (o *InlineResponse200100CountsByStatus) SetOffline(v int32) {
	o.Offline = &v
}

// GetDormant returns the Dormant field value if set, zero value otherwise.
func (o *InlineResponse200100CountsByStatus) GetDormant() int32 {
	if o == nil || isNil(o.Dormant) {
		var ret int32
		return ret
	}
	return *o.Dormant
}

// GetDormantOk returns a tuple with the Dormant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100CountsByStatus) GetDormantOk() (*int32, bool) {
	if o == nil || isNil(o.Dormant) {
    return nil, false
	}
	return o.Dormant, true
}

// HasDormant returns a boolean if a field has been set.
func (o *InlineResponse200100CountsByStatus) HasDormant() bool {
	if o != nil && !isNil(o.Dormant) {
		return true
	}

	return false
}

// SetDormant gets a reference to the given int32 and assigns it to the Dormant field.
func (o *InlineResponse200100CountsByStatus) SetDormant(v int32) {
	o.Dormant = &v
}

func (o InlineResponse200100CountsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !isNil(o.Alerting) {
		toSerialize["alerting"] = o.Alerting
	}
	if !isNil(o.Offline) {
		toSerialize["offline"] = o.Offline
	}
	if !isNil(o.Dormant) {
		toSerialize["dormant"] = o.Dormant
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200100CountsByStatus struct {
	value *InlineResponse200100CountsByStatus
	isSet bool
}

func (v NullableInlineResponse200100CountsByStatus) Get() *InlineResponse200100CountsByStatus {
	return v.value
}

func (v *NullableInlineResponse200100CountsByStatus) Set(val *InlineResponse200100CountsByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200100CountsByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200100CountsByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200100CountsByStatus(val *InlineResponse200100CountsByStatus) *NullableInlineResponse200100CountsByStatus {
	return &NullableInlineResponse200100CountsByStatus{value: val, isSet: true}
}

func (v NullableInlineResponse200100CountsByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200100CountsByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

