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

// checks if the UpdateNetworkApplianceFirewallFirewalledServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallFirewalledServiceRequest{}

// UpdateNetworkApplianceFirewallFirewalledServiceRequest struct for UpdateNetworkApplianceFirewallFirewalledServiceRequest
type UpdateNetworkApplianceFirewallFirewalledServiceRequest struct {
	// A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are \"blocked\" (no remote IPs can access the service), \"restricted\" (only allowed IPs can access the service), and \"unrestriced\" (any remote IP can access the service). This field is required
	Access string `json:"access"`
	// An array of allowed IPs that can access the service. This field is required if \"access\" is set to \"restricted\". Otherwise this field is ignored
	AllowedIps []string `json:"allowedIps,omitempty"`
}

// NewUpdateNetworkApplianceFirewallFirewalledServiceRequest instantiates a new UpdateNetworkApplianceFirewallFirewalledServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallFirewalledServiceRequest(access string) *UpdateNetworkApplianceFirewallFirewalledServiceRequest {
	this := UpdateNetworkApplianceFirewallFirewalledServiceRequest{}
	this.Access = access
	return &this
}

// NewUpdateNetworkApplianceFirewallFirewalledServiceRequestWithDefaults instantiates a new UpdateNetworkApplianceFirewallFirewalledServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallFirewalledServiceRequestWithDefaults() *UpdateNetworkApplianceFirewallFirewalledServiceRequest {
	this := UpdateNetworkApplianceFirewallFirewalledServiceRequest{}
	return &this
}

// GetAccess returns the Access field value
func (o *UpdateNetworkApplianceFirewallFirewalledServiceRequest) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallFirewalledServiceRequest) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *UpdateNetworkApplianceFirewallFirewalledServiceRequest) SetAccess(v string) {
	o.Access = v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallFirewalledServiceRequest) GetAllowedIps() []string {
	if o == nil || IsNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallFirewalledServiceRequest) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedIps) {
		return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallFirewalledServiceRequest) HasAllowedIps() bool {
	if o != nil && !IsNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *UpdateNetworkApplianceFirewallFirewalledServiceRequest) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

func (o UpdateNetworkApplianceFirewallFirewalledServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallFirewalledServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access"] = o.Access
	if !IsNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest struct {
	value *UpdateNetworkApplianceFirewallFirewalledServiceRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest) Get() *UpdateNetworkApplianceFirewallFirewalledServiceRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest) Set(val *UpdateNetworkApplianceFirewallFirewalledServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallFirewalledServiceRequest(val *UpdateNetworkApplianceFirewallFirewalledServiceRequest) *NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest {
	return &NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallFirewalledServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


