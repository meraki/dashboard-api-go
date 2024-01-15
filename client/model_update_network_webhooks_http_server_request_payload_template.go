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

// checks if the UpdateNetworkWebhooksHttpServerRequestPayloadTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWebhooksHttpServerRequestPayloadTemplate{}

// UpdateNetworkWebhooksHttpServerRequestPayloadTemplate The payload template to use when posting data to the HTTP server.
type UpdateNetworkWebhooksHttpServerRequestPayloadTemplate struct {
	// The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Meraki-included templates: for the Webex (included) template use 'wpt_00002'; for the Slack (included) template use 'wpt_00003'; for the Microsoft Teams (included) template use 'wpt_00004'; for the ServiceNow (included) template use 'wpt_00006'
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
}

// NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplate instantiates a new UpdateNetworkWebhooksHttpServerRequestPayloadTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplate() *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate {
	this := UpdateNetworkWebhooksHttpServerRequestPayloadTemplate{}
	return &this
}

// NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults instantiates a new UpdateNetworkWebhooksHttpServerRequestPayloadTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults() *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate {
	this := UpdateNetworkWebhooksHttpServerRequestPayloadTemplate{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateId() string {
	if o == nil || IsNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadTemplateId) {
		return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) HasPayloadTemplateId() bool {
	if o != nil && !IsNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

func (o UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate struct {
	value *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate
	isSet bool
}

func (v NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate) Get() *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate {
	return v.value
}

func (v *NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate) Set(val *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate(val *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) *NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate {
	return &NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate{value: val, isSet: true}
}

func (v NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWebhooksHttpServerRequestPayloadTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


