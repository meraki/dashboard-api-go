# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable custom analytics | [optional] 
**ArtifactId** | Pointer to **string** | The ID of the custom analytics artifact | [optional] 
**Parameters** | Pointer to [**[]DevicesSerialCameraCustomAnalyticsParameters**](DevicesSerialCameraCustomAnalyticsParameters.md) | Parameters for the custom analytics workload | [optional] 

## Methods

### NewInlineObject4

`func NewInlineObject4() *InlineObject4`

NewInlineObject4 instantiates a new InlineObject4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject4WithDefaults

`func NewInlineObject4WithDefaults() *InlineObject4`

NewInlineObject4WithDefaults instantiates a new InlineObject4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject4) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetArtifactId

`func (o *InlineObject4) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *InlineObject4) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *InlineObject4) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *InlineObject4) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetParameters

`func (o *InlineObject4) GetParameters() []DevicesSerialCameraCustomAnalyticsParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *InlineObject4) GetParametersOk() (*[]DevicesSerialCameraCustomAnalyticsParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *InlineObject4) SetParameters(v []DevicesSerialCameraCustomAnalyticsParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *InlineObject4) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


