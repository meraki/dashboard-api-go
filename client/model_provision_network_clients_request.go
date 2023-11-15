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

// checks if the ProvisionNetworkClientsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionNetworkClientsRequest{}

// ProvisionNetworkClientsRequest struct for ProvisionNetworkClientsRequest
type ProvisionNetworkClientsRequest struct {
	// The array of clients to provision
	Clients []ProvisionNetworkClientsRequestClientsInner `json:"clients"`
	// The policy to apply to the specified client. Can be 'Group policy', 'Allowed', 'Blocked', 'Per connection' or 'Normal'. Required.
	DevicePolicy string `json:"devicePolicy"`
	// The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to \"Group policy\". Otherwise this is ignored.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	PoliciesBySecurityAppliance *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance `json:"policiesBySecurityAppliance,omitempty"`
	PoliciesBySsid *ProvisionNetworkClientsRequestPoliciesBySsid `json:"policiesBySsid,omitempty"`
}

// NewProvisionNetworkClientsRequest instantiates a new ProvisionNetworkClientsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionNetworkClientsRequest(clients []ProvisionNetworkClientsRequestClientsInner, devicePolicy string) *ProvisionNetworkClientsRequest {
	this := ProvisionNetworkClientsRequest{}
	this.Clients = clients
	this.DevicePolicy = devicePolicy
	return &this
}

// NewProvisionNetworkClientsRequestWithDefaults instantiates a new ProvisionNetworkClientsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionNetworkClientsRequestWithDefaults() *ProvisionNetworkClientsRequest {
	this := ProvisionNetworkClientsRequest{}
	return &this
}

// GetClients returns the Clients field value
func (o *ProvisionNetworkClientsRequest) GetClients() []ProvisionNetworkClientsRequestClientsInner {
	if o == nil {
		var ret []ProvisionNetworkClientsRequestClientsInner
		return ret
	}

	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequest) GetClientsOk() ([]ProvisionNetworkClientsRequestClientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clients, true
}

// SetClients sets field value
func (o *ProvisionNetworkClientsRequest) SetClients(v []ProvisionNetworkClientsRequestClientsInner) {
	o.Clients = v
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *ProvisionNetworkClientsRequest) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequest) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *ProvisionNetworkClientsRequest) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequest) GetGroupPolicyId() string {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequest) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequest) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *ProvisionNetworkClientsRequest) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetPoliciesBySecurityAppliance returns the PoliciesBySecurityAppliance field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequest) GetPoliciesBySecurityAppliance() ProvisionNetworkClientsRequestPoliciesBySecurityAppliance {
	if o == nil || IsNil(o.PoliciesBySecurityAppliance) {
		var ret ProvisionNetworkClientsRequestPoliciesBySecurityAppliance
		return ret
	}
	return *o.PoliciesBySecurityAppliance
}

// GetPoliciesBySecurityApplianceOk returns a tuple with the PoliciesBySecurityAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequest) GetPoliciesBySecurityApplianceOk() (*ProvisionNetworkClientsRequestPoliciesBySecurityAppliance, bool) {
	if o == nil || IsNil(o.PoliciesBySecurityAppliance) {
		return nil, false
	}
	return o.PoliciesBySecurityAppliance, true
}

// HasPoliciesBySecurityAppliance returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequest) HasPoliciesBySecurityAppliance() bool {
	if o != nil && !IsNil(o.PoliciesBySecurityAppliance) {
		return true
	}

	return false
}

// SetPoliciesBySecurityAppliance gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySecurityAppliance and assigns it to the PoliciesBySecurityAppliance field.
func (o *ProvisionNetworkClientsRequest) SetPoliciesBySecurityAppliance(v ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) {
	o.PoliciesBySecurityAppliance = &v
}

// GetPoliciesBySsid returns the PoliciesBySsid field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequest) GetPoliciesBySsid() ProvisionNetworkClientsRequestPoliciesBySsid {
	if o == nil || IsNil(o.PoliciesBySsid) {
		var ret ProvisionNetworkClientsRequestPoliciesBySsid
		return ret
	}
	return *o.PoliciesBySsid
}

// GetPoliciesBySsidOk returns a tuple with the PoliciesBySsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequest) GetPoliciesBySsidOk() (*ProvisionNetworkClientsRequestPoliciesBySsid, bool) {
	if o == nil || IsNil(o.PoliciesBySsid) {
		return nil, false
	}
	return o.PoliciesBySsid, true
}

// HasPoliciesBySsid returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequest) HasPoliciesBySsid() bool {
	if o != nil && !IsNil(o.PoliciesBySsid) {
		return true
	}

	return false
}

// SetPoliciesBySsid gets a reference to the given ProvisionNetworkClientsRequestPoliciesBySsid and assigns it to the PoliciesBySsid field.
func (o *ProvisionNetworkClientsRequest) SetPoliciesBySsid(v ProvisionNetworkClientsRequestPoliciesBySsid) {
	o.PoliciesBySsid = &v
}

func (o ProvisionNetworkClientsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionNetworkClientsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clients"] = o.Clients
	toSerialize["devicePolicy"] = o.DevicePolicy
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !IsNil(o.PoliciesBySecurityAppliance) {
		toSerialize["policiesBySecurityAppliance"] = o.PoliciesBySecurityAppliance
	}
	if !IsNil(o.PoliciesBySsid) {
		toSerialize["policiesBySsid"] = o.PoliciesBySsid
	}
	return toSerialize, nil
}

type NullableProvisionNetworkClientsRequest struct {
	value *ProvisionNetworkClientsRequest
	isSet bool
}

func (v NullableProvisionNetworkClientsRequest) Get() *ProvisionNetworkClientsRequest {
	return v.value
}

func (v *NullableProvisionNetworkClientsRequest) Set(val *ProvisionNetworkClientsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionNetworkClientsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionNetworkClientsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionNetworkClientsRequest(val *ProvisionNetworkClientsRequest) *NullableProvisionNetworkClientsRequest {
	return &NullableProvisionNetworkClientsRequest{value: val, isSet: true}
}

func (v NullableProvisionNetworkClientsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionNetworkClientsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


