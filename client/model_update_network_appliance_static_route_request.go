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

// UpdateNetworkApplianceStaticRouteRequest struct for UpdateNetworkApplianceStaticRouteRequest
type UpdateNetworkApplianceStaticRouteRequest struct {
	// The name of the static route
	Name *string `json:"name,omitempty"`
	// The subnet of the static route
	Subnet *string `json:"subnet,omitempty"`
	// The gateway IP (next hop) of the static route
	GatewayIp *string `json:"gatewayIp,omitempty"`
	// The gateway IP (next hop) VLAN ID of the static route
	GatewayVlanId *string `json:"gatewayVlanId,omitempty"`
	// The enabled state of the static route
	Enabled *bool `json:"enabled,omitempty"`
	// The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain \"ip\" and \"name\" string fields. See the sample request/response for more details.
	FixedIpAssignments map[string]interface{} `json:"fixedIpAssignments,omitempty"`
	// The DHCP reserved IP ranges on the static route
	ReservedIpRanges []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner `json:"reservedIpRanges,omitempty"`
}

// NewUpdateNetworkApplianceStaticRouteRequest instantiates a new UpdateNetworkApplianceStaticRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceStaticRouteRequest() *UpdateNetworkApplianceStaticRouteRequest {
	this := UpdateNetworkApplianceStaticRouteRequest{}
	return &this
}

// NewUpdateNetworkApplianceStaticRouteRequestWithDefaults instantiates a new UpdateNetworkApplianceStaticRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceStaticRouteRequestWithDefaults() *UpdateNetworkApplianceStaticRouteRequest {
	this := UpdateNetworkApplianceStaticRouteRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkApplianceStaticRouteRequest) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *UpdateNetworkApplianceStaticRouteRequest) SetSubnet(v string) {
	o.Subnet = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayIp() string {
	if o == nil || isNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayIpOk() (*string, bool) {
	if o == nil || isNil(o.GatewayIp) {
    return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) HasGatewayIp() bool {
	if o != nil && !isNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *UpdateNetworkApplianceStaticRouteRequest) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetGatewayVlanId returns the GatewayVlanId field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayVlanId() string {
	if o == nil || isNil(o.GatewayVlanId) {
		var ret string
		return ret
	}
	return *o.GatewayVlanId
}

// GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayVlanIdOk() (*string, bool) {
	if o == nil || isNil(o.GatewayVlanId) {
    return nil, false
	}
	return o.GatewayVlanId, true
}

// HasGatewayVlanId returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) HasGatewayVlanId() bool {
	if o != nil && !isNil(o.GatewayVlanId) {
		return true
	}

	return false
}

// SetGatewayVlanId gets a reference to the given string and assigns it to the GatewayVlanId field.
func (o *UpdateNetworkApplianceStaticRouteRequest) SetGatewayVlanId(v string) {
	o.GatewayVlanId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkApplianceStaticRouteRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetFixedIpAssignments() map[string]interface{} {
	if o == nil || isNil(o.FixedIpAssignments) {
		var ret map[string]interface{}
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetFixedIpAssignmentsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.FixedIpAssignments) {
    return map[string]interface{}{}, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) HasFixedIpAssignments() bool {
	if o != nil && !isNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given map[string]interface{} and assigns it to the FixedIpAssignments field.
func (o *UpdateNetworkApplianceStaticRouteRequest) SetFixedIpAssignments(v map[string]interface{}) {
	o.FixedIpAssignments = v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetReservedIpRanges() []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner {
	if o == nil || isNil(o.ReservedIpRanges) {
		var ret []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) GetReservedIpRangesOk() ([]UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner, bool) {
	if o == nil || isNil(o.ReservedIpRanges) {
    return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceStaticRouteRequest) HasReservedIpRanges() bool {
	if o != nil && !isNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner and assigns it to the ReservedIpRanges field.
func (o *UpdateNetworkApplianceStaticRouteRequest) SetReservedIpRanges(v []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner) {
	o.ReservedIpRanges = v
}

func (o UpdateNetworkApplianceStaticRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.GatewayIp) {
		toSerialize["gatewayIp"] = o.GatewayIp
	}
	if !isNil(o.GatewayVlanId) {
		toSerialize["gatewayVlanId"] = o.GatewayVlanId
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	if !isNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceStaticRouteRequest struct {
	value *UpdateNetworkApplianceStaticRouteRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceStaticRouteRequest) Get() *UpdateNetworkApplianceStaticRouteRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceStaticRouteRequest) Set(val *UpdateNetworkApplianceStaticRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceStaticRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceStaticRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceStaticRouteRequest(val *UpdateNetworkApplianceStaticRouteRequest) *NullableUpdateNetworkApplianceStaticRouteRequest {
	return &NullableUpdateNetworkApplianceStaticRouteRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceStaticRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceStaticRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


