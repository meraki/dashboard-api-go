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

// checks if the UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest{}

// UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest struct for UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest
type UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest struct {
	// Toggle for enabling or disabling active-active AutoVPN
	ActiveActiveAutoVpnEnabled *bool `json:"activeActiveAutoVpnEnabled,omitempty"`
	// The default uplink. Must be one of: 'wan1' or 'wan2'
	DefaultUplink *string `json:"defaultUplink,omitempty"`
	// Toggle for enabling or disabling load balancing
	LoadBalancingEnabled *bool `json:"loadBalancingEnabled,omitempty"`
	FailoverAndFailback *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback `json:"failoverAndFailback,omitempty"`
	// Array of uplink preference rules for WAN traffic
	WanTrafficUplinkPreferences []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner `json:"wanTrafficUplinkPreferences,omitempty"`
	// Array of uplink preference rules for VPN traffic
	VpnTrafficUplinkPreferences []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner `json:"vpnTrafficUplinkPreferences,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest{}
	return &this
}

// GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetActiveActiveAutoVpnEnabled() bool {
	if o == nil || IsNil(o.ActiveActiveAutoVpnEnabled) {
		var ret bool
		return ret
	}
	return *o.ActiveActiveAutoVpnEnabled
}

// GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetActiveActiveAutoVpnEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveActiveAutoVpnEnabled) {
		return nil, false
	}
	return o.ActiveActiveAutoVpnEnabled, true
}

// HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasActiveActiveAutoVpnEnabled() bool {
	if o != nil && !IsNil(o.ActiveActiveAutoVpnEnabled) {
		return true
	}

	return false
}

// SetActiveActiveAutoVpnEnabled gets a reference to the given bool and assigns it to the ActiveActiveAutoVpnEnabled field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetActiveActiveAutoVpnEnabled(v bool) {
	o.ActiveActiveAutoVpnEnabled = &v
}

// GetDefaultUplink returns the DefaultUplink field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetDefaultUplink() string {
	if o == nil || IsNil(o.DefaultUplink) {
		var ret string
		return ret
	}
	return *o.DefaultUplink
}

// GetDefaultUplinkOk returns a tuple with the DefaultUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetDefaultUplinkOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultUplink) {
		return nil, false
	}
	return o.DefaultUplink, true
}

// HasDefaultUplink returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasDefaultUplink() bool {
	if o != nil && !IsNil(o.DefaultUplink) {
		return true
	}

	return false
}

// SetDefaultUplink gets a reference to the given string and assigns it to the DefaultUplink field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetDefaultUplink(v string) {
	o.DefaultUplink = &v
}

// GetLoadBalancingEnabled returns the LoadBalancingEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetLoadBalancingEnabled() bool {
	if o == nil || IsNil(o.LoadBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.LoadBalancingEnabled
}

// GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetLoadBalancingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LoadBalancingEnabled) {
		return nil, false
	}
	return o.LoadBalancingEnabled, true
}

// HasLoadBalancingEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasLoadBalancingEnabled() bool {
	if o != nil && !IsNil(o.LoadBalancingEnabled) {
		return true
	}

	return false
}

// SetLoadBalancingEnabled gets a reference to the given bool and assigns it to the LoadBalancingEnabled field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetLoadBalancingEnabled(v bool) {
	o.LoadBalancingEnabled = &v
}

// GetFailoverAndFailback returns the FailoverAndFailback field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetFailoverAndFailback() UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback {
	if o == nil || IsNil(o.FailoverAndFailback) {
		var ret UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback
		return ret
	}
	return *o.FailoverAndFailback
}

// GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetFailoverAndFailbackOk() (*UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback, bool) {
	if o == nil || IsNil(o.FailoverAndFailback) {
		return nil, false
	}
	return o.FailoverAndFailback, true
}

// HasFailoverAndFailback returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasFailoverAndFailback() bool {
	if o != nil && !IsNil(o.FailoverAndFailback) {
		return true
	}

	return false
}

// SetFailoverAndFailback gets a reference to the given UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback and assigns it to the FailoverAndFailback field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetFailoverAndFailback(v UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) {
	o.FailoverAndFailback = &v
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetWanTrafficUplinkPreferences() []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner {
	if o == nil || IsNil(o.WanTrafficUplinkPreferences) {
		var ret []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetWanTrafficUplinkPreferencesOk() ([]UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner, bool) {
	if o == nil || IsNil(o.WanTrafficUplinkPreferences) {
		return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !IsNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner and assigns it to the WanTrafficUplinkPreferences field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetWanTrafficUplinkPreferences(v []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner) {
	o.WanTrafficUplinkPreferences = v
}

// GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetVpnTrafficUplinkPreferences() []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner {
	if o == nil || IsNil(o.VpnTrafficUplinkPreferences) {
		var ret []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner
		return ret
	}
	return o.VpnTrafficUplinkPreferences
}

// GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetVpnTrafficUplinkPreferencesOk() ([]UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner, bool) {
	if o == nil || IsNil(o.VpnTrafficUplinkPreferences) {
		return nil, false
	}
	return o.VpnTrafficUplinkPreferences, true
}

// HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasVpnTrafficUplinkPreferences() bool {
	if o != nil && !IsNil(o.VpnTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetVpnTrafficUplinkPreferences gets a reference to the given []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner and assigns it to the VpnTrafficUplinkPreferences field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetVpnTrafficUplinkPreferences(v []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner) {
	o.VpnTrafficUplinkPreferences = v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveActiveAutoVpnEnabled) {
		toSerialize["activeActiveAutoVpnEnabled"] = o.ActiveActiveAutoVpnEnabled
	}
	if !IsNil(o.DefaultUplink) {
		toSerialize["defaultUplink"] = o.DefaultUplink
	}
	if !IsNil(o.LoadBalancingEnabled) {
		toSerialize["loadBalancingEnabled"] = o.LoadBalancingEnabled
	}
	if !IsNil(o.FailoverAndFailback) {
		toSerialize["failoverAndFailback"] = o.FailoverAndFailback
	}
	if !IsNil(o.WanTrafficUplinkPreferences) {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	if !IsNil(o.VpnTrafficUplinkPreferences) {
		toSerialize["vpnTrafficUplinkPreferences"] = o.VpnTrafficUplinkPreferences
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) Get() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) Set(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


