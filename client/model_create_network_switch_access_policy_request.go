/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkSwitchAccessPolicyRequest struct for CreateNetworkSwitchAccessPolicyRequest
type CreateNetworkSwitchAccessPolicyRequest struct {
	// Name of the access policy
	Name string `json:"name"`
	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusServers []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner `json:"radiusServers"`
	Radius *GetNetworkSwitchAccessPolicies200ResponseInnerRadius `json:"radius,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	RadiusTestingEnabled bool `json:"radiusTestingEnabled"`
	// Change of authentication for RADIUS re-authentication and disconnection
	RadiusCoaSupportEnabled bool `json:"radiusCoaSupportEnabled"`
	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingEnabled bool `json:"radiusAccountingEnabled"`
	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner `json:"radiusAccountingServers,omitempty"`
	// Acceptable values are `\"\"` for None, or `\"11\"` for Group Policies ACL
	RadiusGroupAttribute *string `json:"radiusGroupAttribute,omitempty"`
	// Choose the Host Mode for the access policy.
	HostMode string `json:"hostMode"`
	// Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	IncreaseAccessSpeed *bool `json:"increaseAccessSpeed,omitempty"`
	// ID for the guest VLAN allow unauthorized devices access to limited network resources
	GuestVlanId *int32 `json:"guestVlanId,omitempty"`
	Dot1x *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x `json:"dot1x,omitempty"`
	// CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
	VoiceVlanClients *bool `json:"voiceVlanClients,omitempty"`
	// Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenEnabled bool `json:"urlRedirectWalledGardenEnabled"`
	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges,omitempty"`
}

// NewCreateNetworkSwitchAccessPolicyRequest instantiates a new CreateNetworkSwitchAccessPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchAccessPolicyRequest(name string, radiusServers []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner, radiusTestingEnabled bool, radiusCoaSupportEnabled bool, radiusAccountingEnabled bool, hostMode string, urlRedirectWalledGardenEnabled bool) *CreateNetworkSwitchAccessPolicyRequest {
	this := CreateNetworkSwitchAccessPolicyRequest{}
	this.Name = name
	this.RadiusServers = radiusServers
	this.RadiusTestingEnabled = radiusTestingEnabled
	this.RadiusCoaSupportEnabled = radiusCoaSupportEnabled
	this.RadiusAccountingEnabled = radiusAccountingEnabled
	this.HostMode = hostMode
	this.UrlRedirectWalledGardenEnabled = urlRedirectWalledGardenEnabled
	return &this
}

// NewCreateNetworkSwitchAccessPolicyRequestWithDefaults instantiates a new CreateNetworkSwitchAccessPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchAccessPolicyRequestWithDefaults() *CreateNetworkSwitchAccessPolicyRequest {
	this := CreateNetworkSwitchAccessPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkSwitchAccessPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkSwitchAccessPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetRadiusServers returns the RadiusServers field value
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusServers() []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner {
	if o == nil {
		var ret []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner
		return ret
	}

	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusServersOk() ([]CreateNetworkSwitchAccessPolicyRequestRadiusServersInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.RadiusServers, true
}

// SetRadiusServers sets field value
func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusServers(v []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) {
	o.RadiusServers = v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadius() GetNetworkSwitchAccessPolicies200ResponseInnerRadius {
	if o == nil || isNil(o.Radius) {
		var ret GetNetworkSwitchAccessPolicies200ResponseInnerRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerRadius, bool) {
	if o == nil || isNil(o.Radius) {
    return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasRadius() bool {
	if o != nil && !isNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given GetNetworkSwitchAccessPolicies200ResponseInnerRadius and assigns it to the Radius field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadius(v GetNetworkSwitchAccessPolicies200ResponseInnerRadius) {
	o.Radius = &v
}

// GetRadiusTestingEnabled returns the RadiusTestingEnabled field value
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusTestingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RadiusTestingEnabled
}

// GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusTestingEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RadiusTestingEnabled, true
}

// SetRadiusTestingEnabled sets field value
func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusTestingEnabled(v bool) {
	o.RadiusTestingEnabled = v
}

// GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field value
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusCoaSupportEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RadiusCoaSupportEnabled
}

// GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusCoaSupportEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RadiusCoaSupportEnabled, true
}

// SetRadiusCoaSupportEnabled sets field value
func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusCoaSupportEnabled(v bool) {
	o.RadiusCoaSupportEnabled = v
}

// GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field value
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RadiusAccountingEnabled
}

// GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RadiusAccountingEnabled, true
}

