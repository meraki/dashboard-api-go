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

// checks if the CreateOrganizationConfigTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationConfigTemplateRequest{}

// CreateOrganizationConfigTemplateRequest struct for CreateOrganizationConfigTemplateRequest
type CreateOrganizationConfigTemplateRequest struct {
	// The name of the configuration template
	Name string `json:"name"`
	// The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template
	TimeZone *string `json:"timeZone,omitempty"`
	// The ID of the network or config template to copy configuration from
	CopyFromNetworkId *string `json:"copyFromNetworkId,omitempty"`
}

// NewCreateOrganizationConfigTemplateRequest instantiates a new CreateOrganizationConfigTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationConfigTemplateRequest(name string) *CreateOrganizationConfigTemplateRequest {
	this := CreateOrganizationConfigTemplateRequest{}
	this.Name = name
	return &this
}

// NewCreateOrganizationConfigTemplateRequestWithDefaults instantiates a new CreateOrganizationConfigTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationConfigTemplateRequestWithDefaults() *CreateOrganizationConfigTemplateRequest {
	this := CreateOrganizationConfigTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationConfigTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationConfigTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationConfigTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *CreateOrganizationConfigTemplateRequest) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationConfigTemplateRequest) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *CreateOrganizationConfigTemplateRequest) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *CreateOrganizationConfigTemplateRequest) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetCopyFromNetworkId returns the CopyFromNetworkId field value if set, zero value otherwise.
func (o *CreateOrganizationConfigTemplateRequest) GetCopyFromNetworkId() string {
	if o == nil || IsNil(o.CopyFromNetworkId) {
		var ret string
		return ret
	}
	return *o.CopyFromNetworkId
}

// GetCopyFromNetworkIdOk returns a tuple with the CopyFromNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationConfigTemplateRequest) GetCopyFromNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.CopyFromNetworkId) {
		return nil, false
	}
	return o.CopyFromNetworkId, true
}

// HasCopyFromNetworkId returns a boolean if a field has been set.
func (o *CreateOrganizationConfigTemplateRequest) HasCopyFromNetworkId() bool {
	if o != nil && !IsNil(o.CopyFromNetworkId) {
		return true
	}

	return false
}

// SetCopyFromNetworkId gets a reference to the given string and assigns it to the CopyFromNetworkId field.
func (o *CreateOrganizationConfigTemplateRequest) SetCopyFromNetworkId(v string) {
	o.CopyFromNetworkId = &v
}

func (o CreateOrganizationConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationConfigTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.CopyFromNetworkId) {
		toSerialize["copyFromNetworkId"] = o.CopyFromNetworkId
	}
	return toSerialize, nil
}

type NullableCreateOrganizationConfigTemplateRequest struct {
	value *CreateOrganizationConfigTemplateRequest
	isSet bool
}

func (v NullableCreateOrganizationConfigTemplateRequest) Get() *CreateOrganizationConfigTemplateRequest {
	return v.value
}

func (v *NullableCreateOrganizationConfigTemplateRequest) Set(val *CreateOrganizationConfigTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationConfigTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationConfigTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationConfigTemplateRequest(val *CreateOrganizationConfigTemplateRequest) *NullableCreateOrganizationConfigTemplateRequest {
	return &NullableCreateOrganizationConfigTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationConfigTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


