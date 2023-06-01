# InlineResponse200103

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the device | [optional] 
**Lat** | Pointer to **float32** | Latitude of the device | [optional] 
**Lng** | Pointer to **float32** | Longitude of the device | [optional] 
**Address** | Pointer to **string** | Physical address of the device | [optional] 
**Notes** | Pointer to **string** | Notes for the device, limited to 255 characters | [optional] 
**Tags** | Pointer to **[]string** | List of tags assigned to the device | [optional] 
**NetworkId** | Pointer to **string** | ID of the network the device belongs to | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**LanIp** | Pointer to **string** | LAN IP address of the device | [optional] 
**Firmware** | Pointer to **string** | Firmware version of the device | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 

## Methods

### NewInlineResponse200103

`func NewInlineResponse200103() *InlineResponse200103`

NewInlineResponse200103 instantiates a new InlineResponse200103 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200103WithDefaults

`func NewInlineResponse200103WithDefaults() *InlineResponse200103`

NewInlineResponse200103WithDefaults instantiates a new InlineResponse200103 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200103) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200103) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200103) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200103) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLat

`func (o *InlineResponse200103) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *InlineResponse200103) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *InlineResponse200103) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *InlineResponse200103) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *InlineResponse200103) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *InlineResponse200103) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *InlineResponse200103) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *InlineResponse200103) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200103) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200103) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200103) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200103) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200103) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200103) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200103) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200103) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200103) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200103) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200103) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200103) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200103) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200103) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200103) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200103) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200103) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200103) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200103) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200103) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200103) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200103) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200103) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200103) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200103) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200103) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200103) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200103) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLanIp

`func (o *InlineResponse200103) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse200103) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse200103) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse200103) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetFirmware

`func (o *InlineResponse200103) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *InlineResponse200103) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *InlineResponse200103) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *InlineResponse200103) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200103) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200103) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200103) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200103) HasProductType() bool`

HasProductType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


