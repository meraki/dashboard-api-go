/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner struct for GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner
type GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner struct {
	// config id
	TrustedAccessConfigId *string `json:"trustedAccessConfigId,omitempty"`
	// time that config was downloaded
	DownloadedAt *string `json:"downloadedAt,omitempty"`
	// time that SCEP completed
	ScepCompletedAt *time.Time `json:"scepCompletedAt,omitempty"`
	// time of last connection
	LastConnectedAt *time.Time `json:"lastConnectedAt,omitempty"`
}

// NewGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner instantiates a new GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner() *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner {
	this := GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner{}
	return &this
}

// NewGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInnerWithDefaults instantiates a new GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInnerWithDefaults() *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner {
	this := GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner{}
	return &this
}

// GetTrustedAccessConfigId returns the TrustedAccessConfigId field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetTrustedAccessConfigId() string {
	if o == nil || isNil(o.TrustedAccessConfigId) {
		var ret string
		return ret
	}
	return *o.TrustedAccessConfigId
}

// GetTrustedAccessConfigIdOk returns a tuple with the TrustedAccessConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetTrustedAccessConfigIdOk() (*string, bool) {
	if o == nil || isNil(o.TrustedAccessConfigId) {
    return nil, false
	}
	return o.TrustedAccessConfigId, true
}

// HasTrustedAccessConfigId returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) HasTrustedAccessConfigId() bool {
	if o != nil && !isNil(o.TrustedAccessConfigId) {
		return true
	}

	return false
}

// SetTrustedAccessConfigId gets a reference to the given string and assigns it to the TrustedAccessConfigId field.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) SetTrustedAccessConfigId(v string) {
	o.TrustedAccessConfigId = &v
}

// GetDownloadedAt returns the DownloadedAt field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetDownloadedAt() string {
	if o == nil || isNil(o.DownloadedAt) {
		var ret string
		return ret
	}
	return *o.DownloadedAt
}

// GetDownloadedAtOk returns a tuple with the DownloadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetDownloadedAtOk() (*string, bool) {
	if o == nil || isNil(o.DownloadedAt) {
    return nil, false
	}
	return o.DownloadedAt, true
}

// HasDownloadedAt returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) HasDownloadedAt() bool {
	if o != nil && !isNil(o.DownloadedAt) {
		return true
	}

	return false
}

// SetDownloadedAt gets a reference to the given string and assigns it to the DownloadedAt field.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) SetDownloadedAt(v string) {
	o.DownloadedAt = &v
}

// GetScepCompletedAt returns the ScepCompletedAt field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetScepCompletedAt() time.Time {
	if o == nil || isNil(o.ScepCompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.ScepCompletedAt
}

// GetScepCompletedAtOk returns a tuple with the ScepCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetScepCompletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScepCompletedAt) {
    return nil, false
	}
	return o.ScepCompletedAt, true
}

// HasScepCompletedAt returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) HasScepCompletedAt() bool {
	if o != nil && !isNil(o.ScepCompletedAt) {
		return true
	}

	return false
}

// SetScepCompletedAt gets a reference to the given time.Time and assigns it to the ScepCompletedAt field.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) SetScepCompletedAt(v time.Time) {
	o.ScepCompletedAt = &v
}

// GetLastConnectedAt returns the LastConnectedAt field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetLastConnectedAt() time.Time {
	if o == nil || isNil(o.LastConnectedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastConnectedAt
}

// GetLastConnectedAtOk returns a tuple with the LastConnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) GetLastConnectedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastConnectedAt) {
    return nil, false
	}
	return o.LastConnectedAt, true
}

// HasLastConnectedAt returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) HasLastConnectedAt() bool {
	if o != nil && !isNil(o.LastConnectedAt) {
		return true
	}

	return false
}

// SetLastConnectedAt gets a reference to the given time.Time and assigns it to the LastConnectedAt field.
func (o *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) SetLastConnectedAt(v time.Time) {
	o.LastConnectedAt = &v
}

func (o GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TrustedAccessConfigId) {
		toSerialize["trustedAccessConfigId"] = o.TrustedAccessConfigId
	}
	if !isNil(o.DownloadedAt) {
		toSerialize["downloadedAt"] = o.DownloadedAt
	}
	if !isNil(o.ScepCompletedAt) {
		toSerialize["scepCompletedAt"] = o.ScepCompletedAt
	}
	if !isNil(o.LastConnectedAt) {
		toSerialize["lastConnectedAt"] = o.LastConnectedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner struct {
	value *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner
	isSet bool
}

func (v NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) Get() *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner {
	return v.value
}

func (v *NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) Set(val *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner(val *GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) *NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner {
	return &NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


