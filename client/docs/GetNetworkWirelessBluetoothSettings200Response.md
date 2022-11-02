# GetNetworkWirelessBluetoothSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanningEnabled** | Pointer to **bool** | Whether APs will scan for Bluetooth enabled clients. | [optional] 
**AdvertisingEnabled** | Pointer to **bool** | Whether APs will advertise beacons. | [optional] 
**Uuid** | Pointer to **string** | The UUID to be used in the beacon identifier. | [optional] 
**MajorMinorAssignmentMode** | Pointer to **string** | The way major and minor number should be assigned to nodes in the network. (&#39;Unique&#39;, &#39;Non-unique&#39;) | [optional] 
**Major** | Pointer to **int32** | The major number to be used in the beacon identifier. Only valid in &#39;Non-unique&#39; mode. | [optional] 
**Minor** | Pointer to **int32** | The minor number to be used in the beacon identifier. Only valid in &#39;Non-unique&#39; mode. | [optional] 

## Methods

### NewGetNetworkWirelessBluetoothSettings200Response

`func NewGetNetworkWirelessBluetoothSettings200Response() *GetNetworkWirelessBluetoothSettings200Response`

NewGetNetworkWirelessBluetoothSettings200Response instantiates a new GetNetworkWirelessBluetoothSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessBluetoothSettings200ResponseWithDefaults

`func NewGetNetworkWirelessBluetoothSettings200ResponseWithDefaults() *GetNetworkWirelessBluetoothSettings200Response`

NewGetNetworkWirelessBluetoothSettings200ResponseWithDefaults instantiates a new GetNetworkWirelessBluetoothSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanningEnabled

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetScanningEnabled() bool`

GetScanningEnabled returns the ScanningEnabled field if non-nil, zero value otherwise.

### GetScanningEnabledOk

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetScanningEnabledOk() (*bool, bool)`

GetScanningEnabledOk returns a tuple with the ScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanningEnabled

`func (o *GetNetworkWirelessBluetoothSettings200Response) SetScanningEnabled(v bool)`

SetScanningEnabled sets ScanningEnabled field to given value.

### HasScanningEnabled

`func (o *GetNetworkWirelessBluetoothSettings200Response) HasScanningEnabled() bool`

HasScanningEnabled returns a boolean if a field has been set.

### GetAdvertisingEnabled

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetAdvertisingEnabled() bool`

GetAdvertisingEnabled returns the AdvertisingEnabled field if non-nil, zero value otherwise.

### GetAdvertisingEnabledOk

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetAdvertisingEnabledOk() (*bool, bool)`

GetAdvertisingEnabledOk returns a tuple with the AdvertisingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisingEnabled

`func (o *GetNetworkWirelessBluetoothSettings200Response) SetAdvertisingEnabled(v bool)`

SetAdvertisingEnabled sets AdvertisingEnabled field to given value.

### HasAdvertisingEnabled

`func (o *GetNetworkWirelessBluetoothSettings200Response) HasAdvertisingEnabled() bool`

HasAdvertisingEnabled returns a boolean if a field has been set.

### GetUuid

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetNetworkWirelessBluetoothSettings200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetNetworkWirelessBluetoothSettings200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMajorMinorAssignmentMode

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajorMinorAssignmentMode() string`

GetMajorMinorAssignmentMode returns the MajorMinorAssignmentMode field if non-nil, zero value otherwise.

### GetMajorMinorAssignmentModeOk

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajorMinorAssignmentModeOk() (*string, bool)`

GetMajorMinorAssignmentModeOk returns a tuple with the MajorMinorAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorMinorAssignmentMode

`func (o *GetNetworkWirelessBluetoothSettings200Response) SetMajorMinorAssignmentMode(v string)`

SetMajorMinorAssignmentMode sets MajorMinorAssignmentMode field to given value.

### HasMajorMinorAssignmentMode

`func (o *GetNetworkWirelessBluetoothSettings200Response) HasMajorMinorAssignmentMode() bool`

HasMajorMinorAssignmentMode returns a boolean if a field has been set.

### GetMajor

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *GetNetworkWirelessBluetoothSettings200Response) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *GetNetworkWirelessBluetoothSettings200Response) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *GetNetworkWirelessBluetoothSettings200Response) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *GetNetworkWirelessBluetoothSettings200Response) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *GetNetworkWirelessBluetoothSettings200Response) HasMinor() bool`

HasMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


