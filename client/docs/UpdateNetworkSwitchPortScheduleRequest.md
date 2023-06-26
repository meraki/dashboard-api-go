# UpdateNetworkSwitchPortScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for your port schedule. | [optional] 
**PortSchedule** | Pointer to [**CreateNetworkSwitchPortScheduleRequestPortSchedule**](CreateNetworkSwitchPortScheduleRequestPortSchedule.md) |  | [optional] 

## Methods

### NewUpdateNetworkSwitchPortScheduleRequest

`func NewUpdateNetworkSwitchPortScheduleRequest() *UpdateNetworkSwitchPortScheduleRequest`

NewUpdateNetworkSwitchPortScheduleRequest instantiates a new UpdateNetworkSwitchPortScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchPortScheduleRequestWithDefaults

`func NewUpdateNetworkSwitchPortScheduleRequestWithDefaults() *UpdateNetworkSwitchPortScheduleRequest`

NewUpdateNetworkSwitchPortScheduleRequestWithDefaults instantiates a new UpdateNetworkSwitchPortScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkSwitchPortScheduleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkSwitchPortScheduleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkSwitchPortScheduleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkSwitchPortScheduleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortSchedule

`func (o *UpdateNetworkSwitchPortScheduleRequest) GetPortSchedule() CreateNetworkSwitchPortScheduleRequestPortSchedule`

GetPortSchedule returns the PortSchedule field if non-nil, zero value otherwise.

### GetPortScheduleOk

`func (o *UpdateNetworkSwitchPortScheduleRequest) GetPortScheduleOk() (*CreateNetworkSwitchPortScheduleRequestPortSchedule, bool)`

GetPortScheduleOk returns a tuple with the PortSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSchedule

`func (o *UpdateNetworkSwitchPortScheduleRequest) SetPortSchedule(v CreateNetworkSwitchPortScheduleRequestPortSchedule)`

SetPortSchedule sets PortSchedule field to given value.

### HasPortSchedule

`func (o *UpdateNetworkSwitchPortScheduleRequest) HasPortSchedule() bool`

HasPortSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


