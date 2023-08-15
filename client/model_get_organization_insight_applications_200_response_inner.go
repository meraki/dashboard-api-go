/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationInsightApplications200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationInsightApplications200ResponseInner{}

// GetOrganizationInsightApplications200ResponseInner struct for GetOrganizationInsightApplications200ResponseInner
type GetOrganizationInsightApplications200ResponseInner struct {
	// Application identifier
	ApplicationId *string `json:"applicationId,omitempty"`
	// Application name
	Name *string `json:"name,omitempty"`
	Thresholds *GetOrganizationInsightApplications200ResponseInnerThresholds `json:"thresholds,omitempty"`
}

// NewGetOrganizationInsightApplications200ResponseInner instantiates a new GetOrganizationInsightApplications200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationInsightApplications200ResponseInner() *GetOrganizationInsightApplications200ResponseInner {
	this := GetOrganizationInsightApplications200ResponseInner{}
	return &this
}

// NewGetOrganizationInsightApplications200ResponseInnerWithDefaults instantiates a new GetOrganizationInsightApplications200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationInsightApplications200ResponseInnerWithDefaults() *GetOrganizationInsightApplications200ResponseInner {
	this := GetOrganizationInsightApplications200ResponseInner{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInner) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInner) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInner) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *GetOrganizationInsightApplications200ResponseInner) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationInsightApplications200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInner) GetThresholds() GetOrganizationInsightApplications200ResponseInnerThresholds {
	if o == nil || IsNil(o.Thresholds) {
		var ret GetOrganizationInsightApplications200ResponseInnerThresholds
		return ret
	}
	return *o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInner) GetThresholdsOk() (*GetOrganizationInsightApplications200ResponseInnerThresholds, bool) {
	if o == nil || IsNil(o.Thresholds) {
		return nil, false
	}
	return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInner) HasThresholds() bool {
	if o != nil && !IsNil(o.Thresholds) {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given GetOrganizationInsightApplications200ResponseInnerThresholds and assigns it to the Thresholds field.
func (o *GetOrganizationInsightApplications200ResponseInner) SetThresholds(v GetOrganizationInsightApplications200ResponseInnerThresholds) {
	o.Thresholds = &v
}

func (o GetOrganizationInsightApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationInsightApplications200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Thresholds) {
		toSerialize["thresholds"] = o.Thresholds
	}
	return toSerialize, nil
}

type NullableGetOrganizationInsightApplications200ResponseInner struct {
	value *GetOrganizationInsightApplications200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationInsightApplications200ResponseInner) Get() *GetOrganizationInsightApplications200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationInsightApplications200ResponseInner) Set(val *GetOrganizationInsightApplications200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationInsightApplications200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationInsightApplications200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationInsightApplications200ResponseInner(val *GetOrganizationInsightApplications200ResponseInner) *NullableGetOrganizationInsightApplications200ResponseInner {
	return &NullableGetOrganizationInsightApplications200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationInsightApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationInsightApplications200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


