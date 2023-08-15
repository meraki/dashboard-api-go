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

// checks if the UpdateOrganizationSamlIdpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationSamlIdpRequest{}

// UpdateOrganizationSamlIdpRequest struct for UpdateOrganizationSamlIdpRequest
type UpdateOrganizationSamlIdpRequest struct {
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint *string `json:"x509certSha1Fingerprint,omitempty"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewUpdateOrganizationSamlIdpRequest instantiates a new UpdateOrganizationSamlIdpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationSamlIdpRequest() *UpdateOrganizationSamlIdpRequest {
	this := UpdateOrganizationSamlIdpRequest{}
	return &this
}

// NewUpdateOrganizationSamlIdpRequestWithDefaults instantiates a new UpdateOrganizationSamlIdpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationSamlIdpRequestWithDefaults() *UpdateOrganizationSamlIdpRequest {
	this := UpdateOrganizationSamlIdpRequest{}
	return &this
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlIdpRequest) GetX509certSha1Fingerprint() string {
	if o == nil || IsNil(o.X509certSha1Fingerprint) {
		var ret string
		return ret
	}
	return *o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlIdpRequest) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.X509certSha1Fingerprint) {
		return nil, false
	}
	return o.X509certSha1Fingerprint, true
}

// HasX509certSha1Fingerprint returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlIdpRequest) HasX509certSha1Fingerprint() bool {
	if o != nil && !IsNil(o.X509certSha1Fingerprint) {
		return true
	}

	return false
}

// SetX509certSha1Fingerprint gets a reference to the given string and assigns it to the X509certSha1Fingerprint field.
func (o *UpdateOrganizationSamlIdpRequest) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = &v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlIdpRequest) GetSloLogoutUrl() string {
	if o == nil || IsNil(o.SloLogoutUrl) {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlIdpRequest) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SloLogoutUrl) {
		return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlIdpRequest) HasSloLogoutUrl() bool {
	if o != nil && !IsNil(o.SloLogoutUrl) {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *UpdateOrganizationSamlIdpRequest) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o UpdateOrganizationSamlIdpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationSamlIdpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.X509certSha1Fingerprint) {
		toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	}
	if !IsNil(o.SloLogoutUrl) {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationSamlIdpRequest struct {
	value *UpdateOrganizationSamlIdpRequest
	isSet bool
}

func (v NullableUpdateOrganizationSamlIdpRequest) Get() *UpdateOrganizationSamlIdpRequest {
	return v.value
}

func (v *NullableUpdateOrganizationSamlIdpRequest) Set(val *UpdateOrganizationSamlIdpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationSamlIdpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationSamlIdpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationSamlIdpRequest(val *UpdateOrganizationSamlIdpRequest) *NullableUpdateOrganizationSamlIdpRequest {
	return &NullableUpdateOrganizationSamlIdpRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationSamlIdpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationSamlIdpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


