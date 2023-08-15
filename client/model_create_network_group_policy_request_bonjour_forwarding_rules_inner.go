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

// checks if the CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner{}

// CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner struct for CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner
type CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner struct {
	// A description for your Bonjour forwarding rule. Optional.
	Description *string `json:"description,omitempty"`
	// The ID of the service VLAN. Required.
	VlanId string `json:"vlanId"`
	// A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'
	Services []string `json:"services"`
}

// NewCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner instantiates a new CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner(vlanId string, services []string) *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner {
	this := CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner{}
	this.VlanId = vlanId
	this.Services = services
	return &this
}

// NewCreateNetworkGroupPolicyRequestBonjourForwardingRulesInnerWithDefaults instantiates a new CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestBonjourForwardingRulesInnerWithDefaults() *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner {
	this := CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) SetDescription(v string) {
	o.Description = &v
}

// GetVlanId returns the VlanId field value
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) GetVlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) GetVlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) SetVlanId(v string) {
	o.VlanId = v
}

// GetServices returns the Services field value
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) GetServicesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Services, true
}

// SetServices sets field value
func (o *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) SetServices(v []string) {
	o.Services = v
}

func (o CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["vlanId"] = o.VlanId
	toSerialize["services"] = o.Services
	return toSerialize, nil
}

type NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner struct {
	value *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) Get() *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) Set(val *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner(val *CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) *NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner {
	return &NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


