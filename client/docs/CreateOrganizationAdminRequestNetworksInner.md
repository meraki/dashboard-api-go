# CreateOrganizationAdminRequestNetworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The network ID | 
**Access** | **string** | The privilege of the dashboard administrator on the network. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;guest-ambassador&#39; or &#39;monitor-only&#39; | 

## Methods

### NewCreateOrganizationAdminRequestNetworksInner

`func NewCreateOrganizationAdminRequestNetworksInner(id string, access string, ) *CreateOrganizationAdminRequestNetworksInner`

NewCreateOrganizationAdminRequestNetworksInner instantiates a new CreateOrganizationAdminRequestNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAdminRequestNetworksInnerWithDefaults

`func NewCreateOrganizationAdminRequestNetworksInnerWithDefaults() *CreateOrganizationAdminRequestNetworksInner`

NewCreateOrganizationAdminRequestNetworksInnerWithDefaults instantiates a new CreateOrganizationAdminRequestNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrganizationAdminRequestNetworksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrganizationAdminRequestNetworksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrganizationAdminRequestNetworksInner) SetId(v string)`

SetId sets Id field to given value.


### GetAccess

`func (o *CreateOrganizationAdminRequestNetworksInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CreateOrganizationAdminRequestNetworksInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CreateOrganizationAdminRequestNetworksInner) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


