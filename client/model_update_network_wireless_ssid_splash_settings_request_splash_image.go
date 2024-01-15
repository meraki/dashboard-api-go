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

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage{}

// UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage The image used in the splash page.
type UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage struct {
	// The MD5 value of the image file. Setting this to null will remove the image from the splash page.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the image file.
	Extension *string `json:"extension,omitempty"`
	Image *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImageImage `json:"image,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashImageWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashImageWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) GetExtension() string {
	if o == nil || IsNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) GetExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) SetExtension(v string) {
	o.Extension = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) GetImage() UpdateNetworkWirelessSsidSplashSettingsRequestSplashImageImage {
	if o == nil || IsNil(o.Image) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSplashImageImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) GetImageOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashImageImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSplashImageImage and assigns it to the Image field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) SetImage(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashImageImage) {
	o.Image = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) ToMap() (map[string]interface{}, error) {
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

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


