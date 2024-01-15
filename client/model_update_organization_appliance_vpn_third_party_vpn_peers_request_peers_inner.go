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

// checks if the UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner{}

// UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner struct for UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner
type UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner struct {
	// The name of the VPN peer
	Name string `json:"name"`
	// [optional] The public IP of the VPN peer
	PublicIp *string `json:"publicIp,omitempty"`
	// The list of the private subnets of the VPN peer
	PrivateSubnets []string `json:"privateSubnets"`
	// [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	LocalId *string `json:"localId,omitempty"`
	// [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
	RemoteId *string `json:"remoteId,omitempty"`
	IpsecPolicies *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies `json:"ipsecPolicies,omitempty"`
	// One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	IpsecPoliciesPreset *string `json:"ipsecPoliciesPreset,omitempty"`
	// The shared secret with the VPN peer
	Secret string `json:"secret"`
	// [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IkeVersion *string `json:"ikeVersion,omitempty"`
	// A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	NetworkTags []string `json:"networkTags,omitempty"`
}

// NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner(name string, privateSubnets []string, secret string) *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner {
	this := UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner{}
	this.Name = name
	this.PrivateSubnets = privateSubnets
	this.Secret = secret
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerWithDefaults instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerWithDefaults() *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner {
	this := UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner{}
	var ikeVersion string = "1"
	this.IkeVersion = &ikeVersion
	return &this
}

// GetName returns the Name field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetName(v string) {
	o.Name = v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPrivateSubnets returns the PrivateSubnets field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPrivateSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPrivateSubnetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateSubnets, true
}

// SetPrivateSubnets sets field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = v
}

// GetLocalId returns the LocalId field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetLocalId() string {
	if o == nil || IsNil(o.LocalId) {
		var ret string
		return ret
	}
	return *o.LocalId
}

// GetLocalIdOk returns a tuple with the LocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetLocalIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocalId) {
		return nil, false
	}
	return o.LocalId, true
}

// HasLocalId returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasLocalId() bool {
	if o != nil && !IsNil(o.LocalId) {
		return true
	}

	return false
}

// SetLocalId gets a reference to the given string and assigns it to the LocalId field.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetLocalId(v string) {
	o.LocalId = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetIpsecPolicies returns the IpsecPolicies field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPolicies() GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies {
	if o == nil || IsNil(o.IpsecPolicies) {
		var ret GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies
		return ret
	}
	return *o.IpsecPolicies
}

// GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPoliciesOk() (*GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies, bool) {
	if o == nil || IsNil(o.IpsecPolicies) {
		return nil, false
	}
	return o.IpsecPolicies, true
}

// HasIpsecPolicies returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasIpsecPolicies() bool {
	if o != nil && !IsNil(o.IpsecPolicies) {
		return true
	}

	return false
}

// SetIpsecPolicies gets a reference to the given GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies and assigns it to the IpsecPolicies field.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetIpsecPolicies(v GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies) {
	o.IpsecPolicies = &v
}

// GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPoliciesPreset() string {
	if o == nil || IsNil(o.IpsecPoliciesPreset) {
		var ret string
		return ret
	}
	return *o.IpsecPoliciesPreset
}

// GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPoliciesPresetOk() (*string, bool) {
	if o == nil || IsNil(o.IpsecPoliciesPreset) {
		return nil, false
	}
	return o.IpsecPoliciesPreset, true
}

// HasIpsecPoliciesPreset returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasIpsecPoliciesPreset() bool {
	if o != nil && !IsNil(o.IpsecPoliciesPreset) {
		return true
	}

	return false
}

// SetIpsecPoliciesPreset gets a reference to the given string and assigns it to the IpsecPoliciesPreset field.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetIpsecPoliciesPreset(v string) {
	o.IpsecPoliciesPreset = &v
}

// GetSecret returns the Secret field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetSecret(v string) {
	o.Secret = v
}

// GetIkeVersion returns the IkeVersion field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIkeVersion() string {
	if o == nil || IsNil(o.IkeVersion) {
		var ret string
		return ret
	}
	return *o.IkeVersion
}

// GetIkeVersionOk returns a tuple with the IkeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIkeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.IkeVersion) {
		return nil, false
	}
	return o.IkeVersion, true
}

// HasIkeVersion returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasIkeVersion() bool {
	if o != nil && !IsNil(o.IkeVersion) {
		return true
	}

	return false
}

// SetIkeVersion gets a reference to the given string and assigns it to the IkeVersion field.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetIkeVersion(v string) {
	o.IkeVersion = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetNetworkTags() []string {
	if o == nil || IsNil(o.NetworkTags) {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkTags) {
		return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasNetworkTags() bool {
	if o != nil && !IsNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

func (o UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	toSerialize["privateSubnets"] = o.PrivateSubnets
	if !IsNil(o.LocalId) {
		toSerialize["localId"] = o.LocalId
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	if !IsNil(o.IpsecPolicies) {
		toSerialize["ipsecPolicies"] = o.IpsecPolicies
	}
	if !IsNil(o.IpsecPoliciesPreset) {
		toSerialize["ipsecPoliciesPreset"] = o.IpsecPoliciesPreset
	}
	toSerialize["secret"] = o.Secret
	if !IsNil(o.IkeVersion) {
		toSerialize["ikeVersion"] = o.IkeVersion
	}
	if !IsNil(o.NetworkTags) {
		toSerialize["networkTags"] = o.NetworkTags
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner struct {
	value *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner
	isSet bool
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) Get() *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner {
	return v.value
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) Set(val *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner(val *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner {
	return &NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner{value: val, isSet: true}
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


