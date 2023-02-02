# InlineResponse20074

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**AvgLatencyMs** | Pointer to **int32** | Average latency in milliseconds | [optional] 

## Methods

### NewInlineResponse20074

`func NewInlineResponse20074() *InlineResponse20074`

NewInlineResponse20074 instantiates a new InlineResponse20074 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074WithDefaults

`func NewInlineResponse20074WithDefaults() *InlineResponse20074`

NewInlineResponse20074WithDefaults instantiates a new InlineResponse20074 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse20074) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse20074) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse20074) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse20074) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse20074) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse20074) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse20074) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse20074) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetAvgLatencyMs

`func (o *InlineResponse20074) GetAvgLatencyMs() int32`

GetAvgLatencyMs returns the AvgLatencyMs field if non-nil, zero value otherwise.

### GetAvgLatencyMsOk

`func (o *InlineResponse20074) GetAvgLatencyMsOk() (*int32, bool)`

GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatencyMs

`func (o *InlineResponse20074) SetAvgLatencyMs(v int32)`

SetAvgLatencyMs sets AvgLatencyMs field to given value.

### HasAvgLatencyMs

`func (o *InlineResponse20074) HasAvgLatencyMs() bool`

HasAvgLatencyMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


