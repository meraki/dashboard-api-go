# UpdateNetworkApplianceRfProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**TwoFourGhzSettings** | Pointer to [**UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings**](UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**UpdateNetworkApplianceRfProfileRequestFiveGhzSettings**](UpdateNetworkApplianceRfProfileRequestFiveGhzSettings.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**CreateNetworkApplianceRfProfileRequestPerSsidSettings**](CreateNetworkApplianceRfProfileRequestPerSsidSettings.md) |  | [optional] 

## Methods

### NewUpdateNetworkApplianceRfProfileRequest

`func NewUpdateNetworkApplianceRfProfileRequest() *UpdateNetworkApplianceRfProfileRequest`

NewUpdateNetworkApplianceRfProfileRequest instantiates a new UpdateNetworkApplianceRfProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceRfProfileRequestWithDefaults

`func NewUpdateNetworkApplianceRfProfileRequestWithDefaults() *UpdateNetworkApplianceRfProfileRequest`

NewUpdateNetworkApplianceRfProfileRequestWithDefaults instantiates a new UpdateNetworkApplianceRfProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkApplianceRfProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkApplianceRfProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkApplianceRfProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkApplianceRfProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) GetTwoFourGhzSettings() UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *UpdateNetworkApplianceRfProfileRequest) GetTwoFourGhzSettingsOk() (*UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) SetTwoFourGhzSettings(v UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) GetFiveGhzSettings() UpdateNetworkApplianceRfProfileRequestFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *UpdateNetworkApplianceRfProfileRequest) GetFiveGhzSettingsOk() (*UpdateNetworkApplianceRfProfileRequestFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) SetFiveGhzSettings(v UpdateNetworkApplianceRfProfileRequestFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) GetPerSsidSettings() CreateNetworkApplianceRfProfileRequestPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *UpdateNetworkApplianceRfProfileRequest) GetPerSsidSettingsOk() (*CreateNetworkApplianceRfProfileRequestPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) SetPerSsidSettings(v CreateNetworkApplianceRfProfileRequestPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *UpdateNetworkApplianceRfProfileRequest) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


