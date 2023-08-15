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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct for CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	// A set of devices to import (or update)
	Devices []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner `json:"devices"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(devices []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{}
	this.Devices = devices
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{}
	return &this
}

// GetDevices returns the Devices field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) GetDevices() []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner {
	if o == nil {
		var ret []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) GetDevicesOk() ([]CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) SetDevices(v []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) {
	o.Devices = v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


