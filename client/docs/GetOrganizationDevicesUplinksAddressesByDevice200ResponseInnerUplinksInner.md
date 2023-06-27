# GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Interface for the device uplink. Available options are: cellular, man1, man2, wan1, wan2 | [optional] 
**Addresses** | Pointer to [**[]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner**](GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner.md) | Available addresses for the interface. | [optional] 

## Methods

### NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner

`func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner`

NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerWithDefaults

`func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerWithDefaults() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner`

NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetAddresses

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) GetAddresses() []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) GetAddressesOk() (*[]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) SetAddresses(v []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


