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

// checks if the GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork{}

// GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork Information on network access to the template
type GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork struct {
	// Indicates whether network admins may modify this template
	AdminsCanModify *bool `json:"adminsCanModify,omitempty"`
}

// NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork() *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork {
	this := GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork{}
	return &this
}

// NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetworkWithDefaults instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetworkWithDefaults() *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork {
	this := GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork{}
	return &this
}

// GetAdminsCanModify returns the AdminsCanModify field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) GetAdminsCanModify() bool {
	if o == nil || IsNil(o.AdminsCanModify) {
		var ret bool
		return ret
	}
	return *o.AdminsCanModify
}

// GetAdminsCanModifyOk returns a tuple with the AdminsCanModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) GetAdminsCanModifyOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminsCanModify) {
		return nil, false
	}
	return o.AdminsCanModify, true
}

// HasAdminsCanModify returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) HasAdminsCanModify() bool {
	if o != nil && !IsNil(o.AdminsCanModify) {
		return true
	}

	return false
}

// SetAdminsCanModify gets a reference to the given bool and assigns it to the AdminsCanModify field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) SetAdminsCanModify(v bool) {
	o.AdminsCanModify = &v
}

func (o GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminsCanModify) {
		toSerialize["adminsCanModify"] = o.AdminsCanModify
	}
	return toSerialize, nil
}

type NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork struct {
	value *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork
	isSet bool
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) Get() *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork {
	return v.value
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) Set(val *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork(val *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork {
	return &NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork{value: val, isSet: true}
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


