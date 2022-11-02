# UpdateNetworkWirelessSsidSchedulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID outage schedule is enabled. | [optional] 
**Ranges** | Pointer to [**[]UpdateNetworkWirelessSsidSchedulesRequestRangesInner**](UpdateNetworkWirelessSsidSchedulesRequestRangesInner.md) | List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence. | [optional] 
**RangesInSeconds** | Pointer to [**[]UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner**](UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner.md) | List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidSchedulesRequest

`func NewUpdateNetworkWirelessSsidSchedulesRequest() *UpdateNetworkWirelessSsidSchedulesRequest`

NewUpdateNetworkWirelessSsidSchedulesRequest instantiates a new UpdateNetworkWirelessSsidSchedulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidSchedulesRequestWithDefaults

`func NewUpdateNetworkWirelessSsidSchedulesRequestWithDefaults() *UpdateNetworkWirelessSsidSchedulesRequest`

NewUpdateNetworkWirelessSsidSchedulesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidSchedulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRanges() []UpdateNetworkWirelessSsidSchedulesRequestRangesInner`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRangesOk() (*[]UpdateNetworkWirelessSsidSchedulesRequestRangesInner, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) SetRanges(v []UpdateNetworkWirelessSsidSchedulesRequestRangesInner)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRangesInSeconds

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRangesInSeconds() []UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner`

GetRangesInSeconds returns the RangesInSeconds field if non-nil, zero value otherwise.

### GetRangesInSecondsOk

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRangesInSecondsOk() (*[]UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner, bool)`

GetRangesInSecondsOk returns a tuple with the RangesInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangesInSeconds

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) SetRangesInSeconds(v []UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner)`

SetRangesInSeconds sets RangesInSeconds field to given value.

### HasRangesInSeconds

`func (o *UpdateNetworkWirelessSsidSchedulesRequest) HasRangesInSeconds() bool`

HasRangesInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


