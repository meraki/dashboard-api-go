# UpdateDeviceCameraCustomAnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable custom analytics | [optional] 
**ArtifactId** | Pointer to **string** | The ID of the custom analytics artifact | [optional] 
**Parameters** | Pointer to [**[]UpdateDeviceCameraCustomAnalyticsRequestParametersInner**](UpdateDeviceCameraCustomAnalyticsRequestParametersInner.md) | Parameters for the custom analytics workload | [optional] 

## Methods

### NewUpdateDeviceCameraCustomAnalyticsRequest

`func NewUpdateDeviceCameraCustomAnalyticsRequest() *UpdateDeviceCameraCustomAnalyticsRequest`

NewUpdateDeviceCameraCustomAnalyticsRequest instantiates a new UpdateDeviceCameraCustomAnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCameraCustomAnalyticsRequestWithDefaults

`func NewUpdateDeviceCameraCustomAnalyticsRequestWithDefaults() *UpdateDeviceCameraCustomAnalyticsRequest`

NewUpdateDeviceCameraCustomAnalyticsRequestWithDefaults instantiates a new UpdateDeviceCameraCustomAnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetArtifactId

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetParameters() []UpdateDeviceCameraCustomAnalyticsRequestParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) GetParametersOk() (*[]UpdateDeviceCameraCustomAnalyticsRequestParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) SetParameters(v []UpdateDeviceCameraCustomAnalyticsRequestParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateDeviceCameraCustomAnalyticsRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


