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

// checks if the UpdateNetworkApplianceTrafficShapingVpnExclusions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingVpnExclusions200Response{}

// UpdateNetworkApplianceTrafficShapingVpnExclusions200Response struct for UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
type UpdateNetworkApplianceTrafficShapingVpnExclusions200Response struct {
	// ID of the network whose VPN exclusion rules are returned.
	NetworkId string `json:"networkId"`
	// Name of the network whose VPN exclusion rules are returned.
	NetworkName string `json:"networkName"`
	// Custom VPN exclusion rules.
	Custom []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner `json:"custom"`
	// Major Application based VPN exclusion rules.
	MajorApplications []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner `json:"majorApplications"`
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusions200Response instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200Response(networkId string, networkName string, custom []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner, majorApplications []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusions200Response{}
	this.NetworkId = networkId
	this.NetworkName = networkName
	this.Custom = custom
	this.MajorApplications = majorApplications
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusions200Response{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetNetworkName returns the NetworkName field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetNetworkName(v string) {
	o.NetworkName = v
}

// GetCustom returns the Custom field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetCustom() []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner {
	if o == nil {
		var ret []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetCustomOk() ([]UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Custom, true
}

// SetCustom sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetCustom(v []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner) {
	o.Custom = v
}

// GetMajorApplications returns the MajorApplications field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetMajorApplications() []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner {
	if o == nil {
		var ret []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner
		return ret
	}

	return o.MajorApplications
}

// GetMajorApplicationsOk returns a tuple with the MajorApplications field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetMajorApplicationsOk() ([]UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.MajorApplications, true
}

// SetMajorApplications sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetMajorApplications(v []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) {
	o.MajorApplications = v
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkId"] = o.NetworkId
	toSerialize["networkName"] = o.NetworkName
	toSerialize["custom"] = o.Custom
	toSerialize["majorApplications"] = o.MajorApplications
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response struct {
	value *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response) Get() *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response) Set(val *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response(val *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response {
	return &NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


