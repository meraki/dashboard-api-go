# CreateNetworkSwitchPortScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for your port schedule. Required | 
**PortSchedule** | Pointer to [**CreateNetworkSwitchPortScheduleRequestPortSchedule**](CreateNetworkSwitchPortScheduleRequestPortSchedule.md) |  | [optional] 

## Methods

### NewCreateNetworkSwitchPortScheduleRequest

`func NewCreateNetworkSwitchPortScheduleRequest(name string, ) *CreateNetworkSwitchPortScheduleRequest`

NewCreateNetworkSwitchPortScheduleRequest instantiates a new CreateNetworkSwitchPortScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchPortScheduleRequestWithDefaults

`func NewCreateNetworkSwitchPortScheduleRequestWithDefaults() *CreateNetworkSwitchPortScheduleRequest`

NewCreateNetworkSwitchPortScheduleRequestWithDefaults instantiates a new CreateNetworkSwitchPortScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkSwitchPortScheduleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkSwitchPortScheduleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkSwitchPortScheduleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPortSchedule

`func (o *CreateNetworkSwitchPortScheduleRequest) GetPortSchedule() CreateNetworkSwitchPortScheduleRequestPortSchedule`

GetPortSchedule returns the PortSchedule field if non-nil, zero value otherwise.

### GetPortScheduleOk

`func (o *CreateNetworkSwitchPortScheduleRequest) GetPortScheduleOk() (*CreateNetworkSwitchPortScheduleRequestPortSchedule, bool)`

GetPortScheduleOk returns a tuple with the PortSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSchedule

`func (o *CreateNetworkSwitchPortScheduleRequest) SetPortSchedule(v CreateNetworkSwitchPortScheduleRequestPortSchedule)`

SetPortSchedule sets PortSchedule field to given value.

### HasPortSchedule

`func (o *CreateNetworkSwitchPortScheduleRequest) HasPortSchedule() bool`

HasPortSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


