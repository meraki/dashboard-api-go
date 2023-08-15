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

// checks if the CreateNetworkCameraQualityRetentionProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraQualityRetentionProfileRequest{}

// CreateNetworkCameraQualityRetentionProfileRequest struct for CreateNetworkCameraQualityRetentionProfileRequest
type CreateNetworkCameraQualityRetentionProfileRequest struct {
	// The name of the new profile. Must be unique. This parameter is required.
	Name string `json:"name"`
	// Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	MotionBasedRetentionEnabled *bool `json:"motionBasedRetentionEnabled,omitempty"`
	// Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	RestrictedBandwidthModeEnabled *bool `json:"restrictedBandwidthModeEnabled,omitempty"`
	// Whether or not to record audio. Can be either true or false. Defaults to false.
	AudioRecordingEnabled *bool `json:"audioRecordingEnabled,omitempty"`
	// Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.
	CloudArchiveEnabled *bool `json:"cloudArchiveEnabled,omitempty"`
	// The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	MotionDetectorVersion *int32 `json:"motionDetectorVersion,omitempty"`
	// Schedule for which this camera will record video, or 'null' to always record.
	ScheduleId *string `json:"scheduleId,omitempty"`
	// The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be one of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 14, 30, 60, 90] days.
	MaxRetentionDays *int32 `json:"maxRetentionDays,omitempty"`
	VideoSettings *CreateNetworkCameraQualityRetentionProfileRequestVideoSettings `json:"videoSettings,omitempty"`
}

// NewCreateNetworkCameraQualityRetentionProfileRequest instantiates a new CreateNetworkCameraQualityRetentionProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequest(name string) *CreateNetworkCameraQualityRetentionProfileRequest {
	this := CreateNetworkCameraQualityRetentionProfileRequest{}
	this.Name = name
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestWithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestWithDefaults() *CreateNetworkCameraQualityRetentionProfileRequest {
	this := CreateNetworkCameraQualityRetentionProfileRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetName(v string) {
	o.Name = v
}

// GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionBasedRetentionEnabled() bool {
	if o == nil || IsNil(o.MotionBasedRetentionEnabled) {
		var ret bool
		return ret
	}
	return *o.MotionBasedRetentionEnabled
}

// GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionBasedRetentionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MotionBasedRetentionEnabled) {
		return nil, false
	}
	return o.MotionBasedRetentionEnabled, true
}

// HasMotionBasedRetentionEnabled returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasMotionBasedRetentionEnabled() bool {
	if o != nil && !IsNil(o.MotionBasedRetentionEnabled) {
		return true
	}

	return false
}

// SetMotionBasedRetentionEnabled gets a reference to the given bool and assigns it to the MotionBasedRetentionEnabled field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetMotionBasedRetentionEnabled(v bool) {
	o.MotionBasedRetentionEnabled = &v
}

// GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetRestrictedBandwidthModeEnabled() bool {
	if o == nil || IsNil(o.RestrictedBandwidthModeEnabled) {
		var ret bool
		return ret
	}
	return *o.RestrictedBandwidthModeEnabled
}

// GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetRestrictedBandwidthModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictedBandwidthModeEnabled) {
		return nil, false
	}
	return o.RestrictedBandwidthModeEnabled, true
}

// HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasRestrictedBandwidthModeEnabled() bool {
	if o != nil && !IsNil(o.RestrictedBandwidthModeEnabled) {
		return true
	}

	return false
}

// SetRestrictedBandwidthModeEnabled gets a reference to the given bool and assigns it to the RestrictedBandwidthModeEnabled field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetRestrictedBandwidthModeEnabled(v bool) {
	o.RestrictedBandwidthModeEnabled = &v
}

// GetAudioRecordingEnabled returns the AudioRecordingEnabled field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetAudioRecordingEnabled() bool {
	if o == nil || IsNil(o.AudioRecordingEnabled) {
		var ret bool
		return ret
	}
	return *o.AudioRecordingEnabled
}

// GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetAudioRecordingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AudioRecordingEnabled) {
		return nil, false
	}
	return o.AudioRecordingEnabled, true
}

// HasAudioRecordingEnabled returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasAudioRecordingEnabled() bool {
	if o != nil && !IsNil(o.AudioRecordingEnabled) {
		return true
	}

	return false
}

// SetAudioRecordingEnabled gets a reference to the given bool and assigns it to the AudioRecordingEnabled field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetAudioRecordingEnabled(v bool) {
	o.AudioRecordingEnabled = &v
}

// GetCloudArchiveEnabled returns the CloudArchiveEnabled field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetCloudArchiveEnabled() bool {
	if o == nil || IsNil(o.CloudArchiveEnabled) {
		var ret bool
		return ret
	}
	return *o.CloudArchiveEnabled
}

// GetCloudArchiveEnabledOk returns a tuple with the CloudArchiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetCloudArchiveEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CloudArchiveEnabled) {
		return nil, false
	}
	return o.CloudArchiveEnabled, true
}

// HasCloudArchiveEnabled returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasCloudArchiveEnabled() bool {
	if o != nil && !IsNil(o.CloudArchiveEnabled) {
		return true
	}

	return false
}

