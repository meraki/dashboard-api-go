# GetNetworkWirelessSignalQualityHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**Snr** | Pointer to **int32** | Signal to noise ratio | [optional] 
**Rssi** | Pointer to **int32** | Received signal strength indicator | [optional] 

## Methods

### NewGetNetworkWirelessSignalQualityHistory200ResponseInner

`func NewGetNetworkWirelessSignalQualityHistory200ResponseInner() *GetNetworkWirelessSignalQualityHistory200ResponseInner`

NewGetNetworkWirelessSignalQualityHistory200ResponseInner instantiates a new GetNetworkWirelessSignalQualityHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSignalQualityHistory200ResponseInnerWithDefaults

`func NewGetNetworkWirelessSignalQualityHistory200ResponseInnerWithDefaults() *GetNetworkWirelessSignalQualityHistory200ResponseInner`

NewGetNetworkWirelessSignalQualityHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessSignalQualityHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetSnr

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetRssi

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasRssi() bool`

HasRssi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


