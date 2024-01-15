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

// checks if the UpdateOrganizationAdminRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationAdminRequest{}

// UpdateOrganizationAdminRequest struct for UpdateOrganizationAdminRequest
type UpdateOrganizationAdminRequest struct {
	// The name of the dashboard administrator
	Name *string `json:"name,omitempty"`
	// The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	OrgAccess *string `json:"orgAccess,omitempty"`
	// The list of tags that the dashboard administrator has privileges on
	Tags []CreateOrganizationAdminRequestTagsInner `json:"tags,omitempty"`
	// The list of networks that the dashboard administrator has privileges on
	Networks []CreateOrganizationAdminRequestNetworksInner `json:"networks,omitempty"`
}

// NewUpdateOrganizationAdminRequest instantiates a new UpdateOrganizationAdminRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationAdminRequest() *UpdateOrganizationAdminRequest {
	this := UpdateOrganizationAdminRequest{}
	return &this
}

// NewUpdateOrganizationAdminRequestWithDefaults instantiates a new UpdateOrganizationAdminRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationAdminRequestWithDefaults() *UpdateOrganizationAdminRequest {
	this := UpdateOrganizationAdminRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationAdminRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdminRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationAdminRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationAdminRequest) SetName(v string) {
	o.Name = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *UpdateOrganizationAdminRequest) GetOrgAccess() string {
	if o == nil || IsNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdminRequest) GetOrgAccessOk() (*string, bool) {
	if o == nil || IsNil(o.OrgAccess) {
		return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *UpdateOrganizationAdminRequest) HasOrgAccess() bool {
	if o != nil && !IsNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *UpdateOrganizationAdminRequest) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateOrganizationAdminRequest) GetTags() []CreateOrganizationAdminRequestTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []CreateOrganizationAdminRequestTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdminRequest) GetTagsOk() ([]CreateOrganizationAdminRequestTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateOrganizationAdminRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []CreateOrganizationAdminRequestTagsInner and assigns it to the Tags field.
func (o *UpdateOrganizationAdminRequest) SetTags(v []CreateOrganizationAdminRequestTagsInner) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *UpdateOrganizationAdminRequest) GetNetworks() []CreateOrganizationAdminRequestNetworksInner {
	if o == nil || IsNil(o.Networks) {
		var ret []CreateOrganizationAdminRequestNetworksInner
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdminRequest) GetNetworksOk() ([]CreateOrganizationAdminRequestNetworksInner, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *UpdateOrganizationAdminRequest) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []CreateOrganizationAdminRequestNetworksInner and assigns it to the Networks field.
func (o *UpdateOrganizationAdminRequest) SetNetworks(v []CreateOrganizationAdminRequestNetworksInner) {
	o.Networks = v
}

func (o UpdateOrganizationAdminRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationAdminRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrgAccess) {
		toSerialize["orgAccess"] = o.OrgAccess
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationAdminRequest struct {
	value *UpdateOrganizationAdminRequest
	isSet bool
}

func (v NullableUpdateOrganizationAdminRequest) Get() *UpdateOrganizationAdminRequest {
	return v.value
}

func (v *NullableUpdateOrganizationAdminRequest) Set(val *UpdateOrganizationAdminRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationAdminRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationAdminRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationAdminRequest(val *UpdateOrganizationAdminRequest) *NullableUpdateOrganizationAdminRequest {
	return &NullableUpdateOrganizationAdminRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationAdminRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationAdminRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


