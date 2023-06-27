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

// UpdateNetworkFirmwareUpgradesStagedStagesRequest struct for UpdateNetworkFirmwareUpgradesStagedStagesRequest
type UpdateNetworkFirmwareUpgradesStagedStagesRequest struct {
	// Array of Staged Upgrade Groups
	Json []UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner `json:"_json,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesStagedStagesRequest instantiates a new UpdateNetworkFirmwareUpgradesStagedStagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedStagesRequest() *UpdateNetworkFirmwareUpgradesStagedStagesRequest {
	this := UpdateNetworkFirmwareUpgradesStagedStagesRequest{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedStagesRequestWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedStagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedStagesRequestWithDefaults() *UpdateNetworkFirmwareUpgradesStagedStagesRequest {
	this := UpdateNetworkFirmwareUpgradesStagedStagesRequest{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequest) GetJson() []UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner {
	if o == nil || isNil(o.Json) {
		var ret []UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner
		return ret
	}
	return o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequest) GetJsonOk() ([]UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner, bool) {
	if o == nil || isNil(o.Json) {
    return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequest) HasJson() bool {
	if o != nil && !isNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given []UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner and assigns it to the Json field.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequest) SetJson(v []UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) {
	o.Json = v
}

func (o UpdateNetworkFirmwareUpgradesStagedStagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Json) {
		toSerialize["_json"] = o.Json
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest struct {
	value *UpdateNetworkFirmwareUpgradesStagedStagesRequest
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest) Get() *UpdateNetworkFirmwareUpgradesStagedStagesRequest {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest) Set(val *UpdateNetworkFirmwareUpgradesStagedStagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedStagesRequest(val *UpdateNetworkFirmwareUpgradesStagedStagesRequest) *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest {
	return &NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

