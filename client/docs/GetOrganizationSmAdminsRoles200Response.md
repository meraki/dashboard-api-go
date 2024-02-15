# GetOrganizationSmAdminsRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]GetOrganizationSmAdminsRoles200ResponseItemsInner**](GetOrganizationSmAdminsRoles200ResponseItemsInner.md) | Array of Limited Access Roles | [optional] 
**Meta** | Pointer to [**GetOrganizationSmAdminsRoles200ResponseMeta**](GetOrganizationSmAdminsRoles200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetOrganizationSmAdminsRoles200Response

`func NewGetOrganizationSmAdminsRoles200Response() *GetOrganizationSmAdminsRoles200Response`

NewGetOrganizationSmAdminsRoles200Response instantiates a new GetOrganizationSmAdminsRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSmAdminsRoles200ResponseWithDefaults

`func NewGetOrganizationSmAdminsRoles200ResponseWithDefaults() *GetOrganizationSmAdminsRoles200Response`

NewGetOrganizationSmAdminsRoles200ResponseWithDefaults instantiates a new GetOrganizationSmAdminsRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetOrganizationSmAdminsRoles200Response) GetItems() []GetOrganizationSmAdminsRoles200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetOrganizationSmAdminsRoles200Response) GetItemsOk() (*[]GetOrganizationSmAdminsRoles200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetOrganizationSmAdminsRoles200Response) SetItems(v []GetOrganizationSmAdminsRoles200ResponseItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetOrganizationSmAdminsRoles200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *GetOrganizationSmAdminsRoles200Response) GetMeta() GetOrganizationSmAdminsRoles200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetOrganizationSmAdminsRoles200Response) GetMetaOk() (*GetOrganizationSmAdminsRoles200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetOrganizationSmAdminsRoles200Response) SetMeta(v GetOrganizationSmAdminsRoles200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetOrganizationSmAdminsRoles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


