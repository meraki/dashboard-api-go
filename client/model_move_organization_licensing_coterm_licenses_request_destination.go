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

// MoveOrganizationLicensingCotermLicensesRequestDestination Destination data for the license move
type MoveOrganizationLicensingCotermLicensesRequestDestination struct {
	// The organization to move the license to
	OrganizationId *string `json:"organizationId,omitempty"`
	// The claim mode of the moved license
	Mode *string `json:"mode,omitempty"`
}

// NewMoveOrganizationLicensingCotermLicensesRequestDestination instantiates a new MoveOrganizationLicensingCotermLicensesRequestDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensingCotermLicensesRequestDestination() *MoveOrganizationLicensingCotermLicensesRequestDestination {
	this := MoveOrganizationLicensingCotermLicensesRequestDestination{}
	return &this
}

// NewMoveOrganizationLicensingCotermLicensesRequestDestinationWithDefaults instantiates a new MoveOrganizationLicensingCotermLicensesRequestDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensingCotermLicensesRequestDestinationWithDefaults() *MoveOrganizationLicensingCotermLicensesRequestDestination {
	this := MoveOrganizationLicensingCotermLicensesRequestDestination{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *MoveOrganizationLicensingCotermLicensesRequestDestination) SetMode(v string) {
	o.Mode = &v
}

func (o MoveOrganizationLicensingCotermLicensesRequestDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableMoveOrganizationLicensingCotermLicensesRequestDestination struct {
	value *MoveOrganizationLicensingCotermLicensesRequestDestination
	isSet bool
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestDestination) Get() *MoveOrganizationLicensingCotermLicensesRequestDestination {
	return v.value
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestDestination) Set(val *MoveOrganizationLicensingCotermLicensesRequestDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensingCotermLicensesRequestDestination(val *MoveOrganizationLicensingCotermLicensesRequestDestination) *NullableMoveOrganizationLicensingCotermLicensesRequestDestination {
	return &NullableMoveOrganizationLicensingCotermLicensesRequestDestination{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


