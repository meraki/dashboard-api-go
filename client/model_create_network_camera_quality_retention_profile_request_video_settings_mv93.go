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

// checks if the CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93{}

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 Quality and resolution for MV93 camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
	Resolution string `json:"resolution"`
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93WithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93WithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) GetQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) GetResolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quality"] = o.Quality
	toSerialize["resolution"] = o.Resolution
	return toSerialize, nil
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93 {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


