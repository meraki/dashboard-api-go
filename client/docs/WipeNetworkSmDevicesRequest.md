# WipeNetworkSmDevicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMac** | Pointer to **string** | The wifiMac of the device to be wiped. | [optional] 
**Id** | Pointer to **string** | The id of the device to be wiped. | [optional] 
**Serial** | Pointer to **string** | The serial of the device to be wiped. | [optional] 
**Pin** | Pointer to **int32** | The pin number (a six digit value) for wiping a macOS device. Required only for macOS devices. | [optional] 

## Methods

### NewWipeNetworkSmDevicesRequest

`func NewWipeNetworkSmDevicesRequest() *WipeNetworkSmDevicesRequest`

NewWipeNetworkSmDevicesRequest instantiates a new WipeNetworkSmDevicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWipeNetworkSmDevicesRequestWithDefaults

`func NewWipeNetworkSmDevicesRequestWithDefaults() *WipeNetworkSmDevicesRequest`

NewWipeNetworkSmDevicesRequestWithDefaults instantiates a new WipeNetworkSmDevicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMac

`func (o *WipeNetworkSmDevicesRequest) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *WipeNetworkSmDevicesRequest) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *WipeNetworkSmDevicesRequest) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *WipeNetworkSmDevicesRequest) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.

### GetId

`func (o *WipeNetworkSmDevicesRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WipeNetworkSmDevicesRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WipeNetworkSmDevicesRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WipeNetworkSmDevicesRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerial

`func (o *WipeNetworkSmDevicesRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *WipeNetworkSmDevicesRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *WipeNetworkSmDevicesRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *WipeNetworkSmDevicesRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetPin

`func (o *WipeNetworkSmDevicesRequest) GetPin() int32`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *WipeNetworkSmDevicesRequest) GetPinOk() (*int32, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *WipeNetworkSmDevicesRequest) SetPin(v int32)`

SetPin sets Pin field to given value.

### HasPin

`func (o *WipeNetworkSmDevicesRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