// SetRadiusAccountingEnabled sets field value
func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusAccountingEnabled(v bool) {
	o.RadiusAccountingEnabled = v
}

// GetRadiusAccountingServers returns the RadiusAccountingServers field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingServers() []CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner {
	if o == nil || isNil(o.RadiusAccountingServers) {
		var ret []CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner
		return ret
	}
	return o.RadiusAccountingServers
}

// GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingServersOk() ([]CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner, bool) {
	if o == nil || isNil(o.RadiusAccountingServers) {
    return nil, false
	}
	return o.RadiusAccountingServers, true
}

// HasRadiusAccountingServers returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasRadiusAccountingServers() bool {
	if o != nil && !isNil(o.RadiusAccountingServers) {
		return true
	}

	return false
}

// SetRadiusAccountingServers gets a reference to the given []CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner and assigns it to the RadiusAccountingServers field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusAccountingServers(v []CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) {
	o.RadiusAccountingServers = v
}

// GetRadiusGroupAttribute returns the RadiusGroupAttribute field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusGroupAttribute() string {
	if o == nil || isNil(o.RadiusGroupAttribute) {
		var ret string
		return ret
	}
	return *o.RadiusGroupAttribute
}

// GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusGroupAttributeOk() (*string, bool) {
	if o == nil || isNil(o.RadiusGroupAttribute) {
    return nil, false
	}
	return o.RadiusGroupAttribute, true
}

// HasRadiusGroupAttribute returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasRadiusGroupAttribute() bool {
	if o != nil && !isNil(o.RadiusGroupAttribute) {
		return true
	}

	return false
}

// SetRadiusGroupAttribute gets a reference to the given string and assigns it to the RadiusGroupAttribute field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusGroupAttribute(v string) {
	o.RadiusGroupAttribute = &v
}

// GetHostMode returns the HostMode field value
func (o *CreateNetworkSwitchAccessPolicyRequest) GetHostMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostMode
}

// GetHostModeOk returns a tuple with the HostMode field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetHostModeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HostMode, true
}

// SetHostMode sets field value
func (o *CreateNetworkSwitchAccessPolicyRequest) SetHostMode(v string) {
	o.HostMode = v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetAccessPolicyType() string {
	if o == nil || isNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessPolicyType) {
    return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasAccessPolicyType() bool {
	if o != nil && !isNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetIncreaseAccessSpeed() bool {
	if o == nil || isNil(o.IncreaseAccessSpeed) {
		var ret bool
		return ret
	}
	return *o.IncreaseAccessSpeed
}

// GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetIncreaseAccessSpeedOk() (*bool, bool) {
	if o == nil || isNil(o.IncreaseAccessSpeed) {
    return nil, false
	}
	return o.IncreaseAccessSpeed, true
}

// HasIncreaseAccessSpeed returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasIncreaseAccessSpeed() bool {
	if o != nil && !isNil(o.IncreaseAccessSpeed) {
		return true
	}

	return false
}

// SetIncreaseAccessSpeed gets a reference to the given bool and assigns it to the IncreaseAccessSpeed field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetIncreaseAccessSpeed(v bool) {
	o.IncreaseAccessSpeed = &v
}

// GetGuestVlanId returns the GuestVlanId field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetGuestVlanId() int32 {
	if o == nil || isNil(o.GuestVlanId) {
		var ret int32
		return ret
	}
	return *o.GuestVlanId
}

// GetGuestVlanIdOk returns a tuple with the GuestVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetGuestVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.GuestVlanId) {
    return nil, false
	}
	return o.GuestVlanId, true
}

// HasGuestVlanId returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasGuestVlanId() bool {
	if o != nil && !isNil(o.GuestVlanId) {
		return true
	}

	return false
}

// SetGuestVlanId gets a reference to the given int32 and assigns it to the GuestVlanId field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetGuestVlanId(v int32) {
	o.GuestVlanId = &v
}

// GetDot1x returns the Dot1x field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetDot1x() GetNetworkSwitchAccessPolicies200ResponseInnerDot1x {
	if o == nil || isNil(o.Dot1x) {
		var ret GetNetworkSwitchAccessPolicies200ResponseInnerDot1x
		return ret
	}
	return *o.Dot1x
}

// GetDot1xOk returns a tuple with the Dot1x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetDot1xOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerDot1x, bool) {
	if o == nil || isNil(o.Dot1x) {
    return nil, false
	}
	return o.Dot1x, true
}

// HasDot1x returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasDot1x() bool {
	if o != nil && !isNil(o.Dot1x) {
		return true
	}

	return false
}

