# CreateNetworkApplianceRfProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new profile. Must be unique. This param is required on creation. | 
**TwoFourGhzSettings** | Pointer to [**CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings**](CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**CreateNetworkApplianceRfProfileRequestFiveGhzSettings**](CreateNetworkApplianceRfProfileRequestFiveGhzSettings.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**CreateNetworkApplianceRfProfileRequestPerSsidSettings**](CreateNetworkApplianceRfProfileRequestPerSsidSettings.md) |  | [optional] 

## Methods

### NewCreateNetworkApplianceRfProfileRequest

`func NewCreateNetworkApplianceRfProfileRequest(name string, ) *CreateNetworkApplianceRfProfileRequest`

NewCreateNetworkApplianceRfProfileRequest instantiates a new CreateNetworkApplianceRfProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkApplianceRfProfileRequestWithDefaults

`func NewCreateNetworkApplianceRfProfileRequestWithDefaults() *CreateNetworkApplianceRfProfileRequest`

NewCreateNetworkApplianceRfProfileRequestWithDefaults instantiates a new CreateNetworkApplianceRfProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkApplianceRfProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkApplianceRfProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkApplianceRfProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTwoFourGhzSettings

`func (o *CreateNetworkApplianceRfProfileRequest) GetTwoFourGhzSettings() CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *CreateNetworkApplianceRfProfileRequest) GetTwoFourGhzSettingsOk() (*CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *CreateNetworkApplianceRfProfileRequest) SetTwoFourGhzSettings(v CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *CreateNetworkApplianceRfProfileRequest) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *CreateNetworkApplianceRfProfileRequest) GetFiveGhzSettings() CreateNetworkApplianceRfProfileRequestFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *CreateNetworkApplianceRfProfileRequest) GetFiveGhzSettingsOk() (*CreateNetworkApplianceRfProfileRequestFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *CreateNetworkApplianceRfProfileRequest) SetFiveGhzSettings(v CreateNetworkApplianceRfProfileRequestFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *CreateNetworkApplianceRfProfileRequest) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *CreateNetworkApplianceRfProfileRequest) GetPerSsidSettings() CreateNetworkApplianceRfProfileRequestPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *CreateNetworkApplianceRfProfileRequest) GetPerSsidSettingsOk() (*CreateNetworkApplianceRfProfileRequestPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *CreateNetworkApplianceRfProfileRequest) SetPerSsidSettings(v CreateNetworkApplianceRfProfileRequestPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *CreateNetworkApplianceRfProfileRequest) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


