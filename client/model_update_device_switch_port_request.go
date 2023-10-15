/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDeviceSwitchPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceSwitchPortRequest{}

// UpdateDeviceSwitchPortRequest struct for UpdateDeviceSwitchPortRequest
type UpdateDeviceSwitchPortRequest struct {
	// The name of the switch port.
	Name *string `json:"name,omitempty"`
	// The list of tags of the switch port.
	Tags []string `json:"tags,omitempty"`
	// The status of the switch port.
	Enabled *bool `json:"enabled,omitempty"`
	// The PoE status of the switch port.
	PoeEnabled *bool `json:"poeEnabled,omitempty"`
	// The type of the switch port ('trunk' or 'access').
	Type *string `json:"type,omitempty"`
	// The VLAN of the switch port. A null value will clear the value set for trunk ports.
	Vlan *int32 `json:"vlan,omitempty"`
	// The voice VLAN of the switch port. Only applicable to access ports.
	VoiceVlan *int32 `json:"voiceVlan,omitempty"`
	// The VLANs allowed on the switch port. Only applicable to trunk ports.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The isolation status of the switch port.
	IsolationEnabled *bool `json:"isolationEnabled,omitempty"`
	// The rapid spanning tree protocol status.
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	StpGuard *string `json:"stpGuard,omitempty"`
	// The link speed for the switch port.
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The ID of the port schedule. A value of null will clear the port schedule.
	PortScheduleId *string `json:"portScheduleId,omitempty"`
	// The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	Udld *string `json:"udld,omitempty"`
	// The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyNumber *int32 `json:"accessPolicyNumber,omitempty"`
	// Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	MacAllowList []string `json:"macAllowList,omitempty"`
	// The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowList []string `json:"stickyMacAllowList,omitempty"`
	// The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int32 `json:"stickyMacAllowListLimit,omitempty"`
	// The storm control status of the switch port.
	StormControlEnabled *bool `json:"stormControlEnabled,omitempty"`
	// The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AdaptivePolicyGroupId *string `json:"adaptivePolicyGroupId,omitempty"`
	// If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PeerSgtCapable *bool `json:"peerSgtCapable,omitempty"`
	// For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	FlexibleStackingEnabled *bool `json:"flexibleStackingEnabled,omitempty"`
	// If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	DaiTrusted *bool `json:"daiTrusted,omitempty"`
	Profile *GetDeviceSwitchPorts200ResponseInnerProfile `json:"profile,omitempty"`
}

// NewUpdateDeviceSwitchPortRequest instantiates a new UpdateDeviceSwitchPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSwitchPortRequest() *UpdateDeviceSwitchPortRequest {
	this := UpdateDeviceSwitchPortRequest{}
	return &this
}

// NewUpdateDeviceSwitchPortRequestWithDefaults instantiates a new UpdateDeviceSwitchPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSwitchPortRequestWithDefaults() *UpdateDeviceSwitchPortRequest {
	this := UpdateDeviceSwitchPortRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDeviceSwitchPortRequest) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateDeviceSwitchPortRequest) SetTags(v []string) {
	o.Tags = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceSwitchPortRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPoeEnabled returns the PoeEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetPoeEnabled() bool {
	if o == nil || IsNil(o.PoeEnabled) {
		var ret bool
		return ret
	}
	return *o.PoeEnabled
}

// GetPoeEnabledOk returns a tuple with the PoeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetPoeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PoeEnabled) {
		return nil, false
	}
	return o.PoeEnabled, true
}

// HasPoeEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasPoeEnabled() bool {
	if o != nil && !IsNil(o.PoeEnabled) {
		return true
	}

	return false
}

// SetPoeEnabled gets a reference to the given bool and assigns it to the PoeEnabled field.
func (o *UpdateDeviceSwitchPortRequest) SetPoeEnabled(v bool) {
	o.PoeEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateDeviceSwitchPortRequest) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *UpdateDeviceSwitchPortRequest) SetVlan(v int32) {
	o.Vlan = &v
}

