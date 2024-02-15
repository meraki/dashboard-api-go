# GetOrganizationSummaryTopNetworksByStatus200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network identifier | [optional] 
**Name** | Pointer to **string** | Network name | [optional] 
**Url** | Pointer to **string** | Network clients list URL | [optional] 
**Tags** | Pointer to **[]string** | Network tags | [optional] 
**Clients** | Pointer to [**GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients**](GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients.md) |  | [optional] 
**Statuses** | Pointer to [**GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses**](GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses.md) |  | [optional] 
**Devices** | Pointer to [**GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices**](GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices.md) |  | [optional] 
**ProductTypes** | Pointer to **[]string** | Product types in network | [optional] 

## Methods

### NewGetOrganizationSummaryTopNetworksByStatus200ResponseInner

`func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInner() *GetOrganizationSummaryTopNetworksByStatus200ResponseInner`

NewGetOrganizationSummaryTopNetworksByStatus200ResponseInner instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerWithDefaults

`func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerWithDefaults() *GetOrganizationSummaryTopNetworksByStatus200ResponseInner`

NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetClients

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetClients() GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetClientsOk() (*GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetClients(v GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetStatuses

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetStatuses() GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetStatusesOk() (*GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetStatuses(v GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetDevices

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetDevices() GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetDevicesOk() (*GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetDevices(v GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetProductTypes

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


