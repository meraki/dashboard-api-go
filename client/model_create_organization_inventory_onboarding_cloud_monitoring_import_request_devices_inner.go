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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner{}

// CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner struct for CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner
type CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner struct {
	// Import ID from the Import operation
	DeviceId string `json:"deviceId"`
	// Device UDI certificate
	Udi string `json:"udi"`
	// Network Id
	NetworkId string `json:"networkId"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner(deviceId string, udi string, networkId string) *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner{}
	this.DeviceId = deviceId
	this.Udi = udi
	this.NetworkId = networkId
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInnerWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInnerWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetUdi returns the Udi field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) GetUdi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Udi
}

// GetUdiOk returns a tuple with the Udi field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) GetUdiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Udi, true
}

// SetUdi sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) SetUdi(v string) {
	o.Udi = v
}

// GetNetworkId returns the NetworkId field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) SetNetworkId(v string) {
	o.NetworkId = v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceId"] = o.DeviceId
	toSerialize["udi"] = o.Udi
	toSerialize["networkId"] = o.NetworkId
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner(val *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


