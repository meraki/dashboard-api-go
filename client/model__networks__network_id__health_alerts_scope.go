/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdHealthAlertsScope The scope of the alert
type NetworksNetworkIdHealthAlertsScope struct {
	// Devices related to the alert
	Devices []NetworksNetworkIdHealthAlertsScopeDevices `json:"devices,omitempty"`
	// Applications related to the alert
	Applications []NetworksNetworkIdHealthAlertsScopeApplications `json:"applications,omitempty"`
	// Peers related to the alert
	Peers []NetworksNetworkIdHealthAlertsScopePeers `json:"peers,omitempty"`
}

// NewNetworksNetworkIdHealthAlertsScope instantiates a new NetworksNetworkIdHealthAlertsScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdHealthAlertsScope() *NetworksNetworkIdHealthAlertsScope {
	this := NetworksNetworkIdHealthAlertsScope{}
	return &this
}

// NewNetworksNetworkIdHealthAlertsScopeWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdHealthAlertsScopeWithDefaults() *NetworksNetworkIdHealthAlertsScope {
	this := NetworksNetworkIdHealthAlertsScope{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScope) GetDevices() []NetworksNetworkIdHealthAlertsScopeDevices {
	if o == nil || isNil(o.Devices) {
		var ret []NetworksNetworkIdHealthAlertsScopeDevices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScope) GetDevicesOk() ([]NetworksNetworkIdHealthAlertsScopeDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScope) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []NetworksNetworkIdHealthAlertsScopeDevices and assigns it to the Devices field.
func (o *NetworksNetworkIdHealthAlertsScope) SetDevices(v []NetworksNetworkIdHealthAlertsScopeDevices) {
	o.Devices = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScope) GetApplications() []NetworksNetworkIdHealthAlertsScopeApplications {
	if o == nil || isNil(o.Applications) {
		var ret []NetworksNetworkIdHealthAlertsScopeApplications
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScope) GetApplicationsOk() ([]NetworksNetworkIdHealthAlertsScopeApplications, bool) {
	if o == nil || isNil(o.Applications) {
    return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScope) HasApplications() bool {
	if o != nil && !isNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []NetworksNetworkIdHealthAlertsScopeApplications and assigns it to the Applications field.
func (o *NetworksNetworkIdHealthAlertsScope) SetApplications(v []NetworksNetworkIdHealthAlertsScopeApplications) {
	o.Applications = v
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScope) GetPeers() []NetworksNetworkIdHealthAlertsScopePeers {
	if o == nil || isNil(o.Peers) {
		var ret []NetworksNetworkIdHealthAlertsScopePeers
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScope) GetPeersOk() ([]NetworksNetworkIdHealthAlertsScopePeers, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScope) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []NetworksNetworkIdHealthAlertsScopePeers and assigns it to the Peers field.
func (o *NetworksNetworkIdHealthAlertsScope) SetPeers(v []NetworksNetworkIdHealthAlertsScopePeers) {
	o.Peers = v
}

func (o NetworksNetworkIdHealthAlertsScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdHealthAlertsScope struct {
	value *NetworksNetworkIdHealthAlertsScope
	isSet bool
}

func (v NullableNetworksNetworkIdHealthAlertsScope) Get() *NetworksNetworkIdHealthAlertsScope {
	return v.value
}

func (v *NullableNetworksNetworkIdHealthAlertsScope) Set(val *NetworksNetworkIdHealthAlertsScope) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdHealthAlertsScope) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdHealthAlertsScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdHealthAlertsScope(val *NetworksNetworkIdHealthAlertsScope) *NullableNetworksNetworkIdHealthAlertsScope {
	return &NullableNetworksNetworkIdHealthAlertsScope{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdHealthAlertsScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdHealthAlertsScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

