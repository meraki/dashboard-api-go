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

// UpdateNetworkSwitchRoutingOspfRequestAreasInner struct for UpdateNetworkSwitchRoutingOspfRequestAreasInner
type UpdateNetworkSwitchRoutingOspfRequestAreasInner struct {
	// OSPF area ID
	AreaId string `json:"areaId"`
	// Name of the OSPF area
	AreaName string `json:"areaName"`
	// Area types in OSPF. Must be one of: [\"normal\", \"stub\", \"nssa\"]
	AreaType string `json:"areaType"`
}

// NewUpdateNetworkSwitchRoutingOspfRequestAreasInner instantiates a new UpdateNetworkSwitchRoutingOspfRequestAreasInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingOspfRequestAreasInner(areaId string, areaName string, areaType string) *UpdateNetworkSwitchRoutingOspfRequestAreasInner {
	this := UpdateNetworkSwitchRoutingOspfRequestAreasInner{}
	this.AreaId = areaId
	this.AreaName = areaName
	this.AreaType = areaType
	return &this
}

// NewUpdateNetworkSwitchRoutingOspfRequestAreasInnerWithDefaults instantiates a new UpdateNetworkSwitchRoutingOspfRequestAreasInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingOspfRequestAreasInnerWithDefaults() *UpdateNetworkSwitchRoutingOspfRequestAreasInner {
	this := UpdateNetworkSwitchRoutingOspfRequestAreasInner{}
	return &this
}

// GetAreaId returns the AreaId field value
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AreaId
}

// GetAreaIdOk returns a tuple with the AreaId field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AreaId, true
}

// SetAreaId sets field value
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) SetAreaId(v string) {
	o.AreaId = v
}

// GetAreaName returns the AreaName field value
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AreaName
}

// GetAreaNameOk returns a tuple with the AreaName field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AreaName, true
}

// SetAreaName sets field value
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) SetAreaName(v string) {
	o.AreaName = v
}

// GetAreaType returns the AreaType field value
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AreaType
}

// GetAreaTypeOk returns a tuple with the AreaType field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AreaType, true
}

// SetAreaType sets field value
func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) SetAreaType(v string) {
	o.AreaType = v
}

func (o UpdateNetworkSwitchRoutingOspfRequestAreasInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["areaId"] = o.AreaId
	}
	if true {
		toSerialize["areaName"] = o.AreaName
	}
	if true {
		toSerialize["areaType"] = o.AreaType
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner struct {
	value *UpdateNetworkSwitchRoutingOspfRequestAreasInner
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner) Get() *UpdateNetworkSwitchRoutingOspfRequestAreasInner {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner) Set(val *UpdateNetworkSwitchRoutingOspfRequestAreasInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingOspfRequestAreasInner(val *UpdateNetworkSwitchRoutingOspfRequestAreasInner) *NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner {
	return &NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestAreasInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


