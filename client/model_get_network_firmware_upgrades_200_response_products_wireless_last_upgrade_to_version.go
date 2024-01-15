/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion{}

// GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion Details of the version the device upgraded to
type GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion struct {
	// Firmware version identifier
	Id *string `json:"id,omitempty"`
	// Name of the firmware version
	Firmware *string `json:"firmware,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
	// Release type of the firmware version
	ReleaseType *string `json:"releaseType,omitempty"`
	// Release date of the firmware version
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion{}
	return &this
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersionWithDefaults instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersionWithDefaults() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetFirmware() string {
	if o == nil || IsNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetFirmwareOk() (*string, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) SetFirmware(v string) {
	o.Firmware = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetReleaseType() string {
	if o == nil || IsNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetReleaseDate() time.Time {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
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

type NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion struct {
	value *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) Get() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) Set(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion {
	return &NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