// SetCloudArchiveEnabled gets a reference to the given bool and assigns it to the CloudArchiveEnabled field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetCloudArchiveEnabled(v bool) {
	o.CloudArchiveEnabled = &v
}

// GetMotionDetectorVersion returns the MotionDetectorVersion field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionDetectorVersion() int32 {
	if o == nil || IsNil(o.MotionDetectorVersion) {
		var ret int32
		return ret
	}
	return *o.MotionDetectorVersion
}

// GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionDetectorVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MotionDetectorVersion) {
		return nil, false
	}
	return o.MotionDetectorVersion, true
}

// HasMotionDetectorVersion returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasMotionDetectorVersion() bool {
	if o != nil && !IsNil(o.MotionDetectorVersion) {
		return true
	}

	return false
}

// SetMotionDetectorVersion gets a reference to the given int32 and assigns it to the MotionDetectorVersion field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetMotionDetectorVersion(v int32) {
	o.MotionDetectorVersion = &v
}

// GetScheduleId returns the ScheduleId field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetScheduleId() string {
	if o == nil || IsNil(o.ScheduleId) {
		var ret string
		return ret
	}
	return *o.ScheduleId
}

// GetScheduleIdOk returns a tuple with the ScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetScheduleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleId) {
		return nil, false
	}
	return o.ScheduleId, true
}

// HasScheduleId returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasScheduleId() bool {
	if o != nil && !IsNil(o.ScheduleId) {
		return true
	}

	return false
}

// SetScheduleId gets a reference to the given string and assigns it to the ScheduleId field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetScheduleId(v string) {
	o.ScheduleId = &v
}

// GetMaxRetentionDays returns the MaxRetentionDays field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMaxRetentionDays() int32 {
	if o == nil || IsNil(o.MaxRetentionDays) {
		var ret int32
		return ret
	}
	return *o.MaxRetentionDays
}

// GetMaxRetentionDaysOk returns a tuple with the MaxRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMaxRetentionDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRetentionDays) {
		return nil, false
	}
	return o.MaxRetentionDays, true
}

// HasMaxRetentionDays returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasMaxRetentionDays() bool {
	if o != nil && !IsNil(o.MaxRetentionDays) {
		return true
	}

	return false
}

// SetMaxRetentionDays gets a reference to the given int32 and assigns it to the MaxRetentionDays field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetMaxRetentionDays(v int32) {
	o.MaxRetentionDays = &v
}

// GetVideoSettings returns the VideoSettings field value if set, zero value otherwise.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetVideoSettings() CreateNetworkCameraQualityRetentionProfileRequestVideoSettings {
	if o == nil || IsNil(o.VideoSettings) {
		var ret CreateNetworkCameraQualityRetentionProfileRequestVideoSettings
		return ret
	}
	return *o.VideoSettings
}

// GetVideoSettingsOk returns a tuple with the VideoSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetVideoSettingsOk() (*CreateNetworkCameraQualityRetentionProfileRequestVideoSettings, bool) {
	if o == nil || IsNil(o.VideoSettings) {
		return nil, false
	}
	return o.VideoSettings, true
}

// HasVideoSettings returns a boolean if a field has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasVideoSettings() bool {
	if o != nil && !IsNil(o.VideoSettings) {
		return true
	}

	return false
}

// SetVideoSettings gets a reference to the given CreateNetworkCameraQualityRetentionProfileRequestVideoSettings and assigns it to the VideoSettings field.
func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetVideoSettings(v CreateNetworkCameraQualityRetentionProfileRequestVideoSettings) {
	o.VideoSettings = &v
}

func (o CreateNetworkCameraQualityRetentionProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraQualityRetentionProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.MotionBasedRetentionEnabled) {
		toSerialize["motionBasedRetentionEnabled"] = o.MotionBasedRetentionEnabled
	}
	if !IsNil(o.RestrictedBandwidthModeEnabled) {
		toSerialize["restrictedBandwidthModeEnabled"] = o.RestrictedBandwidthModeEnabled
	}
	if !IsNil(o.AudioRecordingEnabled) {
		toSerialize["audioRecordingEnabled"] = o.AudioRecordingEnabled
	}
	if !IsNil(o.CloudArchiveEnabled) {
		toSerialize["cloudArchiveEnabled"] = o.CloudArchiveEnabled
	}
	if !IsNil(o.MotionDetectorVersion) {
		toSerialize["motionDetectorVersion"] = o.MotionDetectorVersion
	}
	if !IsNil(o.ScheduleId) {
		toSerialize["scheduleId"] = o.ScheduleId
	}
	if !IsNil(o.MaxRetentionDays) {
		toSerialize["maxRetentionDays"] = o.MaxRetentionDays
	}
	if !IsNil(o.VideoSettings) {
		toSerialize["videoSettings"] = o.VideoSettings
	}
	return toSerialize, nil
}

type NullableCreateNetworkCameraQualityRetentionProfileRequest struct {
	value *CreateNetworkCameraQualityRetentionProfileRequest
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequest) Get() *CreateNetworkCameraQualityRetentionProfileRequest {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequest) Set(val *CreateNetworkCameraQualityRetentionProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequest(val *CreateNetworkCameraQualityRetentionProfileRequest) *NullableCreateNetworkCameraQualityRetentionProfileRequest {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


