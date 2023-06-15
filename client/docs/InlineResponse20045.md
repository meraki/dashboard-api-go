# InlineResponse20045

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

### NewInlineResponse20045

`func NewInlineResponse20045() *InlineResponse20045`

NewInlineResponse20045 instantiates a new InlineResponse20045 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20045WithDefaults

`func NewInlineResponse20045WithDefaults() *InlineResponse20045`

NewInlineResponse20045WithDefaults instantiates a new InlineResponse20045 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20045) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20045) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20045) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20045) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20045) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20045) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20045) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20045) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20045) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20045) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20045) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20045) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse20045) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse20045) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse20045) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse20045) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetWifiMac

`func (o *InlineResponse20045) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *InlineResponse20045) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *InlineResponse20045) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *InlineResponse20045) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.

### GetOsName

`func (o *InlineResponse20045) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *InlineResponse20045) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *InlineResponse20045) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *InlineResponse20045) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetSystemModel

`func (o *InlineResponse20045) GetSystemModel() string`

GetSystemModel returns the SystemModel field if non-nil, zero value otherwise.

### GetSystemModelOk

`func (o *InlineResponse20045) GetSystemModelOk() (*string, bool)`

GetSystemModelOk returns a tuple with the SystemModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModel

`func (o *InlineResponse20045) SetSystemModel(v string)`

SetSystemModel sets SystemModel field to given value.

### HasSystemModel

`func (o *InlineResponse20045) HasSystemModel() bool`

HasSystemModel returns a boolean if a field has been set.

### GetUuid

`func (o *InlineResponse20045) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InlineResponse20045) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InlineResponse20045) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InlineResponse20045) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSerialNumber

`func (o *InlineResponse20045) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *InlineResponse20045) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *InlineResponse20045) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *InlineResponse20045) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse20045) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20045) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20045) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20045) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse20045) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20045) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20045) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20045) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse20045) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse20045) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse20045) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse20045) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


