/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GenerateDeviceCameraSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateDeviceCameraSnapshotRequest{}

// GenerateDeviceCameraSnapshotRequest struct for GenerateDeviceCameraSnapshotRequest
type GenerateDeviceCameraSnapshotRequest struct {
	// [optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// [optional] If set to \"true\" the snapshot will be taken at full sensor resolution. This will error if used with timestamp.
	Fullframe *bool `json:"fullframe,omitempty"`
}

// NewGenerateDeviceCameraSnapshotRequest instantiates a new GenerateDeviceCameraSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDeviceCameraSnapshotRequest() *GenerateDeviceCameraSnapshotRequest {
	this := GenerateDeviceCameraSnapshotRequest{}
	return &this
}

// NewGenerateDeviceCameraSnapshotRequestWithDefaults instantiates a new GenerateDeviceCameraSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDeviceCameraSnapshotRequestWithDefaults() *GenerateDeviceCameraSnapshotRequest {
	this := GenerateDeviceCameraSnapshotRequest{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GenerateDeviceCameraSnapshotRequest) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateDeviceCameraSnapshotRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GenerateDeviceCameraSnapshotRequest) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *GenerateDeviceCameraSnapshotRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetFullframe returns the Fullframe field value if set, zero value otherwise.
func (o *GenerateDeviceCameraSnapshotRequest) GetFullframe() bool {
	if o == nil || IsNil(o.Fullframe) {
		var ret bool
		return ret
	}
	return *o.Fullframe
}

// GetFullframeOk returns a tuple with the Fullframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateDeviceCameraSnapshotRequest) GetFullframeOk() (*bool, bool) {
	if o == nil || IsNil(o.Fullframe) {
		return nil, false
	}
	return o.Fullframe, true
}

// HasFullframe returns a boolean if a field has been set.
func (o *GenerateDeviceCameraSnapshotRequest) HasFullframe() bool {
	if o != nil && !IsNil(o.Fullframe) {
		return true
	}

	return false
}

// SetFullframe gets a reference to the given bool and assigns it to the Fullframe field.
func (o *GenerateDeviceCameraSnapshotRequest) SetFullframe(v bool) {
	o.Fullframe = &v
}

func (o GenerateDeviceCameraSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateDeviceCameraSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Fullframe) {
		toSerialize["fullframe"] = o.Fullframe
	}
	return toSerialize, nil
}

type NullableGenerateDeviceCameraSnapshotRequest struct {
	value *GenerateDeviceCameraSnapshotRequest
	isSet bool
}

func (v NullableGenerateDeviceCameraSnapshotRequest) Get() *GenerateDeviceCameraSnapshotRequest {
	return v.value
}

func (v *NullableGenerateDeviceCameraSnapshotRequest) Set(val *GenerateDeviceCameraSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDeviceCameraSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDeviceCameraSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDeviceCameraSnapshotRequest(val *GenerateDeviceCameraSnapshotRequest) *NullableGenerateDeviceCameraSnapshotRequest {
	return &NullableGenerateDeviceCameraSnapshotRequest{value: val, isSet: true}
}

func (v NullableGenerateDeviceCameraSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateDeviceCameraSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


