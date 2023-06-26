# UpdateNetworkSwitchSettingsRequestPowerExceptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | Serial number of the switch | 
**PowerType** | **string** | Per switch exception (combined, redundant, useNetworkSetting) | 

## Methods

### NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInner

`func NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInner(serial string, powerType string, ) *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner`

NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInner instantiates a new UpdateNetworkSwitchSettingsRequestPowerExceptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInnerWithDefaults

`func NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInnerWithDefaults() *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner`

NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInnerWithDefaults instantiates a new UpdateNetworkSwitchSettingsRequestPowerExceptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetPowerType

`func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetPowerType() string`

GetPowerType returns the PowerType field if non-nil, zero value otherwise.

### GetPowerTypeOk

`func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetPowerTypeOk() (*string, bool)`

GetPowerTypeOk returns a tuple with the PowerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerType

`func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) SetPowerType(v string)`

SetPowerType sets PowerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


