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

// checks if the GetOrganizationFirmwareUpgrades200ResponseInnerToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationFirmwareUpgrades200ResponseInnerToVersion{}

// GetOrganizationFirmwareUpgrades200ResponseInnerToVersion ID of the upgrade's target version
type GetOrganizationFirmwareUpgrades200ResponseInnerToVersion struct {
	// Firmware version ID
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
	// Release type of the firmware version
	ReleaseType *string `json:"releaseType,omitempty"`
	// Release date of the firmware version
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
}

// NewGetOrganizationFirmwareUpgrades200ResponseInnerToVersion instantiates a new GetOrganizationFirmwareUpgrades200ResponseInnerToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationFirmwareUpgrades200ResponseInnerToVersion() *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion {
	this := GetOrganizationFirmwareUpgrades200ResponseInnerToVersion{}
	return &this
}

// NewGetOrganizationFirmwareUpgrades200ResponseInnerToVersionWithDefaults instantiates a new GetOrganizationFirmwareUpgrades200ResponseInnerToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationFirmwareUpgrades200ResponseInnerToVersionWithDefaults() *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion {
	this := GetOrganizationFirmwareUpgrades200ResponseInnerToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetReleaseType() string {
	if o == nil || IsNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetReleaseDate() time.Time {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return toSerialize, nil
}

type NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion struct {
	value *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion
	isSet bool
}

func (v NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion) Get() *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion {
	return v.value
}

func (v *NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion) Set(val *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion(val *GetOrganizationFirmwareUpgrades200ResponseInnerToVersion) *NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion {
	return &NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion{value: val, isSet: true}
}

func (v NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationFirmwareUpgrades200ResponseInnerToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


