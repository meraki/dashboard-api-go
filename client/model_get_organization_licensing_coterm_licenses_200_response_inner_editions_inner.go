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

// checks if the GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner{}

// GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner struct for GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner
type GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner struct {
	// The name of the license edition
	Edition *string `json:"edition,omitempty"`
	// The product type of the license edition
	ProductType *string `json:"productType,omitempty"`
}

// NewGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner instantiates a new GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner() *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner {
	this := GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner{}
	return &this
}

// NewGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInnerWithDefaults instantiates a new GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInnerWithDefaults() *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner {
	this := GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner{}
	return &this
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) GetEdition() string {
	if o == nil || IsNil(o.Edition) {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) GetEditionOk() (*string, bool) {
	if o == nil || IsNil(o.Edition) {
		return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) HasEdition() bool {
	if o != nil && !IsNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) SetEdition(v string) {
	o.Edition = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) SetProductType(v string) {
	o.ProductType = &v
}

func (o GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Edition) {
		toSerialize["edition"] = o.Edition
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return toSerialize, nil
}

type NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner struct {
	value *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner
	isSet bool
}

func (v NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) Get() *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner {
	return v.value
}

func (v *NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) Set(val *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner(val *GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) *NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner {
	return &NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


