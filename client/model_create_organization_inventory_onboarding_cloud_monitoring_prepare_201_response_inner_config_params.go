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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams Params used in order to connect to the device
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams struct {
	Tunnel *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel `json:"tunnel,omitempty"`
	// Static IP Address used to connect to the device
	CloudStaticIp *string `json:"cloudStaticIp,omitempty"`
	User *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser `json:"user,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams{}
	return &this
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) GetTunnel() CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel {
	if o == nil || IsNil(o.Tunnel) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) GetTunnelOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel, bool) {
	if o == nil || IsNil(o.Tunnel) {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) HasTunnel() bool {
	if o != nil && !IsNil(o.Tunnel) {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel and assigns it to the Tunnel field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) SetTunnel(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnel) {
	o.Tunnel = &v
}

// GetCloudStaticIp returns the CloudStaticIp field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) GetCloudStaticIp() string {
	if o == nil || IsNil(o.CloudStaticIp) {
		var ret string
		return ret
	}
	return *o.CloudStaticIp
}

// GetCloudStaticIpOk returns a tuple with the CloudStaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) GetCloudStaticIpOk() (*string, bool) {
	if o == nil || IsNil(o.CloudStaticIp) {
		return nil, false
	}
	return o.CloudStaticIp, true
}

// HasCloudStaticIp returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) HasCloudStaticIp() bool {
	if o != nil && !IsNil(o.CloudStaticIp) {
		return true
	}

	return false
}

// SetCloudStaticIp gets a reference to the given string and assigns it to the CloudStaticIp field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) SetCloudStaticIp(v string) {
	o.CloudStaticIp = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) GetUser() CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser {
	if o == nil || IsNil(o.User) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) GetUserOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser and assigns it to the User field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) SetUser(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsUser) {
	o.User = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tunnel) {
		toSerialize["tunnel"] = o.Tunnel
	}
	if !IsNil(o.CloudStaticIp) {
		toSerialize["cloudStaticIp"] = o.CloudStaticIp
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


