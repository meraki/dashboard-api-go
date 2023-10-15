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

// checks if the UpdateOrganizationAlertsProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationAlertsProfileRequest{}

// UpdateOrganizationAlertsProfileRequest struct for UpdateOrganizationAlertsProfileRequest
type UpdateOrganizationAlertsProfileRequest struct {
	// Is the alert config enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The alert type
	Type *string `json:"type,omitempty"`
	AlertCondition *CreateOrganizationAlertsProfileRequestAlertCondition `json:"alertCondition,omitempty"`
	Recipients *CreateOrganizationAlertsProfileRequestRecipients `json:"recipients,omitempty"`
	// Networks with these tags will be monitored for the alert
	NetworkTags []string `json:"networkTags,omitempty"`
	// User supplied description of the alert
	Description *string `json:"description,omitempty"`
}

// NewUpdateOrganizationAlertsProfileRequest instantiates a new UpdateOrganizationAlertsProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationAlertsProfileRequest() *UpdateOrganizationAlertsProfileRequest {
	this := UpdateOrganizationAlertsProfileRequest{}
	return &this
}

// NewUpdateOrganizationAlertsProfileRequestWithDefaults instantiates a new UpdateOrganizationAlertsProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationAlertsProfileRequestWithDefaults() *UpdateOrganizationAlertsProfileRequest {
	this := UpdateOrganizationAlertsProfileRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateOrganizationAlertsProfileRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAlertsProfileRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationAlertsProfileRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateOrganizationAlertsProfileRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateOrganizationAlertsProfileRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAlertsProfileRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateOrganizationAlertsProfileRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateOrganizationAlertsProfileRequest) SetType(v string) {
	o.Type = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise.
func (o *UpdateOrganizationAlertsProfileRequest) GetAlertCondition() CreateOrganizationAlertsProfileRequestAlertCondition {
	if o == nil || IsNil(o.AlertCondition) {
		var ret CreateOrganizationAlertsProfileRequestAlertCondition
		return ret
	}
	return *o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAlertsProfileRequest) GetAlertConditionOk() (*CreateOrganizationAlertsProfileRequestAlertCondition, bool) {
	if o == nil || IsNil(o.AlertCondition) {
		return nil, false
	}
	return o.AlertCondition, true
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *UpdateOrganizationAlertsProfileRequest) HasAlertCondition() bool {
	if o != nil && !IsNil(o.AlertCondition) {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given CreateOrganizationAlertsProfileRequestAlertCondition and assigns it to the AlertCondition field.
func (o *UpdateOrganizationAlertsProfileRequest) SetAlertCondition(v CreateOrganizationAlertsProfileRequestAlertCondition) {
	o.AlertCondition = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *UpdateOrganizationAlertsProfileRequest) GetRecipients() CreateOrganizationAlertsProfileRequestRecipients {
	if o == nil || IsNil(o.Recipients) {
		var ret CreateOrganizationAlertsProfileRequestRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAlertsProfileRequest) GetRecipientsOk() (*CreateOrganizationAlertsProfileRequestRecipients, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *UpdateOrganizationAlertsProfileRequest) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given CreateOrganizationAlertsProfileRequestRecipients and assigns it to the Recipients field.
func (o *UpdateOrganizationAlertsProfileRequest) SetRecipients(v CreateOrganizationAlertsProfileRequestRecipients) {
	o.Recipients = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *UpdateOrganizationAlertsProfileRequest) GetNetworkTags() []string {
	if o == nil || IsNil(o.NetworkTags) {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAlertsProfileRequest) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkTags) {
		return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *UpdateOrganizationAlertsProfileRequest) HasNetworkTags() bool {
	if o != nil && !IsNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *UpdateOrganizationAlertsProfileRequest) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateOrganizationAlertsProfileRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAlertsProfileRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateOrganizationAlertsProfileRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateOrganizationAlertsProfileRequest) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateOrganizationAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationAlertsProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AlertCondition) {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.NetworkTags) {
		toSerialize["networkTags"] = o.NetworkTags
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationAlertsProfileRequest struct {
	value *UpdateOrganizationAlertsProfileRequest
	isSet bool
}

func (v NullableUpdateOrganizationAlertsProfileRequest) Get() *UpdateOrganizationAlertsProfileRequest {
	return v.value
}

func (v *NullableUpdateOrganizationAlertsProfileRequest) Set(val *UpdateOrganizationAlertsProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationAlertsProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationAlertsProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationAlertsProfileRequest(val *UpdateOrganizationAlertsProfileRequest) *NullableUpdateOrganizationAlertsProfileRequest {
	return &NullableUpdateOrganizationAlertsProfileRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationAlertsProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


