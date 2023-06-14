# GetNetworkWirelessUsageHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**TotalKbps** | Pointer to **int32** | Total usage in kilobytes-per-second | [optional] 
**SentKbps** | Pointer to **int32** | Sent kilobytes-per-second | [optional] 
**ReceivedKbps** | Pointer to **int32** | Received kilobytes-per-second | [optional] 

## Methods

### NewGetNetworkWirelessUsageHistory200ResponseInner

`func NewGetNetworkWirelessUsageHistory200ResponseInner() *GetNetworkWirelessUsageHistory200ResponseInner`

NewGetNetworkWirelessUsageHistory200ResponseInner instantiates a new GetNetworkWirelessUsageHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessUsageHistory200ResponseInnerWithDefaults

`func NewGetNetworkWirelessUsageHistory200ResponseInnerWithDefaults() *GetNetworkWirelessUsageHistory200ResponseInner`

NewGetNetworkWirelessUsageHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessUsageHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetTotalKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetTotalKbps() int32`

GetTotalKbps returns the TotalKbps field if non-nil, zero value otherwise.

### GetTotalKbpsOk

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetTotalKbpsOk() (*int32, bool)`

GetTotalKbpsOk returns a tuple with the TotalKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetTotalKbps(v int32)`

SetTotalKbps sets TotalKbps field to given value.

### HasTotalKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasTotalKbps() bool`

HasTotalKbps returns a boolean if a field has been set.

### GetSentKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetSentKbps() int32`

GetSentKbps returns the SentKbps field if non-nil, zero value otherwise.

### GetSentKbpsOk

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetSentKbpsOk() (*int32, bool)`

GetSentKbpsOk returns a tuple with the SentKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetSentKbps(v int32)`

SetSentKbps sets SentKbps field to given value.

### HasSentKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasSentKbps() bool`

HasSentKbps returns a boolean if a field has been set.

### GetReceivedKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetReceivedKbps() int32`

GetReceivedKbps returns the ReceivedKbps field if non-nil, zero value otherwise.

### GetReceivedKbpsOk

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) GetReceivedKbpsOk() (*int32, bool)`

GetReceivedKbpsOk returns a tuple with the ReceivedKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) SetReceivedKbps(v int32)`

SetReceivedKbps sets ReceivedKbps field to given value.

### HasReceivedKbps

`func (o *GetNetworkWirelessUsageHistory200ResponseInner) HasReceivedKbps() bool`

HasReceivedKbps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


