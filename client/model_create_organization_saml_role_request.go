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

// checks if the CreateOrganizationSamlRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationSamlRoleRequest{}

// CreateOrganizationSamlRoleRequest struct for CreateOrganizationSamlRoleRequest
type CreateOrganizationSamlRoleRequest struct {
	// The role of the SAML administrator
	Role string `json:"role"`
	// The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	OrgAccess string `json:"orgAccess"`
	// The list of tags that the SAML administrator has privileges on
	Tags []CreateOrganizationSamlRoleRequestTagsInner `json:"tags,omitempty"`
	// The list of networks that the SAML administrator has privileges on
	Networks []CreateOrganizationSamlRoleRequestNetworksInner `json:"networks,omitempty"`
}

// NewCreateOrganizationSamlRoleRequest instantiates a new CreateOrganizationSamlRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationSamlRoleRequest(role string, orgAccess string) *CreateOrganizationSamlRoleRequest {
	this := CreateOrganizationSamlRoleRequest{}
	this.Role = role
	this.OrgAccess = orgAccess
	return &this
}

// NewCreateOrganizationSamlRoleRequestWithDefaults instantiates a new CreateOrganizationSamlRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationSamlRoleRequestWithDefaults() *CreateOrganizationSamlRoleRequest {
	this := CreateOrganizationSamlRoleRequest{}
	return &this
}

// GetRole returns the Role field value
func (o *CreateOrganizationSamlRoleRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CreateOrganizationSamlRoleRequest) SetRole(v string) {
	o.Role = v
}

// GetOrgAccess returns the OrgAccess field value
func (o *CreateOrganizationSamlRoleRequest) GetOrgAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequest) GetOrgAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgAccess, true
}

// SetOrgAccess sets field value
func (o *CreateOrganizationSamlRoleRequest) SetOrgAccess(v string) {
	o.OrgAccess = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateOrganizationSamlRoleRequest) GetTags() []CreateOrganizationSamlRoleRequestTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []CreateOrganizationSamlRoleRequestTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequest) GetTagsOk() ([]CreateOrganizationSamlRoleRequestTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateOrganizationSamlRoleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []CreateOrganizationSamlRoleRequestTagsInner and assigns it to the Tags field.
func (o *CreateOrganizationSamlRoleRequest) SetTags(v []CreateOrganizationSamlRoleRequestTagsInner) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *CreateOrganizationSamlRoleRequest) GetNetworks() []CreateOrganizationSamlRoleRequestNetworksInner {
	if o == nil || IsNil(o.Networks) {
		var ret []CreateOrganizationSamlRoleRequestNetworksInner
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequest) GetNetworksOk() ([]CreateOrganizationSamlRoleRequestNetworksInner, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *CreateOrganizationSamlRoleRequest) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []CreateOrganizationSamlRoleRequestNetworksInner and assigns it to the Networks field.
func (o *CreateOrganizationSamlRoleRequest) SetNetworks(v []CreateOrganizationSamlRoleRequestNetworksInner) {
	o.Networks = v
}

func (o CreateOrganizationSamlRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationSamlRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["orgAccess"] = o.OrgAccess
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	return toSerialize, nil
}

type NullableCreateOrganizationSamlRoleRequest struct {
	value *CreateOrganizationSamlRoleRequest
	isSet bool
}

func (v NullableCreateOrganizationSamlRoleRequest) Get() *CreateOrganizationSamlRoleRequest {
	return v.value
}

func (v *NullableCreateOrganizationSamlRoleRequest) Set(val *CreateOrganizationSamlRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationSamlRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationSamlRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationSamlRoleRequest(val *CreateOrganizationSamlRoleRequest) *NullableCreateOrganizationSamlRoleRequest {
	return &NullableCreateOrganizationSamlRoleRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationSamlRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationSamlRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


