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

// checks if the CreateOrganizationCameraRoleRequestAppliedOnDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationCameraRoleRequestAppliedOnDevicesInner{}

// CreateOrganizationCameraRoleRequestAppliedOnDevicesInner struct for CreateOrganizationCameraRoleRequestAppliedOnDevicesInner
type CreateOrganizationCameraRoleRequestAppliedOnDevicesInner struct {
	// Device tag.
	Tag *string `json:"tag,omitempty"`
	// Device id.
	Id *string `json:"id,omitempty"`
	// Network tag scope
	InNetworksWithTag *string `json:"inNetworksWithTag,omitempty"`
	// Network id scope
	InNetworksWithId *string `json:"inNetworksWithId,omitempty"`
	// Permission scope id
	PermissionScopeId string `json:"permissionScopeId"`
}

// NewCreateOrganizationCameraRoleRequestAppliedOnDevicesInner instantiates a new CreateOrganizationCameraRoleRequestAppliedOnDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationCameraRoleRequestAppliedOnDevicesInner(permissionScopeId string) *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner {
	this := CreateOrganizationCameraRoleRequestAppliedOnDevicesInner{}
	this.PermissionScopeId = permissionScopeId
	return &this
}

// NewCreateOrganizationCameraRoleRequestAppliedOnDevicesInnerWithDefaults instantiates a new CreateOrganizationCameraRoleRequestAppliedOnDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationCameraRoleRequestAppliedOnDevicesInnerWithDefaults() *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner {
	this := CreateOrganizationCameraRoleRequestAppliedOnDevicesInner{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) SetTag(v string) {
	o.Tag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) SetId(v string) {
	o.Id = &v
}

// GetInNetworksWithTag returns the InNetworksWithTag field value if set, zero value otherwise.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetInNetworksWithTag() string {
	if o == nil || IsNil(o.InNetworksWithTag) {
		var ret string
		return ret
	}
	return *o.InNetworksWithTag
}

// GetInNetworksWithTagOk returns a tuple with the InNetworksWithTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetInNetworksWithTagOk() (*string, bool) {
	if o == nil || IsNil(o.InNetworksWithTag) {
		return nil, false
	}
	return o.InNetworksWithTag, true
}

// HasInNetworksWithTag returns a boolean if a field has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) HasInNetworksWithTag() bool {
	if o != nil && !IsNil(o.InNetworksWithTag) {
		return true
	}

	return false
}

// SetInNetworksWithTag gets a reference to the given string and assigns it to the InNetworksWithTag field.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) SetInNetworksWithTag(v string) {
	o.InNetworksWithTag = &v
}

// GetInNetworksWithId returns the InNetworksWithId field value if set, zero value otherwise.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetInNetworksWithId() string {
	if o == nil || IsNil(o.InNetworksWithId) {
		var ret string
		return ret
	}
	return *o.InNetworksWithId
}

// GetInNetworksWithIdOk returns a tuple with the InNetworksWithId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetInNetworksWithIdOk() (*string, bool) {
	if o == nil || IsNil(o.InNetworksWithId) {
		return nil, false
	}
	return o.InNetworksWithId, true
}

// HasInNetworksWithId returns a boolean if a field has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) HasInNetworksWithId() bool {
	if o != nil && !IsNil(o.InNetworksWithId) {
		return true
	}

	return false
}

// SetInNetworksWithId gets a reference to the given string and assigns it to the InNetworksWithId field.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) SetInNetworksWithId(v string) {
	o.InNetworksWithId = &v
}

// GetPermissionScopeId returns the PermissionScopeId field value
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetPermissionScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionScopeId
}

// GetPermissionScopeIdOk returns a tuple with the PermissionScopeId field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) GetPermissionScopeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionScopeId, true
}

// SetPermissionScopeId sets field value
func (o *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) SetPermissionScopeId(v string) {
	o.PermissionScopeId = v
}

func (o CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InNetworksWithTag) {
		toSerialize["inNetworksWithTag"] = o.InNetworksWithTag
	}
	if !IsNil(o.InNetworksWithId) {
		toSerialize["inNetworksWithId"] = o.InNetworksWithId
	}
	toSerialize["permissionScopeId"] = o.PermissionScopeId
	return toSerialize, nil
}

type NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner struct {
	value *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner
	isSet bool
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner) Get() *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner {
	return v.value
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner) Set(val *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner(val *CreateOrganizationCameraRoleRequestAppliedOnDevicesInner) *NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner {
	return &NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOnDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


