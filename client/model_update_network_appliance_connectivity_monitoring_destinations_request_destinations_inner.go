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

// UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner struct for UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner
type UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner struct {
	// The IP address to test connectivity with
	Ip string `json:"ip"`
	// Description of the testing destination. Optional, defaults to null
	Description *string `json:"description,omitempty"`
	// Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Default *bool `json:"default,omitempty"`
}

// NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner instantiates a new UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner(ip string) *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner {
	this := UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner{}
	this.Ip = ip
	return &this
}

// NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInnerWithDefaults instantiates a new UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInnerWithDefaults() *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner {
	this := UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner{}
	return &this
}

// GetIp returns the Ip field value
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) GetIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) SetIp(v string) {
	o.Ip = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) SetDescription(v string) {
	o.Description = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
    return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) SetDefault(v bool) {
	o.Default = &v
}

func (o UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner struct {
	value *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) Get() *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) Set(val *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner(val *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner {
	return &NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


