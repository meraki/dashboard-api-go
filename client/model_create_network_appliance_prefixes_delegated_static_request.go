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

// checks if the CreateNetworkAppliancePrefixesDelegatedStaticRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkAppliancePrefixesDelegatedStaticRequest{}

// CreateNetworkAppliancePrefixesDelegatedStaticRequest struct for CreateNetworkAppliancePrefixesDelegatedStaticRequest
type CreateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	// A static IPv6 prefix
	Prefix string `json:"prefix"`
	Origin CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin `json:"origin"`
	// A name or description for the prefix
	Description *string `json:"description,omitempty"`
}

// NewCreateNetworkAppliancePrefixesDelegatedStaticRequest instantiates a new CreateNetworkAppliancePrefixesDelegatedStaticRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkAppliancePrefixesDelegatedStaticRequest(prefix string, origin CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin) *CreateNetworkAppliancePrefixesDelegatedStaticRequest {
	this := CreateNetworkAppliancePrefixesDelegatedStaticRequest{}
	this.Prefix = prefix
	this.Origin = origin
	return &this
}

// NewCreateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults instantiates a new CreateNetworkAppliancePrefixesDelegatedStaticRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults() *CreateNetworkAppliancePrefixesDelegatedStaticRequest {
	this := CreateNetworkAppliancePrefixesDelegatedStaticRequest{}
	return &this
}

// GetPrefix returns the Prefix field value
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) SetPrefix(v string) {
	o.Prefix = v
}

// GetOrigin returns the Origin field value
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetOrigin() CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin {
	if o == nil {
		var ret CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetOriginOk() (*CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) SetOrigin(v CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin) {
	o.Origin = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CreateNetworkAppliancePrefixesDelegatedStaticRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkAppliancePrefixesDelegatedStaticRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prefix"] = o.Prefix
	toSerialize["origin"] = o.Origin
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	value *CreateNetworkAppliancePrefixesDelegatedStaticRequest
	isSet bool
}

func (v NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest) Get() *CreateNetworkAppliancePrefixesDelegatedStaticRequest {
	return v.value
}

func (v *NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest) Set(val *CreateNetworkAppliancePrefixesDelegatedStaticRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkAppliancePrefixesDelegatedStaticRequest(val *CreateNetworkAppliancePrefixesDelegatedStaticRequest) *NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest {
	return &NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkAppliancePrefixesDelegatedStaticRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


