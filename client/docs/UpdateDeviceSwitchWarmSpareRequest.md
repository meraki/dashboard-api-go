# UpdateDeviceSwitchWarmSpareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enable or disable warm spare for a switch | 
**SpareSerial** | Pointer to **string** | Serial number of the warm spare switch | [optional] 

## Methods

### NewUpdateDeviceSwitchWarmSpareRequest

`func NewUpdateDeviceSwitchWarmSpareRequest(enabled bool, ) *UpdateDeviceSwitchWarmSpareRequest`

NewUpdateDeviceSwitchWarmSpareRequest instantiates a new UpdateDeviceSwitchWarmSpareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceSwitchWarmSpareRequestWithDefaults

`func NewUpdateDeviceSwitchWarmSpareRequestWithDefaults() *UpdateDeviceSwitchWarmSpareRequest`

NewUpdateDeviceSwitchWarmSpareRequestWithDefaults instantiates a new UpdateDeviceSwitchWarmSpareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateDeviceSwitchWarmSpareRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDeviceSwitchWarmSpareRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDeviceSwitchWarmSpareRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSpareSerial

`func (o *UpdateDeviceSwitchWarmSpareRequest) GetSpareSerial() string`

GetSpareSerial returns the SpareSerial field if non-nil, zero value otherwise.

### GetSpareSerialOk

`func (o *UpdateDeviceSwitchWarmSpareRequest) GetSpareSerialOk() (*string, bool)`

GetSpareSerialOk returns a tuple with the SpareSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpareSerial

`func (o *UpdateDeviceSwitchWarmSpareRequest) SetSpareSerial(v string)`

SetSpareSerial sets SpareSerial field to given value.

### HasSpareSerial

`func (o *UpdateDeviceSwitchWarmSpareRequest) HasSpareSerial() bool`

HasSpareSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


