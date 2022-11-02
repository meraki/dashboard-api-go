# CreateDeviceLiveToolsPing201ResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**Target** | Pointer to **string** | IP address or FQDN to ping | [optional] 
**Count** | Pointer to **int32** | Number of pings to send | [optional] 

## Methods

### NewCreateDeviceLiveToolsPing201ResponseRequest

`func NewCreateDeviceLiveToolsPing201ResponseRequest() *CreateDeviceLiveToolsPing201ResponseRequest`

NewCreateDeviceLiveToolsPing201ResponseRequest instantiates a new CreateDeviceLiveToolsPing201ResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceLiveToolsPing201ResponseRequestWithDefaults

`func NewCreateDeviceLiveToolsPing201ResponseRequestWithDefaults() *CreateDeviceLiveToolsPing201ResponseRequest`

NewCreateDeviceLiveToolsPing201ResponseRequestWithDefaults instantiates a new CreateDeviceLiveToolsPing201ResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTarget

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCount

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateDeviceLiveToolsPing201ResponseRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