// GetVoiceVlan returns the VoiceVlan field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetVoiceVlan() int32 {
	if o == nil || IsNil(o.VoiceVlan) {
		var ret int32
		return ret
	}
	return *o.VoiceVlan
}

// GetVoiceVlanOk returns a tuple with the VoiceVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetVoiceVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.VoiceVlan) {
		return nil, false
	}
	return o.VoiceVlan, true
}

// HasVoiceVlan returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasVoiceVlan() bool {
	if o != nil && !IsNil(o.VoiceVlan) {
		return true
	}

	return false
}

// SetVoiceVlan gets a reference to the given int32 and assigns it to the VoiceVlan field.
func (o *UpdateDeviceSwitchPortRequest) SetVoiceVlan(v int32) {
	o.VoiceVlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetAllowedVlans() string {
	if o == nil || IsNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetAllowedVlansOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedVlans) {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasAllowedVlans() bool {
	if o != nil && !IsNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *UpdateDeviceSwitchPortRequest) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetIsolationEnabled returns the IsolationEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetIsolationEnabled() bool {
	if o == nil || IsNil(o.IsolationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsolationEnabled
}

// GetIsolationEnabledOk returns a tuple with the IsolationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetIsolationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsolationEnabled) {
		return nil, false
	}
	return o.IsolationEnabled, true
}

// HasIsolationEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasIsolationEnabled() bool {
	if o != nil && !IsNil(o.IsolationEnabled) {
		return true
	}

	return false
}

// SetIsolationEnabled gets a reference to the given bool and assigns it to the IsolationEnabled field.
func (o *UpdateDeviceSwitchPortRequest) SetIsolationEnabled(v bool) {
	o.IsolationEnabled = &v
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetRstpEnabled() bool {
	if o == nil || IsNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RstpEnabled) {
		return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasRstpEnabled() bool {
	if o != nil && !IsNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *UpdateDeviceSwitchPortRequest) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpGuard returns the StpGuard field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetStpGuard() string {
	if o == nil || IsNil(o.StpGuard) {
		var ret string
		return ret
	}
	return *o.StpGuard
}

// GetStpGuardOk returns a tuple with the StpGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetStpGuardOk() (*string, bool) {
	if o == nil || IsNil(o.StpGuard) {
		return nil, false
	}
	return o.StpGuard, true
}

// HasStpGuard returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasStpGuard() bool {
	if o != nil && !IsNil(o.StpGuard) {
		return true
	}

	return false
}

// SetStpGuard gets a reference to the given string and assigns it to the StpGuard field.
func (o *UpdateDeviceSwitchPortRequest) SetStpGuard(v string) {
	o.StpGuard = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetLinkNegotiation() string {
	if o == nil || IsNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || IsNil(o.LinkNegotiation) {
		return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasLinkNegotiation() bool {
	if o != nil && !IsNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *UpdateDeviceSwitchPortRequest) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetPortScheduleId returns the PortScheduleId field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetPortScheduleId() string {
	if o == nil || IsNil(o.PortScheduleId) {
		var ret string
		return ret
	}
	return *o.PortScheduleId
}

// GetPortScheduleIdOk returns a tuple with the PortScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetPortScheduleIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortScheduleId) {
		return nil, false
	}
	return o.PortScheduleId, true
}

// HasPortScheduleId returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasPortScheduleId() bool {
	if o != nil && !IsNil(o.PortScheduleId) {
		return true
	}

	return false
}

// SetPortScheduleId gets a reference to the given string and assigns it to the PortScheduleId field.
func (o *UpdateDeviceSwitchPortRequest) SetPortScheduleId(v string) {
	o.PortScheduleId = &v
}

// GetUdld returns the Udld field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetUdld() string {
	if o == nil || IsNil(o.Udld) {
		var ret string
		return ret
	}
	return *o.Udld
}

