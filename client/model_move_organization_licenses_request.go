/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MoveOrganizationLicensesRequest struct for MoveOrganizationLicensesRequest
type MoveOrganizationLicensesRequest struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId string `json:"destOrganizationId"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds"`
}

// NewMoveOrganizationLicensesRequest instantiates a new MoveOrganizationLicensesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensesRequest(destOrganizationId string, licenseIds []string) *MoveOrganizationLicensesRequest {
	this := MoveOrganizationLicensesRequest{}
	this.DestOrganizationId = destOrganizationId
	this.LicenseIds = licenseIds
	return &this
}

// NewMoveOrganizationLicensesRequestWithDefaults instantiates a new MoveOrganizationLicensesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensesRequestWithDefaults() *MoveOrganizationLicensesRequest {
	this := MoveOrganizationLicensesRequest{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value
func (o *MoveOrganizationLicensesRequest) GetDestOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesRequest) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestOrganizationId, true
}

// SetDestOrganizationId sets field value
func (o *MoveOrganizationLicensesRequest) SetDestOrganizationId(v string) {
	o.DestOrganizationId = v
}

// GetLicenseIds returns the LicenseIds field value
func (o *MoveOrganizationLicensesRequest) GetLicenseIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensesRequest) GetLicenseIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LicenseIds, true
}

// SetLicenseIds sets field value
func (o *MoveOrganizationLicensesRequest) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o MoveOrganizationLicensesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if true {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return json.Marshal(toSerialize)
}

type NullableMoveOrganizationLicensesRequest struct {
	value *MoveOrganizationLicensesRequest
	isSet bool
}

func (v NullableMoveOrganizationLicensesRequest) Get() *MoveOrganizationLicensesRequest {
	return v.value
}

func (v *NullableMoveOrganizationLicensesRequest) Set(val *MoveOrganizationLicensesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensesRequest(val *MoveOrganizationLicensesRequest) *NullableMoveOrganizationLicensesRequest {
	return &NullableMoveOrganizationLicensesRequest{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


