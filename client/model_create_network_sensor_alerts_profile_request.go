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

// checks if the CreateNetworkSensorAlertsProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSensorAlertsProfileRequest{}

// CreateNetworkSensorAlertsProfileRequest struct for CreateNetworkSensorAlertsProfileRequest
type CreateNetworkSensorAlertsProfileRequest struct {
	// Name of the sensor alert profile.
	Name string `json:"name"`
	Schedule *CreateNetworkSensorAlertsProfileRequestSchedule `json:"schedule,omitempty"`
	// List of conditions that will cause the profile to send an alert.
	Conditions []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner `json:"conditions"`
	Recipients *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients `json:"recipients,omitempty"`
	// List of device serials assigned to this sensor alert profile.
	Serials []string `json:"serials,omitempty"`
}

// NewCreateNetworkSensorAlertsProfileRequest instantiates a new CreateNetworkSensorAlertsProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSensorAlertsProfileRequest(name string, conditions []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) *CreateNetworkSensorAlertsProfileRequest {
	this := CreateNetworkSensorAlertsProfileRequest{}
	this.Name = name
	this.Conditions = conditions
	return &this
}

// NewCreateNetworkSensorAlertsProfileRequestWithDefaults instantiates a new CreateNetworkSensorAlertsProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSensorAlertsProfileRequestWithDefaults() *CreateNetworkSensorAlertsProfileRequest {
	this := CreateNetworkSensorAlertsProfileRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkSensorAlertsProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkSensorAlertsProfileRequest) SetName(v string) {
	o.Name = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CreateNetworkSensorAlertsProfileRequest) GetSchedule() CreateNetworkSensorAlertsProfileRequestSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret CreateNetworkSensorAlertsProfileRequestSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) GetScheduleOk() (*CreateNetworkSensorAlertsProfileRequestSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given CreateNetworkSensorAlertsProfileRequestSchedule and assigns it to the Schedule field.
func (o *CreateNetworkSensorAlertsProfileRequest) SetSchedule(v CreateNetworkSensorAlertsProfileRequestSchedule) {
	o.Schedule = &v
}

// GetConditions returns the Conditions field value
func (o *CreateNetworkSensorAlertsProfileRequest) GetConditions() []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner {
	if o == nil {
		var ret []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) GetConditionsOk() ([]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *CreateNetworkSensorAlertsProfileRequest) SetConditions(v []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) {
	o.Conditions = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *CreateNetworkSensorAlertsProfileRequest) GetRecipients() GetNetworkSensorAlertsProfiles200ResponseInnerRecipients {
	if o == nil || IsNil(o.Recipients) {
		var ret GetNetworkSensorAlertsProfiles200ResponseInnerRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) GetRecipientsOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerRecipients, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given GetNetworkSensorAlertsProfiles200ResponseInnerRecipients and assigns it to the Recipients field.
func (o *CreateNetworkSensorAlertsProfileRequest) SetRecipients(v GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) {
	o.Recipients = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *CreateNetworkSensorAlertsProfileRequest) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *CreateNetworkSensorAlertsProfileRequest) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *CreateNetworkSensorAlertsProfileRequest) SetSerials(v []string) {
	o.Serials = v
}

func (o CreateNetworkSensorAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSensorAlertsProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	toSerialize["conditions"] = o.Conditions
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return toSerialize, nil
}

type NullableCreateNetworkSensorAlertsProfileRequest struct {
	value *CreateNetworkSensorAlertsProfileRequest
	isSet bool
}

func (v NullableCreateNetworkSensorAlertsProfileRequest) Get() *CreateNetworkSensorAlertsProfileRequest {
	return v.value
}

func (v *NullableCreateNetworkSensorAlertsProfileRequest) Set(val *CreateNetworkSensorAlertsProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSensorAlertsProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSensorAlertsProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSensorAlertsProfileRequest(val *CreateNetworkSensorAlertsProfileRequest) *NullableCreateNetworkSensorAlertsProfileRequest {
	return &NullableCreateNetworkSensorAlertsProfileRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkSensorAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSensorAlertsProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