// GetUdldOk returns a tuple with the Udld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetUdldOk() (*string, bool) {
	if o == nil || IsNil(o.Udld) {
		return nil, false
	}
	return o.Udld, true
}

// HasUdld returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasUdld() bool {
	if o != nil && !IsNil(o.Udld) {
		return true
	}

	return false
}

// SetUdld gets a reference to the given string and assigns it to the Udld field.
func (o *UpdateDeviceSwitchPortRequest) SetUdld(v string) {
	o.Udld = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetAccessPolicyType() string {
	if o == nil || IsNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicyType) {
		return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasAccessPolicyType() bool {
	if o != nil && !IsNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *UpdateDeviceSwitchPortRequest) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetAccessPolicyNumber returns the AccessPolicyNumber field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetAccessPolicyNumber() int32 {
	if o == nil || IsNil(o.AccessPolicyNumber) {
		var ret int32
		return ret
	}
	return *o.AccessPolicyNumber
}

// GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetAccessPolicyNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessPolicyNumber) {
		return nil, false
	}
	return o.AccessPolicyNumber, true
}

// HasAccessPolicyNumber returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasAccessPolicyNumber() bool {
	if o != nil && !IsNil(o.AccessPolicyNumber) {
		return true
	}

	return false
}

// SetAccessPolicyNumber gets a reference to the given int32 and assigns it to the AccessPolicyNumber field.
func (o *UpdateDeviceSwitchPortRequest) SetAccessPolicyNumber(v int32) {
	o.AccessPolicyNumber = &v
}

// GetMacAllowList returns the MacAllowList field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetMacAllowList() []string {
	if o == nil || IsNil(o.MacAllowList) {
		var ret []string
		return ret
	}
	return o.MacAllowList
}

// GetMacAllowListOk returns a tuple with the MacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetMacAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.MacAllowList) {
		return nil, false
	}
	return o.MacAllowList, true
}

// HasMacAllowList returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasMacAllowList() bool {
	if o != nil && !IsNil(o.MacAllowList) {
		return true
	}

	return false
}

// SetMacAllowList gets a reference to the given []string and assigns it to the MacAllowList field.
func (o *UpdateDeviceSwitchPortRequest) SetMacAllowList(v []string) {
	o.MacAllowList = v
}

// GetStickyMacAllowList returns the StickyMacAllowList field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetStickyMacAllowList() []string {
	if o == nil || IsNil(o.StickyMacAllowList) {
		var ret []string
		return ret
	}
	return o.StickyMacAllowList
}

// GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetStickyMacAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.StickyMacAllowList) {
		return nil, false
	}
	return o.StickyMacAllowList, true
}

// HasStickyMacAllowList returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasStickyMacAllowList() bool {
	if o != nil && !IsNil(o.StickyMacAllowList) {
		return true
	}

	return false
}

// SetStickyMacAllowList gets a reference to the given []string and assigns it to the StickyMacAllowList field.
func (o *UpdateDeviceSwitchPortRequest) SetStickyMacAllowList(v []string) {
	o.StickyMacAllowList = v
}

// GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetStickyMacAllowListLimit() int32 {
	if o == nil || IsNil(o.StickyMacAllowListLimit) {
		var ret int32
		return ret
	}
	return *o.StickyMacAllowListLimit
}

// GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetStickyMacAllowListLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.StickyMacAllowListLimit) {
		return nil, false
	}
	return o.StickyMacAllowListLimit, true
}

// HasStickyMacAllowListLimit returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasStickyMacAllowListLimit() bool {
	if o != nil && !IsNil(o.StickyMacAllowListLimit) {
		return true
	}

	return false
}

// SetStickyMacAllowListLimit gets a reference to the given int32 and assigns it to the StickyMacAllowListLimit field.
func (o *UpdateDeviceSwitchPortRequest) SetStickyMacAllowListLimit(v int32) {
	o.StickyMacAllowListLimit = &v
}

