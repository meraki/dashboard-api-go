# GetOrganizationDevicesUplinksLossAndLatency200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Serial** | Pointer to **string** | Serial of MX device | [optional] 
**Uplink** | Pointer to **string** | Uplink interface (wan1, wan2, or cellular) | [optional] 
**Ip** | Pointer to **string** | IP address of uplink | [optional] 
**TimeSeries** | Pointer to [**[]GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner**](GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner.md) | Loss and latency timeseries data | [optional] 

## Methods

### NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInner

`func NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInner() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner`

NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInner instantiates a new GetOrganizationDevicesUplinksLossAndLatency200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerWithDefaults

`func NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerWithDefaults() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner`

NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksLossAndLatency200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUplink

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetIp

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetTimeSeries

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetTimeSeries() []GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner`

GetTimeSeries returns the TimeSeries field if non-nil, zero value otherwise.

### GetTimeSeriesOk

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetTimeSeriesOk() (*[]GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner, bool)`

GetTimeSeriesOk returns a tuple with the TimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeries

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetTimeSeries(v []GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner)`

SetTimeSeries sets TimeSeries field to given value.

### HasTimeSeries

`func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasTimeSeries() bool`

HasTimeSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


