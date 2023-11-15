/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{}

// CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct for CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest
type CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	// Name of the custom performance class
	Name string `json:"name"`
	// Maximum latency in milliseconds
	MaxLatency *int32 `json:"maxLatency,omitempty"`
	// Maximum jitter in milliseconds
	MaxJitter *int32 `json:"maxJitter,omitempty"`
	// Maximum percentage of packet loss
	MaxLossPercentage *int32 `json:"maxLossPercentage,omitempty"`
}

// NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest instantiates a new CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(name string) *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	this := CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{}
	this.Name = name
	return &this
}

// NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequestWithDefaults instantiates a new CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequestWithDefaults() *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	this := CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetName(v string) {
	o.Name = v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLatency() int32 {
	if o == nil || IsNil(o.MaxLatency) {
		var ret int32
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLatency) {
		return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxLatency() bool {
	if o != nil && !IsNil(o.MaxLatency) {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given int32 and assigns it to the MaxLatency field.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxLatency(v int32) {
	o.MaxLatency = &v
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxJitter() int32 {
	if o == nil || IsNil(o.MaxJitter) {
		var ret int32
		return ret
	}
	return *o.MaxJitter
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxJitterOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxJitter) {
		return nil, false
	}
	return o.MaxJitter, true
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxJitter() bool {
	if o != nil && !IsNil(o.MaxJitter) {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given int32 and assigns it to the MaxJitter field.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxJitter(v int32) {
	o.MaxJitter = &v
}

// GetMaxLossPercentage returns the MaxLossPercentage field value if set, zero value otherwise.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLossPercentage() int32 {
	if o == nil || IsNil(o.MaxLossPercentage) {
		var ret int32
		return ret
	}
	return *o.MaxLossPercentage
}

// GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLossPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLossPercentage) {
		return nil, false
	}
	return o.MaxLossPercentage, true
}

// HasMaxLossPercentage returns a boolean if a field has been set.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxLossPercentage() bool {
	if o != nil && !IsNil(o.MaxLossPercentage) {
		return true
	}

	return false
}

// SetMaxLossPercentage gets a reference to the given int32 and assigns it to the MaxLossPercentage field.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxLossPercentage(v int32) {
	o.MaxLossPercentage = &v
}

func (o CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.MaxLatency) {
		toSerialize["maxLatency"] = o.MaxLatency
	}
	if !IsNil(o.MaxJitter) {
		toSerialize["maxJitter"] = o.MaxJitter
	}
	if !IsNil(o.MaxLossPercentage) {
		toSerialize["maxLossPercentage"] = o.MaxLossPercentage
	}
	return toSerialize, nil
}

type NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	value *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest
	isSet bool
}

func (v NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Get() *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return v.value
}

func (v *NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Set(val *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(val *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) *NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return &NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


