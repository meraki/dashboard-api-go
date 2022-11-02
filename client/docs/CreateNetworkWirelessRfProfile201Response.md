# CreateNetworkWirelessRfProfile201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**NetworkId** | Pointer to **string** | The network ID of the RF Profile | [optional] 
**Name** | Pointer to **string** | The name of the new profile. Must be unique. This param is required on creation. | [optional] 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | Pointer to **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | [optional] 
**ApBandSettings** | Pointer to [**CreateNetworkWirelessRfProfile201ResponseApBandSettings**](CreateNetworkWirelessRfProfile201ResponseApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**CreateNetworkWirelessRfProfileRequestTwoFourGhzSettings**](CreateNetworkWirelessRfProfileRequestTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**CreateNetworkWirelessRfProfileRequestFiveGhzSettings**](CreateNetworkWirelessRfProfileRequestFiveGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**CreateNetworkWirelessRfProfileRequestTransmission**](CreateNetworkWirelessRfProfileRequestTransmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**CreateNetworkWirelessRfProfile201ResponsePerSsidSettings**](CreateNetworkWirelessRfProfile201ResponsePerSsidSettings.md) |  | [optional] 

## Methods

### NewCreateNetworkWirelessRfProfile201Response

`func NewCreateNetworkWirelessRfProfile201Response() *CreateNetworkWirelessRfProfile201Response`

NewCreateNetworkWirelessRfProfile201Response instantiates a new CreateNetworkWirelessRfProfile201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessRfProfile201ResponseWithDefaults

`func NewCreateNetworkWirelessRfProfile201ResponseWithDefaults() *CreateNetworkWirelessRfProfile201Response`

NewCreateNetworkWirelessRfProfile201ResponseWithDefaults instantiates a new CreateNetworkWirelessRfProfile201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNetworkWirelessRfProfile201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkWirelessRfProfile201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkWirelessRfProfile201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *CreateNetworkWirelessRfProfile201Response) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreateNetworkWirelessRfProfile201Response) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CreateNetworkWirelessRfProfile201Response) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkWirelessRfProfile201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkWirelessRfProfile201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkWirelessRfProfile201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientBalancingEnabled

`func (o *CreateNetworkWirelessRfProfile201Response) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *CreateNetworkWirelessRfProfile201Response) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *CreateNetworkWirelessRfProfile201Response) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *CreateNetworkWirelessRfProfile201Response) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *CreateNetworkWirelessRfProfile201Response) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *CreateNetworkWirelessRfProfile201Response) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *CreateNetworkWirelessRfProfile201Response) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *CreateNetworkWirelessRfProfile201Response) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.

### HasBandSelectionType

`func (o *CreateNetworkWirelessRfProfile201Response) HasBandSelectionType() bool`

HasBandSelectionType returns a boolean if a field has been set.

### GetApBandSettings

`func (o *CreateNetworkWirelessRfProfile201Response) GetApBandSettings() CreateNetworkWirelessRfProfile201ResponseApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetApBandSettingsOk() (*CreateNetworkWirelessRfProfile201ResponseApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *CreateNetworkWirelessRfProfile201Response) SetApBandSettings(v CreateNetworkWirelessRfProfile201ResponseApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *CreateNetworkWirelessRfProfile201Response) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *CreateNetworkWirelessRfProfile201Response) GetTwoFourGhzSettings() CreateNetworkWirelessRfProfileRequestTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetTwoFourGhzSettingsOk() (*CreateNetworkWirelessRfProfileRequestTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *CreateNetworkWirelessRfProfile201Response) SetTwoFourGhzSettings(v CreateNetworkWirelessRfProfileRequestTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *CreateNetworkWirelessRfProfile201Response) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *CreateNetworkWirelessRfProfile201Response) GetFiveGhzSettings() CreateNetworkWirelessRfProfileRequestFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetFiveGhzSettingsOk() (*CreateNetworkWirelessRfProfileRequestFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *CreateNetworkWirelessRfProfile201Response) SetFiveGhzSettings(v CreateNetworkWirelessRfProfileRequestFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *CreateNetworkWirelessRfProfile201Response) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *CreateNetworkWirelessRfProfile201Response) GetTransmission() CreateNetworkWirelessRfProfileRequestTransmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetTransmissionOk() (*CreateNetworkWirelessRfProfileRequestTransmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *CreateNetworkWirelessRfProfile201Response) SetTransmission(v CreateNetworkWirelessRfProfileRequestTransmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *CreateNetworkWirelessRfProfile201Response) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *CreateNetworkWirelessRfProfile201Response) GetPerSsidSettings() CreateNetworkWirelessRfProfile201ResponsePerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *CreateNetworkWirelessRfProfile201Response) GetPerSsidSettingsOk() (*CreateNetworkWirelessRfProfile201ResponsePerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *CreateNetworkWirelessRfProfile201Response) SetPerSsidSettings(v CreateNetworkWirelessRfProfile201ResponsePerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *CreateNetworkWirelessRfProfile201Response) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


