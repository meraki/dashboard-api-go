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

// checks if the UpdateNetworkMerakiAuthUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkMerakiAuthUserRequest{}

// UpdateNetworkMerakiAuthUserRequest struct for UpdateNetworkMerakiAuthUserRequest
type UpdateNetworkMerakiAuthUserRequest struct {
	// Name of the user. Only allowed If the user is not Dashboard administrator.
	Name *string `json:"name,omitempty"`
	// The password for this user account. Only allowed If the user is not Dashboard administrator.
	Password *string `json:"password,omitempty"`
	// Whether or not Meraki should email the password to user. Default is false.
	EmailPasswordToUser *bool `json:"emailPasswordToUser,omitempty"`
	// Authorization zones and expiration dates for the user.
	Authorizations []UpdateNetworkMerakiAuthUserRequestAuthorizationsInner `json:"authorizations,omitempty"`
}

// NewUpdateNetworkMerakiAuthUserRequest instantiates a new UpdateNetworkMerakiAuthUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkMerakiAuthUserRequest() *UpdateNetworkMerakiAuthUserRequest {
	this := UpdateNetworkMerakiAuthUserRequest{}
	return &this
}

// NewUpdateNetworkMerakiAuthUserRequestWithDefaults instantiates a new UpdateNetworkMerakiAuthUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkMerakiAuthUserRequestWithDefaults() *UpdateNetworkMerakiAuthUserRequest {
	this := UpdateNetworkMerakiAuthUserRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkMerakiAuthUserRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkMerakiAuthUserRequest) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateNetworkMerakiAuthUserRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateNetworkMerakiAuthUserRequest) SetPassword(v string) {
	o.Password = &v
}

// GetEmailPasswordToUser returns the EmailPasswordToUser field value if set, zero value otherwise.
func (o *UpdateNetworkMerakiAuthUserRequest) GetEmailPasswordToUser() bool {
	if o == nil || IsNil(o.EmailPasswordToUser) {
		var ret bool
		return ret
	}
	return *o.EmailPasswordToUser
}

// GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) GetEmailPasswordToUserOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailPasswordToUser) {
		return nil, false
	}
	return o.EmailPasswordToUser, true
}

// HasEmailPasswordToUser returns a boolean if a field has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) HasEmailPasswordToUser() bool {
	if o != nil && !IsNil(o.EmailPasswordToUser) {
		return true
	}

	return false
}

// SetEmailPasswordToUser gets a reference to the given bool and assigns it to the EmailPasswordToUser field.
func (o *UpdateNetworkMerakiAuthUserRequest) SetEmailPasswordToUser(v bool) {
	o.EmailPasswordToUser = &v
}

// GetAuthorizations returns the Authorizations field value if set, zero value otherwise.
func (o *UpdateNetworkMerakiAuthUserRequest) GetAuthorizations() []UpdateNetworkMerakiAuthUserRequestAuthorizationsInner {
	if o == nil || IsNil(o.Authorizations) {
		var ret []UpdateNetworkMerakiAuthUserRequestAuthorizationsInner
		return ret
	}
	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) GetAuthorizationsOk() ([]UpdateNetworkMerakiAuthUserRequestAuthorizationsInner, bool) {
	if o == nil || IsNil(o.Authorizations) {
		return nil, false
	}
	return o.Authorizations, true
}

// HasAuthorizations returns a boolean if a field has been set.
func (o *UpdateNetworkMerakiAuthUserRequest) HasAuthorizations() bool {
	if o != nil && !IsNil(o.Authorizations) {
		return true
	}

	return false
}

// SetAuthorizations gets a reference to the given []UpdateNetworkMerakiAuthUserRequestAuthorizationsInner and assigns it to the Authorizations field.
func (o *UpdateNetworkMerakiAuthUserRequest) SetAuthorizations(v []UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) {
	o.Authorizations = v
}

func (o UpdateNetworkMerakiAuthUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkMerakiAuthUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.EmailPasswordToUser) {
		toSerialize["emailPasswordToUser"] = o.EmailPasswordToUser
	}
	if !IsNil(o.Authorizations) {
		toSerialize["authorizations"] = o.Authorizations
	}
	return toSerialize, nil
}

type NullableUpdateNetworkMerakiAuthUserRequest struct {
	value *UpdateNetworkMerakiAuthUserRequest
	isSet bool
}

func (v NullableUpdateNetworkMerakiAuthUserRequest) Get() *UpdateNetworkMerakiAuthUserRequest {
	return v.value
}

func (v *NullableUpdateNetworkMerakiAuthUserRequest) Set(val *UpdateNetworkMerakiAuthUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkMerakiAuthUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkMerakiAuthUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkMerakiAuthUserRequest(val *UpdateNetworkMerakiAuthUserRequest) *NullableUpdateNetworkMerakiAuthUserRequest {
	return &NullableUpdateNetworkMerakiAuthUserRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkMerakiAuthUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkMerakiAuthUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


