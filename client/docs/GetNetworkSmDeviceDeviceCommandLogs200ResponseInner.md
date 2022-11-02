# GetNetworkSmDeviceDeviceCommandLogs200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The type of command sent to the device. | [optional] 
**Name** | Pointer to **string** | The name of the device to which the command is sent. | [optional] 
**Details** | Pointer to **string** | A JSON string object containing command details. | [optional] 
**DashboardUser** | Pointer to **string** | The Meraki dashboard user who initiated the command. | [optional] 
**Ts** | Pointer to **string** | The time the command was sent to the device. | [optional] 

## Methods

### NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInner

`func NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInner() *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner`

NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInner instantiates a new GetNetworkSmDeviceDeviceCommandLogs200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInnerWithDefaults

`func NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInnerWithDefaults() *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner`

NewGetNetworkSmDeviceDeviceCommandLogs200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceDeviceCommandLogs200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDetails

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDashboardUser

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDashboardUser() string`

GetDashboardUser returns the DashboardUser field if non-nil, zero value otherwise.

### GetDashboardUserOk

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetDashboardUserOk() (*string, bool)`

GetDashboardUserOk returns a tuple with the DashboardUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUser

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetDashboardUser(v string)`

SetDashboardUser sets DashboardUser field to given value.

### HasDashboardUser

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasDashboardUser() bool`

HasDashboardUser returns a boolean if a field has been set.

### GetTs

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetNetworkSmDeviceDeviceCommandLogs200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


