# NetworksNetworkIdGroupPoliciesScheduling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed. | [optional] 
**Monday** | Pointer to [**NetworksNetworkIdGroupPoliciesSchedulingMonday**](NetworksNetworkIdGroupPoliciesSchedulingMonday.md) |  | [optional] 
**Tuesday** | Pointer to [**NetworksNetworkIdGroupPoliciesSchedulingTuesday**](NetworksNetworkIdGroupPoliciesSchedulingTuesday.md) |  | [optional] 
**Wednesday** | Pointer to [**NetworksNetworkIdGroupPoliciesSchedulingWednesday**](NetworksNetworkIdGroupPoliciesSchedulingWednesday.md) |  | [optional] 
**Thursday** | Pointer to [**NetworksNetworkIdGroupPoliciesSchedulingThursday**](NetworksNetworkIdGroupPoliciesSchedulingThursday.md) |  | [optional] 
**Friday** | Pointer to [**NetworksNetworkIdGroupPoliciesSchedulingFriday**](NetworksNetworkIdGroupPoliciesSchedulingFriday.md) |  | [optional] 
**Saturday** | Pointer to [**NetworksNetworkIdGroupPoliciesSchedulingSaturday**](NetworksNetworkIdGroupPoliciesSchedulingSaturday.md) |  | [optional] 
**Sunday** | Pointer to [**NetworksNetworkIdGroupPoliciesSchedulingSunday**](NetworksNetworkIdGroupPoliciesSchedulingSunday.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdGroupPoliciesScheduling

`func NewNetworksNetworkIdGroupPoliciesScheduling() *NetworksNetworkIdGroupPoliciesScheduling`

NewNetworksNetworkIdGroupPoliciesScheduling instantiates a new NetworksNetworkIdGroupPoliciesScheduling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdGroupPoliciesSchedulingWithDefaults

`func NewNetworksNetworkIdGroupPoliciesSchedulingWithDefaults() *NetworksNetworkIdGroupPoliciesScheduling`

NewNetworksNetworkIdGroupPoliciesSchedulingWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesScheduling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMonday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetMonday() NetworksNetworkIdGroupPoliciesSchedulingMonday`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetMondayOk() (*NetworksNetworkIdGroupPoliciesSchedulingMonday, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetMonday(v NetworksNetworkIdGroupPoliciesSchedulingMonday)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetTuesday() NetworksNetworkIdGroupPoliciesSchedulingTuesday`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetTuesdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingTuesday, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetTuesday(v NetworksNetworkIdGroupPoliciesSchedulingTuesday)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetWednesday() NetworksNetworkIdGroupPoliciesSchedulingWednesday`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetWednesdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingWednesday, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetWednesday(v NetworksNetworkIdGroupPoliciesSchedulingWednesday)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetThursday() NetworksNetworkIdGroupPoliciesSchedulingThursday`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetThursdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingThursday, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetThursday(v NetworksNetworkIdGroupPoliciesSchedulingThursday)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetFriday() NetworksNetworkIdGroupPoliciesSchedulingFriday`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetFridayOk() (*NetworksNetworkIdGroupPoliciesSchedulingFriday, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetFriday(v NetworksNetworkIdGroupPoliciesSchedulingFriday)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetSaturday() NetworksNetworkIdGroupPoliciesSchedulingSaturday`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetSaturdayOk() (*NetworksNetworkIdGroupPoliciesSchedulingSaturday, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetSaturday(v NetworksNetworkIdGroupPoliciesSchedulingSaturday)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetSunday() NetworksNetworkIdGroupPoliciesSchedulingSunday`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *NetworksNetworkIdGroupPoliciesScheduling) GetSundayOk() (*NetworksNetworkIdGroupPoliciesSchedulingSunday, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) SetSunday(v NetworksNetworkIdGroupPoliciesSchedulingSunday)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *NetworksNetworkIdGroupPoliciesScheduling) HasSunday() bool`

HasSunday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


