# CreateNetworkWirelessRfProfile201ResponseApBandSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39; or &#39;5ghz&#39;. Defaults to dual. | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band. Can be either true or false. Defaults to true. | [optional] 

## Methods

### NewCreateNetworkWirelessRfProfile201ResponseApBandSettings

`func NewCreateNetworkWirelessRfProfile201ResponseApBandSettings() *CreateNetworkWirelessRfProfile201ResponseApBandSettings`

NewCreateNetworkWirelessRfProfile201ResponseApBandSettings instantiates a new CreateNetworkWirelessRfProfile201ResponseApBandSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessRfProfile201ResponseApBandSettingsWithDefaults

`func NewCreateNetworkWirelessRfProfile201ResponseApBandSettingsWithDefaults() *CreateNetworkWirelessRfProfile201ResponseApBandSettings`

NewCreateNetworkWirelessRfProfile201ResponseApBandSettingsWithDefaults instantiates a new CreateNetworkWirelessRfProfile201ResponseApBandSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandOperationMode

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *CreateNetworkWirelessRfProfile201ResponseApBandSettings) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


