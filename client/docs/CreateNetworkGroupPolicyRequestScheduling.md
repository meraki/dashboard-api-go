# CreateNetworkGroupPolicyRequestScheduling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed. | [optional] 
**Monday** | Pointer to [**CreateNetworkGroupPolicyRequestSchedulingMonday**](CreateNetworkGroupPolicyRequestSchedulingMonday.md) |  | [optional] 
**Tuesday** | Pointer to [**CreateNetworkGroupPolicyRequestSchedulingTuesday**](CreateNetworkGroupPolicyRequestSchedulingTuesday.md) |  | [optional] 
**Wednesday** | Pointer to [**CreateNetworkGroupPolicyRequestSchedulingWednesday**](CreateNetworkGroupPolicyRequestSchedulingWednesday.md) |  | [optional] 
**Thursday** | Pointer to [**CreateNetworkGroupPolicyRequestSchedulingThursday**](CreateNetworkGroupPolicyRequestSchedulingThursday.md) |  | [optional] 
**Friday** | Pointer to [**CreateNetworkGroupPolicyRequestSchedulingFriday**](CreateNetworkGroupPolicyRequestSchedulingFriday.md) |  | [optional] 
**Saturday** | Pointer to [**CreateNetworkGroupPolicyRequestSchedulingSaturday**](CreateNetworkGroupPolicyRequestSchedulingSaturday.md) |  | [optional] 
**Sunday** | Pointer to [**CreateNetworkGroupPolicyRequestSchedulingSunday**](CreateNetworkGroupPolicyRequestSchedulingSunday.md) |  | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestScheduling

`func NewCreateNetworkGroupPolicyRequestScheduling() *CreateNetworkGroupPolicyRequestScheduling`

NewCreateNetworkGroupPolicyRequestScheduling instantiates a new CreateNetworkGroupPolicyRequestScheduling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestSchedulingWithDefaults

`func NewCreateNetworkGroupPolicyRequestSchedulingWithDefaults() *CreateNetworkGroupPolicyRequestScheduling`

NewCreateNetworkGroupPolicyRequestSchedulingWithDefaults instantiates a new CreateNetworkGroupPolicyRequestScheduling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMonday

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetMonday() CreateNetworkGroupPolicyRequestSchedulingMonday`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetMondayOk() (*CreateNetworkGroupPolicyRequestSchedulingMonday, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetMonday(v CreateNetworkGroupPolicyRequestSchedulingMonday)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetTuesday() CreateNetworkGroupPolicyRequestSchedulingTuesday`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetTuesdayOk() (*CreateNetworkGroupPolicyRequestSchedulingTuesday, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetTuesday(v CreateNetworkGroupPolicyRequestSchedulingTuesday)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetWednesday() CreateNetworkGroupPolicyRequestSchedulingWednesday`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetWednesdayOk() (*CreateNetworkGroupPolicyRequestSchedulingWednesday, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetWednesday(v CreateNetworkGroupPolicyRequestSchedulingWednesday)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetThursday() CreateNetworkGroupPolicyRequestSchedulingThursday`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetThursdayOk() (*CreateNetworkGroupPolicyRequestSchedulingThursday, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetThursday(v CreateNetworkGroupPolicyRequestSchedulingThursday)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetFriday() CreateNetworkGroupPolicyRequestSchedulingFriday`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetFridayOk() (*CreateNetworkGroupPolicyRequestSchedulingFriday, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetFriday(v CreateNetworkGroupPolicyRequestSchedulingFriday)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetSaturday() CreateNetworkGroupPolicyRequestSchedulingSaturday`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetSaturdayOk() (*CreateNetworkGroupPolicyRequestSchedulingSaturday, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetSaturday(v CreateNetworkGroupPolicyRequestSchedulingSaturday)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetSunday() CreateNetworkGroupPolicyRequestSchedulingSunday`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *CreateNetworkGroupPolicyRequestScheduling) GetSundayOk() (*CreateNetworkGroupPolicyRequestSchedulingSunday, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *CreateNetworkGroupPolicyRequestScheduling) SetSunday(v CreateNetworkGroupPolicyRequestSchedulingSunday)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *CreateNetworkGroupPolicyRequestScheduling) HasSunday() bool`

HasSunday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


