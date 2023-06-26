# GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork.md) |  | [optional] 
**ByBand** | Pointer to [**[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner.md) | Channel utilization broken down by band. | [optional] 

## Methods

### NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner

`func NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner() *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInnerWithDefaults

`func NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner`

NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetByBand() []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner`

GetByBand returns the ByBand field if non-nil, zero value otherwise.

### GetByBandOk

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetByBandOk() (*[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner, bool)`

GetByBandOk returns a tuple with the ByBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) SetByBand(v []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner)`

SetByBand sets ByBand field to given value.

### HasByBand

`func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) HasByBand() bool`

HasByBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


