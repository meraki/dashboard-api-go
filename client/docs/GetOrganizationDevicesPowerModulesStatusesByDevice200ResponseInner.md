# GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The device MAC address. | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**Network** | Pointer to [**GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork**](GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork.md) |  | [optional] 
**ProductType** | Pointer to **string** | Device product type. | [optional] 
**Serial** | Pointer to **string** | The device serial number. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device. | [optional] 
**Slots** | Pointer to [**[]GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner**](GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner.md) | Information for the device&#39;s AC power supplies. | [optional] 

## Methods

### NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner

`func NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner`

NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner instantiates a new GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerWithDefaults

`func NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerWithDefaults() *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner`

NewGetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetNetwork() GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) SetNetwork(v GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSlots

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetSlots() []GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) GetSlotsOk() (*[]GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) SetSlots(v []GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInnerSlotsInner)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner) HasSlots() bool`

HasSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


