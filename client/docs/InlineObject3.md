# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable custom analytics | [optional] 
**ArtifactId** | Pointer to **string** | The ID of the custom analytics artifact | [optional] 
**Parameters** | Pointer to [**[]DevicesSerialCameraCustomAnalyticsParameters**](DevicesSerialCameraCustomAnalyticsParameters.md) | Parameters for the custom analytics workload | [optional] 

## Methods

### NewInlineObject3

`func NewInlineObject3() *InlineObject3`

NewInlineObject3 instantiates a new InlineObject3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject3WithDefaults

`func NewInlineObject3WithDefaults() *InlineObject3`

NewInlineObject3WithDefaults instantiates a new InlineObject3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject3) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetArtifactId

`func (o *InlineObject3) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *InlineObject3) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *InlineObject3) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *InlineObject3) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetParameters

`func (o *InlineObject3) GetParameters() []DevicesSerialCameraCustomAnalyticsParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *InlineObject3) GetParametersOk() (*[]DevicesSerialCameraCustomAnalyticsParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *InlineObject3) SetParameters(v []DevicesSerialCameraCustomAnalyticsParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *InlineObject3) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


