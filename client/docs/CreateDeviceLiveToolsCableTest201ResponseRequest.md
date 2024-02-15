# CreateDeviceLiveToolsCableTest201ResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**Ports** | Pointer to **[]string** | A list of ports for which to perform the cable test. | [optional] 

## Methods

### NewCreateDeviceLiveToolsCableTest201ResponseRequest

`func NewCreateDeviceLiveToolsCableTest201ResponseRequest() *CreateDeviceLiveToolsCableTest201ResponseRequest`

NewCreateDeviceLiveToolsCableTest201ResponseRequest instantiates a new CreateDeviceLiveToolsCableTest201ResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceLiveToolsCableTest201ResponseRequestWithDefaults

`func NewCreateDeviceLiveToolsCableTest201ResponseRequestWithDefaults() *CreateDeviceLiveToolsCableTest201ResponseRequest`

NewCreateDeviceLiveToolsCableTest201ResponseRequestWithDefaults instantiates a new CreateDeviceLiveToolsCableTest201ResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetPorts

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


