# GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the channel utilization interval. | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the channel utilization interval. | [optional] 
**Network** | Pointer to [**GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork.md) |  | [optional] 
**ByBand** | Pointer to [**[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner.md) | Channel utilization broken down by band. | [optional] 

## Methods

### NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner

`func NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner() *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInnerWithDefaults

`func NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetByBand() []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner`

GetByBand returns the ByBand field if non-nil, zero value otherwise.

### GetByBandOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetByBandOk() (*[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner, bool)`

GetByBandOk returns a tuple with the ByBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetByBand(v []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner)`

SetByBand sets ByBand field to given value.

### HasByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasByBand() bool`

HasByBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


