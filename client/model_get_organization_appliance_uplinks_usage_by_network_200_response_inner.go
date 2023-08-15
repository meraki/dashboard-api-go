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

// checks if the GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner{}

// GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner struct for GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner
type GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// Uplink usage
	ByUplink []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner `json:"byUplink,omitempty"`
}

// NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner instantiates a new GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner() *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner {
	this := GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner{}
	return &this
}

// NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerWithDefaults instantiates a new GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerWithDefaults() *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner {
	this := GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetByUplink returns the ByUplink field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) GetByUplink() []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner {
	if o == nil || IsNil(o.ByUplink) {
		var ret []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner
		return ret
	}
	return o.ByUplink
}

// GetByUplinkOk returns a tuple with the ByUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) GetByUplinkOk() ([]GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner, bool) {
	if o == nil || IsNil(o.ByUplink) {
		return nil, false
	}
	return o.ByUplink, true
}

// HasByUplink returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) HasByUplink() bool {
	if o != nil && !IsNil(o.ByUplink) {
		return true
	}

	return false
}

// SetByUplink gets a reference to the given []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner and assigns it to the ByUplink field.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) SetByUplink(v []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) {
	o.ByUplink = v
}

func (o GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ByUplink) {
		toSerialize["byUplink"] = o.ByUplink
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner struct {
	value *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) Get() *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) Set(val *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner(val *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner {
	return &NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


