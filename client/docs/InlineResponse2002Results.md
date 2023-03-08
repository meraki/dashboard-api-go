# InlineResponse2002Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** | Number of packets sent | [optional] 
**Received** | Pointer to **int32** | Number of packets received | [optional] 
**Loss** | Pointer to [**InlineResponse2002ResultsLoss**](InlineResponse2002ResultsLoss.md) |  | [optional] 
**Latencies** | Pointer to [**InlineResponse2002ResultsLatencies**](InlineResponse2002ResultsLatencies.md) |  | [optional] 
**Replies** | Pointer to [**[]InlineResponse2002ResultsReplies**](InlineResponse2002ResultsReplies.md) | Received packets | [optional] 

## Methods

### NewInlineResponse2002Results

`func NewInlineResponse2002Results() *InlineResponse2002Results`

NewInlineResponse2002Results instantiates a new InlineResponse2002Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002ResultsWithDefaults

`func NewInlineResponse2002ResultsWithDefaults() *InlineResponse2002Results`

NewInlineResponse2002ResultsWithDefaults instantiates a new InlineResponse2002Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *InlineResponse2002Results) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse2002Results) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse2002Results) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse2002Results) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetReceived

`func (o *InlineResponse2002Results) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *InlineResponse2002Results) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *InlineResponse2002Results) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *InlineResponse2002Results) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetLoss

`func (o *InlineResponse2002Results) GetLoss() InlineResponse2002ResultsLoss`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *InlineResponse2002Results) GetLossOk() (*InlineResponse2002ResultsLoss, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *InlineResponse2002Results) SetLoss(v InlineResponse2002ResultsLoss)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *InlineResponse2002Results) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLatencies

`func (o *InlineResponse2002Results) GetLatencies() InlineResponse2002ResultsLatencies`

GetLatencies returns the Latencies field if non-nil, zero value otherwise.

### GetLatenciesOk

`func (o *InlineResponse2002Results) GetLatenciesOk() (*InlineResponse2002ResultsLatencies, bool)`

GetLatenciesOk returns a tuple with the Latencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencies

`func (o *InlineResponse2002Results) SetLatencies(v InlineResponse2002ResultsLatencies)`

SetLatencies sets Latencies field to given value.

### HasLatencies

`func (o *InlineResponse2002Results) HasLatencies() bool`

HasLatencies returns a boolean if a field has been set.

### GetReplies

`func (o *InlineResponse2002Results) GetReplies() []InlineResponse2002ResultsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *InlineResponse2002Results) GetRepliesOk() (*[]InlineResponse2002ResultsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *InlineResponse2002Results) SetReplies(v []InlineResponse2002ResultsReplies)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *InlineResponse2002Results) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


