# CreateOrganizationPolicyObjectsGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only) | 
**Category** | Pointer to **string** | Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup) | [optional] 
**ObjectIds** | Pointer to **[]int32** | A list of Policy Object ID&#39;s that this NetworkObjectGroup should be associated to (note: these ID&#39;s will replace the existing associated Policy Objects) | [optional] 

## Methods

### NewCreateOrganizationPolicyObjectsGroupRequest

`func NewCreateOrganizationPolicyObjectsGroupRequest(name string, ) *CreateOrganizationPolicyObjectsGroupRequest`

NewCreateOrganizationPolicyObjectsGroupRequest instantiates a new CreateOrganizationPolicyObjectsGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationPolicyObjectsGroupRequestWithDefaults

`func NewCreateOrganizationPolicyObjectsGroupRequestWithDefaults() *CreateOrganizationPolicyObjectsGroupRequest`

NewCreateOrganizationPolicyObjectsGroupRequestWithDefaults instantiates a new CreateOrganizationPolicyObjectsGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationPolicyObjectsGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationPolicyObjectsGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationPolicyObjectsGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *CreateOrganizationPolicyObjectsGroupRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateOrganizationPolicyObjectsGroupRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateOrganizationPolicyObjectsGroupRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateOrganizationPolicyObjectsGroupRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetObjectIds

`func (o *CreateOrganizationPolicyObjectsGroupRequest) GetObjectIds() []int32`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *CreateOrganizationPolicyObjectsGroupRequest) GetObjectIdsOk() (*[]int32, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *CreateOrganizationPolicyObjectsGroupRequest) SetObjectIds(v []int32)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *CreateOrganizationPolicyObjectsGroupRequest) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


