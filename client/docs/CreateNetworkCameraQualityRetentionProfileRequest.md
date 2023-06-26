# CreateNetworkCameraQualityRetentionProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new profile. Must be unique. This parameter is required. | 
**MotionBasedRetentionEnabled** | Pointer to **bool** | Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras. | [optional] 
**RestrictedBandwidthModeEnabled** | Pointer to **bool** | Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras. | [optional] 
**AudioRecordingEnabled** | Pointer to **bool** | Whether or not to record audio. Can be either true or false. Defaults to false. | [optional] 
**CloudArchiveEnabled** | Pointer to **bool** | Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false. | [optional] 
**MotionDetectorVersion** | Pointer to **int32** | The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2. | [optional] 
**ScheduleId** | Pointer to **string** | Schedule for which this camera will record video, or &#39;null&#39; to always record. | [optional] 
**MaxRetentionDays** | Pointer to **int32** | The maximum number of days for which the data will be stored, or &#39;null&#39; to keep data until storage space runs out. If the former, it can be one of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 14, 30, 60, 90] days. | [optional] 
**VideoSettings** | Pointer to [**CreateNetworkCameraQualityRetentionProfileRequestVideoSettings**](CreateNetworkCameraQualityRetentionProfileRequestVideoSettings.md) |  | [optional] 

## Methods

### NewCreateNetworkCameraQualityRetentionProfileRequest

`func NewCreateNetworkCameraQualityRetentionProfileRequest(name string, ) *CreateNetworkCameraQualityRetentionProfileRequest`

NewCreateNetworkCameraQualityRetentionProfileRequest instantiates a new CreateNetworkCameraQualityRetentionProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkCameraQualityRetentionProfileRequestWithDefaults

`func NewCreateNetworkCameraQualityRetentionProfileRequestWithDefaults() *CreateNetworkCameraQualityRetentionProfileRequest`

NewCreateNetworkCameraQualityRetentionProfileRequestWithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMotionBasedRetentionEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionBasedRetentionEnabled() bool`

GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field if non-nil, zero value otherwise.

### GetMotionBasedRetentionEnabledOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionBasedRetentionEnabledOk() (*bool, bool)`

GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionBasedRetentionEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetMotionBasedRetentionEnabled(v bool)`

SetMotionBasedRetentionEnabled sets MotionBasedRetentionEnabled field to given value.

### HasMotionBasedRetentionEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasMotionBasedRetentionEnabled() bool`

HasMotionBasedRetentionEnabled returns a boolean if a field has been set.

### GetRestrictedBandwidthModeEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetRestrictedBandwidthModeEnabled() bool`

GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field if non-nil, zero value otherwise.

### GetRestrictedBandwidthModeEnabledOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetRestrictedBandwidthModeEnabledOk() (*bool, bool)`

GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedBandwidthModeEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetRestrictedBandwidthModeEnabled(v bool)`

SetRestrictedBandwidthModeEnabled sets RestrictedBandwidthModeEnabled field to given value.

### HasRestrictedBandwidthModeEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasRestrictedBandwidthModeEnabled() bool`

HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.

### GetAudioRecordingEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetAudioRecordingEnabled() bool`

GetAudioRecordingEnabled returns the AudioRecordingEnabled field if non-nil, zero value otherwise.

### GetAudioRecordingEnabledOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetAudioRecordingEnabledOk() (*bool, bool)`

GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioRecordingEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetAudioRecordingEnabled(v bool)`

SetAudioRecordingEnabled sets AudioRecordingEnabled field to given value.

### HasAudioRecordingEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasAudioRecordingEnabled() bool`

HasAudioRecordingEnabled returns a boolean if a field has been set.

### GetCloudArchiveEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetCloudArchiveEnabled() bool`

GetCloudArchiveEnabled returns the CloudArchiveEnabled field if non-nil, zero value otherwise.

### GetCloudArchiveEnabledOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetCloudArchiveEnabledOk() (*bool, bool)`

GetCloudArchiveEnabledOk returns a tuple with the CloudArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudArchiveEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetCloudArchiveEnabled(v bool)`

SetCloudArchiveEnabled sets CloudArchiveEnabled field to given value.

### HasCloudArchiveEnabled

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasCloudArchiveEnabled() bool`

HasCloudArchiveEnabled returns a boolean if a field has been set.

### GetMotionDetectorVersion

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionDetectorVersion() int32`

GetMotionDetectorVersion returns the MotionDetectorVersion field if non-nil, zero value otherwise.

### GetMotionDetectorVersionOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMotionDetectorVersionOk() (*int32, bool)`

GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionDetectorVersion

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetMotionDetectorVersion(v int32)`

SetMotionDetectorVersion sets MotionDetectorVersion field to given value.

### HasMotionDetectorVersion

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasMotionDetectorVersion() bool`

HasMotionDetectorVersion returns a boolean if a field has been set.

### GetScheduleId

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetMaxRetentionDays

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMaxRetentionDays() int32`

GetMaxRetentionDays returns the MaxRetentionDays field if non-nil, zero value otherwise.

### GetMaxRetentionDaysOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetMaxRetentionDaysOk() (*int32, bool)`

GetMaxRetentionDaysOk returns a tuple with the MaxRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetentionDays

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetMaxRetentionDays(v int32)`

SetMaxRetentionDays sets MaxRetentionDays field to given value.

### HasMaxRetentionDays

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasMaxRetentionDays() bool`

HasMaxRetentionDays returns a boolean if a field has been set.

### GetVideoSettings

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetVideoSettings() CreateNetworkCameraQualityRetentionProfileRequestVideoSettings`

GetVideoSettings returns the VideoSettings field if non-nil, zero value otherwise.

### GetVideoSettingsOk

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) GetVideoSettingsOk() (*CreateNetworkCameraQualityRetentionProfileRequestVideoSettings, bool)`

GetVideoSettingsOk returns a tuple with the VideoSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSettings

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) SetVideoSettings(v CreateNetworkCameraQualityRetentionProfileRequestVideoSettings)`

SetVideoSettings sets VideoSettings field to given value.

### HasVideoSettings

`func (o *CreateNetworkCameraQualityRetentionProfileRequest) HasVideoSettings() bool`

HasVideoSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


