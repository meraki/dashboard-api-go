# GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number for the device. | [optional] 
**Mac** | Pointer to **string** | The MAC address of the device. | [optional] 
**Network** | Pointer to [**GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork.md) |  | [optional] 
**ByBand** | Pointer to [**[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner.md) | Channel utilization broken down by band. | [optional] 

## Methods

### NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner

`func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerWithDefaults

`func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetByBand() []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner`

GetByBand returns the ByBand field if non-nil, zero value otherwise.

### GetByBandOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetByBandOk() (*[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner, bool)`

GetByBandOk returns a tuple with the ByBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetByBand(v []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner)`

SetByBand sets ByBand field to given value.

### HasByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasByBand() bool`

HasByBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


