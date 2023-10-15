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

// checks if the GetOrganizationAdmins200ResponseInnerNetworksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdmins200ResponseInnerNetworksInner{}

// GetOrganizationAdmins200ResponseInnerNetworksInner struct for GetOrganizationAdmins200ResponseInnerNetworksInner
type GetOrganizationAdmins200ResponseInnerNetworksInner struct {
	// Network ID
	Id *string `json:"id,omitempty"`
	// Admin's level of access to the network
	Access *string `json:"access,omitempty"`
}

// NewGetOrganizationAdmins200ResponseInnerNetworksInner instantiates a new GetOrganizationAdmins200ResponseInnerNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdmins200ResponseInnerNetworksInner() *GetOrganizationAdmins200ResponseInnerNetworksInner {
	this := GetOrganizationAdmins200ResponseInnerNetworksInner{}
	return &this
}

// NewGetOrganizationAdmins200ResponseInnerNetworksInnerWithDefaults instantiates a new GetOrganizationAdmins200ResponseInnerNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdmins200ResponseInnerNetworksInnerWithDefaults() *GetOrganizationAdmins200ResponseInnerNetworksInner {
	this := GetOrganizationAdmins200ResponseInnerNetworksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) SetId(v string) {
	o.Id = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *GetOrganizationAdmins200ResponseInnerNetworksInner) SetAccess(v string) {
	o.Access = &v
}

func (o GetOrganizationAdmins200ResponseInnerNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdmins200ResponseInnerNetworksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdmins200ResponseInnerNetworksInner struct {
	value *GetOrganizationAdmins200ResponseInnerNetworksInner
	isSet bool
}

func (v NullableGetOrganizationAdmins200ResponseInnerNetworksInner) Get() *GetOrganizationAdmins200ResponseInnerNetworksInner {
	return v.value
}

func (v *NullableGetOrganizationAdmins200ResponseInnerNetworksInner) Set(val *GetOrganizationAdmins200ResponseInnerNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdmins200ResponseInnerNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdmins200ResponseInnerNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdmins200ResponseInnerNetworksInner(val *GetOrganizationAdmins200ResponseInnerNetworksInner) *NullableGetOrganizationAdmins200ResponseInnerNetworksInner {
	return &NullableGetOrganizationAdmins200ResponseInnerNetworksInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAdmins200ResponseInnerNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdmins200ResponseInnerNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