// SetDot1x gets a reference to the given GetNetworkSwitchAccessPolicies200ResponseInnerDot1x and assigns it to the Dot1x field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetDot1x(v GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) {
	o.Dot1x = &v
}

// GetVoiceVlanClients returns the VoiceVlanClients field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetVoiceVlanClients() bool {
	if o == nil || isNil(o.VoiceVlanClients) {
		var ret bool
		return ret
	}
	return *o.VoiceVlanClients
}

// GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetVoiceVlanClientsOk() (*bool, bool) {
	if o == nil || isNil(o.VoiceVlanClients) {
    return nil, false
	}
	return o.VoiceVlanClients, true
}

// HasVoiceVlanClients returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasVoiceVlanClients() bool {
	if o != nil && !isNil(o.VoiceVlanClients) {
		return true
	}

	return false
}

// SetVoiceVlanClients gets a reference to the given bool and assigns it to the VoiceVlanClients field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetVoiceVlanClients(v bool) {
	o.VoiceVlanClients = &v
}

// GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field value
func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UrlRedirectWalledGardenEnabled
}

// GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UrlRedirectWalledGardenEnabled, true
}

// SetUrlRedirectWalledGardenEnabled sets field value
func (o *CreateNetworkSwitchAccessPolicyRequest) SetUrlRedirectWalledGardenEnabled(v bool) {
	o.UrlRedirectWalledGardenEnabled = v
}

// GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field value if set, zero value otherwise.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenRanges() []string {
	if o == nil || isNil(o.UrlRedirectWalledGardenRanges) {
		var ret []string
		return ret
	}
	return o.UrlRedirectWalledGardenRanges
}

// GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenRangesOk() ([]string, bool) {
	if o == nil || isNil(o.UrlRedirectWalledGardenRanges) {
    return nil, false
	}
	return o.UrlRedirectWalledGardenRanges, true
}

// HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.
func (o *CreateNetworkSwitchAccessPolicyRequest) HasUrlRedirectWalledGardenRanges() bool {
	if o != nil && !isNil(o.UrlRedirectWalledGardenRanges) {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenRanges gets a reference to the given []string and assigns it to the UrlRedirectWalledGardenRanges field.
func (o *CreateNetworkSwitchAccessPolicyRequest) SetUrlRedirectWalledGardenRanges(v []string) {
	o.UrlRedirectWalledGardenRanges = v
}

func (o CreateNetworkSwitchAccessPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if !isNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if true {
		toSerialize["radiusTestingEnabled"] = o.RadiusTestingEnabled
	}
	if true {
		toSerialize["radiusCoaSupportEnabled"] = o.RadiusCoaSupportEnabled
	}
	if true {
		toSerialize["radiusAccountingEnabled"] = o.RadiusAccountingEnabled
	}
	if !isNil(o.RadiusAccountingServers) {
		toSerialize["radiusAccountingServers"] = o.RadiusAccountingServers
	}
	if !isNil(o.RadiusGroupAttribute) {
		toSerialize["radiusGroupAttribute"] = o.RadiusGroupAttribute
	}
	if true {
		toSerialize["hostMode"] = o.HostMode
	}
	if !isNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !isNil(o.IncreaseAccessSpeed) {
		toSerialize["increaseAccessSpeed"] = o.IncreaseAccessSpeed
	}
	if !isNil(o.GuestVlanId) {
		toSerialize["guestVlanId"] = o.GuestVlanId
	}
	if !isNil(o.Dot1x) {
		toSerialize["dot1x"] = o.Dot1x
	}
	if !isNil(o.VoiceVlanClients) {
		toSerialize["voiceVlanClients"] = o.VoiceVlanClients
	}
	if true {
		toSerialize["urlRedirectWalledGardenEnabled"] = o.UrlRedirectWalledGardenEnabled
	}
	if !isNil(o.UrlRedirectWalledGardenRanges) {
		toSerialize["urlRedirectWalledGardenRanges"] = o.UrlRedirectWalledGardenRanges
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkSwitchAccessPolicyRequest struct {
	value *CreateNetworkSwitchAccessPolicyRequest
	isSet bool
}

func (v NullableCreateNetworkSwitchAccessPolicyRequest) Get() *CreateNetworkSwitchAccessPolicyRequest {
	return v.value
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequest) Set(val *CreateNetworkSwitchAccessPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchAccessPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchAccessPolicyRequest(val *CreateNetworkSwitchAccessPolicyRequest) *NullableCreateNetworkSwitchAccessPolicyRequest {
	return &NullableCreateNetworkSwitchAccessPolicyRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchAccessPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


