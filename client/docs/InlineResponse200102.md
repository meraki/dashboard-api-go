# InlineResponse200102

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Serial** | Pointer to **string** | Serial of MX device | [optional] 
**Uplink** | Pointer to **string** | Uplink interface (wan1, wan2, or cellular) | [optional] 
**Ip** | Pointer to **string** | IP address of uplink | [optional] 
**TimeSeries** | Pointer to [**[]OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries**](OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries.md) | Loss and latency timeseries data | [optional] 

## Methods

### NewInlineResponse200102

`func NewInlineResponse200102() *InlineResponse200102`

NewInlineResponse200102 instantiates a new InlineResponse200102 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102WithDefaults

`func NewInlineResponse200102WithDefaults() *InlineResponse200102`

NewInlineResponse200102WithDefaults instantiates a new InlineResponse200102 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200102) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200102) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200102) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200102) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200102) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200102) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200102) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200102) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200102) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200102) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200102) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200102) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse200102) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse200102) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse200102) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse200102) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetTimeSeries

`func (o *InlineResponse200102) GetTimeSeries() []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries`

GetTimeSeries returns the TimeSeries field if non-nil, zero value otherwise.

### GetTimeSeriesOk

`func (o *InlineResponse200102) GetTimeSeriesOk() (*[]OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries, bool)`

GetTimeSeriesOk returns a tuple with the TimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeries

`func (o *InlineResponse200102) SetTimeSeries(v []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries)`

SetTimeSeries sets TimeSeries field to given value.

### HasTimeSeries

`func (o *InlineResponse200102) HasTimeSeries() bool`

HasTimeSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


