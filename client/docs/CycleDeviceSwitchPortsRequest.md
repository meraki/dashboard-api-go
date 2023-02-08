# CycleDeviceSwitchPortsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | **[]string** | List of switch ports. Example: [1, 2-5, 1_MA-MOD-8X10G_1, 1_MA-MOD-8X10G_2-1_MA-MOD-8X10G_8] | 

## Methods

### NewCycleDeviceSwitchPortsRequest

`func NewCycleDeviceSwitchPortsRequest(ports []string, ) *CycleDeviceSwitchPortsRequest`

NewCycleDeviceSwitchPortsRequest instantiates a new CycleDeviceSwitchPortsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCycleDeviceSwitchPortsRequestWithDefaults

`func NewCycleDeviceSwitchPortsRequestWithDefaults() *CycleDeviceSwitchPortsRequest`

NewCycleDeviceSwitchPortsRequestWithDefaults instantiates a new CycleDeviceSwitchPortsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *CycleDeviceSwitchPortsRequest) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CycleDeviceSwitchPortsRequest) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CycleDeviceSwitchPortsRequest) SetPorts(v []string)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


