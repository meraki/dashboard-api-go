# UpdateNetworkSmDevicesFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMac** | Pointer to **string** | The wifiMac of the device to be modified. | [optional] 
**Id** | Pointer to **string** | The id of the device to be modified. | [optional] 
**Serial** | Pointer to **string** | The serial of the device to be modified. | [optional] 
**DeviceFields** | [**UpdateNetworkSmDevicesFieldsRequestDeviceFields**](UpdateNetworkSmDevicesFieldsRequestDeviceFields.md) |  | 

## Methods

### NewUpdateNetworkSmDevicesFieldsRequest

`func NewUpdateNetworkSmDevicesFieldsRequest(deviceFields UpdateNetworkSmDevicesFieldsRequestDeviceFields, ) *UpdateNetworkSmDevicesFieldsRequest`

NewUpdateNetworkSmDevicesFieldsRequest instantiates a new UpdateNetworkSmDevicesFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSmDevicesFieldsRequestWithDefaults

`func NewUpdateNetworkSmDevicesFieldsRequestWithDefaults() *UpdateNetworkSmDevicesFieldsRequest`

NewUpdateNetworkSmDevicesFieldsRequestWithDefaults instantiates a new UpdateNetworkSmDevicesFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMac

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *UpdateNetworkSmDevicesFieldsRequest) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *UpdateNetworkSmDevicesFieldsRequest) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.

### GetId

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkSmDevicesFieldsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateNetworkSmDevicesFieldsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerial

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *UpdateNetworkSmDevicesFieldsRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *UpdateNetworkSmDevicesFieldsRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetDeviceFields

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetDeviceFields() UpdateNetworkSmDevicesFieldsRequestDeviceFields`

GetDeviceFields returns the DeviceFields field if non-nil, zero value otherwise.

### GetDeviceFieldsOk

`func (o *UpdateNetworkSmDevicesFieldsRequest) GetDeviceFieldsOk() (*UpdateNetworkSmDevicesFieldsRequestDeviceFields, bool)`

GetDeviceFieldsOk returns a tuple with the DeviceFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFields

`func (o *UpdateNetworkSmDevicesFieldsRequest) SetDeviceFields(v UpdateNetworkSmDevicesFieldsRequestDeviceFields)`

SetDeviceFields sets DeviceFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


