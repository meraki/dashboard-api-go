# CreateOrganizationAdaptivePolicyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the group | 
**Sgt** | **int32** | SGT value of the group | 
**Description** | Pointer to **string** | Description of the group (default: \&quot;\&quot;) | [optional] 
**PolicyObjects** | Pointer to [**[]CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner**](CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner.md) | The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group&#39;s SGT value if no other tagging scheme is being used (each requires one unique attribute) (default: []) | [optional] 

## Methods

### NewCreateOrganizationAdaptivePolicyGroupRequest

`func NewCreateOrganizationAdaptivePolicyGroupRequest(name string, sgt int32, ) *CreateOrganizationAdaptivePolicyGroupRequest`

NewCreateOrganizationAdaptivePolicyGroupRequest instantiates a new CreateOrganizationAdaptivePolicyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAdaptivePolicyGroupRequestWithDefaults

`func NewCreateOrganizationAdaptivePolicyGroupRequestWithDefaults() *CreateOrganizationAdaptivePolicyGroupRequest`

NewCreateOrganizationAdaptivePolicyGroupRequestWithDefaults instantiates a new CreateOrganizationAdaptivePolicyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSgt

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetSgt(v int32)`

SetSgt sets Sgt field to given value.


### GetDescription

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjects() []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjectsOk() (*[]CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetPolicyObjects(v []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *CreateOrganizationAdaptivePolicyGroupRequest) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


