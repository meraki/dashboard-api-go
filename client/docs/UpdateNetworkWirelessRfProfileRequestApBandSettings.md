# UpdateNetworkWirelessRfProfileRequestApBandSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39; or &#39;5ghz&#39;. | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band. Can be either true or false. | [optional] 

## Methods

### NewUpdateNetworkWirelessRfProfileRequestApBandSettings

`func NewUpdateNetworkWirelessRfProfileRequestApBandSettings() *UpdateNetworkWirelessRfProfileRequestApBandSettings`

NewUpdateNetworkWirelessRfProfileRequestApBandSettings instantiates a new UpdateNetworkWirelessRfProfileRequestApBandSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessRfProfileRequestApBandSettingsWithDefaults

`func NewUpdateNetworkWirelessRfProfileRequestApBandSettingsWithDefaults() *UpdateNetworkWirelessRfProfileRequestApBandSettings`

NewUpdateNetworkWirelessRfProfileRequestApBandSettingsWithDefaults instantiates a new UpdateNetworkWirelessRfProfileRequestApBandSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandOperationMode

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


