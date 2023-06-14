# CreateNetworkWirelessRfProfileRequestPerSsidSettings8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinBitrate** | Pointer to **float32** | Sets min bitrate (Mbps) of this SSID. Can be one of &#39;1&#39;, &#39;2&#39;, &#39;5.5&#39;, &#39;6&#39;, &#39;9&#39;, &#39;11&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. | [optional] 
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39; or &#39;5ghz&#39;. | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false. | [optional] 

## Methods

### NewCreateNetworkWirelessRfProfileRequestPerSsidSettings8

`func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings8() *CreateNetworkWirelessRfProfileRequestPerSsidSettings8`

NewCreateNetworkWirelessRfProfileRequestPerSsidSettings8 instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessRfProfileRequestPerSsidSettings8WithDefaults

`func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings8WithDefaults() *CreateNetworkWirelessRfProfileRequestPerSsidSettings8`

NewCreateNetworkWirelessRfProfileRequestPerSsidSettings8WithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinBitrate

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) GetMinBitrate() float32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) GetMinBitrateOk() (*float32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) SetMinBitrate(v float32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetBandOperationMode

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings8) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


