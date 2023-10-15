/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions{}

// GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions Descriptions of the early access feature
type GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions struct {
	// Short description
	Short *string `json:"short,omitempty"`
	// Long description
	Long *string `json:"long,omitempty"`
}

// NewGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions instantiates a new GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions() *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions {
	this := GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions{}
	return &this
}

// NewGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptionsWithDefaults instantiates a new GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptionsWithDefaults() *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions {
	this := GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions{}
	return &this
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) GetShort() string {
	if o == nil || IsNil(o.Short) {
		var ret string
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) GetShortOk() (*string, bool) {
	if o == nil || IsNil(o.Short) {
		return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) HasShort() bool {
	if o != nil && !IsNil(o.Short) {
		return true
	}

	return false
}

// SetShort gets a reference to the given string and assigns it to the Short field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) SetShort(v string) {
	o.Short = &v
}

// GetLong returns the Long field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) GetLong() string {
	if o == nil || IsNil(o.Long) {
		var ret string
		return ret
	}
	return *o.Long
}

// GetLongOk returns a tuple with the Long field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) GetLongOk() (*string, bool) {
	if o == nil || IsNil(o.Long) {
		return nil, false
	}
	return o.Long, true
}

// HasLong returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) HasLong() bool {
	if o != nil && !IsNil(o.Long) {
		return true
	}

	return false
}

// SetLong gets a reference to the given string and assigns it to the Long field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) SetLong(v string) {
	o.Long = &v
}

func (o GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Short) {
		toSerialize["short"] = o.Short
	}
	if !IsNil(o.Long) {
		toSerialize["long"] = o.Long
	}
	return toSerialize, nil
}

type NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions struct {
	value *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions
	isSet bool
}

func (v NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) Get() *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions {
	return v.value
}

func (v *NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) Set(val *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions(val *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) *NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions {
	return &NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions{value: val, isSet: true}
}

func (v NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


