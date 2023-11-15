/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateOrganizationSamlRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationSamlRoleRequest{}

// UpdateOrganizationSamlRoleRequest struct for UpdateOrganizationSamlRoleRequest
type UpdateOrganizationSamlRoleRequest struct {
	// The role of the SAML administrator
	Role *string `json:"role,omitempty"`
	// The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	OrgAccess *string `json:"orgAccess,omitempty"`
	// The list of tags that the SAML administrator has privileges on
	Tags []CreateOrganizationSamlRoleRequestTagsInner `json:"tags,omitempty"`
	// The list of networks that the SAML administrator has privileges on
	Networks []CreateOrganizationSamlRoleRequestNetworksInner `json:"networks,omitempty"`
}

// NewUpdateOrganizationSamlRoleRequest instantiates a new UpdateOrganizationSamlRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationSamlRoleRequest() *UpdateOrganizationSamlRoleRequest {
	this := UpdateOrganizationSamlRoleRequest{}
	return &this
}

// NewUpdateOrganizationSamlRoleRequestWithDefaults instantiates a new UpdateOrganizationSamlRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationSamlRoleRequestWithDefaults() *UpdateOrganizationSamlRoleRequest {
	this := UpdateOrganizationSamlRoleRequest{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlRoleRequest) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlRoleRequest) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlRoleRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UpdateOrganizationSamlRoleRequest) SetRole(v string) {
	o.Role = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlRoleRequest) GetOrgAccess() string {
	if o == nil || IsNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlRoleRequest) GetOrgAccessOk() (*string, bool) {
	if o == nil || IsNil(o.OrgAccess) {
		return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlRoleRequest) HasOrgAccess() bool {
	if o != nil && !IsNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *UpdateOrganizationSamlRoleRequest) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlRoleRequest) GetTags() []CreateOrganizationSamlRoleRequestTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []CreateOrganizationSamlRoleRequestTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlRoleRequest) GetTagsOk() ([]CreateOrganizationSamlRoleRequestTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlRoleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []CreateOrganizationSamlRoleRequestTagsInner and assigns it to the Tags field.
func (o *UpdateOrganizationSamlRoleRequest) SetTags(v []CreateOrganizationSamlRoleRequestTagsInner) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlRoleRequest) GetNetworks() []CreateOrganizationSamlRoleRequestNetworksInner {
	if o == nil || IsNil(o.Networks) {
		var ret []CreateOrganizationSamlRoleRequestNetworksInner
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlRoleRequest) GetNetworksOk() ([]CreateOrganizationSamlRoleRequestNetworksInner, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlRoleRequest) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []CreateOrganizationSamlRoleRequestNetworksInner and assigns it to the Networks field.
func (o *UpdateOrganizationSamlRoleRequest) SetNetworks(v []CreateOrganizationSamlRoleRequestNetworksInner) {
	o.Networks = v
}

func (o UpdateOrganizationSamlRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationSamlRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
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

type NullableUpdateOrganizationSamlRoleRequest struct {
	value *UpdateOrganizationSamlRoleRequest
	isSet bool
}

func (v NullableUpdateOrganizationSamlRoleRequest) Get() *UpdateOrganizationSamlRoleRequest {
	return v.value
}

func (v *NullableUpdateOrganizationSamlRoleRequest) Set(val *UpdateOrganizationSamlRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationSamlRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationSamlRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationSamlRoleRequest(val *UpdateOrganizationSamlRoleRequest) *NullableUpdateOrganizationSamlRoleRequest {
	return &NullableUpdateOrganizationSamlRoleRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationSamlRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationSamlRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


