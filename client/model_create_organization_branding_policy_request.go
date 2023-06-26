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

// CreateOrganizationBrandingPolicyRequest struct for CreateOrganizationBrandingPolicyRequest
type CreateOrganizationBrandingPolicyRequest struct {
	// Name of the Dashboard branding policy.
	Name *string `json:"name,omitempty"`
	// Boolean indicating whether this policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	AdminSettings *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings `json:"adminSettings,omitempty"`
	HelpSettings *CreateOrganizationBrandingPolicyRequestHelpSettings `json:"helpSettings,omitempty"`
	CustomLogo *CreateOrganizationBrandingPolicyRequestCustomLogo `json:"customLogo,omitempty"`
}

// NewCreateOrganizationBrandingPolicyRequest instantiates a new CreateOrganizationBrandingPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationBrandingPolicyRequest() *CreateOrganizationBrandingPolicyRequest {
	this := CreateOrganizationBrandingPolicyRequest{}
	return &this
}

// NewCreateOrganizationBrandingPolicyRequestWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationBrandingPolicyRequestWithDefaults() *CreateOrganizationBrandingPolicyRequest {
	this := CreateOrganizationBrandingPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationBrandingPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateOrganizationBrandingPolicyRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdminSettings returns the AdminSettings field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequest) GetAdminSettings() GetOrganizationBrandingPolicies200ResponseInnerAdminSettings {
	if o == nil || isNil(o.AdminSettings) {
		var ret GetOrganizationBrandingPolicies200ResponseInnerAdminSettings
		return ret
	}
	return *o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetAdminSettingsOk() (*GetOrganizationBrandingPolicies200ResponseInnerAdminSettings, bool) {
	if o == nil || isNil(o.AdminSettings) {
    return nil, false
	}
	return o.AdminSettings, true
}

// HasAdminSettings returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequest) HasAdminSettings() bool {
	if o != nil && !isNil(o.AdminSettings) {
		return true
	}

	return false
}

// SetAdminSettings gets a reference to the given GetOrganizationBrandingPolicies200ResponseInnerAdminSettings and assigns it to the AdminSettings field.
func (o *CreateOrganizationBrandingPolicyRequest) SetAdminSettings(v GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) {
	o.AdminSettings = &v
}

// GetHelpSettings returns the HelpSettings field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequest) GetHelpSettings() CreateOrganizationBrandingPolicyRequestHelpSettings {
	if o == nil || isNil(o.HelpSettings) {
		var ret CreateOrganizationBrandingPolicyRequestHelpSettings
		return ret
	}
	return *o.HelpSettings
}

// GetHelpSettingsOk returns a tuple with the HelpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetHelpSettingsOk() (*CreateOrganizationBrandingPolicyRequestHelpSettings, bool) {
	if o == nil || isNil(o.HelpSettings) {
    return nil, false
	}
	return o.HelpSettings, true
}

// HasHelpSettings returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequest) HasHelpSettings() bool {
	if o != nil && !isNil(o.HelpSettings) {
		return true
	}

	return false
}

// SetHelpSettings gets a reference to the given CreateOrganizationBrandingPolicyRequestHelpSettings and assigns it to the HelpSettings field.
func (o *CreateOrganizationBrandingPolicyRequest) SetHelpSettings(v CreateOrganizationBrandingPolicyRequestHelpSettings) {
	o.HelpSettings = &v
}

// GetCustomLogo returns the CustomLogo field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequest) GetCustomLogo() CreateOrganizationBrandingPolicyRequestCustomLogo {
	if o == nil || isNil(o.CustomLogo) {
		var ret CreateOrganizationBrandingPolicyRequestCustomLogo
		return ret
	}
	return *o.CustomLogo
}

// GetCustomLogoOk returns a tuple with the CustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetCustomLogoOk() (*CreateOrganizationBrandingPolicyRequestCustomLogo, bool) {
	if o == nil || isNil(o.CustomLogo) {
    return nil, false
	}
	return o.CustomLogo, true
}

// HasCustomLogo returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequest) HasCustomLogo() bool {
	if o != nil && !isNil(o.CustomLogo) {
		return true
	}

	return false
}

// SetCustomLogo gets a reference to the given CreateOrganizationBrandingPolicyRequestCustomLogo and assigns it to the CustomLogo field.
func (o *CreateOrganizationBrandingPolicyRequest) SetCustomLogo(v CreateOrganizationBrandingPolicyRequestCustomLogo) {
	o.CustomLogo = &v
}

func (o CreateOrganizationBrandingPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AdminSettings) {
		toSerialize["adminSettings"] = o.AdminSettings
	}
	if !isNil(o.HelpSettings) {
		toSerialize["helpSettings"] = o.HelpSettings
	}
	if !isNil(o.CustomLogo) {
		toSerialize["customLogo"] = o.CustomLogo
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationBrandingPolicyRequest struct {
	value *CreateOrganizationBrandingPolicyRequest
	isSet bool
}

func (v NullableCreateOrganizationBrandingPolicyRequest) Get() *CreateOrganizationBrandingPolicyRequest {
	return v.value
}

func (v *NullableCreateOrganizationBrandingPolicyRequest) Set(val *CreateOrganizationBrandingPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationBrandingPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationBrandingPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationBrandingPolicyRequest(val *CreateOrganizationBrandingPolicyRequest) *NullableCreateOrganizationBrandingPolicyRequest {
	return &NullableCreateOrganizationBrandingPolicyRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationBrandingPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationBrandingPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


