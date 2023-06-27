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

// GetNetworkSmDeviceSecurityCenters200ResponseInner struct for GetNetworkSmDeviceSecurityCenters200ResponseInner
type GetNetworkSmDeviceSecurityCenters200ResponseInner struct {
	// Boolean indicating if the device is rooted.
	IsRooted *bool `json:"isRooted,omitempty"`
	// Boolean indicating if the device has Antivirus.
	HasAntiVirus *bool `json:"hasAntiVirus,omitempty"`
	// The name of the Antivirus.
	AntiVirusName *string `json:"antiVirusName,omitempty"`
	// Boolean indicating if the device has a Firewall enabled.
	IsFireWallEnabled *bool `json:"isFireWallEnabled,omitempty"`
	// Boolean indicating if the device has a Firewall installed.
	HasFireWallInstalled *bool `json:"hasFireWallInstalled,omitempty"`
	// The name of the Firewall.
	FireWallName *string `json:"fireWallName,omitempty"`
	// Boolean indicating if the device has disk encryption.
	IsDiskEncrypted *bool `json:"isDiskEncrypted,omitempty"`
	// Boolean indicating if the device has auto login disabled.
	IsAutoLoginDisabled *bool `json:"isAutoLoginDisabled,omitempty"`
	// The Meraki identifier for the security center record.
	Id *string `json:"id,omitempty"`
	// A comma seperated list of procs running on the device.
	RunningProcs *string `json:"runningProcs,omitempty"`
}

// NewGetNetworkSmDeviceSecurityCenters200ResponseInner instantiates a new GetNetworkSmDeviceSecurityCenters200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDeviceSecurityCenters200ResponseInner() *GetNetworkSmDeviceSecurityCenters200ResponseInner {
	this := GetNetworkSmDeviceSecurityCenters200ResponseInner{}
	return &this
}

// NewGetNetworkSmDeviceSecurityCenters200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceSecurityCenters200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDeviceSecurityCenters200ResponseInnerWithDefaults() *GetNetworkSmDeviceSecurityCenters200ResponseInner {
	this := GetNetworkSmDeviceSecurityCenters200ResponseInner{}
	return &this
}

// GetIsRooted returns the IsRooted field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsRooted() bool {
	if o == nil || isNil(o.IsRooted) {
		var ret bool
		return ret
	}
	return *o.IsRooted
}

// GetIsRootedOk returns a tuple with the IsRooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsRootedOk() (*bool, bool) {
	if o == nil || isNil(o.IsRooted) {
    return nil, false
	}
	return o.IsRooted, true
}

// HasIsRooted returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsRooted() bool {
	if o != nil && !isNil(o.IsRooted) {
		return true
	}

	return false
}

// SetIsRooted gets a reference to the given bool and assigns it to the IsRooted field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsRooted(v bool) {
	o.IsRooted = &v
}

// GetHasAntiVirus returns the HasAntiVirus field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasAntiVirus() bool {
	if o == nil || isNil(o.HasAntiVirus) {
		var ret bool
		return ret
	}
	return *o.HasAntiVirus
}

// GetHasAntiVirusOk returns a tuple with the HasAntiVirus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasAntiVirusOk() (*bool, bool) {
	if o == nil || isNil(o.HasAntiVirus) {
    return nil, false
	}
	return o.HasAntiVirus, true
}

// HasHasAntiVirus returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasHasAntiVirus() bool {
	if o != nil && !isNil(o.HasAntiVirus) {
		return true
	}

	return false
}

// SetHasAntiVirus gets a reference to the given bool and assigns it to the HasAntiVirus field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetHasAntiVirus(v bool) {
	o.HasAntiVirus = &v
}

// GetAntiVirusName returns the AntiVirusName field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetAntiVirusName() string {
	if o == nil || isNil(o.AntiVirusName) {
		var ret string
		return ret
	}
	return *o.AntiVirusName
}

// GetAntiVirusNameOk returns a tuple with the AntiVirusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetAntiVirusNameOk() (*string, bool) {
	if o == nil || isNil(o.AntiVirusName) {
    return nil, false
	}
	return o.AntiVirusName, true
}

// HasAntiVirusName returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasAntiVirusName() bool {
	if o != nil && !isNil(o.AntiVirusName) {
		return true
	}

	return false
}

// SetAntiVirusName gets a reference to the given string and assigns it to the AntiVirusName field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetAntiVirusName(v string) {
	o.AntiVirusName = &v
}

// GetIsFireWallEnabled returns the IsFireWallEnabled field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsFireWallEnabled() bool {
	if o == nil || isNil(o.IsFireWallEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFireWallEnabled
}

// GetIsFireWallEnabledOk returns a tuple with the IsFireWallEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsFireWallEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsFireWallEnabled) {
    return nil, false
	}
	return o.IsFireWallEnabled, true
}

// HasIsFireWallEnabled returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsFireWallEnabled() bool {
	if o != nil && !isNil(o.IsFireWallEnabled) {
		return true
	}

	return false
}

// SetIsFireWallEnabled gets a reference to the given bool and assigns it to the IsFireWallEnabled field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsFireWallEnabled(v bool) {
	o.IsFireWallEnabled = &v
}

// GetHasFireWallInstalled returns the HasFireWallInstalled field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasFireWallInstalled() bool {
	if o == nil || isNil(o.HasFireWallInstalled) {
		var ret bool
		return ret
	}
	return *o.HasFireWallInstalled
}