// GetStormControlEnabled returns the StormControlEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetStormControlEnabled() bool {
	if o == nil || IsNil(o.StormControlEnabled) {
		var ret bool
		return ret
	}
	return *o.StormControlEnabled
}

// GetStormControlEnabledOk returns a tuple with the StormControlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetStormControlEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StormControlEnabled) {
		return nil, false
	}
	return o.StormControlEnabled, true
}

// HasStormControlEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasStormControlEnabled() bool {
	if o != nil && !IsNil(o.StormControlEnabled) {
		return true
	}

	return false
}

// SetStormControlEnabled gets a reference to the given bool and assigns it to the StormControlEnabled field.
func (o *UpdateDeviceSwitchPortRequest) SetStormControlEnabled(v bool) {
	o.StormControlEnabled = &v
}

// GetAdaptivePolicyGroupId returns the AdaptivePolicyGroupId field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetAdaptivePolicyGroupId() string {
	if o == nil || IsNil(o.AdaptivePolicyGroupId) {
		var ret string
		return ret
	}
	return *o.AdaptivePolicyGroupId
}

// GetAdaptivePolicyGroupIdOk returns a tuple with the AdaptivePolicyGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetAdaptivePolicyGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdaptivePolicyGroupId) {
		return nil, false
	}
	return o.AdaptivePolicyGroupId, true
}

// HasAdaptivePolicyGroupId returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasAdaptivePolicyGroupId() bool {
	if o != nil && !IsNil(o.AdaptivePolicyGroupId) {
		return true
	}

	return false
}

// SetAdaptivePolicyGroupId gets a reference to the given string and assigns it to the AdaptivePolicyGroupId field.
func (o *UpdateDeviceSwitchPortRequest) SetAdaptivePolicyGroupId(v string) {
	o.AdaptivePolicyGroupId = &v
}

// GetPeerSgtCapable returns the PeerSgtCapable field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetPeerSgtCapable() bool {
	if o == nil || IsNil(o.PeerSgtCapable) {
		var ret bool
		return ret
	}
	return *o.PeerSgtCapable
}

// GetPeerSgtCapableOk returns a tuple with the PeerSgtCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetPeerSgtCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.PeerSgtCapable) {
		return nil, false
	}
	return o.PeerSgtCapable, true
}

// HasPeerSgtCapable returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasPeerSgtCapable() bool {
	if o != nil && !IsNil(o.PeerSgtCapable) {
		return true
	}

	return false
}

// SetPeerSgtCapable gets a reference to the given bool and assigns it to the PeerSgtCapable field.
func (o *UpdateDeviceSwitchPortRequest) SetPeerSgtCapable(v bool) {
	o.PeerSgtCapable = &v
}

// GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetFlexibleStackingEnabled() bool {
	if o == nil || IsNil(o.FlexibleStackingEnabled) {
		var ret bool
		return ret
	}
	return *o.FlexibleStackingEnabled
}

// GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetFlexibleStackingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FlexibleStackingEnabled) {
		return nil, false
	}
	return o.FlexibleStackingEnabled, true
}

// HasFlexibleStackingEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasFlexibleStackingEnabled() bool {
	if o != nil && !IsNil(o.FlexibleStackingEnabled) {
		return true
	}

	return false
}

// SetFlexibleStackingEnabled gets a reference to the given bool and assigns it to the FlexibleStackingEnabled field.
func (o *UpdateDeviceSwitchPortRequest) SetFlexibleStackingEnabled(v bool) {
	o.FlexibleStackingEnabled = &v
}

// GetDaiTrusted returns the DaiTrusted field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetDaiTrusted() bool {
	if o == nil || IsNil(o.DaiTrusted) {
		var ret bool
		return ret
	}
	return *o.DaiTrusted
}

