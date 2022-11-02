# UpdateDeviceCameraQualityAndRetentionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | The ID of a quality and retention profile to assign to the camera. The profile&#39;s settings will override all of the per-camera quality and retention settings. If the value of this parameter is null, any existing profile will be unassigned from the camera. | [optional] 
**MotionBasedRetentionEnabled** | Pointer to **bool** | Boolean indicating if motion-based retention is enabled(true) or disabled(false) on the camera. | [optional] 
**AudioRecordingEnabled** | Pointer to **bool** | Boolean indicating if audio recording is enabled(true) or disabled(false) on the camera | [optional] 
**RestrictedBandwidthModeEnabled** | Pointer to **bool** | Boolean indicating if restricted bandwidth is enabled(true) or disabled(false) on the camera. This setting does not apply to MV2 cameras. | [optional] 
**Quality** | Pointer to **string** | Quality of the camera. Can be one of &#39;Standard&#39;, &#39;High&#39; or &#39;Enhanced&#39;. Not all qualities are supported by every camera model. | [optional] 
**Resolution** | Pointer to **string** | Resolution of the camera. Can be one of &#39;1280x720&#39;, &#39;1920x1080&#39;, &#39;1080x1080&#39; or &#39;2058x2058&#39;. Not all resolutions are supported by every camera model. | [optional] 
**MotionDetectorVersion** | Pointer to **int32** | The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2. | [optional] 

## Methods

### NewUpdateDeviceCameraQualityAndRetentionRequest

`func NewUpdateDeviceCameraQualityAndRetentionRequest() *UpdateDeviceCameraQualityAndRetentionRequest`

NewUpdateDeviceCameraQualityAndRetentionRequest instantiates a new UpdateDeviceCameraQualityAndRetentionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCameraQualityAndRetentionRequestWithDefaults

`func NewUpdateDeviceCameraQualityAndRetentionRequestWithDefaults() *UpdateDeviceCameraQualityAndRetentionRequest`

NewUpdateDeviceCameraQualityAndRetentionRequestWithDefaults instantiates a new UpdateDeviceCameraQualityAndRetentionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetMotionBasedRetentionEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetMotionBasedRetentionEnabled() bool`

GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field if non-nil, zero value otherwise.

### GetMotionBasedRetentionEnabledOk

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetMotionBasedRetentionEnabledOk() (*bool, bool)`

GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionBasedRetentionEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) SetMotionBasedRetentionEnabled(v bool)`

SetMotionBasedRetentionEnabled sets MotionBasedRetentionEnabled field to given value.

### HasMotionBasedRetentionEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) HasMotionBasedRetentionEnabled() bool`

HasMotionBasedRetentionEnabled returns a boolean if a field has been set.

### GetAudioRecordingEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetAudioRecordingEnabled() bool`

GetAudioRecordingEnabled returns the AudioRecordingEnabled field if non-nil, zero value otherwise.

### GetAudioRecordingEnabledOk

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetAudioRecordingEnabledOk() (*bool, bool)`

GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioRecordingEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) SetAudioRecordingEnabled(v bool)`

SetAudioRecordingEnabled sets AudioRecordingEnabled field to given value.

### HasAudioRecordingEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) HasAudioRecordingEnabled() bool`

HasAudioRecordingEnabled returns a boolean if a field has been set.

### GetRestrictedBandwidthModeEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetRestrictedBandwidthModeEnabled() bool`

GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field if non-nil, zero value otherwise.

### GetRestrictedBandwidthModeEnabledOk

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetRestrictedBandwidthModeEnabledOk() (*bool, bool)`

GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedBandwidthModeEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) SetRestrictedBandwidthModeEnabled(v bool)`

SetRestrictedBandwidthModeEnabled sets RestrictedBandwidthModeEnabled field to given value.

### HasRestrictedBandwidthModeEnabled

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) HasRestrictedBandwidthModeEnabled() bool`

HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.

### GetQuality

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetResolution

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetMotionDetectorVersion

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetMotionDetectorVersion() int32`

GetMotionDetectorVersion returns the MotionDetectorVersion field if non-nil, zero value otherwise.

### GetMotionDetectorVersionOk

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) GetMotionDetectorVersionOk() (*int32, bool)`

GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionDetectorVersion

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) SetMotionDetectorVersion(v int32)`

SetMotionDetectorVersion sets MotionDetectorVersion field to given value.

### HasMotionDetectorVersion

`func (o *UpdateDeviceCameraQualityAndRetentionRequest) HasMotionDetectorVersion() bool`

HasMotionDetectorVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


