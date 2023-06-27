# GetOrganizationDevicesProvisioningStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The device MAC address. | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**Network** | Pointer to [**GetOrganizationDevicesAvailabilities200ResponseInnerNetwork**](GetOrganizationDevicesAvailabilities200ResponseInnerNetwork.md) |  | [optional] 
**ProductType** | Pointer to **string** | Device product type. | [optional] 
**Serial** | Pointer to **string** | The device serial number. | [optional] 
**Status** | Pointer to **string** | The device provisioning status. Possible statuses: unprovisioned, incomplete, complete. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device. | [optional] 

## Methods

### NewGetOrganizationDevicesProvisioningStatuses200ResponseInner

`func NewGetOrganizationDevicesProvisioningStatuses200ResponseInner() *GetOrganizationDevicesProvisioningStatuses200ResponseInner`

NewGetOrganizationDevicesProvisioningStatuses200ResponseInner instantiates a new GetOrganizationDevicesProvisioningStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesProvisioningStatuses200ResponseInnerWithDefaults

`func NewGetOrganizationDevicesProvisioningStatuses200ResponseInnerWithDefaults() *GetOrganizationDevicesProvisioningStatuses200ResponseInner`

NewGetOrganizationDevicesProvisioningStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesProvisioningStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetNetwork() GetOrganizationDevicesAvailabilities200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesAvailabilities200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetNetwork(v GetOrganizationDevicesAvailabilities200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


