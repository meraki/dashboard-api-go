# GetDeviceLiveToolsPing200ResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** | Number of packets sent | [optional] 
**Received** | Pointer to **int32** | Number of packets received | [optional] 
**Loss** | Pointer to [**GetDeviceLiveToolsPing200ResponseResultsLoss**](GetDeviceLiveToolsPing200ResponseResultsLoss.md) |  | [optional] 
**Latencies** | Pointer to [**GetDeviceLiveToolsPing200ResponseResultsLatencies**](GetDeviceLiveToolsPing200ResponseResultsLatencies.md) |  | [optional] 
**Replies** | Pointer to [**[]GetDeviceLiveToolsPing200ResponseResultsRepliesInner**](GetDeviceLiveToolsPing200ResponseResultsRepliesInner.md) | Received packets | [optional] 

## Methods

### NewGetDeviceLiveToolsPing200ResponseResults

`func NewGetDeviceLiveToolsPing200ResponseResults() *GetDeviceLiveToolsPing200ResponseResults`

NewGetDeviceLiveToolsPing200ResponseResults instantiates a new GetDeviceLiveToolsPing200ResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceLiveToolsPing200ResponseResultsWithDefaults

`func NewGetDeviceLiveToolsPing200ResponseResultsWithDefaults() *GetDeviceLiveToolsPing200ResponseResults`

NewGetDeviceLiveToolsPing200ResponseResultsWithDefaults instantiates a new GetDeviceLiveToolsPing200ResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetDeviceLiveToolsPing200ResponseResults) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetDeviceLiveToolsPing200ResponseResults) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetReceived

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GetDeviceLiveToolsPing200ResponseResults) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *GetDeviceLiveToolsPing200ResponseResults) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetLoss

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetLoss() GetDeviceLiveToolsPing200ResponseResultsLoss`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetLossOk() (*GetDeviceLiveToolsPing200ResponseResultsLoss, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *GetDeviceLiveToolsPing200ResponseResults) SetLoss(v GetDeviceLiveToolsPing200ResponseResultsLoss)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *GetDeviceLiveToolsPing200ResponseResults) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLatencies

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetLatencies() GetDeviceLiveToolsPing200ResponseResultsLatencies`

GetLatencies returns the Latencies field if non-nil, zero value otherwise.

### GetLatenciesOk

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetLatenciesOk() (*GetDeviceLiveToolsPing200ResponseResultsLatencies, bool)`

GetLatenciesOk returns a tuple with the Latencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencies

`func (o *GetDeviceLiveToolsPing200ResponseResults) SetLatencies(v GetDeviceLiveToolsPing200ResponseResultsLatencies)`

SetLatencies sets Latencies field to given value.

### HasLatencies

`func (o *GetDeviceLiveToolsPing200ResponseResults) HasLatencies() bool`

HasLatencies returns a boolean if a field has been set.

### GetReplies

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetReplies() []GetDeviceLiveToolsPing200ResponseResultsRepliesInner`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *GetDeviceLiveToolsPing200ResponseResults) GetRepliesOk() (*[]GetDeviceLiveToolsPing200ResponseResultsRepliesInner, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *GetDeviceLiveToolsPing200ResponseResults) SetReplies(v []GetDeviceLiveToolsPing200ResponseResultsRepliesInner)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *GetDeviceLiveToolsPing200ResponseResults) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


