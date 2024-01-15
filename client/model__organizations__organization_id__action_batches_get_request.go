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

// checks if the OrganizationsOrganizationIdActionBatchesGetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationIdActionBatchesGetRequest{}

// OrganizationsOrganizationIdActionBatchesGetRequest struct for OrganizationsOrganizationIdActionBatchesGetRequest
type OrganizationsOrganizationIdActionBatchesGetRequest struct {
	Organization *DevicesSerialLiveToolsPingPostRequestOrganization `json:"organization,omitempty"`
	Network *DevicesSerialLiveToolsPingPostRequestOrganization `json:"network,omitempty"`
	SentAt *string `json:"sentAt,omitempty"`
	CallbackId *string `json:"callbackId,omitempty"`
	Message *GetOrganizationActionBatches200ResponseInner `json:"message,omitempty"`
}

// NewOrganizationsOrganizationIdActionBatchesGetRequest instantiates a new OrganizationsOrganizationIdActionBatchesGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdActionBatchesGetRequest() *OrganizationsOrganizationIdActionBatchesGetRequest {
	this := OrganizationsOrganizationIdActionBatchesGetRequest{}
	return &this
}

// NewOrganizationsOrganizationIdActionBatchesGetRequestWithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdActionBatchesGetRequestWithDefaults() *OrganizationsOrganizationIdActionBatchesGetRequest {
	this := OrganizationsOrganizationIdActionBatchesGetRequest{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetOrganization() DevicesSerialLiveToolsPingPostRequestOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret DevicesSerialLiveToolsPingPostRequestOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetOrganizationOk() (*DevicesSerialLiveToolsPingPostRequestOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given DevicesSerialLiveToolsPingPostRequestOrganization and assigns it to the Organization field.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) SetOrganization(v DevicesSerialLiveToolsPingPostRequestOrganization) {
	o.Organization = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetNetwork() DevicesSerialLiveToolsPingPostRequestOrganization {
	if o == nil || IsNil(o.Network) {
		var ret DevicesSerialLiveToolsPingPostRequestOrganization
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetNetworkOk() (*DevicesSerialLiveToolsPingPostRequestOrganization, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given DevicesSerialLiveToolsPingPostRequestOrganization and assigns it to the Network field.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) SetNetwork(v DevicesSerialLiveToolsPingPostRequestOrganization) {
	o.Network = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetSentAt() string {
	if o == nil || IsNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetSentAtOk() (*string, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) SetSentAt(v string) {
	o.SentAt = &v
}

// GetCallbackId returns the CallbackId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetCallbackId() string {
	if o == nil || IsNil(o.CallbackId) {
		var ret string
		return ret
	}
	return *o.CallbackId
}

// GetCallbackIdOk returns a tuple with the CallbackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetCallbackIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackId) {
		return nil, false
	}
	return o.CallbackId, true
}

// HasCallbackId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) HasCallbackId() bool {
	if o != nil && !IsNil(o.CallbackId) {
		return true
	}

	return false
}

// SetCallbackId gets a reference to the given string and assigns it to the CallbackId field.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) SetCallbackId(v string) {
	o.CallbackId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetMessage() GetOrganizationActionBatches200ResponseInner {
	if o == nil || IsNil(o.Message) {
		var ret GetOrganizationActionBatches200ResponseInner
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) GetMessageOk() (*GetOrganizationActionBatches200ResponseInner, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given GetOrganizationActionBatches200ResponseInner and assigns it to the Message field.
func (o *OrganizationsOrganizationIdActionBatchesGetRequest) SetMessage(v GetOrganizationActionBatches200ResponseInner) {
	o.Message = &v
}

func (o OrganizationsOrganizationIdActionBatchesGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationIdActionBatchesGetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !IsNil(o.CallbackId) {
		toSerialize["callbackId"] = o.CallbackId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableOrganizationsOrganizationIdActionBatchesGetRequest struct {
	value *OrganizationsOrganizationIdActionBatchesGetRequest
	isSet bool
}

func (v NullableOrganizationsOrganizationIdActionBatchesGetRequest) Get() *OrganizationsOrganizationIdActionBatchesGetRequest {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdActionBatchesGetRequest) Set(val *OrganizationsOrganizationIdActionBatchesGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdActionBatchesGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdActionBatchesGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdActionBatchesGetRequest(val *OrganizationsOrganizationIdActionBatchesGetRequest) *NullableOrganizationsOrganizationIdActionBatchesGetRequest {
	return &NullableOrganizationsOrganizationIdActionBatchesGetRequest{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdActionBatchesGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdActionBatchesGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


