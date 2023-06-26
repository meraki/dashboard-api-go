# GetNetworkWirelessRfProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the new profile. Must be unique. | [optional] 
**NetworkId** | Pointer to **string** | The network ID of the RF Profile | [optional] 
**Name** | Pointer to **string** | The name of the new profile. Must be unique. This param is required on creation. | [optional] 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | Pointer to **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | [optional] 
**ApBandSettings** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseApBandSettings**](GetNetworkWirelessRfProfiles200ResponseApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings**](GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings**](GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseSixGhzSettings**](GetNetworkWirelessRfProfiles200ResponseSixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseTransmission**](GetNetworkWirelessRfProfiles200ResponseTransmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**GetNetworkWirelessRfProfiles200ResponsePerSsidSettings**](GetNetworkWirelessRfProfiles200ResponsePerSsidSettings.md) |  | [optional] 

## Methods

### NewGetNetworkWirelessRfProfiles200Response

`func NewGetNetworkWirelessRfProfiles200Response() *GetNetworkWirelessRfProfiles200Response`

NewGetNetworkWirelessRfProfiles200Response instantiates a new GetNetworkWirelessRfProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessRfProfiles200ResponseWithDefaults

`func NewGetNetworkWirelessRfProfiles200ResponseWithDefaults() *GetNetworkWirelessRfProfiles200Response`

NewGetNetworkWirelessRfProfiles200ResponseWithDefaults instantiates a new GetNetworkWirelessRfProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkWirelessRfProfiles200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkWirelessRfProfiles200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkWirelessRfProfiles200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetNetworkWirelessRfProfiles200Response) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetNetworkWirelessRfProfiles200Response) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetNetworkWirelessRfProfiles200Response) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkWirelessRfProfiles200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkWirelessRfProfiles200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkWirelessRfProfiles200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientBalancingEnabled

`func (o *GetNetworkWirelessRfProfiles200Response) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *GetNetworkWirelessRfProfiles200Response) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *GetNetworkWirelessRfProfiles200Response) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *GetNetworkWirelessRfProfiles200Response) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *GetNetworkWirelessRfProfiles200Response) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *GetNetworkWirelessRfProfiles200Response) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *GetNetworkWirelessRfProfiles200Response) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *GetNetworkWirelessRfProfiles200Response) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.

### HasBandSelectionType

`func (o *GetNetworkWirelessRfProfiles200Response) HasBandSelectionType() bool`

HasBandSelectionType returns a boolean if a field has been set.

### GetApBandSettings

`func (o *GetNetworkWirelessRfProfiles200Response) GetApBandSettings() GetNetworkWirelessRfProfiles200ResponseApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetApBandSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *GetNetworkWirelessRfProfiles200Response) SetApBandSettings(v GetNetworkWirelessRfProfiles200ResponseApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *GetNetworkWirelessRfProfiles200Response) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) GetTwoFourGhzSettings() GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetTwoFourGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) SetTwoFourGhzSettings(v GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) GetFiveGhzSettings() GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetFiveGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) SetFiveGhzSettings(v GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) GetSixGhzSettings() GetNetworkWirelessRfProfiles200ResponseSixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetSixGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseSixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) SetSixGhzSettings(v GetNetworkWirelessRfProfiles200ResponseSixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *GetNetworkWirelessRfProfiles200Response) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *GetNetworkWirelessRfProfiles200Response) GetTransmission() GetNetworkWirelessRfProfiles200ResponseTransmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetTransmissionOk() (*GetNetworkWirelessRfProfiles200ResponseTransmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *GetNetworkWirelessRfProfiles200Response) SetTransmission(v GetNetworkWirelessRfProfiles200ResponseTransmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *GetNetworkWirelessRfProfiles200Response) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *GetNetworkWirelessRfProfiles200Response) GetPerSsidSettings() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *GetNetworkWirelessRfProfiles200Response) GetPerSsidSettingsOk() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *GetNetworkWirelessRfProfiles200Response) SetPerSsidSettings(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *GetNetworkWirelessRfProfiles200Response) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


