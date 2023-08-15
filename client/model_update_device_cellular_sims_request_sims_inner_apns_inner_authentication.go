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

// checks if the UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication{}

// UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication APN authentication configurations.
type UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication struct {
	// APN auth type.
	Type *string `json:"type,omitempty"`
	// APN username, if type is set.
	Username *string `json:"username,omitempty"`
	// APN password, if type is set (if APN password is not supplied, the password is left unchanged).
	Password *string `json:"password,omitempty"`
}

// NewUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication instantiates a new UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication() *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication {
	this := UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication{}
	var type_ string = "none"
	this.Type = &type_
	return &this
}

// NewUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthenticationWithDefaults instantiates a new UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthenticationWithDefaults() *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication {
	this := UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication{}
	var type_ string = "none"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication struct {
	value *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication
	isSet bool
}

func (v NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) Get() *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication {
	return v.value
}

func (v *NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) Set(val *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication(val *UpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) *NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication {
	return &NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularSimsRequestSimsInnerApnsInnerAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


