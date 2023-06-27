# GetNetworkWirelessLatencyHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**AvgLatencyMs** | Pointer to **int32** | Average latency in milliseconds | [optional] 

## Methods

### NewGetNetworkWirelessLatencyHistory200ResponseInner

`func NewGetNetworkWirelessLatencyHistory200ResponseInner() *GetNetworkWirelessLatencyHistory200ResponseInner`

NewGetNetworkWirelessLatencyHistory200ResponseInner instantiates a new GetNetworkWirelessLatencyHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessLatencyHistory200ResponseInnerWithDefaults

`func NewGetNetworkWirelessLatencyHistory200ResponseInnerWithDefaults() *GetNetworkWirelessLatencyHistory200ResponseInner`

NewGetNetworkWirelessLatencyHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessLatencyHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetAvgLatencyMs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetAvgLatencyMs() int32`

GetAvgLatencyMs returns the AvgLatencyMs field if non-nil, zero value otherwise.

### GetAvgLatencyMsOk

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) GetAvgLatencyMsOk() (*int32, bool)`

GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatencyMs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) SetAvgLatencyMs(v int32)`

SetAvgLatencyMs sets AvgLatencyMs field to given value.

### HasAvgLatencyMs

`func (o *GetNetworkWirelessLatencyHistory200ResponseInner) HasAvgLatencyMs() bool`

HasAvgLatencyMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


