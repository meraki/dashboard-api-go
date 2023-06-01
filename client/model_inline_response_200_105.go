/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200105 struct for InlineResponse200105
type InlineResponse200105 struct {
	// Application identifier
	ApplicationId *string `json:"applicationId,omitempty"`
	// Application name
	Name *string `json:"name,omitempty"`
	Thresholds *OrganizationsOrganizationIdInsightApplicationsThresholds `json:"thresholds,omitempty"`
}

// NewInlineResponse200105 instantiates a new InlineResponse200105 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200105() *InlineResponse200105 {
	this := InlineResponse200105{}
	return &this
}

// NewInlineResponse200105WithDefaults instantiates a new InlineResponse200105 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200105WithDefaults() *InlineResponse200105 {
	this := InlineResponse200105{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *InlineResponse200105) GetApplicationId() string {
	if o == nil || isNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetApplicationIdOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationId) {
    return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *InlineResponse200105) HasApplicationId() bool {
	if o != nil && !isNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *InlineResponse200105) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200105) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200105) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200105) SetName(v string) {
	o.Name = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *InlineResponse200105) GetThresholds() OrganizationsOrganizationIdInsightApplicationsThresholds {
	if o == nil || isNil(o.Thresholds) {
		var ret OrganizationsOrganizationIdInsightApplicationsThresholds
		return ret
	}
	return *o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetThresholdsOk() (*OrganizationsOrganizationIdInsightApplicationsThresholds, bool) {
	if o == nil || isNil(o.Thresholds) {
    return nil, false
	}
	return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *InlineResponse200105) HasThresholds() bool {
	if o != nil && !isNil(o.Thresholds) {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given OrganizationsOrganizationIdInsightApplicationsThresholds and assigns it to the Thresholds field.
func (o *InlineResponse200105) SetThresholds(v OrganizationsOrganizationIdInsightApplicationsThresholds) {
	o.Thresholds = &v
}

func (o InlineResponse200105) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Thresholds) {
		toSerialize["thresholds"] = o.Thresholds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200105 struct {
	value *InlineResponse200105
	isSet bool
}

func (v NullableInlineResponse200105) Get() *InlineResponse200105 {
	return v.value
}

func (v *NullableInlineResponse200105) Set(val *InlineResponse200105) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200105) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200105) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200105(val *InlineResponse200105) *NullableInlineResponse200105 {
	return &NullableInlineResponse200105{value: val, isSet: true}
}

func (v NullableInlineResponse200105) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200105) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