// GetHasFireWallInstalledOk returns a tuple with the HasFireWallInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasFireWallInstalledOk() (*bool, bool) {
	if o == nil || isNil(o.HasFireWallInstalled) {
    return nil, false
	}
	return o.HasFireWallInstalled, true
}

// HasHasFireWallInstalled returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasHasFireWallInstalled() bool {
	if o != nil && !isNil(o.HasFireWallInstalled) {
		return true
	}

	return false
}

// SetHasFireWallInstalled gets a reference to the given bool and assigns it to the HasFireWallInstalled field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetHasFireWallInstalled(v bool) {
	o.HasFireWallInstalled = &v
}

// GetFireWallName returns the FireWallName field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetFireWallName() string {
	if o == nil || isNil(o.FireWallName) {
		var ret string
		return ret
	}
	return *o.FireWallName
}

// GetFireWallNameOk returns a tuple with the FireWallName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetFireWallNameOk() (*string, bool) {
	if o == nil || isNil(o.FireWallName) {
    return nil, false
	}
	return o.FireWallName, true
}

// HasFireWallName returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasFireWallName() bool {
	if o != nil && !isNil(o.FireWallName) {
		return true
	}

	return false
}

// SetFireWallName gets a reference to the given string and assigns it to the FireWallName field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetFireWallName(v string) {
	o.FireWallName = &v
}

// GetIsDiskEncrypted returns the IsDiskEncrypted field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsDiskEncrypted() bool {
	if o == nil || isNil(o.IsDiskEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsDiskEncrypted
}

// GetIsDiskEncryptedOk returns a tuple with the IsDiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsDiskEncryptedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDiskEncrypted) {
    return nil, false
	}
	return o.IsDiskEncrypted, true
}

// HasIsDiskEncrypted returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsDiskEncrypted() bool {
	if o != nil && !isNil(o.IsDiskEncrypted) {
		return true
	}

	return false
}

// SetIsDiskEncrypted gets a reference to the given bool and assigns it to the IsDiskEncrypted field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsDiskEncrypted(v bool) {
	o.IsDiskEncrypted = &v
}

// GetIsAutoLoginDisabled returns the IsAutoLoginDisabled field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsAutoLoginDisabled() bool {
	if o == nil || isNil(o.IsAutoLoginDisabled) {
		var ret bool
		return ret
	}
	return *o.IsAutoLoginDisabled
}

// GetIsAutoLoginDisabledOk returns a tuple with the IsAutoLoginDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsAutoLoginDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsAutoLoginDisabled) {
    return nil, false
	}
	return o.IsAutoLoginDisabled, true
}

// HasIsAutoLoginDisabled returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsAutoLoginDisabled() bool {
	if o != nil && !isNil(o.IsAutoLoginDisabled) {
		return true
	}

	return false
}

// SetIsAutoLoginDisabled gets a reference to the given bool and assigns it to the IsAutoLoginDisabled field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsAutoLoginDisabled(v bool) {
	o.IsAutoLoginDisabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetRunningProcs returns the RunningProcs field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetRunningProcs() string {
	if o == nil || isNil(o.RunningProcs) {
		var ret string
		return ret
	}
	return *o.RunningProcs
}

// GetRunningProcsOk returns a tuple with the RunningProcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetRunningProcsOk() (*string, bool) {
	if o == nil || isNil(o.RunningProcs) {
    return nil, false
	}
	return o.RunningProcs, true
}

// HasRunningProcs returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasRunningProcs() bool {
	if o != nil && !isNil(o.RunningProcs) {
		return true
	}

	return false
}

// SetRunningProcs gets a reference to the given string and assigns it to the RunningProcs field.
func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetRunningProcs(v string) {
	o.RunningProcs = &v
}

func (o GetNetworkSmDeviceSecurityCenters200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsRooted) {
		toSerialize["isRooted"] = o.IsRooted
	}
	if !isNil(o.HasAntiVirus) {
		toSerialize["hasAntiVirus"] = o.HasAntiVirus
	}
	if !isNil(o.AntiVirusName) {
		toSerialize["antiVirusName"] = o.AntiVirusName
	}
	if !isNil(o.IsFireWallEnabled) {
		toSerialize["isFireWallEnabled"] = o.IsFireWallEnabled
	}
	if !isNil(o.HasFireWallInstalled) {
		toSerialize["hasFireWallInstalled"] = o.HasFireWallInstalled
	}
	if !isNil(o.FireWallName) {
		toSerialize["fireWallName"] = o.FireWallName
	}
	if !isNil(o.IsDiskEncrypted) {
		toSerialize["isDiskEncrypted"] = o.IsDiskEncrypted
	}
	if !isNil(o.IsAutoLoginDisabled) {
		toSerialize["isAutoLoginDisabled"] = o.IsAutoLoginDisabled
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.RunningProcs) {
		toSerialize["runningProcs"] = o.RunningProcs
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSmDeviceSecurityCenters200ResponseInner struct {
	value *GetNetworkSmDeviceSecurityCenters200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDeviceSecurityCenters200ResponseInner) Get() *GetNetworkSmDeviceSecurityCenters200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDeviceSecurityCenters200ResponseInner) Set(val *GetNetworkSmDeviceSecurityCenters200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDeviceSecurityCenters200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDeviceSecurityCenters200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDeviceSecurityCenters200ResponseInner(val *GetNetworkSmDeviceSecurityCenters200ResponseInner) *NullableGetNetworkSmDeviceSecurityCenters200ResponseInner {
	return &NullableGetNetworkSmDeviceSecurityCenters200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDeviceSecurityCenters200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDeviceSecurityCenters200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


