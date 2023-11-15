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

// checks if the UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner{}

// UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner struct for UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner
type UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner struct {
	// Protocol for this split tunnel rule.
	Protocol *string `json:"protocol,omitempty"`
	// Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or 'any'.
	DestCidr string `json:"destCidr"`
	// Destination port for this split tunnel rule, (integer in the range 1-65535), or 'any'.
	DestPort *string `json:"destPort,omitempty"`
	// Traffic policy specified for this split tunnel rule, 'allow' or 'deny'.
	Policy string `json:"policy"`
	// Description for this split tunnel rule (optional).
	Comment *string `json:"comment,omitempty"`
}

// NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner(destCidr string, policy string) *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner {
	this := UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner{}
	this.DestCidr = destCidr
	this.Policy = policy
	return &this
}

// NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInnerWithDefaults() *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner {
	this := UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetDestCidr returns the DestCidr field value
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestCidr, true
}

// SetDestCidr sets field value
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetDestCidr(v string) {
	o.DestCidr = v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestPort() string {
	if o == nil || IsNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestPortOk() (*string, bool) {
	if o == nil || IsNil(o.DestPort) {
		return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) HasDestPort() bool {
	if o != nil && !IsNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetDestPort(v string) {
	o.DestPort = &v
}

// GetPolicy returns the Policy field value
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetPolicy(v string) {
	o.Policy = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetComment(v string) {
	o.Comment = &v
}

func (o UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	toSerialize["destCidr"] = o.DestCidr
	if !IsNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	toSerialize["policy"] = o.Policy
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner struct {
	value *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) Get() *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) Set(val *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner(val *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner {
	return &NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


