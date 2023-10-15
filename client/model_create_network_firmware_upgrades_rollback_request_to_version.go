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

// checks if the CreateNetworkFirmwareUpgradesRollbackRequestToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesRollbackRequestToVersion{}

// CreateNetworkFirmwareUpgradesRollbackRequestToVersion Version to downgrade to (if the network has firmware flexibility)
type CreateNetworkFirmwareUpgradesRollbackRequestToVersion struct {
	// The version ID
	Id *string `json:"id,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesRollbackRequestToVersion instantiates a new CreateNetworkFirmwareUpgradesRollbackRequestToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesRollbackRequestToVersion() *CreateNetworkFirmwareUpgradesRollbackRequestToVersion {
	this := CreateNetworkFirmwareUpgradesRollbackRequestToVersion{}
	return &this
}

// NewCreateNetworkFirmwareUpgradesRollbackRequestToVersionWithDefaults instantiates a new CreateNetworkFirmwareUpgradesRollbackRequestToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesRollbackRequestToVersionWithDefaults() *CreateNetworkFirmwareUpgradesRollbackRequestToVersion {
	this := CreateNetworkFirmwareUpgradesRollbackRequestToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollbackRequestToVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequestToVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequestToVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateNetworkFirmwareUpgradesRollbackRequestToVersion) SetId(v string) {
	o.Id = &v
}

func (o CreateNetworkFirmwareUpgradesRollbackRequestToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesRollbackRequestToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion struct {
	value *CreateNetworkFirmwareUpgradesRollbackRequestToVersion
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion) Get() *CreateNetworkFirmwareUpgradesRollbackRequestToVersion {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion) Set(val *CreateNetworkFirmwareUpgradesRollbackRequestToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion(val *CreateNetworkFirmwareUpgradesRollbackRequestToVersion) *NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion {
	return &NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequestToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


