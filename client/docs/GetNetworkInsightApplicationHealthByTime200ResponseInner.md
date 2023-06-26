# GetNetworkInsightApplicationHealthByTime200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**WanGoodput** | Pointer to **int32** | WAN goodput (Number of useful information bits delivered over a WAN per unit of time) | [optional] 
**LanGoodput** | Pointer to **int32** | LAN goodput (Number of useful information bits delivered over a LAN per unit of time) | [optional] 
**WanLatencyMs** | Pointer to **float32** | WAN latency in milliseconds | [optional] 
**LanLatencyMs** | Pointer to **float32** | LAN latency in milliseconds | [optional] 
**WanLossPercent** | Pointer to **float32** | WAN loss percentage | [optional] 
**LanLossPercent** | Pointer to **float32** | LAN loss percentage | [optional] 
**ResponseDuration** | Pointer to **int32** | Duration of the response, in milliseconds | [optional] 
**Sent** | Pointer to **int32** | Sent kilobytes-per-second | [optional] 
**Recv** | Pointer to **int32** | Received kilobytes-per-second | [optional] 
**NumClients** | Pointer to **int32** | Number of clients | [optional] 

## Methods

### NewGetNetworkInsightApplicationHealthByTime200ResponseInner

`func NewGetNetworkInsightApplicationHealthByTime200ResponseInner() *GetNetworkInsightApplicationHealthByTime200ResponseInner`

NewGetNetworkInsightApplicationHealthByTime200ResponseInner instantiates a new GetNetworkInsightApplicationHealthByTime200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkInsightApplicationHealthByTime200ResponseInnerWithDefaults

`func NewGetNetworkInsightApplicationHealthByTime200ResponseInnerWithDefaults() *GetNetworkInsightApplicationHealthByTime200ResponseInner`

NewGetNetworkInsightApplicationHealthByTime200ResponseInnerWithDefaults instantiates a new GetNetworkInsightApplicationHealthByTime200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetWanGoodput

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetWanGoodput() int32`

GetWanGoodput returns the WanGoodput field if non-nil, zero value otherwise.

### GetWanGoodputOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetWanGoodputOk() (*int32, bool)`

GetWanGoodputOk returns a tuple with the WanGoodput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanGoodput

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetWanGoodput(v int32)`

SetWanGoodput sets WanGoodput field to given value.

### HasWanGoodput

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasWanGoodput() bool`

HasWanGoodput returns a boolean if a field has been set.

### GetLanGoodput

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetLanGoodput() int32`

GetLanGoodput returns the LanGoodput field if non-nil, zero value otherwise.

### GetLanGoodputOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetLanGoodputOk() (*int32, bool)`

GetLanGoodputOk returns a tuple with the LanGoodput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanGoodput

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetLanGoodput(v int32)`

SetLanGoodput sets LanGoodput field to given value.

### HasLanGoodput

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasLanGoodput() bool`

HasLanGoodput returns a boolean if a field has been set.

### GetWanLatencyMs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetWanLatencyMs() float32`

GetWanLatencyMs returns the WanLatencyMs field if non-nil, zero value otherwise.

### GetWanLatencyMsOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetWanLatencyMsOk() (*float32, bool)`

GetWanLatencyMsOk returns a tuple with the WanLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanLatencyMs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetWanLatencyMs(v float32)`

SetWanLatencyMs sets WanLatencyMs field to given value.

### HasWanLatencyMs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasWanLatencyMs() bool`

HasWanLatencyMs returns a boolean if a field has been set.

### GetLanLatencyMs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetLanLatencyMs() float32`

GetLanLatencyMs returns the LanLatencyMs field if non-nil, zero value otherwise.

### GetLanLatencyMsOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetLanLatencyMsOk() (*float32, bool)`

GetLanLatencyMsOk returns a tuple with the LanLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanLatencyMs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetLanLatencyMs(v float32)`

SetLanLatencyMs sets LanLatencyMs field to given value.

### HasLanLatencyMs

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasLanLatencyMs() bool`

HasLanLatencyMs returns a boolean if a field has been set.

### GetWanLossPercent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetWanLossPercent() float32`

GetWanLossPercent returns the WanLossPercent field if non-nil, zero value otherwise.

### GetWanLossPercentOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetWanLossPercentOk() (*float32, bool)`

GetWanLossPercentOk returns a tuple with the WanLossPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanLossPercent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetWanLossPercent(v float32)`

SetWanLossPercent sets WanLossPercent field to given value.

### HasWanLossPercent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasWanLossPercent() bool`

HasWanLossPercent returns a boolean if a field has been set.

### GetLanLossPercent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetLanLossPercent() float32`

GetLanLossPercent returns the LanLossPercent field if non-nil, zero value otherwise.

### GetLanLossPercentOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetLanLossPercentOk() (*float32, bool)`

GetLanLossPercentOk returns a tuple with the LanLossPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanLossPercent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetLanLossPercent(v float32)`

SetLanLossPercent sets LanLossPercent field to given value.

### HasLanLossPercent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasLanLossPercent() bool`

HasLanLossPercent returns a boolean if a field has been set.

### GetResponseDuration

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetResponseDuration() int32`

GetResponseDuration returns the ResponseDuration field if non-nil, zero value otherwise.

### GetResponseDurationOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetResponseDurationOk() (*int32, bool)`

GetResponseDurationOk returns a tuple with the ResponseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDuration

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetResponseDuration(v int32)`

SetResponseDuration sets ResponseDuration field to given value.

### HasResponseDuration

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasResponseDuration() bool`

HasResponseDuration returns a boolean if a field has been set.

### GetSent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetRecv

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetNumClients

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *GetNetworkInsightApplicationHealthByTime200ResponseInner) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


