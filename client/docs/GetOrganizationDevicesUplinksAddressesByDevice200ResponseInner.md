# GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The device MAC address. | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**Network** | Pointer to [**GetOrganizationDevicesAvailabilities200ResponseInnerNetwork**](GetOrganizationDevicesAvailabilities200ResponseInnerNetwork.md) |  | [optional] 
**ProductType** | Pointer to **string** | Device product type. | [optional] 
**Serial** | Pointer to **string** | The device serial number. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device. | [optional] 
**Uplinks** | Pointer to [**[]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner**](GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner.md) | List of device uplink addresses information. | [optional] 

## Methods

### NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner

`func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner`

NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerWithDefaults

`func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerWithDefaults() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner`

NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetNetwork() GetOrganizationDevicesAvailabilities200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesAvailabilities200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetNetwork(v GetOrganizationDevicesAvailabilities200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUplinks

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetUplinks() []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetUplinksOk() (*[]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetUplinks(v []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


