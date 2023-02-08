# GetDeviceWirelessBluetoothSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Desired UUID of the beacon. If the value is set to null it will reset to Dashboard&#39;s automatically generated value. | [optional] 
**Major** | Pointer to **int32** | Desired major value of the beacon. If the value is set to null it will reset to Dashboard&#39;s automatically generated value. | [optional] 
**Minor** | Pointer to **int32** | Desired minor value of the beacon. If the value is set to null it will reset to Dashboard&#39;s automatically generated value. | [optional] 

## Methods

### NewGetDeviceWirelessBluetoothSettings200Response

`func NewGetDeviceWirelessBluetoothSettings200Response() *GetDeviceWirelessBluetoothSettings200Response`

NewGetDeviceWirelessBluetoothSettings200Response instantiates a new GetDeviceWirelessBluetoothSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceWirelessBluetoothSettings200ResponseWithDefaults

`func NewGetDeviceWirelessBluetoothSettings200ResponseWithDefaults() *GetDeviceWirelessBluetoothSettings200Response`

NewGetDeviceWirelessBluetoothSettings200ResponseWithDefaults instantiates a new GetDeviceWirelessBluetoothSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *GetDeviceWirelessBluetoothSettings200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDeviceWirelessBluetoothSettings200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDeviceWirelessBluetoothSettings200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDeviceWirelessBluetoothSettings200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMajor

`func (o *GetDeviceWirelessBluetoothSettings200Response) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *GetDeviceWirelessBluetoothSettings200Response) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *GetDeviceWirelessBluetoothSettings200Response) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *GetDeviceWirelessBluetoothSettings200Response) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *GetDeviceWirelessBluetoothSettings200Response) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *GetDeviceWirelessBluetoothSettings200Response) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *GetDeviceWirelessBluetoothSettings200Response) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *GetDeviceWirelessBluetoothSettings200Response) HasMinor() bool`

HasMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


