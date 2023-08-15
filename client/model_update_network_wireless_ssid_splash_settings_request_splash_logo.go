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

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo{}

// UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo The logo used in the splash page.
type UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo struct {
	// The MD5 value of the logo file. Setting this to null will remove the logo from the splash page.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the logo file.
	Extension *string `json:"extension,omitempty"`
	Image *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage `json:"image,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) GetExtension() string {
	if o == nil || IsNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) GetExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) SetExtension(v string) {
	o.Extension = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) GetImage() UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage {
	if o == nil || IsNil(o.Image) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) GetImageOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage and assigns it to the Image field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) SetImage(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogoImage) {
	o.Image = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


