# GetOrganizationDevicesAvailabilities200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The device MAC address. | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**Network** | Pointer to [**GetOrganizationDevicesAvailabilities200ResponseInnerNetwork**](GetOrganizationDevicesAvailabilities200ResponseInnerNetwork.md) |  | [optional] 
**ProductType** | Pointer to **string** | Device product type. | [optional] 
**Serial** | Pointer to **string** | The device serial number. | [optional] 
**Status** | Pointer to **string** | Status of the device. Possible values are: online, alerting, offline, dormant. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device. | [optional] 

## Methods

### NewGetOrganizationDevicesAvailabilities200ResponseInner

`func NewGetOrganizationDevicesAvailabilities200ResponseInner() *GetOrganizationDevicesAvailabilities200ResponseInner`

NewGetOrganizationDevicesAvailabilities200ResponseInner instantiates a new GetOrganizationDevicesAvailabilities200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesAvailabilities200ResponseInnerWithDefaults

`func NewGetOrganizationDevicesAvailabilities200ResponseInnerWithDefaults() *GetOrganizationDevicesAvailabilities200ResponseInner`

NewGetOrganizationDevicesAvailabilities200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesAvailabilities200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetNetwork() GetOrganizationDevicesAvailabilities200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesAvailabilities200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) SetNetwork(v GetOrganizationDevicesAvailabilities200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationDevicesAvailabilities200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


