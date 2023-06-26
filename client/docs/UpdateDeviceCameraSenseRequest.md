# UpdateDeviceCameraSenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenseEnabled** | Pointer to **bool** | Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera | [optional] 
**MqttBrokerId** | Pointer to **string** | The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera | [optional] 
**AudioDetection** | Pointer to [**UpdateDeviceCameraSenseRequestAudioDetection**](UpdateDeviceCameraSenseRequestAudioDetection.md) |  | [optional] 
**DetectionModelId** | Pointer to **string** | The ID of the object detection model | [optional] 

## Methods

### NewUpdateDeviceCameraSenseRequest

`func NewUpdateDeviceCameraSenseRequest() *UpdateDeviceCameraSenseRequest`

NewUpdateDeviceCameraSenseRequest instantiates a new UpdateDeviceCameraSenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCameraSenseRequestWithDefaults

`func NewUpdateDeviceCameraSenseRequestWithDefaults() *UpdateDeviceCameraSenseRequest`

NewUpdateDeviceCameraSenseRequestWithDefaults instantiates a new UpdateDeviceCameraSenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenseEnabled

`func (o *UpdateDeviceCameraSenseRequest) GetSenseEnabled() bool`

GetSenseEnabled returns the SenseEnabled field if non-nil, zero value otherwise.

### GetSenseEnabledOk

`func (o *UpdateDeviceCameraSenseRequest) GetSenseEnabledOk() (*bool, bool)`

GetSenseEnabledOk returns a tuple with the SenseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenseEnabled

`func (o *UpdateDeviceCameraSenseRequest) SetSenseEnabled(v bool)`

SetSenseEnabled sets SenseEnabled field to given value.

### HasSenseEnabled

`func (o *UpdateDeviceCameraSenseRequest) HasSenseEnabled() bool`

HasSenseEnabled returns a boolean if a field has been set.

### GetMqttBrokerId

`func (o *UpdateDeviceCameraSenseRequest) GetMqttBrokerId() string`

GetMqttBrokerId returns the MqttBrokerId field if non-nil, zero value otherwise.

### GetMqttBrokerIdOk

`func (o *UpdateDeviceCameraSenseRequest) GetMqttBrokerIdOk() (*string, bool)`

GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttBrokerId

`func (o *UpdateDeviceCameraSenseRequest) SetMqttBrokerId(v string)`

SetMqttBrokerId sets MqttBrokerId field to given value.

### HasMqttBrokerId

`func (o *UpdateDeviceCameraSenseRequest) HasMqttBrokerId() bool`

HasMqttBrokerId returns a boolean if a field has been set.

### GetAudioDetection

`func (o *UpdateDeviceCameraSenseRequest) GetAudioDetection() UpdateDeviceCameraSenseRequestAudioDetection`

GetAudioDetection returns the AudioDetection field if non-nil, zero value otherwise.

### GetAudioDetectionOk

`func (o *UpdateDeviceCameraSenseRequest) GetAudioDetectionOk() (*UpdateDeviceCameraSenseRequestAudioDetection, bool)`

GetAudioDetectionOk returns a tuple with the AudioDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioDetection

`func (o *UpdateDeviceCameraSenseRequest) SetAudioDetection(v UpdateDeviceCameraSenseRequestAudioDetection)`

SetAudioDetection sets AudioDetection field to given value.

### HasAudioDetection

`func (o *UpdateDeviceCameraSenseRequest) HasAudioDetection() bool`

HasAudioDetection returns a boolean if a field has been set.

### GetDetectionModelId

`func (o *UpdateDeviceCameraSenseRequest) GetDetectionModelId() string`

GetDetectionModelId returns the DetectionModelId field if non-nil, zero value otherwise.

### GetDetectionModelIdOk

`func (o *UpdateDeviceCameraSenseRequest) GetDetectionModelIdOk() (*string, bool)`

GetDetectionModelIdOk returns a tuple with the DetectionModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionModelId

`func (o *UpdateDeviceCameraSenseRequest) SetDetectionModelId(v string)`

SetDetectionModelId sets DetectionModelId field to given value.

### HasDetectionModelId

`func (o *UpdateDeviceCameraSenseRequest) HasDetectionModelId() bool`

HasDetectionModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


