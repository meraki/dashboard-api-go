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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate Root certificate information
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate struct {
	// Public certificate value
	Content *string `json:"content,omitempty"`
	// The name of the server protected by the certificate
	Name *string `json:"name,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificateWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificateWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) SetContent(v string) {
	o.Content = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) SetName(v string) {
	o.Name = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParamsTunnelRootCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


