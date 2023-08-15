/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationCellularGatewayUplinkStatuses200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationCellularGatewayUplinkStatuses200ResponseInner{}

// GetOrganizationCellularGatewayUplinkStatuses200ResponseInner struct for GetOrganizationCellularGatewayUplinkStatuses200ResponseInner
type GetOrganizationCellularGatewayUplinkStatuses200ResponseInner struct {
	// Network Id
	NetworkId *string `json:"networkId,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Device model
	Model *string `json:"model,omitempty"`
	// Last reported time for the device
	LastReportedAt *time.Time `json:"lastReportedAt,omitempty"`
	// Uplinks info
	Uplinks []GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner `json:"uplinks,omitempty"`
}

// NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInner instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInner() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner {
	this := GetOrganizationCellularGatewayUplinkStatuses200ResponseInner{}
	return &this
}

// NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerWithDefaults() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner {
	this := GetOrganizationCellularGatewayUplinkStatuses200ResponseInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetLastReportedAt() time.Time {
	if o == nil || IsNil(o.LastReportedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetLastReportedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastReportedAt) {
		return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasLastReportedAt() bool {
	if o != nil && !IsNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given time.Time and assigns it to the LastReportedAt field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetLastReportedAt(v time.Time) {
	o.LastReportedAt = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetUplinks() []GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner {
	if o == nil || IsNil(o.Uplinks) {
		var ret []GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetUplinksOk() ([]GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner, bool) {
	if o == nil || IsNil(o.Uplinks) {
		return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasUplinks() bool {
	if o != nil && !IsNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner and assigns it to the Uplinks field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetUplinks(v []GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) {
	o.Uplinks = v
}

func (o GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.LastReportedAt) {
		toSerialize["lastReportedAt"] = o.LastReportedAt
	}
	if !IsNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	return toSerialize, nil
}

type NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner struct {
	value *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner) Get() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner) Set(val *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner(val *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner {
	return &NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


