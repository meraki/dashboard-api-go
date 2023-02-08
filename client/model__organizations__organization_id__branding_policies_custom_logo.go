/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdBrandingPoliciesCustomLogo Properties describing the custom logo attached to the branding policy.
type OrganizationsOrganizationIdBrandingPoliciesCustomLogo struct {
	// Whether or not there is a custom logo enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Image *OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage `json:"image,omitempty"`
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogo() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogo{}
	return &this
}

// NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoWithDefaults instantiates a new OrganizationsOrganizationIdBrandingPoliciesCustomLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdBrandingPoliciesCustomLogoWithDefaults() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo {
	this := OrganizationsOrganizationIdBrandingPoliciesCustomLogo{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) GetImage() OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage {
	if o == nil || isNil(o.Image) {
		var ret OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) GetImageOk() (*OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage and assigns it to the Image field.
func (o *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) SetImage(v OrganizationsOrganizationIdBrandingPoliciesCustomLogoImage) {
	o.Image = &v
}

func (o OrganizationsOrganizationIdBrandingPoliciesCustomLogo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo struct {
	value *OrganizationsOrganizationIdBrandingPoliciesCustomLogo
	isSet bool
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo) Get() *OrganizationsOrganizationIdBrandingPoliciesCustomLogo {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo) Set(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo(val *OrganizationsOrganizationIdBrandingPoliciesCustomLogo) *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo {
	return &NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdBrandingPoliciesCustomLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

