# UpdateDeviceCellularSimsRequestSimFailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Failover to secondary SIM (optional) | [optional] 
**Timeout** | Pointer to **int32** | Failover timeout in seconds (optional) | [optional] 

## Methods

### NewUpdateDeviceCellularSimsRequestSimFailover

`func NewUpdateDeviceCellularSimsRequestSimFailover() *UpdateDeviceCellularSimsRequestSimFailover`

NewUpdateDeviceCellularSimsRequestSimFailover instantiates a new UpdateDeviceCellularSimsRequestSimFailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCellularSimsRequestSimFailoverWithDefaults

`func NewUpdateDeviceCellularSimsRequestSimFailoverWithDefaults() *UpdateDeviceCellularSimsRequestSimFailover`

NewUpdateDeviceCellularSimsRequestSimFailoverWithDefaults instantiates a new UpdateDeviceCellularSimsRequestSimFailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateDeviceCellularSimsRequestSimFailover) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDeviceCellularSimsRequestSimFailover) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDeviceCellularSimsRequestSimFailover) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateDeviceCellularSimsRequestSimFailover) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *UpdateDeviceCellularSimsRequestSimFailover) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateDeviceCellularSimsRequestSimFailover) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateDeviceCellularSimsRequestSimFailover) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateDeviceCellularSimsRequestSimFailover) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


