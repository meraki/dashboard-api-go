# GetNetworkSmDevices200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Meraki Id of the device record. | [optional] 
**Name** | Pointer to **string** | The name of the device. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags associated with the device. | [optional] 
**Ssid** | Pointer to **string** | The name of the SSID the device was last connected to. | [optional] 
**WifiMac** | Pointer to **string** | The MAC of the device. | [optional] 
**OsName** | Pointer to **string** | The name of the device OS. | [optional] 
**SystemModel** | Pointer to **string** | The device model. | [optional] 
**Uuid** | Pointer to **string** | The UUID of the device. | [optional] 
**SerialNumber** | Pointer to **string** | The device serial number. | [optional] 
**Serial** | Pointer to **string** | The device serial. | [optional] 
**Ip** | Pointer to **string** | The IP address of the device. | [optional] 
**Notes** | Pointer to **string** | Notes associated with the device. | [optional] 

## Methods

### NewGetNetworkSmDevices200ResponseInner

`func NewGetNetworkSmDevices200ResponseInner() *GetNetworkSmDevices200ResponseInner`

NewGetNetworkSmDevices200ResponseInner instantiates a new GetNetworkSmDevices200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDevices200ResponseInnerWithDefaults

`func NewGetNetworkSmDevices200ResponseInnerWithDefaults() *GetNetworkSmDevices200ResponseInner`

NewGetNetworkSmDevices200ResponseInnerWithDefaults instantiates a new GetNetworkSmDevices200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkSmDevices200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSmDevices200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSmDevices200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSmDevices200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSmDevices200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSmDevices200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSmDevices200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSmDevices200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *GetNetworkSmDevices200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetNetworkSmDevices200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetNetworkSmDevices200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetNetworkSmDevices200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSsid

`func (o *GetNetworkSmDevices200ResponseInner) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *GetNetworkSmDevices200ResponseInner) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *GetNetworkSmDevices200ResponseInner) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *GetNetworkSmDevices200ResponseInner) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetWifiMac

`func (o *GetNetworkSmDevices200ResponseInner) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *GetNetworkSmDevices200ResponseInner) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *GetNetworkSmDevices200ResponseInner) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *GetNetworkSmDevices200ResponseInner) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.

### GetOsName

`func (o *GetNetworkSmDevices200ResponseInner) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *GetNetworkSmDevices200ResponseInner) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *GetNetworkSmDevices200ResponseInner) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *GetNetworkSmDevices200ResponseInner) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetSystemModel

`func (o *GetNetworkSmDevices200ResponseInner) GetSystemModel() string`

GetSystemModel returns the SystemModel field if non-nil, zero value otherwise.

### GetSystemModelOk

`func (o *GetNetworkSmDevices200ResponseInner) GetSystemModelOk() (*string, bool)`

GetSystemModelOk returns a tuple with the SystemModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModel

`func (o *GetNetworkSmDevices200ResponseInner) SetSystemModel(v string)`

SetSystemModel sets SystemModel field to given value.

### HasSystemModel

`func (o *GetNetworkSmDevices200ResponseInner) HasSystemModel() bool`

HasSystemModel returns a boolean if a field has been set.

### GetUuid

`func (o *GetNetworkSmDevices200ResponseInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetNetworkSmDevices200ResponseInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetNetworkSmDevices200ResponseInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetNetworkSmDevices200ResponseInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSerialNumber

`func (o *GetNetworkSmDevices200ResponseInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *GetNetworkSmDevices200ResponseInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *GetNetworkSmDevices200ResponseInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *GetNetworkSmDevices200ResponseInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSerial

`func (o *GetNetworkSmDevices200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetNetworkSmDevices200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetNetworkSmDevices200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetNetworkSmDevices200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetIp

`func (o *GetNetworkSmDevices200ResponseInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetNetworkSmDevices200ResponseInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetNetworkSmDevices200ResponseInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetNetworkSmDevices200ResponseInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNotes

`func (o *GetNetworkSmDevices200ResponseInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetNetworkSmDevices200ResponseInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetNetworkSmDevices200ResponseInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetNetworkSmDevices200ResponseInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


