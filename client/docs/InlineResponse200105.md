# InlineResponse200105

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The device MAC address. | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork**](OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork.md) |  | [optional] 
**ProductType** | Pointer to **string** | Device product type. | [optional] 
**Serial** | Pointer to **string** | The device serial number. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device. | [optional] 
**Slots** | Pointer to [**[]OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots**](OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots.md) | Information for the device&#39;s AC power supplies. | [optional] 

## Methods

### NewInlineResponse200105

`func NewInlineResponse200105() *InlineResponse200105`

NewInlineResponse200105 instantiates a new InlineResponse200105 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200105WithDefaults

`func NewInlineResponse200105WithDefaults() *InlineResponse200105`

NewInlineResponse200105WithDefaults instantiates a new InlineResponse200105 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse200105) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200105) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200105) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200105) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200105) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200105) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200105) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200105) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200105) GetNetwork() OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200105) GetNetworkOk() (*OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200105) SetNetwork(v OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200105) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200105) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200105) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200105) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200105) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200105) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200105) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200105) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200105) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200105) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200105) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200105) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200105) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSlots

`func (o *InlineResponse200105) GetSlots() []OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *InlineResponse200105) GetSlotsOk() (*[]OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *InlineResponse200105) SetSlots(v []OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *InlineResponse200105) HasSlots() bool`

HasSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


