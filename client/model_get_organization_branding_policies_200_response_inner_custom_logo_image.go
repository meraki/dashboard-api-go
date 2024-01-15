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

// checks if the GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage{}

// GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage Properties of the image.
type GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage struct {
	Preview *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview `json:"preview,omitempty"`
}

// NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage instantiates a new GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage {
	this := GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage{}
	return &this
}

// NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImageWithDefaults instantiates a new GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImageWithDefaults() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage {
	this := GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage{}
	return &this
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) GetPreview() GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview {
	if o == nil || IsNil(o.Preview) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) GetPreviewOk() (*GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview, bool) {
	if o == nil || IsNil(o.Preview) {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) HasPreview() bool {
	if o != nil && !IsNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview and assigns it to the Preview field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) SetPreview(v GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) {
	o.Preview = &v
}

func (o GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Preview) {
		toSerialize["preview"] = o.Preview
	}
	return toSerialize, nil
}

type NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage struct {
	value *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage
	isSet bool
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) Get() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage {
	return v.value
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) Set(val *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage(val *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage {
	return &NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage{value: val, isSet: true}
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


