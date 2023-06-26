# GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the channel utilization interval. | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the channel utilization interval. | [optional] 
**Serial** | Pointer to **string** | The serial number for the device. | [optional] 
**Mac** | Pointer to **string** | The MAC address of the device. | [optional] 
**Network** | Pointer to [**GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork.md) |  | [optional] 
**ByBand** | Pointer to [**[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner.md) | Channel utilization broken down by band. | [optional] 

## Methods

### NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner

`func NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner() *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInnerWithDefaults

`func NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetByBand() []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner`

GetByBand returns the ByBand field if non-nil, zero value otherwise.

### GetByBandOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) GetByBandOk() (*[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner, bool)`

GetByBandOk returns a tuple with the ByBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) SetByBand(v []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner)`

SetByBand sets ByBand field to given value.

### HasByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner) HasByBand() bool`

HasByBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


