# CreateDeviceLiveToolsCableTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | **[]string** | A list of ports for which to perform the cable test. | 
**Callback** | Pointer to [**CreateDeviceLiveToolsArpTableRequestCallback**](CreateDeviceLiveToolsArpTableRequestCallback.md) |  | [optional] 

## Methods

### NewCreateDeviceLiveToolsCableTestRequest

`func NewCreateDeviceLiveToolsCableTestRequest(ports []string, ) *CreateDeviceLiveToolsCableTestRequest`

NewCreateDeviceLiveToolsCableTestRequest instantiates a new CreateDeviceLiveToolsCableTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceLiveToolsCableTestRequestWithDefaults

`func NewCreateDeviceLiveToolsCableTestRequestWithDefaults() *CreateDeviceLiveToolsCableTestRequest`

NewCreateDeviceLiveToolsCableTestRequestWithDefaults instantiates a new CreateDeviceLiveToolsCableTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *CreateDeviceLiveToolsCableTestRequest) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CreateDeviceLiveToolsCableTestRequest) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CreateDeviceLiveToolsCableTestRequest) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetCallback

`func (o *CreateDeviceLiveToolsCableTestRequest) GetCallback() CreateDeviceLiveToolsArpTableRequestCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *CreateDeviceLiveToolsCableTestRequest) GetCallbackOk() (*CreateDeviceLiveToolsArpTableRequestCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *CreateDeviceLiveToolsCableTestRequest) SetCallback(v CreateDeviceLiveToolsArpTableRequestCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *CreateDeviceLiveToolsCableTestRequest) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


