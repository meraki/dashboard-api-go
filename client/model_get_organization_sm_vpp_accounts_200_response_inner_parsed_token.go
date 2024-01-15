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

// checks if the GetOrganizationSmVppAccounts200ResponseInnerParsedToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSmVppAccounts200ResponseInnerParsedToken{}

// GetOrganizationSmVppAccounts200ResponseInnerParsedToken The parsed VPP service token
type GetOrganizationSmVppAccounts200ResponseInnerParsedToken struct {
	// The organization name
	OrgName *string `json:"orgName,omitempty"`
	// The hashed token
	HashedToken *string `json:"hashedToken,omitempty"`
	// The expiration time of the token
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// NewGetOrganizationSmVppAccounts200ResponseInnerParsedToken instantiates a new GetOrganizationSmVppAccounts200ResponseInnerParsedToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSmVppAccounts200ResponseInnerParsedToken() *GetOrganizationSmVppAccounts200ResponseInnerParsedToken {
	this := GetOrganizationSmVppAccounts200ResponseInnerParsedToken{}
	return &this
}

// NewGetOrganizationSmVppAccounts200ResponseInnerParsedTokenWithDefaults instantiates a new GetOrganizationSmVppAccounts200ResponseInnerParsedToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSmVppAccounts200ResponseInnerParsedTokenWithDefaults() *GetOrganizationSmVppAccounts200ResponseInnerParsedToken {
	this := GetOrganizationSmVppAccounts200ResponseInnerParsedToken{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) SetOrgName(v string) {
	o.OrgName = &v
}

// GetHashedToken returns the HashedToken field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) GetHashedToken() string {
	if o == nil || IsNil(o.HashedToken) {
		var ret string
		return ret
	}
	return *o.HashedToken
}

// GetHashedTokenOk returns a tuple with the HashedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) GetHashedTokenOk() (*string, bool) {
	if o == nil || IsNil(o.HashedToken) {
		return nil, false
	}
	return o.HashedToken, true
}

// HasHashedToken returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) HasHashedToken() bool {
	if o != nil && !IsNil(o.HashedToken) {
		return true
	}

	return false
}

// SetHashedToken gets a reference to the given string and assigns it to the HashedToken field.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) SetHashedToken(v string) {
	o.HashedToken = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o GetOrganizationSmVppAccounts200ResponseInnerParsedToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSmVppAccounts200ResponseInnerParsedToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !IsNil(o.HashedToken) {
		toSerialize["hashedToken"] = o.HashedToken
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken struct {
	value *GetOrganizationSmVppAccounts200ResponseInnerParsedToken
	isSet bool
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken) Get() *GetOrganizationSmVppAccounts200ResponseInnerParsedToken {
	return v.value
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken) Set(val *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken(val *GetOrganizationSmVppAccounts200ResponseInnerParsedToken) *NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken {
	return &NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken{value: val, isSet: true}
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInnerParsedToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


