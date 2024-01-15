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

// checks if the UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner{}

// UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner struct for UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner
type UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner struct {
	// The name of the custom pie chart item.
	Name string `json:"name"`
	//     The signature type for the custom pie chart item. Can be one of 'host', 'port' or 'ipRange'. 
	Type string `json:"type"`
	//     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item     (see sample request/response for more details). 
	Value string `json:"value"`
}

// NewUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner instantiates a new UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner(name string, type_ string, value string) *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner {
	this := UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner{}
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInnerWithDefaults instantiates a new UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInnerWithDefaults() *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner {
	this := UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) SetValue(v string) {
	o.Value = v
}

func (o UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner struct {
	value *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner
	isSet bool
}

func (v NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) Get() *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner {
	return v.value
}

func (v *NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) Set(val *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner(val *UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) *NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner {
	return &NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


