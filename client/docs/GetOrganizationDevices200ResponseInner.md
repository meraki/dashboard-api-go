# GetOrganizationDevices200ResponseInner

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

### NewGetOrganizationDevices200ResponseInner

`func NewGetOrganizationDevices200ResponseInner() *GetOrganizationDevices200ResponseInner`

NewGetOrganizationDevices200ResponseInner instantiates a new GetOrganizationDevices200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevices200ResponseInnerWithDefaults

`func NewGetOrganizationDevices200ResponseInnerWithDefaults() *GetOrganizationDevices200ResponseInner`

NewGetOrganizationDevices200ResponseInnerWithDefaults instantiates a new GetOrganizationDevices200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetOrganizationDevices200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationDevices200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationDevices200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationDevices200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLat

`func (o *GetOrganizationDevices200ResponseInner) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *GetOrganizationDevices200ResponseInner) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *GetOrganizationDevices200ResponseInner) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *GetOrganizationDevices200ResponseInner) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *GetOrganizationDevices200ResponseInner) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *GetOrganizationDevices200ResponseInner) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *GetOrganizationDevices200ResponseInner) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *GetOrganizationDevices200ResponseInner) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *GetOrganizationDevices200ResponseInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetOrganizationDevices200ResponseInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetOrganizationDevices200ResponseInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetOrganizationDevices200ResponseInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *GetOrganizationDevices200ResponseInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetOrganizationDevices200ResponseInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetOrganizationDevices200ResponseInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetOrganizationDevices200ResponseInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationDevices200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationDevices200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationDevices200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationDevices200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetOrganizationDevices200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationDevices200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationDevices200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationDevices200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevices200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevices200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevices200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevices200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationDevices200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationDevices200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationDevices200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationDevices200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationDevices200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationDevices200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationDevices200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationDevices200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLanIp

`func (o *GetOrganizationDevices200ResponseInner) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *GetOrganizationDevices200ResponseInner) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *GetOrganizationDevices200ResponseInner) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *GetOrganizationDevices200ResponseInner) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetFirmware

`func (o *GetOrganizationDevices200ResponseInner) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *GetOrganizationDevices200ResponseInner) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *GetOrganizationDevices200ResponseInner) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *GetOrganizationDevices200ResponseInner) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationDevices200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationDevices200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationDevices200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationDevices200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


