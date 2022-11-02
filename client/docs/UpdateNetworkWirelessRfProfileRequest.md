# UpdateNetworkWirelessRfProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. | [optional] 
**BandSelectionType** | Pointer to **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. | [optional] 
**ApBandSettings** | Pointer to [**UpdateNetworkWirelessRfProfileRequestApBandSettings**](UpdateNetworkWirelessRfProfileRequestApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings**](UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**UpdateNetworkWirelessRfProfileRequestFiveGhzSettings**](UpdateNetworkWirelessRfProfileRequestFiveGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**CreateNetworkWirelessRfProfileRequestTransmission**](CreateNetworkWirelessRfProfileRequestTransmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**CreateNetworkWirelessRfProfileRequestPerSsidSettings**](CreateNetworkWirelessRfProfileRequestPerSsidSettings.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessRfProfileRequest

`func NewUpdateNetworkWirelessRfProfileRequest() *UpdateNetworkWirelessRfProfileRequest`

NewUpdateNetworkWirelessRfProfileRequest instantiates a new UpdateNetworkWirelessRfProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessRfProfileRequestWithDefaults

`func NewUpdateNetworkWirelessRfProfileRequestWithDefaults() *UpdateNetworkWirelessRfProfileRequest`

NewUpdateNetworkWirelessRfProfileRequestWithDefaults instantiates a new UpdateNetworkWirelessRfProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkWirelessRfProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkWirelessRfProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkWirelessRfProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientBalancingEnabled

`func (o *UpdateNetworkWirelessRfProfileRequest) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *UpdateNetworkWirelessRfProfileRequest) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *UpdateNetworkWirelessRfProfileRequest) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *UpdateNetworkWirelessRfProfileRequest) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *UpdateNetworkWirelessRfProfileRequest) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *UpdateNetworkWirelessRfProfileRequest) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *UpdateNetworkWirelessRfProfileRequest) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *UpdateNetworkWirelessRfProfileRequest) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.

### HasBandSelectionType

`func (o *UpdateNetworkWirelessRfProfileRequest) HasBandSelectionType() bool`

HasBandSelectionType returns a boolean if a field has been set.

### GetApBandSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) GetApBandSettings() UpdateNetworkWirelessRfProfileRequestApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetApBandSettingsOk() (*UpdateNetworkWirelessRfProfileRequestApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) SetApBandSettings(v UpdateNetworkWirelessRfProfileRequestApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettings() UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettingsOk() (*UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) SetTwoFourGhzSettings(v UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) GetFiveGhzSettings() UpdateNetworkWirelessRfProfileRequestFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetFiveGhzSettingsOk() (*UpdateNetworkWirelessRfProfileRequestFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) SetFiveGhzSettings(v UpdateNetworkWirelessRfProfileRequestFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *UpdateNetworkWirelessRfProfileRequest) GetTransmission() CreateNetworkWirelessRfProfileRequestTransmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetTransmissionOk() (*CreateNetworkWirelessRfProfileRequestTransmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *UpdateNetworkWirelessRfProfileRequest) SetTransmission(v CreateNetworkWirelessRfProfileRequestTransmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *UpdateNetworkWirelessRfProfileRequest) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) GetPerSsidSettings() CreateNetworkWirelessRfProfileRequestPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *UpdateNetworkWirelessRfProfileRequest) GetPerSsidSettingsOk() (*CreateNetworkWirelessRfProfileRequestPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) SetPerSsidSettings(v CreateNetworkWirelessRfProfileRequestPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *UpdateNetworkWirelessRfProfileRequest) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