// GetDaiTrustedOk returns a tuple with the DaiTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetDaiTrustedOk() (*bool, bool) {
	if o == nil || IsNil(o.DaiTrusted) {
		return nil, false
	}
	return o.DaiTrusted, true
}

// HasDaiTrusted returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasDaiTrusted() bool {
	if o != nil && !IsNil(o.DaiTrusted) {
		return true
	}

	return false
}

// SetDaiTrusted gets a reference to the given bool and assigns it to the DaiTrusted field.
func (o *UpdateDeviceSwitchPortRequest) SetDaiTrusted(v bool) {
	o.DaiTrusted = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchPortRequest) GetProfile() GetDeviceSwitchPorts200ResponseInnerProfile {
	if o == nil || IsNil(o.Profile) {
		var ret GetDeviceSwitchPorts200ResponseInnerProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchPortRequest) GetProfileOk() (*GetDeviceSwitchPorts200ResponseInnerProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchPortRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given GetDeviceSwitchPorts200ResponseInnerProfile and assigns it to the Profile field.
func (o *UpdateDeviceSwitchPortRequest) SetProfile(v GetDeviceSwitchPorts200ResponseInnerProfile) {
	o.Profile = &v
}

func (o UpdateDeviceSwitchPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceSwitchPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.PoeEnabled) {
		toSerialize["poeEnabled"] = o.PoeEnabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.VoiceVlan) {
		toSerialize["voiceVlan"] = o.VoiceVlan
	}
	if !IsNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	if !IsNil(o.IsolationEnabled) {
		toSerialize["isolationEnabled"] = o.IsolationEnabled
	}
	if !IsNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !IsNil(o.StpGuard) {
		toSerialize["stpGuard"] = o.StpGuard
	}
	if !IsNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	if !IsNil(o.PortScheduleId) {
		toSerialize["portScheduleId"] = o.PortScheduleId
	}
	if !IsNil(o.Udld) {
		toSerialize["udld"] = o.Udld
	}
	if !IsNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !IsNil(o.AccessPolicyNumber) {
		toSerialize["accessPolicyNumber"] = o.AccessPolicyNumber
	}
	if !IsNil(o.MacAllowList) {
		toSerialize["macAllowList"] = o.MacAllowList
	}
	if !IsNil(o.StickyMacAllowList) {
		toSerialize["stickyMacAllowList"] = o.StickyMacAllowList
	}
	if !IsNil(o.StickyMacAllowListLimit) {
		toSerialize["stickyMacAllowListLimit"] = o.StickyMacAllowListLimit
	}
	if !IsNil(o.StormControlEnabled) {
		toSerialize["stormControlEnabled"] = o.StormControlEnabled
	}
	if !IsNil(o.AdaptivePolicyGroupId) {
		toSerialize["adaptivePolicyGroupId"] = o.AdaptivePolicyGroupId
	}
	if !IsNil(o.PeerSgtCapable) {
		toSerialize["peerSgtCapable"] = o.PeerSgtCapable
	}
	if !IsNil(o.FlexibleStackingEnabled) {
		toSerialize["flexibleStackingEnabled"] = o.FlexibleStackingEnabled
	}
	if !IsNil(o.DaiTrusted) {
		toSerialize["daiTrusted"] = o.DaiTrusted
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullableUpdateDeviceSwitchPortRequest struct {
	value *UpdateDeviceSwitchPortRequest
	isSet bool
}

func (v NullableUpdateDeviceSwitchPortRequest) Get() *UpdateDeviceSwitchPortRequest {
	return v.value
}

func (v *NullableUpdateDeviceSwitchPortRequest) Set(val *UpdateDeviceSwitchPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSwitchPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSwitchPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSwitchPortRequest(val *UpdateDeviceSwitchPortRequest) *NullableUpdateDeviceSwitchPortRequest {
	return &NullableUpdateDeviceSwitchPortRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceSwitchPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSwitchPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


