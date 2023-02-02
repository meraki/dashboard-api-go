# InlineResponse2014

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**NetworkId** | Pointer to **string** | The network ID of the RF Profile | [optional] 
**Name** | Pointer to **string** | The name of the new profile. Must be unique. This param is required on creation. | [optional] 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | Pointer to **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | [optional] 
**ApBandSettings** | Pointer to [**InlineResponse2014ApBandSettings**](InlineResponse2014ApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings**](NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**NetworksNetworkIdWirelessRfProfilesFiveGhzSettings**](NetworksNetworkIdWirelessRfProfilesFiveGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**NetworksNetworkIdWirelessRfProfilesTransmission**](NetworksNetworkIdWirelessRfProfilesTransmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**InlineResponse2014PerSsidSettings**](InlineResponse2014PerSsidSettings.md) |  | [optional] 

## Methods

### NewInlineResponse2014

`func NewInlineResponse2014() *InlineResponse2014`

NewInlineResponse2014 instantiates a new InlineResponse2014 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2014WithDefaults

`func NewInlineResponse2014WithDefaults() *InlineResponse2014`

NewInlineResponse2014WithDefaults instantiates a new InlineResponse2014 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2014) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2014) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2014) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2014) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse2014) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse2014) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse2014) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse2014) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2014) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2014) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2014) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2014) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientBalancingEnabled

`func (o *InlineResponse2014) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *InlineResponse2014) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *InlineResponse2014) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *InlineResponse2014) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *InlineResponse2014) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *InlineResponse2014) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *InlineResponse2014) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *InlineResponse2014) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *InlineResponse2014) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *InlineResponse2014) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *InlineResponse2014) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.

### HasBandSelectionType

`func (o *InlineResponse2014) HasBandSelectionType() bool`

HasBandSelectionType returns a boolean if a field has been set.

### GetApBandSettings

`func (o *InlineResponse2014) GetApBandSettings() InlineResponse2014ApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *InlineResponse2014) GetApBandSettingsOk() (*InlineResponse2014ApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *InlineResponse2014) SetApBandSettings(v InlineResponse2014ApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *InlineResponse2014) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *InlineResponse2014) GetTwoFourGhzSettings() NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *InlineResponse2014) GetTwoFourGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *InlineResponse2014) SetTwoFourGhzSettings(v NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *InlineResponse2014) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *InlineResponse2014) GetFiveGhzSettings() NetworksNetworkIdWirelessRfProfilesFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *InlineResponse2014) GetFiveGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *InlineResponse2014) SetFiveGhzSettings(v NetworksNetworkIdWirelessRfProfilesFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *InlineResponse2014) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *InlineResponse2014) GetTransmission() NetworksNetworkIdWirelessRfProfilesTransmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *InlineResponse2014) GetTransmissionOk() (*NetworksNetworkIdWirelessRfProfilesTransmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *InlineResponse2014) SetTransmission(v NetworksNetworkIdWirelessRfProfilesTransmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *InlineResponse2014) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *InlineResponse2014) GetPerSsidSettings() InlineResponse2014PerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *InlineResponse2014) GetPerSsidSettingsOk() (*InlineResponse2014PerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *InlineResponse2014) SetPerSsidSettings(v InlineResponse2014PerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *InlineResponse2014) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


