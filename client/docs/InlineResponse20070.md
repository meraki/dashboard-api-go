# InlineResponse20070

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**ClientCount** | Pointer to **int32** | Number of connected clients | [optional] 

## Methods

### NewInlineResponse20070

`func NewInlineResponse20070() *InlineResponse20070`

NewInlineResponse20070 instantiates a new InlineResponse20070 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20070WithDefaults

`func NewInlineResponse20070WithDefaults() *InlineResponse20070`

NewInlineResponse20070WithDefaults instantiates a new InlineResponse20070 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse20070) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse20070) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse20070) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse20070) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse20070) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse20070) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse20070) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse20070) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetClientCount

`func (o *InlineResponse20070) GetClientCount() int32`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *InlineResponse20070) GetClientCountOk() (*int32, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *InlineResponse20070) SetClientCount(v int32)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *InlineResponse20070) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


