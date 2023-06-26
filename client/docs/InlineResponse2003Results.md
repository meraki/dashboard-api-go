# InlineResponse2003Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** | Number of packets sent | [optional] 
**Received** | Pointer to **int32** | Number of packets received | [optional] 
**Loss** | Pointer to [**InlineResponse2003ResultsLoss**](InlineResponse2003ResultsLoss.md) |  | [optional] 
**Latencies** | Pointer to [**InlineResponse2003ResultsLatencies**](InlineResponse2003ResultsLatencies.md) |  | [optional] 
**Replies** | Pointer to [**[]InlineResponse2003ResultsReplies**](InlineResponse2003ResultsReplies.md) | Received packets | [optional] 

## Methods

### NewInlineResponse2003Results

`func NewInlineResponse2003Results() *InlineResponse2003Results`

NewInlineResponse2003Results instantiates a new InlineResponse2003Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003ResultsWithDefaults

`func NewInlineResponse2003ResultsWithDefaults() *InlineResponse2003Results`

NewInlineResponse2003ResultsWithDefaults instantiates a new InlineResponse2003Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *InlineResponse2003Results) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse2003Results) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse2003Results) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse2003Results) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetReceived

`func (o *InlineResponse2003Results) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *InlineResponse2003Results) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *InlineResponse2003Results) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *InlineResponse2003Results) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetLoss

`func (o *InlineResponse2003Results) GetLoss() InlineResponse2003ResultsLoss`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *InlineResponse2003Results) GetLossOk() (*InlineResponse2003ResultsLoss, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *InlineResponse2003Results) SetLoss(v InlineResponse2003ResultsLoss)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *InlineResponse2003Results) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLatencies

`func (o *InlineResponse2003Results) GetLatencies() InlineResponse2003ResultsLatencies`

GetLatencies returns the Latencies field if non-nil, zero value otherwise.

### GetLatenciesOk

`func (o *InlineResponse2003Results) GetLatenciesOk() (*InlineResponse2003ResultsLatencies, bool)`

GetLatenciesOk returns a tuple with the Latencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencies

`func (o *InlineResponse2003Results) SetLatencies(v InlineResponse2003ResultsLatencies)`

SetLatencies sets Latencies field to given value.

### HasLatencies

`func (o *InlineResponse2003Results) HasLatencies() bool`

HasLatencies returns a boolean if a field has been set.

### GetReplies

`func (o *InlineResponse2003Results) GetReplies() []InlineResponse2003ResultsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *InlineResponse2003Results) GetRepliesOk() (*[]InlineResponse2003ResultsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *InlineResponse2003Results) SetReplies(v []InlineResponse2003ResultsReplies)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *InlineResponse2003Results) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


