# CreateOrganizationAdminRequestTagsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | The name of the tag | 
**Access** | **string** | The privilege of the dashboard administrator on the tag. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;guest-ambassador&#39; or &#39;monitor-only&#39; | 

## Methods

### NewCreateOrganizationAdminRequestTagsInner

`func NewCreateOrganizationAdminRequestTagsInner(tag string, access string, ) *CreateOrganizationAdminRequestTagsInner`

NewCreateOrganizationAdminRequestTagsInner instantiates a new CreateOrganizationAdminRequestTagsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAdminRequestTagsInnerWithDefaults

`func NewCreateOrganizationAdminRequestTagsInnerWithDefaults() *CreateOrganizationAdminRequestTagsInner`

NewCreateOrganizationAdminRequestTagsInnerWithDefaults instantiates a new CreateOrganizationAdminRequestTagsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *CreateOrganizationAdminRequestTagsInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateOrganizationAdminRequestTagsInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateOrganizationAdminRequestTagsInner) SetTag(v string)`

SetTag sets Tag field to given value.


### GetAccess

`func (o *CreateOrganizationAdminRequestTagsInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CreateOrganizationAdminRequestTagsInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CreateOrganizationAdminRequestTagsInner) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


