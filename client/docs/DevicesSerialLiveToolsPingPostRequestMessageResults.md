# DevicesSerialLiveToolsPingPostRequestMessageResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** | Number of packets sent | [optional] 
**Received** | Pointer to **int32** | Number of packets received | [optional] 
**Loss** | Pointer to [**DevicesSerialLiveToolsPingPostRequestMessageResultsLoss**](DevicesSerialLiveToolsPingPostRequestMessageResultsLoss.md) |  | [optional] 
**Latencies** | Pointer to [**DevicesSerialLiveToolsPingPostRequestMessageResultsLatencies**](DevicesSerialLiveToolsPingPostRequestMessageResultsLatencies.md) |  | [optional] 
**Replies** | Pointer to [**[]DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner**](DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner.md) | Received packets | [optional] 

## Methods

### NewDevicesSerialLiveToolsPingPostRequestMessageResults

`func NewDevicesSerialLiveToolsPingPostRequestMessageResults() *DevicesSerialLiveToolsPingPostRequestMessageResults`

NewDevicesSerialLiveToolsPingPostRequestMessageResults instantiates a new DevicesSerialLiveToolsPingPostRequestMessageResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsPingPostRequestMessageResultsWithDefaults

`func NewDevicesSerialLiveToolsPingPostRequestMessageResultsWithDefaults() *DevicesSerialLiveToolsPingPostRequestMessageResults`

NewDevicesSerialLiveToolsPingPostRequestMessageResultsWithDefaults instantiates a new DevicesSerialLiveToolsPingPostRequestMessageResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetReceived

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetLoss

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetLoss() DevicesSerialLiveToolsPingPostRequestMessageResultsLoss`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetLossOk() (*DevicesSerialLiveToolsPingPostRequestMessageResultsLoss, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) SetLoss(v DevicesSerialLiveToolsPingPostRequestMessageResultsLoss)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLatencies

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetLatencies() DevicesSerialLiveToolsPingPostRequestMessageResultsLatencies`

GetLatencies returns the Latencies field if non-nil, zero value otherwise.

### GetLatenciesOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetLatenciesOk() (*DevicesSerialLiveToolsPingPostRequestMessageResultsLatencies, bool)`

GetLatenciesOk returns a tuple with the Latencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencies

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) SetLatencies(v DevicesSerialLiveToolsPingPostRequestMessageResultsLatencies)`

SetLatencies sets Latencies field to given value.

### HasLatencies

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) HasLatencies() bool`

HasLatencies returns a boolean if a field has been set.

### GetReplies

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetReplies() []DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) GetRepliesOk() (*[]DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) SetReplies(v []DevicesSerialLiveToolsPingPostRequestMessageResultsRepliesInner)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *DevicesSerialLiveToolsPingPostRequestMessageResults) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


