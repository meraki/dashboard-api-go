# InlineResponse20033

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

### NewInlineResponse20033

`func NewInlineResponse20033() *InlineResponse20033`

NewInlineResponse20033 instantiates a new InlineResponse20033 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20033WithDefaults

`func NewInlineResponse20033WithDefaults() *InlineResponse20033`

NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse20033) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse20033) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse20033) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse20033) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse20033) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse20033) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse20033) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse20033) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetWanGoodput

`func (o *InlineResponse20033) GetWanGoodput() int32`

GetWanGoodput returns the WanGoodput field if non-nil, zero value otherwise.

### GetWanGoodputOk

`func (o *InlineResponse20033) GetWanGoodputOk() (*int32, bool)`

GetWanGoodputOk returns a tuple with the WanGoodput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanGoodput

`func (o *InlineResponse20033) SetWanGoodput(v int32)`

SetWanGoodput sets WanGoodput field to given value.

### HasWanGoodput

`func (o *InlineResponse20033) HasWanGoodput() bool`

HasWanGoodput returns a boolean if a field has been set.

### GetLanGoodput

`func (o *InlineResponse20033) GetLanGoodput() int32`

GetLanGoodput returns the LanGoodput field if non-nil, zero value otherwise.

### GetLanGoodputOk

`func (o *InlineResponse20033) GetLanGoodputOk() (*int32, bool)`

GetLanGoodputOk returns a tuple with the LanGoodput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanGoodput

`func (o *InlineResponse20033) SetLanGoodput(v int32)`

SetLanGoodput sets LanGoodput field to given value.

### HasLanGoodput

`func (o *InlineResponse20033) HasLanGoodput() bool`

HasLanGoodput returns a boolean if a field has been set.

### GetWanLatencyMs

`func (o *InlineResponse20033) GetWanLatencyMs() float32`

GetWanLatencyMs returns the WanLatencyMs field if non-nil, zero value otherwise.

### GetWanLatencyMsOk

`func (o *InlineResponse20033) GetWanLatencyMsOk() (*float32, bool)`

GetWanLatencyMsOk returns a tuple with the WanLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanLatencyMs

`func (o *InlineResponse20033) SetWanLatencyMs(v float32)`

SetWanLatencyMs sets WanLatencyMs field to given value.

### HasWanLatencyMs

`func (o *InlineResponse20033) HasWanLatencyMs() bool`

HasWanLatencyMs returns a boolean if a field has been set.

### GetLanLatencyMs

`func (o *InlineResponse20033) GetLanLatencyMs() float32`

GetLanLatencyMs returns the LanLatencyMs field if non-nil, zero value otherwise.

### GetLanLatencyMsOk

`func (o *InlineResponse20033) GetLanLatencyMsOk() (*float32, bool)`

GetLanLatencyMsOk returns a tuple with the LanLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanLatencyMs

`func (o *InlineResponse20033) SetLanLatencyMs(v float32)`

SetLanLatencyMs sets LanLatencyMs field to given value.

### HasLanLatencyMs

`func (o *InlineResponse20033) HasLanLatencyMs() bool`

HasLanLatencyMs returns a boolean if a field has been set.

### GetWanLossPercent

`func (o *InlineResponse20033) GetWanLossPercent() float32`

GetWanLossPercent returns the WanLossPercent field if non-nil, zero value otherwise.

### GetWanLossPercentOk

`func (o *InlineResponse20033) GetWanLossPercentOk() (*float32, bool)`

GetWanLossPercentOk returns a tuple with the WanLossPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanLossPercent

`func (o *InlineResponse20033) SetWanLossPercent(v float32)`

SetWanLossPercent sets WanLossPercent field to given value.

### HasWanLossPercent

`func (o *InlineResponse20033) HasWanLossPercent() bool`

HasWanLossPercent returns a boolean if a field has been set.

### GetLanLossPercent

`func (o *InlineResponse20033) GetLanLossPercent() float32`

GetLanLossPercent returns the LanLossPercent field if non-nil, zero value otherwise.

### GetLanLossPercentOk

`func (o *InlineResponse20033) GetLanLossPercentOk() (*float32, bool)`

GetLanLossPercentOk returns a tuple with the LanLossPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanLossPercent

`func (o *InlineResponse20033) SetLanLossPercent(v float32)`

SetLanLossPercent sets LanLossPercent field to given value.

### HasLanLossPercent

`func (o *InlineResponse20033) HasLanLossPercent() bool`

HasLanLossPercent returns a boolean if a field has been set.

### GetResponseDuration

`func (o *InlineResponse20033) GetResponseDuration() int32`

GetResponseDuration returns the ResponseDuration field if non-nil, zero value otherwise.

### GetResponseDurationOk

`func (o *InlineResponse20033) GetResponseDurationOk() (*int32, bool)`

GetResponseDurationOk returns a tuple with the ResponseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDuration

`func (o *InlineResponse20033) SetResponseDuration(v int32)`

SetResponseDuration sets ResponseDuration field to given value.

### HasResponseDuration

`func (o *InlineResponse20033) HasResponseDuration() bool`

HasResponseDuration returns a boolean if a field has been set.

### GetSent

`func (o *InlineResponse20033) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse20033) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse20033) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse20033) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse20033) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse20033) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse20033) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse20033) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetNumClients

`func (o *InlineResponse20033) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *InlineResponse20033) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *InlineResponse20033) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *InlineResponse20033) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


