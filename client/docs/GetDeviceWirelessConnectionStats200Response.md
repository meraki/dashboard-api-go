# GetDeviceWirelessConnectionStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number for the device | [optional] 
**ConnectionStats** | Pointer to [**GetDeviceWirelessConnectionStats200ResponseConnectionStats**](GetDeviceWirelessConnectionStats200ResponseConnectionStats.md) |  | [optional] 

## Methods

### NewGetDeviceWirelessConnectionStats200Response

`func NewGetDeviceWirelessConnectionStats200Response() *GetDeviceWirelessConnectionStats200Response`

NewGetDeviceWirelessConnectionStats200Response instantiates a new GetDeviceWirelessConnectionStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceWirelessConnectionStats200ResponseWithDefaults

`func NewGetDeviceWirelessConnectionStats200ResponseWithDefaults() *GetDeviceWirelessConnectionStats200Response`

NewGetDeviceWirelessConnectionStats200ResponseWithDefaults instantiates a new GetDeviceWirelessConnectionStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetDeviceWirelessConnectionStats200Response) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetDeviceWirelessConnectionStats200Response) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetDeviceWirelessConnectionStats200Response) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetDeviceWirelessConnectionStats200Response) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetConnectionStats

`func (o *GetDeviceWirelessConnectionStats200Response) GetConnectionStats() GetDeviceWirelessConnectionStats200ResponseConnectionStats`

GetConnectionStats returns the ConnectionStats field if non-nil, zero value otherwise.

### GetConnectionStatsOk

`func (o *GetDeviceWirelessConnectionStats200Response) GetConnectionStatsOk() (*GetDeviceWirelessConnectionStats200ResponseConnectionStats, bool)`

GetConnectionStatsOk returns a tuple with the ConnectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStats

`func (o *GetDeviceWirelessConnectionStats200Response) SetConnectionStats(v GetDeviceWirelessConnectionStats200ResponseConnectionStats)`

SetConnectionStats sets ConnectionStats field to given value.

### HasConnectionStats

`func (o *GetDeviceWirelessConnectionStats200Response) HasConnectionStats() bool`

HasConnectionStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


