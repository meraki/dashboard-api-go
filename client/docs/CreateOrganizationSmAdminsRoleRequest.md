# CreateOrganizationSmAdminsRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Limited Access Role | 
**Scope** | Pointer to **string** | The scope of the Limited Access Role | [optional] 
**Tags** | Pointer to **[]string** | The tags of the Limited Access Role | [optional] 

## Methods

### NewCreateOrganizationSmAdminsRoleRequest

`func NewCreateOrganizationSmAdminsRoleRequest(name string, ) *CreateOrganizationSmAdminsRoleRequest`

NewCreateOrganizationSmAdminsRoleRequest instantiates a new CreateOrganizationSmAdminsRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationSmAdminsRoleRequestWithDefaults

`func NewCreateOrganizationSmAdminsRoleRequestWithDefaults() *CreateOrganizationSmAdminsRoleRequest`

NewCreateOrganizationSmAdminsRoleRequestWithDefaults instantiates a new CreateOrganizationSmAdminsRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationSmAdminsRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationSmAdminsRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationSmAdminsRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *CreateOrganizationSmAdminsRoleRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateOrganizationSmAdminsRoleRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateOrganizationSmAdminsRoleRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateOrganizationSmAdminsRoleRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTags

`func (o *CreateOrganizationSmAdminsRoleRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOrganizationSmAdminsRoleRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOrganizationSmAdminsRoleRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateOrganizationSmAdminsRoleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


